// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateImageResponseBody
	GetCode() *string
	SetData(v bool) *UpdateImageResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateImageResponseBody
	GetSuccess() *bool
}

type UpdateImageResponseBody struct {
	// The error code. This field is returned when the request fails. A successful response may not contain this field.
	//
	// example:
	//
	// Image.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the operation was successful. A value of true indicates that the operation was successful.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error or result message. This field describes the cause when the request fails. A successful response may not contain this field.
	//
	// example:
	//
	// The image does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can provide this value for troubleshooting.
	//
	// example:
	//
	// 01A1A962-5503-107B-AC8F-EBA85DF905AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was processed successfully. A value of true indicates that the request was processed successfully. To check whether an asynchronous task is completed, you must query the task status separately.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateImageResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateImageResponseBody) SetCode(v string) *UpdateImageResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateImageResponseBody) SetData(v bool) *UpdateImageResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateImageResponseBody) SetMessage(v string) *UpdateImageResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateImageResponseBody) SetRequestId(v string) *UpdateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageResponseBody) SetSuccess(v bool) *UpdateImageResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateImageResponseBody) Validate() error {
	return dara.Validate(s)
}
