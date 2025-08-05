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
	// The UUID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 66961eea-e659-4225-84c9-9b6da76ec401
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The direction of the traffic to which the access control policy applies.
	//
	// 	- Set the value to **out**.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-xxxxxx
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The new priority of the IPv4 access control policy. You must specify a numeric value for this parameter. The value 1 indicates the highest priority. A larger value indicates a lower priority.
	//
	// > Make sure that the value of this parameter is within the priority range of existing IPv4 access control policies. Otherwise, an error occurs when you call this operation.
	//
	// Before you call this operation, we recommend that you call the DescribeNatFirewallPolicyPriorUsed operation to query the priority range of the IPv4 access control policies in the specified traffic direction.
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
