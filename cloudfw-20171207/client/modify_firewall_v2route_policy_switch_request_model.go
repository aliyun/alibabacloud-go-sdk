// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFirewallV2RoutePolicySwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallId(v string) *ModifyFirewallV2RoutePolicySwitchRequest
	GetFirewallId() *string
	SetLang(v string) *ModifyFirewallV2RoutePolicySwitchRequest
	GetLang() *string
	SetShouldRecover(v string) *ModifyFirewallV2RoutePolicySwitchRequest
	GetShouldRecover() *string
	SetTrFirewallRoutePolicyId(v string) *ModifyFirewallV2RoutePolicySwitchRequest
	GetTrFirewallRoutePolicyId() *string
	SetTrFirewallRoutePolicySwitchStatus(v string) *ModifyFirewallV2RoutePolicySwitchRequest
	GetTrFirewallRoutePolicySwitchStatus() *string
}

type ModifyFirewallV2RoutePolicySwitchRequest struct {
	// The instance ID of the virtual private cloud (VPC) firewalls.
	//
	// **Note**: This parameter is actually required. If it is not provided, the operation returns ErrorParameters (-360103).
	//
	// example:
	//
	// vfw-tr-5b202e7f0be64611****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to restore the traffic redirection configuration. Valid values:
	//
	// - true: Rolls back the route.
	//
	// - false: Revokes the route.
	//
	// example:
	//
	// false
	ShouldRecover *string `json:"ShouldRecover,omitempty" xml:"ShouldRecover,omitempty"`
	// The ID of the firewall route policy.
	//
	// **Note**: This parameter is actually required. If it is not provided or the policy does not exist, the operation returns ErrorTrFirewallPolicyNotFound (-361043). You can call DescribeTrFirewallV2RoutePolicyList to query valid route policy IDs.
	//
	// example:
	//
	// policy-93684cc5caa44b2e****
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
	// The status of the policy switch. Valid values:
	//
	// - open: Enabled.
	//
	// - close: Disabled.
	//
	// **Note**: This parameter is actually required. If it is not provided or the value is not within the preceding range, the operation returns ErrorParameters (-360103).
	//
	// example:
	//
	// open
	TrFirewallRoutePolicySwitchStatus *string `json:"TrFirewallRoutePolicySwitchStatus,omitempty" xml:"TrFirewallRoutePolicySwitchStatus,omitempty"`
}

func (s ModifyFirewallV2RoutePolicySwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirewallV2RoutePolicySwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) GetShouldRecover() *string {
	return s.ShouldRecover
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) GetTrFirewallRoutePolicyId() *string {
	return s.TrFirewallRoutePolicyId
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) GetTrFirewallRoutePolicySwitchStatus() *string {
	return s.TrFirewallRoutePolicySwitchStatus
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) SetFirewallId(v string) *ModifyFirewallV2RoutePolicySwitchRequest {
	s.FirewallId = &v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) SetLang(v string) *ModifyFirewallV2RoutePolicySwitchRequest {
	s.Lang = &v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) SetShouldRecover(v string) *ModifyFirewallV2RoutePolicySwitchRequest {
	s.ShouldRecover = &v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) SetTrFirewallRoutePolicyId(v string) *ModifyFirewallV2RoutePolicySwitchRequest {
	s.TrFirewallRoutePolicyId = &v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) SetTrFirewallRoutePolicySwitchStatus(v string) *ModifyFirewallV2RoutePolicySwitchRequest {
	s.TrFirewallRoutePolicySwitchStatus = &v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) Validate() error {
	return dara.Validate(s)
}
