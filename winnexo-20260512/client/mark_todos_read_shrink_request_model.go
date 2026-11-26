// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarkTodosReadShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *MarkTodosReadShrinkRequest
	GetAppToken() *string
	SetMarkAllRead(v bool) *MarkTodosReadShrinkRequest
	GetMarkAllRead() *bool
	SetOpsUserCode(v string) *MarkTodosReadShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *MarkTodosReadShrinkRequest
	GetOpsUserId() *string
	SetTenantId(v string) *MarkTodosReadShrinkRequest
	GetTenantId() *string
	SetTodoIdsShrink(v string) *MarkTodosReadShrinkRequest
	GetTodoIdsShrink() *string
}

type MarkTodosReadShrinkRequest struct {
	// The platform application token issued by the console. If you specify this parameter, the operation is performed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// Specifies whether to mark all to-do items as read. This parameter takes precedence over todoIds.
	//
	// example:
	//
	// false
	MarkAllRead *bool `json:"markAllRead,omitempty" xml:"markAllRead,omitempty"`
	// The code of the proxy user, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The ID of the proxy user, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The list of to-do IDs. You can specify up to 100 IDs. This parameter is ignored when markAllRead is set to true.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The list of to-do IDs. You can specify up to 100 IDs. This parameter is ignored when markAllRead is set to true.
	//
	// example:
	//
	// ["todo_abc123","todo_def456"]
	TodoIdsShrink *string `json:"todoIds,omitempty" xml:"todoIds,omitempty"`
}

func (s MarkTodosReadShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MarkTodosReadShrinkRequest) GoString() string {
	return s.String()
}

func (s *MarkTodosReadShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *MarkTodosReadShrinkRequest) GetMarkAllRead() *bool {
	return s.MarkAllRead
}

func (s *MarkTodosReadShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *MarkTodosReadShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *MarkTodosReadShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *MarkTodosReadShrinkRequest) GetTodoIdsShrink() *string {
	return s.TodoIdsShrink
}

func (s *MarkTodosReadShrinkRequest) SetAppToken(v string) *MarkTodosReadShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *MarkTodosReadShrinkRequest) SetMarkAllRead(v bool) *MarkTodosReadShrinkRequest {
	s.MarkAllRead = &v
	return s
}

func (s *MarkTodosReadShrinkRequest) SetOpsUserCode(v string) *MarkTodosReadShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *MarkTodosReadShrinkRequest) SetOpsUserId(v string) *MarkTodosReadShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *MarkTodosReadShrinkRequest) SetTenantId(v string) *MarkTodosReadShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *MarkTodosReadShrinkRequest) SetTodoIdsShrink(v string) *MarkTodosReadShrinkRequest {
	s.TodoIdsShrink = &v
	return s
}

func (s *MarkTodosReadShrinkRequest) Validate() error {
	return dara.Validate(s)
}
