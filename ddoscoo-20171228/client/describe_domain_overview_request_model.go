// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainOverviewRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainOverviewRequest
	GetEndTime() *int64
	SetResourceGroupId(v string) *DescribeDomainOverviewRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainOverviewRequest
	GetStartTime() *int64
}

type DescribeDomainOverviewRequest struct {
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1651809600
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1619798400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainOverviewRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainOverviewRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainOverviewRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainOverviewRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainOverviewRequest) SetDomain(v string) *DescribeDomainOverviewRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainOverviewRequest) SetEndTime(v int64) *DescribeDomainOverviewRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainOverviewRequest) SetResourceGroupId(v string) *DescribeDomainOverviewRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainOverviewRequest) SetStartTime(v int64) *DescribeDomainOverviewRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainOverviewRequest) Validate() error {
	return dara.Validate(s)
}
