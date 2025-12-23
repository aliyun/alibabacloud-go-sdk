// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamResourceDiscoveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpamResourceDiscoveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpamResourceDiscoveryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpamResourceDiscoveryResponseBody) *DeleteIpamResourceDiscoveryResponse
	GetBody() *DeleteIpamResourceDiscoveryResponseBody
}

type DeleteIpamResourceDiscoveryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpamResourceDiscoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpamResourceDiscoveryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamResourceDiscoveryResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpamResourceDiscoveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpamResourceDiscoveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpamResourceDiscoveryResponse) GetBody() *DeleteIpamResourceDiscoveryResponseBody {
	return s.Body
}

func (s *DeleteIpamResourceDiscoveryResponse) SetHeaders(v map[string]*string) *DeleteIpamResourceDiscoveryResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpamResourceDiscoveryResponse) SetStatusCode(v int32) *DeleteIpamResourceDiscoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryResponse) SetBody(v *DeleteIpamResourceDiscoveryResponseBody) *DeleteIpamResourceDiscoveryResponse {
	s.Body = v
	return s
}

func (s *DeleteIpamResourceDiscoveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
