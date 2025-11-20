// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GetWorkspaceRequestTenantContext) *GetWorkspaceRequest
	GetTenantContext() *GetWorkspaceRequestTenantContext
	SetWithPermissionRole(v bool) *GetWorkspaceRequest
	GetWithPermissionRole() *bool
	SetWorkspaceId(v string) *GetWorkspaceRequest
	GetWorkspaceId() *string
}

type GetWorkspaceRequest struct {
	TenantContext *GetWorkspaceRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
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

func (s GetWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRequest) GetTenantContext() *GetWorkspaceRequestTenantContext {
	return s.TenantContext
}

func (s *GetWorkspaceRequest) GetWithPermissionRole() *bool {
	return s.WithPermissionRole
}

func (s *GetWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspaceRequest) SetTenantContext(v *GetWorkspaceRequestTenantContext) *GetWorkspaceRequest {
	s.TenantContext = v
	return s
}

func (s *GetWorkspaceRequest) SetWithPermissionRole(v bool) *GetWorkspaceRequest {
	s.WithPermissionRole = &v
	return s
}

func (s *GetWorkspaceRequest) SetWorkspaceId(v string) *GetWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspaceRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetWorkspaceRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkspaceRequestTenantContext) SetTenantId(v string) *GetWorkspaceRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetWorkspaceRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
