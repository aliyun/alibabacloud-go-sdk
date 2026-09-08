// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLSampleEvent interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v string) *RLSampleEvent
	GetDetail() *string
	SetFrom(v string) *RLSampleEvent
	GetFrom() *string
	SetGlobalStep(v string) *RLSampleEvent
	GetGlobalStep() *string
	SetStage(v string) *RLSampleEvent
	GetStage() *string
	SetTimestampMs(v int64) *RLSampleEvent
	GetTimestampMs() *int64
	SetTo(v string) *RLSampleEvent
	GetTo() *string
}

type RLSampleEvent struct {
	// The details. For Megatron rows, the value is rank=..,global_step=..,ppo_epoch=..
	//
	// example:
	//
	// uid_generated
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The event source component. For Megatron rows, the value is "{phase} {status}".
	//
	// example:
	//
	// DataLoader
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The training step to which the event belongs (raw string). For Megatron rows, this is empty because the step is included in Detail.
	//
	// example:
	//
	// 12
	GlobalStep *string `json:"GlobalStep,omitempty" xml:"GlobalStep,omitempty"`
	// The stage. For Megatron rows, this is normalized to TRAIN.
	//
	// example:
	//
	// DATA_PREPROCESS
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// The millisecond timestamp.
	//
	// example:
	//
	// 1787293208012
	TimestampMs *int64 `json:"TimestampMs,omitempty" xml:"TimestampMs,omitempty"`
	// The event target component. For Megatron rows, the value is the function name.
	//
	// example:
	//
	// PPOTrainerV1
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s RLSampleEvent) String() string {
	return dara.Prettify(s)
}

func (s RLSampleEvent) GoString() string {
	return s.String()
}

func (s *RLSampleEvent) GetDetail() *string {
	return s.Detail
}

func (s *RLSampleEvent) GetFrom() *string {
	return s.From
}

func (s *RLSampleEvent) GetGlobalStep() *string {
	return s.GlobalStep
}

func (s *RLSampleEvent) GetStage() *string {
	return s.Stage
}

func (s *RLSampleEvent) GetTimestampMs() *int64 {
	return s.TimestampMs
}

func (s *RLSampleEvent) GetTo() *string {
	return s.To
}

func (s *RLSampleEvent) SetDetail(v string) *RLSampleEvent {
	s.Detail = &v
	return s
}

func (s *RLSampleEvent) SetFrom(v string) *RLSampleEvent {
	s.From = &v
	return s
}

func (s *RLSampleEvent) SetGlobalStep(v string) *RLSampleEvent {
	s.GlobalStep = &v
	return s
}

func (s *RLSampleEvent) SetStage(v string) *RLSampleEvent {
	s.Stage = &v
	return s
}

func (s *RLSampleEvent) SetTimestampMs(v int64) *RLSampleEvent {
	s.TimestampMs = &v
	return s
}

func (s *RLSampleEvent) SetTo(v string) *RLSampleEvent {
	s.To = &v
	return s
}

func (s *RLSampleEvent) Validate() error {
	return dara.Validate(s)
}
