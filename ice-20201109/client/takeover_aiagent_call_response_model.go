// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTakeoverAIAgentCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TakeoverAIAgentCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TakeoverAIAgentCallResponse
	GetStatusCode() *int32
	SetBody(v *TakeoverAIAgentCallResponseBody) *TakeoverAIAgentCallResponse
	GetBody() *TakeoverAIAgentCallResponseBody
}

type TakeoverAIAgentCallResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TakeoverAIAgentCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TakeoverAIAgentCallResponse) String() string {
	return dara.Prettify(s)
}

func (s TakeoverAIAgentCallResponse) GoString() string {
	return s.String()
}

func (s *TakeoverAIAgentCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TakeoverAIAgentCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TakeoverAIAgentCallResponse) GetBody() *TakeoverAIAgentCallResponseBody {
	return s.Body
}

func (s *TakeoverAIAgentCallResponse) SetHeaders(v map[string]*string) *TakeoverAIAgentCallResponse {
	s.Headers = v
	return s
}

func (s *TakeoverAIAgentCallResponse) SetStatusCode(v int32) *TakeoverAIAgentCallResponse {
	s.StatusCode = &v
	return s
}

func (s *TakeoverAIAgentCallResponse) SetBody(v *TakeoverAIAgentCallResponseBody) *TakeoverAIAgentCallResponse {
	s.Body = v
	return s
}

func (s *TakeoverAIAgentCallResponse) Validate() error {
	return dara.Validate(s)
}
