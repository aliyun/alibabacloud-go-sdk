// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaInfoJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaInfoJob(v *GetMediaInfoJobResponseBodyMediaInfoJob) *GetMediaInfoJobResponseBody
	GetMediaInfoJob() *GetMediaInfoJobResponseBodyMediaInfoJob
	SetRequestId(v string) *GetMediaInfoJobResponseBody
	GetRequestId() *string
}

type GetMediaInfoJobResponseBody struct {
	// MediaInfoJobDTO
	MediaInfoJob *GetMediaInfoJobResponseBodyMediaInfoJob `json:"MediaInfoJob,omitempty" xml:"MediaInfoJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaInfoJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaInfoJobResponseBody) GetMediaInfoJob() *GetMediaInfoJobResponseBodyMediaInfoJob {
	return s.MediaInfoJob
}

func (s *GetMediaInfoJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaInfoJobResponseBody) SetMediaInfoJob(v *GetMediaInfoJobResponseBodyMediaInfoJob) *GetMediaInfoJobResponseBody {
	s.MediaInfoJob = v
	return s
}

func (s *GetMediaInfoJobResponseBody) SetRequestId(v string) *GetMediaInfoJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaInfoJobResponseBody) Validate() error {
	if s.MediaInfoJob != nil {
		if err := s.MediaInfoJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaInfoJobResponseBodyMediaInfoJob struct {
	// Indicates whether asynchronous processing was performed.
	//
	// example:
	//
	// true
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The time when the job was complete.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The input of the job.
	Input *GetMediaInfoJobResponseBodyMediaInfoJobInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The job ID.
	//
	// example:
	//
	// ab4802364a2e49208c99efab82dfa8e8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The details of the media information.
	MediaInfoProperty *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty `json:"MediaInfoProperty,omitempty" xml:"MediaInfoProperty,omitempty" type:"Struct"`
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
	ScheduleConfig *GetMediaInfoJobResponseBodyMediaInfoJobScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The state of the job. Valid values: Init (the job is submitted), Success (the job is successful), and Fail (the job failed).
	//
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The job submission information.
	SubmitResultJson map[string]interface{} `json:"SubmitResultJson,omitempty" xml:"SubmitResultJson,omitempty"`
	// The time when the job was submitted.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// The source of the job. Valid values: API, WorkFlow, and Console.
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

func (s GetMediaInfoJobResponseBodyMediaInfoJob) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoJobResponseBodyMediaInfoJob) GoString() string {
	return s.String()
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) GetAsync() *bool {
	return s.Async
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) GetInput() *GetMediaInfoJobResponseBodyMediaInfoJobInput {
	return s.Input
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) GetJobId() *string {
	return s.JobId
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) GetMediaInfoProperty() *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty {
	return s.MediaInfoProperty
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) GetName() *string {
	return s.Name
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) GetScheduleConfig() *GetMediaInfoJobResponseBodyMediaInfoJobScheduleConfig {
	return s.ScheduleConfig
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) GetStatus() *string {
	return s.Status
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) GetSubmitResultJson() map[string]interface{} {
	return s.SubmitResultJson
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) GetUserData() *string {
	return s.UserData
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) SetAsync(v bool) *GetMediaInfoJobResponseBodyMediaInfoJob {
	s.Async = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) SetFinishTime(v string) *GetMediaInfoJobResponseBodyMediaInfoJob {
	s.FinishTime = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) SetInput(v *GetMediaInfoJobResponseBodyMediaInfoJobInput) *GetMediaInfoJobResponseBodyMediaInfoJob {
	s.Input = v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) SetJobId(v string) *GetMediaInfoJobResponseBodyMediaInfoJob {
	s.JobId = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) SetMediaInfoProperty(v *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) *GetMediaInfoJobResponseBodyMediaInfoJob {
	s.MediaInfoProperty = v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) SetName(v string) *GetMediaInfoJobResponseBodyMediaInfoJob {
	s.Name = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) SetRequestId(v string) *GetMediaInfoJobResponseBodyMediaInfoJob {
	s.RequestId = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) SetScheduleConfig(v *GetMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) *GetMediaInfoJobResponseBodyMediaInfoJob {
	s.ScheduleConfig = v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) SetStatus(v string) *GetMediaInfoJobResponseBodyMediaInfoJob {
	s.Status = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) SetSubmitResultJson(v map[string]interface{}) *GetMediaInfoJobResponseBodyMediaInfoJob {
	s.SubmitResultJson = v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) SetSubmitTime(v string) *GetMediaInfoJobResponseBodyMediaInfoJob {
	s.SubmitTime = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) SetTriggerSource(v string) *GetMediaInfoJobResponseBodyMediaInfoJob {
	s.TriggerSource = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) SetUserData(v string) *GetMediaInfoJobResponseBodyMediaInfoJob {
	s.UserData = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJob) Validate() error {
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

type GetMediaInfoJobResponseBodyMediaInfoJobInput struct {
	// The media object. If Type is set to OSS, the URL of an OSS object is returned. Both the OSS and HTTP protocols are supported. If Type is set to Media, the ID of a media asset is returned.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the media object. Valid values:
	//
	// 	- OSS: an Object Storage Service (OSS) object.
	//
	// 	- Media: a media asset.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMediaInfoJobResponseBodyMediaInfoJobInput) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoJobResponseBodyMediaInfoJobInput) GoString() string {
	return s.String()
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobInput) GetMedia() *string {
	return s.Media
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobInput) GetType() *string {
	return s.Type
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobInput) SetMedia(v string) *GetMediaInfoJobResponseBodyMediaInfoJobInput {
	s.Media = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobInput) SetType(v string) *GetMediaInfoJobResponseBodyMediaInfoJobInput {
	s.Type = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobInput) Validate() error {
	return dara.Validate(s)
}

type GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty struct {
	// The information about the audio stream.
	AudioStreamInfoList []*GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	// The basic file information.
	FileBasicInfo *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	// The information about the video stream.
	VideoStreamInfoList []*GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
}

func (s GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) GoString() string {
	return s.String()
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) GetAudioStreamInfoList() []*GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	return s.AudioStreamInfoList
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) GetFileBasicInfo() *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	return s.FileBasicInfo
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) GetVideoStreamInfoList() []*GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	return s.VideoStreamInfoList
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) SetAudioStreamInfoList(v []*GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty {
	s.AudioStreamInfoList = v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) SetFileBasicInfo(v *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty {
	s.FileBasicInfo = v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) SetVideoStreamInfoList(v []*GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty {
	s.VideoStreamInfoList = v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) Validate() error {
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

type GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList struct {
	// The bitrate.
	//
	// example:
	//
	// 0.f
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The sound channel layout.
	//
	// example:
	//
	// stereo
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// The number of sound channels.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The name of the encoding format.
	//
	// example:
	//
	// AAC (Advanced Audio Coding)
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The encoding format.
	//
	// example:
	//
	// aac
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The encoder tag.
	//
	// example:
	//
	// 0x000f
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The name of the encoder tag.
	//
	// example:
	//
	// [15][0][0][0]
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The time base of the encoder.
	//
	// example:
	//
	// 1/44100
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The duration of the stream. Unit: seconds.
	//
	// example:
	//
	// 403.039989
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The sequence number of the stream.
	//
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language of the stream.
	//
	// example:
	//
	// us
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The sample format.
	//
	// example:
	//
	// fltp
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// The sampling rate. Unit: Hz.
	//
	// example:
	//
	// 44100
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// The start time of the stream.
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

func (s GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetChannels() *string {
	return s.Channels
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetBitrate(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetChannelLayout(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetChannels(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecLongName(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecName(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecTag(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecTagString(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecTimeBase(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetDuration(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetIndex(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetLang(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetSampleFmt(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetSampleRate(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetStartTime(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetTimebase(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo struct {
	// The video bitrate.
	//
	// example:
	//
	// 888.563
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The duration of the video. Unit: seconds.
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
	// The file size.
	//
	// example:
	//
	// 31737
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The state of the file.
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
	// The URL of the file.
	//
	// example:
	//
	// http://bucket.oss-cn-shanghai.aliyuncs.com/path/to/file.m3u8
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The name of the video format.
	//
	// example:
	//
	// hls,applehttp
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// The height of the output video.
	//
	// example:
	//
	// 478
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// e520090207114cc7a392d44f0b211574
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The region in which the file resides.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The width of the output video.
	//
	// example:
	//
	// 848
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GoString() string {
	return s.String()
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetBitrate(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetDuration(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileName(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileSize(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileStatus(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileType(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileUrl(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFormatName(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetHeight(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Height = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetMediaId(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.MediaId = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetRegion(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Region = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetWidth(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Width = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList struct {
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
	// The name of the encoding format.
	//
	// example:
	//
	// H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10
	CodecLongName *string `json:"Codec_long_name,omitempty" xml:"Codec_long_name,omitempty"`
	// The encoding format.
	//
	// example:
	//
	// h264
	CodecName *string `json:"Codec_name,omitempty" xml:"Codec_name,omitempty"`
	// The tag of the encoding format.
	//
	// example:
	//
	// 0x001b
	CodecTag *string `json:"Codec_tag,omitempty" xml:"Codec_tag,omitempty"`
	// The tag string of the encoding format.
	//
	// example:
	//
	// [27][0][0][0]
	CodecTagString *string `json:"Codec_tag_string,omitempty" xml:"Codec_tag_string,omitempty"`
	// The time base of the encoder.
	//
	// example:
	//
	// 1/50
	CodecTimeBase *string `json:"Codec_time_base,omitempty" xml:"Codec_time_base,omitempty"`
	// The display aspect ratio.
	//
	// example:
	//
	// 16:9
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// The duration of the file.
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
	// Indicates whether the video stream contains bidirectional frames (B-frames). Valid values:
	//
	// 	- 0: The stream contains no B-frames.
	//
	// 	- 1: The stream contains one B-frame.
	//
	// 	- 2: The stream contains multiple consecutive B-frames.
	//
	// example:
	//
	// 2
	HasBFrames *string `json:"Has_b_frames,omitempty" xml:"Has_b_frames,omitempty"`
	// The height of the output video.
	//
	// example:
	//
	// 478
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The sequence number of the stream.
	//
	// example:
	//
	// 0
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language of the stream.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The codec level.
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
	// The rotation angle of the video image.
	//
	// example:
	//
	// 0
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The aspect ratio of the area from which the sampling points are collected.
	//
	// example:
	//
	// 478:477
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// The start time of the stream.
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
	// The width of the output video.
	//
	// example:
	//
	// 848
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetAvgFps() *string {
	return s.AvgFps
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetBitRate() *string {
	return s.BitRate
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetDar() *string {
	return s.Dar
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetHeight() *string {
	return s.Height
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetLevel() *string {
	return s.Level
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetRotate() *string {
	return s.Rotate
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetSar() *string {
	return s.Sar
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetTimeBase() *string {
	return s.TimeBase
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetWidth() *string {
	return s.Width
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetAvgFps(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.AvgFps = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetBitRate(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.BitRate = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecLongName(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecName(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecTag(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecTagString(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecTimeBase(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetDar(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetDuration(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetFps(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetHasBFrames(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetHeight(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetIndex(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetLang(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetLevel(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetNumFrames(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetPixFmt(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetProfile(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetRotate(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetSar(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetStartTime(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetTimeBase(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.TimeBase = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetWidth(v string) *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetMediaInfoJobResponseBodyMediaInfoJobScheduleConfig struct {
	// The ID of the MPS queue to which the job was submitted.
	//
	// example:
	//
	// e37ebee5d98b4781897f6086e89f9c56
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the job. Valid values: 1 to 10. The greater the value, the higher the priority.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s GetMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) GoString() string {
	return s.String()
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) SetPipelineId(v string) *GetMediaInfoJobResponseBodyMediaInfoJobScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) SetPriority(v int32) *GetMediaInfoJobResponseBodyMediaInfoJobScheduleConfig {
	s.Priority = &v
	return s
}

func (s *GetMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) Validate() error {
	return dara.Validate(s)
}
