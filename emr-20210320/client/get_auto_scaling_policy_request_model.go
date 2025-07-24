// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetAutoScalingPolicyRequest
	GetClusterId() *string
	SetNodeGroupId(v string) *GetAutoScalingPolicyRequest
	GetNodeGroupId() *string
	SetRegionId(v string) *GetAutoScalingPolicyRequest
	GetRegionId() *string
}

type GetAutoScalingPolicyRequest struct {
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

func (s GetAutoScalingPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetAutoScalingPolicyRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetAutoScalingPolicyRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *GetAutoScalingPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAutoScalingPolicyRequest) SetClusterId(v string) *GetAutoScalingPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *GetAutoScalingPolicyRequest) SetNodeGroupId(v string) *GetAutoScalingPolicyRequest {
	s.NodeGroupId = &v
	return s
}

func (s *GetAutoScalingPolicyRequest) SetRegionId(v string) *GetAutoScalingPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GetAutoScalingPolicyRequest) Validate() error {
	return dara.Validate(s)
}
