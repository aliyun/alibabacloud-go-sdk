// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountAndAuthorityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccountAndAuthorityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccountAndAuthorityResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccountAndAuthorityResponseBody) *CreateAccountAndAuthorityResponse
	GetBody() *CreateAccountAndAuthorityResponseBody
}

type CreateAccountAndAuthorityResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccountAndAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccountAndAuthorityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountAndAuthorityResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountAndAuthorityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccountAndAuthorityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccountAndAuthorityResponse) GetBody() *CreateAccountAndAuthorityResponseBody {
	return s.Body
}

func (s *CreateAccountAndAuthorityResponse) SetHeaders(v map[string]*string) *CreateAccountAndAuthorityResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountAndAuthorityResponse) SetStatusCode(v int32) *CreateAccountAndAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccountAndAuthorityResponse) SetBody(v *CreateAccountAndAuthorityResponseBody) *CreateAccountAndAuthorityResponse {
	s.Body = v
	return s
}

func (s *CreateAccountAndAuthorityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
