// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticQpsRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeElasticQpsRecordRequest
	GetEndTime() *int64
	SetIp(v string) *DescribeElasticQpsRecordRequest
	GetIp() *string
	SetStartTime(v int64) *DescribeElasticQpsRecordRequest
	GetStartTime() *int64
}

type DescribeElasticQpsRecordRequest struct {
	// The end of the time range to query. The value is a timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1688140799999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address of the Anti-DDoS Proxy instance to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 203.107.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The beginning of the time range to query. The value is a timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684252800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeElasticQpsRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticQpsRecordRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticQpsRecordRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeElasticQpsRecordRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeElasticQpsRecordRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeElasticQpsRecordRequest) SetEndTime(v int64) *DescribeElasticQpsRecordRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeElasticQpsRecordRequest) SetIp(v string) *DescribeElasticQpsRecordRequest {
	s.Ip = &v
	return s
}

func (s *DescribeElasticQpsRecordRequest) SetStartTime(v int64) *DescribeElasticQpsRecordRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeElasticQpsRecordRequest) Validate() error {
	return dara.Validate(s)
}
