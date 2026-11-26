// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReparseSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ReparseSourceRequest
	GetAppToken() *string
	SetForceSync(v bool) *ReparseSourceRequest
	GetForceSync() *bool
	SetOpsUserCode(v string) *ReparseSourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ReparseSourceRequest
	GetOpsUserId() *string
	SetSourceId(v string) *ReparseSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *ReparseSourceRequest
	GetTenantId() *string
}

type ReparseSourceRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// Specifies whether to synchronously wait for the re-parsing to complete. Default value: false, which indicates asynchronous queuing.
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
	// The ID of the data source to re-parse. This ID is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this value explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReparseSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReparseSourceRequest) GoString() string {
	return s.String()
}

func (s *ReparseSourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ReparseSourceRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *ReparseSourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ReparseSourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ReparseSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReparseSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReparseSourceRequest) SetAppToken(v string) *ReparseSourceRequest {
	s.AppToken = &v
	return s
}

func (s *ReparseSourceRequest) SetForceSync(v bool) *ReparseSourceRequest {
	s.ForceSync = &v
	return s
}

func (s *ReparseSourceRequest) SetOpsUserCode(v string) *ReparseSourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ReparseSourceRequest) SetOpsUserId(v string) *ReparseSourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *ReparseSourceRequest) SetSourceId(v string) *ReparseSourceRequest {
	s.SourceId = &v
	return s
}

func (s *ReparseSourceRequest) SetTenantId(v string) *ReparseSourceRequest {
	s.TenantId = &v
	return s
}

func (s *ReparseSourceRequest) Validate() error {
	return dara.Validate(s)
}
