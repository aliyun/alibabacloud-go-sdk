// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewGroupSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *PreviewGroupSourceRequest
	GetAppToken() *string
	SetGroupId(v string) *PreviewGroupSourceRequest
	GetGroupId() *string
	SetOpsUserCode(v string) *PreviewGroupSourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *PreviewGroupSourceRequest
	GetOpsUserId() *string
	SetSourceId(v string) *PreviewGroupSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *PreviewGroupSourceRequest
	GetTenantId() *string
}

type PreviewGroupSourceRequest struct {
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
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// src_feishu_doc_1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. You can pass this parameter explicitly by using the --tenant-id option in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s PreviewGroupSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewGroupSourceRequest) GoString() string {
	return s.String()
}

func (s *PreviewGroupSourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *PreviewGroupSourceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *PreviewGroupSourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *PreviewGroupSourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *PreviewGroupSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *PreviewGroupSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PreviewGroupSourceRequest) SetAppToken(v string) *PreviewGroupSourceRequest {
	s.AppToken = &v
	return s
}

func (s *PreviewGroupSourceRequest) SetGroupId(v string) *PreviewGroupSourceRequest {
	s.GroupId = &v
	return s
}

func (s *PreviewGroupSourceRequest) SetOpsUserCode(v string) *PreviewGroupSourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *PreviewGroupSourceRequest) SetOpsUserId(v string) *PreviewGroupSourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *PreviewGroupSourceRequest) SetSourceId(v string) *PreviewGroupSourceRequest {
	s.SourceId = &v
	return s
}

func (s *PreviewGroupSourceRequest) SetTenantId(v string) *PreviewGroupSourceRequest {
	s.TenantId = &v
	return s
}

func (s *PreviewGroupSourceRequest) Validate() error {
	return dara.Validate(s)
}
