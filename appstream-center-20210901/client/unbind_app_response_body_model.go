// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnbindAppResponseBody
	GetCode() *string
	SetMessage(v string) *UnbindAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnbindAppResponseBody
	GetSuccess() *bool
}

type UnbindAppResponseBody struct {
	// The error code. This field is returned when the request fails and may not be included in a successful response.
	//
	// example:
	//
	// InternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error or result message. This field describes the cause of a failed request and may not be included in a successful response.
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
	// Indicates whether the request was processed successfully. A value of true indicates that the request was processed successfully. To check whether an asynchronous task is completed, query the task status separately.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindAppResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnbindAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnbindAppResponseBody) SetCode(v string) *UnbindAppResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindAppResponseBody) SetMessage(v string) *UnbindAppResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindAppResponseBody) SetRequestId(v string) *UnbindAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindAppResponseBody) SetSuccess(v bool) *UnbindAppResponseBody {
	s.Success = &v
	return s
}

func (s *UnbindAppResponseBody) Validate() error {
	return dara.Validate(s)
}
