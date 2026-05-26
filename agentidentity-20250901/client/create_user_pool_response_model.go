// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserPoolResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserPoolResponseBody) *CreateUserPoolResponse
	GetBody() *CreateUserPoolResponseBody
}

type CreateUserPoolResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPoolResponse) GoString() string {
	return s.String()
}

func (s *CreateUserPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserPoolResponse) GetBody() *CreateUserPoolResponseBody {
	return s.Body
}

func (s *CreateUserPoolResponse) SetHeaders(v map[string]*string) *CreateUserPoolResponse {
	s.Headers = v
	return s
}

func (s *CreateUserPoolResponse) SetStatusCode(v int32) *CreateUserPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserPoolResponse) SetBody(v *CreateUserPoolResponseBody) *CreateUserPoolResponse {
	s.Body = v
	return s
}

func (s *CreateUserPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
