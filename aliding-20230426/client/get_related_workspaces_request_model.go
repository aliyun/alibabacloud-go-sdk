// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRelatedWorkspacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeRecent(v bool) *GetRelatedWorkspacesRequest
	GetIncludeRecent() *bool
	SetTenantContext(v *GetRelatedWorkspacesRequestTenantContext) *GetRelatedWorkspacesRequest
	GetTenantContext() *GetRelatedWorkspacesRequestTenantContext
}

type GetRelatedWorkspacesRequest struct {
	// example:
	//
	// true
	IncludeRecent *bool                                     `json:"IncludeRecent,omitempty" xml:"IncludeRecent,omitempty"`
	TenantContext *GetRelatedWorkspacesRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetRelatedWorkspacesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRelatedWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesRequest) GetIncludeRecent() *bool {
	return s.IncludeRecent
}

func (s *GetRelatedWorkspacesRequest) GetTenantContext() *GetRelatedWorkspacesRequestTenantContext {
	return s.TenantContext
}

func (s *GetRelatedWorkspacesRequest) SetIncludeRecent(v bool) *GetRelatedWorkspacesRequest {
	s.IncludeRecent = &v
	return s
}

func (s *GetRelatedWorkspacesRequest) SetTenantContext(v *GetRelatedWorkspacesRequestTenantContext) *GetRelatedWorkspacesRequest {
	s.TenantContext = v
	return s
}

func (s *GetRelatedWorkspacesRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRelatedWorkspacesRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetRelatedWorkspacesRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetRelatedWorkspacesRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetRelatedWorkspacesRequestTenantContext) SetTenantId(v string) *GetRelatedWorkspacesRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetRelatedWorkspacesRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
