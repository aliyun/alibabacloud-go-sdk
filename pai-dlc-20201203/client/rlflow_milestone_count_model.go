// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLFlowMilestoneCount interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *RLFlowMilestoneCount
	GetCount() *int32
	SetMilestone(v string) *RLFlowMilestoneCount
	GetMilestone() *string
}

type RLFlowMilestoneCount struct {
	// The number of in-transit trajectories that remain at this milestone.
	//
	// example:
	//
	// 96
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The milestone. Valid values are the same as those of Stuck[].Milestone.
	//
	// example:
	//
	// 生成中
	Milestone *string `json:"Milestone,omitempty" xml:"Milestone,omitempty"`
}

func (s RLFlowMilestoneCount) String() string {
	return dara.Prettify(s)
}

func (s RLFlowMilestoneCount) GoString() string {
	return s.String()
}

func (s *RLFlowMilestoneCount) GetCount() *int32 {
	return s.Count
}

func (s *RLFlowMilestoneCount) GetMilestone() *string {
	return s.Milestone
}

func (s *RLFlowMilestoneCount) SetCount(v int32) *RLFlowMilestoneCount {
	s.Count = &v
	return s
}

func (s *RLFlowMilestoneCount) SetMilestone(v string) *RLFlowMilestoneCount {
	s.Milestone = &v
	return s
}

func (s *RLFlowMilestoneCount) Validate() error {
	return dara.Validate(s)
}
