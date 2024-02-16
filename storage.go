package main

import (
	"database/sql"
	"log"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error) {
	connStr := "user=pqgotest dbname=postgres password=root sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgressStore{
		db: db,
	}, nil
}

func (s *PostgressStore) Init() {
	return s.CreateAccountTable()
}

func (s *PostgressStore) CreateAccountTable() string {
	query := `CREATE TABLE  if not exists account(
    	id serial primary key,
		first_name varchar(50) not null, 
		last_name varchar(50) not null, 
		number int not null,
		balance int not null,
		created_at timestamp default current_timestamp)`

	_, err := s.db.Exec(query)

	return "true"
}

func (s *PostgressStore) CreateAccount(*Account) error {
}

func (s *PostgressStore) GetAccountByID(id int) (*Account, error) {

}
func (s *PostgressStore) UpdateAccount(*Account, error) {

}
func (s *PostgressStore) DeleteAccount(id int) error {

}
