// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateAiOutboundTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TerminateAiOutboundTaskResponseBody
	GetCode() *string
	SetMessage(v string) *TerminateAiOutboundTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *TerminateAiOutboundTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TerminateAiOutboundTaskResponseBody
	GetSuccess() *bool
}

type TerminateAiOutboundTaskResponseBody struct {
	// Status code.
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
	// Request ID.
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

func (s TerminateAiOutboundTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TerminateAiOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateAiOutboundTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *TerminateAiOutboundTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TerminateAiOutboundTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TerminateAiOutboundTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TerminateAiOutboundTaskResponseBody) SetCode(v string) *TerminateAiOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *TerminateAiOutboundTaskResponseBody) SetMessage(v string) *TerminateAiOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *TerminateAiOutboundTaskResponseBody) SetRequestId(v string) *TerminateAiOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminateAiOutboundTaskResponseBody) SetSuccess(v bool) *TerminateAiOutboundTaskResponseBody {
	s.Success = &v
	return s
}

func (s *TerminateAiOutboundTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
