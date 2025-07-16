// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTopDomainsByFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeTopDomainsByFlowRequest
	GetEndTime() *string
	SetLimit(v int64) *DescribeTopDomainsByFlowRequest
	GetLimit() *int64
	SetStartTime(v string) *DescribeTopDomainsByFlowRequest
	GetStartTime() *string
}

type DescribeTopDomainsByFlowRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2019-12-23T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of domain names to query. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The value of StartTime must be in UTC. For example, if the local time is 00:00 on June 1, 2021, set StartTime to 2021-05-31T16:00:00Z.
	//
	// example:
	//
	// 2019-12-22T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTopDomainsByFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopDomainsByFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTopDomainsByFlowRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeTopDomainsByFlowRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTopDomainsByFlowRequest) SetEndTime(v string) *DescribeTopDomainsByFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTopDomainsByFlowRequest) SetLimit(v int64) *DescribeTopDomainsByFlowRequest {
	s.Limit = &v
	return s
}

func (s *DescribeTopDomainsByFlowRequest) SetStartTime(v string) *DescribeTopDomainsByFlowRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTopDomainsByFlowRequest) Validate() error {
	return dara.Validate(s)
}
