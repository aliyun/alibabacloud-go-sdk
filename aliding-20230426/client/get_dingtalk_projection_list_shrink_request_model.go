// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkProjectionListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GetDingtalkProjectionListShrinkRequest
	GetTenantContextShrink() *string
	SetCode(v string) *GetDingtalkProjectionListShrinkRequest
	GetCode() *string
	SetCurrentPage(v int32) *GetDingtalkProjectionListShrinkRequest
	GetCurrentPage() *int32
	SetOrgId(v int64) *GetDingtalkProjectionListShrinkRequest
	GetOrgId() *int64
	SetPageSize(v int32) *GetDingtalkProjectionListShrinkRequest
	GetPageSize() *int32
	SetProjectorWorkNo(v string) *GetDingtalkProjectionListShrinkRequest
	GetProjectorWorkNo() *string
}

type GetDingtalkProjectionListShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
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

func (s GetDingtalkProjectionListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionListShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetDingtalkProjectionListShrinkRequest) GetCode() *string {
	return s.Code
}

func (s *GetDingtalkProjectionListShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetDingtalkProjectionListShrinkRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *GetDingtalkProjectionListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDingtalkProjectionListShrinkRequest) GetProjectorWorkNo() *string {
	return s.ProjectorWorkNo
}

func (s *GetDingtalkProjectionListShrinkRequest) SetTenantContextShrink(v string) *GetDingtalkProjectionListShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetDingtalkProjectionListShrinkRequest) SetCode(v string) *GetDingtalkProjectionListShrinkRequest {
	s.Code = &v
	return s
}

func (s *GetDingtalkProjectionListShrinkRequest) SetCurrentPage(v int32) *GetDingtalkProjectionListShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetDingtalkProjectionListShrinkRequest) SetOrgId(v int64) *GetDingtalkProjectionListShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkProjectionListShrinkRequest) SetPageSize(v int32) *GetDingtalkProjectionListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetDingtalkProjectionListShrinkRequest) SetProjectorWorkNo(v string) *GetDingtalkProjectionListShrinkRequest {
	s.ProjectorWorkNo = &v
	return s
}

func (s *GetDingtalkProjectionListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
