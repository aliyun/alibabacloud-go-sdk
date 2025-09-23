// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQpsListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainQpsListRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainQpsListRequest
	GetEndTime() *int64
	SetInterval(v int64) *DescribeDomainQpsListRequest
	GetInterval() *int64
	SetResourceGroupId(v string) *DescribeDomainQpsListRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainQpsListRequest
	GetStartTime() *int64
}

type DescribeDomainQpsListRequest struct {
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1657162260
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1800
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1672362360
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainQpsListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsListRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainQpsListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainQpsListRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeDomainQpsListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainQpsListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainQpsListRequest) SetDomain(v string) *DescribeDomainQpsListRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainQpsListRequest) SetEndTime(v int64) *DescribeDomainQpsListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQpsListRequest) SetInterval(v int64) *DescribeDomainQpsListRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQpsListRequest) SetResourceGroupId(v string) *DescribeDomainQpsListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainQpsListRequest) SetStartTime(v int64) *DescribeDomainQpsListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsListRequest) Validate() error {
	return dara.Validate(s)
}
