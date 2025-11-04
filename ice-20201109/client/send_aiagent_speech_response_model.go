// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAIAgentSpeechResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendAIAgentSpeechResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendAIAgentSpeechResponse
	GetStatusCode() *int32
	SetBody(v *SendAIAgentSpeechResponseBody) *SendAIAgentSpeechResponse
	GetBody() *SendAIAgentSpeechResponseBody
}

type SendAIAgentSpeechResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendAIAgentSpeechResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendAIAgentSpeechResponse) String() string {
	return dara.Prettify(s)
}

func (s SendAIAgentSpeechResponse) GoString() string {
	return s.String()
}

func (s *SendAIAgentSpeechResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendAIAgentSpeechResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendAIAgentSpeechResponse) GetBody() *SendAIAgentSpeechResponseBody {
	return s.Body
}

func (s *SendAIAgentSpeechResponse) SetHeaders(v map[string]*string) *SendAIAgentSpeechResponse {
	s.Headers = v
	return s
}

func (s *SendAIAgentSpeechResponse) SetStatusCode(v int32) *SendAIAgentSpeechResponse {
	s.StatusCode = &v
	return s
}

func (s *SendAIAgentSpeechResponse) SetBody(v *SendAIAgentSpeechResponseBody) *SendAIAgentSpeechResponse {
	s.Body = v
	return s
}

func (s *SendAIAgentSpeechResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
