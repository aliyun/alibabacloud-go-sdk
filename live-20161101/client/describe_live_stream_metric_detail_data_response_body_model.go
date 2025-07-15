// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamMetricDetailDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveStreamMetricDetailDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamMetricDetailDataResponseBody
	GetEndTime() *string
	SetNextPageToken(v string) *DescribeLiveStreamMetricDetailDataResponseBody
	GetNextPageToken() *string
	SetPageSize(v int32) *DescribeLiveStreamMetricDetailDataResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLiveStreamMetricDetailDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveStreamMetricDetailDataResponseBody
	GetStartTime() *string
	SetStreamDetailData(v *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailData) *DescribeLiveStreamMetricDetailDataResponseBody
	GetStreamDetailData() *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailData
}

type DescribeLiveStreamMetricDetailDataResponseBody struct {
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2015-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The token that determines the start point of the next query. This parameter is returned if more data results are available.
	//
	// example:
	//
	// UjsM9x3aVcJi9a0-ArwJUTTC67C***37C0=
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The number of rows returned.
	//
	// example:
	//
	// 5000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B955107D-E658-4E77-B913-E0AC3D31693E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which data was queried.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The data array returned.
	StreamDetailData *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailData `json:"StreamDetailData,omitempty" xml:"StreamDetailData,omitempty" type:"Struct"`
}

func (s DescribeLiveStreamMetricDetailDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMetricDetailDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) GetStreamDetailData() *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailData {
	return s.StreamDetailData
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) SetDomainName(v string) *DescribeLiveStreamMetricDetailDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) SetEndTime(v string) *DescribeLiveStreamMetricDetailDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) SetNextPageToken(v string) *DescribeLiveStreamMetricDetailDataResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) SetPageSize(v int32) *DescribeLiveStreamMetricDetailDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) SetRequestId(v string) *DescribeLiveStreamMetricDetailDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) SetStartTime(v string) *DescribeLiveStreamMetricDetailDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) SetStreamDetailData(v *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailData) *DescribeLiveStreamMetricDetailDataResponseBody {
	s.StreamDetailData = v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailData struct {
	StreamData []*DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData `json:"StreamData,omitempty" xml:"StreamData,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailData) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailData) GetStreamData() []*DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	return s.StreamData
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailData) SetStreamData(v []*DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailData {
	s.StreamData = v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailData) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData struct {
	// The name of the application.
	//
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The total bandwidth consumed by the stream per minute. Unit: bit/s.
	//
	// example:
	//
	// 423304182.66
	Bps *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The total number of online viewers for the stream per minute.
	//
	// example:
	//
	// 423304182
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The bandwidth over the Flash Video (FLV) protocol. Unit: bit/s.
	//
	// >  This parameter is not returned if no traffic is generated over the protocol.
	//
	// example:
	//
	// 454
	FlvBps *float32 `json:"FlvBps,omitempty" xml:"FlvBps,omitempty"`
	// The number of online viewers over the FLV protocol.
	//
	// >  This parameter is not returned if no traffic is generated over the protocol.
	//
	// example:
	//
	// 32
	FlvCount *int64 `json:"FlvCount,omitempty" xml:"FlvCount,omitempty"`
	// The amount of traffic over the FLV protocol. Unit: bytes.
	//
	// >  This parameter is not returned if no traffic is generated over the protocol.
	//
	// example:
	//
	// 1254
	FlvTraffic *int64 `json:"FlvTraffic,omitempty" xml:"FlvTraffic,omitempty"`
	// The bandwidth over the HTTP Live Streaming (HLS) protocol. Unit: bit/s.
	//
	// >  This parameter is not returned if no traffic is generated over the protocol.
	//
	// example:
	//
	// 4456
	HlsBps *float32 `json:"HlsBps,omitempty" xml:"HlsBps,omitempty"`
	// The number of online viewers over the HLS protocol.
	//
	// >  Currently, this parameter is not supported.
	//
	// example:
	//
	// 56
	HlsCount *int64 `json:"HlsCount,omitempty" xml:"HlsCount,omitempty"`
	// The amount of traffic over the HLS protocol. Unit: bytes.
	//
	// >  This parameter is not returned if no traffic is generated over the protocol.
	//
	// example:
	//
	// 568
	HlsTraffic *int64 `json:"HlsTraffic,omitempty" xml:"HlsTraffic,omitempty"`
	// Number of new connections established per minute.
	//
	// example:
	//
	// 450
	NewConns *string `json:"NewConns,omitempty" xml:"NewConns,omitempty"`
	// The bandwidth over the P2P protocol. Unit: bit/s.
	//
	// >  This parameter is not returned if no traffic is generated over the protocol.
	//
	// example:
	//
	// 6845
	P2pBps *float32 `json:"P2pBps,omitempty" xml:"P2pBps,omitempty"`
	// The number of online viewers over the P2P protocol.
	//
	// >  This parameter is not returned if no traffic is generated over the protocol.
	//
	// example:
	//
	// 78
	P2pCount *int64 `json:"P2pCount,omitempty" xml:"P2pCount,omitempty"`
	// The amount of traffic over the peer-to-peer (P2P) protocol. Unit: bytes.
	//
	// >  This parameter is not returned if no traffic is generated over the protocol.
	//
	// example:
	//
	// 4102
	P2pTraffic *int64 `json:"P2pTraffic,omitempty" xml:"P2pTraffic,omitempty"`
	// The bandwidth over the Real-Time Messaging Protocol (RTMP) protocol. Unit: bit/s.
	//
	// >  This parameter is not returned if no traffic is generated over the protocol.
	//
	// example:
	//
	// 3323
	RtmpBps *float32 `json:"RtmpBps,omitempty" xml:"RtmpBps,omitempty"`
	// The number of online viewers over the RTMP protocol.
	//
	// >  This parameter is not returned if no traffic is generated over the protocol.
	//
	// example:
	//
	// 63
	RtmpCount *int64 `json:"RtmpCount,omitempty" xml:"RtmpCount,omitempty"`
	// The amount of traffic over the RTMP protocol. Unit: bytes.
	//
	// >  This parameter is not returned if no traffic is generated over the protocol.
	//
	// example:
	//
	// 5568
	RtmpTraffic *int64 `json:"RtmpTraffic,omitempty" xml:"RtmpTraffic,omitempty"`
	// The bandwidth over the RTS protocol. Unit: bit/s.
	//
	// >  This parameter is not returned if no traffic is generated over the protocol.
	//
	// example:
	//
	// 2361
	RtsBps *float32 `json:"RtsBps,omitempty" xml:"RtsBps,omitempty"`
	// The number of online viewers over the Real-Time Streaming (RTS) protocol.
	//
	// >  This parameter is not returned if no traffic is generated over the protocol.
	//
	// example:
	//
	// 89
	RtsCount *int64 `json:"RtsCount,omitempty" xml:"RtsCount,omitempty"`
	// The amount of traffic over the RTS protocol. Unit: bytes.
	//
	// >  This parameter is not returned if no traffic is generated over the protocol.
	//
	// example:
	//
	// 2322
	RtsTraffic *int64 `json:"RtsTraffic,omitempty" xml:"RtsTraffic,omitempty"`
	// The name of the stream.
	//
	// example:
	//
	// test.flv
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total amount of traffic consumed by the stream per minute. Unit: bytes.
	//
	// example:
	//
	// 423304182
	Traffic *int64 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetBps() *float32 {
	return s.Bps
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetCount() *int64 {
	return s.Count
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetFlvBps() *float32 {
	return s.FlvBps
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetFlvCount() *int64 {
	return s.FlvCount
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetFlvTraffic() *int64 {
	return s.FlvTraffic
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetHlsBps() *float32 {
	return s.HlsBps
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetHlsCount() *int64 {
	return s.HlsCount
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetHlsTraffic() *int64 {
	return s.HlsTraffic
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetNewConns() *string {
	return s.NewConns
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetP2pBps() *float32 {
	return s.P2pBps
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetP2pCount() *int64 {
	return s.P2pCount
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetP2pTraffic() *int64 {
	return s.P2pTraffic
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetRtmpBps() *float32 {
	return s.RtmpBps
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetRtmpCount() *int64 {
	return s.RtmpCount
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetRtmpTraffic() *int64 {
	return s.RtmpTraffic
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetRtsBps() *float32 {
	return s.RtsBps
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetRtsCount() *int64 {
	return s.RtsCount
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetRtsTraffic() *int64 {
	return s.RtsTraffic
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) GetTraffic() *int64 {
	return s.Traffic
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetAppName(v string) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetBps(v float32) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.Bps = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetCount(v int64) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.Count = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetFlvBps(v float32) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.FlvBps = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetFlvCount(v int64) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.FlvCount = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetFlvTraffic(v int64) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.FlvTraffic = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetHlsBps(v float32) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.HlsBps = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetHlsCount(v int64) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.HlsCount = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetHlsTraffic(v int64) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.HlsTraffic = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetNewConns(v string) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.NewConns = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetP2pBps(v float32) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.P2pBps = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetP2pCount(v int64) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.P2pCount = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetP2pTraffic(v int64) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.P2pTraffic = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetRtmpBps(v float32) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.RtmpBps = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetRtmpCount(v int64) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.RtmpCount = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetRtmpTraffic(v int64) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.RtmpTraffic = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetRtsBps(v float32) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.RtsBps = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetRtsCount(v int64) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.RtsCount = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetRtsTraffic(v int64) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.RtsTraffic = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetStreamName(v string) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetTimeStamp(v string) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) SetTraffic(v int64) *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData {
	s.Traffic = &v
	return s
}

func (s *DescribeLiveStreamMetricDetailDataResponseBodyStreamDetailDataStreamData) Validate() error {
	return dara.Validate(s)
}
