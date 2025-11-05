// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomerResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomerResponseBody) *CreateCustomerResponse
	GetBody() *CreateCustomerResponseBody
}

type CreateCustomerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomerResponse) GetBody() *CreateCustomerResponseBody {
	return s.Body
}

func (s *CreateCustomerResponse) SetHeaders(v map[string]*string) *CreateCustomerResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomerResponse) SetStatusCode(v int32) *CreateCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomerResponse) SetBody(v *CreateCustomerResponseBody) *CreateCustomerResponse {
	s.Body = v
	return s
}

func (s *CreateCustomerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
