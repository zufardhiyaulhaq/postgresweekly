# postgresweekly

Get data from postgres news and create Weekly CRDs based on community-operator and push to git datastore

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square) [![made with Go](https://img.shields.io/badge/made%20with-Go-brightgreen)](http://golang.org) [![Github master branch build](https://img.shields.io/github/workflow/status/zufardhiyaulhaq/postgresweekly/Master)](https://github.com/zufardhiyaulhaq/postgresweekly/actions/workflows/master.yml) [![GitHub issues](https://img.shields.io/github/issues/zufardhiyaulhaq/postgresweekly)](https://github.com/zufardhiyaulhaq/postgresweekly/issues) [![GitHub pull requests](https://img.shields.io/github/issues-pr/zufardhiyaulhaq/postgresweekly)](https://github.com/zufardhiyaulhaq/postgresweekly/pulls)

## Installing the Chart

To install the chart with the release name `my-postgresweekly` and secret `foo`:

```console
kubectl apply -f secret.yaml

helm repo add postgresweekly https://zufardhiyaulhaq.com/postgresweekly/charts/releases/
helm install my-postgresweekly postgresweekly/postgresweekly --values values.yaml
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| cronSchedule | string | `"0 8 * * 0"` |  |
| image.repository | string | `"zufardhiyaulhaq/postgresweekly"` |  |
| image.tag | string | `"v1.0.0"` |  |
| secret | string | `""` |  |
| weekly.community | string | `"postgres"` |  |
| weekly.image_url | string | `"https://kinsta.com/wp-content/uploads/2022/02/postgres-logo.png"` |  |
| weekly.namespace | string | `"community"` |  |
| weekly.tags | string | `"weekly,postgres"` |  |

