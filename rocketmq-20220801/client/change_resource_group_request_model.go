// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ChangeResourceGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ChangeResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *ChangeResourceGroupRequest
	GetResourceId() *string
	SetResourceType(v string) *ChangeResourceGroupRequest
	GetResourceType() *string
}

type ChangeResourceGroupRequest struct {
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The ID of the resource group to which you want to transfer the specified instance.
	//
	// You can call the <props="china">[ListResourceGroups](https://help.aliyun.com/document_detail/158855.html)<props="intl">[ListResourceGroups](https://www.alibabacloud.com/help/resource-management/latest/listresourcegroups) operation to query existing resource groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-9gLOoK****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The resource ID, which is the ID of the ApsaraMQ for RocketMQ instance. This parameter specifies the MSMQ queue instance whose resource group you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// c2c5d1274a8d4317a13bc5b0d4******
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type.
	//
	// Set the value to **instance**, which indicates that the resource is an instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ChangeResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ChangeResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChangeResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ChangeResourceGroupRequest) SetRegionId(v string) *ChangeResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.ResourceGroupId = &v
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
