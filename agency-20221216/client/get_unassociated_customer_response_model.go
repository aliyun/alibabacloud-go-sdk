// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnassociatedCustomerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUnassociatedCustomerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUnassociatedCustomerResponse
	GetStatusCode() *int32
	SetBody(v *GetUnassociatedCustomerResponseBody) *GetUnassociatedCustomerResponse
	GetBody() *GetUnassociatedCustomerResponseBody
}

type GetUnassociatedCustomerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUnassociatedCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUnassociatedCustomerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUnassociatedCustomerResponse) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUnassociatedCustomerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUnassociatedCustomerResponse) GetBody() *GetUnassociatedCustomerResponseBody {
	return s.Body
}

func (s *GetUnassociatedCustomerResponse) SetHeaders(v map[string]*string) *GetUnassociatedCustomerResponse {
	s.Headers = v
	return s
}

func (s *GetUnassociatedCustomerResponse) SetStatusCode(v int32) *GetUnassociatedCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUnassociatedCustomerResponse) SetBody(v *GetUnassociatedCustomerResponseBody) *GetUnassociatedCustomerResponse {
	s.Body = v
	return s
}

func (s *GetUnassociatedCustomerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
