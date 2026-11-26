// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCancelTodosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *BatchCancelTodosRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *BatchCancelTodosRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *BatchCancelTodosRequest
	GetOpsUserId() *string
	SetTenantId(v string) *BatchCancelTodosRequest
	GetTenantId() *string
	SetTodoIds(v []*string) *BatchCancelTodosRequest
	GetTodoIds() []*string
}

type BatchCancelTodosRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user specified by opsUserId or opsUserCode.
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
	// The list of to-do item IDs. You can specify 1 to 100 IDs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The list of to-do item IDs. You can specify 1 to 100 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["todo_abc123"]
	TodoIds []*string `json:"todoIds,omitempty" xml:"todoIds,omitempty" type:"Repeated"`
}

func (s BatchCancelTodosRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCancelTodosRequest) GoString() string {
	return s.String()
}

func (s *BatchCancelTodosRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *BatchCancelTodosRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *BatchCancelTodosRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *BatchCancelTodosRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *BatchCancelTodosRequest) GetTodoIds() []*string {
	return s.TodoIds
}

func (s *BatchCancelTodosRequest) SetAppToken(v string) *BatchCancelTodosRequest {
	s.AppToken = &v
	return s
}

func (s *BatchCancelTodosRequest) SetOpsUserCode(v string) *BatchCancelTodosRequest {
	s.OpsUserCode = &v
	return s
}

func (s *BatchCancelTodosRequest) SetOpsUserId(v string) *BatchCancelTodosRequest {
	s.OpsUserId = &v
	return s
}

func (s *BatchCancelTodosRequest) SetTenantId(v string) *BatchCancelTodosRequest {
	s.TenantId = &v
	return s
}

func (s *BatchCancelTodosRequest) SetTodoIds(v []*string) *BatchCancelTodosRequest {
	s.TodoIds = v
	return s
}

func (s *BatchCancelTodosRequest) Validate() error {
	return dara.Validate(s)
}
