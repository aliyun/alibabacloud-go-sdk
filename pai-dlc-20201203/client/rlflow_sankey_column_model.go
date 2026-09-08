// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLFlowSankeyColumn interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *RLFlowSankeyColumn
	GetCount() *int32
	SetKey(v string) *RLFlowSankeyColumn
	GetKey() *string
	SetLabel(v string) *RLFlowSankeyColumn
	GetLabel() *string
}

type RLFlowSankeyColumn struct {
	// The number of trajectories in the column. The value is monotonized: reaching a later stage implies having passed through all preceding stages.
	//
	// example:
	//
	// 96
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The column identifier. Valid values: gen, run, rollout, sampled, and trained.
	//
	// example:
	//
	// traj
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The Chinese name of the column. Valid values: 轨迹生成, Agent 启动, Rollout 完成, 采样入批, and 完成训练.
	//
	// example:
	//
	// 生成轨迹
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s RLFlowSankeyColumn) String() string {
	return dara.Prettify(s)
}

func (s RLFlowSankeyColumn) GoString() string {
	return s.String()
}

func (s *RLFlowSankeyColumn) GetCount() *int32 {
	return s.Count
}

func (s *RLFlowSankeyColumn) GetKey() *string {
	return s.Key
}

func (s *RLFlowSankeyColumn) GetLabel() *string {
	return s.Label
}

func (s *RLFlowSankeyColumn) SetCount(v int32) *RLFlowSankeyColumn {
	s.Count = &v
	return s
}

func (s *RLFlowSankeyColumn) SetKey(v string) *RLFlowSankeyColumn {
	s.Key = &v
	return s
}

func (s *RLFlowSankeyColumn) SetLabel(v string) *RLFlowSankeyColumn {
	s.Label = &v
	return s
}

func (s *RLFlowSankeyColumn) Validate() error {
	return dara.Validate(s)
}
