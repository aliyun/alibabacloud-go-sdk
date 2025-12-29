// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetWelcomeTextAndMusicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetWelcomeTextAndMusicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetWelcomeTextAndMusicResponse
	GetStatusCode() *int32
	SetBody(v *ResetWelcomeTextAndMusicResponseBody) *ResetWelcomeTextAndMusicResponse
	GetBody() *ResetWelcomeTextAndMusicResponseBody
}

type ResetWelcomeTextAndMusicResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetWelcomeTextAndMusicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetWelcomeTextAndMusicResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetWelcomeTextAndMusicResponse) GoString() string {
	return s.String()
}

func (s *ResetWelcomeTextAndMusicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetWelcomeTextAndMusicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetWelcomeTextAndMusicResponse) GetBody() *ResetWelcomeTextAndMusicResponseBody {
	return s.Body
}

func (s *ResetWelcomeTextAndMusicResponse) SetHeaders(v map[string]*string) *ResetWelcomeTextAndMusicResponse {
	s.Headers = v
	return s
}

func (s *ResetWelcomeTextAndMusicResponse) SetStatusCode(v int32) *ResetWelcomeTextAndMusicResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetWelcomeTextAndMusicResponse) SetBody(v *ResetWelcomeTextAndMusicResponseBody) *ResetWelcomeTextAndMusicResponse {
	s.Body = v
	return s
}

func (s *ResetWelcomeTextAndMusicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
