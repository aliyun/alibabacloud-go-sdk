// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCancelTodosShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *BatchCancelTodosShrinkRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *BatchCancelTodosShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *BatchCancelTodosShrinkRequest
	GetOpsUserId() *string
	SetTenantId(v string) *BatchCancelTodosShrinkRequest
	GetTenantId() *string
	SetTodoIdsShrink(v string) *BatchCancelTodosShrinkRequest
	GetTodoIdsShrink() *string
}

type BatchCancelTodosShrinkRequest struct {
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
	TodoIdsShrink *string `json:"todoIds,omitempty" xml:"todoIds,omitempty"`
}

func (s BatchCancelTodosShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCancelTodosShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCancelTodosShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *BatchCancelTodosShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *BatchCancelTodosShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *BatchCancelTodosShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *BatchCancelTodosShrinkRequest) GetTodoIdsShrink() *string {
	return s.TodoIdsShrink
}

func (s *BatchCancelTodosShrinkRequest) SetAppToken(v string) *BatchCancelTodosShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *BatchCancelTodosShrinkRequest) SetOpsUserCode(v string) *BatchCancelTodosShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *BatchCancelTodosShrinkRequest) SetOpsUserId(v string) *BatchCancelTodosShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *BatchCancelTodosShrinkRequest) SetTenantId(v string) *BatchCancelTodosShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *BatchCancelTodosShrinkRequest) SetTodoIdsShrink(v string) *BatchCancelTodosShrinkRequest {
	s.TodoIdsShrink = &v
	return s
}

func (s *BatchCancelTodosShrinkRequest) Validate() error {
	return dara.Validate(s)
}
