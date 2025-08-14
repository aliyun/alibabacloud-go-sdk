// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaInfoJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaInfoJob(v *SubmitMediaInfoJobResponseBodyMediaInfoJob) *SubmitMediaInfoJobResponseBody
	GetMediaInfoJob() *SubmitMediaInfoJobResponseBodyMediaInfoJob
	SetRequestId(v string) *SubmitMediaInfoJobResponseBody
	GetRequestId() *string
}

type SubmitMediaInfoJobResponseBody struct {
	// MediaInfoJobDTO
	MediaInfoJob *SubmitMediaInfoJobResponseBodyMediaInfoJob `json:"MediaInfoJob,omitempty" xml:"MediaInfoJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitMediaInfoJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBody) GetMediaInfoJob() *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	return s.MediaInfoJob
}

func (s *SubmitMediaInfoJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitMediaInfoJobResponseBody) SetMediaInfoJob(v *SubmitMediaInfoJobResponseBodyMediaInfoJob) *SubmitMediaInfoJobResponseBody {
	s.MediaInfoJob = v
	return s
}

func (s *SubmitMediaInfoJobResponseBody) SetRequestId(v string) *SubmitMediaInfoJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaInfoJobResponseBodyMediaInfoJob struct {
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
	Input *SubmitMediaInfoJobResponseBodyMediaInfoJobInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The job ID.
	//
	// example:
	//
	// ab4802364a2e49208c99efab82dfa8e8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The details of the media information.
	MediaInfoProperty *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty `json:"MediaInfoProperty,omitempty" xml:"MediaInfoProperty,omitempty" type:"Struct"`
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
	ScheduleConfig *SubmitMediaInfoJobResponseBodyMediaInfoJobScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
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

func (s SubmitMediaInfoJobResponseBodyMediaInfoJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJob) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetAsync() *bool {
	return s.Async
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetInput() *SubmitMediaInfoJobResponseBodyMediaInfoJobInput {
	return s.Input
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetMediaInfoProperty() *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty {
	return s.MediaInfoProperty
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetName() *string {
	return s.Name
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetScheduleConfig() *SubmitMediaInfoJobResponseBodyMediaInfoJobScheduleConfig {
	return s.ScheduleConfig
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetStatus() *string {
	return s.Status
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetSubmitResultJson() map[string]interface{} {
	return s.SubmitResultJson
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetAsync(v bool) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.Async = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetFinishTime(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.FinishTime = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetInput(v *SubmitMediaInfoJobResponseBodyMediaInfoJobInput) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.Input = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetJobId(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.JobId = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetMediaInfoProperty(v *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.MediaInfoProperty = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.Name = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetRequestId(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.RequestId = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetScheduleConfig(v *SubmitMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.ScheduleConfig = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetStatus(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.Status = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetSubmitResultJson(v map[string]interface{}) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.SubmitResultJson = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetSubmitTime(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.SubmitTime = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetTriggerSource(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.TriggerSource = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetUserData(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.UserData = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaInfoJobResponseBodyMediaInfoJobInput struct {
	// The media object. If Type is set to OSS, the URL of an OSS object is returned. Both the OSS and HTTP protocols are supported. If Type is set to Media, set this parameter to the ID of a media asset.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the media object. Valid values: OSS and Media. A value of OSS indicates an OSS object. A value of Media indicates a media asset.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobInput) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobInput) GetType() *string {
	return s.Type
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobInput) SetMedia(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobInput {
	s.Media = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobInput) SetType(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobInput {
	s.Type = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobInput) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty struct {
	// The information about the audio stream.
	AudioStreamInfoList []*SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	// The basic file information.
	FileBasicInfo *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	// The information about the video stream.
	VideoStreamInfoList []*SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) GetAudioStreamInfoList() []*SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	return s.AudioStreamInfoList
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) GetFileBasicInfo() *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	return s.FileBasicInfo
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) GetVideoStreamInfoList() []*SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	return s.VideoStreamInfoList
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) SetAudioStreamInfoList(v []*SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty {
	s.AudioStreamInfoList = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) SetFileBasicInfo(v *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty {
	s.FileBasicInfo = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) SetVideoStreamInfoList(v []*SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty {
	s.VideoStreamInfoList = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoProperty) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList struct {
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

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetChannels() *string {
	return s.Channels
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetBitrate(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetChannelLayout(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetChannels(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecLongName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecTag(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecTagString(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetCodecTimeBase(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetDuration(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetIndex(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetLang(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Lang = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetSampleFmt(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetSampleRate(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetStartTime(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) SetTimebase(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyAudioStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo struct {
	// The video bitrate.
	//
	// example:
	//
	// 888.563
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The duration of the video.
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
	// 2b36bd19c13f4145b094c0cad80dbce5
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

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetBitrate(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetDuration(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileSize(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileStatus(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileType(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFileUrl(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetFormatName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetHeight(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Height = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetMediaId(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.MediaId = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetRegion(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Region = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) SetWidth(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo {
	s.Width = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList struct {
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

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetAvgFps() *string {
	return s.AvgFps
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetBitRate() *string {
	return s.BitRate
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetDar() *string {
	return s.Dar
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetHeight() *string {
	return s.Height
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetLevel() *string {
	return s.Level
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetRotate() *string {
	return s.Rotate
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetSar() *string {
	return s.Sar
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetTimeBase() *string {
	return s.TimeBase
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) GetWidth() *string {
	return s.Width
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetAvgFps(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.AvgFps = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetBitRate(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.BitRate = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecLongName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecTag(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecTagString(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetCodecTimeBase(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetDar(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetDuration(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetFps(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetHasBFrames(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetHeight(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetIndex(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetLang(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetLevel(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetNumFrames(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetPixFmt(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetProfile(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetRotate(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetSar(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetStartTime(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetTimeBase(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.TimeBase = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) SetWidth(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMediaInfoPropertyVideoStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaInfoJobResponseBodyMediaInfoJobScheduleConfig struct {
	// The ID of the ApsaraVideo Media Processing (MPS) queue that is used to run the job.
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

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) SetPipelineId(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) SetPriority(v int32) *SubmitMediaInfoJobResponseBodyMediaInfoJobScheduleConfig {
	s.Priority = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobScheduleConfig) Validate() error {
	return dara.Validate(s)
}
