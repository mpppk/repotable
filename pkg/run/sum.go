// Package run provides utilities for calculate run of numbers
package run

type RepoTableConfig struct {
	Repositories []Repository
}

type Repository struct {
	Owner string
	Repo  string
}
