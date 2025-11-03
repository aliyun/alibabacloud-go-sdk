// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomerGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomerGatewayResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomerGatewayResponseBody) *CreateCustomerGatewayResponse
	GetBody() *CreateCustomerGatewayResponseBody
}

type CreateCustomerGatewayResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomerGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomerGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomerGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomerGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomerGatewayResponse) GetBody() *CreateCustomerGatewayResponseBody {
	return s.Body
}

func (s *CreateCustomerGatewayResponse) SetHeaders(v map[string]*string) *CreateCustomerGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomerGatewayResponse) SetStatusCode(v int32) *CreateCustomerGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomerGatewayResponse) SetBody(v *CreateCustomerGatewayResponseBody) *CreateCustomerGatewayResponse {
	s.Body = v
	return s
}

func (s *CreateCustomerGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
