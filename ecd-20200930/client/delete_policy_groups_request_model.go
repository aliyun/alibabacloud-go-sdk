// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyGroupId(v []*string) *DeletePolicyGroupsRequest
	GetPolicyGroupId() []*string
	SetRegionId(v string) *DeletePolicyGroupsRequest
	GetRegionId() *string
}

type DeletePolicyGroupsRequest struct {
	// The cloud computer policy IDs. You can specify 1 to 100 policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-gx2x1dhsmthe9****
	PolicyGroupId []*string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions supported by EDS.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePolicyGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupsRequest) GetPolicyGroupId() []*string {
	return s.PolicyGroupId
}

func (s *DeletePolicyGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePolicyGroupsRequest) SetPolicyGroupId(v []*string) *DeletePolicyGroupsRequest {
	s.PolicyGroupId = v
	return s
}

func (s *DeletePolicyGroupsRequest) SetRegionId(v string) *DeletePolicyGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePolicyGroupsRequest) Validate() error {
	return dara.Validate(s)
}
