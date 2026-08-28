// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeAIAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvokeAIAgentResponseBody
	GetCode() *string
	SetData(v *InvokeAIAgentResponseBodyData) *InvokeAIAgentResponseBody
	GetData() *InvokeAIAgentResponseBodyData
	SetMessage(v string) *InvokeAIAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *InvokeAIAgentResponseBody
	GetRequestId() *string
}

type InvokeAIAgentResponseBody struct {
	// example:
	//
	// 200
	Code *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data *InvokeAIAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A60EE5CA-1294-532A-9775-8D2FD1C6EFBF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s InvokeAIAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeAIAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeAIAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvokeAIAgentResponseBody) GetData() *InvokeAIAgentResponseBodyData {
	return s.Data
}

func (s *InvokeAIAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvokeAIAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeAIAgentResponseBody) SetCode(v string) *InvokeAIAgentResponseBody {
	s.Code = &v
	return s
}

func (s *InvokeAIAgentResponseBody) SetData(v *InvokeAIAgentResponseBodyData) *InvokeAIAgentResponseBody {
	s.Data = v
	return s
}

func (s *InvokeAIAgentResponseBody) SetMessage(v string) *InvokeAIAgentResponseBody {
	s.Message = &v
	return s
}

func (s *InvokeAIAgentResponseBody) SetRequestId(v string) *InvokeAIAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeAIAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InvokeAIAgentResponseBodyData struct {
	// example:
	//
	// {"choices":[{"delta":{"content":"分析结果..."}}]}
	Body    *string            `json:"body,omitempty" xml:"body,omitempty"`
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
}

func (s InvokeAIAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InvokeAIAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *InvokeAIAgentResponseBodyData) GetBody() *string {
	return s.Body
}

func (s *InvokeAIAgentResponseBodyData) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeAIAgentResponseBodyData) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *InvokeAIAgentResponseBodyData) SetBody(v string) *InvokeAIAgentResponseBodyData {
	s.Body = &v
	return s
}

func (s *InvokeAIAgentResponseBodyData) SetHeaders(v map[string]*string) *InvokeAIAgentResponseBodyData {
	s.Headers = v
	return s
}

func (s *InvokeAIAgentResponseBodyData) SetHttpCode(v int32) *InvokeAIAgentResponseBodyData {
	s.HttpCode = &v
	return s
}

func (s *InvokeAIAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
