// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryKnowledgeBaseFailedSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *RetryKnowledgeBaseFailedSourcesRequest
	GetAppToken() *string
	SetDirectoryId(v string) *RetryKnowledgeBaseFailedSourcesRequest
	GetDirectoryId() *string
	SetOpsUserCode(v string) *RetryKnowledgeBaseFailedSourcesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *RetryKnowledgeBaseFailedSourcesRequest
	GetOpsUserId() *string
	SetTenantId(v string) *RetryKnowledgeBaseFailedSourcesRequest
	GetTenantId() *string
}

type RetryKnowledgeBaseFailedSourcesRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the enterprise knowledge base directory. Failed resources in subdirectories are recursively included.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
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
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this value explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RetryKnowledgeBaseFailedSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseFailedSourcesRequest) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) SetAppToken(v string) *RetryKnowledgeBaseFailedSourcesRequest {
	s.AppToken = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) SetDirectoryId(v string) *RetryKnowledgeBaseFailedSourcesRequest {
	s.DirectoryId = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) SetOpsUserCode(v string) *RetryKnowledgeBaseFailedSourcesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) SetOpsUserId(v string) *RetryKnowledgeBaseFailedSourcesRequest {
	s.OpsUserId = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) SetTenantId(v string) *RetryKnowledgeBaseFailedSourcesRequest {
	s.TenantId = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesRequest) Validate() error {
	return dara.Validate(s)
}
