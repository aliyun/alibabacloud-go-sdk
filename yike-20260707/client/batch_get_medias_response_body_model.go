// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetMediasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoredList(v []*string) *BatchGetMediasResponseBody
	GetIgnoredList() []*string
	SetMediaInfos(v []*BatchGetMediasResponseBodyMediaInfos) *BatchGetMediasResponseBody
	GetMediaInfos() []*BatchGetMediasResponseBodyMediaInfos
	SetRequestId(v string) *BatchGetMediasResponseBody
	GetRequestId() *string
}

type BatchGetMediasResponseBody struct {
	IgnoredList []*string                               `json:"IgnoredList,omitempty" xml:"IgnoredList,omitempty" type:"Repeated"`
	MediaInfos  []*BatchGetMediasResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetMediasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediasResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetMediasResponseBody) GetIgnoredList() []*string {
	return s.IgnoredList
}

func (s *BatchGetMediasResponseBody) GetMediaInfos() []*BatchGetMediasResponseBodyMediaInfos {
	return s.MediaInfos
}

func (s *BatchGetMediasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetMediasResponseBody) SetIgnoredList(v []*string) *BatchGetMediasResponseBody {
	s.IgnoredList = v
	return s
}

func (s *BatchGetMediasResponseBody) SetMediaInfos(v []*BatchGetMediasResponseBodyMediaInfos) *BatchGetMediasResponseBody {
	s.MediaInfos = v
	return s
}

func (s *BatchGetMediasResponseBody) SetRequestId(v string) *BatchGetMediasResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetMediasResponseBody) Validate() error {
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

type BatchGetMediasResponseBodyMediaInfos struct {
	FileInfoList     []*BatchGetMediasResponseBodyMediaInfosFileInfoList   `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	MediaBasicInfo   *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo   `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	MediaDynamicInfo *BatchGetMediasResponseBodyMediaInfosMediaDynamicInfo `json:"MediaDynamicInfo,omitempty" xml:"MediaDynamicInfo,omitempty" type:"Struct"`
	// example:
	//
	// ******c48fb37407365d4f2cd8******
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s BatchGetMediasResponseBodyMediaInfos) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediasResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *BatchGetMediasResponseBodyMediaInfos) GetFileInfoList() []*BatchGetMediasResponseBodyMediaInfosFileInfoList {
	return s.FileInfoList
}

func (s *BatchGetMediasResponseBodyMediaInfos) GetMediaBasicInfo() *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	return s.MediaBasicInfo
}

func (s *BatchGetMediasResponseBodyMediaInfos) GetMediaDynamicInfo() *BatchGetMediasResponseBodyMediaInfosMediaDynamicInfo {
	return s.MediaDynamicInfo
}

func (s *BatchGetMediasResponseBodyMediaInfos) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetMediasResponseBodyMediaInfos) SetFileInfoList(v []*BatchGetMediasResponseBodyMediaInfosFileInfoList) *BatchGetMediasResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfos) SetMediaBasicInfo(v *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) *BatchGetMediasResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfos) SetMediaDynamicInfo(v *BatchGetMediasResponseBodyMediaInfosMediaDynamicInfo) *BatchGetMediasResponseBodyMediaInfos {
	s.MediaDynamicInfo = v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfos) SetMediaId(v string) *BatchGetMediasResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfos) Validate() error {
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

type BatchGetMediasResponseBodyMediaInfosFileInfoList struct {
	AudioStreamInfoList    []*BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList    `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	FileBasicInfo          *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo            `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	SubtitleStreamInfoList []*BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList `json:"SubtitleStreamInfoList,omitempty" xml:"SubtitleStreamInfoList,omitempty" type:"Repeated"`
	VideoStreamInfoList    []*BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList    `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
}

func (s BatchGetMediasResponseBodyMediaInfosFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediasResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoList) GetAudioStreamInfoList() []*BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	return s.AudioStreamInfoList
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoList) GetFileBasicInfo() *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo {
	return s.FileBasicInfo
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoList) GetSubtitleStreamInfoList() []*BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	return s.SubtitleStreamInfoList
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoList) GetVideoStreamInfoList() []*BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	return s.VideoStreamInfoList
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoList) SetAudioStreamInfoList(v []*BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) *BatchGetMediasResponseBodyMediaInfosFileInfoList {
	s.AudioStreamInfoList = v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) *BatchGetMediasResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoList) SetSubtitleStreamInfoList(v []*BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) *BatchGetMediasResponseBodyMediaInfosFileInfoList {
	s.SubtitleStreamInfoList = v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoList) SetVideoStreamInfoList(v []*BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) *BatchGetMediasResponseBodyMediaInfosFileInfoList {
	s.VideoStreamInfoList = v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoList) Validate() error {
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

type BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList struct {
	// example:
	//
	// 48236800
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// -
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=fe58c6512a1c59524c199577c833abee23f9a16bc549815ca157c46d28fe6ffa
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// example:
	//
	// xx
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// example:
	//
	// H264
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// example:
	//
	// xx
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// example:
	//
	// xx
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// example:
	//
	// xx
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// example:
	//
	// 15
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 32
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
	// 32
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// example:
	//
	// {\\"ApiKey\\":\\"c0358c6e51c1013b446fdeb21a3a5d1c\\",\\"AppId\\":\\"5b347bfb\\",\\"ApiSecret\\":\\"a9872e2342952e248727798f642936b6\\"}
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// example:
	//
	// xx
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// example:
	//
	// 0.01
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// example:
	//
	// 1779850920
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// -
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetChannels() *string {
	return s.Channels
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetBitrate(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetChannelLayout(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetChannels(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetCodecLongName(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetCodecName(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetCodecTag(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetCodecTagString(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetCodecTimeBase(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetDuration(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetFps(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Fps = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetIndex(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetLang(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Lang = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetNumFrames(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetProfile(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Profile = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetSampleFmt(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetSampleRate(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetStartTime(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) SetTimebase(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListAudioStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
	// example:
	//
	// 30
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// 200
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// example.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 191
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
	// jpg
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// example:
	//
	// 416
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 640
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList struct {
	// example:
	//
	// xx
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// example:
	//
	// H264
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// example:
	//
	// xx
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// example:
	//
	// xx
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// example:
	//
	// xx
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// example:
	//
	// 6
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
	// 1767953790
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// xx
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GoString() string {
	return s.String()
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetCodecLongName(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetCodecName(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetCodecTag(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetCodecTagString(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetCodecTimeBase(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetDuration(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.Duration = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetIndex(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.Index = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetLang(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.Lang = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetStartTime(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) SetTimebase(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListSubtitleStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList struct {
	// example:
	//
	// xx
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// example:
	//
	// 23736607
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// xx
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// example:
	//
	// H264
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// example:
	//
	// xx
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// example:
	//
	// xx
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// example:
	//
	// xx
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// example:
	//
	// xx
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// example:
	//
	// 6
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// xx
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// example:
	//
	// xx
	HasBFrames *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	// example:
	//
	// 0
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
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
	// loose
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// xx
	NbFrames *string `json:"Nb_frames,omitempty" xml:"Nb_frames,omitempty"`
	// example:
	//
	// xx
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// example:
	//
	// xx
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// example:
	//
	// {\\"ApiKey\\":\\"c0358c6e51c1013b446fdeb21a3a5d1c\\",\\"AppId\\":\\"5b347bfb\\",\\"ApiSecret\\":\\"a9872e2342952e248727798f642936b6\\"}
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// example:
	//
	// xx
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// example:
	//
	// xx
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// example:
	//
	// 1779850920
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// xx
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// example:
	//
	// 720
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetDar() *string {
	return s.Dar
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetHeight() *string {
	return s.Height
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetLevel() *string {
	return s.Level
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetNbFrames() *string {
	return s.NbFrames
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetRotate() *string {
	return s.Rotate
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetSar() *string {
	return s.Sar
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) GetWidth() *string {
	return s.Width
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetAvgFPS(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.AvgFPS = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetBitrate(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetCodecLongName(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetCodecName(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetCodecTag(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetCodecTagString(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetCodecTimeBase(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetDar(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetDuration(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetFps(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetHasBFrames(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetHeight(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetIndex(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetLang(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetLevel(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetNbFrames(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.NbFrames = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetNumFrames(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetPixFmt(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetProfile(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetRotate(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetSar(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetStartTime(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetTimebase(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) SetWidth(v string) *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosFileInfoListVideoStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type BatchGetMediasResponseBodyMediaInfosMediaBasicInfo struct {
	CategoryId   *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// https://dtlive-bj.oss-cn-beijing.aliyuncs.com/cover/01e1271d-ff4f-4689-9c20-e1df81486859_open_live_cover.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// example:
	//
	// 2020-12-26T04:11:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// sample_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// https://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// example:
	//
	// *****64623a94eca8516569c8f*****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// tag1，tag2
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 2021-01-08T16:52:04Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// []
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
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
	// {}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetCategoryName() *string {
	return s.CategoryName
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetCoverURL() *string {
	return s.CoverURL
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetDescription() *string {
	return s.Description
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetInputURL() *string {
	return s.InputURL
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetMediaTags() *string {
	return s.MediaTags
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetMediaType() *string {
	return s.MediaType
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetSnapshots() *string {
	return s.Snapshots
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetSource() *string {
	return s.Source
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetSpriteImages() *string {
	return s.SpriteImages
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetStatus() *string {
	return s.Status
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) GetUserData() *string {
	return s.UserData
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetCategoryId(v int64) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.CategoryId = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetCategoryName(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.CategoryName = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetSpriteImages(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaBasicInfo) Validate() error {
	return dara.Validate(s)
}

type BatchGetMediasResponseBodyMediaInfosMediaDynamicInfo struct {
	DynamicMetaData *BatchGetMediasResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData `json:"DynamicMetaData,omitempty" xml:"DynamicMetaData,omitempty" type:"Struct"`
}

func (s BatchGetMediasResponseBodyMediaInfosMediaDynamicInfo) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediasResponseBodyMediaInfosMediaDynamicInfo) GoString() string {
	return s.String()
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaDynamicInfo) GetDynamicMetaData() *BatchGetMediasResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData {
	return s.DynamicMetaData
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaDynamicInfo) SetDynamicMetaData(v *BatchGetMediasResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData) *BatchGetMediasResponseBodyMediaInfosMediaDynamicInfo {
	s.DynamicMetaData = v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaDynamicInfo) Validate() error {
	if s.DynamicMetaData != nil {
		if err := s.DynamicMetaData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchGetMediasResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData struct {
	// example:
	//
	// 05D92F7EE52363AE3C95FB23EC56611929613720
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s BatchGetMediasResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediasResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData) GoString() string {
	return s.String()
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData) GetData() *string {
	return s.Data
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData) SetData(v string) *BatchGetMediasResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData {
	s.Data = &v
	return s
}

func (s *BatchGetMediasResponseBodyMediaInfosMediaDynamicInfoDynamicMetaData) Validate() error {
	return dara.Validate(s)
}
