// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAfterAnswerDelayPlaybackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveAfterAnswerDelayPlaybackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveAfterAnswerDelayPlaybackResponse
	GetStatusCode() *int32
	SetBody(v *SaveAfterAnswerDelayPlaybackResponseBody) *SaveAfterAnswerDelayPlaybackResponse
	GetBody() *SaveAfterAnswerDelayPlaybackResponseBody
}

type SaveAfterAnswerDelayPlaybackResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveAfterAnswerDelayPlaybackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveAfterAnswerDelayPlaybackResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveAfterAnswerDelayPlaybackResponse) GoString() string {
	return s.String()
}

func (s *SaveAfterAnswerDelayPlaybackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveAfterAnswerDelayPlaybackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveAfterAnswerDelayPlaybackResponse) GetBody() *SaveAfterAnswerDelayPlaybackResponseBody {
	return s.Body
}

func (s *SaveAfterAnswerDelayPlaybackResponse) SetHeaders(v map[string]*string) *SaveAfterAnswerDelayPlaybackResponse {
	s.Headers = v
	return s
}

func (s *SaveAfterAnswerDelayPlaybackResponse) SetStatusCode(v int32) *SaveAfterAnswerDelayPlaybackResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveAfterAnswerDelayPlaybackResponse) SetBody(v *SaveAfterAnswerDelayPlaybackResponseBody) *SaveAfterAnswerDelayPlaybackResponse {
	s.Body = v
	return s
}

func (s *SaveAfterAnswerDelayPlaybackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
