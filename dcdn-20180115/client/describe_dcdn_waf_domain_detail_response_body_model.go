// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafDomainDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v *DescribeDcdnWafDomainDetailResponseBodyDomain) *DescribeDcdnWafDomainDetailResponseBody
	GetDomain() *DescribeDcdnWafDomainDetailResponseBodyDomain
	SetRequestId(v string) *DescribeDcdnWafDomainDetailResponseBody
	GetRequestId() *string
}

type DescribeDcdnWafDomainDetailResponseBody struct {
	// The information about the accelerated domain name.
	Domain *DescribeDcdnWafDomainDetailResponseBodyDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 153ca2cd-3c01-44be-82C0-64dbc6c88630
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnWafDomainDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainDetailResponseBody) GetDomain() *DescribeDcdnWafDomainDetailResponseBodyDomain {
	return s.Domain
}

func (s *DescribeDcdnWafDomainDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafDomainDetailResponseBody) SetDomain(v *DescribeDcdnWafDomainDetailResponseBodyDomain) *DescribeDcdnWafDomainDetailResponseBody {
	s.Domain = v
	return s
}

func (s *DescribeDcdnWafDomainDetailResponseBody) SetRequestId(v string) *DescribeDcdnWafDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafDomainDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafDomainDetailResponseBodyDomain struct {
	// The types of the protection policies.
	DefenseScenes []*DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes `json:"DefenseScenes,omitempty" xml:"DefenseScenes,omitempty" type:"Repeated"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDcdnWafDomainDetailResponseBodyDomain) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDomainDetailResponseBodyDomain) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainDetailResponseBodyDomain) GetDefenseScenes() []*DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes {
	return s.DefenseScenes
}

func (s *DescribeDcdnWafDomainDetailResponseBodyDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnWafDomainDetailResponseBodyDomain) SetDefenseScenes(v []*DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes) *DescribeDcdnWafDomainDetailResponseBodyDomain {
	s.DefenseScenes = v
	return s
}

func (s *DescribeDcdnWafDomainDetailResponseBodyDomain) SetDomainName(v string) *DescribeDcdnWafDomainDetailResponseBodyDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnWafDomainDetailResponseBodyDomain) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes struct {
	// The type of the protection policy. Valid values:
	//
	// 	- waf_group: basic web protection
	//
	// 	- custom_acl: custom protection
	//
	// 	- whitelist: whitelist
	//
	// example:
	//
	// waf_group
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The ID of the protection policy.
	//
	// example:
	//
	// 10000002
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The IDs of the protection policy.
	//
	// example:
	//
	// 10000001,10000004
	PolicyIds *string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty"`
}

func (s DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes) GetPolicyIds() *string {
	return s.PolicyIds
}

func (s *DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes) SetDefenseScene(v string) *DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes) SetPolicyId(v int64) *DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes {
	s.PolicyId = &v
	return s
}

func (s *DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes) SetPolicyIds(v string) *DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes {
	s.PolicyIds = &v
	return s
}

func (s *DescribeDcdnWafDomainDetailResponseBodyDomainDefenseScenes) Validate() error {
	return dara.Validate(s)
}
