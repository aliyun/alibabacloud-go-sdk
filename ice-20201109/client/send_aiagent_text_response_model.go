// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAIAgentTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendAIAgentTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendAIAgentTextResponse
	GetStatusCode() *int32
	SetBody(v *SendAIAgentTextResponseBody) *SendAIAgentTextResponse
	GetBody() *SendAIAgentTextResponseBody
}

type SendAIAgentTextResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendAIAgentTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendAIAgentTextResponse) String() string {
	return dara.Prettify(s)
}

func (s SendAIAgentTextResponse) GoString() string {
	return s.String()
}

func (s *SendAIAgentTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendAIAgentTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendAIAgentTextResponse) GetBody() *SendAIAgentTextResponseBody {
	return s.Body
}

func (s *SendAIAgentTextResponse) SetHeaders(v map[string]*string) *SendAIAgentTextResponse {
	s.Headers = v
	return s
}

func (s *SendAIAgentTextResponse) SetStatusCode(v int32) *SendAIAgentTextResponse {
	s.StatusCode = &v
	return s
}

func (s *SendAIAgentTextResponse) SetBody(v *SendAIAgentTextResponseBody) *SendAIAgentTextResponse {
	s.Body = v
	return s
}

func (s *SendAIAgentTextResponse) Validate() error {
	return dara.Validate(s)
}
