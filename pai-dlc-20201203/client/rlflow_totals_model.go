// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLFlowTotals interface {
	dara.Model
	String() string
	GoString() string
	SetInflight(v int32) *RLFlowTotals
	GetInflight() *int32
	SetRewarded(v int32) *RLFlowTotals
	GetRewarded() *int32
	SetSampled(v int32) *RLFlowTotals
	GetSampled() *int32
	SetTrained(v int32) *RLFlowTotals
	GetTrained() *int32
	SetTrajs(v int32) *RLFlowTotals
	GetTrajs() *int32
	SetUids(v int32) *RLFlowTotals
	GetUids() *int32
}

type RLFlowTotals struct {
	// The number of in-flight trajectories (no desired state).
	//
	// example:
	//
	// 0
	Inflight *int32 `json:"Inflight,omitempty" xml:"Inflight,omitempty"`
	// The number of trajectories that have completed reward scoring (hit reward_score_computed).
	//
	// example:
	//
	// 96
	Rewarded *int32 `json:"Rewarded,omitempty" xml:"Rewarded,omitempty"`
	// The number of trajectories sampled into a batch by the trainer (hit sampled_from_replay_buffer).
	//
	// example:
	//
	// 96
	Sampled *int32 `json:"Sampled,omitempty" xml:"Sampled,omitempty"`
	// The number of trajectories that have completed training (hit actor_parameters_updated).
	//
	// example:
	//
	// 96
	Trained *int32 `json:"Trained,omitempty" xml:"Trained,omitempty"`
	// The total number of trajectories in the window.
	//
	// example:
	//
	// 96
	Trajs *int32 `json:"Trajs,omitempty" xml:"Trajs,omitempty"`
	// The number of sample UIDs that appear in the window.
	//
	// example:
	//
	// 24
	Uids *int32 `json:"Uids,omitempty" xml:"Uids,omitempty"`
}

func (s RLFlowTotals) String() string {
	return dara.Prettify(s)
}

func (s RLFlowTotals) GoString() string {
	return s.String()
}

func (s *RLFlowTotals) GetInflight() *int32 {
	return s.Inflight
}

func (s *RLFlowTotals) GetRewarded() *int32 {
	return s.Rewarded
}

func (s *RLFlowTotals) GetSampled() *int32 {
	return s.Sampled
}

func (s *RLFlowTotals) GetTrained() *int32 {
	return s.Trained
}

func (s *RLFlowTotals) GetTrajs() *int32 {
	return s.Trajs
}

func (s *RLFlowTotals) GetUids() *int32 {
	return s.Uids
}

func (s *RLFlowTotals) SetInflight(v int32) *RLFlowTotals {
	s.Inflight = &v
	return s
}

func (s *RLFlowTotals) SetRewarded(v int32) *RLFlowTotals {
	s.Rewarded = &v
	return s
}

func (s *RLFlowTotals) SetSampled(v int32) *RLFlowTotals {
	s.Sampled = &v
	return s
}

func (s *RLFlowTotals) SetTrained(v int32) *RLFlowTotals {
	s.Trained = &v
	return s
}

func (s *RLFlowTotals) SetTrajs(v int32) *RLFlowTotals {
	s.Trajs = &v
	return s
}

func (s *RLFlowTotals) SetUids(v int32) *RLFlowTotals {
	s.Uids = &v
	return s
}

func (s *RLFlowTotals) Validate() error {
	return dara.Validate(s)
}
