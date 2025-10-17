// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyForStreamAccessTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyForStreamAccessTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyForStreamAccessTokenResponse
	GetStatusCode() *int32
	SetBody(v *ApplyForStreamAccessTokenResponseBody) *ApplyForStreamAccessTokenResponse
	GetBody() *ApplyForStreamAccessTokenResponseBody
}

type ApplyForStreamAccessTokenResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyForStreamAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyForStreamAccessTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyForStreamAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *ApplyForStreamAccessTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyForStreamAccessTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyForStreamAccessTokenResponse) GetBody() *ApplyForStreamAccessTokenResponseBody {
	return s.Body
}

func (s *ApplyForStreamAccessTokenResponse) SetHeaders(v map[string]*string) *ApplyForStreamAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *ApplyForStreamAccessTokenResponse) SetStatusCode(v int32) *ApplyForStreamAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyForStreamAccessTokenResponse) SetBody(v *ApplyForStreamAccessTokenResponseBody) *ApplyForStreamAccessTokenResponse {
	s.Body = v
	return s
}

func (s *ApplyForStreamAccessTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
