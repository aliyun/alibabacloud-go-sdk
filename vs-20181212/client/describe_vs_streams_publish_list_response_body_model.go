// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsStreamsPublishListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *DescribeVsStreamsPublishListResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeVsStreamsPublishListResponseBody
	GetPageSize() *int32
	SetPublishInfo(v *DescribeVsStreamsPublishListResponseBodyPublishInfo) *DescribeVsStreamsPublishListResponseBody
	GetPublishInfo() *DescribeVsStreamsPublishListResponseBodyPublishInfo
	SetRequestId(v string) *DescribeVsStreamsPublishListResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *DescribeVsStreamsPublishListResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeVsStreamsPublishListResponseBody
	GetTotalPage() *int32
}

type DescribeVsStreamsPublishListResponseBody struct {
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 3000
	PageSize    *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PublishInfo *DescribeVsStreamsPublishListResponseBodyPublishInfo `json:"PublishInfo,omitempty" xml:"PublishInfo,omitempty" type:"Struct"`
	// example:
	//
	// 119F7639-4646-51A4-B6C1-300D391C0104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeVsStreamsPublishListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsStreamsPublishListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsPublishListResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeVsStreamsPublishListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVsStreamsPublishListResponseBody) GetPublishInfo() *DescribeVsStreamsPublishListResponseBodyPublishInfo {
	return s.PublishInfo
}

func (s *DescribeVsStreamsPublishListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsStreamsPublishListResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeVsStreamsPublishListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeVsStreamsPublishListResponseBody) SetPageNum(v int32) *DescribeVsStreamsPublishListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBody) SetPageSize(v int32) *DescribeVsStreamsPublishListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBody) SetPublishInfo(v *DescribeVsStreamsPublishListResponseBodyPublishInfo) *DescribeVsStreamsPublishListResponseBody {
	s.PublishInfo = v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBody) SetRequestId(v string) *DescribeVsStreamsPublishListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBody) SetTotalNum(v int32) *DescribeVsStreamsPublishListResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBody) SetTotalPage(v int32) *DescribeVsStreamsPublishListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVsStreamsPublishListResponseBodyPublishInfo struct {
	LiveStreamPublishInfo []*DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo `json:"LiveStreamPublishInfo,omitempty" xml:"LiveStreamPublishInfo,omitempty" type:"Repeated"`
}

func (s DescribeVsStreamsPublishListResponseBodyPublishInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsStreamsPublishListResponseBodyPublishInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfo) GetLiveStreamPublishInfo() []*DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	return s.LiveStreamPublishInfo
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfo) SetLiveStreamPublishInfo(v []*DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) *DescribeVsStreamsPublishListResponseBodyPublishInfo {
	s.LiveStreamPublishInfo = v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo struct {
	// example:
	//
	// xxxApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 192.168.0.1
	ClientAddr *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 192.168.0.1
	EdgeNodeAddr *string `json:"EdgeNodeAddr,omitempty" xml:"EdgeNodeAddr,omitempty"`
	// example:
	//
	// example.aliyundoc.com
	PublishDomain *string `json:"PublishDomain,omitempty" xml:"PublishDomain,omitempty"`
	// example:
	//
	// 2016-06-29T19:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// example:
	//
	// center
	PublishType *string `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	// example:
	//
	// https://example.aliyundoc.com/xxxApp/3402000****320000001.m3u8
	PublishUrl *string `json:"PublishUrl,omitempty" xml:"PublishUrl,omitempty"`
	// example:
	//
	// 2016-06-29T19:00:00Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// example:
	//
	// xxxStream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// example:
	//
	// https://example.aliyundoc.com/xxxApp/3402000****320000001.m3u8
	StreamUrl *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
	// example:
	//
	// 3888920****8138204-cn-qingdao
	TranscodeId *string `json:"TranscodeId,omitempty" xml:"TranscodeId,omitempty"`
	// example:
	//
	// yes
	Transcoded *string `json:"Transcoded,omitempty" xml:"Transcoded,omitempty"`
}

func (s DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetAppName() *string {
	return s.AppName
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetClientAddr() *string {
	return s.ClientAddr
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetEdgeNodeAddr() *string {
	return s.EdgeNodeAddr
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetPublishDomain() *string {
	return s.PublishDomain
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetPublishTime() *string {
	return s.PublishTime
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetPublishType() *string {
	return s.PublishType
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetPublishUrl() *string {
	return s.PublishUrl
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetStopTime() *string {
	return s.StopTime
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetStreamUrl() *string {
	return s.StreamUrl
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetTranscodeId() *string {
	return s.TranscodeId
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetTranscoded() *string {
	return s.Transcoded
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetAppName(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.AppName = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetClientAddr(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.ClientAddr = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetDomainName(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetEdgeNodeAddr(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.EdgeNodeAddr = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetPublishDomain(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.PublishDomain = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetPublishTime(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.PublishTime = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetPublishType(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.PublishType = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetPublishUrl(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.PublishUrl = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetStopTime(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.StopTime = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetStreamName(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetStreamUrl(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.StreamUrl = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetTranscodeId(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.TranscodeId = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetTranscoded(v string) *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.Transcoded = &v
	return s
}

func (s *DescribeVsStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) Validate() error {
	return dara.Validate(s)
}
