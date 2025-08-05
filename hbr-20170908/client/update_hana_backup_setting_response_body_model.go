// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHanaBackupSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateHanaBackupSettingResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateHanaBackupSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHanaBackupSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateHanaBackupSettingResponseBody
	GetSuccess() *bool
}

type UpdateHanaBackupSettingResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4892D474-9A4A-5298-BCD3-E46112A1EFD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateHanaBackupSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHanaBackupSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHanaBackupSettingResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateHanaBackupSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHanaBackupSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHanaBackupSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateHanaBackupSettingResponseBody) SetCode(v string) *UpdateHanaBackupSettingResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHanaBackupSettingResponseBody) SetMessage(v string) *UpdateHanaBackupSettingResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHanaBackupSettingResponseBody) SetRequestId(v string) *UpdateHanaBackupSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHanaBackupSettingResponseBody) SetSuccess(v bool) *UpdateHanaBackupSettingResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateHanaBackupSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
