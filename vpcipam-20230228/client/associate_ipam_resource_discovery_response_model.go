// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateIpamResourceDiscoveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateIpamResourceDiscoveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateIpamResourceDiscoveryResponse
	GetStatusCode() *int32
	SetBody(v *AssociateIpamResourceDiscoveryResponseBody) *AssociateIpamResourceDiscoveryResponse
	GetBody() *AssociateIpamResourceDiscoveryResponseBody
}

type AssociateIpamResourceDiscoveryResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateIpamResourceDiscoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateIpamResourceDiscoveryResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateIpamResourceDiscoveryResponse) GoString() string {
	return s.String()
}

func (s *AssociateIpamResourceDiscoveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateIpamResourceDiscoveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateIpamResourceDiscoveryResponse) GetBody() *AssociateIpamResourceDiscoveryResponseBody {
	return s.Body
}

func (s *AssociateIpamResourceDiscoveryResponse) SetHeaders(v map[string]*string) *AssociateIpamResourceDiscoveryResponse {
	s.Headers = v
	return s
}

func (s *AssociateIpamResourceDiscoveryResponse) SetStatusCode(v int32) *AssociateIpamResourceDiscoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryResponse) SetBody(v *AssociateIpamResourceDiscoveryResponseBody) *AssociateIpamResourceDiscoveryResponse {
	s.Body = v
	return s
}

func (s *AssociateIpamResourceDiscoveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
