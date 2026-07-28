// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewResourceGroupId(v string) *ChangeResourceGroupRequest
	GetNewResourceGroupId() *string
	SetRegionId(v string) *ChangeResourceGroupRequest
	GetRegionId() *string
	SetResourceId(v string) *ChangeResourceGroupRequest
	GetResourceId() *string
	SetResourceType(v string) *ChangeResourceGroupRequest
	GetResourceType() *string
}

type ChangeResourceGroupRequest struct {
	// The ID of the new resource group.
	//
	// > A resource group is a mechanism for managing resources by group within an Alibaba Cloud account. Resource groups help you address the complexity of resource grouping and authorization management within a single cloud account. For more information, see [What is Resource Management](https://help.aliyun.com/document_detail/94475.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-bp1drpcfz9srr393h****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The region ID of the resource group that you want to modify.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Express Connect circuit resource whose resource group you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp16qjewdsunr41m1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid values:
	//
	// - **PHYSICALCONNECTION**: Express Connect circuit instance.
	//
	// - **VIRTUALBORDERROUTER**: Virtual Border Router.
	//
	// - **ROUTERINTERFACE**: VBR uplink.
	//
	// - **TRAFFICQOS**: QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// PHYSICALCONNECTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *ChangeResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ChangeResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChangeResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetRegionId(v string) *ChangeResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *ChangeResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
