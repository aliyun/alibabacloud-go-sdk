// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWelcomeTextAndMusicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWelcomeTextAndMusicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWelcomeTextAndMusicResponse
	GetStatusCode() *int32
	SetBody(v *GetWelcomeTextAndMusicResponseBody) *GetWelcomeTextAndMusicResponse
	GetBody() *GetWelcomeTextAndMusicResponseBody
}

type GetWelcomeTextAndMusicResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWelcomeTextAndMusicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWelcomeTextAndMusicResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWelcomeTextAndMusicResponse) GoString() string {
	return s.String()
}

func (s *GetWelcomeTextAndMusicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWelcomeTextAndMusicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWelcomeTextAndMusicResponse) GetBody() *GetWelcomeTextAndMusicResponseBody {
	return s.Body
}

func (s *GetWelcomeTextAndMusicResponse) SetHeaders(v map[string]*string) *GetWelcomeTextAndMusicResponse {
	s.Headers = v
	return s
}

func (s *GetWelcomeTextAndMusicResponse) SetStatusCode(v int32) *GetWelcomeTextAndMusicResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWelcomeTextAndMusicResponse) SetBody(v *GetWelcomeTextAndMusicResponseBody) *GetWelcomeTextAndMusicResponse {
	s.Body = v
	return s
}

func (s *GetWelcomeTextAndMusicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
