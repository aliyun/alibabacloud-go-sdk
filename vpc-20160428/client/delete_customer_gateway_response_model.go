// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomerGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomerGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomerGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomerGatewayResponseBody) *DeleteCustomerGatewayResponse
	GetBody() *DeleteCustomerGatewayResponseBody
}

type DeleteCustomerGatewayResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomerGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomerGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomerGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomerGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomerGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomerGatewayResponse) GetBody() *DeleteCustomerGatewayResponseBody {
	return s.Body
}

func (s *DeleteCustomerGatewayResponse) SetHeaders(v map[string]*string) *DeleteCustomerGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomerGatewayResponse) SetStatusCode(v int32) *DeleteCustomerGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomerGatewayResponse) SetBody(v *DeleteCustomerGatewayResponseBody) *DeleteCustomerGatewayResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomerGatewayResponse) Validate() error {
	return dara.Validate(s)
}
