// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMdsCubeTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMdsCubeTemplatesRequest
	GetAppId() *string
	SetKeyword(v string) *ListMdsCubeTemplatesRequest
	GetKeyword() *string
	SetPageNum(v int32) *ListMdsCubeTemplatesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMdsCubeTemplatesRequest
	GetPageSize() *int32
	SetTenantId(v string) *ListMdsCubeTemplatesRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListMdsCubeTemplatesRequest
	GetWorkspaceId() *string
}

type ListMdsCubeTemplatesRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
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
	// ZXCXMAHQ-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// dev
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMdsCubeTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListMdsCubeTemplatesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMdsCubeTemplatesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMdsCubeTemplatesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMdsCubeTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMdsCubeTemplatesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMdsCubeTemplatesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMdsCubeTemplatesRequest) SetAppId(v string) *ListMdsCubeTemplatesRequest {
	s.AppId = &v
	return s
}

func (s *ListMdsCubeTemplatesRequest) SetKeyword(v string) *ListMdsCubeTemplatesRequest {
	s.Keyword = &v
	return s
}

func (s *ListMdsCubeTemplatesRequest) SetPageNum(v int32) *ListMdsCubeTemplatesRequest {
	s.PageNum = &v
	return s
}

func (s *ListMdsCubeTemplatesRequest) SetPageSize(v int32) *ListMdsCubeTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMdsCubeTemplatesRequest) SetTenantId(v string) *ListMdsCubeTemplatesRequest {
	s.TenantId = &v
	return s
}

func (s *ListMdsCubeTemplatesRequest) SetWorkspaceId(v string) *ListMdsCubeTemplatesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMdsCubeTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
