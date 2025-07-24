// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAutoScalingPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *PutAutoScalingPolicyRequest
	GetClusterId() *string
	SetConstraints(v *ScalingConstraints) *PutAutoScalingPolicyRequest
	GetConstraints() *ScalingConstraints
	SetNodeGroupId(v string) *PutAutoScalingPolicyRequest
	GetNodeGroupId() *string
	SetRegionId(v string) *PutAutoScalingPolicyRequest
	GetRegionId() *string
	SetScalingRules(v []*ScalingRule) *PutAutoScalingPolicyRequest
	GetScalingRules() []*ScalingRule
}

type PutAutoScalingPolicyRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The constraints on the maximum and minimum numbers of nodes in a node group.
	Constraints *ScalingConstraints `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The ID of the node group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description list of auto scaling rules. Number of elements in the array: 0 to 100.
	ScalingRules []*ScalingRule `json:"ScalingRules,omitempty" xml:"ScalingRules,omitempty" type:"Repeated"`
}

func (s PutAutoScalingPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s PutAutoScalingPolicyRequest) GoString() string {
	return s.String()
}

func (s *PutAutoScalingPolicyRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *PutAutoScalingPolicyRequest) GetConstraints() *ScalingConstraints {
	return s.Constraints
}

func (s *PutAutoScalingPolicyRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *PutAutoScalingPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PutAutoScalingPolicyRequest) GetScalingRules() []*ScalingRule {
	return s.ScalingRules
}

func (s *PutAutoScalingPolicyRequest) SetClusterId(v string) *PutAutoScalingPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *PutAutoScalingPolicyRequest) SetConstraints(v *ScalingConstraints) *PutAutoScalingPolicyRequest {
	s.Constraints = v
	return s
}

func (s *PutAutoScalingPolicyRequest) SetNodeGroupId(v string) *PutAutoScalingPolicyRequest {
	s.NodeGroupId = &v
	return s
}

func (s *PutAutoScalingPolicyRequest) SetRegionId(v string) *PutAutoScalingPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *PutAutoScalingPolicyRequest) SetScalingRules(v []*ScalingRule) *PutAutoScalingPolicyRequest {
	s.ScalingRules = v
	return s
}

func (s *PutAutoScalingPolicyRequest) Validate() error {
	return dara.Validate(s)
}
