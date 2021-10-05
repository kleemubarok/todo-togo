package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-togo/entity"
)

type IStatusRepo interface {
	SelectStatus(status entity.Status) (entity.Status, error)
	SelectAllStatus() ([]entity.Status, error)
}

func NewStatusRepo(dbParam *sqlx.DB) IStatusRepo {
	return &SRepo{
		db: dbParam,
	}
}

func (t *SRepo) SelectStatus(s entity.Status) (entity.Status, error) {
	res := entity.Status{}
	err := t.db.Get(&res, "SELECT * FROM status WHERE id=$1",s.StatusID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (t *SRepo) SelectAllStatus() ([]entity.Status, error) {
	var res []entity.Status
	err := t.db.Select(&res,"SELECT * FROM status")
	if err != nil {
		return nil, err
	}

	return res, nil
}