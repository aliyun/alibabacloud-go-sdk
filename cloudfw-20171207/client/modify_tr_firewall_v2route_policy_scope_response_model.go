// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTrFirewallV2RoutePolicyScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTrFirewallV2RoutePolicyScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTrFirewallV2RoutePolicyScopeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTrFirewallV2RoutePolicyScopeResponseBody) *ModifyTrFirewallV2RoutePolicyScopeResponse
	GetBody() *ModifyTrFirewallV2RoutePolicyScopeResponseBody
}

type ModifyTrFirewallV2RoutePolicyScopeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTrFirewallV2RoutePolicyScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTrFirewallV2RoutePolicyScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrFirewallV2RoutePolicyScopeResponse) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponse) GetBody() *ModifyTrFirewallV2RoutePolicyScopeResponseBody {
	return s.Body
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponse) SetHeaders(v map[string]*string) *ModifyTrFirewallV2RoutePolicyScopeResponse {
	s.Headers = v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponse) SetStatusCode(v int32) *ModifyTrFirewallV2RoutePolicyScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponse) SetBody(v *ModifyTrFirewallV2RoutePolicyScopeResponseBody) *ModifyTrFirewallV2RoutePolicyScopeResponse {
	s.Body = v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponse) Validate() error {
	return dara.Validate(s)
}
