// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddACLRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *AddACLRuleResponseBody
	GetAclId() *string
	SetAcrId(v string) *AddACLRuleResponseBody
	GetAcrId() *string
	SetDescription(v string) *AddACLRuleResponseBody
	GetDescription() *string
	SetDestCidr(v string) *AddACLRuleResponseBody
	GetDestCidr() *string
	SetDestPortRange(v string) *AddACLRuleResponseBody
	GetDestPortRange() *string
	SetDirection(v string) *AddACLRuleResponseBody
	GetDirection() *string
	SetDpiGroupIds(v *AddACLRuleResponseBodyDpiGroupIds) *AddACLRuleResponseBody
	GetDpiGroupIds() *AddACLRuleResponseBodyDpiGroupIds
	SetDpiSignatureIds(v *AddACLRuleResponseBodyDpiSignatureIds) *AddACLRuleResponseBody
	GetDpiSignatureIds() *AddACLRuleResponseBodyDpiSignatureIds
	SetGmtCreate(v int64) *AddACLRuleResponseBody
	GetGmtCreate() *int64
	SetIpProtocol(v string) *AddACLRuleResponseBody
	GetIpProtocol() *string
	SetName(v string) *AddACLRuleResponseBody
	GetName() *string
	SetPolicy(v string) *AddACLRuleResponseBody
	GetPolicy() *string
	SetPriority(v int32) *AddACLRuleResponseBody
	GetPriority() *int32
	SetRequestId(v string) *AddACLRuleResponseBody
	GetRequestId() *string
	SetSourceCidr(v string) *AddACLRuleResponseBody
	GetSourceCidr() *string
	SetSourcePortRange(v string) *AddACLRuleResponseBody
	GetSourcePortRange() *string
	SetType(v string) *AddACLRuleResponseBody
	GetType() *string
}

type AddACLRuleResponseBody struct {
	// The ID of the ACL.
	//
	// example:
	//
	// acl-xhwhyuo43l0*******
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ID of the ACL rule.
	//
	// example:
	//
	// acr-c1hkd054qywi******
	AcrId *string `json:"AcrId,omitempty" xml:"AcrId,omitempty"`
	// The description of the ACL rule.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block.
	//
	// For example: 192.168.10.0/24.
	//
	// example:
	//
	// 192.168.10.0/24
	DestCidr *string `json:"DestCidr,omitempty" xml:"DestCidr,omitempty"`
	// The destination port range.
	//
	// example:
	//
	// 1/65535
	DestPortRange *string `json:"DestPortRange,omitempty" xml:"DestPortRange,omitempty"`
	// The direction of traffic to which the ACL rule applies. Valid values:
	//
	// - **in**: inbound. Traffic from an external network to the local branch where the SAG instance is deployed.
	//
	// - **out**: outbound. Traffic from the local branch where the SAG instance is deployed to an external network.
	//
	// example:
	//
	// out
	Direction       *string                                `json:"Direction,omitempty" xml:"Direction,omitempty"`
	DpiGroupIds     *AddACLRuleResponseBodyDpiGroupIds     `json:"DpiGroupIds,omitempty" xml:"DpiGroupIds,omitempty" type:"Struct"`
	DpiSignatureIds *AddACLRuleResponseBodyDpiSignatureIds `json:"DpiSignatureIds,omitempty" xml:"DpiSignatureIds,omitempty" type:"Struct"`
	// The UNIX timestamp when the ACL rule was created.
	//
	// This value is a long integer. If rules have the same priority, the one with the earlier timestamp takes precedence.
	//
	// example:
	//
	// 1553766882689
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The protocol to which the ACL rule applies.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The name of the ACL rule.
	//
	// example:
	//
	// doctest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The authorization policy of the ACL rule.
	//
	// - **accept**: allows access.
	//
	// - **drop**: denies access.
	//
	// example:
	//
	// drop
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The priority of the ACL rule.
	//
	// A smaller value indicates a higher priority. If multiple rules have the same priority, the rule that is first delivered to the Smart Access Gateway device takes precedence.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 880F84CB-9B54-4413-A8A3-8832C82D1BC4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source CIDR block.
	//
	// The source CIDR block is in the CIDR format. For example: 192.168.1.0/24.
	//
	// example:
	//
	// 192.168.20.0/24
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The source port range.
	//
	// example:
	//
	// 1/65535
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The type of the ACL rule.
	//
	// - **LAN**: private network. The ACL rule controls traffic on private networks.
	//
	// - **WAN**: public network. The ACL rule controls traffic on public networks.
	//
	// example:
	//
	// LAN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddACLRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddACLRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddACLRuleResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *AddACLRuleResponseBody) GetAcrId() *string {
	return s.AcrId
}

func (s *AddACLRuleResponseBody) GetDescription() *string {
	return s.Description
}

func (s *AddACLRuleResponseBody) GetDestCidr() *string {
	return s.DestCidr
}

func (s *AddACLRuleResponseBody) GetDestPortRange() *string {
	return s.DestPortRange
}

func (s *AddACLRuleResponseBody) GetDirection() *string {
	return s.Direction
}

func (s *AddACLRuleResponseBody) GetDpiGroupIds() *AddACLRuleResponseBodyDpiGroupIds {
	return s.DpiGroupIds
}

func (s *AddACLRuleResponseBody) GetDpiSignatureIds() *AddACLRuleResponseBodyDpiSignatureIds {
	return s.DpiSignatureIds
}

func (s *AddACLRuleResponseBody) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *AddACLRuleResponseBody) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *AddACLRuleResponseBody) GetName() *string {
	return s.Name
}

func (s *AddACLRuleResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *AddACLRuleResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *AddACLRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddACLRuleResponseBody) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *AddACLRuleResponseBody) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *AddACLRuleResponseBody) GetType() *string {
	return s.Type
}

func (s *AddACLRuleResponseBody) SetAclId(v string) *AddACLRuleResponseBody {
	s.AclId = &v
	return s
}

func (s *AddACLRuleResponseBody) SetAcrId(v string) *AddACLRuleResponseBody {
	s.AcrId = &v
	return s
}

func (s *AddACLRuleResponseBody) SetDescription(v string) *AddACLRuleResponseBody {
	s.Description = &v
	return s
}

func (s *AddACLRuleResponseBody) SetDestCidr(v string) *AddACLRuleResponseBody {
	s.DestCidr = &v
	return s
}

func (s *AddACLRuleResponseBody) SetDestPortRange(v string) *AddACLRuleResponseBody {
	s.DestPortRange = &v
	return s
}

func (s *AddACLRuleResponseBody) SetDirection(v string) *AddACLRuleResponseBody {
	s.Direction = &v
	return s
}

func (s *AddACLRuleResponseBody) SetDpiGroupIds(v *AddACLRuleResponseBodyDpiGroupIds) *AddACLRuleResponseBody {
	s.DpiGroupIds = v
	return s
}

func (s *AddACLRuleResponseBody) SetDpiSignatureIds(v *AddACLRuleResponseBodyDpiSignatureIds) *AddACLRuleResponseBody {
	s.DpiSignatureIds = v
	return s
}

func (s *AddACLRuleResponseBody) SetGmtCreate(v int64) *AddACLRuleResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *AddACLRuleResponseBody) SetIpProtocol(v string) *AddACLRuleResponseBody {
	s.IpProtocol = &v
	return s
}

func (s *AddACLRuleResponseBody) SetName(v string) *AddACLRuleResponseBody {
	s.Name = &v
	return s
}

func (s *AddACLRuleResponseBody) SetPolicy(v string) *AddACLRuleResponseBody {
	s.Policy = &v
	return s
}

func (s *AddACLRuleResponseBody) SetPriority(v int32) *AddACLRuleResponseBody {
	s.Priority = &v
	return s
}

func (s *AddACLRuleResponseBody) SetRequestId(v string) *AddACLRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddACLRuleResponseBody) SetSourceCidr(v string) *AddACLRuleResponseBody {
	s.SourceCidr = &v
	return s
}

func (s *AddACLRuleResponseBody) SetSourcePortRange(v string) *AddACLRuleResponseBody {
	s.SourcePortRange = &v
	return s
}

func (s *AddACLRuleResponseBody) SetType(v string) *AddACLRuleResponseBody {
	s.Type = &v
	return s
}

func (s *AddACLRuleResponseBody) Validate() error {
	if s.DpiGroupIds != nil {
		if err := s.DpiGroupIds.Validate(); err != nil {
			return err
		}
	}
	if s.DpiSignatureIds != nil {
		if err := s.DpiSignatureIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddACLRuleResponseBodyDpiGroupIds struct {
	DpiGroupId []*string `json:"DpiGroupId,omitempty" xml:"DpiGroupId,omitempty" type:"Repeated"`
}

func (s AddACLRuleResponseBodyDpiGroupIds) String() string {
	return dara.Prettify(s)
}

func (s AddACLRuleResponseBodyDpiGroupIds) GoString() string {
	return s.String()
}

func (s *AddACLRuleResponseBodyDpiGroupIds) GetDpiGroupId() []*string {
	return s.DpiGroupId
}

func (s *AddACLRuleResponseBodyDpiGroupIds) SetDpiGroupId(v []*string) *AddACLRuleResponseBodyDpiGroupIds {
	s.DpiGroupId = v
	return s
}

func (s *AddACLRuleResponseBodyDpiGroupIds) Validate() error {
	return dara.Validate(s)
}

type AddACLRuleResponseBodyDpiSignatureIds struct {
	DpiSignatureId []*string `json:"DpiSignatureId,omitempty" xml:"DpiSignatureId,omitempty" type:"Repeated"`
}

func (s AddACLRuleResponseBodyDpiSignatureIds) String() string {
	return dara.Prettify(s)
}

func (s AddACLRuleResponseBodyDpiSignatureIds) GoString() string {
	return s.String()
}

func (s *AddACLRuleResponseBodyDpiSignatureIds) GetDpiSignatureId() []*string {
	return s.DpiSignatureId
}

func (s *AddACLRuleResponseBodyDpiSignatureIds) SetDpiSignatureId(v []*string) *AddACLRuleResponseBodyDpiSignatureIds {
	s.DpiSignatureId = v
	return s
}

func (s *AddACLRuleResponseBodyDpiSignatureIds) Validate() error {
	return dara.Validate(s)
}
