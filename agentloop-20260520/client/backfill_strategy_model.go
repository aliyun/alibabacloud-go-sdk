// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackfillStrategy interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *BackfillStrategy
	GetEnabled() *bool
	SetEndTime(v int64) *BackfillStrategy
	GetEndTime() *int64
	SetImmediate(v bool) *BackfillStrategy
	GetImmediate() *bool
	SetStartTime(v int64) *BackfillStrategy
	GetStartTime() *int64
}

type BackfillStrategy struct {
	Enabled   *bool  `json:"enabled,omitempty" xml:"enabled,omitempty"`
	EndTime   *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Immediate *bool  `json:"immediate,omitempty" xml:"immediate,omitempty"`
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s BackfillStrategy) String() string {
	return dara.Prettify(s)
}

func (s BackfillStrategy) GoString() string {
	return s.String()
}

func (s *BackfillStrategy) GetEnabled() *bool {
	return s.Enabled
}

func (s *BackfillStrategy) GetEndTime() *int64 {
	return s.EndTime
}

func (s *BackfillStrategy) GetImmediate() *bool {
	return s.Immediate
}

func (s *BackfillStrategy) GetStartTime() *int64 {
	return s.StartTime
}

func (s *BackfillStrategy) SetEnabled(v bool) *BackfillStrategy {
	s.Enabled = &v
	return s
}

func (s *BackfillStrategy) SetEndTime(v int64) *BackfillStrategy {
	s.EndTime = &v
	return s
}

func (s *BackfillStrategy) SetImmediate(v bool) *BackfillStrategy {
	s.Immediate = &v
	return s
}

func (s *BackfillStrategy) SetStartTime(v int64) *BackfillStrategy {
	s.StartTime = &v
	return s
}

func (s *BackfillStrategy) Validate() error {
	return dara.Validate(s)
}
