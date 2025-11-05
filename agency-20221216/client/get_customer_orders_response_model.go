// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomerOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomerOrdersResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomerOrdersResponseBody) *GetCustomerOrdersResponse
	GetBody() *GetCustomerOrdersResponseBody
}

type GetCustomerOrdersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerOrdersResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomerOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomerOrdersResponse) GetBody() *GetCustomerOrdersResponseBody {
	return s.Body
}

func (s *GetCustomerOrdersResponse) SetHeaders(v map[string]*string) *GetCustomerOrdersResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerOrdersResponse) SetStatusCode(v int32) *GetCustomerOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerOrdersResponse) SetBody(v *GetCustomerOrdersResponseBody) *GetCustomerOrdersResponse {
	s.Body = v
	return s
}

func (s *GetCustomerOrdersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
