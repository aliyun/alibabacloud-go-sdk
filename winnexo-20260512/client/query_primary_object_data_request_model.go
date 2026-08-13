// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPrimaryObjectDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *QueryPrimaryObjectDataRequest
	GetKeyword() *string
	SetOnlyFavorites(v bool) *QueryPrimaryObjectDataRequest
	GetOnlyFavorites() *bool
	SetOperatingObjectName(v string) *QueryPrimaryObjectDataRequest
	GetOperatingObjectName() *string
	SetPage(v int64) *QueryPrimaryObjectDataRequest
	GetPage() *int64
	SetPageSize(v int64) *QueryPrimaryObjectDataRequest
	GetPageSize() *int64
	SetTenantId(v string) *QueryPrimaryObjectDataRequest
	GetTenantId() *string
}

type QueryPrimaryObjectDataRequest struct {
	// 关键字搜索（固定匹配 name；若 schema 定义 description，则同时匹配 description）
	//
	// example:
	//
	// 示例关键词
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 仅返回关注的主对象；false 或不传则返回全部对象（包含 isFavorited 标识）
	//
	// example:
	//
	// false
	OnlyFavorites *bool `json:"onlyFavorites,omitempty" xml:"onlyFavorites,omitempty"`
	// 运营对象名称（如 customer_1）
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 页码（从 1 开始）
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
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryPrimaryObjectDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPrimaryObjectDataRequest) GoString() string {
	return s.String()
}

func (s *QueryPrimaryObjectDataRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryPrimaryObjectDataRequest) GetOnlyFavorites() *bool {
	return s.OnlyFavorites
}

func (s *QueryPrimaryObjectDataRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *QueryPrimaryObjectDataRequest) GetPage() *int64 {
	return s.Page
}

func (s *QueryPrimaryObjectDataRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryPrimaryObjectDataRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryPrimaryObjectDataRequest) SetKeyword(v string) *QueryPrimaryObjectDataRequest {
	s.Keyword = &v
	return s
}

func (s *QueryPrimaryObjectDataRequest) SetOnlyFavorites(v bool) *QueryPrimaryObjectDataRequest {
	s.OnlyFavorites = &v
	return s
}

func (s *QueryPrimaryObjectDataRequest) SetOperatingObjectName(v string) *QueryPrimaryObjectDataRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *QueryPrimaryObjectDataRequest) SetPage(v int64) *QueryPrimaryObjectDataRequest {
	s.Page = &v
	return s
}

func (s *QueryPrimaryObjectDataRequest) SetPageSize(v int64) *QueryPrimaryObjectDataRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPrimaryObjectDataRequest) SetTenantId(v string) *QueryPrimaryObjectDataRequest {
	s.TenantId = &v
	return s
}

func (s *QueryPrimaryObjectDataRequest) Validate() error {
	return dara.Validate(s)
}
