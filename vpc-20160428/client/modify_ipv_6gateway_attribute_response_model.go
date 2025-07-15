// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpv6GatewayAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIpv6GatewayAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIpv6GatewayAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIpv6GatewayAttributeResponseBody) *ModifyIpv6GatewayAttributeResponse
	GetBody() *ModifyIpv6GatewayAttributeResponseBody
}

type ModifyIpv6GatewayAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIpv6GatewayAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIpv6GatewayAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpv6GatewayAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpv6GatewayAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIpv6GatewayAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIpv6GatewayAttributeResponse) GetBody() *ModifyIpv6GatewayAttributeResponseBody {
	return s.Body
}

func (s *ModifyIpv6GatewayAttributeResponse) SetHeaders(v map[string]*string) *ModifyIpv6GatewayAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpv6GatewayAttributeResponse) SetStatusCode(v int32) *ModifyIpv6GatewayAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIpv6GatewayAttributeResponse) SetBody(v *ModifyIpv6GatewayAttributeResponseBody) *ModifyIpv6GatewayAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyIpv6GatewayAttributeResponse) Validate() error {
	return dara.Validate(s)
}
