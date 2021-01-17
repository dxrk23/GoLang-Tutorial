package postgres

import (
	"dimash/snippetbox/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SnippetModel struct{
	Pool *pgxpool.Pool
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}

