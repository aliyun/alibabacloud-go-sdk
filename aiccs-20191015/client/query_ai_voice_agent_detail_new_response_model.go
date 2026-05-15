// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiVoiceAgentDetailNewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAiVoiceAgentDetailNewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAiVoiceAgentDetailNewResponse
	GetStatusCode() *int32
	SetBody(v *QueryAiVoiceAgentDetailNewResponseBody) *QueryAiVoiceAgentDetailNewResponse
	GetBody() *QueryAiVoiceAgentDetailNewResponseBody
}

type QueryAiVoiceAgentDetailNewResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAiVoiceAgentDetailNewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAiVoiceAgentDetailNewResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailNewResponse) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailNewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAiVoiceAgentDetailNewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAiVoiceAgentDetailNewResponse) GetBody() *QueryAiVoiceAgentDetailNewResponseBody {
	return s.Body
}

func (s *QueryAiVoiceAgentDetailNewResponse) SetHeaders(v map[string]*string) *QueryAiVoiceAgentDetailNewResponse {
	s.Headers = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponse) SetStatusCode(v int32) *QueryAiVoiceAgentDetailNewResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponse) SetBody(v *QueryAiVoiceAgentDetailNewResponseBody) *QueryAiVoiceAgentDetailNewResponse {
	s.Body = v
	return s
}

func (s *QueryAiVoiceAgentDetailNewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
