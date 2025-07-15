// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpv6AddressAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIpv6AddressAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIpv6AddressAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIpv6AddressAttributeResponseBody) *ModifyIpv6AddressAttributeResponse
	GetBody() *ModifyIpv6AddressAttributeResponseBody
}

type ModifyIpv6AddressAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIpv6AddressAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIpv6AddressAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpv6AddressAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpv6AddressAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIpv6AddressAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIpv6AddressAttributeResponse) GetBody() *ModifyIpv6AddressAttributeResponseBody {
	return s.Body
}

func (s *ModifyIpv6AddressAttributeResponse) SetHeaders(v map[string]*string) *ModifyIpv6AddressAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpv6AddressAttributeResponse) SetStatusCode(v int32) *ModifyIpv6AddressAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIpv6AddressAttributeResponse) SetBody(v *ModifyIpv6AddressAttributeResponseBody) *ModifyIpv6AddressAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyIpv6AddressAttributeResponse) Validate() error {
	return dara.Validate(s)
}
