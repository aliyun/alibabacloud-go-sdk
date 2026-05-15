// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiVoiceAgentDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAiVoiceAgentDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAiVoiceAgentDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryAiVoiceAgentDetailResponseBody) *QueryAiVoiceAgentDetailResponse
	GetBody() *QueryAiVoiceAgentDetailResponseBody
}

type QueryAiVoiceAgentDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAiVoiceAgentDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAiVoiceAgentDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAiVoiceAgentDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryAiVoiceAgentDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAiVoiceAgentDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAiVoiceAgentDetailResponse) GetBody() *QueryAiVoiceAgentDetailResponseBody {
	return s.Body
}

func (s *QueryAiVoiceAgentDetailResponse) SetHeaders(v map[string]*string) *QueryAiVoiceAgentDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponse) SetStatusCode(v int32) *QueryAiVoiceAgentDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAiVoiceAgentDetailResponse) SetBody(v *QueryAiVoiceAgentDetailResponseBody) *QueryAiVoiceAgentDetailResponse {
	s.Body = v
	return s
}

func (s *QueryAiVoiceAgentDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
