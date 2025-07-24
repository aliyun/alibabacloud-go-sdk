// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAutoScalingPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *RemoveAutoScalingPolicyRequest
	GetClusterId() *string
	SetNodeGroupId(v string) *RemoveAutoScalingPolicyRequest
	GetNodeGroupId() *string
	SetRegionId(v string) *RemoveAutoScalingPolicyRequest
	GetRegionId() *string
}

type RemoveAutoScalingPolicyRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
}

func (s RemoveAutoScalingPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveAutoScalingPolicyRequest) GoString() string {
	return s.String()
}

func (s *RemoveAutoScalingPolicyRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RemoveAutoScalingPolicyRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *RemoveAutoScalingPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveAutoScalingPolicyRequest) SetClusterId(v string) *RemoveAutoScalingPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *RemoveAutoScalingPolicyRequest) SetNodeGroupId(v string) *RemoveAutoScalingPolicyRequest {
	s.NodeGroupId = &v
	return s
}

func (s *RemoveAutoScalingPolicyRequest) SetRegionId(v string) *RemoveAutoScalingPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveAutoScalingPolicyRequest) Validate() error {
	return dara.Validate(s)
}
