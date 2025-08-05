// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallPolicyBackUpAssociationListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyAssociationBackupConfigs(v []*DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) *DescribeTrFirewallPolicyBackUpAssociationListResponseBody
	GetPolicyAssociationBackupConfigs() []*DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs
	SetRequestId(v string) *DescribeTrFirewallPolicyBackUpAssociationListResponseBody
	GetRequestId() *string
}

type DescribeTrFirewallPolicyBackUpAssociationListResponseBody struct {
	// The route tables.
	PolicyAssociationBackupConfigs []*DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs `json:"PolicyAssociationBackupConfigs,omitempty" xml:"PolicyAssociationBackupConfigs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C264A756-9B48-57E3-B312-716941E146C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTrFirewallPolicyBackUpAssociationListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallPolicyBackUpAssociationListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBody) GetPolicyAssociationBackupConfigs() []*DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs {
	return s.PolicyAssociationBackupConfigs
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBody) SetPolicyAssociationBackupConfigs(v []*DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) *DescribeTrFirewallPolicyBackUpAssociationListResponseBody {
	s.PolicyAssociationBackupConfigs = v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBody) SetRequestId(v string) *DescribeTrFirewallPolicyBackUpAssociationListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs struct {
	// The ID of the traffic redirection instance.
	//
	// example:
	//
	// vpc-wz9grb8ng3y7h7lf2****
	CandidateId *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	// The name of the traffic redirection instance.
	//
	// example:
	//
	// test
	CandidateName *string `json:"CandidateName,omitempty" xml:"CandidateName,omitempty"`
	// The type of the traffic redirection instance.
	//
	// example:
	//
	// VPC
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
	// The route table that is used after traffic redirection.
	//
	// example:
	//
	// vtb-wz9898grickmh5j09****
	CurrentRouteTableId *string `json:"CurrentRouteTableId,omitempty" xml:"CurrentRouteTableId,omitempty"`
	// The ID of the route table.
	//
	// example:
	//
	// vtb-wz9slp3s7m4qrzvnq****
	OriginalRouteTableId *string `json:"OriginalRouteTableId,omitempty" xml:"OriginalRouteTableId,omitempty"`
}

func (s DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) GetCandidateId() *string {
	return s.CandidateId
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) GetCandidateName() *string {
	return s.CandidateName
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) GetCandidateType() *string {
	return s.CandidateType
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) GetCurrentRouteTableId() *string {
	return s.CurrentRouteTableId
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) GetOriginalRouteTableId() *string {
	return s.OriginalRouteTableId
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) SetCandidateId(v string) *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs {
	s.CandidateId = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) SetCandidateName(v string) *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs {
	s.CandidateName = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) SetCandidateType(v string) *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs {
	s.CandidateType = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) SetCurrentRouteTableId(v string) *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs {
	s.CurrentRouteTableId = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) SetOriginalRouteTableId(v string) *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs {
	s.OriginalRouteTableId = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) Validate() error {
	return dara.Validate(s)
}
