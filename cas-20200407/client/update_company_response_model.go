// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCompanyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCompanyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCompanyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCompanyResponseBody) *UpdateCompanyResponse
	GetBody() *UpdateCompanyResponseBody
}

type UpdateCompanyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCompanyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCompanyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompanyResponse) GoString() string {
	return s.String()
}

func (s *UpdateCompanyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCompanyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCompanyResponse) GetBody() *UpdateCompanyResponseBody {
	return s.Body
}

func (s *UpdateCompanyResponse) SetHeaders(v map[string]*string) *UpdateCompanyResponse {
	s.Headers = v
	return s
}

func (s *UpdateCompanyResponse) SetStatusCode(v int32) *UpdateCompanyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCompanyResponse) SetBody(v *UpdateCompanyResponseBody) *UpdateCompanyResponse {
	s.Body = v
	return s
}

func (s *UpdateCompanyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
