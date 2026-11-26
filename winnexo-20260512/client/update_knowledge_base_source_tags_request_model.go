// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseSourceTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateKnowledgeBaseSourceTagsRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *UpdateKnowledgeBaseSourceTagsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateKnowledgeBaseSourceTagsRequest
	GetOpsUserId() *string
	SetSourceId(v string) *UpdateKnowledgeBaseSourceTagsRequest
	GetSourceId() *string
	SetSourceTags(v string) *UpdateKnowledgeBaseSourceTagsRequest
	GetSourceTags() *string
	SetTenantId(v string) *UpdateKnowledgeBaseSourceTagsRequest
	GetTenantId() *string
}

type UpdateKnowledgeBaseSourceTagsRequest struct {
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
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The resource tags. This is an optional parameter that accepts a JSON string list, such as ["tagA","tagB"].
	//
	// example:
	//
	// string_value
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateKnowledgeBaseSourceTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSourceTagsRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) SetAppToken(v string) *UpdateKnowledgeBaseSourceTagsRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) SetOpsUserCode(v string) *UpdateKnowledgeBaseSourceTagsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) SetOpsUserId(v string) *UpdateKnowledgeBaseSourceTagsRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) SetSourceId(v string) *UpdateKnowledgeBaseSourceTagsRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) SetSourceTags(v string) *UpdateKnowledgeBaseSourceTagsRequest {
	s.SourceTags = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) SetTenantId(v string) *UpdateKnowledgeBaseSourceTagsRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) Validate() error {
	return dara.Validate(s)
}
