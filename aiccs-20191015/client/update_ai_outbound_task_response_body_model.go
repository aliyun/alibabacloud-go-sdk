// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiOutboundTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAiOutboundTaskResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateAiOutboundTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAiOutboundTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAiOutboundTaskResponseBody
	GetSuccess() *bool
}

type UpdateAiOutboundTaskResponseBody struct {
	// The status code.
	//
	// example:
	//
	// ok
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API invocation succeeded.
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAiOutboundTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAiOutboundTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAiOutboundTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAiOutboundTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAiOutboundTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAiOutboundTaskResponseBody) SetCode(v string) *UpdateAiOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAiOutboundTaskResponseBody) SetMessage(v string) *UpdateAiOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAiOutboundTaskResponseBody) SetRequestId(v string) *UpdateAiOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAiOutboundTaskResponseBody) SetSuccess(v bool) *UpdateAiOutboundTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAiOutboundTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
