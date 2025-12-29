// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushWelcomeTextAndMusicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushWelcomeTextAndMusicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushWelcomeTextAndMusicResponse
	GetStatusCode() *int32
	SetBody(v *PushWelcomeTextAndMusicResponseBody) *PushWelcomeTextAndMusicResponse
	GetBody() *PushWelcomeTextAndMusicResponseBody
}

type PushWelcomeTextAndMusicResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushWelcomeTextAndMusicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushWelcomeTextAndMusicResponse) String() string {
	return dara.Prettify(s)
}

func (s PushWelcomeTextAndMusicResponse) GoString() string {
	return s.String()
}

func (s *PushWelcomeTextAndMusicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushWelcomeTextAndMusicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushWelcomeTextAndMusicResponse) GetBody() *PushWelcomeTextAndMusicResponseBody {
	return s.Body
}

func (s *PushWelcomeTextAndMusicResponse) SetHeaders(v map[string]*string) *PushWelcomeTextAndMusicResponse {
	s.Headers = v
	return s
}

func (s *PushWelcomeTextAndMusicResponse) SetStatusCode(v int32) *PushWelcomeTextAndMusicResponse {
	s.StatusCode = &v
	return s
}

func (s *PushWelcomeTextAndMusicResponse) SetBody(v *PushWelcomeTextAndMusicResponseBody) *PushWelcomeTextAndMusicResponse {
	s.Body = v
	return s
}

func (s *PushWelcomeTextAndMusicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
