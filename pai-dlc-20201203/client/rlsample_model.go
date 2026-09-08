// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLSample interface {
	dara.Model
	String() string
	GoString() string
	SetLatestDetail(v string) *RLSample
	GetLatestDetail() *string
	SetLatestStage(v string) *RLSample
	GetLatestStage() *string
	SetLatestStatus(v string) *RLSample
	GetLatestStatus() *string
	SetLatestTimestampMs(v int64) *RLSample
	GetLatestTimestampMs() *int64
	SetPromptUid(v string) *RLSample
	GetPromptUid() *string
	SetSampleIndex(v string) *RLSample
	GetSampleIndex() *string
	SetTerminalState(v string) *RLSample
	GetTerminalState() *string
	SetTraceCount(v int64) *RLSample
	GetTraceCount() *int64
}

type RLSample struct {
	// The detail of the latest event.
	//
	// example:
	//
	// actor_parameters_updated
	LatestDetail *string `json:"LatestDetail,omitempty" xml:"LatestDetail,omitempty"`
	// The stage of the latest event.
	//
	// example:
	//
	// TRAIN_UPDATE
	LatestStage *string `json:"LatestStage,omitempty" xml:"LatestStage,omitempty"`
	// The latest sample_status.
	//
	// example:
	//
	// COMPLETE
	LatestStatus *string `json:"LatestStatus,omitempty" xml:"LatestStatus,omitempty"`
	// The millisecond timestamp of the latest event.
	//
	// example:
	//
	// 1787293215480
	LatestTimestampMs *int64 `json:"LatestTimestampMs,omitempty" xml:"LatestTimestampMs,omitempty"`
	// The sample UID.
	//
	// example:
	//
	// 321fa56f-e1e5-4eb3-8047-db7a230c9a75
	PromptUid *string `json:"PromptUid,omitempty" xml:"PromptUid,omitempty"`
	// The trajectory ordinal number (numeric string).
	//
	// example:
	//
	// 2
	SampleIndex *string `json:"SampleIndex,omitempty" xml:"SampleIndex,omitempty"`
	// The desired state. Valid values: trained (training completed) and empty string (in progress). The current frame does not perform oversampling, so discarded and cancelled do not occur.
	//
	// example:
	//
	// trained
	TerminalState *string `json:"TerminalState,omitempty" xml:"TerminalState,omitempty"`
	// The number of trace rows for the trajectory, including B/C type allocations.
	//
	// example:
	//
	// 10
	TraceCount *int64 `json:"TraceCount,omitempty" xml:"TraceCount,omitempty"`
}

func (s RLSample) String() string {
	return dara.Prettify(s)
}

func (s RLSample) GoString() string {
	return s.String()
}

func (s *RLSample) GetLatestDetail() *string {
	return s.LatestDetail
}

func (s *RLSample) GetLatestStage() *string {
	return s.LatestStage
}

func (s *RLSample) GetLatestStatus() *string {
	return s.LatestStatus
}

func (s *RLSample) GetLatestTimestampMs() *int64 {
	return s.LatestTimestampMs
}

func (s *RLSample) GetPromptUid() *string {
	return s.PromptUid
}

func (s *RLSample) GetSampleIndex() *string {
	return s.SampleIndex
}

func (s *RLSample) GetTerminalState() *string {
	return s.TerminalState
}

func (s *RLSample) GetTraceCount() *int64 {
	return s.TraceCount
}

func (s *RLSample) SetLatestDetail(v string) *RLSample {
	s.LatestDetail = &v
	return s
}

func (s *RLSample) SetLatestStage(v string) *RLSample {
	s.LatestStage = &v
	return s
}

func (s *RLSample) SetLatestStatus(v string) *RLSample {
	s.LatestStatus = &v
	return s
}

func (s *RLSample) SetLatestTimestampMs(v int64) *RLSample {
	s.LatestTimestampMs = &v
	return s
}

func (s *RLSample) SetPromptUid(v string) *RLSample {
	s.PromptUid = &v
	return s
}

func (s *RLSample) SetSampleIndex(v string) *RLSample {
	s.SampleIndex = &v
	return s
}

func (s *RLSample) SetTerminalState(v string) *RLSample {
	s.TerminalState = &v
	return s
}

func (s *RLSample) SetTraceCount(v int64) *RLSample {
	s.TraceCount = &v
	return s
}

func (s *RLSample) Validate() error {
	return dara.Validate(s)
}
