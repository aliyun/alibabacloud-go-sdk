// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTodoDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetTodoDetailRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *GetTodoDetailRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetTodoDetailRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetTodoDetailRequest
	GetTenantId() *string
	SetTodoId(v string) *GetTodoDetailRequest
	GetTodoId() *string
}

type GetTodoDetailRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the to-do item.
	//
	// This parameter is required.
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
}

func (s GetTodoDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTodoDetailRequest) GoString() string {
	return s.String()
}

func (s *GetTodoDetailRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetTodoDetailRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetTodoDetailRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetTodoDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetTodoDetailRequest) GetTodoId() *string {
	return s.TodoId
}

func (s *GetTodoDetailRequest) SetAppToken(v string) *GetTodoDetailRequest {
	s.AppToken = &v
	return s
}

func (s *GetTodoDetailRequest) SetOpsUserCode(v string) *GetTodoDetailRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetTodoDetailRequest) SetOpsUserId(v string) *GetTodoDetailRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetTodoDetailRequest) SetTenantId(v string) *GetTodoDetailRequest {
	s.TenantId = &v
	return s
}

func (s *GetTodoDetailRequest) SetTodoId(v string) *GetTodoDetailRequest {
	s.TodoId = &v
	return s
}

func (s *GetTodoDetailRequest) Validate() error {
	return dara.Validate(s)
}
