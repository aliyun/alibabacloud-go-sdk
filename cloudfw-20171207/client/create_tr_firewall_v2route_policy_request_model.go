// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrFirewallV2RoutePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestCandidateList(v []*CreateTrFirewallV2RoutePolicyRequestDestCandidateList) *CreateTrFirewallV2RoutePolicyRequest
	GetDestCandidateList() []*CreateTrFirewallV2RoutePolicyRequestDestCandidateList
	SetFirewallId(v string) *CreateTrFirewallV2RoutePolicyRequest
	GetFirewallId() *string
	SetLang(v string) *CreateTrFirewallV2RoutePolicyRequest
	GetLang() *string
	SetPolicyDescription(v string) *CreateTrFirewallV2RoutePolicyRequest
	GetPolicyDescription() *string
	SetPolicyName(v string) *CreateTrFirewallV2RoutePolicyRequest
	GetPolicyName() *string
	SetPolicyType(v string) *CreateTrFirewallV2RoutePolicyRequest
	GetPolicyType() *string
	SetSrcCandidateList(v []*CreateTrFirewallV2RoutePolicyRequestSrcCandidateList) *CreateTrFirewallV2RoutePolicyRequest
	GetSrcCandidateList() []*CreateTrFirewallV2RoutePolicyRequestSrcCandidateList
}

type CreateTrFirewallV2RoutePolicyRequest struct {
	// The secondary traffic redirection instances.
	DestCandidateList []*CreateTrFirewallV2RoutePolicyRequestDestCandidateList `json:"DestCandidateList,omitempty" xml:"DestCandidateList,omitempty" type:"Repeated"`
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-tr-f8ce36689b224f77****
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
	// The description of the traffic redirection instance.
	//
	// example:
	//
	// test
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The name of the traffic redirection instance.
	//
	// example:
	//
	// TEST_VPC_FW
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the traffic redirection scenario of the VPC firewall. Valid values:
	//
	// 	- **fullmesh**: interconnected instances
	//
	// 	- **one_to_one**: instance to instance
	//
	// 	- **end_to_end**: instance to instances
	//
	// example:
	//
	// fullmesh
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The primary traffic redirection instances.
	SrcCandidateList []*CreateTrFirewallV2RoutePolicyRequestSrcCandidateList `json:"SrcCandidateList,omitempty" xml:"SrcCandidateList,omitempty" type:"Repeated"`
}

func (s CreateTrFirewallV2RoutePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTrFirewallV2RoutePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2RoutePolicyRequest) GetDestCandidateList() []*CreateTrFirewallV2RoutePolicyRequestDestCandidateList {
	return s.DestCandidateList
}

func (s *CreateTrFirewallV2RoutePolicyRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *CreateTrFirewallV2RoutePolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateTrFirewallV2RoutePolicyRequest) GetPolicyDescription() *string {
	return s.PolicyDescription
}

func (s *CreateTrFirewallV2RoutePolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateTrFirewallV2RoutePolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateTrFirewallV2RoutePolicyRequest) GetSrcCandidateList() []*CreateTrFirewallV2RoutePolicyRequestSrcCandidateList {
	return s.SrcCandidateList
}

func (s *CreateTrFirewallV2RoutePolicyRequest) SetDestCandidateList(v []*CreateTrFirewallV2RoutePolicyRequestDestCandidateList) *CreateTrFirewallV2RoutePolicyRequest {
	s.DestCandidateList = v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequest) SetFirewallId(v string) *CreateTrFirewallV2RoutePolicyRequest {
	s.FirewallId = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequest) SetLang(v string) *CreateTrFirewallV2RoutePolicyRequest {
	s.Lang = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequest) SetPolicyDescription(v string) *CreateTrFirewallV2RoutePolicyRequest {
	s.PolicyDescription = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequest) SetPolicyName(v string) *CreateTrFirewallV2RoutePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequest) SetPolicyType(v string) *CreateTrFirewallV2RoutePolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequest) SetSrcCandidateList(v []*CreateTrFirewallV2RoutePolicyRequestSrcCandidateList) *CreateTrFirewallV2RoutePolicyRequest {
	s.SrcCandidateList = v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequest) Validate() error {
	if s.DestCandidateList != nil {
		for _, item := range s.DestCandidateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SrcCandidateList != nil {
		for _, item := range s.SrcCandidateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTrFirewallV2RoutePolicyRequestDestCandidateList struct {
	// The ID of the traffic redirection instance.
	//
	// example:
	//
	// vpc-2ze9epancaw8t4sha****
	CandidateId *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	// The type of the traffic redirection instance.
	//
	// example:
	//
	// VPC
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
}

func (s CreateTrFirewallV2RoutePolicyRequestDestCandidateList) String() string {
	return dara.Prettify(s)
}

func (s CreateTrFirewallV2RoutePolicyRequestDestCandidateList) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2RoutePolicyRequestDestCandidateList) GetCandidateId() *string {
	return s.CandidateId
}

func (s *CreateTrFirewallV2RoutePolicyRequestDestCandidateList) GetCandidateType() *string {
	return s.CandidateType
}

func (s *CreateTrFirewallV2RoutePolicyRequestDestCandidateList) SetCandidateId(v string) *CreateTrFirewallV2RoutePolicyRequestDestCandidateList {
	s.CandidateId = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequestDestCandidateList) SetCandidateType(v string) *CreateTrFirewallV2RoutePolicyRequestDestCandidateList {
	s.CandidateType = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequestDestCandidateList) Validate() error {
	return dara.Validate(s)
}

type CreateTrFirewallV2RoutePolicyRequestSrcCandidateList struct {
	// The ID of the traffic redirection instance.
	//
	// example:
	//
	// vpc-2ze9epancaw8t4sha****
	CandidateId *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	// The type of the traffic redirection instance.
	//
	// example:
	//
	// VPC
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
}

func (s CreateTrFirewallV2RoutePolicyRequestSrcCandidateList) String() string {
	return dara.Prettify(s)
}

func (s CreateTrFirewallV2RoutePolicyRequestSrcCandidateList) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2RoutePolicyRequestSrcCandidateList) GetCandidateId() *string {
	return s.CandidateId
}

func (s *CreateTrFirewallV2RoutePolicyRequestSrcCandidateList) GetCandidateType() *string {
	return s.CandidateType
}

func (s *CreateTrFirewallV2RoutePolicyRequestSrcCandidateList) SetCandidateId(v string) *CreateTrFirewallV2RoutePolicyRequestSrcCandidateList {
	s.CandidateId = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequestSrcCandidateList) SetCandidateType(v string) *CreateTrFirewallV2RoutePolicyRequestSrcCandidateList {
	s.CandidateType = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequestSrcCandidateList) Validate() error {
	return dara.Validate(s)
}
