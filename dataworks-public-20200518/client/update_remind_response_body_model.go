// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRemindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateRemindResponseBody
	GetData() *bool
	SetErrorCode(v string) *UpdateRemindResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateRemindResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateRemindResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateRemindResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRemindResponseBody
	GetSuccess() *bool
}

type UpdateRemindResponseBody struct {
	// Indicates whether the modification to the custom alert rule succeeds.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRemindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRemindResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRemindResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateRemindResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateRemindResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateRemindResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateRemindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRemindResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRemindResponseBody) SetData(v bool) *UpdateRemindResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRemindResponseBody) SetErrorCode(v string) *UpdateRemindResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateRemindResponseBody) SetErrorMessage(v string) *UpdateRemindResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateRemindResponseBody) SetHttpStatusCode(v int32) *UpdateRemindResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateRemindResponseBody) SetRequestId(v string) *UpdateRemindResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRemindResponseBody) SetSuccess(v bool) *UpdateRemindResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRemindResponseBody) Validate() error {
	return dara.Validate(s)
}
