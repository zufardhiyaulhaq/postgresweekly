package postgresweekly

import (
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	communityv1alpha1 "github.com/zufardhiyaulhaq/community-operator-v2/api/v1alpha1"
	"github.com/zufardhiyaulhaq/postgresweekly/pkg/scrappers"
)

type PostgresWeekly struct {
	URL string
}

func NewPostgresWeekly(URL string) scrappers.Scrapper {
	return PostgresWeekly{
		URL: URL,
	}
}

func (g PostgresWeekly) GetName() (string, error) {
	response, err := http.Get(g.URL)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}

	text := document.Find("div#content").ChildrenFiltered("table:first-of-type").First().Find("td:first-of-type").First().Find("p").Text()
	regex, err := regexp.Compile(`\d+`)
	if err != nil {
		return "", err
	}

	name := "Postgres Weekly " + regex.FindString(text)

	return name, nil
}

func (g PostgresWeekly) GetArticles() ([]communityv1alpha1.WeeklySpec_Article, error) {
	var articles []communityv1alpha1.WeeklySpec_Article

	response, err := http.Get(g.URL)
	if err != nil {
		return articles, err
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return articles, err
	}

	document.Find("div#content").Each(func(i int, div *goquery.Selection) {
		div.Find("span.mainlink").Each(func(i int, span *goquery.Selection) {
			var article communityv1alpha1.WeeklySpec_Article
			article.Title = span.Find("a").Text()
			article.Type = "News"
			article.Url, _ = span.Find("a").Attr("href")

			articles = append(articles, article)
		})
	})

	return articles, nil
}
