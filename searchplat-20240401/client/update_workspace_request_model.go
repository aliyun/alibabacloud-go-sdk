// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateWorkspaceRequest
	GetDescription() *string
	SetResourceGroupId(v string) *UpdateWorkspaceRequest
	GetResourceGroupId() *string
}

type UpdateWorkspaceRequest struct {
	// The description.
	//
	// example:
	//
	// test2024106271022
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzivjfrlpyn3y
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s UpdateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkspaceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateWorkspaceRequest) SetDescription(v string) *UpdateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetResourceGroupId(v string) *UpdateWorkspaceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
