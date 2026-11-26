// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTodoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *TransferTodoShrinkRequest
	GetAppToken() *string
	SetComment(v string) *TransferTodoShrinkRequest
	GetComment() *string
	SetOpsUserCode(v string) *TransferTodoShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *TransferTodoShrinkRequest
	GetOpsUserId() *string
	SetTargetUserIdsShrink(v string) *TransferTodoShrinkRequest
	GetTargetUserIdsShrink() *string
	SetTenantId(v string) *TransferTodoShrinkRequest
	GetTenantId() *string
	SetTodoId(v string) *TransferTodoShrinkRequest
	GetTodoId() *string
}

type TransferTodoShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The handover note.
	//
	// example:
	//
	// Please continue processing
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
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
	// The list of target user IDs for the transfer (1 to 100).
	//
	// This parameter is required.
	//
	// example:
	//
	// 43
	TargetUserIdsShrink *string `json:"targetUserIds,omitempty" xml:"targetUserIds,omitempty"`
	// The to-do item ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The list of target user IDs for the transfer (1 to 100).
	//
	// This parameter is required.
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
}

func (s TransferTodoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferTodoShrinkRequest) GoString() string {
	return s.String()
}

func (s *TransferTodoShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *TransferTodoShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *TransferTodoShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *TransferTodoShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *TransferTodoShrinkRequest) GetTargetUserIdsShrink() *string {
	return s.TargetUserIdsShrink
}

func (s *TransferTodoShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *TransferTodoShrinkRequest) GetTodoId() *string {
	return s.TodoId
}

func (s *TransferTodoShrinkRequest) SetAppToken(v string) *TransferTodoShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *TransferTodoShrinkRequest) SetComment(v string) *TransferTodoShrinkRequest {
	s.Comment = &v
	return s
}

func (s *TransferTodoShrinkRequest) SetOpsUserCode(v string) *TransferTodoShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *TransferTodoShrinkRequest) SetOpsUserId(v string) *TransferTodoShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *TransferTodoShrinkRequest) SetTargetUserIdsShrink(v string) *TransferTodoShrinkRequest {
	s.TargetUserIdsShrink = &v
	return s
}

func (s *TransferTodoShrinkRequest) SetTenantId(v string) *TransferTodoShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *TransferTodoShrinkRequest) SetTodoId(v string) *TransferTodoShrinkRequest {
	s.TodoId = &v
	return s
}

func (s *TransferTodoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
