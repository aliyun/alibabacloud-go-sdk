// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunStrategies interface {
	dara.Model
	String() string
	GoString() string
	SetBackfill(v *BackfillStrategy) *RunStrategies
	GetBackfill() *BackfillStrategy
	SetContinuous(v *ContinuousStrategy) *RunStrategies
	GetContinuous() *ContinuousStrategy
}

type RunStrategies struct {
	// The historical batch backfill policy. Backfill is enabled when the object exists and enabled is not explicitly set to false.
	//
	// example:
	//
	// {"enabled":true,"startTime":1782816000000,"endTime":1782902400000,"immediate":false}
	Backfill *BackfillStrategy `json:"backfill,omitempty" xml:"backfill,omitempty"`
	// The continuous evaluation policy. Continuous evaluation is enabled when the object exists and enabled is not explicitly set to false.
	//
	// example:
	//
	// {"enabled":true,"intervalUnit":"HOUR","intervalValue":1,"dataDelayMinutes":5}
	Continuous *ContinuousStrategy `json:"continuous,omitempty" xml:"continuous,omitempty"`
}

func (s RunStrategies) String() string {
	return dara.Prettify(s)
}

func (s RunStrategies) GoString() string {
	return s.String()
}

func (s *RunStrategies) GetBackfill() *BackfillStrategy {
	return s.Backfill
}

func (s *RunStrategies) GetContinuous() *ContinuousStrategy {
	return s.Continuous
}

func (s *RunStrategies) SetBackfill(v *BackfillStrategy) *RunStrategies {
	s.Backfill = v
	return s
}

func (s *RunStrategies) SetContinuous(v *ContinuousStrategy) *RunStrategies {
	s.Continuous = v
	return s
}

func (s *RunStrategies) Validate() error {
	if s.Backfill != nil {
		if err := s.Backfill.Validate(); err != nil {
			return err
		}
	}
	if s.Continuous != nil {
		if err := s.Continuous.Validate(); err != nil {
			return err
		}
	}
	return nil
}
