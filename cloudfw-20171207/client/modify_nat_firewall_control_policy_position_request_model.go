// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatFirewallControlPolicyPositionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuid(v string) *ModifyNatFirewallControlPolicyPositionRequest
	GetAclUuid() *string
	SetDirection(v string) *ModifyNatFirewallControlPolicyPositionRequest
	GetDirection() *string
	SetLang(v string) *ModifyNatFirewallControlPolicyPositionRequest
	GetLang() *string
	SetNatGatewayId(v string) *ModifyNatFirewallControlPolicyPositionRequest
	GetNatGatewayId() *string
	SetNewOrder(v int32) *ModifyNatFirewallControlPolicyPositionRequest
	GetNewOrder() *int32
}

type ModifyNatFirewallControlPolicyPositionRequest struct {
	// The unique ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 66461eea-e659-4225-84c9-*****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The traffic direction of the access control policy.
	//
	// Valid values:
	//
	// - **out**: outbound traffic.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The NAT gateway ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-xxxxxx
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The policy priority for the NAT firewall IPv4 access control policy. A value of 1 indicates the highest priority. A larger value indicates a lower priority.
	//
	// > The policy priority value must be within the range of existing NAT firewall IPv4 policy priorities. Otherwise, an error occurs when you call this operation.
	//
	// Before you call this operation, call DescribeNatFirewallPolicyPriorUsed to query the priority range of IPv4 policies for the specified traffic direction of the NAT firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	NewOrder *int32 `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
}

func (s ModifyNatFirewallControlPolicyPositionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatFirewallControlPolicyPositionRequest) GoString() string {
	return s.String()
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) GetNewOrder() *int32 {
	return s.NewOrder
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) SetAclUuid(v string) *ModifyNatFirewallControlPolicyPositionRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) SetDirection(v string) *ModifyNatFirewallControlPolicyPositionRequest {
	s.Direction = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) SetLang(v string) *ModifyNatFirewallControlPolicyPositionRequest {
	s.Lang = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) SetNatGatewayId(v string) *ModifyNatFirewallControlPolicyPositionRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) SetNewOrder(v int32) *ModifyNatFirewallControlPolicyPositionRequest {
	s.NewOrder = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) Validate() error {
	return dara.Validate(s)
}
