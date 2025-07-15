// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpv4GatewayAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIpv4GatewayAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIpv4GatewayAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIpv4GatewayAttributeResponseBody) *UpdateIpv4GatewayAttributeResponse
	GetBody() *UpdateIpv4GatewayAttributeResponseBody
}

type UpdateIpv4GatewayAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpv4GatewayAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpv4GatewayAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpv4GatewayAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpv4GatewayAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIpv4GatewayAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIpv4GatewayAttributeResponse) GetBody() *UpdateIpv4GatewayAttributeResponseBody {
	return s.Body
}

func (s *UpdateIpv4GatewayAttributeResponse) SetHeaders(v map[string]*string) *UpdateIpv4GatewayAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpv4GatewayAttributeResponse) SetStatusCode(v int32) *UpdateIpv4GatewayAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpv4GatewayAttributeResponse) SetBody(v *UpdateIpv4GatewayAttributeResponseBody) *UpdateIpv4GatewayAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateIpv4GatewayAttributeResponse) Validate() error {
	return dara.Validate(s)
}
