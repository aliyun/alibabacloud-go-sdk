// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOption(v *GetWorkspacesRequestOption) *GetWorkspacesRequest
	GetOption() *GetWorkspacesRequestOption
	SetTenantContext(v *GetWorkspacesRequestTenantContext) *GetWorkspacesRequest
	GetTenantContext() *GetWorkspacesRequestTenantContext
	SetWorkspaceIds(v []*string) *GetWorkspacesRequest
	GetWorkspaceIds() []*string
}

type GetWorkspacesRequest struct {
	Option        *GetWorkspacesRequestOption        `json:"Option,omitempty" xml:"Option,omitempty" type:"Struct"`
	TenantContext *GetWorkspacesRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// workspace_id
	WorkspaceIds []*string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty" type:"Repeated"`
}

func (s GetWorkspacesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspacesRequest) GetOption() *GetWorkspacesRequestOption {
	return s.Option
}

func (s *GetWorkspacesRequest) GetTenantContext() *GetWorkspacesRequestTenantContext {
	return s.TenantContext
}

func (s *GetWorkspacesRequest) GetWorkspaceIds() []*string {
	return s.WorkspaceIds
}

func (s *GetWorkspacesRequest) SetOption(v *GetWorkspacesRequestOption) *GetWorkspacesRequest {
	s.Option = v
	return s
}

func (s *GetWorkspacesRequest) SetTenantContext(v *GetWorkspacesRequestTenantContext) *GetWorkspacesRequest {
	s.TenantContext = v
	return s
}

func (s *GetWorkspacesRequest) SetWorkspaceIds(v []*string) *GetWorkspacesRequest {
	s.WorkspaceIds = v
	return s
}

func (s *GetWorkspacesRequest) Validate() error {
	if s.Option != nil {
		if err := s.Option.Validate(); err != nil {
			return err
		}
	}
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspacesRequestOption struct {
	// example:
	//
	// true
	WithPermissionRole *bool `json:"WithPermissionRole,omitempty" xml:"WithPermissionRole,omitempty"`
}

func (s GetWorkspacesRequestOption) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacesRequestOption) GoString() string {
	return s.String()
}

func (s *GetWorkspacesRequestOption) GetWithPermissionRole() *bool {
	return s.WithPermissionRole
}

func (s *GetWorkspacesRequestOption) SetWithPermissionRole(v bool) *GetWorkspacesRequestOption {
	s.WithPermissionRole = &v
	return s
}

func (s *GetWorkspacesRequestOption) Validate() error {
	return dara.Validate(s)
}

type GetWorkspacesRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetWorkspacesRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacesRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetWorkspacesRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkspacesRequestTenantContext) SetTenantId(v string) *GetWorkspacesRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetWorkspacesRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
