// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLFlowFunnelStage interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *RLFlowFunnelStage
	GetCount() *int32
	SetKey(v string) *RLFlowFunnelStage
	GetKey() *string
	SetLabel(v string) *RLFlowFunnelStage
	GetLabel() *string
	SetPct(v float64) *RLFlowFunnelStage
	GetPct() *float64
}

type RLFlowFunnelStage struct {
	// The number of trajectories that reach this level.
	//
	// example:
	//
	// 96
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The level identifier. Valid values: traj, dispatch, run, rollout, reward, sampled, and trained.
	//
	// example:
	//
	// traj
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The Chinese name of the level. Valid values: 生成轨迹, 下发到 Worker, Agent 启动, Rollout 完成, reward 打分, 采样入批, and 完成训练.
	//
	// example:
	//
	// 生成轨迹
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The percentage relative to the first traj level.
	//
	// example:
	//
	// 100
	Pct *float64 `json:"Pct,omitempty" xml:"Pct,omitempty"`
}

func (s RLFlowFunnelStage) String() string {
	return dara.Prettify(s)
}

func (s RLFlowFunnelStage) GoString() string {
	return s.String()
}

func (s *RLFlowFunnelStage) GetCount() *int32 {
	return s.Count
}

func (s *RLFlowFunnelStage) GetKey() *string {
	return s.Key
}

func (s *RLFlowFunnelStage) GetLabel() *string {
	return s.Label
}

func (s *RLFlowFunnelStage) GetPct() *float64 {
	return s.Pct
}

func (s *RLFlowFunnelStage) SetCount(v int32) *RLFlowFunnelStage {
	s.Count = &v
	return s
}

func (s *RLFlowFunnelStage) SetKey(v string) *RLFlowFunnelStage {
	s.Key = &v
	return s
}

func (s *RLFlowFunnelStage) SetLabel(v string) *RLFlowFunnelStage {
	s.Label = &v
	return s
}

func (s *RLFlowFunnelStage) SetPct(v float64) *RLFlowFunnelStage {
	s.Pct = &v
	return s
}

func (s *RLFlowFunnelStage) Validate() error {
	return dara.Validate(s)
}
