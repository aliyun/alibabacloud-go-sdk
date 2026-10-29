// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsTopDomainsByFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeVsTopDomainsByFlowRequest
	GetEndTime() *string
	SetLimit(v int64) *DescribeVsTopDomainsByFlowRequest
	GetLimit() *int64
	SetOwnerId(v int64) *DescribeVsTopDomainsByFlowRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVsTopDomainsByFlowRequest
	GetStartTime() *string
}

type DescribeVsTopDomainsByFlowRequest struct {
	// The end time of the data to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.<br>Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2018-12-10T18:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of domains to return. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 3
	Limit   *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The start time of the data to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.<br>The minimum data granularity is 5 minutes.<br>If you do not specify this parameter, data of the current month is queried by default.
	//
	// example:
	//
	// 2021-12-12T10:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsTopDomainsByFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsTopDomainsByFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsTopDomainsByFlowRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsTopDomainsByFlowRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeVsTopDomainsByFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsTopDomainsByFlowRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsTopDomainsByFlowRequest) SetEndTime(v string) *DescribeVsTopDomainsByFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowRequest) SetLimit(v int64) *DescribeVsTopDomainsByFlowRequest {
	s.Limit = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowRequest) SetOwnerId(v int64) *DescribeVsTopDomainsByFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowRequest) SetStartTime(v string) *DescribeVsTopDomainsByFlowRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowRequest) Validate() error {
	return dara.Validate(s)
}
