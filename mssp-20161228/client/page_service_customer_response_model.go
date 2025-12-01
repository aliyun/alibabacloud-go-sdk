// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageServiceCustomerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PageServiceCustomerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PageServiceCustomerResponse
	GetStatusCode() *int32
	SetBody(v *PageServiceCustomerResponseBody) *PageServiceCustomerResponse
	GetBody() *PageServiceCustomerResponseBody
}

type PageServiceCustomerResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageServiceCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageServiceCustomerResponse) String() string {
	return dara.Prettify(s)
}

func (s PageServiceCustomerResponse) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PageServiceCustomerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PageServiceCustomerResponse) GetBody() *PageServiceCustomerResponseBody {
	return s.Body
}

func (s *PageServiceCustomerResponse) SetHeaders(v map[string]*string) *PageServiceCustomerResponse {
	s.Headers = v
	return s
}

func (s *PageServiceCustomerResponse) SetStatusCode(v int32) *PageServiceCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *PageServiceCustomerResponse) SetBody(v *PageServiceCustomerResponseBody) *PageServiceCustomerResponse {
	s.Body = v
	return s
}

func (s *PageServiceCustomerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
