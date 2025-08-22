// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnTopDomainsByFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDcdnTopDomainsByFlowRequest
	GetEndTime() *string
	SetLimit(v int64) *DescribeDcdnTopDomainsByFlowRequest
	GetLimit() *int64
	SetStartTime(v string) *DescribeDcdnTopDomainsByFlowRequest
	GetStartTime() *string
}

type DescribeDcdnTopDomainsByFlowRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The end time must be later than the start time.
	//
	// example:
	//
	// 2016-03-14T07:34:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of domain names to return. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The start of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2016-03-01T04:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnTopDomainsByFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnTopDomainsByFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTopDomainsByFlowRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnTopDomainsByFlowRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeDcdnTopDomainsByFlowRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnTopDomainsByFlowRequest) SetEndTime(v string) *DescribeDcdnTopDomainsByFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowRequest) SetLimit(v int64) *DescribeDcdnTopDomainsByFlowRequest {
	s.Limit = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowRequest) SetStartTime(v string) *DescribeDcdnTopDomainsByFlowRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowRequest) Validate() error {
	return dara.Validate(s)
}
