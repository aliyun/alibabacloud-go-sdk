// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *AddWorkspaceShrinkRequest
	GetName() *string
	SetOptionShrink(v string) *AddWorkspaceShrinkRequest
	GetOptionShrink() *string
	SetTenantContextShrink(v string) *AddWorkspaceShrinkRequest
	GetTenantContextShrink() *string
}

type AddWorkspaceShrinkRequest struct {
	// This parameter is required.
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OptionShrink        *string `json:"Option,omitempty" xml:"Option,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s AddWorkspaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *AddWorkspaceShrinkRequest) GetOptionShrink() *string {
	return s.OptionShrink
}

func (s *AddWorkspaceShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *AddWorkspaceShrinkRequest) SetName(v string) *AddWorkspaceShrinkRequest {
	s.Name = &v
	return s
}

func (s *AddWorkspaceShrinkRequest) SetOptionShrink(v string) *AddWorkspaceShrinkRequest {
	s.OptionShrink = &v
	return s
}

func (s *AddWorkspaceShrinkRequest) SetTenantContextShrink(v string) *AddWorkspaceShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *AddWorkspaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
