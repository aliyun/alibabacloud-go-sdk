// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpTrafficRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEip(v string) *DescribeIpTrafficRequest
	GetEip() *string
	SetEndTime(v int64) *DescribeIpTrafficRequest
	GetEndTime() *int64
	SetInterval(v int32) *DescribeIpTrafficRequest
	GetInterval() *int32
	SetPort(v int32) *DescribeIpTrafficRequest
	GetPort() *int32
	SetQueryProtocol(v string) *DescribeIpTrafficRequest
	GetQueryProtocol() *string
	SetResourceGroupId(v string) *DescribeIpTrafficRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeIpTrafficRequest
	GetSourceIp() *string
	SetStartTime(v int64) *DescribeIpTrafficRequest
	GetStartTime() *int64
}

type DescribeIpTrafficRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1.1
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1536734120
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 233
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// http
	QueryProtocol *string `json:"QueryProtocol,omitempty" xml:"QueryProtocol,omitempty"`
	// example:
	//
	// xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1536734112
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeIpTrafficRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpTrafficRequest) GetEip() *string {
	return s.Eip
}

func (s *DescribeIpTrafficRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeIpTrafficRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeIpTrafficRequest) GetPort() *int32 {
	return s.Port
}

func (s *DescribeIpTrafficRequest) GetQueryProtocol() *string {
	return s.QueryProtocol
}

func (s *DescribeIpTrafficRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeIpTrafficRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeIpTrafficRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeIpTrafficRequest) SetEip(v string) *DescribeIpTrafficRequest {
	s.Eip = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetEndTime(v int64) *DescribeIpTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetInterval(v int32) *DescribeIpTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetPort(v int32) *DescribeIpTrafficRequest {
	s.Port = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetQueryProtocol(v string) *DescribeIpTrafficRequest {
	s.QueryProtocol = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetResourceGroupId(v string) *DescribeIpTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetSourceIp(v string) *DescribeIpTrafficRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetStartTime(v int64) *DescribeIpTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeIpTrafficRequest) Validate() error {
	return dara.Validate(s)
}
