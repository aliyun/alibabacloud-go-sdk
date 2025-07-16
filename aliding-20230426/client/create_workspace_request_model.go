// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateWorkspaceRequest
	GetDescription() *string
	SetName(v string) *CreateWorkspaceRequest
	GetName() *string
	SetTenantContext(v *CreateWorkspaceRequestTenantContext) *CreateWorkspaceRequest
	GetTenantContext() *CreateWorkspaceRequestTenantContext
}

type CreateWorkspaceRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name          *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContext *CreateWorkspaceRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s CreateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkspaceRequest) GetName() *string {
	return s.Name
}

func (s *CreateWorkspaceRequest) GetTenantContext() *CreateWorkspaceRequestTenantContext {
	return s.TenantContext
}

func (s *CreateWorkspaceRequest) SetDescription(v string) *CreateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceRequest) SetName(v string) *CreateWorkspaceRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceRequest) SetTenantContext(v *CreateWorkspaceRequestTenantContext) *CreateWorkspaceRequest {
	s.TenantContext = v
	return s
}

func (s *CreateWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateWorkspaceRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateWorkspaceRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateWorkspaceRequestTenantContext) SetTenantId(v string) *CreateWorkspaceRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateWorkspaceRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
