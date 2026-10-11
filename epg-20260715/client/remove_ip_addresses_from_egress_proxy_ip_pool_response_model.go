// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIpAddressesFromEgressProxyIpPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveIpAddressesFromEgressProxyIpPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveIpAddressesFromEgressProxyIpPoolResponse
	GetStatusCode() *int32
	SetBody(v *RemoveIpAddressesFromEgressProxyIpPoolResponseBody) *RemoveIpAddressesFromEgressProxyIpPoolResponse
	GetBody() *RemoveIpAddressesFromEgressProxyIpPoolResponseBody
}

type RemoveIpAddressesFromEgressProxyIpPoolResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveIpAddressesFromEgressProxyIpPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveIpAddressesFromEgressProxyIpPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpAddressesFromEgressProxyIpPoolResponse) GoString() string {
	return s.String()
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolResponse) GetBody() *RemoveIpAddressesFromEgressProxyIpPoolResponseBody {
	return s.Body
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolResponse) SetHeaders(v map[string]*string) *RemoveIpAddressesFromEgressProxyIpPoolResponse {
	s.Headers = v
	return s
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolResponse) SetStatusCode(v int32) *RemoveIpAddressesFromEgressProxyIpPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolResponse) SetBody(v *RemoveIpAddressesFromEgressProxyIpPoolResponseBody) *RemoveIpAddressesFromEgressProxyIpPoolResponse {
	s.Body = v
	return s
}

func (s *RemoveIpAddressesFromEgressProxyIpPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
