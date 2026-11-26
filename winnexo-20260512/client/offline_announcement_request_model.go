// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineAnnouncementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnouncementId(v int64) *OfflineAnnouncementRequest
	GetAnnouncementId() *int64
	SetAppToken(v string) *OfflineAnnouncementRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *OfflineAnnouncementRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *OfflineAnnouncementRequest
	GetOpsUserId() *string
	SetTenantId(v string) *OfflineAnnouncementRequest
	GetTenantId() *string
}

type OfflineAnnouncementRequest struct {
	// The business ID of the announcement.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001
	AnnouncementId *int64 `json:"announcementId,omitempty" xml:"announcementId,omitempty"`
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
	// The tenant ID. This is a common parameter. Pass it explicitly in winnexo-cli by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s OfflineAnnouncementRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineAnnouncementRequest) GoString() string {
	return s.String()
}

func (s *OfflineAnnouncementRequest) GetAnnouncementId() *int64 {
	return s.AnnouncementId
}

func (s *OfflineAnnouncementRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *OfflineAnnouncementRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *OfflineAnnouncementRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *OfflineAnnouncementRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *OfflineAnnouncementRequest) SetAnnouncementId(v int64) *OfflineAnnouncementRequest {
	s.AnnouncementId = &v
	return s
}

func (s *OfflineAnnouncementRequest) SetAppToken(v string) *OfflineAnnouncementRequest {
	s.AppToken = &v
	return s
}

func (s *OfflineAnnouncementRequest) SetOpsUserCode(v string) *OfflineAnnouncementRequest {
	s.OpsUserCode = &v
	return s
}

func (s *OfflineAnnouncementRequest) SetOpsUserId(v string) *OfflineAnnouncementRequest {
	s.OpsUserId = &v
	return s
}

func (s *OfflineAnnouncementRequest) SetTenantId(v string) *OfflineAnnouncementRequest {
	s.TenantId = &v
	return s
}

func (s *OfflineAnnouncementRequest) Validate() error {
	return dara.Validate(s)
}
