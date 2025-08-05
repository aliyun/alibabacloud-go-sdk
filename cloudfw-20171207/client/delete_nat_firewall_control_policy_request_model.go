// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatFirewallControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuid(v string) *DeleteNatFirewallControlPolicyRequest
	GetAclUuid() *string
	SetDirection(v string) *DeleteNatFirewallControlPolicyRequest
	GetDirection() *string
	SetLang(v string) *DeleteNatFirewallControlPolicyRequest
	GetLang() *string
	SetNatGatewayId(v string) *DeleteNatFirewallControlPolicyRequest
	GetNatGatewayId() *string
}

type DeleteNatFirewallControlPolicyRequest struct {
	// The UUID of the access control policy.
	//
	// To delete an access control policy, you must provide the ID of the policy. You can call the DescribeNatFirewallControlPolicy operation to query the UUIDs of access control policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// b6c8f905-2eb6-442a-ba35-9416e9dbb2c5
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The direction of the traffic to which the access control policy applies.
	//
	// Valid values:
	//
	// 	- **out**: outbound traffic
	//
	// This parameter is required.
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
}

func (s DeleteNatFirewallControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteNatFirewallControlPolicyRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *DeleteNatFirewallControlPolicyRequest) GetDirection() *string {
	return s.Direction
}

func (s *DeleteNatFirewallControlPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteNatFirewallControlPolicyRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DeleteNatFirewallControlPolicyRequest) SetAclUuid(v string) *DeleteNatFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyRequest) SetDirection(v string) *DeleteNatFirewallControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyRequest) SetLang(v string) *DeleteNatFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyRequest) SetNatGatewayId(v string) *DeleteNatFirewallControlPolicyRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
