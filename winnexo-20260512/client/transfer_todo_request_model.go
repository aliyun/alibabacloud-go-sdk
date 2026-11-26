// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTodoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *TransferTodoRequest
	GetAppToken() *string
	SetComment(v string) *TransferTodoRequest
	GetComment() *string
	SetOpsUserCode(v string) *TransferTodoRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *TransferTodoRequest
	GetOpsUserId() *string
	SetTargetUserIds(v []*string) *TransferTodoRequest
	GetTargetUserIds() []*string
	SetTenantId(v string) *TransferTodoRequest
	GetTenantId() *string
	SetTodoId(v string) *TransferTodoRequest
	GetTodoId() *string
}

type TransferTodoRequest struct {
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
	TargetUserIds []*string `json:"targetUserIds,omitempty" xml:"targetUserIds,omitempty" type:"Repeated"`
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

func (s TransferTodoRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferTodoRequest) GoString() string {
	return s.String()
}

func (s *TransferTodoRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *TransferTodoRequest) GetComment() *string {
	return s.Comment
}

func (s *TransferTodoRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *TransferTodoRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *TransferTodoRequest) GetTargetUserIds() []*string {
	return s.TargetUserIds
}

func (s *TransferTodoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *TransferTodoRequest) GetTodoId() *string {
	return s.TodoId
}

func (s *TransferTodoRequest) SetAppToken(v string) *TransferTodoRequest {
	s.AppToken = &v
	return s
}

func (s *TransferTodoRequest) SetComment(v string) *TransferTodoRequest {
	s.Comment = &v
	return s
}

func (s *TransferTodoRequest) SetOpsUserCode(v string) *TransferTodoRequest {
	s.OpsUserCode = &v
	return s
}

func (s *TransferTodoRequest) SetOpsUserId(v string) *TransferTodoRequest {
	s.OpsUserId = &v
	return s
}

func (s *TransferTodoRequest) SetTargetUserIds(v []*string) *TransferTodoRequest {
	s.TargetUserIds = v
	return s
}

func (s *TransferTodoRequest) SetTenantId(v string) *TransferTodoRequest {
	s.TenantId = &v
	return s
}

func (s *TransferTodoRequest) SetTodoId(v string) *TransferTodoRequest {
	s.TodoId = &v
	return s
}

func (s *TransferTodoRequest) Validate() error {
	return dara.Validate(s)
}
