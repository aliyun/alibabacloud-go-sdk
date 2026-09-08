// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLTrajectory interface {
	dara.Model
	String() string
	GoString() string
	SetLatestTimestampMs(v int64) *RLTrajectory
	GetLatestTimestampMs() *int64
	SetSampleIndex(v string) *RLTrajectory
	GetSampleIndex() *string
	SetTerminalState(v string) *RLTrajectory
	GetTerminalState() *string
	SetTraceCount(v int64) *RLTrajectory
	GetTraceCount() *int64
}

type RLTrajectory struct {
	// The latest event millisecond UNIX timestamp.
	//
	// example:
	//
	// 1787293215480
	LatestTimestampMs *int64 `json:"LatestTimestampMs,omitempty" xml:"LatestTimestampMs,omitempty"`
	// The trajectory ordinal number.
	//
	// example:
	//
	// 2
	SampleIndex *string `json:"SampleIndex,omitempty" xml:"SampleIndex,omitempty"`
	// The desired state. Valid values:
	//
	// - trained: Training is complete.
	//
	// - Empty string: In progress.
	//
	// The current frame does not perform oversampling, so discarded and cancelled do not occur.
	//
	// example:
	//
	// trained
	TerminalState *string `json:"TerminalState,omitempty" xml:"TerminalState,omitempty"`
	// The number of trace rows.
	//
	// example:
	//
	// 10
	TraceCount *int64 `json:"TraceCount,omitempty" xml:"TraceCount,omitempty"`
}

func (s RLTrajectory) String() string {
	return dara.Prettify(s)
}

func (s RLTrajectory) GoString() string {
	return s.String()
}

func (s *RLTrajectory) GetLatestTimestampMs() *int64 {
	return s.LatestTimestampMs
}

func (s *RLTrajectory) GetSampleIndex() *string {
	return s.SampleIndex
}

func (s *RLTrajectory) GetTerminalState() *string {
	return s.TerminalState
}

func (s *RLTrajectory) GetTraceCount() *int64 {
	return s.TraceCount
}

func (s *RLTrajectory) SetLatestTimestampMs(v int64) *RLTrajectory {
	s.LatestTimestampMs = &v
	return s
}

func (s *RLTrajectory) SetSampleIndex(v string) *RLTrajectory {
	s.SampleIndex = &v
	return s
}

func (s *RLTrajectory) SetTerminalState(v string) *RLTrajectory {
	s.TerminalState = &v
	return s
}

func (s *RLTrajectory) SetTraceCount(v int64) *RLTrajectory {
	s.TraceCount = &v
	return s
}

func (s *RLTrajectory) Validate() error {
	return dara.Validate(s)
}
