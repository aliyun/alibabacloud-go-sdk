// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByOrgIdAndStaffIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrgId(v int64) *GetUserIdByOrgIdAndStaffIdRequest
	GetOrgId() *int64
	SetTenantContext(v *GetUserIdByOrgIdAndStaffIdRequestTenantContext) *GetUserIdByOrgIdAndStaffIdRequest
	GetTenantContext() *GetUserIdByOrgIdAndStaffIdRequestTenantContext
}

type GetUserIdByOrgIdAndStaffIdRequest struct {
	// example:
	//
	// 123456
	OrgId         *int64                                          `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	TenantContext *GetUserIdByOrgIdAndStaffIdRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetUserIdByOrgIdAndStaffIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOrgIdAndStaffIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByOrgIdAndStaffIdRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *GetUserIdByOrgIdAndStaffIdRequest) GetTenantContext() *GetUserIdByOrgIdAndStaffIdRequestTenantContext {
	return s.TenantContext
}

func (s *GetUserIdByOrgIdAndStaffIdRequest) SetOrgId(v int64) *GetUserIdByOrgIdAndStaffIdRequest {
	s.OrgId = &v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdRequest) SetTenantContext(v *GetUserIdByOrgIdAndStaffIdRequestTenantContext) *GetUserIdByOrgIdAndStaffIdRequest {
	s.TenantContext = v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserIdByOrgIdAndStaffIdRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetUserIdByOrgIdAndStaffIdRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOrgIdAndStaffIdRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetUserIdByOrgIdAndStaffIdRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserIdByOrgIdAndStaffIdRequestTenantContext) SetTenantId(v string) *GetUserIdByOrgIdAndStaffIdRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
