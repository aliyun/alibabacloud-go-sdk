// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSourceContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateSourceContentRequest
	GetAppToken() *string
	SetContent(v string) *UpdateSourceContentRequest
	GetContent() *string
	SetForceSync(v bool) *UpdateSourceContentRequest
	GetForceSync() *bool
	SetOpsUserCode(v string) *UpdateSourceContentRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateSourceContentRequest
	GetOpsUserId() *string
	SetSourceId(v string) *UpdateSourceContentRequest
	GetSourceId() *string
	SetTenantId(v string) *UpdateSourceContentRequest
	GetTenantId() *string
}

type UpdateSourceContentRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The returned content.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Specifies whether to force synchronous processing.
	//
	// example:
	//
	// false
	ForceSync *bool `json:"forceSync,omitempty" xml:"forceSync,omitempty"`
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
	// The ID of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateSourceContentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSourceContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateSourceContentRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateSourceContentRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateSourceContentRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *UpdateSourceContentRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateSourceContentRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateSourceContentRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateSourceContentRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateSourceContentRequest) SetAppToken(v string) *UpdateSourceContentRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateSourceContentRequest) SetContent(v string) *UpdateSourceContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateSourceContentRequest) SetForceSync(v bool) *UpdateSourceContentRequest {
	s.ForceSync = &v
	return s
}

func (s *UpdateSourceContentRequest) SetOpsUserCode(v string) *UpdateSourceContentRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateSourceContentRequest) SetOpsUserId(v string) *UpdateSourceContentRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateSourceContentRequest) SetSourceId(v string) *UpdateSourceContentRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateSourceContentRequest) SetTenantId(v string) *UpdateSourceContentRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateSourceContentRequest) Validate() error {
	return dara.Validate(s)
}
