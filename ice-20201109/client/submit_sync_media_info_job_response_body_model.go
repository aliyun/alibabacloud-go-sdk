// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSyncMediaInfoJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaInfoJob(v *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) *SubmitSyncMediaInfoJobResponseBody
	GetMediaInfoJob() *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob
	SetRequestId(v string) *SubmitSyncMediaInfoJobResponseBody
	GetRequestId() *string
}

type SubmitSyncMediaInfoJobResponseBody struct {
	// MediaInfoJobDTO
	MediaInfoJob *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob `json:"MediaInfoJob,omitempty" xml:"MediaInfoJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSyncMediaInfoJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSyncMediaInfoJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSyncMediaInfoJobResponseBody) GetMediaInfoJob() *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob {
	return s.MediaInfoJob
}

func (s *SubmitSyncMediaInfoJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSyncMediaInfoJobResponseBody) SetMediaInfoJob(v *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) *SubmitSyncMediaInfoJobResponseBody {
	s.MediaInfoJob = v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBody) SetRequestId(v string) *SubmitSyncMediaInfoJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitSyncMediaInfoJobResponseBodyMediaInfoJob struct {
	// Indicates whether asynchronous processing was performed.
	//
	// example:
	//
	// true
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The time when the job was complete. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The input of the job.
	Input *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The job ID.
	//
	// example:
	//
	// ab4802364a2e49208c99efab82dfa8e8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The details of the media information.
	MediaInfoProperty *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty `json:"MediaInfoProperty,omitempty" xml:"MediaInfoProperty,omitempty" type:"Struct"`
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
	ScheduleConfig *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The state of the job. Valid values: Init (the job is submitted), Success (the job is successful), and Fail (the job failed).
	//
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The job submission information.
	//
	// example:
	//
	// {}
	SubmitResultJson map[string]interface{} `json:"SubmitResultJson,omitempty" xml:"SubmitResultJson,omitempty"`
	// The time when the job was submitted. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
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

func (s SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) GoString() string {
	return s.String()
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) GetAsync() *bool {
	return s.Async
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) GetInput() *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobInput {
	return s.Input
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) GetMediaInfoProperty() *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty {
	return s.MediaInfoProperty
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) GetName() *string {
	return s.Name
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) GetScheduleConfig() *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobScheduleConfig {
	return s.ScheduleConfig
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) GetStatus() *string {
	return s.Status
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) GetSubmitResultJson() map[string]interface{} {
	return s.SubmitResultJson
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) SetAsync(v bool) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob {
	s.Async = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) SetFinishTime(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob {
	s.FinishTime = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) SetInput(v *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobInput) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob {
	s.Input = v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) SetJobId(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob {
	s.JobId = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) SetMediaInfoProperty(v *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob {
	s.MediaInfoProperty = v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) SetName(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob {
	s.Name = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) SetRequestId(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob {
	s.RequestId = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) SetScheduleConfig(v *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob {
	s.ScheduleConfig = v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) SetStatus(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob {
	s.Status = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) SetSubmitResultJson(v map[string]interface{}) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob {
	s.SubmitResultJson = v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) SetSubmitTime(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob {
	s.SubmitTime = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) SetTriggerSource(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob {
	s.TriggerSource = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) SetUserData(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob {
	s.UserData = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJob) Validate() error {
	return dara.Validate(s)
}

type SubmitSyncMediaInfoJobResponseBodyMediaInfoJobInput struct {
	// The media object. If Type is set to OSS, set this parameter to the URL of an OSS object. Both the OSS and HTTP protocols are supported. If Type is set to Media, set this parameter to the ID of a media asset.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the media object. Valid values:
	//
	// 	- OSS: an OSS object.
	//
	// 	- Media: a media asset.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitSyncMediaInfoJobResponseBodyMediaInfoJobInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitSyncMediaInfoJobResponseBodyMediaInfoJobInput) GoString() string {
	return s.String()
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobInput) GetType() *string {
	return s.Type
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobInput) SetMedia(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobInput {
	s.Media = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobInput) SetType(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobInput {
	s.Type = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobInput) Validate() error {
	return dara.Validate(s)
}

type SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty struct {
	// The information about the audio stream.
	AudioStreamInfoList []*SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	// The basic file information.
	FileBasicInfo *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	// The information about the video stream.
	VideoStreamInfoList []*SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
}

func (s SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) String() string {
	return dara.Prettify(s)
}

func (s SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) GoString() string {
	return s.String()
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) GetAudioStreamInfoList() []*SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	return s.AudioStreamInfoList
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) GetFileBasicInfo() *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	return s.FileBasicInfo
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) GetVideoStreamInfoList() []*SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	return s.VideoStreamInfoList
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) SetAudioStreamInfoList(v []*SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty {
	s.AudioStreamInfoList = v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) SetFileBasicInfo(v *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty {
	s.FileBasicInfo = v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) SetVideoStreamInfoList(v []*SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty {
	s.VideoStreamInfoList = v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) Validate() error {
	return dara.Validate(s)
}

type SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList struct {
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
	// The duration of the file.
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

func (s SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetChannels() *string {
	return s.Channels
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetBitrate(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetChannelLayout(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetChannels(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecLongName(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecName(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecTag(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecTagString(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecTimeBase(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetDuration(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetIndex(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetLang(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Lang = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetSampleFmt(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetSampleRate(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetStartTime(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetTimebase(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo struct {
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
	// The file size. Unit: bytes.
	//
	// example:
	//
	// 31737
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The state of the file. Valid values:
	//
	// 	- Normal
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
	// 999e68259c924f52a6be603cbb3f91cc
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

func (s SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GoString() string {
	return s.String()
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetBitrate(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetDuration(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileName(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileSize(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileStatus(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileType(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileUrl(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFormatName(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetHeight(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Height = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetMediaId(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.MediaId = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetRegion(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Region = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetWidth(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Width = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList struct {
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

func (s SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetAvgFps() *string {
	return s.AvgFps
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetBitRate() *string {
	return s.BitRate
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetDar() *string {
	return s.Dar
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetHeight() *string {
	return s.Height
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetLevel() *string {
	return s.Level
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetRotate() *string {
	return s.Rotate
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetSar() *string {
	return s.Sar
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetTimeBase() *string {
	return s.TimeBase
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetWidth() *string {
	return s.Width
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetAvgFps(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.AvgFps = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetBitRate(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.BitRate = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecLongName(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecName(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecTag(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecTagString(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecTimeBase(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetDar(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetDuration(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetFps(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetHasBFrames(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetHeight(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetIndex(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetLang(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetLevel(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetNumFrames(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetPixFmt(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetProfile(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetRotate(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetSar(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetStartTime(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetTimeBase(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.TimeBase = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetWidth(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type SubmitSyncMediaInfoJobResponseBodyMediaInfoJobScheduleConfig struct {
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

func (s SubmitSyncMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSyncMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) SetPipelineId(v string) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) SetPriority(v int32) *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobScheduleConfig {
	s.Priority = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) Validate() error {
	return dara.Validate(s)
}
