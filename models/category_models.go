package models

import "database/sql"

type CategoryResponse struct {
	ID				int				`json:"id"`
	CategoryName	sql.NullString			`json:"category_name"`
}

type CategoryRequest struct {
	ID				int				`json:"id"`
	CategoryName 	string	`json:"category_name"`
}