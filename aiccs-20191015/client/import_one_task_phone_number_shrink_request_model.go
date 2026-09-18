// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportOneTaskPhoneNumberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptionType(v int64) *ImportOneTaskPhoneNumberShrinkRequest
	GetEncryptionType() *int64
	SetExtension(v string) *ImportOneTaskPhoneNumberShrinkRequest
	GetExtension() *string
	SetOutId(v string) *ImportOneTaskPhoneNumberShrinkRequest
	GetOutId() *string
	SetOwnerId(v int64) *ImportOneTaskPhoneNumberShrinkRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *ImportOneTaskPhoneNumberShrinkRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *ImportOneTaskPhoneNumberShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ImportOneTaskPhoneNumberShrinkRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *ImportOneTaskPhoneNumberShrinkRequest
	GetTaskId() *int64
	SetVariablesShrink(v string) *ImportOneTaskPhoneNumberShrinkRequest
	GetVariablesShrink() *string
}

type ImportOneTaskPhoneNumberShrinkRequest struct {
	// example:
	//
	// 81
	EncryptionType *int64 `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The extension number.
	//
	// example:
	//
	// 示例值示例值示例值
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The external serial number. We recommend that you use a unique ID. The value cannot exceed 128 characters.
	//
	// example:
	//
	// 94ba739b-xxxx-ef91-335d-4be006c34899
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The called phone number.
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
	// The variable list in Map format.
	//
	// > Variable format for engine-based voice call tasks:
	//
	// > - {"startWordParam.VariableKey1":"VariableValue1","promptParam.VariableKey2":"VariableValue2","bizParam.VariableKey3":"VariableValue3"}
	//
	// example:
	//
	// {"VariableKey1":"VariableValue1","VariableKey2":"VariableValue2"}
	//
	// For example values of engine-based voice call tasks, refer to the description on the left
	VariablesShrink *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s ImportOneTaskPhoneNumberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportOneTaskPhoneNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) GetEncryptionType() *int64 {
	return s.EncryptionType
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) GetExtension() *string {
	return s.Extension
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) GetOutId() *string {
	return s.OutId
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) GetVariablesShrink() *string {
	return s.VariablesShrink
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) SetEncryptionType(v int64) *ImportOneTaskPhoneNumberShrinkRequest {
	s.EncryptionType = &v
	return s
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) SetExtension(v string) *ImportOneTaskPhoneNumberShrinkRequest {
	s.Extension = &v
	return s
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) SetOutId(v string) *ImportOneTaskPhoneNumberShrinkRequest {
	s.OutId = &v
	return s
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) SetOwnerId(v int64) *ImportOneTaskPhoneNumberShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) SetPhoneNumber(v string) *ImportOneTaskPhoneNumberShrinkRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) SetResourceOwnerAccount(v string) *ImportOneTaskPhoneNumberShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) SetResourceOwnerId(v int64) *ImportOneTaskPhoneNumberShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) SetTaskId(v int64) *ImportOneTaskPhoneNumberShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) SetVariablesShrink(v string) *ImportOneTaskPhoneNumberShrinkRequest {
	s.VariablesShrink = &v
	return s
}

func (s *ImportOneTaskPhoneNumberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
