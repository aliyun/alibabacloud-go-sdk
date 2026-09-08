// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLFlowTurns interface {
	dara.Model
	String() string
	GoString() string
	SetAvg(v float64) *RLFlowTurns
	GetAvg() *float64
	SetCount(v int32) *RLFlowTurns
	GetCount() *int32
	SetMax(v int32) *RLFlowTurns
	GetMax() *int32
	SetP50(v int32) *RLFlowTurns
	GetP50() *int32
	SetP90(v int32) *RLFlowTurns
	GetP90() *int32
}

type RLFlowTurns struct {
	// 平均生成轮数
	//
	// example:
	//
	// 1.9
	Avg *float64 `json:"Avg,omitempty" xml:"Avg,omitempty"`
	// 参与统计的轨迹数
	//
	// example:
	//
	// 96
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 最大生成轮数
	//
	// example:
	//
	// 9.2
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// 生成轮数 P50
	//
	// example:
	//
	// 0.8
	P50 *int32 `json:"P50,omitempty" xml:"P50,omitempty"`
	// 生成轮数 P90
	//
	// example:
	//
	// 5.9
	P90 *int32 `json:"P90,omitempty" xml:"P90,omitempty"`
}

func (s RLFlowTurns) String() string {
	return dara.Prettify(s)
}

func (s RLFlowTurns) GoString() string {
	return s.String()
}

func (s *RLFlowTurns) GetAvg() *float64 {
	return s.Avg
}

func (s *RLFlowTurns) GetCount() *int32 {
	return s.Count
}

func (s *RLFlowTurns) GetMax() *int32 {
	return s.Max
}

func (s *RLFlowTurns) GetP50() *int32 {
	return s.P50
}

func (s *RLFlowTurns) GetP90() *int32 {
	return s.P90
}

func (s *RLFlowTurns) SetAvg(v float64) *RLFlowTurns {
	s.Avg = &v
	return s
}

func (s *RLFlowTurns) SetCount(v int32) *RLFlowTurns {
	s.Count = &v
	return s
}

func (s *RLFlowTurns) SetMax(v int32) *RLFlowTurns {
	s.Max = &v
	return s
}

func (s *RLFlowTurns) SetP50(v int32) *RLFlowTurns {
	s.P50 = &v
	return s
}

func (s *RLFlowTurns) SetP90(v int32) *RLFlowTurns {
	s.P90 = &v
	return s
}

func (s *RLFlowTurns) Validate() error {
	return dara.Validate(s)
}
