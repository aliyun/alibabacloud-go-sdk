// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryDirectoryFailedSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *RetryDirectoryFailedSourcesRequest
	GetAppToken() *string
	SetDirectoryId(v string) *RetryDirectoryFailedSourcesRequest
	GetDirectoryId() *string
	SetOpsUserCode(v string) *RetryDirectoryFailedSourcesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *RetryDirectoryFailedSourcesRequest
	GetOpsUserId() *string
	SetTenantId(v string) *RetryDirectoryFailedSourcesRequest
	GetTenantId() *string
}

type RetryDirectoryFailedSourcesRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the enterprise knowledge base folder. Failed resources in subfolders are recursively included.
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

func (s RetryDirectoryFailedSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryDirectoryFailedSourcesRequest) GoString() string {
	return s.String()
}

func (s *RetryDirectoryFailedSourcesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RetryDirectoryFailedSourcesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *RetryDirectoryFailedSourcesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *RetryDirectoryFailedSourcesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *RetryDirectoryFailedSourcesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RetryDirectoryFailedSourcesRequest) SetAppToken(v string) *RetryDirectoryFailedSourcesRequest {
	s.AppToken = &v
	return s
}

func (s *RetryDirectoryFailedSourcesRequest) SetDirectoryId(v string) *RetryDirectoryFailedSourcesRequest {
	s.DirectoryId = &v
	return s
}

func (s *RetryDirectoryFailedSourcesRequest) SetOpsUserCode(v string) *RetryDirectoryFailedSourcesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *RetryDirectoryFailedSourcesRequest) SetOpsUserId(v string) *RetryDirectoryFailedSourcesRequest {
	s.OpsUserId = &v
	return s
}

func (s *RetryDirectoryFailedSourcesRequest) SetTenantId(v string) *RetryDirectoryFailedSourcesRequest {
	s.TenantId = &v
	return s
}

func (s *RetryDirectoryFailedSourcesRequest) Validate() error {
	return dara.Validate(s)
}
