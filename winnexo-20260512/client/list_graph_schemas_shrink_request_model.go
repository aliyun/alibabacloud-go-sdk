// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGraphSchemasShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListGraphSchemasShrinkRequest
	GetAppToken() *string
	SetKeyword(v string) *ListGraphSchemasShrinkRequest
	GetKeyword() *string
	SetOpsUserCode(v string) *ListGraphSchemasShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListGraphSchemasShrinkRequest
	GetOpsUserId() *string
	SetSemanticTagsShrink(v string) *ListGraphSchemasShrinkRequest
	GetSemanticTagsShrink() *string
	SetTenantId(v string) *ListGraphSchemasShrinkRequest
	GetTenantId() *string
}

type ListGraphSchemasShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The keyword for fuzzy match of component data in the form.
	//
	// example:
	//
	// crm
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
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
	// The semantic tags used for filtering. A graph is retained if any tag matches.
	//
	// example:
	//
	// ["Sales"]
	SemanticTagsShrink *string `json:"semanticTags,omitempty" xml:"semanticTags,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListGraphSchemasShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGraphSchemasShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGraphSchemasShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListGraphSchemasShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListGraphSchemasShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListGraphSchemasShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListGraphSchemasShrinkRequest) GetSemanticTagsShrink() *string {
	return s.SemanticTagsShrink
}

func (s *ListGraphSchemasShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListGraphSchemasShrinkRequest) SetAppToken(v string) *ListGraphSchemasShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *ListGraphSchemasShrinkRequest) SetKeyword(v string) *ListGraphSchemasShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListGraphSchemasShrinkRequest) SetOpsUserCode(v string) *ListGraphSchemasShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListGraphSchemasShrinkRequest) SetOpsUserId(v string) *ListGraphSchemasShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListGraphSchemasShrinkRequest) SetSemanticTagsShrink(v string) *ListGraphSchemasShrinkRequest {
	s.SemanticTagsShrink = &v
	return s
}

func (s *ListGraphSchemasShrinkRequest) SetTenantId(v string) *ListGraphSchemasShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListGraphSchemasShrinkRequest) Validate() error {
	return dara.Validate(s)
}
