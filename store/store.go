package store

import (
	"errors"

	"githum.com/hikobend/go_todo/entity"
)

var (
	Tasks = &TaskStore{Tasks: map[entity.TaskID]*entity.Task{}}

	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	LastID entity.TaskID
	Tasks  map[entity.TaskID]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

func (ts *TaskStore) Get(id entity.TaskID) (*entity.Task, error) {
	if ts, ok := ts.Tasks[id]; ok {
		return ts, nil
	}
	return nil, ErrNotFound
}
