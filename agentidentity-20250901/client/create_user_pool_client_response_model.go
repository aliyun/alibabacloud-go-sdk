// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserPoolClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserPoolClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserPoolClientResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserPoolClientResponseBody) *CreateUserPoolClientResponse
	GetBody() *CreateUserPoolClientResponseBody
}

type CreateUserPoolClientResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserPoolClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserPoolClientResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPoolClientResponse) GoString() string {
	return s.String()
}

func (s *CreateUserPoolClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserPoolClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserPoolClientResponse) GetBody() *CreateUserPoolClientResponseBody {
	return s.Body
}

func (s *CreateUserPoolClientResponse) SetHeaders(v map[string]*string) *CreateUserPoolClientResponse {
	s.Headers = v
	return s
}

func (s *CreateUserPoolClientResponse) SetStatusCode(v int32) *CreateUserPoolClientResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserPoolClientResponse) SetBody(v *CreateUserPoolClientResponseBody) *CreateUserPoolClientResponse {
	s.Body = v
	return s
}

func (s *CreateUserPoolClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
