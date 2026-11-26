// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarkTodosReadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *MarkTodosReadRequest
	GetAppToken() *string
	SetMarkAllRead(v bool) *MarkTodosReadRequest
	GetMarkAllRead() *bool
	SetOpsUserCode(v string) *MarkTodosReadRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *MarkTodosReadRequest
	GetOpsUserId() *string
	SetTenantId(v string) *MarkTodosReadRequest
	GetTenantId() *string
	SetTodoIds(v []*string) *MarkTodosReadRequest
	GetTodoIds() []*string
}

type MarkTodosReadRequest struct {
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
	TodoIds []*string `json:"todoIds,omitempty" xml:"todoIds,omitempty" type:"Repeated"`
}

func (s MarkTodosReadRequest) String() string {
	return dara.Prettify(s)
}

func (s MarkTodosReadRequest) GoString() string {
	return s.String()
}

func (s *MarkTodosReadRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *MarkTodosReadRequest) GetMarkAllRead() *bool {
	return s.MarkAllRead
}

func (s *MarkTodosReadRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *MarkTodosReadRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *MarkTodosReadRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *MarkTodosReadRequest) GetTodoIds() []*string {
	return s.TodoIds
}

func (s *MarkTodosReadRequest) SetAppToken(v string) *MarkTodosReadRequest {
	s.AppToken = &v
	return s
}

func (s *MarkTodosReadRequest) SetMarkAllRead(v bool) *MarkTodosReadRequest {
	s.MarkAllRead = &v
	return s
}

func (s *MarkTodosReadRequest) SetOpsUserCode(v string) *MarkTodosReadRequest {
	s.OpsUserCode = &v
	return s
}

func (s *MarkTodosReadRequest) SetOpsUserId(v string) *MarkTodosReadRequest {
	s.OpsUserId = &v
	return s
}

func (s *MarkTodosReadRequest) SetTenantId(v string) *MarkTodosReadRequest {
	s.TenantId = &v
	return s
}

func (s *MarkTodosReadRequest) SetTodoIds(v []*string) *MarkTodosReadRequest {
	s.TodoIds = v
	return s
}

func (s *MarkTodosReadRequest) Validate() error {
	return dara.Validate(s)
}
