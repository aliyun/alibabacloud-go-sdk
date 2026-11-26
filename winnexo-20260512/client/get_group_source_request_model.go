// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetGroupSourceRequest
	GetAppToken() *string
	SetGroupId(v string) *GetGroupSourceRequest
	GetGroupId() *string
	SetOpsUserCode(v string) *GetGroupSourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetGroupSourceRequest
	GetOpsUserId() *string
	SetSourceId(v string) *GetGroupSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *GetGroupSourceRequest
	GetTenantId() *string
}

type GetGroupSourceRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleGroupId
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
	// The ID of the personal FILE data source to be replaced. The ID is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetGroupSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupSourceRequest) GoString() string {
	return s.String()
}

func (s *GetGroupSourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetGroupSourceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetGroupSourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetGroupSourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetGroupSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *GetGroupSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetGroupSourceRequest) SetAppToken(v string) *GetGroupSourceRequest {
	s.AppToken = &v
	return s
}

func (s *GetGroupSourceRequest) SetGroupId(v string) *GetGroupSourceRequest {
	s.GroupId = &v
	return s
}

func (s *GetGroupSourceRequest) SetOpsUserCode(v string) *GetGroupSourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetGroupSourceRequest) SetOpsUserId(v string) *GetGroupSourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetGroupSourceRequest) SetSourceId(v string) *GetGroupSourceRequest {
	s.SourceId = &v
	return s
}

func (s *GetGroupSourceRequest) SetTenantId(v string) *GetGroupSourceRequest {
	s.TenantId = &v
	return s
}

func (s *GetGroupSourceRequest) Validate() error {
	return dara.Validate(s)
}
