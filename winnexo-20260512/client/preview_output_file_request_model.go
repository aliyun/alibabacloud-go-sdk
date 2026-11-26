// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewOutputFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *PreviewOutputFileRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *PreviewOutputFileRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *PreviewOutputFileRequest
	GetOpsUserId() *string
	SetOutputItemId(v string) *PreviewOutputFileRequest
	GetOutputItemId() *string
	SetTenantId(v string) *PreviewOutputFileRequest
	GetTenantId() *string
	SetVersionId(v string) *PreviewOutputFileRequest
	GetVersionId() *string
}

type PreviewOutputFileRequest struct {
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
	// The output item ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleOutputItemId
	OutputItemId *string `json:"outputItemId,omitempty" xml:"outputItemId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The version referenced by the alias.
	//
	// example:
	//
	// exampleVersionId
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s PreviewOutputFileRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewOutputFileRequest) GoString() string {
	return s.String()
}

func (s *PreviewOutputFileRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *PreviewOutputFileRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *PreviewOutputFileRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *PreviewOutputFileRequest) GetOutputItemId() *string {
	return s.OutputItemId
}

func (s *PreviewOutputFileRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PreviewOutputFileRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *PreviewOutputFileRequest) SetAppToken(v string) *PreviewOutputFileRequest {
	s.AppToken = &v
	return s
}

func (s *PreviewOutputFileRequest) SetOpsUserCode(v string) *PreviewOutputFileRequest {
	s.OpsUserCode = &v
	return s
}

func (s *PreviewOutputFileRequest) SetOpsUserId(v string) *PreviewOutputFileRequest {
	s.OpsUserId = &v
	return s
}

func (s *PreviewOutputFileRequest) SetOutputItemId(v string) *PreviewOutputFileRequest {
	s.OutputItemId = &v
	return s
}

func (s *PreviewOutputFileRequest) SetTenantId(v string) *PreviewOutputFileRequest {
	s.TenantId = &v
	return s
}

func (s *PreviewOutputFileRequest) SetVersionId(v string) *PreviewOutputFileRequest {
	s.VersionId = &v
	return s
}

func (s *PreviewOutputFileRequest) Validate() error {
	return dara.Validate(s)
}
