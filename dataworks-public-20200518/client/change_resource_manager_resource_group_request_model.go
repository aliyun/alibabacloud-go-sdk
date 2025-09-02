// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceManagerResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v string) *ChangeResourceManagerResourceGroupRequest
	GetResourceId() *string
	SetResourceManagerResourceGroupId(v string) *ChangeResourceManagerResourceGroupRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceType(v string) *ChangeResourceManagerResourceGroupRequest
	GetResourceType() *string
}

type ChangeResourceManagerResourceGroupRequest struct {
	// The ID of the resource type.
	//
	// 	- If you set ResourceType to project, set this parameter to the value of ProjectIdentifier. You can call the [ListProjects](https://help.aliyun.com/document_detail/2780068.html) operation to obtain the value of ProjectIdentifier.
	//
	// 	- If you set ResourceType to tenantresourcegroup, set this parameter to the value of ResourceGroupType. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/2780075.html) operation to obtain the value of ResourceGroupType. Only the values 7, 8, and 9 are valid.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the new resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- project: workspace. If you want to change the resource group that you specify when you activate DataWorks, set the value to project.
	//
	// 	- tenantresourcegroup: exclusive resource group. If you want to change the resource group that you specify when you purchase a DataWorks exclusive resource group, set the value to tenantresourcegroup.
	//
	// This parameter is required.
	//
	// example:
	//
	// project
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceManagerResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceManagerResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceManagerResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChangeResourceManagerResourceGroupRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ChangeResourceManagerResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ChangeResourceManagerResourceGroupRequest) SetResourceId(v string) *ChangeResourceManagerResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceManagerResourceGroupRequest) SetResourceManagerResourceGroupId(v string) *ChangeResourceManagerResourceGroupRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ChangeResourceManagerResourceGroupRequest) SetResourceType(v string) *ChangeResourceManagerResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *ChangeResourceManagerResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
