// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetMediaInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoredList(v []*string) *BatchGetMediaInfosResponseBody
	GetIgnoredList() []*string
	SetMediaInfos(v []*BatchGetMediaInfosResponseBodyMediaInfos) *BatchGetMediaInfosResponseBody
	GetMediaInfos() []*BatchGetMediaInfosResponseBodyMediaInfos
	SetRequestId(v string) *BatchGetMediaInfosResponseBody
	GetRequestId() *string
}

type BatchGetMediaInfosResponseBody struct {
	// A list of media IDs for which information could not be retrieved.
	IgnoredList []*string `json:"IgnoredList,omitempty" xml:"IgnoredList,omitempty" type:"Repeated"`
	// A list of media assets.
	MediaInfos []*BatchGetMediaInfosResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetMediaInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBody) GetIgnoredList() []*string {
	return s.IgnoredList
}

func (s *BatchGetMediaInfosResponseBody) GetMediaInfos() []*BatchGetMediaInfosResponseBodyMediaInfos {
	return s.MediaInfos
}

func (s *BatchGetMediaInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetMediaInfosResponseBody) SetIgnoredList(v []*string) *BatchGetMediaInfosResponseBody {
	s.IgnoredList = v
	return s
}

func (s *BatchGetMediaInfosResponseBody) SetMediaInfos(v []*BatchGetMediaInfosResponseBodyMediaInfos) *BatchGetMediaInfosResponseBody {
	s.MediaInfos = v
	return s
}

func (s *BatchGetMediaInfosResponseBody) SetRequestId(v string) *BatchGetMediaInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBody) Validate() error {
	if s.MediaInfos != nil {
		for _, item := range s.MediaInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetMediaInfosResponseBodyMediaInfos struct {
	// A list of basic file information.
	FileInfoList []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	// The basic information about the media asset.
	MediaBasicInfo *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// The dynamic information about the media asset.
	MediaDynamicInfo *BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfo `json:"MediaDynamicInfo,omitempty" xml:"MediaDynamicInfo,omitempty" type:"Struct"`
	// The media ID.
	//
	// example:
	//
	// ******c48fb37407365d4f2cd8******
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfos) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) GetFileInfoList() []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoList {
	return s.FileInfoList
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) GetMediaBasicInfo() *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) GetMediaDynamicInfo() *BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfo {
	return s.MediaDynamicInfo
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) SetFileInfoList(v []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) *BatchGetMediaInfosResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) SetMediaBasicInfo(v *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) *BatchGetMediaInfosResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) SetMediaDynamicInfo(v *BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfo) *BatchGetMediaInfosResponseBodyMediaInfos {
	s.MediaDynamicInfo = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) SetMediaId(v string) *BatchGetMediaInfosResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) Validate() error {
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

type BatchGetMediaInfosResponseBodyMediaInfosFileInfoList struct {
	// The audio streams.
	AudioStreamInfoList []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	// The basic information about the file, such as the duration and file size.
	FileBasicInfo *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	// The subtitle streams.
	SubtitleStreamInfoList []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList `json:"SubtitleStreamInfoList,omitempty" xml:"SubtitleStreamInfoList,omitempty" type:"Repeated"`
	// The video streams.
	VideoStreamInfoList []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) GetAudioStreamInfoList() []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	return s.AudioStreamInfoList
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) GetFileBasicInfo() *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	return s.FileBasicInfo
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) GetSubtitleStreamInfoList() []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	return s.SubtitleStreamInfoList
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) GetVideoStreamInfoList() []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	return s.VideoStreamInfoList
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) SetAudioStreamInfoList(v []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList {
	s.AudioStreamInfoList = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) SetSubtitleStreamInfoList(v []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList {
	s.SubtitleStreamInfoList = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) SetVideoStreamInfoList(v []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList {
	s.VideoStreamInfoList = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) Validate() error {
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

type BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList struct {
	// The bitrate.
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The channel layout.
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// The number of audio channels.
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The full name of the codec.
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the codec.
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The codec tag.
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The codec tag string.
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The time base of the codec.
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The duration.
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The frame rate.
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The index of the stream.
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The total number of frames.
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// The profile.
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The sample format.
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// The sample rate.
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// The start time.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetChannels() *string {
	return s.Channels
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetBitrate(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetChannelLayout(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetChannels(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetCodecLongName(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetCodecName(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetCodecTag(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetCodecTagString(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetCodecTimeBase(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetDuration(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetFps(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Fps = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetIndex(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetLang(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Lang = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetNumFrames(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetProfile(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Profile = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetSampleFmt(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetSampleRate(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetStartTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetTimebase(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListAudioStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
	// The bitrate.
	//
	// example:
	//
	// 1132.68
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The duration.
	//
	// example:
	//
	// 200
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The file name.
	//
	// example:
	//
	// example
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file size, in bytes.
	//
	// example:
	//
	// 30611502
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The file status.
	//
	// example:
	//
	// Normal
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// The file type.
	//
	// example:
	//
	// source_file
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The OSS URL of the file.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The container format.
	//
	// example:
	//
	// mov,mp4,m4a,3gp,3g2,mj2
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// The height.
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The region where the file is stored.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The width.
	//
	// example:
	//
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList struct {
	// The full name of the codec.
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the codec.
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The codec tag.
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The codec tag string.
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The time base of the codec.
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The duration.
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The index of the stream.
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start time.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetCodecLongName(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetCodecName(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetCodecTag(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetCodecTagString(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetCodecTimeBase(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetDuration(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.Duration = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetIndex(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.Index = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetLang(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.Lang = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetStartTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetTimebase(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList struct {
	// The average frame rate.
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// The bitrate.
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The full name of the codec.
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the codec.
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The codec tag.
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The codec tag string.
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The time base of the codec.
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The display aspect ratio (DAR).
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// The duration.
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The frame rate.
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// Indicates whether B-frames exist.
	HasBFrames *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	// The height of the video.
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The index of the stream.
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The level.
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// This parameter is an alias for `NumFrames`.
	NbFrames *string `json:"Nb_frames,omitempty" xml:"Nb_frames,omitempty"`
	// The total number of frames.
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// The pixel format.
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The profile.
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The rotation angle.
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The sample aspect ratio (SAR).
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// The start time.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// The width of the video.
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetDar() *string {
	return s.Dar
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetHeight() *string {
	return s.Height
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetLevel() *string {
	return s.Level
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetNbFrames() *string {
	return s.NbFrames
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetRotate() *string {
	return s.Rotate
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetSar() *string {
	return s.Sar
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetWidth() *string {
	return s.Width
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetAvgFPS(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.AvgFPS = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetBitrate(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetCodecLongName(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetCodecName(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetCodecTag(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetCodecTagString(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetCodecTimeBase(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetDar(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetDuration(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetFps(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetHasBFrames(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetHeight(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetIndex(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetLang(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetLevel(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetNbFrames(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.NbFrames = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetNumFrames(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetPixFmt(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetProfile(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetRotate(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetSar(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetStartTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetTimebase(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetWidth(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListVideoStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo struct {
	// The business associated with the media asset.
	//
	// example:
	//
	// ICE
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The business type of the media asset.
	//
	// example:
	//
	// general
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The category.
	//
	// example:
	//
	// category1
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The cover URL.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.png?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time the media asset was created.
	//
	// example:
	//
	// 2020-12-26T04:11:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time the media asset was deleted.
	//
	// example:
	//
	// 2020-12-26T04:11:10Z
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// The description.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the media asset in its source system.
	//
	// example:
	//
	// https://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The media ID.
	//
	// example:
	//
	// ******c48fb37407365d4f2cd8******
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The tags.
	//
	// example:
	//
	// tag1, tag2
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// The media type. Valid values:
	//
	// - `Image`
	//
	// - `Video`
	//
	// - `Audio`
	//
	// - `Text`
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The time the media asset was last modified.
	//
	// example:
	//
	// 2020-12-26T04:11:12Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The snapshots.
	//
	// example:
	//
	// [
	//
	//     "http://example-bucket.oss-cn-shanghai.aliyuncs.com/snapshot-00001.png?Expires=&OSSAccessKeyId=&Signature=&security-token=",
	//
	//     "http://example-bucket.oss-cn-shanghai.aliyuncs.com/snapshot-00002.jpg?Expires=&OSSAccessKeyId=&Signature=&security-token=",
	//
	//     "http://example-bucket.oss-cn-shanghai.aliyuncs.com/snapshot-00003.jpg?Expires=&OSSAccessKeyId=&Signature=&security-token="
	//
	// ]
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// The source. Valid values:
	//
	// - `OSS`
	//
	// - `VOD`
	//
	// example:
	//
	// oss
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The sprite images.
	//
	// example:
	//
	// [{"bucket":"example-bucket","count":"32","iceJobId":"******83ec44d58b2069def2e******","location":"oss-cn-shanghai","snapshotRegular":"example/example-{Count}.jpg","spriteRegular":"example/example-{TileCount}.jpg","templateId":"******e438b14ff39293eaec25******","tileCount":"1"}]
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
	// The status of the media asset.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title.
	//
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The transcoding status.
	//
	// example:
	//
	// Init
	TranscodeStatus *string `json:"TranscodeStatus,omitempty" xml:"TranscodeStatus,omitempty"`
	// The user data.
	//
	// example:
	//
	// userDataTest
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetBiz() *string {
	return s.Biz
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetBusinessType() *string {
	return s.BusinessType
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetCategory() *string {
	return s.Category
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetInputURL() *string {
	return s.InputURL
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetSnapshots() *string {
	return s.Snapshots
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetSpriteImages() *string {
	return s.SpriteImages
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetTranscodeStatus() *string {
	return s.TranscodeStatus
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetBiz(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Biz = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetBusinessType(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetCategory(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetDeletedTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetSpriteImages(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetTranscodeStatus(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}

type BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfo struct {
	// The type of dynamic metadata. Valid values:
	//
	// - `"ai"`: Standardized data derived from raw AI results.
	//
	// - `"user-defined"`: The user-defined metadata.
	//
	// - `"system"`: The system-generated data.
	DynamicMetaData *BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData `json:"DynamicMetaData,omitempty" xml:"DynamicMetaData,omitempty" type:"Struct"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfo) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfo) GetDynamicMetaData() *BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData {
	return s.DynamicMetaData
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfo) SetDynamicMetaData(v *BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData) *BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfo {
	s.DynamicMetaData = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfo) Validate() error {
	if s.DynamicMetaData != nil {
		if err := s.DynamicMetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData struct {
	// The content of the dynamic metadata.
	//
	// example:
	//
	// system
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData) GetData() *string {
	return s.Data
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData) SetData(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData {
	s.Data = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData) Validate() error {
	return dara.Validate(s)
}
