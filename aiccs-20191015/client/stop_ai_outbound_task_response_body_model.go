// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAiOutboundTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopAiOutboundTaskResponseBody
	GetCode() *string
	SetMessage(v string) *StopAiOutboundTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopAiOutboundTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopAiOutboundTaskResponseBody
	GetSuccess() *bool
}

type StopAiOutboundTaskResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopAiOutboundTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopAiOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopAiOutboundTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopAiOutboundTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopAiOutboundTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopAiOutboundTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopAiOutboundTaskResponseBody) SetCode(v string) *StopAiOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopAiOutboundTaskResponseBody) SetMessage(v string) *StopAiOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopAiOutboundTaskResponseBody) SetRequestId(v string) *StopAiOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopAiOutboundTaskResponseBody) SetSuccess(v bool) *StopAiOutboundTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StopAiOutboundTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
