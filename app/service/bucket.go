package service

import (
	"github.com/juju/ratelimit"
	"time"
)

//global bucket
type Msg struct {
	Status		bool			`json:"status"`
	Detail		string			`json:"detail"`
	Available	int64			`json:"available"`
	Rate		float64			`json:"rate"`
	Capacity	int64			`json:"capacity"`
	MaxWait		time.Duration	`json:"max_wait"`
}

type Bucket struct {
	Rate		float64
	Capacity	int64
	MaxWait		time.Duration
	Bucket		*ratelimit.Bucket
}

func NewBucket(rate float64,capacity int64,maxWait time.Duration) *Bucket {
	return &Bucket{
		rate,
		capacity,
		maxWait,
		ratelimit.NewBucketWithRate(rate,capacity),
	}
}

func (b *Bucket) GetToken(count int64) *Msg {
	return &Msg{
		b.Bucket.WaitMaxDuration(count,b.MaxWait),
		"",
		b.Bucket.Available(),
		b.Rate,
		b.Capacity,
		b.MaxWait,
	}
}

//api bucket

//user bucket