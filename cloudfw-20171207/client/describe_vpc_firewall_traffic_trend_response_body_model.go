// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallTrafficTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvgInBps(v int64) *DescribeVpcFirewallTrafficTrendResponseBody
	GetAvgInBps() *int64
	SetAvgOutBps(v int64) *DescribeVpcFirewallTrafficTrendResponseBody
	GetAvgOutBps() *int64
	SetAvgSession(v int64) *DescribeVpcFirewallTrafficTrendResponseBody
	GetAvgSession() *int64
	SetAvgTotalBps(v int64) *DescribeVpcFirewallTrafficTrendResponseBody
	GetAvgTotalBps() *int64
	SetDataList(v []*DescribeVpcFirewallTrafficTrendResponseBodyDataList) *DescribeVpcFirewallTrafficTrendResponseBody
	GetDataList() []*DescribeVpcFirewallTrafficTrendResponseBodyDataList
	SetMaxBandwidthTime(v int64) *DescribeVpcFirewallTrafficTrendResponseBody
	GetMaxBandwidthTime() *int64
	SetMaxInBps(v int64) *DescribeVpcFirewallTrafficTrendResponseBody
	GetMaxInBps() *int64
	SetMaxOutBps(v int64) *DescribeVpcFirewallTrafficTrendResponseBody
	GetMaxOutBps() *int64
	SetMaxSession(v int64) *DescribeVpcFirewallTrafficTrendResponseBody
	GetMaxSession() *int64
	SetMaxTotalBps(v int64) *DescribeVpcFirewallTrafficTrendResponseBody
	GetMaxTotalBps() *int64
	SetRequestId(v string) *DescribeVpcFirewallTrafficTrendResponseBody
	GetRequestId() *string
	SetTotalBytes(v int64) *DescribeVpcFirewallTrafficTrendResponseBody
	GetTotalBytes() *int64
	SetTotalInBytes(v int64) *DescribeVpcFirewallTrafficTrendResponseBody
	GetTotalInBytes() *int64
	SetTotalOutBytes(v int64) *DescribeVpcFirewallTrafficTrendResponseBody
	GetTotalOutBytes() *int64
	SetTotalSession(v int64) *DescribeVpcFirewallTrafficTrendResponseBody
	GetTotalSession() *int64
}

type DescribeVpcFirewallTrafficTrendResponseBody struct {
	// The average inbound network throughput. Unit: bit/s.
	//
	// example:
	//
	// 1264110
	AvgInBps *int64 `json:"AvgInBps,omitempty" xml:"AvgInBps,omitempty"`
	// The average outbound network throughput. Unit: bit/s.
	//
	// example:
	//
	// 68915
	AvgOutBps *int64 `json:"AvgOutBps,omitempty" xml:"AvgOutBps,omitempty"`
	// The average number of requests.
	//
	// example:
	//
	// 1995
	AvgSession *int64 `json:"AvgSession,omitempty" xml:"AvgSession,omitempty"`
	// The average total network throughput in both the outbound and inbound directions. Unit: bit/s.
	//
	// example:
	//
	// 34291
	AvgTotalBps *int64 `json:"AvgTotalBps,omitempty" xml:"AvgTotalBps,omitempty"`
	// The data list.
	DataList []*DescribeVpcFirewallTrafficTrendResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The timestamp when the peak bandwidth occurred. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1768008060
	MaxBandwidthTime *int64 `json:"MaxBandwidthTime,omitempty" xml:"MaxBandwidthTime,omitempty"`
	// The peak inbound network throughput. Unit: bit/s.
	//
	// example:
	//
	// 1436
	MaxInBps *int64 `json:"MaxInBps,omitempty" xml:"MaxInBps,omitempty"`
	// The peak outbound network throughput. Unit: bit/s.
	//
	// example:
	//
	// 2128
	MaxOutBps *int64 `json:"MaxOutBps,omitempty" xml:"MaxOutBps,omitempty"`
	// The peak number of requests.
	//
	// example:
	//
	// 2003
	MaxSession *int64 `json:"MaxSession,omitempty" xml:"MaxSession,omitempty"`
	// The peak total network throughput in both the outbound and inbound directions. Unit: bit/s.
	//
	// example:
	//
	// 61947852
	MaxTotalBps *int64 `json:"MaxTotalBps,omitempty" xml:"MaxTotalBps,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 45F8B9E6-8583-56B3-A127-1B421C9E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total traffic. Unit: bytes.
	//
	// example:
	//
	// 141688156232
	TotalBytes *int64 `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// The total inbound network throughput. Unit: bytes.
	//
	// example:
	//
	// 2659635037
	TotalInBytes *int64 `json:"TotalInBytes,omitempty" xml:"TotalInBytes,omitempty"`
	// The total outbound network throughput. Unit: bytes.
	//
	// example:
	//
	// 399762701
	TotalOutBytes *int64 `json:"TotalOutBytes,omitempty" xml:"TotalOutBytes,omitempty"`
	// The total number of requests.
	//
	// example:
	//
	// 1078757
	TotalSession *int64 `json:"TotalSession,omitempty" xml:"TotalSession,omitempty"`
}

func (s DescribeVpcFirewallTrafficTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallTrafficTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetAvgInBps() *int64 {
	return s.AvgInBps
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetAvgOutBps() *int64 {
	return s.AvgOutBps
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetAvgSession() *int64 {
	return s.AvgSession
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetAvgTotalBps() *int64 {
	return s.AvgTotalBps
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetDataList() []*DescribeVpcFirewallTrafficTrendResponseBodyDataList {
	return s.DataList
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetMaxBandwidthTime() *int64 {
	return s.MaxBandwidthTime
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetMaxInBps() *int64 {
	return s.MaxInBps
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetMaxOutBps() *int64 {
	return s.MaxOutBps
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetMaxSession() *int64 {
	return s.MaxSession
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetMaxTotalBps() *int64 {
	return s.MaxTotalBps
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetTotalInBytes() *int64 {
	return s.TotalInBytes
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetTotalOutBytes() *int64 {
	return s.TotalOutBytes
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) GetTotalSession() *int64 {
	return s.TotalSession
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetAvgInBps(v int64) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.AvgInBps = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetAvgOutBps(v int64) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.AvgOutBps = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetAvgSession(v int64) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.AvgSession = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetAvgTotalBps(v int64) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.AvgTotalBps = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetDataList(v []*DescribeVpcFirewallTrafficTrendResponseBodyDataList) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetMaxBandwidthTime(v int64) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.MaxBandwidthTime = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetMaxInBps(v int64) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.MaxInBps = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetMaxOutBps(v int64) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.MaxOutBps = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetMaxSession(v int64) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.MaxSession = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetMaxTotalBps(v int64) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.MaxTotalBps = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetRequestId(v string) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetTotalBytes(v int64) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.TotalBytes = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetTotalInBytes(v int64) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.TotalInBytes = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetTotalOutBytes(v int64) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.TotalOutBytes = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) SetTotalSession(v int64) *DescribeVpcFirewallTrafficTrendResponseBody {
	s.TotalSession = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcFirewallTrafficTrendResponseBodyDataList struct {
	// The inbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 187
	InBps *int64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// The inbound traffic. Unit: bytes.
	//
	// example:
	//
	// 32
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// The inbound packet forwarding rate. Unit: pps.
	//
	// example:
	//
	// 2
	InPps *int64 `json:"InPps,omitempty" xml:"InPps,omitempty"`
	// The number of new connections.
	//
	// example:
	//
	// 27
	NewConn *int64 `json:"NewConn,omitempty" xml:"NewConn,omitempty"`
	// The outbound traffic. Unit: bytes.
	//
	// example:
	//
	// 45
	OutBps *int64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// The total outbound network throughput. Unit: bytes.
	//
	// example:
	//
	// 230
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The outbound packet forwarding rate. Unit: pps.
	//
	// example:
	//
	// 2
	OutPps *int64 `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
	// The number of sessions.
	//
	// example:
	//
	// 27
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// The time when the traffic occurred. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1758470400
	Time *int32 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeVpcFirewallTrafficTrendResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallTrafficTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) GetInBps() *int64 {
	return s.InBps
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) GetInPps() *int64 {
	return s.InPps
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) GetNewConn() *int64 {
	return s.NewConn
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) GetOutBps() *int64 {
	return s.OutBps
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) GetOutPps() *int64 {
	return s.OutPps
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) GetTime() *int32 {
	return s.Time
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) SetInBps(v int64) *DescribeVpcFirewallTrafficTrendResponseBodyDataList {
	s.InBps = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) SetInBytes(v int64) *DescribeVpcFirewallTrafficTrendResponseBodyDataList {
	s.InBytes = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) SetInPps(v int64) *DescribeVpcFirewallTrafficTrendResponseBodyDataList {
	s.InPps = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) SetNewConn(v int64) *DescribeVpcFirewallTrafficTrendResponseBodyDataList {
	s.NewConn = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) SetOutBps(v int64) *DescribeVpcFirewallTrafficTrendResponseBodyDataList {
	s.OutBps = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) SetOutBytes(v int64) *DescribeVpcFirewallTrafficTrendResponseBodyDataList {
	s.OutBytes = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) SetOutPps(v int64) *DescribeVpcFirewallTrafficTrendResponseBodyDataList {
	s.OutPps = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) SetSessionCount(v int64) *DescribeVpcFirewallTrafficTrendResponseBodyDataList {
	s.SessionCount = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) SetTime(v int32) *DescribeVpcFirewallTrafficTrendResponseBodyDataList {
	s.Time = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
