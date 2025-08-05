// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTrFirewallV2RoutePolicyScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTrFirewallV2RoutePolicyScopeResponseBody
	GetRequestId() *string
	SetTrFirewallRoutePolicyId(v string) *ModifyTrFirewallV2RoutePolicyScopeResponseBody
	GetTrFirewallRoutePolicyId() *string
}

type ModifyTrFirewallV2RoutePolicyScopeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E7D4D635-0C70-5CEB-A609-851E94D51FBB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the routing policy.
	//
	// example:
	//
	// policy-4d724d0139df48f18091
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s ModifyTrFirewallV2RoutePolicyScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrFirewallV2RoutePolicyScopeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponseBody) GetTrFirewallRoutePolicyId() *string {
	return s.TrFirewallRoutePolicyId
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponseBody) SetRequestId(v string) *ModifyTrFirewallV2RoutePolicyScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponseBody) SetTrFirewallRoutePolicyId(v string) *ModifyTrFirewallV2RoutePolicyScopeResponseBody {
	s.TrFirewallRoutePolicyId = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
