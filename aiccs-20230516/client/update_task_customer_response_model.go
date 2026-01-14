// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskCustomerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskCustomerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskCustomerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskCustomerResponseBody) *UpdateTaskCustomerResponse
	GetBody() *UpdateTaskCustomerResponseBody
}

type UpdateTaskCustomerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskCustomerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskCustomerResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskCustomerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskCustomerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskCustomerResponse) GetBody() *UpdateTaskCustomerResponseBody {
	return s.Body
}

func (s *UpdateTaskCustomerResponse) SetHeaders(v map[string]*string) *UpdateTaskCustomerResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskCustomerResponse) SetStatusCode(v int32) *UpdateTaskCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskCustomerResponse) SetBody(v *UpdateTaskCustomerResponseBody) *UpdateTaskCustomerResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskCustomerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
