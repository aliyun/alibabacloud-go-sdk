// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamDiscoveredIpAddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpamDiscoveredIpAddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpamDiscoveredIpAddressesResponse
	GetStatusCode() *int32
	SetBody(v *ListIpamDiscoveredIpAddressesResponseBody) *ListIpamDiscoveredIpAddressesResponse
	GetBody() *ListIpamDiscoveredIpAddressesResponseBody
}

type ListIpamDiscoveredIpAddressesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamDiscoveredIpAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamDiscoveredIpAddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpamDiscoveredIpAddressesResponse) GoString() string {
	return s.String()
}

func (s *ListIpamDiscoveredIpAddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpamDiscoveredIpAddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpamDiscoveredIpAddressesResponse) GetBody() *ListIpamDiscoveredIpAddressesResponseBody {
	return s.Body
}

func (s *ListIpamDiscoveredIpAddressesResponse) SetHeaders(v map[string]*string) *ListIpamDiscoveredIpAddressesResponse {
	s.Headers = v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponse) SetStatusCode(v int32) *ListIpamDiscoveredIpAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponse) SetBody(v *ListIpamDiscoveredIpAddressesResponseBody) *ListIpamDiscoveredIpAddressesResponse {
	s.Body = v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
