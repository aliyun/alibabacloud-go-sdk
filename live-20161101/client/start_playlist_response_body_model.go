// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPlaylistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProgramId(v string) *StartPlaylistResponseBody
	GetProgramId() *string
	SetRequestId(v string) *StartPlaylistResponseBody
	GetRequestId() *string
	SetStreamInfo(v *StartPlaylistResponseBodyStreamInfo) *StartPlaylistResponseBody
	GetStreamInfo() *StartPlaylistResponseBodyStreamInfo
}

type StartPlaylistResponseBody struct {
	// The ID of the episode list. You can use the ID as a request parameter in the API operation that is used to stop playing the episode list.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b96****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the live stream.
	StreamInfo *StartPlaylistResponseBodyStreamInfo `json:"StreamInfo,omitempty" xml:"StreamInfo,omitempty" type:"Struct"`
}

func (s StartPlaylistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartPlaylistResponseBody) GoString() string {
	return s.String()
}

func (s *StartPlaylistResponseBody) GetProgramId() *string {
	return s.ProgramId
}

func (s *StartPlaylistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartPlaylistResponseBody) GetStreamInfo() *StartPlaylistResponseBodyStreamInfo {
	return s.StreamInfo
}

func (s *StartPlaylistResponseBody) SetProgramId(v string) *StartPlaylistResponseBody {
	s.ProgramId = &v
	return s
}

func (s *StartPlaylistResponseBody) SetRequestId(v string) *StartPlaylistResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartPlaylistResponseBody) SetStreamInfo(v *StartPlaylistResponseBodyStreamInfo) *StartPlaylistResponseBody {
	s.StreamInfo = v
	return s
}

func (s *StartPlaylistResponseBody) Validate() error {
	if s.StreamInfo != nil {
		if err := s.StreamInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartPlaylistResponseBodyStreamInfo struct {
	// The name of the application.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The streaming URLs.
	Streams *StartPlaylistResponseBodyStreamInfoStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Struct"`
}

func (s StartPlaylistResponseBodyStreamInfo) String() string {
	return dara.Prettify(s)
}

func (s StartPlaylistResponseBodyStreamInfo) GoString() string {
	return s.String()
}

func (s *StartPlaylistResponseBodyStreamInfo) GetAppName() *string {
	return s.AppName
}

func (s *StartPlaylistResponseBodyStreamInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *StartPlaylistResponseBodyStreamInfo) GetStreamName() *string {
	return s.StreamName
}

func (s *StartPlaylistResponseBodyStreamInfo) GetStreams() *StartPlaylistResponseBodyStreamInfoStreams {
	return s.Streams
}

func (s *StartPlaylistResponseBodyStreamInfo) SetAppName(v string) *StartPlaylistResponseBodyStreamInfo {
	s.AppName = &v
	return s
}

func (s *StartPlaylistResponseBodyStreamInfo) SetDomainName(v string) *StartPlaylistResponseBodyStreamInfo {
	s.DomainName = &v
	return s
}

func (s *StartPlaylistResponseBodyStreamInfo) SetStreamName(v string) *StartPlaylistResponseBodyStreamInfo {
	s.StreamName = &v
	return s
}

func (s *StartPlaylistResponseBodyStreamInfo) SetStreams(v *StartPlaylistResponseBodyStreamInfoStreams) *StartPlaylistResponseBodyStreamInfo {
	s.Streams = v
	return s
}

func (s *StartPlaylistResponseBodyStreamInfo) Validate() error {
	if s.Streams != nil {
		if err := s.Streams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartPlaylistResponseBodyStreamInfoStreams struct {
	Stream []*StartPlaylistResponseBodyStreamInfoStreamsStream `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
}

func (s StartPlaylistResponseBodyStreamInfoStreams) String() string {
	return dara.Prettify(s)
}

func (s StartPlaylistResponseBodyStreamInfoStreams) GoString() string {
	return s.String()
}

func (s *StartPlaylistResponseBodyStreamInfoStreams) GetStream() []*StartPlaylistResponseBodyStreamInfoStreamsStream {
	return s.Stream
}

func (s *StartPlaylistResponseBodyStreamInfoStreams) SetStream(v []*StartPlaylistResponseBodyStreamInfoStreamsStream) *StartPlaylistResponseBodyStreamInfoStreams {
	s.Stream = v
	return s
}

func (s *StartPlaylistResponseBodyStreamInfoStreams) Validate() error {
	if s.Stream != nil {
		for _, item := range s.Stream {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StartPlaylistResponseBodyStreamInfoStreamsStream struct {
	// The streaming URL in the Flash Video (FLV) format.
	//
	// example:
	//
	// http://aliyundoc.com/caster/liveStream****.flv?auth_key=1612772224-0-0-3632be7cd9907169e8b09e91099c****
	PullFlvUrl *string `json:"PullFlvUrl,omitempty" xml:"PullFlvUrl,omitempty"`
	// The streaming URL in the Real-Time Messaging Protocol (RTMP) format.
	//
	// example:
	//
	// rtmp:///aliyundoc.com/caster/liveStream****?auth_key=1612772224-0-0-4404ca59c0246226d49d01f734b1****
	PullM3U8Url *string `json:"PullM3U8Url,omitempty" xml:"PullM3U8Url,omitempty"`
	// The streaming URL in the M3U8 format.
	//
	// example:
	//
	// http://aliyundoc.com/caster/liveStream****.m3u8?auth_key=1612772224-0-0-919a023a127156fe82e3562c3b3b****
	PullRtmpUrl *string `json:"PullRtmpUrl,omitempty" xml:"PullRtmpUrl,omitempty"`
	// The video quality of the live stream. Valid values: **original**: original quality
	//
	// example:
	//
	// original
	Quality *string `json:"Quality,omitempty" xml:"Quality,omitempty"`
}

func (s StartPlaylistResponseBodyStreamInfoStreamsStream) String() string {
	return dara.Prettify(s)
}

func (s StartPlaylistResponseBodyStreamInfoStreamsStream) GoString() string {
	return s.String()
}

func (s *StartPlaylistResponseBodyStreamInfoStreamsStream) GetPullFlvUrl() *string {
	return s.PullFlvUrl
}

func (s *StartPlaylistResponseBodyStreamInfoStreamsStream) GetPullM3U8Url() *string {
	return s.PullM3U8Url
}

func (s *StartPlaylistResponseBodyStreamInfoStreamsStream) GetPullRtmpUrl() *string {
	return s.PullRtmpUrl
}

func (s *StartPlaylistResponseBodyStreamInfoStreamsStream) GetQuality() *string {
	return s.Quality
}

func (s *StartPlaylistResponseBodyStreamInfoStreamsStream) SetPullFlvUrl(v string) *StartPlaylistResponseBodyStreamInfoStreamsStream {
	s.PullFlvUrl = &v
	return s
}

func (s *StartPlaylistResponseBodyStreamInfoStreamsStream) SetPullM3U8Url(v string) *StartPlaylistResponseBodyStreamInfoStreamsStream {
	s.PullM3U8Url = &v
	return s
}

func (s *StartPlaylistResponseBodyStreamInfoStreamsStream) SetPullRtmpUrl(v string) *StartPlaylistResponseBodyStreamInfoStreamsStream {
	s.PullRtmpUrl = &v
	return s
}

func (s *StartPlaylistResponseBodyStreamInfoStreamsStream) SetQuality(v string) *StartPlaylistResponseBodyStreamInfoStreamsStream {
	s.Quality = &v
	return s
}

func (s *StartPlaylistResponseBodyStreamInfoStreamsStream) Validate() error {
	return dara.Validate(s)
}
