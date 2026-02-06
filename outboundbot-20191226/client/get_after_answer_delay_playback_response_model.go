// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAfterAnswerDelayPlaybackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAfterAnswerDelayPlaybackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAfterAnswerDelayPlaybackResponse
	GetStatusCode() *int32
	SetBody(v *GetAfterAnswerDelayPlaybackResponseBody) *GetAfterAnswerDelayPlaybackResponse
	GetBody() *GetAfterAnswerDelayPlaybackResponseBody
}

type GetAfterAnswerDelayPlaybackResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAfterAnswerDelayPlaybackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAfterAnswerDelayPlaybackResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAfterAnswerDelayPlaybackResponse) GoString() string {
	return s.String()
}

func (s *GetAfterAnswerDelayPlaybackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAfterAnswerDelayPlaybackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAfterAnswerDelayPlaybackResponse) GetBody() *GetAfterAnswerDelayPlaybackResponseBody {
	return s.Body
}

func (s *GetAfterAnswerDelayPlaybackResponse) SetHeaders(v map[string]*string) *GetAfterAnswerDelayPlaybackResponse {
	s.Headers = v
	return s
}

func (s *GetAfterAnswerDelayPlaybackResponse) SetStatusCode(v int32) *GetAfterAnswerDelayPlaybackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAfterAnswerDelayPlaybackResponse) SetBody(v *GetAfterAnswerDelayPlaybackResponseBody) *GetAfterAnswerDelayPlaybackResponse {
	s.Body = v
	return s
}

func (s *GetAfterAnswerDelayPlaybackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
