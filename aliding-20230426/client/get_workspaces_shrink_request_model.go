// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspacesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOptionShrink(v string) *GetWorkspacesShrinkRequest
	GetOptionShrink() *string
	SetTenantContextShrink(v string) *GetWorkspacesShrinkRequest
	GetTenantContextShrink() *string
	SetWorkspaceIdsShrink(v string) *GetWorkspacesShrinkRequest
	GetWorkspaceIdsShrink() *string
}

type GetWorkspacesShrinkRequest struct {
	OptionShrink        *string `json:"Option,omitempty" xml:"Option,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// workspace_id
	WorkspaceIdsShrink *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
}

func (s GetWorkspacesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspacesShrinkRequest) GetOptionShrink() *string {
	return s.OptionShrink
}

func (s *GetWorkspacesShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetWorkspacesShrinkRequest) GetWorkspaceIdsShrink() *string {
	return s.WorkspaceIdsShrink
}

func (s *GetWorkspacesShrinkRequest) SetOptionShrink(v string) *GetWorkspacesShrinkRequest {
	s.OptionShrink = &v
	return s
}

func (s *GetWorkspacesShrinkRequest) SetTenantContextShrink(v string) *GetWorkspacesShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetWorkspacesShrinkRequest) SetWorkspaceIdsShrink(v string) *GetWorkspacesShrinkRequest {
	s.WorkspaceIdsShrink = &v
	return s
}

func (s *GetWorkspacesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
