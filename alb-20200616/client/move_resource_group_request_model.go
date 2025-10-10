// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewResourceGroupId(v string) *MoveResourceGroupRequest
	GetNewResourceGroupId() *string
	SetResourceId(v string) *MoveResourceGroupRequest
	GetResourceId() *string
	SetResourceType(v string) *MoveResourceGroupRequest
	GetResourceType() *string
}

type MoveResourceGroupRequest struct {
	// The ID of the resource group to which you want to transfer the cloud resource.
	//
	// >  You can use resource groups to manage resources within your Alibaba Cloud account by group. This helps you resolve issues such as resource grouping and permission management for your Alibaba Cloud account. For more information, see [What is resource management?](https://help.aliyun.com/document_detail/94475.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-9gLOoK****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-hp34s2h0xx1ht4nwo****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- **loadbalancer**: Application Load Balancer (ALB) instance
	//
	// 	- **acl**: access control list (ACL)
	//
	// 	- **securitypolicy**: security policy
	//
	// 	- **servergroup**: server group
	//
	// This parameter is required.
	//
	// example:
	//
	// ACL
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *MoveResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *MoveResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *MoveResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
