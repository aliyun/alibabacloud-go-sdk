// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAssetIPTrafficInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeUserAssetIPTrafficInfoResponseBody
	GetEndTime() *int64
	SetInBps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody
	GetInBps() *int64
	SetInPps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody
	GetInPps() *int64
	SetNewConn(v int64) *DescribeUserAssetIPTrafficInfoResponseBody
	GetNewConn() *int64
	SetOutBps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody
	GetOutBps() *int64
	SetOutPps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody
	GetOutPps() *int64
	SetRequestId(v string) *DescribeUserAssetIPTrafficInfoResponseBody
	GetRequestId() *string
	SetSessionCount(v int64) *DescribeUserAssetIPTrafficInfoResponseBody
	GetSessionCount() *int64
	SetStartTime(v int64) *DescribeUserAssetIPTrafficInfoResponseBody
	GetStartTime() *int64
}

type DescribeUserAssetIPTrafficInfoResponseBody struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1656923760
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The network throughput, which indicates the inbound traffic rate. Unit: bit/s.
	//
	// example:
	//
	// 4520
	InBps *int64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// The inbound network throughput, which indicates the number of packets that are sent inbound per second. Unit: packets per second (pps).
	//
	// example:
	//
	// 233
	InPps *int64 `json:"InPps,omitempty" xml:"InPps,omitempty"`
	// The new connection creation rate.
	//
	// example:
	//
	// 43
	NewConn *int64 `json:"NewConn,omitempty" xml:"NewConn,omitempty"`
	// The network throughput, which indicates the outbound traffic rate. Unit: bit/s.
	//
	// example:
	//
	// 4180
	OutBps *int64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// The outbound network throughput, which indicates the number of packets that are sent outbound per second. Unit: pps.
	//
	// example:
	//
	// 224
	OutPps *int64 `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of requests.
	//
	// example:
	//
	// 50
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1656837360
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUserAssetIPTrafficInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAssetIPTrafficInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) GetInBps() *int64 {
	return s.InBps
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) GetInPps() *int64 {
	return s.InPps
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) GetNewConn() *int64 {
	return s.NewConn
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) GetOutBps() *int64 {
	return s.OutBps
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) GetOutPps() *int64 {
	return s.OutPps
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetEndTime(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetInBps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.InBps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetInPps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.InPps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetNewConn(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.NewConn = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetOutBps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.OutBps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetOutPps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.OutPps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetRequestId(v string) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetSessionCount(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.SessionCount = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetStartTime(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
