// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetTrafficTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvgInBps(v int64) *DescribeInternetTrafficTrendResponseBody
	GetAvgInBps() *int64
	SetAvgOutBps(v int64) *DescribeInternetTrafficTrendResponseBody
	GetAvgOutBps() *int64
	SetAvgSession(v int64) *DescribeInternetTrafficTrendResponseBody
	GetAvgSession() *int64
	SetAvgTotalBps(v int64) *DescribeInternetTrafficTrendResponseBody
	GetAvgTotalBps() *int64
	SetDataList(v []*DescribeInternetTrafficTrendResponseBodyDataList) *DescribeInternetTrafficTrendResponseBody
	GetDataList() []*DescribeInternetTrafficTrendResponseBodyDataList
	SetMaxBandwidthTime(v int64) *DescribeInternetTrafficTrendResponseBody
	GetMaxBandwidthTime() *int64
	SetMaxDayExceedBytes(v int64) *DescribeInternetTrafficTrendResponseBody
	GetMaxDayExceedBytes() *int64
	SetMaxInBps(v int64) *DescribeInternetTrafficTrendResponseBody
	GetMaxInBps() *int64
	SetMaxOutBps(v int64) *DescribeInternetTrafficTrendResponseBody
	GetMaxOutBps() *int64
	SetMaxSession(v int64) *DescribeInternetTrafficTrendResponseBody
	GetMaxSession() *int64
	SetMaxTotalBps(v int64) *DescribeInternetTrafficTrendResponseBody
	GetMaxTotalBps() *int64
	SetRequestId(v string) *DescribeInternetTrafficTrendResponseBody
	GetRequestId() *string
	SetTotalBytes(v int64) *DescribeInternetTrafficTrendResponseBody
	GetTotalBytes() *int64
	SetTotalExceedBytes(v int64) *DescribeInternetTrafficTrendResponseBody
	GetTotalExceedBytes() *int64
	SetTotalInBytes(v int64) *DescribeInternetTrafficTrendResponseBody
	GetTotalInBytes() *int64
	SetTotalOutBytes(v int64) *DescribeInternetTrafficTrendResponseBody
	GetTotalOutBytes() *int64
	SetTotalSession(v int64) *DescribeInternetTrafficTrendResponseBody
	GetTotalSession() *int64
}

type DescribeInternetTrafficTrendResponseBody struct {
	// The average network throughput in the inbound direction. Unit: bps.
	//
	// example:
	//
	// 6114152
	AvgInBps *int64 `json:"AvgInBps,omitempty" xml:"AvgInBps,omitempty"`
	// The average network throughput in the outbound direction. Unit: bps.
	//
	// example:
	//
	// 70148993
	AvgOutBps *int64 `json:"AvgOutBps,omitempty" xml:"AvgOutBps,omitempty"`
	// The average number of requests.
	//
	// example:
	//
	// 79013
	AvgSession *int64 `json:"AvgSession,omitempty" xml:"AvgSession,omitempty"`
	// The average network throughput in both outbound and inbound directions. Unit: bps.
	//
	// example:
	//
	// 2306
	AvgTotalBps *int64 `json:"AvgTotalBps,omitempty" xml:"AvgTotalBps,omitempty"`
	// The list of traffic statistics information.
	DataList []*DescribeInternetTrafficTrendResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The timestamp when the peak bandwidth occurred. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1672736400
	MaxBandwidthTime *int64 `json:"MaxBandwidthTime,omitempty" xml:"MaxBandwidthTime,omitempty"`
	// The maximum daily excess traffic.
	//
	// example:
	//
	// 873
	MaxDayExceedBytes *int64 `json:"MaxDayExceedBytes,omitempty" xml:"MaxDayExceedBytes,omitempty"`
	// The peak network throughput in the inbound direction. Unit: bps.
	//
	// example:
	//
	// 10275643
	MaxInBps *int64 `json:"MaxInBps,omitempty" xml:"MaxInBps,omitempty"`
	// The peak network throughput in the outbound direction. Unit: bps.
	//
	// example:
	//
	// 395188
	MaxOutBps *int64 `json:"MaxOutBps,omitempty" xml:"MaxOutBps,omitempty"`
	// The peak number of requests.
	//
	// example:
	//
	// 931641
	MaxSession *int64 `json:"MaxSession,omitempty" xml:"MaxSession,omitempty"`
	// The peak network throughput in both outbound and inbound directions. Unit: bps.
	//
	// example:
	//
	// 89783147
	MaxTotalBps *int64 `json:"MaxTotalBps,omitempty" xml:"MaxTotalBps,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C7E837BE-0379-565E-B7B4-DE595C8D337C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total network traffic in both outbound and inbound directions. Unit: bytes.
	//
	// example:
	//
	// 963227674958
	TotalBytes *int64 `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// The total excess traffic.
	//
	// example:
	//
	// 4243873
	TotalExceedBytes *int64 `json:"TotalExceedBytes,omitempty" xml:"TotalExceedBytes,omitempty"`
	// The total network traffic in the inbound direction. Unit: bytes.
	//
	// example:
	//
	// 41536824243873
	TotalInBytes *int64 `json:"TotalInBytes,omitempty" xml:"TotalInBytes,omitempty"`
	// The total network traffic in the outbound direction. Unit: bytes.
	//
	// example:
	//
	// 2660894567178
	TotalOutBytes *int64 `json:"TotalOutBytes,omitempty" xml:"TotalOutBytes,omitempty"`
	// The total number of requests.
	//
	// example:
	//
	// 1061449
	TotalSession *int64 `json:"TotalSession,omitempty" xml:"TotalSession,omitempty"`
}

func (s DescribeInternetTrafficTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetTrafficTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInternetTrafficTrendResponseBody) GetAvgInBps() *int64 {
	return s.AvgInBps
}

func (s *DescribeInternetTrafficTrendResponseBody) GetAvgOutBps() *int64 {
	return s.AvgOutBps
}

func (s *DescribeInternetTrafficTrendResponseBody) GetAvgSession() *int64 {
	return s.AvgSession
}

func (s *DescribeInternetTrafficTrendResponseBody) GetAvgTotalBps() *int64 {
	return s.AvgTotalBps
}

func (s *DescribeInternetTrafficTrendResponseBody) GetDataList() []*DescribeInternetTrafficTrendResponseBodyDataList {
	return s.DataList
}

func (s *DescribeInternetTrafficTrendResponseBody) GetMaxBandwidthTime() *int64 {
	return s.MaxBandwidthTime
}

func (s *DescribeInternetTrafficTrendResponseBody) GetMaxDayExceedBytes() *int64 {
	return s.MaxDayExceedBytes
}

func (s *DescribeInternetTrafficTrendResponseBody) GetMaxInBps() *int64 {
	return s.MaxInBps
}

func (s *DescribeInternetTrafficTrendResponseBody) GetMaxOutBps() *int64 {
	return s.MaxOutBps
}

func (s *DescribeInternetTrafficTrendResponseBody) GetMaxSession() *int64 {
	return s.MaxSession
}

func (s *DescribeInternetTrafficTrendResponseBody) GetMaxTotalBps() *int64 {
	return s.MaxTotalBps
}

func (s *DescribeInternetTrafficTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInternetTrafficTrendResponseBody) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *DescribeInternetTrafficTrendResponseBody) GetTotalExceedBytes() *int64 {
	return s.TotalExceedBytes
}

func (s *DescribeInternetTrafficTrendResponseBody) GetTotalInBytes() *int64 {
	return s.TotalInBytes
}

func (s *DescribeInternetTrafficTrendResponseBody) GetTotalOutBytes() *int64 {
	return s.TotalOutBytes
}

func (s *DescribeInternetTrafficTrendResponseBody) GetTotalSession() *int64 {
	return s.TotalSession
}

func (s *DescribeInternetTrafficTrendResponseBody) SetAvgInBps(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.AvgInBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetAvgOutBps(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.AvgOutBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetAvgSession(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.AvgSession = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetAvgTotalBps(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.AvgTotalBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetDataList(v []*DescribeInternetTrafficTrendResponseBodyDataList) *DescribeInternetTrafficTrendResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetMaxBandwidthTime(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.MaxBandwidthTime = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetMaxDayExceedBytes(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.MaxDayExceedBytes = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetMaxInBps(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.MaxInBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetMaxOutBps(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.MaxOutBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetMaxSession(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.MaxSession = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetMaxTotalBps(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.MaxTotalBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetRequestId(v string) *DescribeInternetTrafficTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetTotalBytes(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.TotalBytes = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetTotalExceedBytes(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.TotalExceedBytes = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetTotalInBytes(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.TotalInBytes = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetTotalOutBytes(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.TotalOutBytes = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetTotalSession(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.TotalSession = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) Validate() error {
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

type DescribeInternetTrafficTrendResponseBodyDataList struct {
	// The network throughput in the inbound direction. Unit: bps.
	//
	// example:
	//
	// 187
	InBps *int64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// The total network traffic in the inbound direction. Unit: bytes.
	//
	// example:
	//
	// 235
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// The packet rate in the inbound direction. Unit: pps.
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
	// The network throughput in the outbound direction. Unit: bps.
	//
	// example:
	//
	// 45
	OutBps *int64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// The total network traffic in the outbound direction. Unit: bytes.
	//
	// example:
	//
	// 1123
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The packet rate in the outbound direction. Unit: pps.
	//
	// example:
	//
	// 2
	OutPps *int64 `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
	// The number of requests.
	//
	// example:
	//
	// 27
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// The time when the traffic occurred. The value is a UNIX timestamp in seconds.
	//
	// If the data for this time point is not fully processed, the values of other fields are -1.
	//
	// example:
	//
	// 1659405600
	Time *int32 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The total network throughput in both outbound and inbound directions. Unit: bps.
	//
	// example:
	//
	// 232
	TotalBps *int64 `json:"TotalBps,omitempty" xml:"TotalBps,omitempty"`
}

func (s DescribeInternetTrafficTrendResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetTrafficTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) GetInBps() *int64 {
	return s.InBps
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) GetInPps() *int64 {
	return s.InPps
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) GetNewConn() *int64 {
	return s.NewConn
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) GetOutBps() *int64 {
	return s.OutBps
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) GetOutPps() *int64 {
	return s.OutPps
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) GetTime() *int32 {
	return s.Time
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) GetTotalBps() *int64 {
	return s.TotalBps
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetInBps(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.InBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetInBytes(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.InBytes = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetInPps(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.InPps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetNewConn(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.NewConn = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetOutBps(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.OutBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetOutBytes(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.OutBytes = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetOutPps(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.OutPps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetSessionCount(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.SessionCount = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetTime(v int32) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.Time = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetTotalBps(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.TotalBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
