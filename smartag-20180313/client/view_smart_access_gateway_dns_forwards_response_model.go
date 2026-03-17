// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayDnsForwardsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ViewSmartAccessGatewayDnsForwardsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ViewSmartAccessGatewayDnsForwardsResponse
	GetStatusCode() *int32
	SetBody(v *ViewSmartAccessGatewayDnsForwardsResponseBody) *ViewSmartAccessGatewayDnsForwardsResponse
	GetBody() *ViewSmartAccessGatewayDnsForwardsResponseBody
}

type ViewSmartAccessGatewayDnsForwardsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ViewSmartAccessGatewayDnsForwardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ViewSmartAccessGatewayDnsForwardsResponse) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayDnsForwardsResponse) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayDnsForwardsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ViewSmartAccessGatewayDnsForwardsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ViewSmartAccessGatewayDnsForwardsResponse) GetBody() *ViewSmartAccessGatewayDnsForwardsResponseBody {
	return s.Body
}

func (s *ViewSmartAccessGatewayDnsForwardsResponse) SetHeaders(v map[string]*string) *ViewSmartAccessGatewayDnsForwardsResponse {
	s.Headers = v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponse) SetStatusCode(v int32) *ViewSmartAccessGatewayDnsForwardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponse) SetBody(v *ViewSmartAccessGatewayDnsForwardsResponseBody) *ViewSmartAccessGatewayDnsForwardsResponse {
	s.Body = v
	return s
}

func (s *ViewSmartAccessGatewayDnsForwardsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
