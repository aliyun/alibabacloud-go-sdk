// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTodoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetTodoRequest
	GetAppToken() *string
	SetMessageId(v string) *GetTodoRequest
	GetMessageId() *string
	SetOpsUserCode(v string) *GetTodoRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetTodoRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetTodoRequest
	GetTenantId() *string
	SetTodoId(v string) *GetTodoRequest
	GetTodoId() *string
	SetWorkItemId(v string) *GetTodoRequest
	GetWorkItemId() *string
}

type GetTodoRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the assistant message. This parameter is used to query multiple to-do items by message.
	//
	// example:
	//
	// msg-1
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
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
	// The ID of the to-do item (get).
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
	// The ID of the work item. This parameter is used to query multiple to-do items by source.
	//
	// example:
	//
	// wi_1
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s GetTodoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTodoRequest) GoString() string {
	return s.String()
}

func (s *GetTodoRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetTodoRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *GetTodoRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetTodoRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetTodoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetTodoRequest) GetTodoId() *string {
	return s.TodoId
}

func (s *GetTodoRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetTodoRequest) SetAppToken(v string) *GetTodoRequest {
	s.AppToken = &v
	return s
}

func (s *GetTodoRequest) SetMessageId(v string) *GetTodoRequest {
	s.MessageId = &v
	return s
}

func (s *GetTodoRequest) SetOpsUserCode(v string) *GetTodoRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetTodoRequest) SetOpsUserId(v string) *GetTodoRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetTodoRequest) SetTenantId(v string) *GetTodoRequest {
	s.TenantId = &v
	return s
}

func (s *GetTodoRequest) SetTodoId(v string) *GetTodoRequest {
	s.TodoId = &v
	return s
}

func (s *GetTodoRequest) SetWorkItemId(v string) *GetTodoRequest {
	s.WorkItemId = &v
	return s
}

func (s *GetTodoRequest) Validate() error {
	return dara.Validate(s)
}
