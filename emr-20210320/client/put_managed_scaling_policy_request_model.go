// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutManagedScalingPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *PutManagedScalingPolicyRequest
	GetClusterId() *string
	SetConstraints(v *ManagedScalingConstraints) *PutManagedScalingPolicyRequest
	GetConstraints() *ManagedScalingConstraints
	SetRegionId(v string) *PutManagedScalingPolicyRequest
	GetRegionId() *string
}

type PutManagedScalingPolicyRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The constrains on the maximum and minimum numbers of nodes in a node group.
	Constraints *ManagedScalingConstraints `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s PutManagedScalingPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s PutManagedScalingPolicyRequest) GoString() string {
	return s.String()
}

func (s *PutManagedScalingPolicyRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *PutManagedScalingPolicyRequest) GetConstraints() *ManagedScalingConstraints {
	return s.Constraints
}

func (s *PutManagedScalingPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PutManagedScalingPolicyRequest) SetClusterId(v string) *PutManagedScalingPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *PutManagedScalingPolicyRequest) SetConstraints(v *ManagedScalingConstraints) *PutManagedScalingPolicyRequest {
	s.Constraints = v
	return s
}

func (s *PutManagedScalingPolicyRequest) SetRegionId(v string) *PutManagedScalingPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *PutManagedScalingPolicyRequest) Validate() error {
	return dara.Validate(s)
}
