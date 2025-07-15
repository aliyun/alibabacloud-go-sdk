// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomerGatewayAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCustomerGatewayAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCustomerGatewayAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCustomerGatewayAttributeResponseBody) *ModifyCustomerGatewayAttributeResponse
	GetBody() *ModifyCustomerGatewayAttributeResponseBody
}

type ModifyCustomerGatewayAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCustomerGatewayAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCustomerGatewayAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomerGatewayAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyCustomerGatewayAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCustomerGatewayAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCustomerGatewayAttributeResponse) GetBody() *ModifyCustomerGatewayAttributeResponseBody {
	return s.Body
}

func (s *ModifyCustomerGatewayAttributeResponse) SetHeaders(v map[string]*string) *ModifyCustomerGatewayAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyCustomerGatewayAttributeResponse) SetStatusCode(v int32) *ModifyCustomerGatewayAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeResponse) SetBody(v *ModifyCustomerGatewayAttributeResponseBody) *ModifyCustomerGatewayAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyCustomerGatewayAttributeResponse) Validate() error {
	return dara.Validate(s)
}
