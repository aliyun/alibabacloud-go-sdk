// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMdsCubeResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMdsCubeResourcesRequest
	GetAppId() *string
	SetPageNum(v int32) *ListMdsCubeResourcesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMdsCubeResourcesRequest
	GetPageSize() *int32
	SetTemplateId(v string) *ListMdsCubeResourcesRequest
	GetTemplateId() *string
	SetTenantId(v string) *ListMdsCubeResourcesRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListMdsCubeResourcesRequest
	GetWorkspaceId() *string
	SetTest(v string) *ListMdsCubeResourcesRequest
	GetTest() *string
}

type ListMdsCubeResourcesRequest struct {
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
	// test12435
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// ZXCXMAHQ-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// dev
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// test
	Test *string `json:"test,omitempty" xml:"test,omitempty"`
}

func (s ListMdsCubeResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListMdsCubeResourcesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMdsCubeResourcesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMdsCubeResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMdsCubeResourcesRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListMdsCubeResourcesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMdsCubeResourcesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMdsCubeResourcesRequest) GetTest() *string {
	return s.Test
}

func (s *ListMdsCubeResourcesRequest) SetAppId(v string) *ListMdsCubeResourcesRequest {
	s.AppId = &v
	return s
}

func (s *ListMdsCubeResourcesRequest) SetPageNum(v int32) *ListMdsCubeResourcesRequest {
	s.PageNum = &v
	return s
}

func (s *ListMdsCubeResourcesRequest) SetPageSize(v int32) *ListMdsCubeResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMdsCubeResourcesRequest) SetTemplateId(v string) *ListMdsCubeResourcesRequest {
	s.TemplateId = &v
	return s
}

func (s *ListMdsCubeResourcesRequest) SetTenantId(v string) *ListMdsCubeResourcesRequest {
	s.TenantId = &v
	return s
}

func (s *ListMdsCubeResourcesRequest) SetWorkspaceId(v string) *ListMdsCubeResourcesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMdsCubeResourcesRequest) SetTest(v string) *ListMdsCubeResourcesRequest {
	s.Test = &v
	return s
}

func (s *ListMdsCubeResourcesRequest) Validate() error {
	return dara.Validate(s)
}
