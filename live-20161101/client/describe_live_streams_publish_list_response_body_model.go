// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsPublishListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *DescribeLiveStreamsPublishListResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveStreamsPublishListResponseBody
	GetPageSize() *int32
	SetPublishInfo(v *DescribeLiveStreamsPublishListResponseBodyPublishInfo) *DescribeLiveStreamsPublishListResponseBody
	GetPublishInfo() *DescribeLiveStreamsPublishListResponseBodyPublishInfo
	SetRequestId(v string) *DescribeLiveStreamsPublishListResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *DescribeLiveStreamsPublishListResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeLiveStreamsPublishListResponseBody
	GetTotalPage() *int32
}

type DescribeLiveStreamsPublishListResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 2
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize    *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PublishInfo *DescribeLiveStreamsPublishListResponseBodyPublishInfo `json:"PublishInfo,omitempty" xml:"PublishInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 40A4F36D-A7CC-473A-88E7-154F92242566
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that meet the specified conditions.
	//
	// example:
	//
	// 11
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeLiveStreamsPublishListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsPublishListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsPublishListResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveStreamsPublishListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamsPublishListResponseBody) GetPublishInfo() *DescribeLiveStreamsPublishListResponseBodyPublishInfo {
	return s.PublishInfo
}

func (s *DescribeLiveStreamsPublishListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamsPublishListResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeLiveStreamsPublishListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeLiveStreamsPublishListResponseBody) SetPageNum(v int32) *DescribeLiveStreamsPublishListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBody) SetPageSize(v int32) *DescribeLiveStreamsPublishListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBody) SetPublishInfo(v *DescribeLiveStreamsPublishListResponseBodyPublishInfo) *DescribeLiveStreamsPublishListResponseBody {
	s.PublishInfo = v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBody) SetRequestId(v string) *DescribeLiveStreamsPublishListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBody) SetTotalNum(v int32) *DescribeLiveStreamsPublishListResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBody) SetTotalPage(v int32) *DescribeLiveStreamsPublishListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBody) Validate() error {
	if s.PublishInfo != nil {
		if err := s.PublishInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveStreamsPublishListResponseBodyPublishInfo struct {
	LiveStreamPublishInfo []*DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo `json:"LiveStreamPublishInfo,omitempty" xml:"LiveStreamPublishInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamsPublishListResponseBodyPublishInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsPublishListResponseBodyPublishInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfo) GetLiveStreamPublishInfo() []*DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	return s.LiveStreamPublishInfo
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfo) SetLiveStreamPublishInfo(v []*DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) *DescribeLiveStreamsPublishListResponseBodyPublishInfo {
	s.LiveStreamPublishInfo = v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfo) Validate() error {
	if s.LiveStreamPublishInfo != nil {
		for _, item := range s.LiveStreamPublishInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo struct {
	AliInnerErrorFlags *string `json:"AliInnerErrorFlags,omitempty" xml:"AliInnerErrorFlags,omitempty"`
	AppName            *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ClientAddr         *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	DomainName         *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EdgeNodeAddr       *string `json:"EdgeNodeAddr,omitempty" xml:"EdgeNodeAddr,omitempty"`
	PublishDomain      *string `json:"PublishDomain,omitempty" xml:"PublishDomain,omitempty"`
	PublishTime        *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	PublishType        *string `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	PublishUrl         *string `json:"PublishUrl,omitempty" xml:"PublishUrl,omitempty"`
	StopTime           *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	StreamName         *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	StreamUrl          *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
	TranscodeId        *string `json:"TranscodeId,omitempty" xml:"TranscodeId,omitempty"`
	Transcoded         *string `json:"Transcoded,omitempty" xml:"Transcoded,omitempty"`
}

func (s DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetAliInnerErrorFlags() *string {
	return s.AliInnerErrorFlags
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetClientAddr() *string {
	return s.ClientAddr
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetEdgeNodeAddr() *string {
	return s.EdgeNodeAddr
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetPublishDomain() *string {
	return s.PublishDomain
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetPublishTime() *string {
	return s.PublishTime
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetPublishType() *string {
	return s.PublishType
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetPublishUrl() *string {
	return s.PublishUrl
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetStopTime() *string {
	return s.StopTime
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetStreamUrl() *string {
	return s.StreamUrl
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetTranscodeId() *string {
	return s.TranscodeId
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) GetTranscoded() *string {
	return s.Transcoded
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetAliInnerErrorFlags(v string) *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.AliInnerErrorFlags = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetAppName(v string) *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetClientAddr(v string) *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.ClientAddr = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetDomainName(v string) *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetEdgeNodeAddr(v string) *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.EdgeNodeAddr = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetPublishDomain(v string) *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.PublishDomain = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetPublishTime(v string) *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.PublishTime = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetPublishType(v string) *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.PublishType = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetPublishUrl(v string) *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.PublishUrl = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetStopTime(v string) *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.StopTime = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetStreamName(v string) *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetStreamUrl(v string) *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.StreamUrl = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetTranscodeId(v string) *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.TranscodeId = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) SetTranscoded(v string) *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo {
	s.Transcoded = &v
	return s
}

func (s *DescribeLiveStreamsPublishListResponseBodyPublishInfoLiveStreamPublishInfo) Validate() error {
	return dara.Validate(s)
}
