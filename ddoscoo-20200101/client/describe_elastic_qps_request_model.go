// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticQpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeElasticQpsRequest
	GetEndTime() *int64
	SetInterval(v string) *DescribeElasticQpsRequest
	GetInterval() *string
	SetIp(v string) *DescribeElasticQpsRequest
	GetIp() *string
	SetRegion(v string) *DescribeElasticQpsRequest
	GetRegion() *string
	SetStartTime(v int64) *DescribeElasticQpsRequest
	GetStartTime() *int64
}

type DescribeElasticQpsRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// >  This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684339200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The sampling interval. Unit: seconds. The value must be a multiple of 60. Default value: 60. Unit: seconds. The query result varies depending on the sampling interval.
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The IP address of the Anti-DDoS Proxy instance to query.
	//
	// example:
	//
	// 203.107.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The type of the service. Valid values:
	//
	// 	- **cn**: Anti-DDoS Proxy (Chinese Mainland)
	//
	// 	- **cn-hongkong**: Anti-DDoS Proxy (Outside Chinese Mainland)
	//
	// This parameter is required.
	//
	// example:
	//
	// cn
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// >  This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684252800
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeElasticQpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticQpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticQpsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeElasticQpsRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeElasticQpsRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeElasticQpsRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeElasticQpsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeElasticQpsRequest) SetEndTime(v int64) *DescribeElasticQpsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeElasticQpsRequest) SetInterval(v string) *DescribeElasticQpsRequest {
	s.Interval = &v
	return s
}

func (s *DescribeElasticQpsRequest) SetIp(v string) *DescribeElasticQpsRequest {
	s.Ip = &v
	return s
}

func (s *DescribeElasticQpsRequest) SetRegion(v string) *DescribeElasticQpsRequest {
	s.Region = &v
	return s
}

func (s *DescribeElasticQpsRequest) SetStartTime(v int64) *DescribeElasticQpsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeElasticQpsRequest) Validate() error {
	return dara.Validate(s)
}
