// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublicMediaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaInfo(v *GetPublicMediaInfoResponseBodyMediaInfo) *GetPublicMediaInfoResponseBody
	GetMediaInfo() *GetPublicMediaInfoResponseBodyMediaInfo
	SetRequestId(v string) *GetPublicMediaInfoResponseBody
	GetRequestId() *string
}

type GetPublicMediaInfoResponseBody struct {
	MediaInfo *GetPublicMediaInfoResponseBodyMediaInfo `json:"MediaInfo,omitempty" xml:"MediaInfo,omitempty" type:"Struct"`
	// RequestId
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPublicMediaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPublicMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicMediaInfoResponseBody) GetMediaInfo() *GetPublicMediaInfoResponseBodyMediaInfo {
	return s.MediaInfo
}

func (s *GetPublicMediaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPublicMediaInfoResponseBody) SetMediaInfo(v *GetPublicMediaInfoResponseBodyMediaInfo) *GetPublicMediaInfoResponseBody {
	s.MediaInfo = v
	return s
}

func (s *GetPublicMediaInfoResponseBody) SetRequestId(v string) *GetPublicMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublicMediaInfoResponseBody) Validate() error {
	if s.MediaInfo != nil {
		if err := s.MediaInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPublicMediaInfoResponseBodyMediaInfo struct {
	DynamicMetaData *GetPublicMediaInfoResponseBodyMediaInfoDynamicMetaData `json:"DynamicMetaData,omitempty" xml:"DynamicMetaData,omitempty" type:"Struct"`
	// FileInfos
	FileInfoList []*GetPublicMediaInfoResponseBodyMediaInfoFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	// BasicInfo
	MediaBasicInfo *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// example:
	//
	// icepublic-****14e501538aeef0a3140176f6****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s GetPublicMediaInfoResponseBodyMediaInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPublicMediaInfoResponseBodyMediaInfo) GoString() string {
	return s.String()
}

func (s *GetPublicMediaInfoResponseBodyMediaInfo) GetDynamicMetaData() *GetPublicMediaInfoResponseBodyMediaInfoDynamicMetaData {
	return s.DynamicMetaData
}

func (s *GetPublicMediaInfoResponseBodyMediaInfo) GetFileInfoList() []*GetPublicMediaInfoResponseBodyMediaInfoFileInfoList {
	return s.FileInfoList
}

func (s *GetPublicMediaInfoResponseBodyMediaInfo) GetMediaBasicInfo() *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *GetPublicMediaInfoResponseBodyMediaInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *GetPublicMediaInfoResponseBodyMediaInfo) SetDynamicMetaData(v *GetPublicMediaInfoResponseBodyMediaInfoDynamicMetaData) *GetPublicMediaInfoResponseBodyMediaInfo {
	s.DynamicMetaData = v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfo) SetFileInfoList(v []*GetPublicMediaInfoResponseBodyMediaInfoFileInfoList) *GetPublicMediaInfoResponseBodyMediaInfo {
	s.FileInfoList = v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfo) SetMediaBasicInfo(v *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) *GetPublicMediaInfoResponseBodyMediaInfo {
	s.MediaBasicInfo = v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfo) SetMediaId(v string) *GetPublicMediaInfoResponseBodyMediaInfo {
	s.MediaId = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfo) Validate() error {
	if s.DynamicMetaData != nil {
		if err := s.DynamicMetaData.Validate(); err != nil {
			return err
		}
	}
	if s.FileInfoList != nil {
		for _, item := range s.FileInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MediaBasicInfo != nil {
		if err := s.MediaBasicInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPublicMediaInfoResponseBodyMediaInfoDynamicMetaData struct {
	// example:
	//
	// {"AuditionUrl": "http://example-bucket.cdn.domain.com/example.mp4", "AuditionCount": 3}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPublicMediaInfoResponseBodyMediaInfoDynamicMetaData) String() string {
	return dara.Prettify(s)
}

func (s GetPublicMediaInfoResponseBodyMediaInfoDynamicMetaData) GoString() string {
	return s.String()
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoDynamicMetaData) GetData() *string {
	return s.Data
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoDynamicMetaData) GetType() *string {
	return s.Type
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoDynamicMetaData) SetData(v string) *GetPublicMediaInfoResponseBodyMediaInfoDynamicMetaData {
	s.Data = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoDynamicMetaData) SetType(v string) *GetPublicMediaInfoResponseBodyMediaInfoDynamicMetaData {
	s.Type = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoDynamicMetaData) Validate() error {
	return dara.Validate(s)
}

type GetPublicMediaInfoResponseBodyMediaInfoFileInfoList struct {
	AudioStreamInfoList    []*GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList    `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	FileBasicInfo          *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo            `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	SubtitleStreamInfoList []*GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList `json:"SubtitleStreamInfoList,omitempty" xml:"SubtitleStreamInfoList,omitempty" type:"Repeated"`
	VideoStreamInfoList    []*GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList    `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
}

func (s GetPublicMediaInfoResponseBodyMediaInfoFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPublicMediaInfoResponseBodyMediaInfoFileInfoList) GoString() string {
	return s.String()
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoList) GetAudioStreamInfoList() []*GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	return s.AudioStreamInfoList
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoList) GetFileBasicInfo() *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	return s.FileBasicInfo
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoList) GetSubtitleStreamInfoList() []*GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	return s.SubtitleStreamInfoList
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoList) GetVideoStreamInfoList() []*GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	return s.VideoStreamInfoList
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoList) SetAudioStreamInfoList(v []*GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoList {
	s.AudioStreamInfoList = v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoList) SetFileBasicInfo(v *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoList) SetSubtitleStreamInfoList(v []*GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoList {
	s.SubtitleStreamInfoList = v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoList) SetVideoStreamInfoList(v []*GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoList {
	s.VideoStreamInfoList = v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoList) Validate() error {
	if s.AudioStreamInfoList != nil {
		for _, item := range s.AudioStreamInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FileBasicInfo != nil {
		if err := s.FileBasicInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SubtitleStreamInfoList != nil {
		for _, item := range s.SubtitleStreamInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoStreamInfoList != nil {
		for _, item := range s.VideoStreamInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList struct {
	// example:
	//
	// 192.0
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// stereo
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// example:
	//
	// AAC (Advanced Audio Coding)
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// example:
	//
	// aac
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// example:
	//
	// 0x6134706d
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// example:
	//
	// mp4a
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// example:
	//
	// 1/44100
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// example:
	//
	// 16.2
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 10
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// und
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 162
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// example:
	//
	// High
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// example:
	//
	// fltp
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// example:
	//
	// 44100
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// example:
	//
	// 0.000000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1/44100
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetChannels() *string {
	return s.Channels
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetBitrate(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetChannelLayout(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetChannels(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecLongName(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecName(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTag(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTagString(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTimeBase(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetDuration(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetFps(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Fps = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetIndex(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetLang(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetNumFrames(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetProfile(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Profile = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetSampleFmt(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetSampleRate(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetStartTime(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetTimebase(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo struct {
	// example:
	//
	// 192.0
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// 16.2
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// example.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 27007
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// Normal
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// example:
	//
	// source_file
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// http://example-bucket.cdn.domain.com/example.mp4
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// mp4
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// example:
	//
	// 0
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 0
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetBitrate(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetDuration(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileName(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileSize(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileStatus(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileType(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileUrl(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFormatName(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetHeight(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetRegion(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetWidth(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList struct {
	// example:
	//
	// SubRip Text
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// example:
	//
	// srt
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// example:
	//
	// unicode
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// example:
	//
	// unicode
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// example:
	//
	// 29.97
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// und
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 0
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 30
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecLongName(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecName(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTag(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTagString(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTimeBase(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetDuration(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetIndex(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetLang(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetStartTime(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetTimebase(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList struct {
	// example:
	//
	// 24.0
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// example:
	//
	// 1001.594
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// example:
	//
	// h264
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// example:
	//
	// 0x0000
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// example:
	//
	// [0][0][0][0]
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// example:
	//
	// 1/48
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// example:
	//
	// 0:1
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// example:
	//
	// 216.206706
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 24.0
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// example:
	//
	// 2
	HasBFrames *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	// example:
	//
	// 540
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// und
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 30
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 5184
	NbFrames *string `json:"Nb_frames,omitempty" xml:"Nb_frames,omitempty"`
	// example:
	//
	// 5184
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// example:
	//
	// High
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// example:
	//
	// 0
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// example:
	//
	// 0:1
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// example:
	//
	// 0.081706
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1/12288
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// example:
	//
	// 960
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetDar() *string {
	return s.Dar
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetHeight() *string {
	return s.Height
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetLevel() *string {
	return s.Level
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetNbFrames() *string {
	return s.NbFrames
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetRotate() *string {
	return s.Rotate
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetSar() *string {
	return s.Sar
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetWidth() *string {
	return s.Width
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetAvgFPS(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.AvgFPS = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetBitrate(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecLongName(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecName(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTag(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTagString(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTimeBase(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetDar(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetDuration(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetFps(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetHasBFrames(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetHeight(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetIndex(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetLang(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetLevel(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetNbFrames(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.NbFrames = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetNumFrames(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetPixFmt(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetProfile(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetRotate(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetSar(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetStartTime(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetTimebase(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetWidth(v string) *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo struct {
	// example:
	//
	// general
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// category
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.png?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// example:
	//
	// 2020-12-26T04:11:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2020-12-26T04:11:15Z
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// MediaId
	//
	// example:
	//
	// icepublic-****14e501538aeef0a3140176f6****
	MediaId   *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 2020-12-26T04:11:10Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// oss
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// [{"bucket":"example-bucket","count":"32","iceJobId":"******83ec44d58b2069def2e******","location":"oss-cn-shanghai","snapshotRegular":"example/example-{Count}.jpg","spriteRegular":"example/example-{TileCount}.jpg","templateId":"******e438b14ff39293eaec25******","tileCount":"1"}]
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// {"key":"value"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetCategory() *string {
	return s.Category
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetSpriteImages() *string {
	return s.SpriteImages
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetBusinessType(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCategory(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCoverURL(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCreateTime(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetDeletedTime(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetDescription(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaId(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaTags(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaType(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetModifiedTime(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetSource(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetSpriteImages(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetStatus(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetTitle(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetUserData(v string) *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *GetPublicMediaInfoResponseBodyMediaInfoMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}
