// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLFlowStuckItem interface {
	dara.Model
	String() string
	GoString() string
	SetIdleSec(v int64) *RLFlowStuckItem
	GetIdleSec() *int64
	SetLastTsMs(v int64) *RLFlowStuckItem
	GetLastTsMs() *int64
	SetMilestone(v string) *RLFlowStuckItem
	GetMilestone() *string
	SetNTurns(v int32) *RLFlowStuckItem
	GetNTurns() *int32
	SetPromptUid(v string) *RLFlowStuckItem
	GetPromptUid() *string
	SetSampleIndex(v string) *RLFlowStuckItem
	GetSampleIndex() *string
}

type RLFlowStuckItem struct {
	// The number of idle seconds since NowMs. This value is the descending sort key of the Stuck list.
	//
	// example:
	//
	// 0
	IdleSec *int64 `json:"IdleSec,omitempty" xml:"IdleSec,omitempty"`
	// The UNIX timestamp of the last event, in milliseconds.
	//
	// example:
	//
	// 1787293215480
	LastTsMs *int64 `json:"LastTsMs,omitempty" xml:"LastTsMs,omitempty"`
	// The current milestone where the entry is staying. Valid values:
	//
	// - 已生成未下发: Generated but not delivered.
	//
	// - 已下发未启动: Delivered but not started.
	//
	// - 已启动待生成: Started and pending generation.
	//
	// - 生成中: Generating.
	//
	// - Rollout完成待打分: Rollout completed and pending scoring.
	//
	// - 已打分待采样: Scored and pending sampling.
	//
	// - 已采样待训练: Sampled and pending training.
	//
	// example:
	//
	// 生成中
	Milestone *string `json:"Milestone,omitempty" xml:"Milestone,omitempty"`
	// The number of completed generation rounds.
	//
	// example:
	//
	// 3
	NTurns *int32 `json:"NTurns,omitempty" xml:"NTurns,omitempty"`
	// The UID of the sample.
	//
	// example:
	//
	// 321fa56f-e1e5-4eb3-8047-db7a230c9a75
	PromptUid *string `json:"PromptUid,omitempty" xml:"PromptUid,omitempty"`
	// The ordinal number of the trajectory.
	//
	// example:
	//
	// 2
	SampleIndex *string `json:"SampleIndex,omitempty" xml:"SampleIndex,omitempty"`
}

func (s RLFlowStuckItem) String() string {
	return dara.Prettify(s)
}

func (s RLFlowStuckItem) GoString() string {
	return s.String()
}

func (s *RLFlowStuckItem) GetIdleSec() *int64 {
	return s.IdleSec
}

func (s *RLFlowStuckItem) GetLastTsMs() *int64 {
	return s.LastTsMs
}

func (s *RLFlowStuckItem) GetMilestone() *string {
	return s.Milestone
}

func (s *RLFlowStuckItem) GetNTurns() *int32 {
	return s.NTurns
}

func (s *RLFlowStuckItem) GetPromptUid() *string {
	return s.PromptUid
}

func (s *RLFlowStuckItem) GetSampleIndex() *string {
	return s.SampleIndex
}

func (s *RLFlowStuckItem) SetIdleSec(v int64) *RLFlowStuckItem {
	s.IdleSec = &v
	return s
}

func (s *RLFlowStuckItem) SetLastTsMs(v int64) *RLFlowStuckItem {
	s.LastTsMs = &v
	return s
}

func (s *RLFlowStuckItem) SetMilestone(v string) *RLFlowStuckItem {
	s.Milestone = &v
	return s
}

func (s *RLFlowStuckItem) SetNTurns(v int32) *RLFlowStuckItem {
	s.NTurns = &v
	return s
}

func (s *RLFlowStuckItem) SetPromptUid(v string) *RLFlowStuckItem {
	s.PromptUid = &v
	return s
}

func (s *RLFlowStuckItem) SetSampleIndex(v string) *RLFlowStuckItem {
	s.SampleIndex = &v
	return s
}

func (s *RLFlowStuckItem) Validate() error {
	return dara.Validate(s)
}
