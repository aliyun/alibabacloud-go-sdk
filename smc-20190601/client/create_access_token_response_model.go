// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccessTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccessTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccessTokenResponseBody) *CreateAccessTokenResponse
	GetBody() *CreateAccessTokenResponseBody
}

type CreateAccessTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccessTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccessTokenResponse) GetBody() *CreateAccessTokenResponseBody {
	return s.Body
}

func (s *CreateAccessTokenResponse) SetHeaders(v map[string]*string) *CreateAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessTokenResponse) SetStatusCode(v int32) *CreateAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessTokenResponse) SetBody(v *CreateAccessTokenResponseBody) *CreateAccessTokenResponse {
	s.Body = v
	return s
}

func (s *CreateAccessTokenResponse) Validate() error {
	return dara.Validate(s)
}
