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
	SetStartTime(v int64) *BackfillStrategy
	GetStartTime() *int64
}

type BackfillStrategy struct {
	// Specifies whether the backfill policy is enabled. If this parameter is not specified or is set to true, the policy is enabled. If this parameter is set to false, the policy is disabled but the configuration is retained.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The end of the backfill time range, in UNIX millisecond timestamp. Provide a complete time range when you need to manually start a backfill.
	//
	// example:
	//
	// 1782902400000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start of the backfill time range, in UNIX millisecond timestamp. Provide a complete time range when you need to manually start a backfill.
	//
	// example:
	//
	// 1782816000000
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

func (s *BackfillStrategy) SetStartTime(v int64) *BackfillStrategy {
	s.StartTime = &v
	return s
}

func (s *BackfillStrategy) Validate() error {
	return dara.Validate(s)
}
