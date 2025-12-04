// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVccFlowInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *ListVccFlowInfosRequest
	GetDirection() *string
	SetFrom(v int64) *ListVccFlowInfosRequest
	GetFrom() *int64
	SetMetricName(v string) *ListVccFlowInfosRequest
	GetMetricName() *string
	SetRegionId(v string) *ListVccFlowInfosRequest
	GetRegionId() *string
	SetTo(v int64) *ListVccFlowInfosRequest
	GetTo() *int64
	SetVccId(v string) *ListVccFlowInfosRequest
	GetVccId() *string
}

type ListVccFlowInfosRequest struct {
	// Direction
	//
	// Valid value:
	//
	// 	- IN: inbound.
	//
	// 	- OUT: the outbound.
	//
	// example:
	//
	// OUT
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The start time. The default value is 5 minutes ago.
	//
	// example:
	//
	// 1667727514000
	From *int64 `json:"From,omitempty" xml:"From,omitempty"`
	// Metric
	//
	// Valid value:
	//
	// 	- totalPacketsRate: Total Packet Rate.
	//
	// 	- dropBytesRate: the of the stream drop rate.
	//
	// 	- dropPacketsRate: Dropped Packet Rate.
	//
	// 	- totalBytesRate: the total streaming rate.
	//
	// 	- passBytesRate: by stream rate.
	//
	// 	- passPacketsRate: by packet rate.
	//
	// example:
	//
	// passBytesRate
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The end time. The default time is the current time.
	//
	// example:
	//
	// 1689749749000
	To *int64 `json:"To,omitempty" xml:"To,omitempty"`
	// Lingjun Connection ID
	//
	// example:
	//
	// vcc-cn-zvp2******
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s ListVccFlowInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVccFlowInfosRequest) GoString() string {
	return s.String()
}

func (s *ListVccFlowInfosRequest) GetDirection() *string {
	return s.Direction
}

func (s *ListVccFlowInfosRequest) GetFrom() *int64 {
	return s.From
}

func (s *ListVccFlowInfosRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *ListVccFlowInfosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVccFlowInfosRequest) GetTo() *int64 {
	return s.To
}

func (s *ListVccFlowInfosRequest) GetVccId() *string {
	return s.VccId
}

func (s *ListVccFlowInfosRequest) SetDirection(v string) *ListVccFlowInfosRequest {
	s.Direction = &v
	return s
}

func (s *ListVccFlowInfosRequest) SetFrom(v int64) *ListVccFlowInfosRequest {
	s.From = &v
	return s
}

func (s *ListVccFlowInfosRequest) SetMetricName(v string) *ListVccFlowInfosRequest {
	s.MetricName = &v
	return s
}

func (s *ListVccFlowInfosRequest) SetRegionId(v string) *ListVccFlowInfosRequest {
	s.RegionId = &v
	return s
}

func (s *ListVccFlowInfosRequest) SetTo(v int64) *ListVccFlowInfosRequest {
	s.To = &v
	return s
}

func (s *ListVccFlowInfosRequest) SetVccId(v string) *ListVccFlowInfosRequest {
	s.VccId = &v
	return s
}

func (s *ListVccFlowInfosRequest) Validate() error {
	return dara.Validate(s)
}
