// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsOnlineListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOnlineInfo(v *DescribeLiveStreamsOnlineListResponseBodyOnlineInfo) *DescribeLiveStreamsOnlineListResponseBody
	GetOnlineInfo() *DescribeLiveStreamsOnlineListResponseBodyOnlineInfo
	SetPageNum(v int32) *DescribeLiveStreamsOnlineListResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveStreamsOnlineListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLiveStreamsOnlineListResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *DescribeLiveStreamsOnlineListResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeLiveStreamsOnlineListResponseBody
	GetTotalPage() *int32
}

type DescribeLiveStreamsOnlineListResponseBody struct {
	// The information about the live streams that are being ingested.
	OnlineInfo *DescribeLiveStreamsOnlineListResponseBodyOnlineInfo `json:"OnlineInfo,omitempty" xml:"OnlineInfo,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5c6a2a0df228-4a64- af62-20e91b9676b3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of streams that meet the specified conditions.
	//
	// example:
	//
	// 11
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeLiveStreamsOnlineListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsOnlineListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsOnlineListResponseBody) GetOnlineInfo() *DescribeLiveStreamsOnlineListResponseBodyOnlineInfo {
	return s.OnlineInfo
}

func (s *DescribeLiveStreamsOnlineListResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveStreamsOnlineListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamsOnlineListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamsOnlineListResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeLiveStreamsOnlineListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeLiveStreamsOnlineListResponseBody) SetOnlineInfo(v *DescribeLiveStreamsOnlineListResponseBodyOnlineInfo) *DescribeLiveStreamsOnlineListResponseBody {
	s.OnlineInfo = v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBody) SetPageNum(v int32) *DescribeLiveStreamsOnlineListResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBody) SetPageSize(v int32) *DescribeLiveStreamsOnlineListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBody) SetRequestId(v string) *DescribeLiveStreamsOnlineListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBody) SetTotalNum(v int32) *DescribeLiveStreamsOnlineListResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBody) SetTotalPage(v int32) *DescribeLiveStreamsOnlineListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBody) Validate() error {
	if s.OnlineInfo != nil {
		if err := s.OnlineInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveStreamsOnlineListResponseBodyOnlineInfo struct {
	LiveStreamOnlineInfo []*DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo `json:"LiveStreamOnlineInfo,omitempty" xml:"LiveStreamOnlineInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamsOnlineListResponseBodyOnlineInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsOnlineListResponseBodyOnlineInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfo) GetLiveStreamOnlineInfo() []*DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	return s.LiveStreamOnlineInfo
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfo) SetLiveStreamOnlineInfo(v []*DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfo {
	s.LiveStreamOnlineInfo = v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfo) Validate() error {
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

type DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the audio codec.
	//
	// example:
	//
	// 10
	AudioCodecId *int32 `json:"AudioCodecId,omitempty" xml:"AudioCodecId,omitempty"`
	// The audio bitrate of the live stream. Unit: Kbit/s.
	//
	// >  This parameter can be returned after you submit a ticket for whitelist configuration. For more information about how to submit a ticket, see Contact us.
	//
	// example:
	//
	// 600
	AudioDataRate *int32 `json:"AudioDataRate,omitempty" xml:"AudioDataRate,omitempty"`
	// The IP address of the client for stream ingest.
	//
	// example:
	//
	// 106.11.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The frame rate. Unit: FPS.
	//
	// example:
	//
	// 15
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// The height of the video resolution. Unit: pixels.
	//
	// example:
	//
	// 720
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The ingest domain. If live center ingest was used, the streaming domain is returned.
	//
	// example:
	//
	// demo.aliyundoc.com
	PublishDomain *string `json:"PublishDomain,omitempty" xml:"PublishDomain,omitempty"`
	// The start time of stream ingest. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-08-09T02:37:59Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The ingest type. Valid values:
	//
	// 	- **edge**: edge ingest.
	//
	// 	- **center**: live center ingest.
	//
	// example:
	//
	// edge
	PublishType *string `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	// The complete URL that was used to ingest the stream.
	//
	// example:
	//
	// rtmp://demo.aliyundoc.com/live/test****
	PublishUrl *string `json:"PublishUrl,omitempty" xml:"PublishUrl,omitempty"`
	// The IP address of the ingest node.
	//
	// example:
	//
	// 120.221.XX.XX
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// Indicates whether the stream was transcoded. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// no
	Transcoded *string `json:"Transcoded,omitempty" xml:"Transcoded,omitempty"`
	// The ID of the video codec.
	//
	// example:
	//
	// 7
	VideoCodecId *int32 `json:"VideoCodecId,omitempty" xml:"VideoCodecId,omitempty"`
	// The video bitrate of the live stream. Unit: Kbit/s.
	//
	// >  This parameter can be returned after you submit a ticket for whitelist configuration. For more information about how to submit a ticket, see Contact us.
	//
	// example:
	//
	// 600
	VideoDataRate *int32 `json:"VideoDataRate,omitempty" xml:"VideoDataRate,omitempty"`
	// The width of the video resolution. Unit: pixels.
	//
	// example:
	//
	// 1280
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetAudioCodecId() *int32 {
	return s.AudioCodecId
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetAudioDataRate() *int32 {
	return s.AudioDataRate
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetFrameRate() *int32 {
	return s.FrameRate
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetHeight() *int32 {
	return s.Height
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetPublishDomain() *string {
	return s.PublishDomain
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetPublishTime() *string {
	return s.PublishTime
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetPublishType() *string {
	return s.PublishType
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetPublishUrl() *string {
	return s.PublishUrl
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetServerIp() *string {
	return s.ServerIp
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetTranscoded() *string {
	return s.Transcoded
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetVideoCodecId() *int32 {
	return s.VideoCodecId
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetVideoDataRate() *int32 {
	return s.VideoDataRate
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) GetWidth() *int32 {
	return s.Width
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetAppName(v string) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetAudioCodecId(v int32) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.AudioCodecId = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetAudioDataRate(v int32) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.AudioDataRate = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetClientIp(v string) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.ClientIp = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetDomainName(v string) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetFrameRate(v int32) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.FrameRate = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetHeight(v int32) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.Height = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetPublishDomain(v string) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.PublishDomain = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetPublishTime(v string) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.PublishTime = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetPublishType(v string) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.PublishType = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetPublishUrl(v string) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.PublishUrl = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetServerIp(v string) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.ServerIp = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetStreamName(v string) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetTranscoded(v string) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.Transcoded = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetVideoCodecId(v int32) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.VideoCodecId = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetVideoDataRate(v int32) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.VideoDataRate = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) SetWidth(v int32) *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo {
	s.Width = &v
	return s
}

func (s *DescribeLiveStreamsOnlineListResponseBodyOnlineInfoLiveStreamOnlineInfo) Validate() error {
	return dara.Validate(s)
}
