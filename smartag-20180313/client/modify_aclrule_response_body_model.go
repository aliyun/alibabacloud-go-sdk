// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyACLRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *ModifyACLRuleResponseBody
	GetAclId() *string
	SetAcrId(v string) *ModifyACLRuleResponseBody
	GetAcrId() *string
	SetDescription(v string) *ModifyACLRuleResponseBody
	GetDescription() *string
	SetDestCidr(v string) *ModifyACLRuleResponseBody
	GetDestCidr() *string
	SetDestPortRange(v string) *ModifyACLRuleResponseBody
	GetDestPortRange() *string
	SetDirection(v string) *ModifyACLRuleResponseBody
	GetDirection() *string
	SetDpiGroupIds(v *ModifyACLRuleResponseBodyDpiGroupIds) *ModifyACLRuleResponseBody
	GetDpiGroupIds() *ModifyACLRuleResponseBodyDpiGroupIds
	SetDpiSignatureIds(v *ModifyACLRuleResponseBodyDpiSignatureIds) *ModifyACLRuleResponseBody
	GetDpiSignatureIds() *ModifyACLRuleResponseBodyDpiSignatureIds
	SetGmtCreate(v int64) *ModifyACLRuleResponseBody
	GetGmtCreate() *int64
	SetIpProtocol(v string) *ModifyACLRuleResponseBody
	GetIpProtocol() *string
	SetName(v string) *ModifyACLRuleResponseBody
	GetName() *string
	SetPolicy(v string) *ModifyACLRuleResponseBody
	GetPolicy() *string
	SetPriority(v int32) *ModifyACLRuleResponseBody
	GetPriority() *int32
	SetRequestId(v string) *ModifyACLRuleResponseBody
	GetRequestId() *string
	SetSourceCidr(v string) *ModifyACLRuleResponseBody
	GetSourceCidr() *string
	SetSourcePortRange(v string) *ModifyACLRuleResponseBody
	GetSourcePortRange() *string
}

type ModifyACLRuleResponseBody struct {
	// The ID of the ACL instance.
	//
	// example:
	//
	// acl-jdc7tir4fkplwr****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ID of the access control rule.
	//
	// example:
	//
	// acr-r8hezn2pi39s5a****
	AcrId *string `json:"AcrId,omitempty" xml:"AcrId,omitempty"`
	// The description of the access control rule.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block.
	//
	// The destination CIDR block is in CIDR format. For example: 192.168.10.0/24.
	//
	// example:
	//
	// 0.0.0.0/0
	DestCidr *string `json:"DestCidr,omitempty" xml:"DestCidr,omitempty"`
	// The destination port range.
	//
	// example:
	//
	// -1/-1
	DestPortRange *string `json:"DestPortRange,omitempty" xml:"DestPortRange,omitempty"`
	// The direction in which the access control rule is applied. Valid values:
	//
	// - **in**: inbound. This is the direction of traffic from an external network to the on-premises network where the Smart Access Gateway instance is deployed.
	//
	// - **out**: outbound. This is the direction of traffic from the on-premises network where the Smart Access Gateway instance is deployed to an external network.
	//
	// example:
	//
	// in
	Direction       *string                                   `json:"Direction,omitempty" xml:"Direction,omitempty"`
	DpiGroupIds     *ModifyACLRuleResponseBodyDpiGroupIds     `json:"DpiGroupIds,omitempty" xml:"DpiGroupIds,omitempty" type:"Struct"`
	DpiSignatureIds *ModifyACLRuleResponseBodyDpiSignatureIds `json:"DpiSignatureIds,omitempty" xml:"DpiSignatureIds,omitempty" type:"Struct"`
	// The timestamp that indicates when the access control rule was created.
	//
	// The value is a UNIX timestamp. If rules have the same priority, the one with the smaller timestamp takes precedence.
	//
	// example:
	//
	// 1553777700000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The protocol used by the access control rule.
	//
	// example:
	//
	// ALL
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The name of the access control rule.
	//
	// example:
	//
	// doctest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The authorization policy of the access control rule.
	//
	// - **accept**: allows access.
	//
	// - **drop**: denies access.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The priority of the access control rule.
	//
	// A smaller value indicates a higher priority. If rules have the same priority, the one that is first delivered to the Smart Access Gateway device takes precedence.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7F3DD2C1-0F6B-4575-9106-B2D50DF7A711
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source CIDR block.
	//
	// The source CIDR block is in CIDR format. For example: 192.168.1.0/24.
	//
	// example:
	//
	// 0.0.0.0/0
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The source port range.
	//
	// example:
	//
	// -1/-1
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s ModifyACLRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyACLRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyACLRuleResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *ModifyACLRuleResponseBody) GetAcrId() *string {
	return s.AcrId
}

func (s *ModifyACLRuleResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ModifyACLRuleResponseBody) GetDestCidr() *string {
	return s.DestCidr
}

func (s *ModifyACLRuleResponseBody) GetDestPortRange() *string {
	return s.DestPortRange
}

func (s *ModifyACLRuleResponseBody) GetDirection() *string {
	return s.Direction
}

func (s *ModifyACLRuleResponseBody) GetDpiGroupIds() *ModifyACLRuleResponseBodyDpiGroupIds {
	return s.DpiGroupIds
}

func (s *ModifyACLRuleResponseBody) GetDpiSignatureIds() *ModifyACLRuleResponseBodyDpiSignatureIds {
	return s.DpiSignatureIds
}

func (s *ModifyACLRuleResponseBody) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ModifyACLRuleResponseBody) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifyACLRuleResponseBody) GetName() *string {
	return s.Name
}

func (s *ModifyACLRuleResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *ModifyACLRuleResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyACLRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyACLRuleResponseBody) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *ModifyACLRuleResponseBody) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *ModifyACLRuleResponseBody) SetAclId(v string) *ModifyACLRuleResponseBody {
	s.AclId = &v
	return s
}

func (s *ModifyACLRuleResponseBody) SetAcrId(v string) *ModifyACLRuleResponseBody {
	s.AcrId = &v
	return s
}

func (s *ModifyACLRuleResponseBody) SetDescription(v string) *ModifyACLRuleResponseBody {
	s.Description = &v
	return s
}

func (s *ModifyACLRuleResponseBody) SetDestCidr(v string) *ModifyACLRuleResponseBody {
	s.DestCidr = &v
	return s
}

func (s *ModifyACLRuleResponseBody) SetDestPortRange(v string) *ModifyACLRuleResponseBody {
	s.DestPortRange = &v
	return s
}

func (s *ModifyACLRuleResponseBody) SetDirection(v string) *ModifyACLRuleResponseBody {
	s.Direction = &v
	return s
}

func (s *ModifyACLRuleResponseBody) SetDpiGroupIds(v *ModifyACLRuleResponseBodyDpiGroupIds) *ModifyACLRuleResponseBody {
	s.DpiGroupIds = v
	return s
}

func (s *ModifyACLRuleResponseBody) SetDpiSignatureIds(v *ModifyACLRuleResponseBodyDpiSignatureIds) *ModifyACLRuleResponseBody {
	s.DpiSignatureIds = v
	return s
}

func (s *ModifyACLRuleResponseBody) SetGmtCreate(v int64) *ModifyACLRuleResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *ModifyACLRuleResponseBody) SetIpProtocol(v string) *ModifyACLRuleResponseBody {
	s.IpProtocol = &v
	return s
}

func (s *ModifyACLRuleResponseBody) SetName(v string) *ModifyACLRuleResponseBody {
	s.Name = &v
	return s
}

func (s *ModifyACLRuleResponseBody) SetPolicy(v string) *ModifyACLRuleResponseBody {
	s.Policy = &v
	return s
}

func (s *ModifyACLRuleResponseBody) SetPriority(v int32) *ModifyACLRuleResponseBody {
	s.Priority = &v
	return s
}

func (s *ModifyACLRuleResponseBody) SetRequestId(v string) *ModifyACLRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyACLRuleResponseBody) SetSourceCidr(v string) *ModifyACLRuleResponseBody {
	s.SourceCidr = &v
	return s
}

func (s *ModifyACLRuleResponseBody) SetSourcePortRange(v string) *ModifyACLRuleResponseBody {
	s.SourcePortRange = &v
	return s
}

func (s *ModifyACLRuleResponseBody) Validate() error {
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

type ModifyACLRuleResponseBodyDpiGroupIds struct {
	DpiGroupId []*string `json:"DpiGroupId,omitempty" xml:"DpiGroupId,omitempty" type:"Repeated"`
}

func (s ModifyACLRuleResponseBodyDpiGroupIds) String() string {
	return dara.Prettify(s)
}

func (s ModifyACLRuleResponseBodyDpiGroupIds) GoString() string {
	return s.String()
}

func (s *ModifyACLRuleResponseBodyDpiGroupIds) GetDpiGroupId() []*string {
	return s.DpiGroupId
}

func (s *ModifyACLRuleResponseBodyDpiGroupIds) SetDpiGroupId(v []*string) *ModifyACLRuleResponseBodyDpiGroupIds {
	s.DpiGroupId = v
	return s
}

func (s *ModifyACLRuleResponseBodyDpiGroupIds) Validate() error {
	return dara.Validate(s)
}

type ModifyACLRuleResponseBodyDpiSignatureIds struct {
	DpiSignatureId []*string `json:"DpiSignatureId,omitempty" xml:"DpiSignatureId,omitempty" type:"Repeated"`
}

func (s ModifyACLRuleResponseBodyDpiSignatureIds) String() string {
	return dara.Prettify(s)
}

func (s ModifyACLRuleResponseBodyDpiSignatureIds) GoString() string {
	return s.String()
}

func (s *ModifyACLRuleResponseBodyDpiSignatureIds) GetDpiSignatureId() []*string {
	return s.DpiSignatureId
}

func (s *ModifyACLRuleResponseBodyDpiSignatureIds) SetDpiSignatureId(v []*string) *ModifyACLRuleResponseBodyDpiSignatureIds {
	s.DpiSignatureId = v
	return s
}

func (s *ModifyACLRuleResponseBodyDpiSignatureIds) Validate() error {
	return dara.Validate(s)
}
