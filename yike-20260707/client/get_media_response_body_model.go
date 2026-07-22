// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaInfo(v *GetMediaResponseBodyMediaInfo) *GetMediaResponseBody
	GetMediaInfo() *GetMediaResponseBodyMediaInfo
	SetRequestId(v string) *GetMediaResponseBody
	GetRequestId() *string
}

type GetMediaResponseBody struct {
	MediaInfo *GetMediaResponseBodyMediaInfo `json:"MediaInfo,omitempty" xml:"MediaInfo,omitempty" type:"Struct"`
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaResponseBody) GetMediaInfo() *GetMediaResponseBodyMediaInfo {
	return s.MediaInfo
}

func (s *GetMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaResponseBody) SetMediaInfo(v *GetMediaResponseBodyMediaInfo) *GetMediaResponseBody {
	s.MediaInfo = v
	return s
}

func (s *GetMediaResponseBody) SetRequestId(v string) *GetMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaResponseBody) Validate() error {
	if s.MediaInfo != nil {
		if err := s.MediaInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaResponseBodyMediaInfo struct {
	FileInfoList     []*GetMediaResponseBodyMediaInfoFileInfoList   `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	MediaBasicInfo   *GetMediaResponseBodyMediaInfoMediaBasicInfo   `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	MediaDynamicInfo *GetMediaResponseBodyMediaInfoMediaDynamicInfo `json:"MediaDynamicInfo,omitempty" xml:"MediaDynamicInfo,omitempty" type:"Struct"`
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s GetMediaResponseBodyMediaInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResponseBodyMediaInfo) GoString() string {
	return s.String()
}

func (s *GetMediaResponseBodyMediaInfo) GetFileInfoList() []*GetMediaResponseBodyMediaInfoFileInfoList {
	return s.FileInfoList
}

func (s *GetMediaResponseBodyMediaInfo) GetMediaBasicInfo() *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *GetMediaResponseBodyMediaInfo) GetMediaDynamicInfo() *GetMediaResponseBodyMediaInfoMediaDynamicInfo {
	return s.MediaDynamicInfo
}

func (s *GetMediaResponseBodyMediaInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaResponseBodyMediaInfo) SetFileInfoList(v []*GetMediaResponseBodyMediaInfoFileInfoList) *GetMediaResponseBodyMediaInfo {
	s.FileInfoList = v
	return s
}

func (s *GetMediaResponseBodyMediaInfo) SetMediaBasicInfo(v *GetMediaResponseBodyMediaInfoMediaBasicInfo) *GetMediaResponseBodyMediaInfo {
	s.MediaBasicInfo = v
	return s
}

func (s *GetMediaResponseBodyMediaInfo) SetMediaDynamicInfo(v *GetMediaResponseBodyMediaInfoMediaDynamicInfo) *GetMediaResponseBodyMediaInfo {
	s.MediaDynamicInfo = v
	return s
}

func (s *GetMediaResponseBodyMediaInfo) SetMediaId(v string) *GetMediaResponseBodyMediaInfo {
	s.MediaId = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfo) Validate() error {
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
	if s.MediaDynamicInfo != nil {
		if err := s.MediaDynamicInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaResponseBodyMediaInfoFileInfoList struct {
	AudioStreamInfoList    []*GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList    `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	FileBasicInfo          *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo            `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	SubtitleStreamInfoList []*GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList `json:"SubtitleStreamInfoList,omitempty" xml:"SubtitleStreamInfoList,omitempty" type:"Repeated"`
	VideoStreamInfoList    []*GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList    `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
}

func (s GetMediaResponseBodyMediaInfoFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResponseBodyMediaInfoFileInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaResponseBodyMediaInfoFileInfoList) GetAudioStreamInfoList() []*GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	return s.AudioStreamInfoList
}

func (s *GetMediaResponseBodyMediaInfoFileInfoList) GetFileBasicInfo() *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo {
	return s.FileBasicInfo
}

func (s *GetMediaResponseBodyMediaInfoFileInfoList) GetSubtitleStreamInfoList() []*GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	return s.SubtitleStreamInfoList
}

func (s *GetMediaResponseBodyMediaInfoFileInfoList) GetVideoStreamInfoList() []*GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	return s.VideoStreamInfoList
}

func (s *GetMediaResponseBodyMediaInfoFileInfoList) SetAudioStreamInfoList(v []*GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) *GetMediaResponseBodyMediaInfoFileInfoList {
	s.AudioStreamInfoList = v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoList) SetFileBasicInfo(v *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) *GetMediaResponseBodyMediaInfoFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoList) SetSubtitleStreamInfoList(v []*GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) *GetMediaResponseBodyMediaInfoFileInfoList {
	s.SubtitleStreamInfoList = v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoList) SetVideoStreamInfoList(v []*GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) *GetMediaResponseBodyMediaInfoFileInfoList {
	s.VideoStreamInfoList = v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoList) Validate() error {
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

type GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList struct {
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
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 10
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
	// 2026-02-04T02:13:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1/44100
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetChannels() *string {
	return s.Channels
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetBitrate(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetChannelLayout(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetChannels(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecLongName(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecName(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTag(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTagString(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTimeBase(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetDuration(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetFps(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Fps = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetIndex(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetLang(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetNumFrames(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetProfile(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Profile = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetSampleFmt(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetSampleRate(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetStartTime(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetTimebase(v string) *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListAudioStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo struct {
	// example:
	//
	// 30
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// 2020-12-26T04:11:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 6
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
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// mp4
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// example:
	//
	// 540
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 2020-12-26T04:11:10Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 960
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) SetBitrate(v string) *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) SetCreateTime(v string) *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) SetDuration(v string) *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileName(v string) *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileSize(v string) *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileStatus(v string) *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileType(v string) *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileUrl(v string) *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFormatName(v string) *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) SetHeight(v string) *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) SetModifiedTime(v string) *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) SetRegion(v string) *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) SetWidth(v string) *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList struct {
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
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 2025-03-07T01:30Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 30
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecLongName(v string) *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecName(v string) *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTag(v string) *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTagString(v string) *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTimeBase(v string) *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetDuration(v string) *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetIndex(v string) *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetLang(v string) *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetStartTime(v string) *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetTimebase(v string) *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList struct {
	// example:
	//
	// 24.0
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// example:
	//
	// 20
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
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 30
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// xxx
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
	// 2025-08-04T12:00:00Z
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

func (s GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetDar() *string {
	return s.Dar
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetHeight() *string {
	return s.Height
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetLevel() *string {
	return s.Level
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetNbFrames() *string {
	return s.NbFrames
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetRotate() *string {
	return s.Rotate
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetSar() *string {
	return s.Sar
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GetWidth() *string {
	return s.Width
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetAvgFPS(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.AvgFPS = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetBitrate(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecLongName(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecName(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTag(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTagString(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTimeBase(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetDar(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetDuration(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetFps(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetHasBFrames(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetHeight(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetIndex(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetLang(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetLevel(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetNbFrames(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.NbFrames = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetNumFrames(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetPixFmt(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetProfile(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetRotate(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetSar(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetStartTime(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetTimebase(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetWidth(v string) *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoFileInfoListVideoStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetMediaResponseBodyMediaInfoMediaBasicInfo struct {
	// example:
	//
	// general
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	CategoryId   *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// https://dtlive-vip.oss-cn-shanghai.aliyuncs.com/cover/4e88a055-75fc-4ff5-8b8a-f32224917514_open_live_cover.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// example:
	//
	// 2020-12-26T04:11:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1586676
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// https://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// pitaya,prd-wuxi,248
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// example:
	//
	// image
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 2020-12-26T04:11:08Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// []
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// example:
	//
	// /
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// []
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// T32
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// general
	UploadSource *string `json:"UploadSource,omitempty" xml:"UploadSource,omitempty"`
	// example:
	//
	// {}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetMediaResponseBodyMediaInfoMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResponseBodyMediaInfoMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetCategoryName() *string {
	return s.CategoryName
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetEntityId() *string {
	return s.EntityId
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetInputURL() *string {
	return s.InputURL
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetSnapshots() *string {
	return s.Snapshots
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetSpriteImages() *string {
	return s.SpriteImages
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetUploadSource() *string {
	return s.UploadSource
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetBusinessType(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetCategoryId(v int64) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.CategoryId = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetCategoryName(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.CategoryName = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetCoverURL(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetCreateTime(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetDescription(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetEntityId(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.EntityId = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetInputURL(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetMediaId(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetMediaTags(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetMediaType(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetModifiedTime(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetSnapshots(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetSource(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetSpriteImages(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetStatus(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetTitle(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetUploadSource(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.UploadSource = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) SetUserData(v string) *GetMediaResponseBodyMediaInfoMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}

type GetMediaResponseBodyMediaInfoMediaDynamicInfo struct {
	DynamicMetaData *GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData `json:"DynamicMetaData,omitempty" xml:"DynamicMetaData,omitempty" type:"Struct"`
}

func (s GetMediaResponseBodyMediaInfoMediaDynamicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResponseBodyMediaInfoMediaDynamicInfo) GoString() string {
	return s.String()
}

func (s *GetMediaResponseBodyMediaInfoMediaDynamicInfo) GetDynamicMetaData() *GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData {
	return s.DynamicMetaData
}

func (s *GetMediaResponseBodyMediaInfoMediaDynamicInfo) SetDynamicMetaData(v *GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData) *GetMediaResponseBodyMediaInfoMediaDynamicInfo {
	s.DynamicMetaData = v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaDynamicInfo) Validate() error {
	if s.DynamicMetaData != nil {
		if err := s.DynamicMetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData struct {
	// example:
	//
	// {}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 4614131
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData) GoString() string {
	return s.String()
}

func (s *GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData) GetData() *string {
	return s.Data
}

func (s *GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData) GetEntityId() *string {
	return s.EntityId
}

func (s *GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData) GetType() *string {
	return s.Type
}

func (s *GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData) SetData(v string) *GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData {
	s.Data = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData) SetEntityId(v string) *GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData {
	s.EntityId = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData) SetType(v string) *GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData {
	s.Type = &v
	return s
}

func (s *GetMediaResponseBodyMediaInfoMediaDynamicInfoDynamicMetaData) Validate() error {
	return dara.Validate(s)
}
