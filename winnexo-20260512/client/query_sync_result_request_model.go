// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySyncResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *QuerySyncResultRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *QuerySyncResultRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *QuerySyncResultRequest
	GetOpsUserId() *string
	SetTaskId(v int64) *QuerySyncResultRequest
	GetTaskId() *int64
	SetTenantId(v string) *QuerySyncResultRequest
	GetTenantId() *string
}

type QuerySyncResultRequest struct {
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
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QuerySyncResultRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySyncResultRequest) GoString() string {
	return s.String()
}

func (s *QuerySyncResultRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *QuerySyncResultRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *QuerySyncResultRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *QuerySyncResultRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QuerySyncResultRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QuerySyncResultRequest) SetAppToken(v string) *QuerySyncResultRequest {
	s.AppToken = &v
	return s
}

func (s *QuerySyncResultRequest) SetOpsUserCode(v string) *QuerySyncResultRequest {
	s.OpsUserCode = &v
	return s
}

func (s *QuerySyncResultRequest) SetOpsUserId(v string) *QuerySyncResultRequest {
	s.OpsUserId = &v
	return s
}

func (s *QuerySyncResultRequest) SetTaskId(v int64) *QuerySyncResultRequest {
	s.TaskId = &v
	return s
}

func (s *QuerySyncResultRequest) SetTenantId(v string) *QuerySyncResultRequest {
	s.TenantId = &v
	return s
}

func (s *QuerySyncResultRequest) Validate() error {
	return dara.Validate(s)
}
