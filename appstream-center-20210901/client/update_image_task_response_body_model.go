// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateImageTaskResponseBody
	GetCode() *string
	SetData(v bool) *UpdateImageTaskResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateImageTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateImageTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateImageTaskResponseBody
	GetSuccess() *bool
}

type UpdateImageTaskResponseBody struct {
	// The error code. Returned when the request fails. A successful response may not contain this field.
	//
	// example:
	//
	// InternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the operation is successful. A value of true indicates that the operation is successful.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error or result message. Used to describe the cause when the request fails. A successful response may not contain this field.
	//
	// example:
	//
	// Internal error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can provide this value for troubleshooting.
	//
	// example:
	//
	// 01A1A962-5503-107B-AC8F-EBA85DF905AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. A value of true indicates that the request is successful. To check whether an asynchronous task is completed, you must query the task status separately.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateImageTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateImageTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateImageTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateImageTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateImageTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateImageTaskResponseBody) SetCode(v string) *UpdateImageTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateImageTaskResponseBody) SetData(v bool) *UpdateImageTaskResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateImageTaskResponseBody) SetMessage(v string) *UpdateImageTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateImageTaskResponseBody) SetRequestId(v string) *UpdateImageTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageTaskResponseBody) SetSuccess(v bool) *UpdateImageTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateImageTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
