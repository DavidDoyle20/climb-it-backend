// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: habits.sql

package database

import (
	"context"
	"time"
)

const createHabitForUser = `-- name: CreateHabitForUser :one
INSERT INTO habits (id, created_at, updated_at, name, user_id, start_date, end_date)
VALUES (
    ?,
    datetime('now'),
    datetime('now'),
    ?,
    ?,
    ?,
    ?
)
RETURNING id, created_at, updated_at, name, user_id, start_date, end_date
`

type CreateHabitForUserParams struct {
	ID        string
	Name      string
	UserID    string
	StartDate time.Time
	EndDate   time.Time
}

func (q *Queries) CreateHabitForUser(ctx context.Context, arg CreateHabitForUserParams) (Habit, error) {
	row := q.db.QueryRowContext(ctx, createHabitForUser,
		arg.ID,
		arg.Name,
		arg.UserID,
		arg.StartDate,
		arg.EndDate,
	)
	var i Habit
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.UserID,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const getHabit = `-- name: GetHabit :one
SELECT id, created_at, updated_at, name, user_id, start_date, end_date FROM habits
WHERE id = ?
`

func (q *Queries) GetHabit(ctx context.Context, id string) (Habit, error) {
	row := q.db.QueryRowContext(ctx, getHabit, id)
	var i Habit
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.UserID,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const getUserHabits = `-- name: GetUserHabits :many
SELECT id, created_at, updated_at, name, user_id, start_date, end_date FROM habits
WHERE user_id = ?
`

func (q *Queries) GetUserHabits(ctx context.Context, userID string) ([]Habit, error) {
	rows, err := q.db.QueryContext(ctx, getUserHabits, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Habit
	for rows.Next() {
		var i Habit
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.UserID,
			&i.StartDate,
			&i.EndDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeHabit = `-- name: RemoveHabit :exec
DELETE FROM habits
WHERE id = ?
`

func (q *Queries) RemoveHabit(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, removeHabit, id)
	return err
}
