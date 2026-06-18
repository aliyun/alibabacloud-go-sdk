// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmSmartCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *LlmSmartCallResponseBody
	GetCallId() *string
	SetCode(v string) *LlmSmartCallResponseBody
	GetCode() *string
	SetMessage(v string) *LlmSmartCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *LlmSmartCallResponseBody
	GetRequestId() *string
}

type LlmSmartCallResponseBody struct {
	// Unique receipt ID for this call.
	//
	// example:
	//
	// 125165515***^11195613****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// Request status code. A return value of "OK" indicates that the request succeeded.
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
	// Id of the request
	//
	// example:
	//
	// D6A51251-F7C4-596A-9F45-3C3219A5450D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LlmSmartCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallResponseBody) GoString() string {
	return s.String()
}

func (s *LlmSmartCallResponseBody) GetCallId() *string {
	return s.CallId
}

func (s *LlmSmartCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *LlmSmartCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LlmSmartCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LlmSmartCallResponseBody) SetCallId(v string) *LlmSmartCallResponseBody {
	s.CallId = &v
	return s
}

func (s *LlmSmartCallResponseBody) SetCode(v string) *LlmSmartCallResponseBody {
	s.Code = &v
	return s
}

func (s *LlmSmartCallResponseBody) SetMessage(v string) *LlmSmartCallResponseBody {
	s.Message = &v
	return s
}

func (s *LlmSmartCallResponseBody) SetRequestId(v string) *LlmSmartCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *LlmSmartCallResponseBody) Validate() error {
	return dara.Validate(s)
}
