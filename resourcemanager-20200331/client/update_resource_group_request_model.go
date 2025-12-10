// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewDisplayName(v string) *UpdateResourceGroupRequest
	GetNewDisplayName() *string
	SetResourceGroupId(v string) *UpdateResourceGroupRequest
	GetResourceGroupId() *string
}

type UpdateResourceGroupRequest struct {
	// The display name of the resource group.
	//
	// The name must be 1 to 50 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// project
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	// The ID of the resource group.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-9gLOoK****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupRequest) GetNewDisplayName() *string {
	return s.NewDisplayName
}

func (s *UpdateResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateResourceGroupRequest) SetNewDisplayName(v string) *UpdateResourceGroupRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateResourceGroupRequest) SetResourceGroupId(v string) *UpdateResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
