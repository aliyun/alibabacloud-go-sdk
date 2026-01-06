// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkProjectionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GetDingtalkProjectionListRequestTenantContext) *GetDingtalkProjectionListRequest
	GetTenantContext() *GetDingtalkProjectionListRequestTenantContext
	SetCode(v string) *GetDingtalkProjectionListRequest
	GetCode() *string
	SetCurrentPage(v int32) *GetDingtalkProjectionListRequest
	GetCurrentPage() *int32
	SetOrgId(v int64) *GetDingtalkProjectionListRequest
	GetOrgId() *int64
	SetPageSize(v int32) *GetDingtalkProjectionListRequest
	GetPageSize() *int32
	SetProjectorWorkNo(v string) *GetDingtalkProjectionListRequest
	GetProjectorWorkNo() *string
}

type GetDingtalkProjectionListRequest struct {
	TenantContext *GetDingtalkProjectionListRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// PROJ001
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 123456789
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 343242
	ProjectorWorkNo *string `json:"projectorWorkNo,omitempty" xml:"projectorWorkNo,omitempty"`
}

func (s GetDingtalkProjectionListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionListRequest) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionListRequest) GetTenantContext() *GetDingtalkProjectionListRequestTenantContext {
	return s.TenantContext
}

func (s *GetDingtalkProjectionListRequest) GetCode() *string {
	return s.Code
}

func (s *GetDingtalkProjectionListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetDingtalkProjectionListRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *GetDingtalkProjectionListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDingtalkProjectionListRequest) GetProjectorWorkNo() *string {
	return s.ProjectorWorkNo
}

func (s *GetDingtalkProjectionListRequest) SetTenantContext(v *GetDingtalkProjectionListRequestTenantContext) *GetDingtalkProjectionListRequest {
	s.TenantContext = v
	return s
}

func (s *GetDingtalkProjectionListRequest) SetCode(v string) *GetDingtalkProjectionListRequest {
	s.Code = &v
	return s
}

func (s *GetDingtalkProjectionListRequest) SetCurrentPage(v int32) *GetDingtalkProjectionListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetDingtalkProjectionListRequest) SetOrgId(v int64) *GetDingtalkProjectionListRequest {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkProjectionListRequest) SetPageSize(v int32) *GetDingtalkProjectionListRequest {
	s.PageSize = &v
	return s
}

func (s *GetDingtalkProjectionListRequest) SetProjectorWorkNo(v string) *GetDingtalkProjectionListRequest {
	s.ProjectorWorkNo = &v
	return s
}

func (s *GetDingtalkProjectionListRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDingtalkProjectionListRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetDingtalkProjectionListRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionListRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionListRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDingtalkProjectionListRequestTenantContext) SetTenantId(v string) *GetDingtalkProjectionListRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetDingtalkProjectionListRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
