// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicIpAddressPoolAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePublicIpAddressPoolAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePublicIpAddressPoolAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePublicIpAddressPoolAttributeResponseBody) *UpdatePublicIpAddressPoolAttributeResponse
	GetBody() *UpdatePublicIpAddressPoolAttributeResponseBody
}

type UpdatePublicIpAddressPoolAttributeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePublicIpAddressPoolAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePublicIpAddressPoolAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicIpAddressPoolAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdatePublicIpAddressPoolAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePublicIpAddressPoolAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePublicIpAddressPoolAttributeResponse) GetBody() *UpdatePublicIpAddressPoolAttributeResponseBody {
	return s.Body
}

func (s *UpdatePublicIpAddressPoolAttributeResponse) SetHeaders(v map[string]*string) *UpdatePublicIpAddressPoolAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdatePublicIpAddressPoolAttributeResponse) SetStatusCode(v int32) *UpdatePublicIpAddressPoolAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePublicIpAddressPoolAttributeResponse) SetBody(v *UpdatePublicIpAddressPoolAttributeResponseBody) *UpdatePublicIpAddressPoolAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdatePublicIpAddressPoolAttributeResponse) Validate() error {
	return dara.Validate(s)
}
