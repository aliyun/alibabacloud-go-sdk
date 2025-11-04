// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertOutputDetailFileMeta interface {
	dara.Model
	String() string
	GoString() string
	SetAudioStreamInfoList(v []*MediaConvertOutputDetailFileMetaAudioStreamInfoList) *MediaConvertOutputDetailFileMeta
	GetAudioStreamInfoList() []*MediaConvertOutputDetailFileMetaAudioStreamInfoList
	SetFileBasicInfo(v *MediaConvertOutputDetailFileMetaFileBasicInfo) *MediaConvertOutputDetailFileMeta
	GetFileBasicInfo() *MediaConvertOutputDetailFileMetaFileBasicInfo
	SetVideoStreamInfoList(v []*MediaConvertOutputDetailFileMetaVideoStreamInfoList) *MediaConvertOutputDetailFileMeta
	GetVideoStreamInfoList() []*MediaConvertOutputDetailFileMetaVideoStreamInfoList
}

type MediaConvertOutputDetailFileMeta struct {
	AudioStreamInfoList []*MediaConvertOutputDetailFileMetaAudioStreamInfoList `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	FileBasicInfo       *MediaConvertOutputDetailFileMetaFileBasicInfo         `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	VideoStreamInfoList []*MediaConvertOutputDetailFileMetaVideoStreamInfoList `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
}

func (s MediaConvertOutputDetailFileMeta) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertOutputDetailFileMeta) GoString() string {
	return s.String()
}

func (s *MediaConvertOutputDetailFileMeta) GetAudioStreamInfoList() []*MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	return s.AudioStreamInfoList
}

func (s *MediaConvertOutputDetailFileMeta) GetFileBasicInfo() *MediaConvertOutputDetailFileMetaFileBasicInfo {
	return s.FileBasicInfo
}

func (s *MediaConvertOutputDetailFileMeta) GetVideoStreamInfoList() []*MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	return s.VideoStreamInfoList
}

func (s *MediaConvertOutputDetailFileMeta) SetAudioStreamInfoList(v []*MediaConvertOutputDetailFileMetaAudioStreamInfoList) *MediaConvertOutputDetailFileMeta {
	s.AudioStreamInfoList = v
	return s
}

func (s *MediaConvertOutputDetailFileMeta) SetFileBasicInfo(v *MediaConvertOutputDetailFileMetaFileBasicInfo) *MediaConvertOutputDetailFileMeta {
	s.FileBasicInfo = v
	return s
}

func (s *MediaConvertOutputDetailFileMeta) SetVideoStreamInfoList(v []*MediaConvertOutputDetailFileMetaVideoStreamInfoList) *MediaConvertOutputDetailFileMeta {
	s.VideoStreamInfoList = v
	return s
}

func (s *MediaConvertOutputDetailFileMeta) Validate() error {
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

type MediaConvertOutputDetailFileMetaAudioStreamInfoList struct {
	// example:
	//
	// 0.f
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
	// 0x000f
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// example:
	//
	// [15][0][0][0]
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// example:
	//
	// 1/44100
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// example:
	//
	// 403.039989
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	// 1.473556
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1/90000
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s MediaConvertOutputDetailFileMetaAudioStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertOutputDetailFileMetaAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetChannels() *string {
	return s.Channels
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetBitrate(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetChannelLayout(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetChannels(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetCodecLongName(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetCodecName(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetCodecTag(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetCodecTagString(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetCodecTimeBase(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetDuration(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetIndex(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetLang(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.Lang = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetSampleFmt(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetSampleRate(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetStartTime(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) SetTimebase(v string) *MediaConvertOutputDetailFileMetaAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaAudioStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type MediaConvertOutputDetailFileMetaFileBasicInfo struct {
	// example:
	//
	// 888.563
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// 403.039999
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// file.m3u8
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 31737
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
	// http://bucket.oss-cn-shanghai.aliyuncs.com/path/to/file.m3u8
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// hls,applehttp
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// example:
	//
	// 478
	Height  *string `json:"Height,omitempty" xml:"Height,omitempty"`
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 848
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s MediaConvertOutputDetailFileMetaFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertOutputDetailFileMetaFileBasicInfo) GoString() string {
	return s.String()
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) SetBitrate(v string) *MediaConvertOutputDetailFileMetaFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) SetDuration(v string) *MediaConvertOutputDetailFileMetaFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) SetFileName(v string) *MediaConvertOutputDetailFileMetaFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) SetFileSize(v string) *MediaConvertOutputDetailFileMetaFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) SetFileStatus(v string) *MediaConvertOutputDetailFileMetaFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) SetFileType(v string) *MediaConvertOutputDetailFileMetaFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) SetFileUrl(v string) *MediaConvertOutputDetailFileMetaFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) SetFormatName(v string) *MediaConvertOutputDetailFileMetaFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) SetHeight(v string) *MediaConvertOutputDetailFileMetaFileBasicInfo {
	s.Height = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) SetMediaId(v string) *MediaConvertOutputDetailFileMetaFileBasicInfo {
	s.MediaId = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) SetRegion(v string) *MediaConvertOutputDetailFileMetaFileBasicInfo {
	s.Region = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) SetWidth(v string) *MediaConvertOutputDetailFileMetaFileBasicInfo {
	s.Width = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type MediaConvertOutputDetailFileMetaVideoStreamInfoList struct {
	// example:
	//
	// 25.0
	AvgFps *string `json:"Avg_fps,omitempty" xml:"Avg_fps,omitempty"`
	// example:
	//
	// 888.563
	BitRate *string `json:"Bit_rate,omitempty" xml:"Bit_rate,omitempty"`
	// example:
	//
	// H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10
	CodecLongName *string `json:"Codec_long_name,omitempty" xml:"Codec_long_name,omitempty"`
	// example:
	//
	// h264
	CodecName *string `json:"Codec_name,omitempty" xml:"Codec_name,omitempty"`
	// example:
	//
	// 0x001b
	CodecTag *string `json:"Codec_tag,omitempty" xml:"Codec_tag,omitempty"`
	// example:
	//
	// [27][0][0][0]
	CodecTagString *string `json:"Codec_tag_string,omitempty" xml:"Codec_tag_string,omitempty"`
	CodecTimeBase  *string `json:"Codec_time_base,omitempty" xml:"Codec_time_base,omitempty"`
	// example:
	//
	// 16:9
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// example:
	//
	// 403.039989
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 25.0
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// example:
	//
	// 2
	HasBFrames *string `json:"Has_b_frames,omitempty" xml:"Has_b_frames,omitempty"`
	// example:
	//
	// 478
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 31
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 10040
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
	// 478:477
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// example:
	//
	// 1.473556
	StartTime *string `json:"Start_time,omitempty" xml:"Start_time,omitempty"`
	// example:
	//
	// 1/90000
	TimeBase *string `json:"Time_base,omitempty" xml:"Time_base,omitempty"`
	// example:
	//
	// 848
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s MediaConvertOutputDetailFileMetaVideoStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertOutputDetailFileMetaVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetAvgFps() *string {
	return s.AvgFps
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetBitRate() *string {
	return s.BitRate
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetDar() *string {
	return s.Dar
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetHeight() *string {
	return s.Height
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetLevel() *string {
	return s.Level
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetRotate() *string {
	return s.Rotate
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetSar() *string {
	return s.Sar
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetTimeBase() *string {
	return s.TimeBase
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) GetWidth() *string {
	return s.Width
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetAvgFps(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.AvgFps = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetBitRate(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.BitRate = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetCodecLongName(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetCodecName(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetCodecTag(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetCodecTagString(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetCodecTimeBase(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetDar(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetDuration(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetFps(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetHasBFrames(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetHeight(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetIndex(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetLang(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetLevel(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetNumFrames(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetPixFmt(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetProfile(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetRotate(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetSar(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetStartTime(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetTimeBase(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.TimeBase = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) SetWidth(v string) *MediaConvertOutputDetailFileMetaVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *MediaConvertOutputDetailFileMetaVideoStreamInfoList) Validate() error {
	return dara.Validate(s)
}
