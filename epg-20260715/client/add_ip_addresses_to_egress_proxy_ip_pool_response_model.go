// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpAddressesToEgressProxyIpPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddIpAddressesToEgressProxyIpPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddIpAddressesToEgressProxyIpPoolResponse
	GetStatusCode() *int32
	SetBody(v *AddIpAddressesToEgressProxyIpPoolResponseBody) *AddIpAddressesToEgressProxyIpPoolResponse
	GetBody() *AddIpAddressesToEgressProxyIpPoolResponseBody
}

type AddIpAddressesToEgressProxyIpPoolResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddIpAddressesToEgressProxyIpPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddIpAddressesToEgressProxyIpPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s AddIpAddressesToEgressProxyIpPoolResponse) GoString() string {
	return s.String()
}

func (s *AddIpAddressesToEgressProxyIpPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddIpAddressesToEgressProxyIpPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddIpAddressesToEgressProxyIpPoolResponse) GetBody() *AddIpAddressesToEgressProxyIpPoolResponseBody {
	return s.Body
}

func (s *AddIpAddressesToEgressProxyIpPoolResponse) SetHeaders(v map[string]*string) *AddIpAddressesToEgressProxyIpPoolResponse {
	s.Headers = v
	return s
}

func (s *AddIpAddressesToEgressProxyIpPoolResponse) SetStatusCode(v int32) *AddIpAddressesToEgressProxyIpPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIpAddressesToEgressProxyIpPoolResponse) SetBody(v *AddIpAddressesToEgressProxyIpPoolResponseBody) *AddIpAddressesToEgressProxyIpPoolResponse {
	s.Body = v
	return s
}

func (s *AddIpAddressesToEgressProxyIpPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
