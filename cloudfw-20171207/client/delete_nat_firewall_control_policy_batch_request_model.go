// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatFirewallControlPolicyBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuidList(v []*string) *DeleteNatFirewallControlPolicyBatchRequest
	GetAclUuidList() []*string
	SetDirection(v string) *DeleteNatFirewallControlPolicyBatchRequest
	GetDirection() *string
	SetLang(v string) *DeleteNatFirewallControlPolicyBatchRequest
	GetLang() *string
	SetNatGatewayId(v string) *DeleteNatFirewallControlPolicyBatchRequest
	GetNatGatewayId() *string
}

type DeleteNatFirewallControlPolicyBatchRequest struct {
	// The UUIDs of access control policies.
	//
	// This parameter is required.
	AclUuidList []*string `json:"AclUuidList,omitempty" xml:"AclUuidList,omitempty" type:"Repeated"`
	// The direction of the traffic to which the access control policy applies. Valid values:
	//
	// 	- **out**: outbound traffic
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
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
	// ngw-uf6j0426ap74vd6vrb676
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
}

func (s DeleteNatFirewallControlPolicyBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatFirewallControlPolicyBatchRequest) GoString() string {
	return s.String()
}

func (s *DeleteNatFirewallControlPolicyBatchRequest) GetAclUuidList() []*string {
	return s.AclUuidList
}

func (s *DeleteNatFirewallControlPolicyBatchRequest) GetDirection() *string {
	return s.Direction
}

func (s *DeleteNatFirewallControlPolicyBatchRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteNatFirewallControlPolicyBatchRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DeleteNatFirewallControlPolicyBatchRequest) SetAclUuidList(v []*string) *DeleteNatFirewallControlPolicyBatchRequest {
	s.AclUuidList = v
	return s
}

func (s *DeleteNatFirewallControlPolicyBatchRequest) SetDirection(v string) *DeleteNatFirewallControlPolicyBatchRequest {
	s.Direction = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyBatchRequest) SetLang(v string) *DeleteNatFirewallControlPolicyBatchRequest {
	s.Lang = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyBatchRequest) SetNatGatewayId(v string) *DeleteNatFirewallControlPolicyBatchRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyBatchRequest) Validate() error {
	return dara.Validate(s)
}
