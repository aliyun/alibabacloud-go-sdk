// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserStreamMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveUserStreamMetricDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveUserStreamMetricDataResponseBody
	GetEndTime() *string
	SetPageNumber(v int64) *DescribeLiveUserStreamMetricDataResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeLiveUserStreamMetricDataResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeLiveUserStreamMetricDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveUserStreamMetricDataResponseBody
	GetStartTime() *string
	SetStreamDetailData(v []*DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) *DescribeLiveUserStreamMetricDataResponseBody
	GetStreamDetailData() []*DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData
	SetTotalCount(v int64) *DescribeLiveUserStreamMetricDataResponseBody
	GetTotalCount() *int64
}

type DescribeLiveUserStreamMetricDataResponseBody struct {
	// example:
	//
	// test.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// YYYY-MM-DDThh:mm:ssZ
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime        *string                                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamDetailData []*DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData `json:"StreamDetailData,omitempty" xml:"StreamDetailData,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLiveUserStreamMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserStreamMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) GetStreamDetailData() []*DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	return s.StreamDetailData
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) SetDomainName(v string) *DescribeLiveUserStreamMetricDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) SetEndTime(v string) *DescribeLiveUserStreamMetricDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) SetPageNumber(v int64) *DescribeLiveUserStreamMetricDataResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) SetPageSize(v int64) *DescribeLiveUserStreamMetricDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) SetRequestId(v string) *DescribeLiveUserStreamMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) SetStartTime(v string) *DescribeLiveUserStreamMetricDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) SetStreamDetailData(v []*DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) *DescribeLiveUserStreamMetricDataResponseBody {
	s.StreamDetailData = v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) SetTotalCount(v int64) *DescribeLiveUserStreamMetricDataResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData struct {
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 423304182.66
	Bps *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// example:
	//
	// 423304182
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 454
	FlvBps *float32 `json:"FlvBps,omitempty" xml:"FlvBps,omitempty"`
	// example:
	//
	// 32
	FlvCount *int64 `json:"FlvCount,omitempty" xml:"FlvCount,omitempty"`
	// example:
	//
	// 1254
	FlvTraffic *float64 `json:"FlvTraffic,omitempty" xml:"FlvTraffic,omitempty"`
	// example:
	//
	// 4456
	HlsBps *float32 `json:"HlsBps,omitempty" xml:"HlsBps,omitempty"`
	// example:
	//
	// 56
	HlsCount *int64 `json:"HlsCount,omitempty" xml:"HlsCount,omitempty"`
	// example:
	//
	// 568
	HlsTraffic *float64 `json:"HlsTraffic,omitempty" xml:"HlsTraffic,omitempty"`
	// example:
	//
	// 450
	NewConns *int64 `json:"NewConns,omitempty" xml:"NewConns,omitempty"`
	// example:
	//
	// 6845
	P2pBps *float32 `json:"P2pBps,omitempty" xml:"P2pBps,omitempty"`
	// example:
	//
	// 78
	P2pCount *int64 `json:"P2pCount,omitempty" xml:"P2pCount,omitempty"`
	// example:
	//
	// 4102
	P2pTraffic *float64 `json:"P2pTraffic,omitempty" xml:"P2pTraffic,omitempty"`
	// example:
	//
	// 3323
	RtmpBps *float32 `json:"RtmpBps,omitempty" xml:"RtmpBps,omitempty"`
	// example:
	//
	// 63
	RtmpCount *int64 `json:"RtmpCount,omitempty" xml:"RtmpCount,omitempty"`
	// example:
	//
	// 5568
	RtmpTraffic *float64 `json:"RtmpTraffic,omitempty" xml:"RtmpTraffic,omitempty"`
	// example:
	//
	// 2361
	RtsBps *float64 `json:"RtsBps,omitempty" xml:"RtsBps,omitempty"`
	// example:
	//
	// 89
	RtsCount *int64 `json:"RtsCount,omitempty" xml:"RtsCount,omitempty"`
	// example:
	//
	// 2322
	RtsTraffic *float64 `json:"RtsTraffic,omitempty" xml:"RtsTraffic,omitempty"`
	// example:
	//
	// test.flv
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// example:
	//
	// 2015-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// example:
	//
	// 423304182
	Traffic *float64 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetBps() *float32 {
	return s.Bps
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetCount() *int64 {
	return s.Count
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetFlvBps() *float32 {
	return s.FlvBps
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetFlvCount() *int64 {
	return s.FlvCount
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetFlvTraffic() *float64 {
	return s.FlvTraffic
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetHlsBps() *float32 {
	return s.HlsBps
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetHlsCount() *int64 {
	return s.HlsCount
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetHlsTraffic() *float64 {
	return s.HlsTraffic
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetNewConns() *int64 {
	return s.NewConns
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetP2pBps() *float32 {
	return s.P2pBps
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetP2pCount() *int64 {
	return s.P2pCount
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetP2pTraffic() *float64 {
	return s.P2pTraffic
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetRtmpBps() *float32 {
	return s.RtmpBps
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetRtmpCount() *int64 {
	return s.RtmpCount
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetRtmpTraffic() *float64 {
	return s.RtmpTraffic
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetRtsBps() *float64 {
	return s.RtsBps
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetRtsCount() *int64 {
	return s.RtsCount
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetRtsTraffic() *float64 {
	return s.RtsTraffic
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) GetTraffic() *float64 {
	return s.Traffic
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetAppName(v string) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.AppName = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetBps(v float32) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.Bps = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetCount(v int64) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.Count = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetFlvBps(v float32) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.FlvBps = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetFlvCount(v int64) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.FlvCount = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetFlvTraffic(v float64) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.FlvTraffic = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetHlsBps(v float32) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.HlsBps = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetHlsCount(v int64) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.HlsCount = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetHlsTraffic(v float64) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.HlsTraffic = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetNewConns(v int64) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.NewConns = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetP2pBps(v float32) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.P2pBps = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetP2pCount(v int64) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.P2pCount = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetP2pTraffic(v float64) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.P2pTraffic = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetRtmpBps(v float32) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.RtmpBps = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetRtmpCount(v int64) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.RtmpCount = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetRtmpTraffic(v float64) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.RtmpTraffic = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetRtsBps(v float64) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.RtsBps = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetRtsCount(v int64) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.RtsCount = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetRtsTraffic(v float64) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.RtsTraffic = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetStreamName(v string) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetTimeStamp(v string) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) SetTraffic(v float64) *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData {
	s.Traffic = &v
	return s
}

func (s *DescribeLiveUserStreamMetricDataResponseBodyStreamDetailData) Validate() error {
	return dara.Validate(s)
}
