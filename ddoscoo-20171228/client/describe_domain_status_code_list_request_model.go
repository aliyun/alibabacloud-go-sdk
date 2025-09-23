// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainStatusCodeListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainStatusCodeListRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainStatusCodeListRequest
	GetEndTime() *int64
	SetInterval(v int64) *DescribeDomainStatusCodeListRequest
	GetInterval() *int64
	SetQueryType(v string) *DescribeDomainStatusCodeListRequest
	GetQueryType() *string
	SetResourceGroupId(v string) *DescribeDomainStatusCodeListRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainStatusCodeListRequest
	GetStartTime() *int64
}

type DescribeDomainStatusCodeListRequest struct {
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1583683200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1800
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// upstream
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1582992000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainStatusCodeListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatusCodeListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeListRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainStatusCodeListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainStatusCodeListRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeDomainStatusCodeListRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeDomainStatusCodeListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainStatusCodeListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainStatusCodeListRequest) SetDomain(v string) *DescribeDomainStatusCodeListRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetEndTime(v int64) *DescribeDomainStatusCodeListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetInterval(v int64) *DescribeDomainStatusCodeListRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetQueryType(v string) *DescribeDomainStatusCodeListRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetResourceGroupId(v string) *DescribeDomainStatusCodeListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetStartTime(v int64) *DescribeDomainStatusCodeListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) Validate() error {
	return dara.Validate(s)
}
