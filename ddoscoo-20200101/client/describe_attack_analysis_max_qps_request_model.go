// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttackAnalysisMaxQpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeAttackAnalysisMaxQpsRequest
	GetEndTime() *int64
	SetIp(v string) *DescribeAttackAnalysisMaxQpsRequest
	GetIp() *string
	SetStartTime(v int64) *DescribeAttackAnalysisMaxQpsRequest
	GetStartTime() *int64
}

type DescribeAttackAnalysisMaxQpsRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1619798400
	EndTime *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Ip      *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1622476799
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAttackAnalysisMaxQpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttackAnalysisMaxQpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAttackAnalysisMaxQpsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAttackAnalysisMaxQpsRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeAttackAnalysisMaxQpsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAttackAnalysisMaxQpsRequest) SetEndTime(v int64) *DescribeAttackAnalysisMaxQpsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAttackAnalysisMaxQpsRequest) SetIp(v string) *DescribeAttackAnalysisMaxQpsRequest {
	s.Ip = &v
	return s
}

func (s *DescribeAttackAnalysisMaxQpsRequest) SetStartTime(v int64) *DescribeAttackAnalysisMaxQpsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAttackAnalysisMaxQpsRequest) Validate() error {
	return dara.Validate(s)
}
