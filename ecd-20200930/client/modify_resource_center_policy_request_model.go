// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceCenterPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyGroupIds(v []*string) *ModifyResourceCenterPolicyRequest
	GetPolicyGroupIds() []*string
	SetPolicyGroupType(v string) *ModifyResourceCenterPolicyRequest
	GetPolicyGroupType() *string
	SetProductType(v string) *ModifyResourceCenterPolicyRequest
	GetProductType() *string
	SetResourceIds(v []*string) *ModifyResourceCenterPolicyRequest
	GetResourceIds() []*string
	SetResourceRegionId(v string) *ModifyResourceCenterPolicyRequest
	GetResourceRegionId() *string
	SetResourceType(v string) *ModifyResourceCenterPolicyRequest
	GetResourceType() *string
}

type ModifyResourceCenterPolicyRequest struct {
	// The IDs of the cloud computer policies that you want to associate with cloud computers.
	//
	// >  You can specify up to one cloud computer policy that takes effect globally, and up to four cloud computer policies that apply to specific IP addresses. If multiple cloud computer policies are configured for global enforcement, only the earliest-associated policy will take effect
	//
	// This parameter is required.
	PolicyGroupIds []*string `json:"PolicyGroupIds,omitempty" xml:"PolicyGroupIds,omitempty" type:"Repeated"`
	// The policy type.
	//
	// Valid values:
	//
	// 	- general: a general policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// general
	PolicyGroupType *string `json:"PolicyGroupType,omitempty" xml:"PolicyGroupType,omitempty"`
	// The service type.
	//
	// Valid values:
	//
	// 	- app: cloud applications.
	//
	// 	- resourceGroup: resource groups.
	//
	// 	- desktop: cloud computers.
	//
	// 	- desktopGroup: cloud computer shares.
	//
	// This parameter is required.
	//
	// example:
	//
	// desktop
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The resource IDs. You can specify up to 100 resource IDs.
	//
	// This parameter is required.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
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

func (s ModifyResourceCenterPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceCenterPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyResourceCenterPolicyRequest) GetPolicyGroupIds() []*string {
	return s.PolicyGroupIds
}

func (s *ModifyResourceCenterPolicyRequest) GetPolicyGroupType() *string {
	return s.PolicyGroupType
}

func (s *ModifyResourceCenterPolicyRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyResourceCenterPolicyRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *ModifyResourceCenterPolicyRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ModifyResourceCenterPolicyRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ModifyResourceCenterPolicyRequest) SetPolicyGroupIds(v []*string) *ModifyResourceCenterPolicyRequest {
	s.PolicyGroupIds = v
	return s
}

func (s *ModifyResourceCenterPolicyRequest) SetPolicyGroupType(v string) *ModifyResourceCenterPolicyRequest {
	s.PolicyGroupType = &v
	return s
}

func (s *ModifyResourceCenterPolicyRequest) SetProductType(v string) *ModifyResourceCenterPolicyRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyResourceCenterPolicyRequest) SetResourceIds(v []*string) *ModifyResourceCenterPolicyRequest {
	s.ResourceIds = v
	return s
}

func (s *ModifyResourceCenterPolicyRequest) SetResourceRegionId(v string) *ModifyResourceCenterPolicyRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *ModifyResourceCenterPolicyRequest) SetResourceType(v string) *ModifyResourceCenterPolicyRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyResourceCenterPolicyRequest) Validate() error {
	return dara.Validate(s)
}
