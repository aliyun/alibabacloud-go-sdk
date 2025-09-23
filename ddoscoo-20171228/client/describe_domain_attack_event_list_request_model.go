// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAttackEventListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainAttackEventListRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainAttackEventListRequest
	GetEndTime() *int64
	SetPageNo(v int32) *DescribeDomainAttackEventListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeDomainAttackEventListRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeDomainAttackEventListRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainAttackEventListRequest
	GetStartTime() *int64
}

type DescribeDomainAttackEventListRequest struct {
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1668740400
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1681966800
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainAttackEventListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAttackEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventListRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainAttackEventListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainAttackEventListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeDomainAttackEventListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDomainAttackEventListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainAttackEventListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainAttackEventListRequest) SetDomain(v string) *DescribeDomainAttackEventListRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAttackEventListRequest) SetEndTime(v int64) *DescribeDomainAttackEventListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAttackEventListRequest) SetPageNo(v int32) *DescribeDomainAttackEventListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeDomainAttackEventListRequest) SetPageSize(v int32) *DescribeDomainAttackEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainAttackEventListRequest) SetResourceGroupId(v string) *DescribeDomainAttackEventListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainAttackEventListRequest) SetStartTime(v int64) *DescribeDomainAttackEventListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAttackEventListRequest) Validate() error {
	return dara.Validate(s)
}
