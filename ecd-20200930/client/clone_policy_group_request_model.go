// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClonePolicyGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ClonePolicyGroupRequest
	GetName() *string
	SetPolicyGroupId(v string) *ClonePolicyGroupRequest
	GetPolicyGroupId() *string
	SetRegionId(v string) *ClonePolicyGroupRequest
	GetRegionId() *string
}

type ClonePolicyGroupRequest struct {
	// The name of the cloud computer policy that you want to create.
	//
	// This parameter is required.
	//
	// example:
	//
	// testPolicyGroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the destination cloud computer policy that you want to clone.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-gx2x1dhsmthe9****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions supported by Elastic Desktop Service (EDS).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ClonePolicyGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ClonePolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ClonePolicyGroupRequest) GetName() *string {
	return s.Name
}

func (s *ClonePolicyGroupRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *ClonePolicyGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ClonePolicyGroupRequest) SetName(v string) *ClonePolicyGroupRequest {
	s.Name = &v
	return s
}

func (s *ClonePolicyGroupRequest) SetPolicyGroupId(v string) *ClonePolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ClonePolicyGroupRequest) SetRegionId(v string) *ClonePolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ClonePolicyGroupRequest) Validate() error {
	return dara.Validate(s)
}
