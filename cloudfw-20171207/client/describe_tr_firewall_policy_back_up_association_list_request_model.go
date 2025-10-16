// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallPolicyBackUpAssociationListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCandidateList(v []*DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList) *DescribeTrFirewallPolicyBackUpAssociationListRequest
	GetCandidateList() []*DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList
	SetFirewallId(v string) *DescribeTrFirewallPolicyBackUpAssociationListRequest
	GetFirewallId() *string
	SetLang(v string) *DescribeTrFirewallPolicyBackUpAssociationListRequest
	GetLang() *string
	SetTrFirewallRoutePolicyId(v string) *DescribeTrFirewallPolicyBackUpAssociationListRequest
	GetTrFirewallRoutePolicyId() *string
}

type DescribeTrFirewallPolicyBackUpAssociationListRequest struct {
	// The traffic redirection instances.
	CandidateList []*DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList `json:"CandidateList,omitempty" xml:"CandidateList,omitempty" type:"Repeated"`
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-tr-8b268ce1b26e4c68****
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
	// policy-5dcafb12ff794a56****
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s DescribeTrFirewallPolicyBackUpAssociationListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallPolicyBackUpAssociationListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequest) GetCandidateList() []*DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList {
	return s.CandidateList
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequest) GetTrFirewallRoutePolicyId() *string {
	return s.TrFirewallRoutePolicyId
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequest) SetCandidateList(v []*DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList) *DescribeTrFirewallPolicyBackUpAssociationListRequest {
	s.CandidateList = v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequest) SetFirewallId(v string) *DescribeTrFirewallPolicyBackUpAssociationListRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequest) SetLang(v string) *DescribeTrFirewallPolicyBackUpAssociationListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequest) SetTrFirewallRoutePolicyId(v string) *DescribeTrFirewallPolicyBackUpAssociationListRequest {
	s.TrFirewallRoutePolicyId = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequest) Validate() error {
	if s.CandidateList != nil {
		for _, item := range s.CandidateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList struct {
	// The ID of the traffic redirection instance.
	//
	// example:
	//
	// vpc-wz9grb8ng3y7h7lf2****
	CandidateId *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	// The type of the traffic redirection instance.
	//
	// example:
	//
	// VPC
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
}

func (s DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList) GetCandidateId() *string {
	return s.CandidateId
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList) GetCandidateType() *string {
	return s.CandidateType
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList) SetCandidateId(v string) *DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList {
	s.CandidateId = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList) SetCandidateType(v string) *DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList {
	s.CandidateType = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList) Validate() error {
	return dara.Validate(s)
}
