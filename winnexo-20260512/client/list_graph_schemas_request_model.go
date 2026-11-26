// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGraphSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListGraphSchemasRequest
	GetAppToken() *string
	SetKeyword(v string) *ListGraphSchemasRequest
	GetKeyword() *string
	SetOpsUserCode(v string) *ListGraphSchemasRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListGraphSchemasRequest
	GetOpsUserId() *string
	SetSemanticTags(v []*string) *ListGraphSchemasRequest
	GetSemanticTags() []*string
	SetTenantId(v string) *ListGraphSchemasRequest
	GetTenantId() *string
}

type ListGraphSchemasRequest struct {
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
	SemanticTags []*string `json:"semanticTags,omitempty" xml:"semanticTags,omitempty" type:"Repeated"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListGraphSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGraphSchemasRequest) GoString() string {
	return s.String()
}

func (s *ListGraphSchemasRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListGraphSchemasRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListGraphSchemasRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListGraphSchemasRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListGraphSchemasRequest) GetSemanticTags() []*string {
	return s.SemanticTags
}

func (s *ListGraphSchemasRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListGraphSchemasRequest) SetAppToken(v string) *ListGraphSchemasRequest {
	s.AppToken = &v
	return s
}

func (s *ListGraphSchemasRequest) SetKeyword(v string) *ListGraphSchemasRequest {
	s.Keyword = &v
	return s
}

func (s *ListGraphSchemasRequest) SetOpsUserCode(v string) *ListGraphSchemasRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListGraphSchemasRequest) SetOpsUserId(v string) *ListGraphSchemasRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListGraphSchemasRequest) SetSemanticTags(v []*string) *ListGraphSchemasRequest {
	s.SemanticTags = v
	return s
}

func (s *ListGraphSchemasRequest) SetTenantId(v string) *ListGraphSchemasRequest {
	s.TenantId = &v
	return s
}

func (s *ListGraphSchemasRequest) Validate() error {
	return dara.Validate(s)
}
