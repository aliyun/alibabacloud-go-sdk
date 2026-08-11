// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaInfoJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v []*ListMediaInfoJobsResponseBodyJobs) *ListMediaInfoJobsResponseBody
	GetJobs() []*ListMediaInfoJobsResponseBodyJobs
	SetNextPageToken(v string) *ListMediaInfoJobsResponseBody
	GetNextPageToken() *string
	SetRequestId(v string) *ListMediaInfoJobsResponseBody
	GetRequestId() *string
}

type ListMediaInfoJobsResponseBody struct {
	// The list of media information jobs.
	Jobs []*ListMediaInfoJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The token used to retrieve the next page. Leave this parameter empty for the first request. The token for the next page is returned after the first query.
	//
	// example:
	//
	// 019daf5780f74831b0e1a767c9f1c178
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMediaInfoJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMediaInfoJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaInfoJobsResponseBody) GetJobs() []*ListMediaInfoJobsResponseBodyJobs {
	return s.Jobs
}

func (s *ListMediaInfoJobsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListMediaInfoJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaInfoJobsResponseBody) SetJobs(v []*ListMediaInfoJobsResponseBodyJobs) *ListMediaInfoJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListMediaInfoJobsResponseBody) SetNextPageToken(v string) *ListMediaInfoJobsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListMediaInfoJobsResponseBody) SetRequestId(v string) *ListMediaInfoJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaInfoJobsResponseBody) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMediaInfoJobsResponseBodyJobs struct {
	// Indicates whether the job is processed asynchronously.
	//
	// example:
	//
	// true
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The job completion time. The time is in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The job input.
	Input *ListMediaInfoJobsResponseBodyJobsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The job ID.
	//
	// example:
	//
	// ab4802364a2e49208c99efab82df****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The media information details.
	MediaInfoProperty *ListMediaInfoJobsResponseBodyJobsMediaInfoProperty `json:"MediaInfoProperty,omitempty" xml:"MediaInfoProperty,omitempty" type:"Struct"`
	// The job name.
	//
	// example:
	//
	// job-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4879B9DE-E4B6-19DC-91F5-9D5F4DCE4168
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scheduling information.
	ScheduleConfig *ListMediaInfoJobsResponseBodyJobsScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The task status.
	//
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The job submission information.
	SubmitResultJson map[string]interface{} `json:"SubmitResultJson,omitempty" xml:"SubmitResultJson,omitempty"`
	// The job submission time. The time is in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// The job source.
	//
	// example:
	//
	// API
	TriggerSource *string `json:"TriggerSource,omitempty" xml:"TriggerSource,omitempty"`
	// The user data.
	//
	// example:
	//
	// user-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ListMediaInfoJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListMediaInfoJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListMediaInfoJobsResponseBodyJobs) GetAsync() *bool {
	return s.Async
}

func (s *ListMediaInfoJobsResponseBodyJobs) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListMediaInfoJobsResponseBodyJobs) GetInput() *ListMediaInfoJobsResponseBodyJobsInput {
	return s.Input
}

func (s *ListMediaInfoJobsResponseBodyJobs) GetJobId() *string {
	return s.JobId
}

func (s *ListMediaInfoJobsResponseBodyJobs) GetMediaInfoProperty() *ListMediaInfoJobsResponseBodyJobsMediaInfoProperty {
	return s.MediaInfoProperty
}

func (s *ListMediaInfoJobsResponseBodyJobs) GetName() *string {
	return s.Name
}

func (s *ListMediaInfoJobsResponseBodyJobs) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaInfoJobsResponseBodyJobs) GetScheduleConfig() *ListMediaInfoJobsResponseBodyJobsScheduleConfig {
	return s.ScheduleConfig
}

func (s *ListMediaInfoJobsResponseBodyJobs) GetStatus() *string {
	return s.Status
}

func (s *ListMediaInfoJobsResponseBodyJobs) GetSubmitResultJson() map[string]interface{} {
	return s.SubmitResultJson
}

func (s *ListMediaInfoJobsResponseBodyJobs) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *ListMediaInfoJobsResponseBodyJobs) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *ListMediaInfoJobsResponseBodyJobs) GetUserData() *string {
	return s.UserData
}

func (s *ListMediaInfoJobsResponseBodyJobs) SetAsync(v bool) *ListMediaInfoJobsResponseBodyJobs {
	s.Async = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobs) SetFinishTime(v string) *ListMediaInfoJobsResponseBodyJobs {
	s.FinishTime = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobs) SetInput(v *ListMediaInfoJobsResponseBodyJobsInput) *ListMediaInfoJobsResponseBodyJobs {
	s.Input = v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobs) SetJobId(v string) *ListMediaInfoJobsResponseBodyJobs {
	s.JobId = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobs) SetMediaInfoProperty(v *ListMediaInfoJobsResponseBodyJobsMediaInfoProperty) *ListMediaInfoJobsResponseBodyJobs {
	s.MediaInfoProperty = v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobs) SetName(v string) *ListMediaInfoJobsResponseBodyJobs {
	s.Name = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobs) SetRequestId(v string) *ListMediaInfoJobsResponseBodyJobs {
	s.RequestId = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobs) SetScheduleConfig(v *ListMediaInfoJobsResponseBodyJobsScheduleConfig) *ListMediaInfoJobsResponseBodyJobs {
	s.ScheduleConfig = v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobs) SetStatus(v string) *ListMediaInfoJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobs) SetSubmitResultJson(v map[string]interface{}) *ListMediaInfoJobsResponseBodyJobs {
	s.SubmitResultJson = v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobs) SetSubmitTime(v string) *ListMediaInfoJobsResponseBodyJobs {
	s.SubmitTime = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobs) SetTriggerSource(v string) *ListMediaInfoJobsResponseBodyJobs {
	s.TriggerSource = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobs) SetUserData(v string) *ListMediaInfoJobsResponseBodyJobs {
	s.UserData = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobs) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.MediaInfoProperty != nil {
		if err := s.MediaInfoProperty.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleConfig != nil {
		if err := s.ScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMediaInfoJobsResponseBodyJobsInput struct {
	// The media value:
	//
	// - If type is set to OSS, the value is a URL that supports the OSS protocol and HTTP protocol.
	//
	// - If type is set to Media, the value is a media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The media object type.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMediaInfoJobsResponseBodyJobsInput) String() string {
	return dara.Prettify(s)
}

func (s ListMediaInfoJobsResponseBodyJobsInput) GoString() string {
	return s.String()
}

func (s *ListMediaInfoJobsResponseBodyJobsInput) GetMedia() *string {
	return s.Media
}

func (s *ListMediaInfoJobsResponseBodyJobsInput) GetType() *string {
	return s.Type
}

func (s *ListMediaInfoJobsResponseBodyJobsInput) SetMedia(v string) *ListMediaInfoJobsResponseBodyJobsInput {
	s.Media = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsInput) SetType(v string) *ListMediaInfoJobsResponseBodyJobsInput {
	s.Type = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsInput) Validate() error {
	return dara.Validate(s)
}

type ListMediaInfoJobsResponseBodyJobsMediaInfoProperty struct {
	// The audio stream information.
	AudioStreamInfoList []*ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	// The basic file information.
	FileBasicInfo *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	// The video stream information.
	VideoStreamInfoList []*ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
}

func (s ListMediaInfoJobsResponseBodyJobsMediaInfoProperty) String() string {
	return dara.Prettify(s)
}

func (s ListMediaInfoJobsResponseBodyJobsMediaInfoProperty) GoString() string {
	return s.String()
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoProperty) GetAudioStreamInfoList() []*ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	return s.AudioStreamInfoList
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoProperty) GetFileBasicInfo() *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo {
	return s.FileBasicInfo
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoProperty) GetVideoStreamInfoList() []*ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	return s.VideoStreamInfoList
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoProperty) SetAudioStreamInfoList(v []*ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) *ListMediaInfoJobsResponseBodyJobsMediaInfoProperty {
	s.AudioStreamInfoList = v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoProperty) SetFileBasicInfo(v *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) *ListMediaInfoJobsResponseBodyJobsMediaInfoProperty {
	s.FileBasicInfo = v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoProperty) SetVideoStreamInfoList(v []*ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) *ListMediaInfoJobsResponseBodyJobsMediaInfoProperty {
	s.VideoStreamInfoList = v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoProperty) Validate() error {
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

type ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList struct {
	// The bitrate.
	//
	// example:
	//
	// 0.f
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The channel layout.
	//
	// example:
	//
	// stereo
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// The number of audio channels.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The codec format name.
	//
	// example:
	//
	// AAC (Advanced Audio Coding)
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The codec format.
	//
	// example:
	//
	// aac
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The codec tag.
	//
	// example:
	//
	// 0x000f
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The codec tag string.
	//
	// example:
	//
	// [15][0][0][0]
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The codec time base.
	//
	// example:
	//
	// 1/44100
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The duration. Unit: seconds.
	//
	// example:
	//
	// 403.039989
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The stream index.
	//
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language.
	//
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The sample format.
	//
	// example:
	//
	// fltp
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// The sample rate. Unit: Hz.
	//
	// example:
	//
	// 44100
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1.473556
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	//
	// example:
	//
	// 1/90000
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetChannels() *string {
	return s.Channels
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetBitrate(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetChannelLayout(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetChannels(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetCodecLongName(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetCodecName(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetCodecTag(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetCodecTagString(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetCodecTimeBase(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetDuration(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetIndex(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetLang(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.Lang = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetSampleFmt(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetSampleRate(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetStartTime(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) SetTimebase(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyAudioStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo struct {
	// The video bitrate.
	//
	// example:
	//
	// 888.563
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The video duration. Unit: seconds.
	//
	// example:
	//
	// 403.039999
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The file name.
	//
	// example:
	//
	// file.m3u8
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file size. Unit: bytes.
	//
	// example:
	//
	// 31737
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The file status.
	//
	// example:
	//
	// Normal
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// The file type. Valid values:
	//
	// example:
	//
	// source_file
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The file URL.
	//
	// example:
	//
	// http://bucket.oss-cn-shanghai.aliyuncs.com/path/to/file.m3u8
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The video format name.
	//
	// example:
	//
	// hls,applehttp
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// The height.
	//
	// example:
	//
	// 478
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The media asset ID.
	//
	// example:
	//
	// 4765337007f571edbfdf81848c01****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
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
	// 848
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) GoString() string {
	return s.String()
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) SetBitrate(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) SetDuration(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) SetFileName(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) SetFileSize(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) SetFileStatus(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) SetFileType(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) SetFileUrl(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) SetFormatName(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) SetHeight(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo {
	s.Height = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) SetMediaId(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo {
	s.MediaId = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) SetRegion(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo {
	s.Region = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) SetWidth(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo {
	s.Width = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList struct {
	// The average frame rate.
	//
	// example:
	//
	// 25.0
	AvgFps *string `json:"Avg_fps,omitempty" xml:"Avg_fps,omitempty"`
	// The bitrate.
	//
	// example:
	//
	// 888.563
	BitRate *string `json:"Bit_rate,omitempty" xml:"Bit_rate,omitempty"`
	// The codec format name.
	//
	// example:
	//
	// H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10
	CodecLongName *string `json:"Codec_long_name,omitempty" xml:"Codec_long_name,omitempty"`
	// The codec format.
	//
	// example:
	//
	// h264
	CodecName *string `json:"Codec_name,omitempty" xml:"Codec_name,omitempty"`
	// The codec tag.
	//
	// example:
	//
	// 0x001b
	CodecTag *string `json:"Codec_tag,omitempty" xml:"Codec_tag,omitempty"`
	// The codec tag string.
	//
	// example:
	//
	// [27][0][0][0]
	CodecTagString *string `json:"Codec_tag_string,omitempty" xml:"Codec_tag_string,omitempty"`
	// The codec time base.
	//
	// example:
	//
	// 1/50
	CodecTimeBase *string `json:"Codec_time_base,omitempty" xml:"Codec_time_base,omitempty"`
	// The display aspect ratio (DAR).
	//
	// example:
	//
	// 16:9
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// The duration. Unit: seconds.
	//
	// example:
	//
	// 403.039989
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The frame rate.
	//
	// example:
	//
	// 25.0
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// Indicates whether B-frames exist. Valid values:
	//
	// - 0: no B-frames.
	//
	// - 1: one B-frame.
	//
	// - 2: multiple consecutive B-frames.
	//
	// example:
	//
	// 2
	HasBFrames *string `json:"Has_b_frames,omitempty" xml:"Has_b_frames,omitempty"`
	// The height.
	//
	// example:
	//
	// 478
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The stream index.
	//
	// example:
	//
	// 0
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language.
	//
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The encoding level.
	//
	// example:
	//
	// 31
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The total number of frames.
	//
	// example:
	//
	// 10040
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// The pixel format.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The encoder profile.
	//
	// example:
	//
	// High
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The video rotation angle.
	//
	// example:
	//
	// 0
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The sample aspect ratio (SAR).
	//
	// example:
	//
	// 478:477
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1.473556
	StartTime *string `json:"Start_time,omitempty" xml:"Start_time,omitempty"`
	// The time base.
	//
	// example:
	//
	// 1/90000
	TimeBase *string `json:"Time_base,omitempty" xml:"Time_base,omitempty"`
	// The width.
	//
	// example:
	//
	// 848
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetAvgFps() *string {
	return s.AvgFps
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetBitRate() *string {
	return s.BitRate
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetDar() *string {
	return s.Dar
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetHeight() *string {
	return s.Height
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetLevel() *string {
	return s.Level
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetRotate() *string {
	return s.Rotate
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetSar() *string {
	return s.Sar
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetTimeBase() *string {
	return s.TimeBase
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) GetWidth() *string {
	return s.Width
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetAvgFps(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.AvgFps = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetBitRate(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.BitRate = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetCodecLongName(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetCodecName(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetCodecTag(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetCodecTagString(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetCodecTimeBase(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetDar(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetDuration(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetFps(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetHasBFrames(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetHeight(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetIndex(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetLang(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetLevel(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetNumFrames(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetPixFmt(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetProfile(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetRotate(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetSar(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetStartTime(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetTimeBase(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.TimeBase = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) SetWidth(v string) *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsMediaInfoPropertyVideoStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type ListMediaInfoJobsResponseBodyJobsScheduleConfig struct {
	// The pipeline ID.
	//
	// example:
	//
	// e37ebee5d98b4781897f6086e89f****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The job priority. A larger value indicates a higher priority. Valid values: 1 to 10.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s ListMediaInfoJobsResponseBodyJobsScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s ListMediaInfoJobsResponseBodyJobsScheduleConfig) GoString() string {
	return s.String()
}

func (s *ListMediaInfoJobsResponseBodyJobsScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListMediaInfoJobsResponseBodyJobsScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *ListMediaInfoJobsResponseBodyJobsScheduleConfig) SetPipelineId(v string) *ListMediaInfoJobsResponseBodyJobsScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsScheduleConfig) SetPriority(v int32) *ListMediaInfoJobsResponseBodyJobsScheduleConfig {
	s.Priority = &v
	return s
}

func (s *ListMediaInfoJobsResponseBodyJobsScheduleConfig) Validate() error {
	return dara.Validate(s)
}
