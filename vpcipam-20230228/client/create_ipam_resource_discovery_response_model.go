// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamResourceDiscoveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIpamResourceDiscoveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIpamResourceDiscoveryResponse
	GetStatusCode() *int32
	SetBody(v *CreateIpamResourceDiscoveryResponseBody) *CreateIpamResourceDiscoveryResponse
	GetBody() *CreateIpamResourceDiscoveryResponseBody
}

type CreateIpamResourceDiscoveryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpamResourceDiscoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpamResourceDiscoveryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamResourceDiscoveryResponse) GoString() string {
	return s.String()
}

func (s *CreateIpamResourceDiscoveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIpamResourceDiscoveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIpamResourceDiscoveryResponse) GetBody() *CreateIpamResourceDiscoveryResponseBody {
	return s.Body
}

func (s *CreateIpamResourceDiscoveryResponse) SetHeaders(v map[string]*string) *CreateIpamResourceDiscoveryResponse {
	s.Headers = v
	return s
}

func (s *CreateIpamResourceDiscoveryResponse) SetStatusCode(v int32) *CreateIpamResourceDiscoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpamResourceDiscoveryResponse) SetBody(v *CreateIpamResourceDiscoveryResponseBody) *CreateIpamResourceDiscoveryResponse {
	s.Body = v
	return s
}

func (s *CreateIpamResourceDiscoveryResponse) Validate() error {
	return dara.Validate(s)
}
