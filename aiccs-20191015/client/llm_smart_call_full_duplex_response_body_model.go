// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmSmartCallFullDuplexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *LlmSmartCallFullDuplexResponseBody
	GetCallId() *string
	SetCode(v string) *LlmSmartCallFullDuplexResponseBody
	GetCode() *string
	SetMessage(v string) *LlmSmartCallFullDuplexResponseBody
	GetMessage() *string
	SetRequestId(v string) *LlmSmartCallFullDuplexResponseBody
	GetRequestId() *string
}

type LlmSmartCallFullDuplexResponseBody struct {
	// example:
	//
	// 153955119976^140696759976
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LlmSmartCallFullDuplexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallFullDuplexResponseBody) GoString() string {
	return s.String()
}

func (s *LlmSmartCallFullDuplexResponseBody) GetCallId() *string {
	return s.CallId
}

func (s *LlmSmartCallFullDuplexResponseBody) GetCode() *string {
	return s.Code
}

func (s *LlmSmartCallFullDuplexResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LlmSmartCallFullDuplexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LlmSmartCallFullDuplexResponseBody) SetCallId(v string) *LlmSmartCallFullDuplexResponseBody {
	s.CallId = &v
	return s
}

func (s *LlmSmartCallFullDuplexResponseBody) SetCode(v string) *LlmSmartCallFullDuplexResponseBody {
	s.Code = &v
	return s
}

func (s *LlmSmartCallFullDuplexResponseBody) SetMessage(v string) *LlmSmartCallFullDuplexResponseBody {
	s.Message = &v
	return s
}

func (s *LlmSmartCallFullDuplexResponseBody) SetRequestId(v string) *LlmSmartCallFullDuplexResponseBody {
	s.RequestId = &v
	return s
}

func (s *LlmSmartCallFullDuplexResponseBody) Validate() error {
	return dara.Validate(s)
}
