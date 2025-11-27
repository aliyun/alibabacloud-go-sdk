// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsStreamsOnlineListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOnlineInfo(v *DescribeVsStreamsOnlineListResponseBodyOnlineInfo) *DescribeVsStreamsOnlineListResponseBody
	GetOnlineInfo() *DescribeVsStreamsOnlineListResponseBodyOnlineInfo
	SetPageNum(v int32) *DescribeVsStreamsOnlineListResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeVsStreamsOnlineListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVsStreamsOnlineListResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *DescribeVsStreamsOnlineListResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeVsStreamsOnlineListResponseBody
	GetTotalPage() *int32
}

type DescribeVsStreamsOnlineListResponseBody struct {
	OnlineInfo *DescribeVsStreamsOnlineListResponseBodyOnlineInfo `json:"OnlineInfo,omitempty" xml:"OnlineInfo,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// B31FC4AD-3592-573E-8063-878F722B322A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeVsStreamsOnlineListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsStreamsOnlineListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsOnlineListResponseBody) GetOnlineInfo() *DescribeVsStreamsOnlineListResponseBodyOnlineInfo {
	return s.OnlineInfo
}

func (s *DescribeVsStreamsOnlineListResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeVsStreamsOnlineListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVsStreamsOnlineListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsStreamsOnlineListResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeVsStreamsOnlineListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeVsStreamsOnlineListResponseBody) SetOnlineInfo(v *DescribeVsStreamsOnlineListResponseBodyOnlineInfo) *DescribeVsStreamsOnlineListResponseBody {
	s.OnlineInfo = v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBody) SetPageNum(v int32) *DescribeVsStreamsOnlineListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBody) SetPageSize(v int32) *DescribeVsStreamsOnlineListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBody) SetRequestId(v string) *DescribeVsStreamsOnlineListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBody) SetTotalNum(v int32) *DescribeVsStreamsOnlineListResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBody) SetTotalPage(v int32) *DescribeVsStreamsOnlineListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBody) Validate() error {
	if s.OnlineInfo != nil {
		if err := s.OnlineInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVsStreamsOnlineListResponseBodyOnlineInfo struct {
	LiveStreamOnlineInfo []*DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo `json:"LiveStreamOnlineInfo,omitempty" xml:"LiveStreamOnlineInfo,omitempty" type:"Repeated"`
}

func (s DescribeVsStreamsOnlineListResponseBodyOnlineInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsStreamsOnlineListResponseBodyOnlineInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfo) GetLiveStreamOnlineInfo() []*DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	return s.LiveStreamOnlineInfo
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfo) SetLiveStreamOnlineInfo(v []*DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) *DescribeVsStreamsOnlineListResponseBodyOnlineInfo {
	s.LiveStreamOnlineInfo = v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfo) Validate() error {
	if s.LiveStreamOnlineInfo != nil {
		for _, item := range s.LiveStreamOnlineInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo struct {
	// example:
	//
	// xxxApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// push.example.com
	PublishDomain *string `json:"PublishDomain,omitempty" xml:"PublishDomain,omitempty"`
	// example:
	//
	// 2015-12-02T06:58:04Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// example:
	//
	// edge
	PublishType *string `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	// example:
	//
	// rtmp://example.com/xchen
	PublishUrl *string `json:"PublishUrl,omitempty" xml:"PublishUrl,omitempty"`
	// example:
	//
	// testxchen_small
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// example:
	//
	// 123
	TranscodeId *string `json:"TranscodeId,omitempty" xml:"TranscodeId,omitempty"`
	// example:
	//
	// no
	Transcoded *string `json:"Transcoded,omitempty" xml:"Transcoded,omitempty"`
}

func (s DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetAppName() *string {
	return s.AppName
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetPublishDomain() *string {
	return s.PublishDomain
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetPublishTime() *string {
	return s.PublishTime
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetPublishType() *string {
	return s.PublishType
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetPublishUrl() *string {
	return s.PublishUrl
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetTranscodeId() *string {
	return s.TranscodeId
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetTranscoded() *string {
	return s.Transcoded
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetAppName(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.AppName = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetDomainName(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetPublishDomain(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.PublishDomain = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetPublishTime(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.PublishTime = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetPublishType(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.PublishType = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetPublishUrl(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.PublishUrl = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetStreamName(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetTranscodeId(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.TranscodeId = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetTranscoded(v string) *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.Transcoded = &v
	return s
}

func (s *DescribeVsStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) Validate() error {
	return dara.Validate(s)
}
