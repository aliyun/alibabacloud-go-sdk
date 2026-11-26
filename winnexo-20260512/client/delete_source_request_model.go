// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *DeleteSourceRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *DeleteSourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *DeleteSourceRequest
	GetOpsUserId() *string
	SetSourceId(v string) *DeleteSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *DeleteSourceRequest
	GetTenantId() *string
}

type DeleteSourceRequest struct {
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
	// The unique identifier on the business system side, that is, the business ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 781
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 21577
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *DeleteSourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *DeleteSourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *DeleteSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *DeleteSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteSourceRequest) SetAppToken(v string) *DeleteSourceRequest {
	s.AppToken = &v
	return s
}

func (s *DeleteSourceRequest) SetOpsUserCode(v string) *DeleteSourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *DeleteSourceRequest) SetOpsUserId(v string) *DeleteSourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *DeleteSourceRequest) SetSourceId(v string) *DeleteSourceRequest {
	s.SourceId = &v
	return s
}

func (s *DeleteSourceRequest) SetTenantId(v string) *DeleteSourceRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteSourceRequest) Validate() error {
	return dara.Validate(s)
}
