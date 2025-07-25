// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamResourceDiscoveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIpamResourceDiscoveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIpamResourceDiscoveryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIpamResourceDiscoveryResponseBody) *UpdateIpamResourceDiscoveryResponse
	GetBody() *UpdateIpamResourceDiscoveryResponseBody
}

type UpdateIpamResourceDiscoveryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpamResourceDiscoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpamResourceDiscoveryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamResourceDiscoveryResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpamResourceDiscoveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIpamResourceDiscoveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIpamResourceDiscoveryResponse) GetBody() *UpdateIpamResourceDiscoveryResponseBody {
	return s.Body
}

func (s *UpdateIpamResourceDiscoveryResponse) SetHeaders(v map[string]*string) *UpdateIpamResourceDiscoveryResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpamResourceDiscoveryResponse) SetStatusCode(v int32) *UpdateIpamResourceDiscoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryResponse) SetBody(v *UpdateIpamResourceDiscoveryResponseBody) *UpdateIpamResourceDiscoveryResponse {
	s.Body = v
	return s
}

func (s *UpdateIpamResourceDiscoveryResponse) Validate() error {
	return dara.Validate(s)
}
