// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFirewallV2RoutePolicySwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFirewallV2RoutePolicySwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFirewallV2RoutePolicySwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFirewallV2RoutePolicySwitchResponseBody) *ModifyFirewallV2RoutePolicySwitchResponse
	GetBody() *ModifyFirewallV2RoutePolicySwitchResponseBody
}

type ModifyFirewallV2RoutePolicySwitchResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFirewallV2RoutePolicySwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFirewallV2RoutePolicySwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirewallV2RoutePolicySwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyFirewallV2RoutePolicySwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFirewallV2RoutePolicySwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFirewallV2RoutePolicySwitchResponse) GetBody() *ModifyFirewallV2RoutePolicySwitchResponseBody {
	return s.Body
}

func (s *ModifyFirewallV2RoutePolicySwitchResponse) SetHeaders(v map[string]*string) *ModifyFirewallV2RoutePolicySwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchResponse) SetStatusCode(v int32) *ModifyFirewallV2RoutePolicySwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchResponse) SetBody(v *ModifyFirewallV2RoutePolicySwitchResponseBody) *ModifyFirewallV2RoutePolicySwitchResponse {
	s.Body = v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
