// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveTopDomainsByFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeLiveTopDomainsByFlowRequest
	GetEndTime() *string
	SetLimit(v int64) *DescribeLiveTopDomainsByFlowRequest
	GetLimit() *int64
	SetOwnerId(v int64) *DescribeLiveTopDomainsByFlowRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveTopDomainsByFlowRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveTopDomainsByFlowRequest
	GetStartTime() *string
}

type DescribeLiveTopDomainsByFlowRequest struct {
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2018-03-20T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of domain names that you want to retrieve. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	Limit    *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The minimum data granularity is 5 minutes. If you do not specify this parameter, the data of the current month is returned.
	//
	// example:
	//
	// 2018-03-17T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveTopDomainsByFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTopDomainsByFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveTopDomainsByFlowRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveTopDomainsByFlowRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeLiveTopDomainsByFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveTopDomainsByFlowRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveTopDomainsByFlowRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveTopDomainsByFlowRequest) SetEndTime(v string) *DescribeLiveTopDomainsByFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowRequest) SetLimit(v int64) *DescribeLiveTopDomainsByFlowRequest {
	s.Limit = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowRequest) SetOwnerId(v int64) *DescribeLiveTopDomainsByFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowRequest) SetRegionId(v string) *DescribeLiveTopDomainsByFlowRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowRequest) SetStartTime(v string) *DescribeLiveTopDomainsByFlowRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowRequest) Validate() error {
	return dara.Validate(s)
}
