// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizeEgress(v []*ModifySecurityGroupAttributeRequestAuthorizeEgress) *ModifySecurityGroupAttributeRequest
	GetAuthorizeEgress() []*ModifySecurityGroupAttributeRequestAuthorizeEgress
	SetAuthorizeIngress(v []*ModifySecurityGroupAttributeRequestAuthorizeIngress) *ModifySecurityGroupAttributeRequest
	GetAuthorizeIngress() []*ModifySecurityGroupAttributeRequestAuthorizeIngress
	SetOfficeSiteId(v string) *ModifySecurityGroupAttributeRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *ModifySecurityGroupAttributeRequest
	GetRegionId() *string
	SetRevokeEgress(v []*ModifySecurityGroupAttributeRequestRevokeEgress) *ModifySecurityGroupAttributeRequest
	GetRevokeEgress() []*ModifySecurityGroupAttributeRequestRevokeEgress
	SetRevokeIngress(v []*ModifySecurityGroupAttributeRequestRevokeIngress) *ModifySecurityGroupAttributeRequest
	GetRevokeIngress() []*ModifySecurityGroupAttributeRequestRevokeIngress
}

type ModifySecurityGroupAttributeRequest struct {
	AuthorizeEgress  []*ModifySecurityGroupAttributeRequestAuthorizeEgress  `json:"AuthorizeEgress,omitempty" xml:"AuthorizeEgress,omitempty" type:"Repeated"`
	AuthorizeIngress []*ModifySecurityGroupAttributeRequestAuthorizeIngress `json:"AuthorizeIngress,omitempty" xml:"AuthorizeIngress,omitempty" type:"Repeated"`
	// This parameter is required.
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// This parameter is required.
	RegionId      *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RevokeEgress  []*ModifySecurityGroupAttributeRequestRevokeEgress  `json:"RevokeEgress,omitempty" xml:"RevokeEgress,omitempty" type:"Repeated"`
	RevokeIngress []*ModifySecurityGroupAttributeRequestRevokeIngress `json:"RevokeIngress,omitempty" xml:"RevokeIngress,omitempty" type:"Repeated"`
}

func (s ModifySecurityGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupAttributeRequest) GetAuthorizeEgress() []*ModifySecurityGroupAttributeRequestAuthorizeEgress {
	return s.AuthorizeEgress
}

func (s *ModifySecurityGroupAttributeRequest) GetAuthorizeIngress() []*ModifySecurityGroupAttributeRequestAuthorizeIngress {
	return s.AuthorizeIngress
}

func (s *ModifySecurityGroupAttributeRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ModifySecurityGroupAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySecurityGroupAttributeRequest) GetRevokeEgress() []*ModifySecurityGroupAttributeRequestRevokeEgress {
	return s.RevokeEgress
}

func (s *ModifySecurityGroupAttributeRequest) GetRevokeIngress() []*ModifySecurityGroupAttributeRequestRevokeIngress {
	return s.RevokeIngress
}

func (s *ModifySecurityGroupAttributeRequest) SetAuthorizeEgress(v []*ModifySecurityGroupAttributeRequestAuthorizeEgress) *ModifySecurityGroupAttributeRequest {
	s.AuthorizeEgress = v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetAuthorizeIngress(v []*ModifySecurityGroupAttributeRequestAuthorizeIngress) *ModifySecurityGroupAttributeRequest {
	s.AuthorizeIngress = v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetOfficeSiteId(v string) *ModifySecurityGroupAttributeRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetRegionId(v string) *ModifySecurityGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetRevokeEgress(v []*ModifySecurityGroupAttributeRequestRevokeEgress) *ModifySecurityGroupAttributeRequest {
	s.RevokeEgress = v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetRevokeIngress(v []*ModifySecurityGroupAttributeRequestRevokeIngress) *ModifySecurityGroupAttributeRequest {
	s.RevokeIngress = v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) Validate() error {
	if s.AuthorizeEgress != nil {
		for _, item := range s.AuthorizeEgress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AuthorizeIngress != nil {
		for _, item := range s.AuthorizeIngress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RevokeEgress != nil {
		for _, item := range s.RevokeEgress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RevokeIngress != nil {
		for _, item := range s.RevokeIngress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifySecurityGroupAttributeRequestAuthorizeEgress struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestCidrIp      *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	NicType         *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	Priority        *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SourceCidrIp    *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s ModifySecurityGroupAttributeRequestAuthorizeEgress) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupAttributeRequestAuthorizeEgress) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) GetDescription() *string {
	return s.Description
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) GetNicType() *string {
	return s.NicType
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) GetPolicy() *string {
	return s.Policy
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) GetPriority() *string {
	return s.Priority
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) SetDescription(v string) *ModifySecurityGroupAttributeRequestAuthorizeEgress {
	s.Description = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) SetDestCidrIp(v string) *ModifySecurityGroupAttributeRequestAuthorizeEgress {
	s.DestCidrIp = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) SetIpProtocol(v string) *ModifySecurityGroupAttributeRequestAuthorizeEgress {
	s.IpProtocol = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) SetNicType(v string) *ModifySecurityGroupAttributeRequestAuthorizeEgress {
	s.NicType = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) SetPolicy(v string) *ModifySecurityGroupAttributeRequestAuthorizeEgress {
	s.Policy = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) SetPortRange(v string) *ModifySecurityGroupAttributeRequestAuthorizeEgress {
	s.PortRange = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) SetPriority(v string) *ModifySecurityGroupAttributeRequestAuthorizeEgress {
	s.Priority = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) SetSourceCidrIp(v string) *ModifySecurityGroupAttributeRequestAuthorizeEgress {
	s.SourceCidrIp = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) SetSourcePortRange(v string) *ModifySecurityGroupAttributeRequestAuthorizeEgress {
	s.SourcePortRange = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeEgress) Validate() error {
	return dara.Validate(s)
}

type ModifySecurityGroupAttributeRequestAuthorizeIngress struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestCidrIp      *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	NicType         *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	Priority        *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SourceCidrIp    *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s ModifySecurityGroupAttributeRequestAuthorizeIngress) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupAttributeRequestAuthorizeIngress) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) GetDescription() *string {
	return s.Description
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) GetNicType() *string {
	return s.NicType
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) GetPolicy() *string {
	return s.Policy
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) GetPriority() *string {
	return s.Priority
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) SetDescription(v string) *ModifySecurityGroupAttributeRequestAuthorizeIngress {
	s.Description = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) SetDestCidrIp(v string) *ModifySecurityGroupAttributeRequestAuthorizeIngress {
	s.DestCidrIp = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) SetIpProtocol(v string) *ModifySecurityGroupAttributeRequestAuthorizeIngress {
	s.IpProtocol = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) SetNicType(v string) *ModifySecurityGroupAttributeRequestAuthorizeIngress {
	s.NicType = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) SetPolicy(v string) *ModifySecurityGroupAttributeRequestAuthorizeIngress {
	s.Policy = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) SetPortRange(v string) *ModifySecurityGroupAttributeRequestAuthorizeIngress {
	s.PortRange = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) SetPriority(v string) *ModifySecurityGroupAttributeRequestAuthorizeIngress {
	s.Priority = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) SetSourceCidrIp(v string) *ModifySecurityGroupAttributeRequestAuthorizeIngress {
	s.SourceCidrIp = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) SetSourcePortRange(v string) *ModifySecurityGroupAttributeRequestAuthorizeIngress {
	s.SourcePortRange = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestAuthorizeIngress) Validate() error {
	return dara.Validate(s)
}

type ModifySecurityGroupAttributeRequestRevokeEgress struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestCidrIp      *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	NicType         *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	Priority        *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SourceCidrIp    *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s ModifySecurityGroupAttributeRequestRevokeEgress) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupAttributeRequestRevokeEgress) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) GetDescription() *string {
	return s.Description
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) GetNicType() *string {
	return s.NicType
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) GetPolicy() *string {
	return s.Policy
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) GetPriority() *string {
	return s.Priority
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) SetDescription(v string) *ModifySecurityGroupAttributeRequestRevokeEgress {
	s.Description = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) SetDestCidrIp(v string) *ModifySecurityGroupAttributeRequestRevokeEgress {
	s.DestCidrIp = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) SetIpProtocol(v string) *ModifySecurityGroupAttributeRequestRevokeEgress {
	s.IpProtocol = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) SetNicType(v string) *ModifySecurityGroupAttributeRequestRevokeEgress {
	s.NicType = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) SetPolicy(v string) *ModifySecurityGroupAttributeRequestRevokeEgress {
	s.Policy = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) SetPortRange(v string) *ModifySecurityGroupAttributeRequestRevokeEgress {
	s.PortRange = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) SetPriority(v string) *ModifySecurityGroupAttributeRequestRevokeEgress {
	s.Priority = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) SetSourceCidrIp(v string) *ModifySecurityGroupAttributeRequestRevokeEgress {
	s.SourceCidrIp = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) SetSourcePortRange(v string) *ModifySecurityGroupAttributeRequestRevokeEgress {
	s.SourcePortRange = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeEgress) Validate() error {
	return dara.Validate(s)
}

type ModifySecurityGroupAttributeRequestRevokeIngress struct {
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestCidrIp      *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	IpProtocol      *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	NicType         *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	Priority        *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SourceCidrIp    *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s ModifySecurityGroupAttributeRequestRevokeIngress) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupAttributeRequestRevokeIngress) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) GetDescription() *string {
	return s.Description
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) GetNicType() *string {
	return s.NicType
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) GetPolicy() *string {
	return s.Policy
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) GetPriority() *string {
	return s.Priority
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) SetDescription(v string) *ModifySecurityGroupAttributeRequestRevokeIngress {
	s.Description = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) SetDestCidrIp(v string) *ModifySecurityGroupAttributeRequestRevokeIngress {
	s.DestCidrIp = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) SetIpProtocol(v string) *ModifySecurityGroupAttributeRequestRevokeIngress {
	s.IpProtocol = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) SetNicType(v string) *ModifySecurityGroupAttributeRequestRevokeIngress {
	s.NicType = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) SetPolicy(v string) *ModifySecurityGroupAttributeRequestRevokeIngress {
	s.Policy = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) SetPortRange(v string) *ModifySecurityGroupAttributeRequestRevokeIngress {
	s.PortRange = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) SetPriority(v string) *ModifySecurityGroupAttributeRequestRevokeIngress {
	s.Priority = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) SetSourceCidrIp(v string) *ModifySecurityGroupAttributeRequestRevokeIngress {
	s.SourceCidrIp = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) SetSourcePortRange(v string) *ModifySecurityGroupAttributeRequestRevokeIngress {
	s.SourcePortRange = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequestRevokeIngress) Validate() error {
	return dara.Validate(s)
}
