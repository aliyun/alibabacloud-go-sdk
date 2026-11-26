// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPrimaryObjectDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *QueryPrimaryObjectDataRequest
	GetAppToken() *string
	SetKeyword(v string) *QueryPrimaryObjectDataRequest
	GetKeyword() *string
	SetOnlyFavorites(v bool) *QueryPrimaryObjectDataRequest
	GetOnlyFavorites() *bool
	SetOperatingObjectName(v string) *QueryPrimaryObjectDataRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *QueryPrimaryObjectDataRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *QueryPrimaryObjectDataRequest
	GetOpsUserId() *string
	SetPage(v int64) *QueryPrimaryObjectDataRequest
	GetPage() *int64
	SetPageSize(v int64) *QueryPrimaryObjectDataRequest
	GetPageSize() *int64
	SetTenantId(v string) *QueryPrimaryObjectDataRequest
	GetTenantId() *string
}

type QueryPrimaryObjectDataRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The keyword for search. The keyword is matched against the name field. If the schema defines a description field, the keyword is also matched against the description field.
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
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The page number. Pages start from page 1.
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
	// The tenant ID. This is a common parameter. You can explicitly pass this parameter by using --tenant-id in winnexo-cli.
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

func (s *QueryPrimaryObjectDataRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *QueryPrimaryObjectDataRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *QueryPrimaryObjectDataRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *QueryPrimaryObjectDataRequest) SetAppToken(v string) *QueryPrimaryObjectDataRequest {
	s.AppToken = &v
	return s
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

func (s *QueryPrimaryObjectDataRequest) SetOpsUserCode(v string) *QueryPrimaryObjectDataRequest {
	s.OpsUserCode = &v
	return s
}

func (s *QueryPrimaryObjectDataRequest) SetOpsUserId(v string) *QueryPrimaryObjectDataRequest {
	s.OpsUserId = &v
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
