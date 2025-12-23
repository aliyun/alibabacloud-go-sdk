// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateIpamResourceDiscoveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateIpamResourceDiscoveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateIpamResourceDiscoveryResponse
	GetStatusCode() *int32
	SetBody(v *DissociateIpamResourceDiscoveryResponseBody) *DissociateIpamResourceDiscoveryResponse
	GetBody() *DissociateIpamResourceDiscoveryResponseBody
}

type DissociateIpamResourceDiscoveryResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateIpamResourceDiscoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateIpamResourceDiscoveryResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateIpamResourceDiscoveryResponse) GoString() string {
	return s.String()
}

func (s *DissociateIpamResourceDiscoveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateIpamResourceDiscoveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateIpamResourceDiscoveryResponse) GetBody() *DissociateIpamResourceDiscoveryResponseBody {
	return s.Body
}

func (s *DissociateIpamResourceDiscoveryResponse) SetHeaders(v map[string]*string) *DissociateIpamResourceDiscoveryResponse {
	s.Headers = v
	return s
}

func (s *DissociateIpamResourceDiscoveryResponse) SetStatusCode(v int32) *DissociateIpamResourceDiscoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryResponse) SetBody(v *DissociateIpamResourceDiscoveryResponseBody) *DissociateIpamResourceDiscoveryResponse {
	s.Body = v
	return s
}

func (s *DissociateIpamResourceDiscoveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
