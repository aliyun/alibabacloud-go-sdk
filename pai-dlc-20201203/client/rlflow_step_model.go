// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLFlowStep interface {
	dara.Model
	String() string
	GoString() string
	SetBufferWaitP50(v float64) *RLFlowStep
	GetBufferWaitP50() *float64
	SetGapSec(v float64) *RLFlowStep
	GetGapSec() *float64
	SetIdleSec(v float64) *RLFlowStep
	GetIdleSec() *float64
	SetNSamples(v int32) *RLFlowStep
	GetNSamples() *int32
	SetNTrajs(v int32) *RLFlowStep
	GetNTrajs() *int32
	SetProdEndMs(v int64) *RLFlowStep
	GetProdEndMs() *int64
	SetProdStartMs(v int64) *RLFlowStep
	GetProdStartMs() *int64
	SetRolloutP50(v float64) *RLFlowStep
	GetRolloutP50() *float64
	SetRolloutSec(v float64) *RLFlowStep
	GetRolloutSec() *float64
	SetStep(v int64) *RLFlowStep
	GetStep() *int64
	SetTFwdStartMs(v int64) *RLFlowStep
	GetTFwdStartMs() *int64
	SetTOptEndMs(v int64) *RLFlowStep
	GetTOptEndMs() *int64
	SetTRolloutEndMs(v int64) *RLFlowStep
	GetTRolloutEndMs() *int64
	SetTRolloutStartMs(v int64) *RLFlowStep
	GetTRolloutStartMs() *int64
	SetTTrainEndMs(v int64) *RLFlowStep
	GetTTrainEndMs() *int64
	SetTTrainStartMs(v int64) *RLFlowStep
	GetTTrainStartMs() *int64
	SetTUpdateMs(v int64) *RLFlowStep
	GetTUpdateMs() *int64
	SetTrainSec(v float64) *RLFlowStep
	GetTrainSec() *float64
}

type RLFlowStep struct {
	// The P50 latency in seconds from when trajectories of the step enter the buffer to when batching occurs.
	//
	// example:
	//
	// 0.8
	BufferWaitP50 *float64 `json:"BufferWaitP50,omitempty" xml:"BufferWaitP50,omitempty"`
	// The gap duration in seconds, calculated as train started − rollout finished. This represents batching or transfer wait time.
	//
	// example:
	//
	// 0
	GapSec *float64 `json:"GapSec,omitempty" xml:"GapSec,omitempty"`
	// The training idle time in seconds, calculated as the current step training start − the previous step training end. If no marker is present, the value falls back to the current step forward computation start − the previous step optimizer end. A value greater than 0 indicates that the trainer is waiting for data.
	//
	// example:
	//
	// 0
	IdleSec *float64 `json:"IdleSec,omitempty" xml:"IdleSec,omitempty"`
	// The number of samples (UIDs) consumed by the step.
	//
	// example:
	//
	// 24
	NSamples *int32 `json:"NSamples,omitempty" xml:"NSamples,omitempty"`
	// The number of trajectories executed in the step.
	//
	// example:
	//
	// 96
	NTrajs *int32 `json:"NTrajs,omitempty" xml:"NTrajs,omitempty"`
	// The latest time when trajectories of the step enter the buffer, in milliseconds.
	//
	// example:
	//
	// 1787474487713
	ProdEndMs *int64 `json:"ProdEndMs,omitempty" xml:"ProdEndMs,omitempty"`
	// The earliest time when trajectories of the step enter the buffer, in milliseconds.
	//
	// example:
	//
	// 1787474487713
	ProdStartMs *int64 `json:"ProdStartMs,omitempty" xml:"ProdStartMs,omitempty"`
	// The P50 latency in seconds from when trajectories of the step start execution to when they enter the buffer.
	//
	// example:
	//
	// 0.8
	RolloutP50 *float64 `json:"RolloutP50,omitempty" xml:"RolloutP50,omitempty"`
	// The rollout duration in seconds, calculated as rollout finished − rollout started. This value is null if no marker is present.
	//
	// example:
	//
	// 0
	RolloutSec *float64 `json:"RolloutSec,omitempty" xml:"RolloutSec,omitempty"`
	// The global step ordinal number.
	//
	// example:
	//
	// 3
	Step *int64 `json:"Step,omitempty" xml:"Step,omitempty"`
	// The forward computation start time, in milliseconds.
	//
	// example:
	//
	// 1787474487713
	TFwdStartMs *int64 `json:"TFwdStartMs,omitempty" xml:"TFwdStartMs,omitempty"`
	// The optimizer end time, in milliseconds.
	//
	// example:
	//
	// 1787474487713
	TOptEndMs *int64 `json:"TOptEndMs,omitempty" xml:"TOptEndMs,omitempty"`
	// The node operation log "Step N rollout finished" time, in milliseconds.
	//
	// example:
	//
	// 1787474487713
	TRolloutEndMs *int64 `json:"TRolloutEndMs,omitempty" xml:"TRolloutEndMs,omitempty"`
	// The node operation log "Step N rollout started" time, in milliseconds (taken from agent_collect_time).
	//
	// example:
	//
	// 1787474487713
	TRolloutStartMs *int64 `json:"TRolloutStartMs,omitempty" xml:"TRolloutStartMs,omitempty"`
	// The node operation log "Step N train finished" time, in milliseconds.
	//
	// example:
	//
	// 1787474487713
	TTrainEndMs *int64 `json:"TTrainEndMs,omitempty" xml:"TTrainEndMs,omitempty"`
	// The node operation log "Step N train started" time, in milliseconds.
	//
	// example:
	//
	// 1787474487713
	TTrainStartMs *int64 `json:"TTrainStartMs,omitempty" xml:"TTrainStartMs,omitempty"`
	// The TRAIN_UPDATE (parameter update) time, in milliseconds.
	//
	// example:
	//
	// 1787474487713
	TUpdateMs *int64 `json:"TUpdateMs,omitempty" xml:"TUpdateMs,omitempty"`
	// The training duration in seconds. This value is preferentially calculated as train finished − train started. If no marker is present, the value falls back to the duration from batching to training completion. This value is null if global_step is duplicated because of a job restart.
	//
	// example:
	//
	// 0
	TrainSec *float64 `json:"TrainSec,omitempty" xml:"TrainSec,omitempty"`
}

func (s RLFlowStep) String() string {
	return dara.Prettify(s)
}

func (s RLFlowStep) GoString() string {
	return s.String()
}

func (s *RLFlowStep) GetBufferWaitP50() *float64 {
	return s.BufferWaitP50
}

func (s *RLFlowStep) GetGapSec() *float64 {
	return s.GapSec
}

func (s *RLFlowStep) GetIdleSec() *float64 {
	return s.IdleSec
}

func (s *RLFlowStep) GetNSamples() *int32 {
	return s.NSamples
}

func (s *RLFlowStep) GetNTrajs() *int32 {
	return s.NTrajs
}

func (s *RLFlowStep) GetProdEndMs() *int64 {
	return s.ProdEndMs
}

func (s *RLFlowStep) GetProdStartMs() *int64 {
	return s.ProdStartMs
}

func (s *RLFlowStep) GetRolloutP50() *float64 {
	return s.RolloutP50
}

func (s *RLFlowStep) GetRolloutSec() *float64 {
	return s.RolloutSec
}

func (s *RLFlowStep) GetStep() *int64 {
	return s.Step
}

func (s *RLFlowStep) GetTFwdStartMs() *int64 {
	return s.TFwdStartMs
}

func (s *RLFlowStep) GetTOptEndMs() *int64 {
	return s.TOptEndMs
}

func (s *RLFlowStep) GetTRolloutEndMs() *int64 {
	return s.TRolloutEndMs
}

func (s *RLFlowStep) GetTRolloutStartMs() *int64 {
	return s.TRolloutStartMs
}

func (s *RLFlowStep) GetTTrainEndMs() *int64 {
	return s.TTrainEndMs
}

func (s *RLFlowStep) GetTTrainStartMs() *int64 {
	return s.TTrainStartMs
}

func (s *RLFlowStep) GetTUpdateMs() *int64 {
	return s.TUpdateMs
}

func (s *RLFlowStep) GetTrainSec() *float64 {
	return s.TrainSec
}

func (s *RLFlowStep) SetBufferWaitP50(v float64) *RLFlowStep {
	s.BufferWaitP50 = &v
	return s
}

func (s *RLFlowStep) SetGapSec(v float64) *RLFlowStep {
	s.GapSec = &v
	return s
}

func (s *RLFlowStep) SetIdleSec(v float64) *RLFlowStep {
	s.IdleSec = &v
	return s
}

func (s *RLFlowStep) SetNSamples(v int32) *RLFlowStep {
	s.NSamples = &v
	return s
}

func (s *RLFlowStep) SetNTrajs(v int32) *RLFlowStep {
	s.NTrajs = &v
	return s
}

func (s *RLFlowStep) SetProdEndMs(v int64) *RLFlowStep {
	s.ProdEndMs = &v
	return s
}

func (s *RLFlowStep) SetProdStartMs(v int64) *RLFlowStep {
	s.ProdStartMs = &v
	return s
}

func (s *RLFlowStep) SetRolloutP50(v float64) *RLFlowStep {
	s.RolloutP50 = &v
	return s
}

func (s *RLFlowStep) SetRolloutSec(v float64) *RLFlowStep {
	s.RolloutSec = &v
	return s
}

func (s *RLFlowStep) SetStep(v int64) *RLFlowStep {
	s.Step = &v
	return s
}

func (s *RLFlowStep) SetTFwdStartMs(v int64) *RLFlowStep {
	s.TFwdStartMs = &v
	return s
}

func (s *RLFlowStep) SetTOptEndMs(v int64) *RLFlowStep {
	s.TOptEndMs = &v
	return s
}

func (s *RLFlowStep) SetTRolloutEndMs(v int64) *RLFlowStep {
	s.TRolloutEndMs = &v
	return s
}

func (s *RLFlowStep) SetTRolloutStartMs(v int64) *RLFlowStep {
	s.TRolloutStartMs = &v
	return s
}

func (s *RLFlowStep) SetTTrainEndMs(v int64) *RLFlowStep {
	s.TTrainEndMs = &v
	return s
}

func (s *RLFlowStep) SetTTrainStartMs(v int64) *RLFlowStep {
	s.TTrainStartMs = &v
	return s
}

func (s *RLFlowStep) SetTUpdateMs(v int64) *RLFlowStep {
	s.TUpdateMs = &v
	return s
}

func (s *RLFlowStep) SetTrainSec(v float64) *RLFlowStep {
	s.TrainSec = &v
	return s
}

func (s *RLFlowStep) Validate() error {
	return dara.Validate(s)
}
