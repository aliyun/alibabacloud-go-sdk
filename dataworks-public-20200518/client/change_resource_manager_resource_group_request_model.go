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
	// The ID of the resource. The value of this parameter depends on the value of the ResourceType parameter:
	//
	// - If ResourceType is set to project, this parameter specifies the name of the workspace (ProjectIdentifier). You can call the [ListProjects](https://help.aliyun.com/document_detail/2780068.html) operation to obtain the workspace name.
	//
	// - If ResourceType is set to tenantresourcegroup, this parameter specifies the identifier of the exclusive resource group (Identifier). You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/2780075.html) operation to obtain the identifier. This applies only to resource groups of type 7, 8, or 9.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the destination resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The resource type. Valid values:
	//
	// - project: a workspace. Select this value to change the resource group for a DataWorks edition.
	//
	// - tenantresourcegroup: an exclusive resource group. Select this value to change the resource group for a DataWorks exclusive resource group.
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
