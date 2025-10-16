// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallV2RoutePoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFirewallV2RoutePoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFirewallV2RoutePoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFirewallV2RoutePoliciesResponseBody) *DeleteFirewallV2RoutePoliciesResponse
	GetBody() *DeleteFirewallV2RoutePoliciesResponseBody
}

type DeleteFirewallV2RoutePoliciesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFirewallV2RoutePoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFirewallV2RoutePoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallV2RoutePoliciesResponse) GoString() string {
	return s.String()
}

func (s *DeleteFirewallV2RoutePoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFirewallV2RoutePoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFirewallV2RoutePoliciesResponse) GetBody() *DeleteFirewallV2RoutePoliciesResponseBody {
	return s.Body
}

func (s *DeleteFirewallV2RoutePoliciesResponse) SetHeaders(v map[string]*string) *DeleteFirewallV2RoutePoliciesResponse {
	s.Headers = v
	return s
}

func (s *DeleteFirewallV2RoutePoliciesResponse) SetStatusCode(v int32) *DeleteFirewallV2RoutePoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFirewallV2RoutePoliciesResponse) SetBody(v *DeleteFirewallV2RoutePoliciesResponseBody) *DeleteFirewallV2RoutePoliciesResponse {
	s.Body = v
	return s
}

func (s *DeleteFirewallV2RoutePoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
