// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenterPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v int32) *DeleteCenterPolicyRequest
	GetBusinessType() *int32
	SetPolicyGroupIds(v []*string) *DeleteCenterPolicyRequest
	GetPolicyGroupIds() []*string
	SetRegionId(v string) *DeleteCenterPolicyRequest
	GetRegionId() *string
	SetResourceType(v string) *DeleteCenterPolicyRequest
	GetResourceType() *string
}

type DeleteCenterPolicyRequest struct {
	// The business type.
	//
	// Valid values:
	//
	// 	- 1: public cloud.
	//
	// 	- 8: commercial edition.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The policy IDs.
	//
	// This parameter is required.
	PolicyGroupIds []*string `json:"PolicyGroupIds,omitempty" xml:"PolicyGroupIds,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- app: cloud applications.
	//
	// 	- desktop: cloud computers.
	//
	// This parameter is required.
	//
	// example:
	//
	// desktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DeleteCenterPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenterPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenterPolicyRequest) GetBusinessType() *int32 {
	return s.BusinessType
}

func (s *DeleteCenterPolicyRequest) GetPolicyGroupIds() []*string {
	return s.PolicyGroupIds
}

func (s *DeleteCenterPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCenterPolicyRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteCenterPolicyRequest) SetBusinessType(v int32) *DeleteCenterPolicyRequest {
	s.BusinessType = &v
	return s
}

func (s *DeleteCenterPolicyRequest) SetPolicyGroupIds(v []*string) *DeleteCenterPolicyRequest {
	s.PolicyGroupIds = v
	return s
}

func (s *DeleteCenterPolicyRequest) SetRegionId(v string) *DeleteCenterPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCenterPolicyRequest) SetResourceType(v string) *DeleteCenterPolicyRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteCenterPolicyRequest) Validate() error {
	return dara.Validate(s)
}
