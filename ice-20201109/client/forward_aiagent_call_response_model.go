// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForwardAIAgentCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ForwardAIAgentCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ForwardAIAgentCallResponse
	GetStatusCode() *int32
	SetBody(v *ForwardAIAgentCallResponseBody) *ForwardAIAgentCallResponse
	GetBody() *ForwardAIAgentCallResponseBody
}

type ForwardAIAgentCallResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForwardAIAgentCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ForwardAIAgentCallResponse) String() string {
	return dara.Prettify(s)
}

func (s ForwardAIAgentCallResponse) GoString() string {
	return s.String()
}

func (s *ForwardAIAgentCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ForwardAIAgentCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ForwardAIAgentCallResponse) GetBody() *ForwardAIAgentCallResponseBody {
	return s.Body
}

func (s *ForwardAIAgentCallResponse) SetHeaders(v map[string]*string) *ForwardAIAgentCallResponse {
	s.Headers = v
	return s
}

func (s *ForwardAIAgentCallResponse) SetStatusCode(v int32) *ForwardAIAgentCallResponse {
	s.StatusCode = &v
	return s
}

func (s *ForwardAIAgentCallResponse) SetBody(v *ForwardAIAgentCallResponseBody) *ForwardAIAgentCallResponse {
	s.Body = v
	return s
}

func (s *ForwardAIAgentCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
