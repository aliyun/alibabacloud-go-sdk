// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLFlowSankeyExit interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *RLFlowSankeyExit
	GetCount() *int32
	SetFrom(v string) *RLFlowSankeyExit
	GetFrom() *string
	SetFromIdx(v int32) *RLFlowSankeyExit
	GetFromIdx() *int32
	SetLabel(v string) *RLFlowSankeyExit
	GetLabel() *string
}

type RLFlowSankeyExit struct {
	// The number of trajectories on the outflow edge.
	//
	// example:
	//
	// 96
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The key of the outflow source column.
	//
	// example:
	//
	// DataLoader
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The index of the outflow source column (0-based).
	//
	// example:
	//
	// 0
	FromIdx *int32 `json:"FromIdx,omitempty" xml:"FromIdx,omitempty"`
	// The Chinese name of the outflow destination. Valid values vary by the column where the outflow is located: 在途·未下发 / 在途·生成中 / 在途·待采样 / 在途·待训练.
	//
	// example:
	//
	// 生成轨迹
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s RLFlowSankeyExit) String() string {
	return dara.Prettify(s)
}

func (s RLFlowSankeyExit) GoString() string {
	return s.String()
}

func (s *RLFlowSankeyExit) GetCount() *int32 {
	return s.Count
}

func (s *RLFlowSankeyExit) GetFrom() *string {
	return s.From
}

func (s *RLFlowSankeyExit) GetFromIdx() *int32 {
	return s.FromIdx
}

func (s *RLFlowSankeyExit) GetLabel() *string {
	return s.Label
}

func (s *RLFlowSankeyExit) SetCount(v int32) *RLFlowSankeyExit {
	s.Count = &v
	return s
}

func (s *RLFlowSankeyExit) SetFrom(v string) *RLFlowSankeyExit {
	s.From = &v
	return s
}

func (s *RLFlowSankeyExit) SetFromIdx(v int32) *RLFlowSankeyExit {
	s.FromIdx = &v
	return s
}

func (s *RLFlowSankeyExit) SetLabel(v string) *RLFlowSankeyExit {
	s.Label = &v
	return s
}

func (s *RLFlowSankeyExit) Validate() error {
	return dara.Validate(s)
}
