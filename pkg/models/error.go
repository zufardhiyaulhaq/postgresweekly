package models

type PostgresweeklyNameNotFoundError struct{}

func (k *PostgresweeklyNameNotFoundError) Error() string {
	return "Postgresweekly name not found"
}
