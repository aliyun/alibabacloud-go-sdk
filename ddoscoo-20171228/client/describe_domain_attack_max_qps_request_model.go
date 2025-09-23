// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAttackMaxQpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainAttackMaxQpsRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainAttackMaxQpsRequest
	GetEndTime() *int64
	SetResourceGroupId(v string) *DescribeDomainAttackMaxQpsRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainAttackMaxQpsRequest
	GetStartTime() *int64
}

type DescribeDomainAttackMaxQpsRequest struct {
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1667801940
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1657562370
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainAttackMaxQpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAttackMaxQpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackMaxQpsRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainAttackMaxQpsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainAttackMaxQpsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainAttackMaxQpsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainAttackMaxQpsRequest) SetDomain(v string) *DescribeDomainAttackMaxQpsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAttackMaxQpsRequest) SetEndTime(v int64) *DescribeDomainAttackMaxQpsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAttackMaxQpsRequest) SetResourceGroupId(v string) *DescribeDomainAttackMaxQpsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainAttackMaxQpsRequest) SetStartTime(v int64) *DescribeDomainAttackMaxQpsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAttackMaxQpsRequest) Validate() error {
	return dara.Validate(s)
}
