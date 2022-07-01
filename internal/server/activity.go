package server

import (
	"fmt"
	"sync"
	"time"
)

type Activity struct {
	Time        time.time
	Description string
	ID          uint64
}

type Activities struct {
	activities []Activity
}

func (c *Activities) Insert(activity Acitivity) uint64 {
	activity.ID = uint64(len(c.activities))
	c.activities = append(c.activities, activity)
	return activity.ID
}

func (c *Activities) Retrieve(id uint64) (Activity, error) {
	var ErrIDNotFound = fmt.Errorf("ID not found")
	if id >= uint64(len(c.activities)) {
		return Activity{}, ErrIDNotFound
	}

	return c.activities[id], nil
}
