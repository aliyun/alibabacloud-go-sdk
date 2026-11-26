// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTodoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CancelTodoRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *CancelTodoRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CancelTodoRequest
	GetOpsUserId() *string
	SetTenantId(v string) *CancelTodoRequest
	GetTenantId() *string
	SetTodoId(v string) *CancelTodoRequest
	GetTodoId() *string
	SetWorkItemId(v string) *CancelTodoRequest
	GetWorkItemId() *string
}

type CancelTodoRequest struct {
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
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the to-do item. You must specify either this parameter or workItemId.
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
	// The ID of the work item. You can use this parameter to cancel to-do items in batches by source. You must specify either this parameter or todoId.
	//
	// example:
	//
	// wi_1
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s CancelTodoRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelTodoRequest) GoString() string {
	return s.String()
}

func (s *CancelTodoRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CancelTodoRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CancelTodoRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CancelTodoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CancelTodoRequest) GetTodoId() *string {
	return s.TodoId
}

func (s *CancelTodoRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *CancelTodoRequest) SetAppToken(v string) *CancelTodoRequest {
	s.AppToken = &v
	return s
}

func (s *CancelTodoRequest) SetOpsUserCode(v string) *CancelTodoRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CancelTodoRequest) SetOpsUserId(v string) *CancelTodoRequest {
	s.OpsUserId = &v
	return s
}

func (s *CancelTodoRequest) SetTenantId(v string) *CancelTodoRequest {
	s.TenantId = &v
	return s
}

func (s *CancelTodoRequest) SetTodoId(v string) *CancelTodoRequest {
	s.TodoId = &v
	return s
}

func (s *CancelTodoRequest) SetWorkItemId(v string) *CancelTodoRequest {
	s.WorkItemId = &v
	return s
}

func (s *CancelTodoRequest) Validate() error {
	return dara.Validate(s)
}
