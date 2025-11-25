// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopAttackListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttackList(v []*DescribeDomainTopAttackListResponseBodyAttackList) *DescribeDomainTopAttackListResponseBody
	GetAttackList() []*DescribeDomainTopAttackListResponseBodyAttackList
	SetRequestId(v string) *DescribeDomainTopAttackListResponseBody
	GetRequestId() *string
}

type DescribeDomainTopAttackListResponseBody struct {
	// The peak QPS of the website.
	AttackList []*DescribeDomainTopAttackListResponseBodyAttackList `json:"AttackList,omitempty" xml:"AttackList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainTopAttackListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopAttackListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopAttackListResponseBody) GetAttackList() []*DescribeDomainTopAttackListResponseBodyAttackList {
	return s.AttackList
}

func (s *DescribeDomainTopAttackListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainTopAttackListResponseBody) SetAttackList(v []*DescribeDomainTopAttackListResponseBodyAttackList) *DescribeDomainTopAttackListResponseBody {
	s.AttackList = v
	return s
}

func (s *DescribeDomainTopAttackListResponseBody) SetRequestId(v string) *DescribeDomainTopAttackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainTopAttackListResponseBody) Validate() error {
	if s.AttackList != nil {
		for _, item := range s.AttackList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainTopAttackListResponseBodyAttackList struct {
	// The attack QPS. Unit: QPS
	//
	// example:
	//
	// 0
	Attack *int64 `json:"Attack,omitempty" xml:"Attack,omitempty"`
	// The number of all QPS, which includes normal and attack QPS. Unit: QPS.
	//
	// example:
	//
	// 294
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The domain name of the website.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainTopAttackListResponseBodyAttackList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopAttackListResponseBodyAttackList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopAttackListResponseBodyAttackList) GetAttack() *int64 {
	return s.Attack
}

func (s *DescribeDomainTopAttackListResponseBodyAttackList) GetCount() *int64 {
	return s.Count
}

func (s *DescribeDomainTopAttackListResponseBodyAttackList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainTopAttackListResponseBodyAttackList) SetAttack(v int64) *DescribeDomainTopAttackListResponseBodyAttackList {
	s.Attack = &v
	return s
}

func (s *DescribeDomainTopAttackListResponseBodyAttackList) SetCount(v int64) *DescribeDomainTopAttackListResponseBodyAttackList {
	s.Count = &v
	return s
}

func (s *DescribeDomainTopAttackListResponseBodyAttackList) SetDomain(v string) *DescribeDomainTopAttackListResponseBodyAttackList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainTopAttackListResponseBodyAttackList) Validate() error {
	return dara.Validate(s)
}
