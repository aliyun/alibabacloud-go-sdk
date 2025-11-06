// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMdsCubeTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMdsCubeTasksRequest
	GetAppId() *string
	SetPageNum(v int32) *ListMdsCubeTasksRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMdsCubeTasksRequest
	GetPageSize() *int32
	SetTemplateResourceId(v int64) *ListMdsCubeTasksRequest
	GetTemplateResourceId() *int64
	SetTenantId(v string) *ListMdsCubeTasksRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListMdsCubeTasksRequest
	GetWorkspaceId() *string
}

type ListMdsCubeTasksRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TemplateResourceId *int64 `json:"TemplateResourceId,omitempty" xml:"TemplateResourceId,omitempty"`
	// example:
	//
	// ZXCXMAHQ-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// dev
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMdsCubeTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeTasksRequest) GoString() string {
	return s.String()
}

func (s *ListMdsCubeTasksRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMdsCubeTasksRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMdsCubeTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMdsCubeTasksRequest) GetTemplateResourceId() *int64 {
	return s.TemplateResourceId
}

func (s *ListMdsCubeTasksRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMdsCubeTasksRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMdsCubeTasksRequest) SetAppId(v string) *ListMdsCubeTasksRequest {
	s.AppId = &v
	return s
}

func (s *ListMdsCubeTasksRequest) SetPageNum(v int32) *ListMdsCubeTasksRequest {
	s.PageNum = &v
	return s
}

func (s *ListMdsCubeTasksRequest) SetPageSize(v int32) *ListMdsCubeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListMdsCubeTasksRequest) SetTemplateResourceId(v int64) *ListMdsCubeTasksRequest {
	s.TemplateResourceId = &v
	return s
}

func (s *ListMdsCubeTasksRequest) SetTenantId(v string) *ListMdsCubeTasksRequest {
	s.TenantId = &v
	return s
}

func (s *ListMdsCubeTasksRequest) SetWorkspaceId(v string) *ListMdsCubeTasksRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMdsCubeTasksRequest) Validate() error {
	return dara.Validate(s)
}
