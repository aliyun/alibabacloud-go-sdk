// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatePageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListTemplatePageRequest
	GetAppId() *string
	SetCurrentPage(v int32) *ListTemplatePageRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListTemplatePageRequest
	GetPageSize() *int32
	SetTenantId(v string) *ListTemplatePageRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListTemplatePageRequest
	GetWorkspaceId() *string
}

type ListTemplatePageRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// RTHDCODI
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTemplatePageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatePageRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatePageRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListTemplatePageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTemplatePageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTemplatePageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListTemplatePageRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListTemplatePageRequest) SetAppId(v string) *ListTemplatePageRequest {
	s.AppId = &v
	return s
}

func (s *ListTemplatePageRequest) SetCurrentPage(v int32) *ListTemplatePageRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTemplatePageRequest) SetPageSize(v int32) *ListTemplatePageRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatePageRequest) SetTenantId(v string) *ListTemplatePageRequest {
	s.TenantId = &v
	return s
}

func (s *ListTemplatePageRequest) SetWorkspaceId(v string) *ListTemplatePageRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListTemplatePageRequest) Validate() error {
	return dara.Validate(s)
}
