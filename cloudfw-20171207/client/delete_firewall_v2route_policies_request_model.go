// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallV2RoutePoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallId(v string) *DeleteFirewallV2RoutePoliciesRequest
	GetFirewallId() *string
	SetLang(v string) *DeleteFirewallV2RoutePoliciesRequest
	GetLang() *string
	SetTrFirewallRoutePolicyId(v string) *DeleteFirewallV2RoutePoliciesRequest
	GetTrFirewallRoutePolicyId() *string
}

type DeleteFirewallV2RoutePoliciesRequest struct {
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-tr-d5ba592ac6c84aff****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the routing policy.
	//
	// example:
	//
	// policy-2d06d3568fd74d60****
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s DeleteFirewallV2RoutePoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallV2RoutePoliciesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFirewallV2RoutePoliciesRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DeleteFirewallV2RoutePoliciesRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteFirewallV2RoutePoliciesRequest) GetTrFirewallRoutePolicyId() *string {
	return s.TrFirewallRoutePolicyId
}

func (s *DeleteFirewallV2RoutePoliciesRequest) SetFirewallId(v string) *DeleteFirewallV2RoutePoliciesRequest {
	s.FirewallId = &v
	return s
}

func (s *DeleteFirewallV2RoutePoliciesRequest) SetLang(v string) *DeleteFirewallV2RoutePoliciesRequest {
	s.Lang = &v
	return s
}

func (s *DeleteFirewallV2RoutePoliciesRequest) SetTrFirewallRoutePolicyId(v string) *DeleteFirewallV2RoutePoliciesRequest {
	s.TrFirewallRoutePolicyId = &v
	return s
}

func (s *DeleteFirewallV2RoutePoliciesRequest) Validate() error {
	return dara.Validate(s)
}
