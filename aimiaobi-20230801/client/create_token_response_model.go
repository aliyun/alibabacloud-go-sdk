// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateTokenResponseBody) *CreateTokenResponse
	GetBody() *CreateTokenResponseBody
}

type CreateTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTokenResponse) GetBody() *CreateTokenResponseBody {
	return s.Body
}

func (s *CreateTokenResponse) SetHeaders(v map[string]*string) *CreateTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateTokenResponse) SetStatusCode(v int32) *CreateTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTokenResponse) SetBody(v *CreateTokenResponseBody) *CreateTokenResponse {
	s.Body = v
	return s
}

func (s *CreateTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
