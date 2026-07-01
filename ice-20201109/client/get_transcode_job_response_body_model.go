// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTranscodeJobResponseBody
	GetRequestId() *string
	SetTranscodeParentJob(v *GetTranscodeJobResponseBodyTranscodeParentJob) *GetTranscodeJobResponseBody
	GetTranscodeParentJob() *GetTranscodeJobResponseBodyTranscodeParentJob
}

type GetTranscodeJobResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 9EDC30DC-0050-5459-B788-F761B2BE359B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// TranscodeParentJobWithSubJobDTO
	TranscodeParentJob *GetTranscodeJobResponseBodyTranscodeParentJob `json:"TranscodeParentJob,omitempty" xml:"TranscodeParentJob,omitempty" type:"Struct"`
}

func (s GetTranscodeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTranscodeJobResponseBody) GetTranscodeParentJob() *GetTranscodeJobResponseBodyTranscodeParentJob {
	return s.TranscodeParentJob
}

func (s *GetTranscodeJobResponseBody) SetRequestId(v string) *GetTranscodeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTranscodeJobResponseBody) SetTranscodeParentJob(v *GetTranscodeJobResponseBodyTranscodeParentJob) *GetTranscodeJobResponseBody {
	s.TranscodeParentJob = v
	return s
}

func (s *GetTranscodeJobResponseBody) Validate() error {
	if s.TranscodeParentJob != nil {
		if err := s.TranscodeParentJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJob struct {
	// Job creation time. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Job completion time. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// Job input group. A single input represents a transcoding job. Multiple inputs represent an audio-video merging job.
	InputGroup []*GetTranscodeJobResponseBodyTranscodeParentJobInputGroup `json:"InputGroup,omitempty" xml:"InputGroup,omitempty" type:"Repeated"`
	// Number of sub-jobs.
	//
	// example:
	//
	// 1
	JobCount *int32 `json:"JobCount,omitempty" xml:"JobCount,omitempty"`
	// Job name.
	//
	// example:
	//
	// transcode-job
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Job output group.
	OutputGroup []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroup `json:"OutputGroup,omitempty" xml:"OutputGroup,omitempty" type:"Repeated"`
	// Parent job ID.
	//
	// example:
	//
	// 8b2198504dd340b7b3c9842a74fc9baa
	ParentJobId *string `json:"ParentJobId,omitempty" xml:"ParentJobId,omitempty"`
	// Job completion percentage.
	//
	// example:
	//
	// 0
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// Request ID at job submission.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Job scheduling configuration.
	ScheduleConfig *GetTranscodeJobResponseBodyTranscodeParentJobScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// Job status.
	//
	// - Success: All sub-jobs completed successfully.
	//
	// - Fail: All sub-jobs failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Job submission time. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// List of sub-jobs.
	TranscodeJobList []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList `json:"TranscodeJobList,omitempty" xml:"TranscodeJobList,omitempty" type:"Repeated"`
	// Job source.
	//
	// - API: API
	//
	// - WorkFlow: Workflow
	//
	// - Console: Console
	//
	// example:
	//
	// API
	TriggerSource *string `json:"TriggerSource,omitempty" xml:"TriggerSource,omitempty"`
	// User data.
	//
	// example:
	//
	// user-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJob) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJob) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetInputGroup() []*GetTranscodeJobResponseBodyTranscodeParentJobInputGroup {
	return s.InputGroup
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetJobCount() *int32 {
	return s.JobCount
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetName() *string {
	return s.Name
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetOutputGroup() []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroup {
	return s.OutputGroup
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetParentJobId() *string {
	return s.ParentJobId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetPercent() *int32 {
	return s.Percent
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetScheduleConfig() *GetTranscodeJobResponseBodyTranscodeParentJobScheduleConfig {
	return s.ScheduleConfig
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetStatus() *string {
	return s.Status
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetTranscodeJobList() []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	return s.TranscodeJobList
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) GetUserData() *string {
	return s.UserData
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetCreateTime(v string) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.CreateTime = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetFinishTime(v string) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.FinishTime = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetInputGroup(v []*GetTranscodeJobResponseBodyTranscodeParentJobInputGroup) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.InputGroup = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetJobCount(v int32) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.JobCount = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetName(v string) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.Name = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetOutputGroup(v []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroup) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.OutputGroup = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetParentJobId(v string) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.ParentJobId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetPercent(v int32) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.Percent = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetRequestId(v string) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.RequestId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetScheduleConfig(v *GetTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.ScheduleConfig = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetStatus(v string) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.Status = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetSubmitTime(v string) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.SubmitTime = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetTranscodeJobList(v []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.TranscodeJobList = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetTriggerSource(v string) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.TriggerSource = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) SetUserData(v string) *GetTranscodeJobResponseBodyTranscodeParentJob {
	s.UserData = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJob) Validate() error {
	if s.InputGroup != nil {
		for _, item := range s.InputGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutputGroup != nil {
		for _, item := range s.OutputGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScheduleConfig != nil {
		if err := s.ScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TranscodeJobList != nil {
		for _, item := range s.TranscodeJobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobInputGroup struct {
	// Media value:
	//
	// - If Type is OSS, this field is a URL that supports OSS and HTTP protocols.
	//
	// - If Type is Media, this field is a media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// Media object type.
	//
	// Valid values:
	//
	// - OSS: OSS file.
	//
	// - Media: Media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobInputGroup) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobInputGroup) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobInputGroup) GetMedia() *string {
	return s.Media
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobInputGroup) GetType() *string {
	return s.Type
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobInputGroup) SetMedia(v string) *GetTranscodeJobResponseBodyTranscodeParentJobInputGroup {
	s.Media = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobInputGroup) SetType(v string) *GetTranscodeJobResponseBodyTranscodeParentJobInputGroup {
	s.Type = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobInputGroup) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroup struct {
	// Output media configuration.
	Output *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// Job processing configuration.
	ProcessConfig *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig `json:"ProcessConfig,omitempty" xml:"ProcessConfig,omitempty" type:"Struct"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroup) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroup) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroup) GetOutput() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput {
	return s.Output
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroup) GetProcessConfig() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig {
	return s.ProcessConfig
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroup) SetOutput(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroup {
	s.Output = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroup) SetProcessConfig(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroup {
	s.ProcessConfig = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroup) Validate() error {
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.ProcessConfig != nil {
		if err := s.ProcessConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput struct {
	// Media value:
	//
	// - If Type is OSS, this field is a URL that supports OSS and HTTP protocols.
	//
	// - If Type is Media, this field is a media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// Output stream URL.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
	// Media object type.
	//
	// Valid values:
	//
	// - OSS: OSS file.
	//
	// - Media: Media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) GetMedia() *string {
	return s.Media
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) GetType() *string {
	return s.Type
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) SetMedia(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput {
	s.Media = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) SetOutputUrl(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput {
	s.OutputUrl = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) SetType(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput {
	s.Type = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig struct {
	// Multi-input merging configuration
	CombineConfigs []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs `json:"CombineConfigs,omitempty" xml:"CombineConfigs,omitempty" type:"Repeated"`
	// Encryption configuration.
	Encryption *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	// Image watermark configuration.
	ImageWatermarks []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks `json:"ImageWatermarks,omitempty" xml:"ImageWatermarks,omitempty" type:"Repeated"`
	// Subtitle embedding configuration.
	Subtitles []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// Text watermark configuration.
	TextWatermarks []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks `json:"TextWatermarks,omitempty" xml:"TextWatermarks,omitempty" type:"Repeated"`
	// Transcoding configuration.
	Transcode *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode `json:"Transcode,omitempty" xml:"Transcode,omitempty" type:"Struct"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) GetCombineConfigs() []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs {
	return s.CombineConfigs
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) GetEncryption() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption {
	return s.Encryption
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) GetImageWatermarks() []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks {
	return s.ImageWatermarks
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) GetSubtitles() []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles {
	return s.Subtitles
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) GetTextWatermarks() []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks {
	return s.TextWatermarks
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) GetTranscode() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode {
	return s.Transcode
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) SetCombineConfigs(v []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig {
	s.CombineConfigs = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) SetEncryption(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig {
	s.Encryption = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) SetImageWatermarks(v []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig {
	s.ImageWatermarks = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) SetSubtitles(v []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig {
	s.Subtitles = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) SetTextWatermarks(v []*GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig {
	s.TextWatermarks = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) SetTranscode(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig {
	s.Transcode = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) Validate() error {
	if s.CombineConfigs != nil {
		for _, item := range s.CombineConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			return err
		}
	}
	if s.ImageWatermarks != nil {
		for _, item := range s.ImageWatermarks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Subtitles != nil {
		for _, item := range s.Subtitles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TextWatermarks != nil {
		for _, item := range s.TextWatermarks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Transcode != nil {
		if err := s.Transcode.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs struct {
	// Audio stream index
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 or exclude
	AudioIndex *string `json:"AudioIndex,omitempty" xml:"AudioIndex,omitempty"`
	// Duration of the input stream. Default: video duration.
	//
	// example:
	//
	// 20.0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Start time of the input stream. Default: 0
	//
	// example:
	//
	// 0.0
	Start *float64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// Video stream index
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 or exclude
	VideoIndex *string `json:"VideoIndex,omitempty" xml:"VideoIndex,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) GetAudioIndex() *string {
	return s.AudioIndex
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) GetDuration() *float64 {
	return s.Duration
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) GetStart() *float64 {
	return s.Start
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) GetVideoIndex() *string {
	return s.VideoIndex
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) SetAudioIndex(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs {
	s.AudioIndex = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) SetDuration(v float64) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs {
	s.Duration = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) SetStart(v float64) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs {
	s.Start = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) SetVideoIndex(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs {
	s.VideoIndex = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption struct {
	// Ciphertext of the encryption key for standard encryption.
	//
	// example:
	//
	// MTYi00NDU0LTg5O****
	CipherText *string `json:"CipherText,omitempty" xml:"CipherText,omitempty"`
	// Decryption service endpoint for standard encryption.
	//
	// example:
	//
	// https://sample.com/path?CipherText=MTYi00NDU0LTg5O****
	DecryptKeyUri *string `json:"DecryptKeyUri,omitempty" xml:"DecryptKeyUri,omitempty"`
	// Encryption type.
	//
	// example:
	//
	// PrivateEncryption
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// Key service type. Only KMS and Base64 are supported.
	//
	// example:
	//
	// KMS
	KeyServiceType *string `json:"KeyServiceType,omitempty" xml:"KeyServiceType,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) GetCipherText() *string {
	return s.CipherText
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) GetDecryptKeyUri() *string {
	return s.DecryptKeyUri
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) GetEncryptType() *string {
	return s.EncryptType
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) GetKeyServiceType() *string {
	return s.KeyServiceType
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) SetCipherText(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption {
	s.CipherText = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) SetDecryptKeyUri(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption {
	s.DecryptKeyUri = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) SetEncryptType(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption {
	s.EncryptType = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) SetKeyServiceType(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption {
	s.KeyServiceType = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks struct {
	// Override parameters. If specified, these parameters override the corresponding template parameters.
	OverwriteParams *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// Template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) GetOverwriteParams() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	return s.OverwriteParams
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) SetOverwriteParams(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks {
	s.OverwriteParams = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) SetTemplateId(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks {
	s.TemplateId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams struct {
	// Watermark position, x.
	//
	// example:
	//
	// 10
	Dx *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// Watermark position, y.
	//
	// example:
	//
	// 10
	Dy *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// Watermark image file.
	File *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// Height.
	//
	// example:
	//
	// 32
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// Reference position: TopLeft, TopRight, BottomLeft, BottomRight. Default: TopLeft.
	//
	// example:
	//
	// TopLeft
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// Display time settings.
	Timeline *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline `json:"Timeline,omitempty" xml:"Timeline,omitempty" type:"Struct"`
	// Width.
	//
	// example:
	//
	// 32
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GetDx() *string {
	return s.Dx
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GetDy() *string {
	return s.Dy
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GetFile() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile {
	return s.File
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GetHeight() *string {
	return s.Height
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GetReferPos() *string {
	return s.ReferPos
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GetTimeline() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline {
	return s.Timeline
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GetWidth() *string {
	return s.Width
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) SetDx(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Dx = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) SetDy(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Dy = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) SetFile(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.File = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) SetHeight(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Height = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) SetReferPos(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.ReferPos = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) SetTimeline(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Timeline = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) SetWidth(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Width = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	if s.Timeline != nil {
		if err := s.Timeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile struct {
	// Media value:
	//
	// - If Type is OSS, this field is a URL that supports OSS and HTTP protocols.
	//
	// - If Type is Media, this field is a media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// Media object type.
	//
	// Valid values:
	//
	// - OSS: OSS file.
	//
	// - Media: Media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) GetMedia() *string {
	return s.Media
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) GetType() *string {
	return s.Type
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) SetMedia(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile {
	s.Media = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) SetType(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile {
	s.Type = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline struct {
	// Display duration in seconds or "ToEND".
	//
	// example:
	//
	// ToEND
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Start time.
	//
	// example:
	//
	// 00:00:05
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) GetDuration() *string {
	return s.Duration
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) GetStart() *string {
	return s.Start
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) SetDuration(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline {
	s.Duration = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) SetStart(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline {
	s.Start = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles struct {
	// Override parameters. If specified, these parameters override the corresponding template parameters.
	OverwriteParams *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// Template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) GetOverwriteParams() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams {
	return s.OverwriteParams
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) SetOverwriteParams(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles {
	s.OverwriteParams = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) SetTemplateId(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles {
	s.TemplateId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams struct {
	// File encoding format.
	//
	// example:
	//
	// UTF-8
	CharEnc *string `json:"CharEnc,omitempty" xml:"CharEnc,omitempty"`
	// Subtitle file.
	File *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// Subtitle file format.
	//
	// example:
	//
	// vtt
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) GetCharEnc() *string {
	return s.CharEnc
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) GetFile() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile {
	return s.File
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) GetFormat() *string {
	return s.Format
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) SetCharEnc(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams {
	s.CharEnc = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) SetFile(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams {
	s.File = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) SetFormat(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams {
	s.Format = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile struct {
	// Media value:
	//
	// - If Type is OSS, this field is a URL that supports OSS and HTTP protocols.
	//
	// - If Type is Media, this field is a media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// Media object type.
	//
	// Valid values:
	//
	// - OSS: OSS file.
	//
	// - Media: Media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) GetMedia() *string {
	return s.Media
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) GetType() *string {
	return s.Type
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) SetMedia(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile {
	s.Media = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) SetType(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile {
	s.Type = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks struct {
	// Override parameters. If specified, these parameters override the corresponding template parameters.
	OverwriteParams *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// Template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) GetOverwriteParams() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	return s.OverwriteParams
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) SetOverwriteParams(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks {
	s.OverwriteParams = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) SetTemplateId(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks {
	s.TemplateId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams struct {
	// Adjust font size based on output video size. Valid values: true / false. Default: false.
	//
	// example:
	//
	// false
	Adaptive *string `json:"Adaptive,omitempty" xml:"Adaptive,omitempty"`
	// Border color.
	//
	// example:
	//
	// #006400
	BorderColor *string `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// Border width.
	//
	// example:
	//
	// 0
	BorderWidth *int32 `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	// Watermark text. Do not Base64 encode. The string must be UTF-8 encoded.
	//
	// example:
	//
	// test watermark
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Transparency.
	//
	// example:
	//
	// 1.0
	FontAlpha *string `json:"FontAlpha,omitempty" xml:"FontAlpha,omitempty"`
	// Color.
	//
	// example:
	//
	// #006400
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// Font.
	//
	// example:
	//
	// SimSun
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// Font size.
	//
	// example:
	//
	// 16
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// Watermark position, distance from the left.
	//
	// example:
	//
	// 10
	Left *string `json:"Left,omitempty" xml:"Left,omitempty"`
	// Watermark position, distance from the top.
	//
	// example:
	//
	// 10
	Top *string `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetAdaptive() *string {
	return s.Adaptive
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetBorderColor() *string {
	return s.BorderColor
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetBorderWidth() *int32 {
	return s.BorderWidth
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetContent() *string {
	return s.Content
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontAlpha() *string {
	return s.FontAlpha
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontColor() *string {
	return s.FontColor
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontName() *string {
	return s.FontName
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontSize() *int32 {
	return s.FontSize
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetLeft() *string {
	return s.Left
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetTop() *string {
	return s.Top
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetAdaptive(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Adaptive = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetBorderColor(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.BorderColor = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetBorderWidth(v int32) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.BorderWidth = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetContent(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Content = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontAlpha(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontAlpha = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontColor(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontColor = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontName(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontName = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontSize(v int32) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontSize = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetLeft(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Left = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetTop(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Top = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode struct {
	// Override parameters. If specified, these parameters override the corresponding template parameters.
	OverwriteParams *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// Template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) GetOverwriteParams() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams {
	return s.OverwriteParams
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) SetOverwriteParams(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode {
	s.OverwriteParams = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) SetTemplateId(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode {
	s.TemplateId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams struct {
	// Audio settings.
	Audio *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// Container format settings.
	Container *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// Muxing settings.
	MuxConfig *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	// Conditional transcoding parameters.
	TransConfig *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	// Video settings.
	Video *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) GetAudio() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	return s.Audio
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) GetContainer() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer {
	return s.Container
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) GetMuxConfig() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig {
	return s.MuxConfig
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) GetTransConfig() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	return s.TransConfig
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) GetVideo() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	return s.Video
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) SetAudio(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams {
	s.Audio = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) SetContainer(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams {
	s.Container = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) SetMuxConfig(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams {
	s.MuxConfig = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) SetTransConfig(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams {
	s.TransConfig = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) SetVideo(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams {
	s.Video = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) Validate() error {
	if s.Audio != nil {
		if err := s.Audio.Validate(); err != nil {
			return err
		}
	}
	if s.Container != nil {
		if err := s.Container.Validate(); err != nil {
			return err
		}
	}
	if s.MuxConfig != nil {
		if err := s.MuxConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TransConfig != nil {
		if err := s.TransConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio struct {
	// Audio bitrate of the output file.
	//
	// - Valid range: [8, 1000]
	//
	// - Unit: Kbps
	//
	// - Default value: 128
	//
	// example:
	//
	// 128
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// Number of sound channels.
	//
	// Default value: 2
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// Audio codec format: AAC, MP3, VORBIS, FLAC.
	//
	// Default value: AAC
	//
	// example:
	//
	// AAC
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// Audio encoding preset.
	//
	// When Codec is AAC, valid values are aac_low, aac_he, aac_he_v2, aac_ld, aac_eld.
	//
	// example:
	//
	// aac_low
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Whether to remove the audio stream.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// Sample rate.
	//
	// - Default value: 44100
	//
	// - Supported values: 22050, 32000, 44100, 48000, 96000,
	//
	// - Unit: Hz
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	// Volume control.
	Volume *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Struct"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetChannels() *string {
	return s.Channels
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetCodec() *string {
	return s.Codec
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetProfile() *string {
	return s.Profile
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetRemove() *string {
	return s.Remove
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetVolume() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	return s.Volume
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetBitrate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Bitrate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetChannels(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Channels = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetCodec(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Codec = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetProfile(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Profile = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetRemove(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Remove = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetSamplerate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Samplerate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetVolume(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Volume = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) Validate() error {
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume struct {
	// Target volume.
	//
	// example:
	//
	// -6
	IntegratedLoudnessTarget *string `json:"IntegratedLoudnessTarget,omitempty" xml:"IntegratedLoudnessTarget,omitempty"`
	// Volume range.
	//
	// example:
	//
	// 8
	LoudnessRangeTarget *string `json:"LoudnessRangeTarget,omitempty" xml:"LoudnessRangeTarget,omitempty"`
	// Volume adjustment method.
	//
	// example:
	//
	// auto
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// Peak volume.
	//
	// example:
	//
	// -1
	TruePeak *string `json:"TruePeak,omitempty" xml:"TruePeak,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetIntegratedLoudnessTarget() *string {
	return s.IntegratedLoudnessTarget
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetLoudnessRangeTarget() *string {
	return s.LoudnessRangeTarget
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetMethod() *string {
	return s.Method
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetTruePeak() *string {
	return s.TruePeak
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetIntegratedLoudnessTarget(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.IntegratedLoudnessTarget = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetLoudnessRangeTarget(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.LoudnessRangeTarget = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetMethod(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.Method = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetTruePeak(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.TruePeak = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer struct {
	// Container format.
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer) GetFormat() *string {
	return s.Format
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer) SetFormat(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer {
	s.Format = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig struct {
	// Segmentation settings.
	Segment *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) GetSegment() *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	return s.Segment
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) SetSegment(v *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig {
	s.Segment = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) Validate() error {
	if s.Segment != nil {
		if err := s.Segment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment struct {
	// Segment length.
	//
	// example:
	//
	// 10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Forced segmentation time points.
	//
	// example:
	//
	// 2,3
	ForceSegTime *string `json:"ForceSegTime,omitempty" xml:"ForceSegTime,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GetForceSegTime() *string {
	return s.ForceSegTime
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) SetDuration(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) SetForceSegTime(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	s.ForceSegTime = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig struct {
	// Resolution adjustment method. This parameter takes effect only when both Width and Height are specified. It can be used together with LongShortMode.
	//
	// Valid values: rescale, crop, pad, none.
	//
	// Default value: none.
	//
	// example:
	//
	// none
	AdjDarMethod *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	// Whether to check audio bitrate. IsCheckAudioBitrate and IsCheckAudioBitrateFail are mutually exclusive. IsCheckAudioBitrateFail takes precedence.
	//
	// - true: Check. If the input audio bitrate is lower than the output setting, transcode using the input audio bitrate.
	//
	// - false: Do not check.
	//
	// Default value:
	//
	// - If this parameter is empty and the codec differs from the input source: false.
	//
	// - If this parameter is empty and the codec matches the input source: true.
	//
	// example:
	//
	// true
	IsCheckAudioBitrate *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	// Whether to check audio bitrate. IsCheckAudioBitrate and IsCheckAudioBitrateFail are mutually exclusive. This parameter takes precedence.
	//
	// - true: Check. If the input audio bitrate is lower than the output setting, return a transcoding failure.
	//
	// - false: Do not check.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	// Whether to check video resolution. IsCheckReso and IsCheckResoFail are mutually exclusive. IsCheckResoFail takes precedence.
	//
	// - true: Check. If the input video resolution (width or height) is smaller than the output setting, transcode using the input video resolution.
	//
	// - false: Do not check.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckReso *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	// Whether to check video resolution. IsCheckReso and IsCheckResoFail are mutually exclusive. This parameter takes precedence.
	//
	// - true: Check. If the input video resolution (width or height) is smaller than the output setting, return a transcoding failure.
	//
	// - false: Do not check.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckResoFail *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	// Whether to check video bitrate. IsCheckVideoBitrate and IsCheckVideoBitrateFail are mutually exclusive. IsCheckVideoBitrateFail takes precedence.
	//
	// - true: Check. If the input video bitrate is lower than the output setting, transcode using the input video bitrate.
	//
	// - false: Do not check.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckVideoBitrate *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	// Whether to check video bitrate. IsCheckVideoBitrate and IsCheckVideoBitrateFail are mutually exclusive. This parameter takes precedence.
	//
	// - true: Check. If the input video bitrate is lower than the output setting, return a transcoding failure.
	//
	// - false: Do not check.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	// Video transcoding mode. Valid values:
	//
	// - onepass: generally used for ABR. Faster encoding than twopass.
	//
	// - twopass: generally used for VBR. Slower encoding than onepass.
	//
	// - CBR: constant bitrate mode.
	//
	// Default value: onepass.
	//
	// example:
	//
	// onepass
	TransMode *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetAdjDarMethod() *string {
	return s.AdjDarMethod
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckAudioBitrate() *string {
	return s.IsCheckAudioBitrate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckAudioBitrateFail() *string {
	return s.IsCheckAudioBitrateFail
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckReso() *string {
	return s.IsCheckReso
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckResoFail() *string {
	return s.IsCheckResoFail
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckVideoBitrate() *string {
	return s.IsCheckVideoBitrate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckVideoBitrateFail() *string {
	return s.IsCheckVideoBitrateFail
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetAdjDarMethod(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.AdjDarMethod = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckAudioBitrate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckAudioBitrateFail(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckReso(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckResoFail(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckVideoBitrate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckVideoBitrateFail(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetTransMode(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.TransMode = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo struct {
	// Maximum ABR bitrate (valid only for narrow-high 1)
	//
	// - Valid range: [10, 50000],
	//
	// - Unit: Kbps.
	//
	// example:
	//
	// 6000
	AbrMax *string `json:"AbrMax,omitempty" xml:"AbrMax,omitempty"`
	// Video average bitrate.
	//
	// - Valid range: [10, 50000].
	//
	// - Unit: Kbps.
	//
	// example:
	//
	// 3000
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// Buffer size
	//
	// - Valid range: [1000, 128000]
	//
	// - Default value: 6000
	//
	// - Unit: Kb
	//
	// example:
	//
	// 6000
	Bufsize *string `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	// Encoding format.
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// Bitrate-quality control factor.
	//
	// - Valid range: [0, 51].
	//
	// - Default value: 23 for H.264, 26 for H.265.
	//
	// - If Crf is set, the Bitrate setting is ignored.
	//
	// example:
	//
	// 23
	Crf *string `json:"Crf,omitempty" xml:"Crf,omitempty"`
	// Video cropping.
	//
	// Two methods are supported.
	//
	// Automatically detect and crop black bars by setting this parameter to "border".
	//
	// Custom cropping in the format: width:height:left:top.
	//
	// Example: 1280:800:0:140
	//
	// example:
	//
	// 1280:800:0:140
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// Frame rate.
	//
	// - Valid range: (0, 60].
	//
	// - If the input file frame rate exceeds 60, use 60.
	//
	// - Default value: input file frame rate.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// Maximum number of frames between keyframes.
	//
	// - Valid range: [1, 1080000].
	//
	// - Default value: 250.
	//
	// example:
	//
	// 250
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// Height.
	//
	// - Valid range: [128, 4096].
	//
	// - Unit: px.
	//
	// - Default value: original video height.
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// Whether to enable automatic rotation (long-short edge mode).
	//
	// example:
	//
	// false
	LongShortMode *string `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
	// Peak video bitrate
	//
	// - Valid range: [10, 50000]
	//
	// - Unit: Kbps.
	//
	// example:
	//
	// 9000
	Maxrate *string `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	// Black bar padding parameters
	//
	// - Format: width:height:left:top.
	//
	// - Example: 1280:800:0:140
	//
	// example:
	//
	// 1280:800:0:140
	Pad *string `json:"Pad,omitempty" xml:"Pad,omitempty"`
	// Video color format.
	//
	// Valid values: yuv420p, yuvj420p, and other standard color formats.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// Video encoder preset. This parameter is supported only for H.264.
	//
	// Supported values: veryfast, fast, medium, slow, slower.
	//
	// Default value: medium.
	//
	// example:
	//
	// medium
	Preset *string `json:"Preset,omitempty" xml:"Preset,omitempty"`
	// Encoding profile.
	//
	// Supported values: baseline, main, high.
	//
	// - baseline: for mobile devices.
	//
	// - main: for standard-resolution devices.
	//
	// - high: for high-resolution devices.
	//
	// Default value: high.
	//
	// example:
	//
	// Main
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Whether to remove the video.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// Scan mode.
	//
	// Supported values: interlaced, progressive.
	//
	// example:
	//
	// progressive
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// Width.
	//
	// - Valid range: [128, 4096].
	//
	// - Unit: px.
	//
	// - Default value: original video width.
	//
	// example:
	//
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetAbrMax() *string {
	return s.AbrMax
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetCodec() *string {
	return s.Codec
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetCrf() *string {
	return s.Crf
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetCrop() *string {
	return s.Crop
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetFps() *string {
	return s.Fps
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetGop() *string {
	return s.Gop
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetHeight() *string {
	return s.Height
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetLongShortMode() *string {
	return s.LongShortMode
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetPad() *string {
	return s.Pad
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetPreset() *string {
	return s.Preset
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetProfile() *string {
	return s.Profile
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetRemove() *string {
	return s.Remove
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetWidth() *string {
	return s.Width
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetAbrMax(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.AbrMax = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetBitrate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Bitrate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetBufsize(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Bufsize = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetCodec(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Codec = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetCrf(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Crf = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetCrop(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Crop = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetFps(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Fps = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetGop(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Gop = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetHeight(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Height = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetLongShortMode(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.LongShortMode = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetMaxrate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Maxrate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetPad(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Pad = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetPixFmt(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.PixFmt = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetPreset(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Preset = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetProfile(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Profile = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetRemove(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Remove = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetScanMode(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.ScanMode = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetWidth(v string) *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Width = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobScheduleConfig struct {
	// Pipeline ID.
	//
	// example:
	//
	// e37ebee5d98b4781897f6086e89f9c56
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// Job priority. Higher numbers indicate higher priority. Valid range: 1–10.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) SetPipelineId(v string) *GetTranscodeJobResponseBodyTranscodeParentJobScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) SetPriority(v int32) *GetTranscodeJobResponseBodyTranscodeParentJobScheduleConfig {
	s.Priority = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList struct {
	// Job creation time.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Job completion time.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// Job input group. A single input represents a transcoding job. Multiple inputs represent an audio-video merging job.
	InputGroup []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup `json:"InputGroup,omitempty" xml:"InputGroup,omitempty" type:"Repeated"`
	// Sub-job ID.
	//
	// example:
	//
	// 7d6a7e0d4db2457a8d45ff5d43e1bf0a
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Index of the sub-job within the entire job.
	//
	// example:
	//
	// 0
	JobIndex *int32 `json:"JobIndex,omitempty" xml:"JobIndex,omitempty"`
	// Job name.
	//
	// example:
	//
	// transcode-job
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Generated video media information.
	OutFileMeta *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta `json:"OutFileMeta,omitempty" xml:"OutFileMeta,omitempty" type:"Struct"`
	// Output media configuration.
	Output *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// Parent job ID.
	//
	// example:
	//
	// 8b2198504dd340b7b3c9842a74fc9baa
	ParentJobId *string `json:"ParentJobId,omitempty" xml:"ParentJobId,omitempty"`
	// Transcoding processing configuration.
	ProcessConfig *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig `json:"ProcessConfig,omitempty" xml:"ProcessConfig,omitempty" type:"Struct"`
	// Request ID at job submission.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Job scheduling information.
	ScheduleConfig *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// Transcoding job status
	//
	// - Init: Submitted.
	//
	// - Processing: Transcoding.
	//
	// - Success: Transcoding succeeded.
	//
	// - Fail: Transcoding failed.
	//
	// - Deleted: Deleted.
	//
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Job submission result.
	//
	// example:
	//
	// {}
	SubmitResultJson map[string]interface{} `json:"SubmitResultJson,omitempty" xml:"SubmitResultJson,omitempty"`
	// Job submission time.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// User data.
	//
	// example:
	//
	// user-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetInputGroup() []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup {
	return s.InputGroup
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetJobId() *string {
	return s.JobId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetJobIndex() *int32 {
	return s.JobIndex
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetName() *string {
	return s.Name
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetOutFileMeta() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta {
	return s.OutFileMeta
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetOutput() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput {
	return s.Output
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetParentJobId() *string {
	return s.ParentJobId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetProcessConfig() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	return s.ProcessConfig
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetScheduleConfig() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig {
	return s.ScheduleConfig
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetStatus() *string {
	return s.Status
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetSubmitResultJson() map[string]interface{} {
	return s.SubmitResultJson
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetUserData() *string {
	return s.UserData
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetCreateTime(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.CreateTime = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetFinishTime(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.FinishTime = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetInputGroup(v []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.InputGroup = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetJobId(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.JobId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetJobIndex(v int32) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.JobIndex = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetName(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.Name = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetOutFileMeta(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.OutFileMeta = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetOutput(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.Output = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetParentJobId(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.ParentJobId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetProcessConfig(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.ProcessConfig = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetRequestId(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.RequestId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetScheduleConfig(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.ScheduleConfig = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetStatus(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.Status = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetSubmitResultJson(v map[string]interface{}) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.SubmitResultJson = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetSubmitTime(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.SubmitTime = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetUserData(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.UserData = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) Validate() error {
	if s.InputGroup != nil {
		for _, item := range s.InputGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutFileMeta != nil {
		if err := s.OutFileMeta.Validate(); err != nil {
			return err
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.ProcessConfig != nil {
		if err := s.ProcessConfig.Validate(); err != nil {
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

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup struct {
	// Media stream URL. Specify this parameter only when transcoding a media stream.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	InputUrl *string `json:"InputUrl,omitempty" xml:"InputUrl,omitempty"`
	// Media value:
	//
	// - If Type is OSS, this field is a URL that supports OSS and HTTP protocols.
	//
	// - If Type is Media, this field is a media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// Media object type.
	//
	// Valid values:
	//
	// - OSS: OSS file.
	//
	// - Media: Media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) GetInputUrl() *string {
	return s.InputUrl
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) GetMedia() *string {
	return s.Media
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) GetType() *string {
	return s.Type
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) SetInputUrl(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup {
	s.InputUrl = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) SetMedia(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup {
	s.Media = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) SetType(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup {
	s.Type = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta struct {
	// Audio stream information.
	AudioStreamInfoList []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	// Basic file information.
	FileBasicInfo *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	// Video stream information.
	VideoStreamInfoList []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) GetAudioStreamInfoList() []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	return s.AudioStreamInfoList
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) GetFileBasicInfo() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	return s.FileBasicInfo
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) GetVideoStreamInfoList() []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	return s.VideoStreamInfoList
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) SetAudioStreamInfoList(v []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta {
	s.AudioStreamInfoList = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) SetFileBasicInfo(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta {
	s.FileBasicInfo = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) SetVideoStreamInfoList(v []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta {
	s.VideoStreamInfoList = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) Validate() error {
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

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList struct {
	// Bitrate.
	//
	// example:
	//
	// 0.f
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// Channel layout.
	//
	// example:
	//
	// stereo
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// Number of sound channels.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// Encoding format name.
	//
	// example:
	//
	// AAC (Advanced Audio Coding)
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// Encoding format.
	//
	// example:
	//
	// aac
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// Codec tag.
	//
	// example:
	//
	// 0x000f
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// Codec tag name.
	//
	// example:
	//
	// [15][0][0][0]
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// Codec time base.
	//
	// example:
	//
	// 1/44100
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// Duration (unit: seconds).
	//
	// example:
	//
	// 403.039989
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Stream index.
	//
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// Language.
	//
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Sample format.
	//
	// example:
	//
	// fltp
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// Sample rate (unit: Hz).
	//
	// example:
	//
	// 44100
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// Start time.
	//
	// example:
	//
	// 1.473556
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Time base.
	//
	// example:
	//
	// 1/90000
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetChannels() *string {
	return s.Channels
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetBitrate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetChannelLayout(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetChannels(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetCodecLongName(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetCodecName(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetCodecTag(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetCodecTagString(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetCodecTimeBase(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetDuration(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetIndex(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetLang(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetSampleFmt(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetSampleRate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetStartTime(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetTimebase(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo struct {
	// Video bitrate.
	//
	// example:
	//
	// 888.563
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// Video duration (unit: seconds).
	//
	// example:
	//
	// 403.039999
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// File name.
	//
	// example:
	//
	// file.m3u8
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// File size. Unit: Byte.
	//
	// example:
	//
	// 31737
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// File status.
	//
	// example:
	//
	// Normal
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// File type. Valid values: source_file, transcode_file
	//
	// example:
	//
	// source_file
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// File URL.
	//
	// example:
	//
	// http://bucket.oss-cn-shanghai.aliyuncs.com/path/to/file.m3u8
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// Video format name.
	//
	// example:
	//
	// hls,applehttp
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// Height.
	//
	// example:
	//
	// 478
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// Media asset ID.
	//
	// example:
	//
	// 486c2890096871edba6f81848c016303
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// File region.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Width.
	//
	// example:
	//
	// 848
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetBitrate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetDuration(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetFileName(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetFileSize(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetFileStatus(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetFileType(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetFileUrl(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetFormatName(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetHeight(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.Height = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetMediaId(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.MediaId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetRegion(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.Region = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetWidth(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.Width = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList struct {
	// Average frame rate.
	//
	// example:
	//
	// 25.0
	AvgFps *string `json:"Avg_fps,omitempty" xml:"Avg_fps,omitempty"`
	// Bitrate.
	//
	// example:
	//
	// 888.563
	BitRate *string `json:"Bit_rate,omitempty" xml:"Bit_rate,omitempty"`
	// Encoding format name.
	//
	// example:
	//
	// H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10
	CodecLongName *string `json:"Codec_long_name,omitempty" xml:"Codec_long_name,omitempty"`
	// Encoding format.
	//
	// example:
	//
	// h264
	CodecName *string `json:"Codec_name,omitempty" xml:"Codec_name,omitempty"`
	// Codec tag.
	//
	// example:
	//
	// 0x001b
	CodecTag *string `json:"Codec_tag,omitempty" xml:"Codec_tag,omitempty"`
	// Codec tag string.
	//
	// example:
	//
	// [27][0][0][0]
	CodecTagString *string `json:"Codec_tag_string,omitempty" xml:"Codec_tag_string,omitempty"`
	// Codec time base.
	//
	// example:
	//
	// 1/50
	CodecTimeBase *string `json:"Codec_time_base,omitempty" xml:"Codec_time_base,omitempty"`
	// Display aspect ratio.
	//
	// example:
	//
	// 16:9
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// Duration (unit: seconds).
	//
	// example:
	//
	// 403.039989
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Frame rate.
	//
	// example:
	//
	// 25.0
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// Whether B-frames exist.
	//
	// Valid values:
	//
	// - 0: No B-frames,
	//
	// - 1: One B-frame,
	//
	// - 2: Multiple consecutive B-frames.
	//
	// example:
	//
	// 2
	HasBFrames *string `json:"Has_b_frames,omitempty" xml:"Has_b_frames,omitempty"`
	// Height.
	//
	// example:
	//
	// 478
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// Stream index.
	//
	// example:
	//
	// 0
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// Language.
	//
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Encoding level.
	//
	// example:
	//
	// 31
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Total number of frames.
	//
	// example:
	//
	// 10040
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// Color storage format.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// Encoder preset.
	//
	// example:
	//
	// High
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Video rotation angle.
	//
	// Valid values: 0, 90, 180, 270.
	//
	// Default value: 0
	//
	// example:
	//
	// 0
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// Sample aspect ratio.
	//
	// example:
	//
	// 478:477
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// Start time.
	//
	// example:
	//
	// 1.473556
	StartTime *string `json:"Start_time,omitempty" xml:"Start_time,omitempty"`
	// Time base.
	//
	// example:
	//
	// 1/90000
	TimeBase *string `json:"Time_base,omitempty" xml:"Time_base,omitempty"`
	// Width.
	//
	// example:
	//
	// 848
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetAvgFps() *string {
	return s.AvgFps
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetBitRate() *string {
	return s.BitRate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetDar() *string {
	return s.Dar
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetHeight() *string {
	return s.Height
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetLevel() *string {
	return s.Level
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetRotate() *string {
	return s.Rotate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetSar() *string {
	return s.Sar
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetTimeBase() *string {
	return s.TimeBase
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetWidth() *string {
	return s.Width
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetAvgFps(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.AvgFps = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetBitRate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.BitRate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetCodecLongName(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetCodecName(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetCodecTag(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetCodecTagString(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetCodecTimeBase(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetDar(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetDuration(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetFps(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetHasBFrames(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetHeight(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetIndex(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetLang(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetLevel(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetNumFrames(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetPixFmt(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetProfile(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetRotate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetSar(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetStartTime(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetTimeBase(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.TimeBase = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetWidth(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput struct {
	// Media value:
	//
	// - If Type is OSS, this field is a URL that supports OSS and HTTP protocols.
	//
	// - If Type is Media, this field is a media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// Transcoded stream URL. Specify this parameter only when outputting to a media stream.
	//
	// example:
	//
	// oss://bucket/path/to/{MediaId}/{JobId}.mp4
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
	// Media object type.
	//
	// Valid values:
	//
	// - OSS: OSS file.
	//
	// - Media: Media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) GetMedia() *string {
	return s.Media
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) GetType() *string {
	return s.Type
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) SetMedia(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput {
	s.Media = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) SetOutputUrl(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput {
	s.OutputUrl = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) SetType(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput {
	s.Type = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig struct {
	// Multi-input merging configuration
	CombineConfigs []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs `json:"CombineConfigs,omitempty" xml:"CombineConfigs,omitempty" type:"Repeated"`
	// Encryption configuration.
	Encryption *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	// Image watermark configuration.
	ImageWatermarks []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks `json:"ImageWatermarks,omitempty" xml:"ImageWatermarks,omitempty" type:"Repeated"`
	// Whether to inherit tags from the input stream.
	//
	// This parameter does not take effect if the input is not a media asset.
	//
	// Default value: false
	//
	// example:
	//
	// true
	IsInheritTags *bool `json:"IsInheritTags,omitempty" xml:"IsInheritTags,omitempty"`
	// Subtitle embedding configuration.
	Subtitles []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// Text watermark configuration.
	TextWatermarks []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks `json:"TextWatermarks,omitempty" xml:"TextWatermarks,omitempty" type:"Repeated"`
	// Transcoding configuration.
	Transcode *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode `json:"Transcode,omitempty" xml:"Transcode,omitempty" type:"Struct"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GetCombineConfigs() []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs {
	return s.CombineConfigs
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GetEncryption() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption {
	return s.Encryption
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GetImageWatermarks() []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks {
	return s.ImageWatermarks
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GetIsInheritTags() *bool {
	return s.IsInheritTags
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GetSubtitles() []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles {
	return s.Subtitles
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GetTextWatermarks() []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks {
	return s.TextWatermarks
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GetTranscode() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode {
	return s.Transcode
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) SetCombineConfigs(v []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	s.CombineConfigs = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) SetEncryption(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	s.Encryption = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) SetImageWatermarks(v []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	s.ImageWatermarks = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) SetIsInheritTags(v bool) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	s.IsInheritTags = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) SetSubtitles(v []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	s.Subtitles = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) SetTextWatermarks(v []*GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	s.TextWatermarks = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) SetTranscode(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	s.Transcode = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) Validate() error {
	if s.CombineConfigs != nil {
		for _, item := range s.CombineConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			return err
		}
	}
	if s.ImageWatermarks != nil {
		for _, item := range s.ImageWatermarks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Subtitles != nil {
		for _, item := range s.Subtitles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TextWatermarks != nil {
		for _, item := range s.TextWatermarks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Transcode != nil {
		if err := s.Transcode.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs struct {
	// Audio stream index
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 or exclude
	AudioIndex *string `json:"AudioIndex,omitempty" xml:"AudioIndex,omitempty"`
	// Duration of the input stream. Default: video duration.
	//
	// example:
	//
	// 20.0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Start time of the input stream. Default: 0
	//
	// example:
	//
	// 0.0
	Start *float64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// Video stream index
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 or exclude
	VideoIndex *string `json:"VideoIndex,omitempty" xml:"VideoIndex,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) GetAudioIndex() *string {
	return s.AudioIndex
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) GetDuration() *float64 {
	return s.Duration
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) GetStart() *float64 {
	return s.Start
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) GetVideoIndex() *string {
	return s.VideoIndex
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) SetAudioIndex(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs {
	s.AudioIndex = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) SetDuration(v float64) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs {
	s.Duration = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) SetStart(v float64) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs {
	s.Start = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) SetVideoIndex(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs {
	s.VideoIndex = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption struct {
	// Ciphertext of the encryption key for standard encryption.
	//
	// example:
	//
	// MTYi00NDU0LTg5O****
	CipherText *string `json:"CipherText,omitempty" xml:"CipherText,omitempty"`
	// Decryption service endpoint for standard encryption.
	//
	// example:
	//
	// https://sample.com/path?CipherText=MTYi00NDU0LTg5O****
	DecryptKeyUri *string `json:"DecryptKeyUri,omitempty" xml:"DecryptKeyUri,omitempty"`
	// Encryption type.
	//
	// example:
	//
	// PrivateEncryption
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// Key service type. Only KMS and Base64 are supported.
	//
	// example:
	//
	// KMS
	KeyServiceType *string `json:"KeyServiceType,omitempty" xml:"KeyServiceType,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) GetCipherText() *string {
	return s.CipherText
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) GetDecryptKeyUri() *string {
	return s.DecryptKeyUri
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) GetEncryptType() *string {
	return s.EncryptType
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) GetKeyServiceType() *string {
	return s.KeyServiceType
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) SetCipherText(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption {
	s.CipherText = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) SetDecryptKeyUri(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption {
	s.DecryptKeyUri = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) SetEncryptType(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption {
	s.EncryptType = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) SetKeyServiceType(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption {
	s.KeyServiceType = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks struct {
	// Override parameters. If specified, these parameters override the corresponding template parameters.
	OverwriteParams *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// Template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) GetOverwriteParams() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	return s.OverwriteParams
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) SetOverwriteParams(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks {
	s.OverwriteParams = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) SetTemplateId(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks {
	s.TemplateId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams struct {
	// Watermark position, x.
	//
	// example:
	//
	// 10
	Dx *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// Watermark position, y.
	//
	// example:
	//
	// 10
	Dy *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// Watermark image file.
	File *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The height of the image watermark.
	//
	// example:
	//
	// 32
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// Reference position: TopLeft, TopRight, BottomLeft, BottomRight.<br>
	//
	// Default value: TopLeft.<br>
	//
	// example:
	//
	// TopLeft
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// Display time settings.
	Timeline *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline `json:"Timeline,omitempty" xml:"Timeline,omitempty" type:"Struct"`
	// Width.
	//
	// example:
	//
	// 32
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GetDx() *string {
	return s.Dx
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GetDy() *string {
	return s.Dy
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GetFile() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile {
	return s.File
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GetHeight() *string {
	return s.Height
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GetReferPos() *string {
	return s.ReferPos
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GetTimeline() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline {
	return s.Timeline
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GetWidth() *string {
	return s.Width
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) SetDx(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	s.Dx = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) SetDy(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	s.Dy = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) SetFile(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	s.File = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) SetHeight(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	s.Height = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) SetReferPos(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	s.ReferPos = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) SetTimeline(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	s.Timeline = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) SetWidth(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	s.Width = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	if s.Timeline != nil {
		if err := s.Timeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile struct {
	// Media value:
	//
	// - If Type is OSS, this field is a URL that supports OSS and HTTP protocols.
	//
	// - If Type is Media, this field is a media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// Media object type.
	//
	// Valid values:
	//
	// - OSS: OSS file.
	//
	// - Media: Media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) GetMedia() *string {
	return s.Media
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) GetType() *string {
	return s.Type
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) SetMedia(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile {
	s.Media = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) SetType(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile {
	s.Type = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline struct {
	// Display duration in seconds or "ToEND".
	//
	// example:
	//
	// ToEND
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Start time.
	//
	// example:
	//
	// 00:00:05
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) GetDuration() *string {
	return s.Duration
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) GetStart() *string {
	return s.Start
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) SetDuration(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline {
	s.Duration = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) SetStart(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline {
	s.Start = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles struct {
	// Override parameters. If specified, these parameters override the corresponding template parameters.
	OverwriteParams *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// Template ID
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) GetOverwriteParams() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams {
	return s.OverwriteParams
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) SetOverwriteParams(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles {
	s.OverwriteParams = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) SetTemplateId(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles {
	s.TemplateId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams struct {
	// File encoding format.
	//
	// example:
	//
	// UTF-8
	CharEnc *string `json:"CharEnc,omitempty" xml:"CharEnc,omitempty"`
	// Subtitle file.
	File *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// Subtitle file format.
	//
	// example:
	//
	// vtt
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) GetCharEnc() *string {
	return s.CharEnc
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) GetFile() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile {
	return s.File
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) GetFormat() *string {
	return s.Format
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) SetCharEnc(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams {
	s.CharEnc = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) SetFile(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams {
	s.File = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) SetFormat(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams {
	s.Format = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile struct {
	// Media value:
	//
	// - If Type is OSS, this field is a URL that supports OSS and HTTP protocols.
	//
	// - If Type is Media, this field is a media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// Media object type.
	//
	// Valid values:
	//
	// - OSS: OSS file.
	//
	// - Media: Media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) GetMedia() *string {
	return s.Media
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) GetType() *string {
	return s.Type
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) SetMedia(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile {
	s.Media = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) SetType(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile {
	s.Type = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks struct {
	// Override parameters. If specified, these parameters override the corresponding template parameters.
	OverwriteParams *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// Template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) GetOverwriteParams() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	return s.OverwriteParams
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) SetOverwriteParams(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks {
	s.OverwriteParams = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) SetTemplateId(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks {
	s.TemplateId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams struct {
	// Adjust font size based on output video size. Valid values: true / false. Default: false.
	//
	// example:
	//
	// false
	Adaptive *string `json:"Adaptive,omitempty" xml:"Adaptive,omitempty"`
	// Border color.
	//
	// example:
	//
	// #006400
	BorderColor *string `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// Border width.
	//
	// example:
	//
	// 0
	BorderWidth *int32 `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	// Watermark text. Do not Base64 encode. The string must be UTF-8 encoded.
	//
	// example:
	//
	// Test watermark
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Transparency.
	//
	// example:
	//
	// 1.0
	FontAlpha *string `json:"FontAlpha,omitempty" xml:"FontAlpha,omitempty"`
	// Color.
	//
	// example:
	//
	// #006400
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// Font.
	//
	// example:
	//
	// SimSun
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// Font size.
	//
	// example:
	//
	// 16
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// Watermark position, distance from the left.
	//
	// example:
	//
	// 10
	Left *string `json:"Left,omitempty" xml:"Left,omitempty"`
	// Watermark position, distance from the top.
	//
	// example:
	//
	// 10
	Top *string `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetAdaptive() *string {
	return s.Adaptive
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetBorderColor() *string {
	return s.BorderColor
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetBorderWidth() *int32 {
	return s.BorderWidth
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetContent() *string {
	return s.Content
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetFontAlpha() *string {
	return s.FontAlpha
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetFontColor() *string {
	return s.FontColor
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetFontName() *string {
	return s.FontName
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetFontSize() *int32 {
	return s.FontSize
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetLeft() *string {
	return s.Left
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetTop() *string {
	return s.Top
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetAdaptive(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.Adaptive = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetBorderColor(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.BorderColor = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetBorderWidth(v int32) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.BorderWidth = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetContent(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.Content = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetFontAlpha(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.FontAlpha = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetFontColor(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.FontColor = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetFontName(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.FontName = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetFontSize(v int32) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.FontSize = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetLeft(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.Left = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetTop(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.Top = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode struct {
	// Override parameters. If specified, these parameters override the corresponding template parameters.
	OverwriteParams *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// Template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) GetOverwriteParams() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams {
	return s.OverwriteParams
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) SetOverwriteParams(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode {
	s.OverwriteParams = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) SetTemplateId(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode {
	s.TemplateId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams struct {
	// Audio settings.
	Audio *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// Container format settings.
	Container *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// Muxing settings.
	MuxConfig *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	Tags      map[string]*string                                                                                           `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Conditional transcoding parameters.
	TransConfig *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	// Video settings.
	Video *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) GetAudio() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	return s.Audio
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) GetContainer() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer {
	return s.Container
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) GetMuxConfig() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig {
	return s.MuxConfig
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) GetTags() map[string]*string {
	return s.Tags
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) GetTransConfig() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	return s.TransConfig
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) GetVideo() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	return s.Video
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) SetAudio(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams {
	s.Audio = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) SetContainer(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams {
	s.Container = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) SetMuxConfig(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams {
	s.MuxConfig = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) SetTags(v map[string]*string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams {
	s.Tags = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) SetTransConfig(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams {
	s.TransConfig = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) SetVideo(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams {
	s.Video = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) Validate() error {
	if s.Audio != nil {
		if err := s.Audio.Validate(); err != nil {
			return err
		}
	}
	if s.Container != nil {
		if err := s.Container.Validate(); err != nil {
			return err
		}
	}
	if s.MuxConfig != nil {
		if err := s.MuxConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TransConfig != nil {
		if err := s.TransConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio struct {
	// Audio bitrate of the output file.
	//
	// - Valid range: [8, 1000]
	//
	// - Unit: Kbps
	//
	// - Default value: 128
	//
	// example:
	//
	// 128
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// Number of sound channels.
	//
	// Default value: 2
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// Audio codec format: AAC, MP3, VORBIS, FLAC.
	//
	// Default value: AAC
	//
	// Default value: AAC.
	//
	// example:
	//
	// AAC
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// Audio encoding preset.
	//
	// When Codec is AAC, valid values are aac_low, aac_he, aac_he_v2, aac_ld, aac_eld.
	//
	// example:
	//
	// aac_low
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Whether to remove the audio stream.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// Sample rate.
	//
	// - Supported values: 22050, 32000, 44100, 48000, 96000,
	//
	// - Unit: Hz
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	// Volume control.
	Volume *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Struct"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GetChannels() *string {
	return s.Channels
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GetCodec() *string {
	return s.Codec
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GetProfile() *string {
	return s.Profile
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GetRemove() *string {
	return s.Remove
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GetVolume() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume {
	return s.Volume
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) SetBitrate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	s.Bitrate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) SetChannels(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	s.Channels = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) SetCodec(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	s.Codec = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) SetProfile(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	s.Profile = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) SetRemove(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	s.Remove = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) SetSamplerate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	s.Samplerate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) SetVolume(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	s.Volume = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) Validate() error {
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume struct {
	// Target volume.
	//
	// example:
	//
	// -6
	IntegratedLoudnessTarget *string `json:"IntegratedLoudnessTarget,omitempty" xml:"IntegratedLoudnessTarget,omitempty"`
	// Volume range.
	//
	// example:
	//
	// 8
	LoudnessRangeTarget *string `json:"LoudnessRangeTarget,omitempty" xml:"LoudnessRangeTarget,omitempty"`
	// Volume adjustment method.
	//
	// example:
	//
	// auto
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// Peak volume.
	//
	// example:
	//
	// -1
	TruePeak *string `json:"TruePeak,omitempty" xml:"TruePeak,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) GetIntegratedLoudnessTarget() *string {
	return s.IntegratedLoudnessTarget
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) GetLoudnessRangeTarget() *string {
	return s.LoudnessRangeTarget
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) GetMethod() *string {
	return s.Method
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) GetTruePeak() *string {
	return s.TruePeak
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) SetIntegratedLoudnessTarget(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.IntegratedLoudnessTarget = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) SetLoudnessRangeTarget(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.LoudnessRangeTarget = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) SetMethod(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.Method = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) SetTruePeak(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.TruePeak = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer struct {
	// Container format.
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer) GetFormat() *string {
	return s.Format
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer) SetFormat(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer {
	s.Format = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig struct {
	// Segmentation settings.
	Segment *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig) GetSegment() *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	return s.Segment
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig) SetSegment(v *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig {
	s.Segment = v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig) Validate() error {
	if s.Segment != nil {
		if err := s.Segment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment struct {
	// Segment length.
	//
	// example:
	//
	// 10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Forced segmentation time points.
	//
	// example:
	//
	// 2,3
	ForceSegTime *string `json:"ForceSegTime,omitempty" xml:"ForceSegTime,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GetForceSegTime() *string {
	return s.ForceSegTime
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) SetDuration(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) SetForceSegTime(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	s.ForceSegTime = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig struct {
	// Resolution adjustment method. This parameter takes effect only when both Width and Height are specified. It can be used together with LongShortMode.
	//
	// Valid values: rescale, crop, pad, none.
	//
	// Default value: none.
	//
	// Example: See how to set resolution.
	//
	// example:
	//
	// none
	AdjDarMethod *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	// Whether to check audio bitrate. IsCheckAudioBitrate and IsCheckAudioBitrateFail are mutually exclusive. IsCheckAudioBitrateFail takes precedence.
	//
	// - true: Check. If the input audio bitrate is lower than the output setting, transcode using the input audio bitrate.
	//
	// - false: Do not check.
	//
	// Default value:
	//
	// - If this parameter is empty and the codec differs from the input source: false.
	//
	// - If this parameter is empty and the codec matches the input source: true.
	//
	// example:
	//
	// true
	IsCheckAudioBitrate *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	// Whether to check audio bitrate. IsCheckAudioBitrate and IsCheckAudioBitrateFail are mutually exclusive. This parameter takes precedence.
	//
	// - true: Check. If the input audio bitrate is lower than the output setting, return a transcoding failure.
	//
	// - false: Do not check.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	// Whether to check video resolution. IsCheckReso and IsCheckResoFail are mutually exclusive. IsCheckResoFail takes precedence.
	//
	// - true: Check. If the input video resolution (width or height) is smaller than the output setting, transcode using the input video resolution.
	//
	// - false: Do not check.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckReso *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	// Whether to check video resolution. IsCheckReso and IsCheckResoFail are mutually exclusive. This parameter takes precedence.
	//
	// - true: Check. If the input video resolution (width or height) is smaller than the output setting, return a transcoding failure.
	//
	// - false: Do not check.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckResoFail *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	// Whether to check video bitrate. IsCheckVideoBitrate and IsCheckVideoBitrateFail are mutually exclusive. IsCheckVideoBitrateFail takes precedence.
	//
	// - true: Check. If the input video bitrate is lower than the output setting, transcode using the input video bitrate.
	//
	// - false: Do not check.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckVideoBitrate *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	// Whether to check video bitrate. IsCheckVideoBitrate and IsCheckVideoBitrateFail are mutually exclusive. This parameter takes precedence.
	//
	// - true: Check. If the input video bitrate is lower than the output setting, return a transcoding failure.
	//
	// - false: Do not check.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	// Video transcoding mode. Valid values:
	//
	// - onepass: generally used for ABR. Faster encoding than twopass.
	//
	// - twopass: generally used for VBR. Slower encoding than onepass.
	//
	// - CBR: constant bitrate mode.
	//
	// Default value: onepass.
	//
	// example:
	//
	// onepass
	TransMode *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetAdjDarMethod() *string {
	return s.AdjDarMethod
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckAudioBitrate() *string {
	return s.IsCheckAudioBitrate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckAudioBitrateFail() *string {
	return s.IsCheckAudioBitrateFail
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckReso() *string {
	return s.IsCheckReso
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckResoFail() *string {
	return s.IsCheckResoFail
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckVideoBitrate() *string {
	return s.IsCheckVideoBitrate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckVideoBitrateFail() *string {
	return s.IsCheckVideoBitrateFail
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetAdjDarMethod(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.AdjDarMethod = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckAudioBitrate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckAudioBitrateFail(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckReso(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckResoFail(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckVideoBitrate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckVideoBitrateFail(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetTransMode(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.TransMode = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo struct {
	// Maximum ABR bitrate (valid only for narrow-high 1):
	//
	// - Valid range: [10, 50000].
	//
	// - Unit: Kbps.
	//
	// example:
	//
	// 6000
	AbrMax *string `json:"AbrMax,omitempty" xml:"AbrMax,omitempty"`
	// Video average bitrate.
	//
	// - Valid range: [10, 50000].
	//
	// - Unit: Kbps.
	//
	// example:
	//
	// 3000
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// Buffer size:
	//
	// - Valid range: [1000, 128000].
	//
	// - Default value: 6000.
	//
	// - Unit: Kb.
	//
	// example:
	//
	// 6000
	Bufsize *string `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	// Encoding format.
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// Bitrate-quality control factor.
	//
	// - Valid range: [0, 51].
	//
	// - Default value: 23 for H.264, 26 for H.265.
	//
	// If Crf is set, the Bitrate setting is ignored.
	//
	// example:
	//
	// 23
	Crf *string `json:"Crf,omitempty" xml:"Crf,omitempty"`
	// Video cropping. Two methods are supported:
	//
	// - Automatically detect and crop black bars by setting this parameter to "border".
	//
	// - Custom cropping in the format: width:height:left:top.
	//
	// example:
	//
	// 1280:800:0:140
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// Frame rate.
	//
	// - Valid range: (0, 60].
	//
	//   If the input file frame rate exceeds 60, use 60.
	//
	// - Default value: input file frame rate.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// Maximum number of frames between keyframes.
	//
	// - Valid range: [1, 1080000].
	//
	// - Default value: 250.
	//
	// example:
	//
	// 250
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// Height.
	//
	// - Valid range: [128, 4096].
	//
	// - Unit: px.
	//
	// Default value: original video height.
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// Whether to enable automatic rotation (long-short edge mode).
	//
	// example:
	//
	// false
	LongShortMode *string `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
	// Peak video bitrate:
	//
	// - Valid range: [10, 50000].
	//
	// - Unit: Kbps.
	//
	// example:
	//
	// 9000
	Maxrate *string `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	// Add black bars to the video in the format: width:height:left:top.
	//
	// example:
	//
	// 1280:800:0:140
	Pad *string `json:"Pad,omitempty" xml:"Pad,omitempty"`
	// Video color format.
	//
	// Valid values: yuv420p, yuvj420p, and other standard color formats.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// Video encoder preset. This parameter is supported only for H.264.
	//
	// Supported values: veryfast, fast, medium, slow, slower.
	//
	// Default value: medium.
	//
	// Default value: medium.
	//
	// example:
	//
	// medium
	Preset *string `json:"Preset,omitempty" xml:"Preset,omitempty"`
	// Encoding profile.
	//
	// Supported values: baseline, main, high.
	//
	// - baseline: for mobile devices.
	//
	// - main: for standard-resolution devices.
	//
	// - high: for high-resolution devices.
	//
	// Default value: high.
	//
	// example:
	//
	// Main
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Whether to remove the video.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// Scan mode.
	//
	// Supported values: interlaced, progressive.
	//
	// example:
	//
	// progressive
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// Width.
	//
	// - Valid range: [128, 4096].
	//
	// - Unit: px.
	//
	// Default value: original video width.
	//
	// example:
	//
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetAbrMax() *string {
	return s.AbrMax
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetCodec() *string {
	return s.Codec
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetCrf() *string {
	return s.Crf
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetCrop() *string {
	return s.Crop
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetFps() *string {
	return s.Fps
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetGop() *string {
	return s.Gop
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetHeight() *string {
	return s.Height
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetLongShortMode() *string {
	return s.LongShortMode
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetPad() *string {
	return s.Pad
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetPreset() *string {
	return s.Preset
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetProfile() *string {
	return s.Profile
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetRemove() *string {
	return s.Remove
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetWidth() *string {
	return s.Width
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetAbrMax(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.AbrMax = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetBitrate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Bitrate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetBufsize(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Bufsize = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetCodec(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Codec = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetCrf(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Crf = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetCrop(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Crop = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetFps(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Fps = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetGop(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Gop = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetHeight(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Height = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetLongShortMode(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.LongShortMode = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetMaxrate(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Maxrate = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetPad(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Pad = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetPixFmt(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.PixFmt = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetPreset(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Preset = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetProfile(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Profile = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetRemove(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Remove = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetScanMode(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.ScanMode = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetWidth(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Width = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig struct {
	// Pipeline ID.
	//
	// example:
	//
	// e37ebee5d98b4781897f6086e89f9c56
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// Job priority. Higher numbers indicate higher priority. Valid range: 1–10.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) SetPipelineId(v string) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) SetPriority(v int32) *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig {
	s.Priority = &v
	return s
}

func (s *GetTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) Validate() error {
	return dara.Validate(s)
}
