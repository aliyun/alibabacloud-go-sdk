// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAiOutboundTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartAiOutboundTaskResponseBody
	GetCode() *string
	SetMessage(v string) *StartAiOutboundTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartAiOutboundTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartAiOutboundTaskResponseBody
	GetSuccess() *bool
}

type StartAiOutboundTaskResponseBody struct {
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
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

func (s StartAiOutboundTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAiOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartAiOutboundTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartAiOutboundTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartAiOutboundTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAiOutboundTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartAiOutboundTaskResponseBody) SetCode(v string) *StartAiOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StartAiOutboundTaskResponseBody) SetMessage(v string) *StartAiOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StartAiOutboundTaskResponseBody) SetRequestId(v string) *StartAiOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAiOutboundTaskResponseBody) SetSuccess(v bool) *StartAiOutboundTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StartAiOutboundTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
