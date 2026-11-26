// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReparseGroupSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ReparseGroupSourceRequest
	GetAppToken() *string
	SetForceSync(v bool) *ReparseGroupSourceRequest
	GetForceSync() *bool
	SetGroupId(v string) *ReparseGroupSourceRequest
	GetGroupId() *string
	SetOpsUserCode(v string) *ReparseGroupSourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ReparseGroupSourceRequest
	GetOpsUserId() *string
	SetSourceId(v string) *ReparseGroupSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *ReparseGroupSourceRequest
	GetTenantId() *string
}

type ReparseGroupSourceRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// Specifies whether to synchronously wait for the re-parsing to complete. Default value: false, which indicates that the request is asynchronously queued.
	//
	// example:
	//
	// false
	ForceSync *bool `json:"forceSync,omitempty" xml:"forceSync,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
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
	// The unique identifier on the business system side, which is the business ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// source_example
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this parameter explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ReparseGroupSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReparseGroupSourceRequest) GoString() string {
	return s.String()
}

func (s *ReparseGroupSourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ReparseGroupSourceRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *ReparseGroupSourceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ReparseGroupSourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ReparseGroupSourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ReparseGroupSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ReparseGroupSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReparseGroupSourceRequest) SetAppToken(v string) *ReparseGroupSourceRequest {
	s.AppToken = &v
	return s
}

func (s *ReparseGroupSourceRequest) SetForceSync(v bool) *ReparseGroupSourceRequest {
	s.ForceSync = &v
	return s
}

func (s *ReparseGroupSourceRequest) SetGroupId(v string) *ReparseGroupSourceRequest {
	s.GroupId = &v
	return s
}

func (s *ReparseGroupSourceRequest) SetOpsUserCode(v string) *ReparseGroupSourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ReparseGroupSourceRequest) SetOpsUserId(v string) *ReparseGroupSourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *ReparseGroupSourceRequest) SetSourceId(v string) *ReparseGroupSourceRequest {
	s.SourceId = &v
	return s
}

func (s *ReparseGroupSourceRequest) SetTenantId(v string) *ReparseGroupSourceRequest {
	s.TenantId = &v
	return s
}

func (s *ReparseGroupSourceRequest) Validate() error {
	return dara.Validate(s)
}
