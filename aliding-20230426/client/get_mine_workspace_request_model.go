// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMineWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequest(v map[string]interface{}) *GetMineWorkspaceRequest
	GetRequest() map[string]interface{}
	SetTenantContext(v *GetMineWorkspaceRequestTenantContext) *GetMineWorkspaceRequest
	GetTenantContext() *GetMineWorkspaceRequestTenantContext
}

type GetMineWorkspaceRequest struct {
	Request       map[string]interface{}                `json:"Request,omitempty" xml:"Request,omitempty"`
	TenantContext *GetMineWorkspaceRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetMineWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMineWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceRequest) GetRequest() map[string]interface{} {
	return s.Request
}

func (s *GetMineWorkspaceRequest) GetTenantContext() *GetMineWorkspaceRequestTenantContext {
	return s.TenantContext
}

func (s *GetMineWorkspaceRequest) SetRequest(v map[string]interface{}) *GetMineWorkspaceRequest {
	s.Request = v
	return s
}

func (s *GetMineWorkspaceRequest) SetTenantContext(v *GetMineWorkspaceRequestTenantContext) *GetMineWorkspaceRequest {
	s.TenantContext = v
	return s
}

func (s *GetMineWorkspaceRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMineWorkspaceRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetMineWorkspaceRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetMineWorkspaceRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMineWorkspaceRequestTenantContext) SetTenantId(v string) *GetMineWorkspaceRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetMineWorkspaceRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
