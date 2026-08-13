// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutputFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItemType(v string) *ListOutputFilesRequest
	GetItemType() *string
	SetKeyword(v string) *ListOutputFilesRequest
	GetKeyword() *string
	SetOperatingObjectName(v string) *ListOutputFilesRequest
	GetOperatingObjectName() *string
	SetPage(v int64) *ListOutputFilesRequest
	GetPage() *int64
	SetPageSize(v int64) *ListOutputFilesRequest
	GetPageSize() *int64
	SetSharedOnly(v bool) *ListOutputFilesRequest
	GetSharedOnly() *bool
	SetTenantId(v string) *ListOutputFilesRequest
	GetTenantId() *string
}

type ListOutputFilesRequest struct {
	// 产出明细类型: ppt/html/document/picture/slides/video/audio/email/others
	//
	// example:
	//
	// ppt
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// 关键词搜索，匹配产出标题或明细名称
	//
	// example:
	//
	// string_value
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 数字员工（运营对象）名称，按名称过滤
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 页码，从 1 开始
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页数量，范围 1-100
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 是否仅展示开启分享的产出和产出明细
	//
	// example:
	//
	// False
	SharedOnly *bool `json:"sharedOnly,omitempty" xml:"sharedOnly,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListOutputFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOutputFilesRequest) GoString() string {
	return s.String()
}

func (s *ListOutputFilesRequest) GetItemType() *string {
	return s.ItemType
}

func (s *ListOutputFilesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListOutputFilesRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListOutputFilesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListOutputFilesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOutputFilesRequest) GetSharedOnly() *bool {
	return s.SharedOnly
}

func (s *ListOutputFilesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListOutputFilesRequest) SetItemType(v string) *ListOutputFilesRequest {
	s.ItemType = &v
	return s
}

func (s *ListOutputFilesRequest) SetKeyword(v string) *ListOutputFilesRequest {
	s.Keyword = &v
	return s
}

func (s *ListOutputFilesRequest) SetOperatingObjectName(v string) *ListOutputFilesRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListOutputFilesRequest) SetPage(v int64) *ListOutputFilesRequest {
	s.Page = &v
	return s
}

func (s *ListOutputFilesRequest) SetPageSize(v int64) *ListOutputFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOutputFilesRequest) SetSharedOnly(v bool) *ListOutputFilesRequest {
	s.SharedOnly = &v
	return s
}

func (s *ListOutputFilesRequest) SetTenantId(v string) *ListOutputFilesRequest {
	s.TenantId = &v
	return s
}

func (s *ListOutputFilesRequest) Validate() error {
	return dara.Validate(s)
}
