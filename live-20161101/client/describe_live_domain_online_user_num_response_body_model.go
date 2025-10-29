// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainOnlineUserNumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOnlineUserInfo(v *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfo) *DescribeLiveDomainOnlineUserNumResponseBody
	GetOnlineUserInfo() *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfo
	SetRequestId(v string) *DescribeLiveDomainOnlineUserNumResponseBody
	GetRequestId() *string
	SetStreamCount(v int32) *DescribeLiveDomainOnlineUserNumResponseBody
	GetStreamCount() *int32
	SetUserCount(v int32) *DescribeLiveDomainOnlineUserNumResponseBody
	GetUserCount() *int32
}

type DescribeLiveDomainOnlineUserNumResponseBody struct {
	// The information about the streams.
	OnlineUserInfo *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfo `json:"OnlineUserInfo,omitempty" xml:"OnlineUserInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 3A3A8C3D-F8B2-4FBF-9319-771A11B855FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of streams.
	//
	// example:
	//
	// 1
	StreamCount *int32 `json:"StreamCount,omitempty" xml:"StreamCount,omitempty"`
	// The total number of online users at the specified point in time for all the live streams under the main streaming domain.
	//
	// example:
	//
	// 1
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s DescribeLiveDomainOnlineUserNumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainOnlineUserNumResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainOnlineUserNumResponseBody) GetOnlineUserInfo() *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfo {
	return s.OnlineUserInfo
}

func (s *DescribeLiveDomainOnlineUserNumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainOnlineUserNumResponseBody) GetStreamCount() *int32 {
	return s.StreamCount
}

func (s *DescribeLiveDomainOnlineUserNumResponseBody) GetUserCount() *int32 {
	return s.UserCount
}

func (s *DescribeLiveDomainOnlineUserNumResponseBody) SetOnlineUserInfo(v *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfo) *DescribeLiveDomainOnlineUserNumResponseBody {
	s.OnlineUserInfo = v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponseBody) SetRequestId(v string) *DescribeLiveDomainOnlineUserNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponseBody) SetStreamCount(v int32) *DescribeLiveDomainOnlineUserNumResponseBody {
	s.StreamCount = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponseBody) SetUserCount(v int32) *DescribeLiveDomainOnlineUserNumResponseBody {
	s.UserCount = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponseBody) Validate() error {
	if s.OnlineUserInfo != nil {
		if err := s.OnlineUserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfo struct {
	LiveStreamOnlineUserNumInfo []*DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo `json:"LiveStreamOnlineUserNumInfo,omitempty" xml:"LiveStreamOnlineUserNumInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfo) GetLiveStreamOnlineUserNumInfo() []*DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo {
	return s.LiveStreamOnlineUserNumInfo
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfo) SetLiveStreamOnlineUserNumInfo(v []*DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo) *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfo {
	s.LiveStreamOnlineUserNumInfo = v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfo) Validate() error {
	if s.LiveStreamOnlineUserNumInfo != nil {
		for _, item := range s.LiveStreamOnlineUserNumInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo struct {
	// The statistics on the stream.
	Infos *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos `json:"Infos,omitempty" xml:"Infos,omitempty" type:"Struct"`
	// The name of the stream.
	//
	// example:
	//
	// rtmp://example.com/test/liveStream****_3_1
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo) GetInfos() *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos {
	return s.Infos
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo) SetInfos(v *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos) *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo {
	s.Infos = v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo) SetStreamName(v string) *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfo) Validate() error {
	if s.Infos != nil {
		if err := s.Infos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos struct {
	Info []*DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo `json:"Info,omitempty" xml:"Info,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos) GetInfo() []*DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo {
	return s.Info
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos) SetInfo(v []*DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos {
	s.Info = v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfos) Validate() error {
	if s.Info != nil {
		for _, item := range s.Info {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo struct {
	// The transcoding template. A value of origin indicates that the stream is a source stream.
	//
	// example:
	//
	// origin
	TranscodeTemplate *string `json:"TranscodeTemplate,omitempty" xml:"TranscodeTemplate,omitempty"`
	// The number of online users for the stream, which can be a source stream or transcoded stream.
	//
	// example:
	//
	// 1
	UserNumber *int64 `json:"UserNumber,omitempty" xml:"UserNumber,omitempty"`
}

func (s DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) GetTranscodeTemplate() *string {
	return s.TranscodeTemplate
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) GetUserNumber() *int64 {
	return s.UserNumber
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) SetTranscodeTemplate(v string) *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo {
	s.TranscodeTemplate = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) SetUserNumber(v int64) *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo {
	s.UserNumber = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponseBodyOnlineUserInfoLiveStreamOnlineUserNumInfoInfosInfo) Validate() error {
	return dara.Validate(s)
}
