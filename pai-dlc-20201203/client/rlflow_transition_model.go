// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLFlowTransition interface {
	dara.Model
	String() string
	GoString() string
	SetAvg(v float64) *RLFlowTransition
	GetAvg() *float64
	SetCount(v int32) *RLFlowTransition
	GetCount() *int32
	SetKey(v string) *RLFlowTransition
	GetKey() *string
	SetLabel(v string) *RLFlowTransition
	GetLabel() *string
	SetMax(v float64) *RLFlowTransition
	GetMax() *float64
	SetP50(v float64) *RLFlowTransition
	GetP50() *float64
	SetP90(v float64) *RLFlowTransition
	GetP90() *float64
	SetP99(v float64) *RLFlowTransition
	GetP99() *float64
	SetSlowest(v []*RLFlowSlowestItem) *RLFlowTransition
	GetSlowest() []*RLFlowSlowestItem
}

type RLFlowTransition struct {
	// The average duration in seconds.
	//
	// example:
	//
	// 1.9
	Avg *float64 `json:"Avg,omitempty" xml:"Avg,omitempty"`
	// The number of trajectories included in the statistics.
	//
	// example:
	//
	// 96
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The phase identifier. Valid values: dispatch_wait, start_wait, env_prepare, generation, agent_finish, reward, buffer_wait, logprob, ref_logprob, advantage, update, and e2e.
	//
	// example:
	//
	// traj
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The Chinese name of the phase.
	//
	// example:
	//
	// 生成轨迹
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The maximum duration in seconds.
	//
	// example:
	//
	// 9.2
	Max *float64 `json:"Max,omitempty" xml:"Max,omitempty"`
	// The P50 duration in seconds.
	//
	// example:
	//
	// 0.8
	P50 *float64 `json:"P50,omitempty" xml:"P50,omitempty"`
	// The P90 duration in seconds.
	//
	// example:
	//
	// 5.9
	P90 *float64 `json:"P90,omitempty" xml:"P90,omitempty"`
	// The P99 duration in seconds.
	//
	// example:
	//
	// 8.6
	P99 *float64 `json:"P99,omitempty" xml:"P99,omitempty"`
	// The slowest 5 trajectories.
	//
	// example:
	//
	// [{"PromptUid":"321fa56f-e1e5-4eb3-8047-db7a230c9a75","SampleIndex":"2","Sec":9.2}]
	Slowest []*RLFlowSlowestItem `json:"Slowest,omitempty" xml:"Slowest,omitempty" type:"Repeated"`
}

func (s RLFlowTransition) String() string {
	return dara.Prettify(s)
}

func (s RLFlowTransition) GoString() string {
	return s.String()
}

func (s *RLFlowTransition) GetAvg() *float64 {
	return s.Avg
}

func (s *RLFlowTransition) GetCount() *int32 {
	return s.Count
}

func (s *RLFlowTransition) GetKey() *string {
	return s.Key
}

func (s *RLFlowTransition) GetLabel() *string {
	return s.Label
}

func (s *RLFlowTransition) GetMax() *float64 {
	return s.Max
}

func (s *RLFlowTransition) GetP50() *float64 {
	return s.P50
}

func (s *RLFlowTransition) GetP90() *float64 {
	return s.P90
}

func (s *RLFlowTransition) GetP99() *float64 {
	return s.P99
}

func (s *RLFlowTransition) GetSlowest() []*RLFlowSlowestItem {
	return s.Slowest
}

func (s *RLFlowTransition) SetAvg(v float64) *RLFlowTransition {
	s.Avg = &v
	return s
}

func (s *RLFlowTransition) SetCount(v int32) *RLFlowTransition {
	s.Count = &v
	return s
}

func (s *RLFlowTransition) SetKey(v string) *RLFlowTransition {
	s.Key = &v
	return s
}

func (s *RLFlowTransition) SetLabel(v string) *RLFlowTransition {
	s.Label = &v
	return s
}

func (s *RLFlowTransition) SetMax(v float64) *RLFlowTransition {
	s.Max = &v
	return s
}

func (s *RLFlowTransition) SetP50(v float64) *RLFlowTransition {
	s.P50 = &v
	return s
}

func (s *RLFlowTransition) SetP90(v float64) *RLFlowTransition {
	s.P90 = &v
	return s
}

func (s *RLFlowTransition) SetP99(v float64) *RLFlowTransition {
	s.P99 = &v
	return s
}

func (s *RLFlowTransition) SetSlowest(v []*RLFlowSlowestItem) *RLFlowTransition {
	s.Slowest = v
	return s
}

func (s *RLFlowTransition) Validate() error {
	if s.Slowest != nil {
		for _, item := range s.Slowest {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
