// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GetWorkspaceShrinkRequest
	GetTenantContextShrink() *string
	SetWithPermissionRole(v bool) *GetWorkspaceShrinkRequest
	GetWithPermissionRole() *bool
	SetWorkspaceId(v string) *GetWorkspaceShrinkRequest
	GetWorkspaceId() *string
}

type GetWorkspaceShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// false
	WithPermissionRole *bool `json:"WithPermissionRole,omitempty" xml:"WithPermissionRole,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MJ0pDSKMV9dO20E4
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetWorkspaceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetWorkspaceShrinkRequest) GetWithPermissionRole() *bool {
	return s.WithPermissionRole
}

func (s *GetWorkspaceShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspaceShrinkRequest) SetTenantContextShrink(v string) *GetWorkspaceShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetWorkspaceShrinkRequest) SetWithPermissionRole(v bool) *GetWorkspaceShrinkRequest {
	s.WithPermissionRole = &v
	return s
}

func (s *GetWorkspaceShrinkRequest) SetWorkspaceId(v string) *GetWorkspaceShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
