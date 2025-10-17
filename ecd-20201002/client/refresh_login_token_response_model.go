// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshLoginTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshLoginTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshLoginTokenResponse
	GetStatusCode() *int32
	SetBody(v *RefreshLoginTokenResponseBody) *RefreshLoginTokenResponse
	GetBody() *RefreshLoginTokenResponseBody
}

type RefreshLoginTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshLoginTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshLoginTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshLoginTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshLoginTokenResponse) GetBody() *RefreshLoginTokenResponseBody {
	return s.Body
}

func (s *RefreshLoginTokenResponse) SetHeaders(v map[string]*string) *RefreshLoginTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshLoginTokenResponse) SetStatusCode(v int32) *RefreshLoginTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshLoginTokenResponse) SetBody(v *RefreshLoginTokenResponseBody) *RefreshLoginTokenResponse {
	s.Body = v
	return s
}

func (s *RefreshLoginTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
