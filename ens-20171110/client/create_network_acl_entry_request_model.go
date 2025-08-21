// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkAclEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *CreateNetworkAclEntryRequest
	GetCidrBlock() *string
	SetDescription(v string) *CreateNetworkAclEntryRequest
	GetDescription() *string
	SetDestinationCidrBlock(v string) *CreateNetworkAclEntryRequest
	GetDestinationCidrBlock() *string
	SetDirection(v string) *CreateNetworkAclEntryRequest
	GetDirection() *string
	SetNetworkAclEntryName(v string) *CreateNetworkAclEntryRequest
	GetNetworkAclEntryName() *string
	SetNetworkAclId(v string) *CreateNetworkAclEntryRequest
	GetNetworkAclId() *string
	SetPolicy(v string) *CreateNetworkAclEntryRequest
	GetPolicy() *string
	SetPortRange(v string) *CreateNetworkAclEntryRequest
	GetPortRange() *string
	SetPriority(v int32) *CreateNetworkAclEntryRequest
	GetPriority() *int32
	SetProtocol(v string) *CreateNetworkAclEntryRequest
	GetProtocol() *string
}

type CreateNetworkAclEntryRequest struct {
	// The source CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description of the network ACL.
	//
	// The description must be 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// This is my NetworkAcl.
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The direction in which the rule is applied. Valid values:
	//
	// 	- **ingress**
	//
	// 	- **egress**
	//
	// This parameter is required.
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The name of the rule.
	//
	// The name must be 1 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// acl-1
	NetworkAclEntryName *string `json:"NetworkAclEntryName,omitempty" xml:"NetworkAclEntryName,omitempty"`
	// The ID of the network ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-bp1lhl0taikrbgnh****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// The action that is performed on network traffic that matches the rule. Valid values:
	//
	// 	- **accept**: allows network traffic.
	//
	// 	- **drop**: blocks network traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The port range.
	//
	// 	- If you set **Protocol*	- to **all*	- or **icmp**, set this parameter to -1/-1, which specifies all ports.
	//
	// 	- If you set **Protocol*	- to **tcp*	- or **udp**, the port can be **1 to 65535**. You can set this parameter to **1/200*	- or **80/80**, which specifies ports 1 to 200 or port 80.
	//
	// This parameter is required.
	//
	// example:
	//
	// -1/-1
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the rule. Valid values: **1 to 100**. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **icmp**: ICMP
	//
	// 	- **tcp**: TCP
	//
	// 	- **udp**: UDP
	//
	// 	- **all**: all protocols
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateNetworkAclEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclEntryRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateNetworkAclEntryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkAclEntryRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *CreateNetworkAclEntryRequest) GetDirection() *string {
	return s.Direction
}

func (s *CreateNetworkAclEntryRequest) GetNetworkAclEntryName() *string {
	return s.NetworkAclEntryName
}

func (s *CreateNetworkAclEntryRequest) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *CreateNetworkAclEntryRequest) GetPolicy() *string {
	return s.Policy
}

func (s *CreateNetworkAclEntryRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *CreateNetworkAclEntryRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateNetworkAclEntryRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateNetworkAclEntryRequest) SetCidrBlock(v string) *CreateNetworkAclEntryRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateNetworkAclEntryRequest) SetDescription(v string) *CreateNetworkAclEntryRequest {
	s.Description = &v
	return s
}

func (s *CreateNetworkAclEntryRequest) SetDestinationCidrBlock(v string) *CreateNetworkAclEntryRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateNetworkAclEntryRequest) SetDirection(v string) *CreateNetworkAclEntryRequest {
	s.Direction = &v
	return s
}

func (s *CreateNetworkAclEntryRequest) SetNetworkAclEntryName(v string) *CreateNetworkAclEntryRequest {
	s.NetworkAclEntryName = &v
	return s
}

func (s *CreateNetworkAclEntryRequest) SetNetworkAclId(v string) *CreateNetworkAclEntryRequest {
	s.NetworkAclId = &v
	return s
}

func (s *CreateNetworkAclEntryRequest) SetPolicy(v string) *CreateNetworkAclEntryRequest {
	s.Policy = &v
	return s
}

func (s *CreateNetworkAclEntryRequest) SetPortRange(v string) *CreateNetworkAclEntryRequest {
	s.PortRange = &v
	return s
}

func (s *CreateNetworkAclEntryRequest) SetPriority(v int32) *CreateNetworkAclEntryRequest {
	s.Priority = &v
	return s
}

func (s *CreateNetworkAclEntryRequest) SetProtocol(v string) *CreateNetworkAclEntryRequest {
	s.Protocol = &v
	return s
}

func (s *CreateNetworkAclEntryRequest) Validate() error {
	return dara.Validate(s)
}
