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
	// The keyword for searching. The keyword is matched against the name field by default. If the schema defines a description field, the keyword is also matched against the description field.
	//
	// example:
	//
	// SampleKeyword
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// Specifies whether to return only favorited primary objects. If this parameter is set to false or not specified, all objects are returned, including the isFavorited flag.
	//
	// example:
	//
	// false
	OnlyFavorites *bool `json:"onlyFavorites,omitempty" xml:"onlyFavorites,omitempty"`
	// The operating object name, such as customer_1.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tenant ID. This is a common parameter. You can explicitly pass it in winnexo-cli by using --tenant-id.
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
