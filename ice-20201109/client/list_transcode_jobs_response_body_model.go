// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTranscodeJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v []*ListTranscodeJobsResponseBodyJobs) *ListTranscodeJobsResponseBody
	GetJobs() []*ListTranscodeJobsResponseBodyJobs
	SetNextPageToken(v string) *ListTranscodeJobsResponseBody
	GetNextPageToken() *string
	SetRequestId(v string) *ListTranscodeJobsResponseBody
	GetRequestId() *string
}

type ListTranscodeJobsResponseBody struct {
	// The list of jobs.
	Jobs []*ListTranscodeJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. The token of the next page is returned after you call this operation for the first time.
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

func (s ListTranscodeJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBody) GetJobs() []*ListTranscodeJobsResponseBodyJobs {
	return s.Jobs
}

func (s *ListTranscodeJobsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListTranscodeJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTranscodeJobsResponseBody) SetJobs(v []*ListTranscodeJobsResponseBodyJobs) *ListTranscodeJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListTranscodeJobsResponseBody) SetNextPageToken(v string) *ListTranscodeJobsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListTranscodeJobsResponseBody) SetRequestId(v string) *ListTranscodeJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTranscodeJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobs struct {
	// The time when the job was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the job was complete. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The input group of the job. An input of a single file indicates a transcoding job. An input of multiple files indicates an audio and video stream merge job.
	InputGroup []*ListTranscodeJobsResponseBodyJobsInputGroup `json:"InputGroup,omitempty" xml:"InputGroup,omitempty" type:"Repeated"`
	// The number of subjobs.
	//
	// example:
	//
	// 1
	JobCount *int32 `json:"JobCount,omitempty" xml:"JobCount,omitempty"`
	// The job name.
	//
	// example:
	//
	// transcode-job
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output group of the job.
	OutputGroup []*ListTranscodeJobsResponseBodyJobsOutputGroup `json:"OutputGroup,omitempty" xml:"OutputGroup,omitempty" type:"Repeated"`
	// The main job ID.
	//
	// example:
	//
	// 8b2198504dd340b7b3c9842a74fc9baa
	ParentJobId *string `json:"ParentJobId,omitempty" xml:"ParentJobId,omitempty"`
	// The completion percentage of the job.
	//
	// example:
	//
	// 0
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The ID of the request that submitted the job.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scheduling configuration of the job.
	ScheduleConfig *ListTranscodeJobsResponseBodyJobsScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The state of the job.
	//
	// 	- Success: At least one of the subjobs is successful.
	//
	// 	- Fail: All subjobs failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the job was submitted. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// The source of the job. Valid values:
	//
	// 	- API
	//
	// 	- WorkFlow
	//
	// 	- Console
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

func (s ListTranscodeJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTranscodeJobsResponseBodyJobs) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListTranscodeJobsResponseBodyJobs) GetInputGroup() []*ListTranscodeJobsResponseBodyJobsInputGroup {
	return s.InputGroup
}

func (s *ListTranscodeJobsResponseBodyJobs) GetJobCount() *int32 {
	return s.JobCount
}

func (s *ListTranscodeJobsResponseBodyJobs) GetName() *string {
	return s.Name
}

func (s *ListTranscodeJobsResponseBodyJobs) GetOutputGroup() []*ListTranscodeJobsResponseBodyJobsOutputGroup {
	return s.OutputGroup
}

func (s *ListTranscodeJobsResponseBodyJobs) GetParentJobId() *string {
	return s.ParentJobId
}

func (s *ListTranscodeJobsResponseBodyJobs) GetPercent() *int32 {
	return s.Percent
}

func (s *ListTranscodeJobsResponseBodyJobs) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTranscodeJobsResponseBodyJobs) GetScheduleConfig() *ListTranscodeJobsResponseBodyJobsScheduleConfig {
	return s.ScheduleConfig
}

func (s *ListTranscodeJobsResponseBodyJobs) GetStatus() *string {
	return s.Status
}

func (s *ListTranscodeJobsResponseBodyJobs) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *ListTranscodeJobsResponseBodyJobs) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *ListTranscodeJobsResponseBodyJobs) GetUserData() *string {
	return s.UserData
}

func (s *ListTranscodeJobsResponseBodyJobs) SetCreateTime(v string) *ListTranscodeJobsResponseBodyJobs {
	s.CreateTime = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobs) SetFinishTime(v string) *ListTranscodeJobsResponseBodyJobs {
	s.FinishTime = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobs) SetInputGroup(v []*ListTranscodeJobsResponseBodyJobsInputGroup) *ListTranscodeJobsResponseBodyJobs {
	s.InputGroup = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobs) SetJobCount(v int32) *ListTranscodeJobsResponseBodyJobs {
	s.JobCount = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobs) SetName(v string) *ListTranscodeJobsResponseBodyJobs {
	s.Name = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobs) SetOutputGroup(v []*ListTranscodeJobsResponseBodyJobsOutputGroup) *ListTranscodeJobsResponseBodyJobs {
	s.OutputGroup = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobs) SetParentJobId(v string) *ListTranscodeJobsResponseBodyJobs {
	s.ParentJobId = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobs) SetPercent(v int32) *ListTranscodeJobsResponseBodyJobs {
	s.Percent = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobs) SetRequestId(v string) *ListTranscodeJobsResponseBodyJobs {
	s.RequestId = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobs) SetScheduleConfig(v *ListTranscodeJobsResponseBodyJobsScheduleConfig) *ListTranscodeJobsResponseBodyJobs {
	s.ScheduleConfig = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobs) SetStatus(v string) *ListTranscodeJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobs) SetSubmitTime(v string) *ListTranscodeJobsResponseBodyJobs {
	s.SubmitTime = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobs) SetTriggerSource(v string) *ListTranscodeJobsResponseBodyJobs {
	s.TriggerSource = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobs) SetUserData(v string) *ListTranscodeJobsResponseBodyJobs {
	s.UserData = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobs) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsInputGroup struct {
	// The URL of the media asset. This parameter is specified only when the media asset is transcoded.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	InputUrl *string `json:"InputUrl,omitempty" xml:"InputUrl,omitempty"`
	// The media object.
	//
	// 	- If Type is set to OSS, the URL of an OSS object is returned. Both the OSS and HTTP protocols are supported.
	//
	// 	- If Type is set to Media, the ID of a media asset is returned.
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

func (s ListTranscodeJobsResponseBodyJobsInputGroup) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsInputGroup) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsInputGroup) GetInputUrl() *string {
	return s.InputUrl
}

func (s *ListTranscodeJobsResponseBodyJobsInputGroup) GetMedia() *string {
	return s.Media
}

func (s *ListTranscodeJobsResponseBodyJobsInputGroup) GetType() *string {
	return s.Type
}

func (s *ListTranscodeJobsResponseBodyJobsInputGroup) SetInputUrl(v string) *ListTranscodeJobsResponseBodyJobsInputGroup {
	s.InputUrl = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsInputGroup) SetMedia(v string) *ListTranscodeJobsResponseBodyJobsInputGroup {
	s.Media = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsInputGroup) SetType(v string) *ListTranscodeJobsResponseBodyJobsInputGroup {
	s.Type = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsInputGroup) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroup struct {
	// The output file configuration.
	Output *ListTranscodeJobsResponseBodyJobsOutputGroupOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The job processing configuration.
	ProcessConfig *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig `json:"ProcessConfig,omitempty" xml:"ProcessConfig,omitempty" type:"Struct"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroup) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroup) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroup) GetOutput() *ListTranscodeJobsResponseBodyJobsOutputGroupOutput {
	return s.Output
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroup) GetProcessConfig() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig {
	return s.ProcessConfig
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroup) SetOutput(v *ListTranscodeJobsResponseBodyJobsOutputGroupOutput) *ListTranscodeJobsResponseBodyJobsOutputGroup {
	s.Output = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroup) SetProcessConfig(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) *ListTranscodeJobsResponseBodyJobsOutputGroup {
	s.ProcessConfig = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroup) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupOutput struct {
	// The media object. If Type is set to OSS, the URL of an OSS object is returned. Both the OSS and HTTP protocols are supported. If Type is set to Media, the ID of a media asset is returned.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The URL of the transcoded output stream. This parameter is required only when the output is a media asset.
	//
	// example:
	//
	// oss://bucket/path/to/{MediaId}/{JobId}.mp4
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
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

func (s ListTranscodeJobsResponseBodyJobsOutputGroupOutput) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupOutput) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupOutput) GetMedia() *string {
	return s.Media
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupOutput) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupOutput) GetType() *string {
	return s.Type
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupOutput) SetMedia(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupOutput {
	s.Media = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupOutput) SetOutputUrl(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupOutput {
	s.OutputUrl = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupOutput) SetType(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupOutput {
	s.Type = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupOutput) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig struct {
	// The multi-input stream merge configuration.
	CombineConfigs []*ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs `json:"CombineConfigs,omitempty" xml:"CombineConfigs,omitempty" type:"Repeated"`
	// The encryption settings.
	Encryption *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	// The watermark configuration for an image.
	ImageWatermarks []*ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarks `json:"ImageWatermarks,omitempty" xml:"ImageWatermarks,omitempty" type:"Repeated"`
	// Indicates whether the tags of the input stream are inherited in the output stream. This parameter does not take effect when the input is not a media asset. Default value: false.
	IsInheritTags *bool `json:"IsInheritTags,omitempty" xml:"IsInheritTags,omitempty"`
	// The subtitle configuration.
	Subtitles []*ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitles `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// The configurations of the text watermarks.
	TextWatermarks []*ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarks `json:"TextWatermarks,omitempty" xml:"TextWatermarks,omitempty" type:"Repeated"`
	// The transcoding configuration.
	Transcode *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscode `json:"Transcode,omitempty" xml:"Transcode,omitempty" type:"Struct"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) GetCombineConfigs() []*ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs {
	return s.CombineConfigs
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) GetEncryption() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption {
	return s.Encryption
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) GetImageWatermarks() []*ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarks {
	return s.ImageWatermarks
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) GetIsInheritTags() *bool {
	return s.IsInheritTags
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) GetSubtitles() []*ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitles {
	return s.Subtitles
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) GetTextWatermarks() []*ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarks {
	return s.TextWatermarks
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) GetTranscode() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscode {
	return s.Transcode
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) SetCombineConfigs(v []*ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig {
	s.CombineConfigs = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) SetEncryption(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig {
	s.Encryption = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) SetImageWatermarks(v []*ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarks) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig {
	s.ImageWatermarks = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) SetIsInheritTags(v bool) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig {
	s.IsInheritTags = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) SetSubtitles(v []*ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitles) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig {
	s.Subtitles = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) SetTextWatermarks(v []*ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarks) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig {
	s.TextWatermarks = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) SetTranscode(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscode) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig {
	s.Transcode = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfig) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs struct {
	// The audio stream index.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 或 exclude
	AudioIndex *string `json:"AudioIndex,omitempty" xml:"AudioIndex,omitempty"`
	// The duration of the input stream. The default value is the duration of the video.
	//
	// example:
	//
	// 20.0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start time of the input stream. Default value: 0.
	//
	// example:
	//
	// 0.0
	Start *float64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The video stream index.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 或 exclude
	VideoIndex *string `json:"VideoIndex,omitempty" xml:"VideoIndex,omitempty"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs) GetAudioIndex() *string {
	return s.AudioIndex
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs) GetDuration() *float64 {
	return s.Duration
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs) GetStart() *float64 {
	return s.Start
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs) GetVideoIndex() *string {
	return s.VideoIndex
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs) SetAudioIndex(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs {
	s.AudioIndex = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs) SetDuration(v float64) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs {
	s.Duration = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs) SetStart(v float64) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs {
	s.Start = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs) SetVideoIndex(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs {
	s.VideoIndex = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigCombineConfigs) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption struct {
	// The ciphertext of HTTP Live Streaming (HLS) encryption.
	//
	// example:
	//
	// MTYi00NDU0LTg5O****
	CipherText *string `json:"CipherText,omitempty" xml:"CipherText,omitempty"`
	// The endpoint of the decryption service for HLS encryption.
	//
	// example:
	//
	// https://sample.com/path?CipherText=MTYi00NDU0LTg5O****
	DecryptKeyUri *string `json:"DecryptKeyUri,omitempty" xml:"DecryptKeyUri,omitempty"`
	// The encryption type.
	//
	// example:
	//
	// PrivateEncryption
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption) GetCipherText() *string {
	return s.CipherText
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption) GetDecryptKeyUri() *string {
	return s.DecryptKeyUri
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption) GetEncryptType() *string {
	return s.EncryptType
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption) SetCipherText(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption {
	s.CipherText = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption) SetDecryptKeyUri(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption {
	s.DecryptKeyUri = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption) SetEncryptType(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption {
	s.EncryptType = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigEncryption) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarks struct {
	// The parameters that are used to overwrite the corresponding parameters of the template.
	OverwriteParams *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarks) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarks) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarks) GetOverwriteParams() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams {
	return s.OverwriteParams
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarks) SetOverwriteParams(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarks {
	s.OverwriteParams = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarks) SetTemplateId(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarks {
	s.TemplateId = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarks) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams struct {
	// The position of the watermark on the x-axis.
	//
	// example:
	//
	// 10
	Dx *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// The position of the watermark on the y-axis.
	//
	// example:
	//
	// 10
	Dy *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The watermark image file.
	File *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The height of the output video.
	//
	// example:
	//
	// 32
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The reference position of the watermark. Valid values: TopLeft, TopRight, BottomLeft, and BottomRight. Default value: TopLeft.
	//
	// example:
	//
	// TopLeft
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// The timeline settings.
	Timeline *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline `json:"Timeline,omitempty" xml:"Timeline,omitempty" type:"Struct"`
	// The width of the output video.
	//
	// example:
	//
	// 32
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) GetDx() *string {
	return s.Dx
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) GetDy() *string {
	return s.Dy
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) GetFile() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsFile {
	return s.File
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) GetHeight() *string {
	return s.Height
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) GetReferPos() *string {
	return s.ReferPos
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) GetTimeline() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline {
	return s.Timeline
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) GetWidth() *string {
	return s.Width
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) SetDx(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Dx = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) SetDy(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Dy = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) SetFile(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.File = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) SetHeight(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Height = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) SetReferPos(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.ReferPos = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) SetTimeline(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Timeline = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) SetWidth(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Width = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParams) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsFile struct {
	// The media object. If Type is set to OSS, the URL of an OSS object is returned. Both the OSS and HTTP protocols are supported. If Type is set to Media, the ID of a media asset is returned.
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

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) GetMedia() *string {
	return s.Media
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) GetType() *string {
	return s.Type
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) SetMedia(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsFile {
	s.Media = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) SetType(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsFile {
	s.Type = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline struct {
	// The duration of the stream. Valid values: the number of seconds or "ToEND".
	//
	// example:
	//
	// ToEND
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The beginning of the time range for which data was queried.
	//
	// example:
	//
	// 00:00:05
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) GetDuration() *string {
	return s.Duration
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) GetStart() *string {
	return s.Start
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) SetDuration(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline {
	s.Duration = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) SetStart(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline {
	s.Start = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitles struct {
	// The parameters that are used to overwrite the corresponding parameters of the template.
	OverwriteParams *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitles) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitles) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitles) GetOverwriteParams() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams {
	return s.OverwriteParams
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitles) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitles) SetOverwriteParams(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitles {
	s.OverwriteParams = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitles) SetTemplateId(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitles {
	s.TemplateId = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitles) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams struct {
	// The file encoding format.
	//
	// example:
	//
	// UTF-8
	CharEnc *string `json:"CharEnc,omitempty" xml:"CharEnc,omitempty"`
	// The subtitle file.
	File *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The format of the subtitle file.
	//
	// example:
	//
	// vtt
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams) GetCharEnc() *string {
	return s.CharEnc
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams) GetFile() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParamsFile {
	return s.File
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams) GetFormat() *string {
	return s.Format
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams) SetCharEnc(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams {
	s.CharEnc = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams) SetFile(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParamsFile) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams {
	s.File = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams) SetFormat(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams {
	s.Format = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParams) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParamsFile struct {
	// The media object.
	//
	// 	- If Type is set to OSS, the URL of an OSS object is returned. Both the OSS and HTTP protocols are supported.
	//
	// 	- If Type is set to Media, the ID of a media asset is returned.
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

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParamsFile) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParamsFile) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParamsFile) GetMedia() *string {
	return s.Media
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParamsFile) GetType() *string {
	return s.Type
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParamsFile) SetMedia(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParamsFile {
	s.Media = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParamsFile) SetType(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParamsFile {
	s.Type = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigSubtitlesOverwriteParamsFile) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarks struct {
	// The parameters that are used to overwrite the corresponding parameters of the template.
	OverwriteParams *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarks) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarks) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarks) GetOverwriteParams() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams {
	return s.OverwriteParams
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarks) SetOverwriteParams(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarks {
	s.OverwriteParams = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarks) SetTemplateId(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarks {
	s.TemplateId = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarks) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams struct {
	// Indicates whether the text size was adjusted based on the output video dimensions. true / false, default: false
	//
	// example:
	//
	// false
	Adaptive *string `json:"Adaptive,omitempty" xml:"Adaptive,omitempty"`
	// The border color.
	//
	// example:
	//
	// #006400
	BorderColor *string `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// The border width.
	//
	// example:
	//
	// 0
	BorderWidth *int32 `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	// The watermark text. Base64 encoding is not required. The string must be encoded in UTF-8.
	//
	// example:
	//
	// 测试水印
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The transparency of the watermark.
	//
	// example:
	//
	// 1.0
	FontAlpha *string `json:"FontAlpha,omitempty" xml:"FontAlpha,omitempty"`
	// The color of the text.
	//
	// example:
	//
	// #006400
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// The font of the text.
	//
	// example:
	//
	// SimSun
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// The size of the text.
	//
	// example:
	//
	// 16
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// The distance of the watermark from the left edge.
	//
	// example:
	//
	// 10
	Left *string `json:"Left,omitempty" xml:"Left,omitempty"`
	// The distance of the watermark from the top edge.
	//
	// example:
	//
	// 10
	Top *string `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) GetAdaptive() *string {
	return s.Adaptive
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) GetBorderColor() *string {
	return s.BorderColor
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) GetBorderWidth() *int32 {
	return s.BorderWidth
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) GetContent() *string {
	return s.Content
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontAlpha() *string {
	return s.FontAlpha
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontColor() *string {
	return s.FontColor
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontName() *string {
	return s.FontName
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontSize() *int32 {
	return s.FontSize
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) GetLeft() *string {
	return s.Left
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) GetTop() *string {
	return s.Top
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) SetAdaptive(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Adaptive = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) SetBorderColor(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.BorderColor = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) SetBorderWidth(v int32) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.BorderWidth = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) SetContent(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Content = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontAlpha(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontAlpha = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontColor(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontColor = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontName(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontName = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontSize(v int32) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontSize = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) SetLeft(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Left = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) SetTop(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Top = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTextWatermarksOverwriteParams) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscode struct {
	// The parameters that are used to overwrite the corresponding parameters of the template.
	OverwriteParams *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscode) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscode) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscode) GetOverwriteParams() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams {
	return s.OverwriteParams
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscode) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscode) SetOverwriteParams(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscode {
	s.OverwriteParams = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscode) SetTemplateId(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscode {
	s.TemplateId = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscode) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams struct {
	// The audio settings.
	Audio *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// The encapsulation format settings.
	Container *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The encapsulation settings.
	MuxConfig *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	Tags      map[string]*string                                                                          `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The video settings.
	Video *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams) GetAudio() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	return s.Audio
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams) GetContainer() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsContainer {
	return s.Container
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams) GetMuxConfig() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig {
	return s.MuxConfig
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams) GetTags() map[string]*string {
	return s.Tags
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams) GetVideo() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	return s.Video
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams) SetAudio(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams {
	s.Audio = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams) SetContainer(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsContainer) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams {
	s.Container = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams) SetMuxConfig(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams {
	s.MuxConfig = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams) SetTags(v map[string]*string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams {
	s.Tags = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams) SetVideo(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams {
	s.Video = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParams) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio struct {
	// The audio bitrate of the output file.
	//
	// 	- Valid values: [8,1000].
	//
	// 	- Unit: Kbit/s.
	//
	// 	- Default value: 128.
	//
	// example:
	//
	// 128
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The number of sound channels. Default value: 2.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The audio codec. Valid values: AAC, MP3, VORBIS, and FLAC. Default value: AAC.
	//
	// example:
	//
	// AAC
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The audio codec profile. If the Codec parameter is set to AAC, the valid values are aac_low, aac_he, aac_he_v2, aac_ld, and aac_eld.
	//
	// example:
	//
	// aac_low
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Indicates whether the audio stream is deleted.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The sampling rate.
	//
	// 	- Default value: 44100.
	//
	// 	- Valid values: 22050, 32000, 44100, 48000, and 96000.
	//
	// 	- Unit: Hz.
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	// The volume configurations.
	Volume *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Struct"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetChannels() *string {
	return s.Channels
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetCodec() *string {
	return s.Codec
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetProfile() *string {
	return s.Profile
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetRemove() *string {
	return s.Remove
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetVolume() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	return s.Volume
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetBitrate(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Bitrate = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetChannels(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Channels = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetCodec(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Codec = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetProfile(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Profile = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetRemove(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Remove = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetSamplerate(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Samplerate = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetVolume(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Volume = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudio) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume struct {
	// The output volume.
	//
	// example:
	//
	// -6
	IntegratedLoudnessTarget *string `json:"IntegratedLoudnessTarget,omitempty" xml:"IntegratedLoudnessTarget,omitempty"`
	// The volume range.
	//
	// example:
	//
	// 8
	LoudnessRangeTarget *string `json:"LoudnessRangeTarget,omitempty" xml:"LoudnessRangeTarget,omitempty"`
	// The volume adjustment method. Valid values:
	//
	// example:
	//
	// auto
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The peak volume.
	//
	// example:
	//
	// -1
	TruePeak *string `json:"TruePeak,omitempty" xml:"TruePeak,omitempty"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetIntegratedLoudnessTarget() *string {
	return s.IntegratedLoudnessTarget
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetLoudnessRangeTarget() *string {
	return s.LoudnessRangeTarget
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetMethod() *string {
	return s.Method
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetTruePeak() *string {
	return s.TruePeak
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetIntegratedLoudnessTarget(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.IntegratedLoudnessTarget = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetLoudnessRangeTarget(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.LoudnessRangeTarget = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetMethod(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.Method = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetTruePeak(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.TruePeak = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsContainer struct {
	// The container format.
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsContainer) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsContainer) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsContainer) GetFormat() *string {
	return s.Format
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsContainer) SetFormat(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsContainer {
	s.Format = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsContainer) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig struct {
	// The segment settings.
	Segment *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) GetSegment() *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	return s.Segment
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) SetSegment(v *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig {
	s.Segment = v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment struct {
	// The segment length.
	//
	// example:
	//
	// 10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The forced segmentation point in time.
	//
	// example:
	//
	// 2,3
	ForceSegTime *string `json:"ForceSegTime,omitempty" xml:"ForceSegTime,omitempty"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GetForceSegTime() *string {
	return s.ForceSegTime
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) SetDuration(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) SetForceSegTime(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	s.ForceSegTime = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo struct {
	// The maximum adaptive bitrate (ABR). This parameter takes effect only for Narrowband HD 1.0. Valid values: [10,50000]. Unit: Kbit/s.
	//
	// example:
	//
	// 6000
	AbrMax *string `json:"AbrMax,omitempty" xml:"AbrMax,omitempty"`
	// The average bitrate of the video.
	//
	// 	- Valid values: [10,50000].
	//
	// 	- Unit: Kbit/s.
	//
	// example:
	//
	// 3000
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The buffer size.
	//
	// 	- Valid values: [1000,128000].
	//
	// 	- Default value: 6000.
	//
	// 	- Unit: KB.
	//
	// example:
	//
	// 6000
	Bufsize *string `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	// The encoding format.
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The constant rate factor.
	//
	// 	- Valid values: [0,51].
	//
	// 	- Default value: 23 if the encoding format is H.264, or 26 if the encoding format is H.265.
	//
	// If this parameter is set, the value of Bitrate becomes invalid.
	//
	// example:
	//
	// 23
	Crf *string `json:"Crf,omitempty" xml:"Crf,omitempty"`
	// The method of video cropping. Valid values:
	//
	// 	- border: automatically detects and removes black bars.
	//
	// 	- A value in the width:height:left:top format: crops the videos based on the custom settings. Example: 1280:800:0:140.
	//
	// example:
	//
	// 1280:800:0:140
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// The frame rate.
	//
	// 	- Valid values: (0,60].
	//
	// 	- The value is 60 if the frame rate of the input video exceeds 60.
	//
	// 	- Default value: the frame rate of the input video.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The maximum number of frames between two keyframes.
	//
	// 	- Valid values: [1,1080000].
	//
	// 	- Default value: 250.
	//
	// example:
	//
	// 250
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height of the output video.
	//
	// 	- Valid values: [128,4096].
	//
	// 	- Unit: pixels.
	//
	// 	- Default value: the height of the input video.
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// Indicates whether the auto-rotate screen feature is enabled.
	//
	// example:
	//
	// false
	LongShortMode *string `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
	// The maximum bitrate of the output video. Valid values: [10,50000]. Unit: Kbit/s.
	//
	// example:
	//
	// 9000
	Maxrate *string `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	// The black bars added to the video.
	//
	// 	- Format: width:height:left:top.
	//
	// 	- Example: 1280:800:0:140.
	//
	// example:
	//
	// 1280:800:0:140
	Pad *string `json:"Pad,omitempty" xml:"Pad,omitempty"`
	// The pixel format of the video. Valid values: standard pixel formats such as yuv420p and yuvj420p.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The preset video algorithm. This parameter takes effect only if the encoding format is H.264. Valid values: veryfast, fast, medium, slow, and slower. Default value: medium.
	//
	// example:
	//
	// medium
	Preset *string `json:"Preset,omitempty" xml:"Preset,omitempty"`
	// The encoding profile. Valid values: baseline, main, and high.
	//
	// 	- baseline: applicable to mobile devices.
	//
	// 	- main: applicable to standard-definition devices.
	//
	// 	- high: applicable to high-definition devices.
	//
	// Default value: high.
	//
	// example:
	//
	// Main
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Indicates whether the video was removed.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The scan mode. Valid values: interlaced and progressive.
	//
	// example:
	//
	// progressive
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// The width of the output video.
	//
	// 	- Valid values: [128,4096].
	//
	// 	- Unit: pixels.
	//
	// 	- Default value: the width of the input video.
	//
	// example:
	//
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetAbrMax() *string {
	return s.AbrMax
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetCodec() *string {
	return s.Codec
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetCrf() *string {
	return s.Crf
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetCrop() *string {
	return s.Crop
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetFps() *string {
	return s.Fps
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetGop() *string {
	return s.Gop
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetHeight() *string {
	return s.Height
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetLongShortMode() *string {
	return s.LongShortMode
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetPad() *string {
	return s.Pad
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetPreset() *string {
	return s.Preset
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetProfile() *string {
	return s.Profile
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetRemove() *string {
	return s.Remove
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetWidth() *string {
	return s.Width
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetAbrMax(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.AbrMax = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetBitrate(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Bitrate = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetBufsize(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Bufsize = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetCodec(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Codec = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetCrf(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Crf = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetCrop(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Crop = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetFps(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Fps = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetGop(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Gop = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetHeight(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Height = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetLongShortMode(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.LongShortMode = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetMaxrate(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Maxrate = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetPad(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Pad = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetPixFmt(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.PixFmt = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetPreset(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Preset = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetProfile(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Profile = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetRemove(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Remove = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetScanMode(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.ScanMode = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetWidth(v string) *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Width = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsOutputGroupProcessConfigTranscodeOverwriteParamsVideo) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeJobsResponseBodyJobsScheduleConfig struct {
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

func (s ListTranscodeJobsResponseBodyJobsScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponseBodyJobsScheduleConfig) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponseBodyJobsScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListTranscodeJobsResponseBodyJobsScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *ListTranscodeJobsResponseBodyJobsScheduleConfig) SetPipelineId(v string) *ListTranscodeJobsResponseBodyJobsScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsScheduleConfig) SetPriority(v int32) *ListTranscodeJobsResponseBodyJobsScheduleConfig {
	s.Priority = &v
	return s
}

func (s *ListTranscodeJobsResponseBodyJobsScheduleConfig) Validate() error {
	return dara.Validate(s)
}
