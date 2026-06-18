// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportOneTaskPhoneNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptionType(v int64) *ImportOneTaskPhoneNumberRequest
	GetEncryptionType() *int64
	SetOutId(v string) *ImportOneTaskPhoneNumberRequest
	GetOutId() *string
	SetOwnerId(v int64) *ImportOneTaskPhoneNumberRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *ImportOneTaskPhoneNumberRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *ImportOneTaskPhoneNumberRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ImportOneTaskPhoneNumberRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *ImportOneTaskPhoneNumberRequest
	GetTaskId() *int64
	SetVariables(v map[string]interface{}) *ImportOneTaskPhoneNumberRequest
	GetVariables() map[string]interface{}
}

type ImportOneTaskPhoneNumberRequest struct {
	// example:
	//
	// 81
	EncryptionType *int64 `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The external ID. We recommend that you use a unique ID to ensure idempotency. The value cannot exceed 128 characters.
	//
	// example:
	//
	// 94ba739b-xxxx-ef91-335d-4be006c34899
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number of the callee.
	//
	// This parameter is required.
	//
	// example:
	//
	// 180******
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1231231212****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// A list of variables in a map.
	//
	// > The format of variables for an engine-based call task is as follows:
	//
	// >
	//
	// > - {"startWordParam.variable_key1":"variable_value1","promptParam.variable_key2":"variable_value2","bizParam.variable_key3":"variable_value3"}
	//
	// example:
	//
	// {"变量key1":"变量值1","变量key2":"变量值2"}
	//
	// 引擎呼叫任务示例值请看左侧描述
	Variables map[string]interface{} `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s ImportOneTaskPhoneNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportOneTaskPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *ImportOneTaskPhoneNumberRequest) GetEncryptionType() *int64 {
	return s.EncryptionType
}

func (s *ImportOneTaskPhoneNumberRequest) GetOutId() *string {
	return s.OutId
}

func (s *ImportOneTaskPhoneNumberRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ImportOneTaskPhoneNumberRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ImportOneTaskPhoneNumberRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ImportOneTaskPhoneNumberRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ImportOneTaskPhoneNumberRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ImportOneTaskPhoneNumberRequest) GetVariables() map[string]interface{} {
	return s.Variables
}

func (s *ImportOneTaskPhoneNumberRequest) SetEncryptionType(v int64) *ImportOneTaskPhoneNumberRequest {
	s.EncryptionType = &v
	return s
}

func (s *ImportOneTaskPhoneNumberRequest) SetOutId(v string) *ImportOneTaskPhoneNumberRequest {
	s.OutId = &v
	return s
}

func (s *ImportOneTaskPhoneNumberRequest) SetOwnerId(v int64) *ImportOneTaskPhoneNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportOneTaskPhoneNumberRequest) SetPhoneNumber(v string) *ImportOneTaskPhoneNumberRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ImportOneTaskPhoneNumberRequest) SetResourceOwnerAccount(v string) *ImportOneTaskPhoneNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportOneTaskPhoneNumberRequest) SetResourceOwnerId(v int64) *ImportOneTaskPhoneNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportOneTaskPhoneNumberRequest) SetTaskId(v int64) *ImportOneTaskPhoneNumberRequest {
	s.TaskId = &v
	return s
}

func (s *ImportOneTaskPhoneNumberRequest) SetVariables(v map[string]interface{}) *ImportOneTaskPhoneNumberRequest {
	s.Variables = v
	return s
}

func (s *ImportOneTaskPhoneNumberRequest) Validate() error {
	return dara.Validate(s)
}
