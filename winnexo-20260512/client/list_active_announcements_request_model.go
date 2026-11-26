// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActiveAnnouncementsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListActiveAnnouncementsRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *ListActiveAnnouncementsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListActiveAnnouncementsRequest
	GetOpsUserId() *string
	SetPageNumber(v int64) *ListActiveAnnouncementsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListActiveAnnouncementsRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListActiveAnnouncementsRequest
	GetTenantId() *string
}

type ListActiveAnnouncementsRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
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
	// The page number for paginated queries.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 100. Maximum value: 500.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListActiveAnnouncementsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListActiveAnnouncementsRequest) GoString() string {
	return s.String()
}

func (s *ListActiveAnnouncementsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListActiveAnnouncementsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListActiveAnnouncementsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListActiveAnnouncementsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListActiveAnnouncementsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListActiveAnnouncementsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListActiveAnnouncementsRequest) SetAppToken(v string) *ListActiveAnnouncementsRequest {
	s.AppToken = &v
	return s
}

func (s *ListActiveAnnouncementsRequest) SetOpsUserCode(v string) *ListActiveAnnouncementsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListActiveAnnouncementsRequest) SetOpsUserId(v string) *ListActiveAnnouncementsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListActiveAnnouncementsRequest) SetPageNumber(v int64) *ListActiveAnnouncementsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListActiveAnnouncementsRequest) SetPageSize(v int64) *ListActiveAnnouncementsRequest {
	s.PageSize = &v
	return s
}

func (s *ListActiveAnnouncementsRequest) SetTenantId(v string) *ListActiveAnnouncementsRequest {
	s.TenantId = &v
	return s
}

func (s *ListActiveAnnouncementsRequest) Validate() error {
	return dara.Validate(s)
}
