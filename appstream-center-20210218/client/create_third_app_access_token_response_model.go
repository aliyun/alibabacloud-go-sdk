// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateThirdAppAccessTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateThirdAppAccessTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateThirdAppAccessTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateThirdAppAccessTokenResponseBody) *CreateThirdAppAccessTokenResponse
	GetBody() *CreateThirdAppAccessTokenResponseBody
}

type CreateThirdAppAccessTokenResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateThirdAppAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateThirdAppAccessTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateThirdAppAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateThirdAppAccessTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateThirdAppAccessTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateThirdAppAccessTokenResponse) GetBody() *CreateThirdAppAccessTokenResponseBody {
	return s.Body
}

func (s *CreateThirdAppAccessTokenResponse) SetHeaders(v map[string]*string) *CreateThirdAppAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateThirdAppAccessTokenResponse) SetStatusCode(v int32) *CreateThirdAppAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateThirdAppAccessTokenResponse) SetBody(v *CreateThirdAppAccessTokenResponseBody) *CreateThirdAppAccessTokenResponse {
	s.Body = v
	return s
}

func (s *CreateThirdAppAccessTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
