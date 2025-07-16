// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateWorkspaceShrinkRequest
	GetDescription() *string
	SetName(v string) *CreateWorkspaceShrinkRequest
	GetName() *string
	SetTenantContextShrink(v string) *CreateWorkspaceShrinkRequest
	GetTenantContextShrink() *string
}

type CreateWorkspaceShrinkRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s CreateWorkspaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkspaceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateWorkspaceShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateWorkspaceShrinkRequest) SetDescription(v string) *CreateWorkspaceShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceShrinkRequest) SetName(v string) *CreateWorkspaceShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceShrinkRequest) SetTenantContextShrink(v string) *CreateWorkspaceShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateWorkspaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
