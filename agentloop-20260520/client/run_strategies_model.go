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
	Backfill   *BackfillStrategy   `json:"backfill,omitempty" xml:"backfill,omitempty"`
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
