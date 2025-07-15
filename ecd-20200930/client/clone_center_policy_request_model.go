// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneCenterPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v int32) *CloneCenterPolicyRequest
	GetBusinessType() *int32
	SetName(v string) *CloneCenterPolicyRequest
	GetName() *string
	SetPolicyGroupId(v string) *CloneCenterPolicyRequest
	GetPolicyGroupId() *string
	SetRegionId(v string) *CloneCenterPolicyRequest
	GetRegionId() *string
	SetResourceType(v string) *CloneCenterPolicyRequest
	GetResourceType() *string
}

type CloneCenterPolicyRequest struct {
	// The business type.
	//
	// Valid values:
	//
	// 	- 1: public cloud
	//
	// 	- 8: commercial edition.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The name of the cloud computer policy that you want to clone.
	//
	// This parameter is required.
	//
	// example:
	//
	// testPolicyGroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the cloud computer policy that you want to clone.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-gx2x1dhsmthe9****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
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

func (s CloneCenterPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneCenterPolicyRequest) GoString() string {
	return s.String()
}

func (s *CloneCenterPolicyRequest) GetBusinessType() *int32 {
	return s.BusinessType
}

func (s *CloneCenterPolicyRequest) GetName() *string {
	return s.Name
}

func (s *CloneCenterPolicyRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *CloneCenterPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CloneCenterPolicyRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CloneCenterPolicyRequest) SetBusinessType(v int32) *CloneCenterPolicyRequest {
	s.BusinessType = &v
	return s
}

func (s *CloneCenterPolicyRequest) SetName(v string) *CloneCenterPolicyRequest {
	s.Name = &v
	return s
}

func (s *CloneCenterPolicyRequest) SetPolicyGroupId(v string) *CloneCenterPolicyRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CloneCenterPolicyRequest) SetRegionId(v string) *CloneCenterPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CloneCenterPolicyRequest) SetResourceType(v string) *CloneCenterPolicyRequest {
	s.ResourceType = &v
	return s
}

func (s *CloneCenterPolicyRequest) Validate() error {
	return dara.Validate(s)
}
