// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewPersonalSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *PreviewPersonalSourceRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *PreviewPersonalSourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *PreviewPersonalSourceRequest
	GetOpsUserId() *string
	SetSourceId(v string) *PreviewPersonalSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *PreviewPersonalSourceRequest
	GetTenantId() *string
}

type PreviewPersonalSourceRequest struct {
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
	// The data source ID, which is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. The winnexo-cli passes this parameter explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s PreviewPersonalSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewPersonalSourceRequest) GoString() string {
	return s.String()
}

func (s *PreviewPersonalSourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *PreviewPersonalSourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *PreviewPersonalSourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *PreviewPersonalSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *PreviewPersonalSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PreviewPersonalSourceRequest) SetAppToken(v string) *PreviewPersonalSourceRequest {
	s.AppToken = &v
	return s
}

func (s *PreviewPersonalSourceRequest) SetOpsUserCode(v string) *PreviewPersonalSourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *PreviewPersonalSourceRequest) SetOpsUserId(v string) *PreviewPersonalSourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *PreviewPersonalSourceRequest) SetSourceId(v string) *PreviewPersonalSourceRequest {
	s.SourceId = &v
	return s
}

func (s *PreviewPersonalSourceRequest) SetTenantId(v string) *PreviewPersonalSourceRequest {
	s.TenantId = &v
	return s
}

func (s *PreviewPersonalSourceRequest) Validate() error {
	return dara.Validate(s)
}
