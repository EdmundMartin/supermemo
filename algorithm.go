package supermemo

import (
	"fmt"
	"math"
	"time"
)

type Note struct {
	Ease        float64
	Interval    float64
	Repetitions int
	ReviewDate  time.Time
}

func (n *Note) calculateIntervalReps(quality int) (float64, int) {
	if quality < 3 {
		return 1, 0
	}

	if n.Repetitions == 0 {
		return 1, n.Repetitions + 1
	}

	if n.Repetitions == 1 {
		return 6, n.Repetitions + 1
	}

	ease := math.Ceil(n.Interval * n.Ease)

	return ease, n.Repetitions + 1
}

func (n *Note) Review(quality int, reviewDate time.Time) {
	n.Interval, n.Repetitions = n.calculateIntervalReps(quality)

	qual := float64(quality)
	n.Ease += 0.1 - (5-qual)*(0.08+(5-qual)*0.02)
	n.Ease = math.Max(1.3, n.Ease)

	delta := int(n.Interval)
	n.ReviewDate = reviewDate.AddDate(0, 0, delta)
}

func (n *Note) String() string {
	return fmt.Sprintf("<Note: Ease: %f, Reps: %d, Interval: %f>", n.Ease, n.Repetitions, n.Interval)
}

func FirstReview(quality int) *Note {
	note := &Note{
		Ease:        2.5,
		Interval:    0,
		Repetitions: 0,
	}
	note.Review(quality, time.Now())
	return note
}

