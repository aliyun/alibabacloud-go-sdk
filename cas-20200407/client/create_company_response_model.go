// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCompanyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCompanyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCompanyResponse
	GetStatusCode() *int32
	SetBody(v *CreateCompanyResponseBody) *CreateCompanyResponse
	GetBody() *CreateCompanyResponseBody
}

type CreateCompanyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCompanyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCompanyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCompanyResponse) GoString() string {
	return s.String()
}

func (s *CreateCompanyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCompanyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCompanyResponse) GetBody() *CreateCompanyResponseBody {
	return s.Body
}

func (s *CreateCompanyResponse) SetHeaders(v map[string]*string) *CreateCompanyResponse {
	s.Headers = v
	return s
}

func (s *CreateCompanyResponse) SetStatusCode(v int32) *CreateCompanyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCompanyResponse) SetBody(v *CreateCompanyResponseBody) *CreateCompanyResponse {
	s.Body = v
	return s
}

func (s *CreateCompanyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
