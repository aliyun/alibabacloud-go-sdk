// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisibleKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListVisibleKnowledgeBasesRequest
	GetAppToken() *string
	SetOperatingObjectName(v string) *ListVisibleKnowledgeBasesRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *ListVisibleKnowledgeBasesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListVisibleKnowledgeBasesRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ListVisibleKnowledgeBasesRequest
	GetTenantId() *string
}

type ListVisibleKnowledgeBasesRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The name of the digital employee (operating object name).
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
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
	// The tenant ID.
	//
	// example:
	//
	// 1729094555111072
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListVisibleKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBasesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListVisibleKnowledgeBasesRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListVisibleKnowledgeBasesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListVisibleKnowledgeBasesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListVisibleKnowledgeBasesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVisibleKnowledgeBasesRequest) SetAppToken(v string) *ListVisibleKnowledgeBasesRequest {
	s.AppToken = &v
	return s
}

func (s *ListVisibleKnowledgeBasesRequest) SetOperatingObjectName(v string) *ListVisibleKnowledgeBasesRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListVisibleKnowledgeBasesRequest) SetOpsUserCode(v string) *ListVisibleKnowledgeBasesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListVisibleKnowledgeBasesRequest) SetOpsUserId(v string) *ListVisibleKnowledgeBasesRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListVisibleKnowledgeBasesRequest) SetTenantId(v string) *ListVisibleKnowledgeBasesRequest {
	s.TenantId = &v
	return s
}

func (s *ListVisibleKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
