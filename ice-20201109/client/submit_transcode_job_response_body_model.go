// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTranscodeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitTranscodeJobResponseBody
	GetRequestId() *string
	SetTranscodeParentJob(v *SubmitTranscodeJobResponseBodyTranscodeParentJob) *SubmitTranscodeJobResponseBody
	GetTranscodeParentJob() *SubmitTranscodeJobResponseBodyTranscodeParentJob
}

type SubmitTranscodeJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// TranscodeParentJobWithSubJobDTO
	TranscodeParentJob *SubmitTranscodeJobResponseBodyTranscodeParentJob `json:"TranscodeParentJob,omitempty" xml:"TranscodeParentJob,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTranscodeJobResponseBody) GetTranscodeParentJob() *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	return s.TranscodeParentJob
}

func (s *SubmitTranscodeJobResponseBody) SetRequestId(v string) *SubmitTranscodeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBody) SetTranscodeParentJob(v *SubmitTranscodeJobResponseBodyTranscodeParentJob) *SubmitTranscodeJobResponseBody {
	s.TranscodeParentJob = v
	return s
}

func (s *SubmitTranscodeJobResponseBody) Validate() error {
	if s.TranscodeParentJob != nil {
		if err := s.TranscodeParentJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJob struct {
	// The time when the job was created. The format is yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the job finished. The format is yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The input group for the job. A single input creates a transcoding job. Multiple inputs create a job to merge audio and video streams.
	InputGroup []*SubmitTranscodeJobResponseBodyTranscodeParentJobInputGroup `json:"InputGroup,omitempty" xml:"InputGroup,omitempty" type:"Repeated"`
	// The number of sub-jobs.
	//
	// example:
	//
	// 1
	JobCount *int32 `json:"JobCount,omitempty" xml:"JobCount,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// transcode-job
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output group of the job.
	OutputGroup []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroup `json:"OutputGroup,omitempty" xml:"OutputGroup,omitempty" type:"Repeated"`
	// The ID of the parent job.
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
	// The ID of the request.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job scheduling configuration.
	ScheduleConfig *SubmitTranscodeJobResponseBodyTranscodeParentJobScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The status of the job. Success: At least one sub-job succeeded after all sub-jobs are complete. Fail: All sub-jobs failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the job was submitted. The format is yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// The list of sub-jobs.
	TranscodeJobList []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList `json:"TranscodeJobList,omitempty" xml:"TranscodeJobList,omitempty" type:"Repeated"`
	// The source of the job. Valid values: \\`API\\`, \\`WorkFlow\\`, and \\`Console\\`.
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

func (s SubmitTranscodeJobResponseBodyTranscodeParentJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJob) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetInputGroup() []*SubmitTranscodeJobResponseBodyTranscodeParentJobInputGroup {
	return s.InputGroup
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetJobCount() *int32 {
	return s.JobCount
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetName() *string {
	return s.Name
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetOutputGroup() []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroup {
	return s.OutputGroup
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetParentJobId() *string {
	return s.ParentJobId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetPercent() *int32 {
	return s.Percent
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetScheduleConfig() *SubmitTranscodeJobResponseBodyTranscodeParentJobScheduleConfig {
	return s.ScheduleConfig
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetStatus() *string {
	return s.Status
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetTranscodeJobList() []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	return s.TranscodeJobList
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) GetUserData() *string {
	return s.UserData
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetCreateTime(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.CreateTime = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetFinishTime(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.FinishTime = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetInputGroup(v []*SubmitTranscodeJobResponseBodyTranscodeParentJobInputGroup) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.InputGroup = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetJobCount(v int32) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.JobCount = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetName(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.Name = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetOutputGroup(v []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroup) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.OutputGroup = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetParentJobId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.ParentJobId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetPercent(v int32) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.Percent = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetRequestId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.RequestId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetScheduleConfig(v *SubmitTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.ScheduleConfig = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetStatus(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.Status = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetSubmitTime(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.SubmitTime = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetTranscodeJobList(v []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.TranscodeJobList = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetTriggerSource(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.TriggerSource = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) SetUserData(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJob {
	s.UserData = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJob) Validate() error {
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

type SubmitTranscodeJobResponseBodyTranscodeParentJobInputGroup struct {
	// The value of the media asset:
	//
	// - If type is OSS, this is a URL. Both the OSS and HTTP protocols are supported.
	//
	// - If type is Media, this is the media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the media object.
	//
	// Valid values:
	//
	// - OSS: an OSS file.
	//
	// - Media: a media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobInputGroup) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobInputGroup) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobInputGroup) GetMedia() *string {
	return s.Media
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobInputGroup) GetType() *string {
	return s.Type
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobInputGroup) SetMedia(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobInputGroup {
	s.Media = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobInputGroup) SetType(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobInputGroup {
	s.Type = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobInputGroup) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroup struct {
	// The output media configuration.
	Output *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The job processing configuration.
	ProcessConfig *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig `json:"ProcessConfig,omitempty" xml:"ProcessConfig,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroup) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroup) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroup) GetOutput() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput {
	return s.Output
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroup) GetProcessConfig() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig {
	return s.ProcessConfig
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroup) SetOutput(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroup {
	s.Output = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroup) SetProcessConfig(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroup {
	s.ProcessConfig = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroup) Validate() error {
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

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput struct {
	// The value of the media asset:
	//
	// - If type is set to OSS, the value is a URL. The OSS and HTTP protocols are supported.
	//
	// - If type is set to Media, the value is the media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The media object type.
	//
	// Valid values:
	//
	// - OSS: An OSS file.
	//
	// - Media: A media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) GetMedia() *string {
	return s.Media
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) GetType() *string {
	return s.Type
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) SetMedia(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput {
	s.Media = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) SetType(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput {
	s.Type = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupOutput) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig struct {
	// The settings for combining multiple input streams.
	CombineConfigs []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs `json:"CombineConfigs,omitempty" xml:"CombineConfigs,omitempty" type:"Repeated"`
	// The encryption configuration.
	Encryption *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	// The image watermark settings.
	ImageWatermarks []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks `json:"ImageWatermarks,omitempty" xml:"ImageWatermarks,omitempty" type:"Repeated"`
	// The caption burn-in configuration.
	Subtitles []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// The text watermark configurations.
	TextWatermarks []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks `json:"TextWatermarks,omitempty" xml:"TextWatermarks,omitempty" type:"Repeated"`
	// The transcoding configuration.
	Transcode *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode `json:"Transcode,omitempty" xml:"Transcode,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) GetCombineConfigs() []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs {
	return s.CombineConfigs
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) GetEncryption() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption {
	return s.Encryption
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) GetImageWatermarks() []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks {
	return s.ImageWatermarks
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) GetSubtitles() []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles {
	return s.Subtitles
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) GetTextWatermarks() []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks {
	return s.TextWatermarks
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) GetTranscode() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode {
	return s.Transcode
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) SetCombineConfigs(v []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig {
	s.CombineConfigs = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) SetEncryption(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig {
	s.Encryption = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) SetImageWatermarks(v []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig {
	s.ImageWatermarks = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) SetSubtitles(v []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig {
	s.Subtitles = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) SetTextWatermarks(v []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig {
	s.TextWatermarks = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) SetTranscode(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig {
	s.Transcode = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfig) Validate() error {
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

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs struct {
	// The index of the audio stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 或 exclude
	AudioIndex *string `json:"AudioIndex,omitempty" xml:"AudioIndex,omitempty"`
	// The duration of the input stream. By default, this is the duration of the video.
	//
	// example:
	//
	// 20.0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start time of the input stream. The default value is 0.
	//
	// example:
	//
	// 0.0
	Start *float64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The index of the video stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 或 exclude
	VideoIndex *string `json:"VideoIndex,omitempty" xml:"VideoIndex,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) GetAudioIndex() *string {
	return s.AudioIndex
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) GetDuration() *float64 {
	return s.Duration
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) GetStart() *float64 {
	return s.Start
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) GetVideoIndex() *string {
	return s.VideoIndex
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) SetAudioIndex(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs {
	s.AudioIndex = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) SetDuration(v float64) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs {
	s.Duration = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) SetStart(v float64) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs {
	s.Start = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) SetVideoIndex(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs {
	s.VideoIndex = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption struct {
	// The ciphertext of the key for standard encryption.
	//
	// example:
	//
	// MTYi00NDU0LTg5O****
	CipherText *string `json:"CipherText,omitempty" xml:"CipherText,omitempty"`
	// The decryption endpoint for standard encryption.
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
	// The type of the key service. Only KMS and Base64 are supported.
	//
	// example:
	//
	// KMS
	KeyServiceType *string `json:"KeyServiceType,omitempty" xml:"KeyServiceType,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) GetCipherText() *string {
	return s.CipherText
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) GetDecryptKeyUri() *string {
	return s.DecryptKeyUri
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) GetEncryptType() *string {
	return s.EncryptType
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) GetKeyServiceType() *string {
	return s.KeyServiceType
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) SetCipherText(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption {
	s.CipherText = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) SetDecryptKeyUri(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption {
	s.DecryptKeyUri = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) SetEncryptType(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption {
	s.EncryptType = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) SetKeyServiceType(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption {
	s.KeyServiceType = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks struct {
	// The parameters that, when specified, overwrite the corresponding parameters in the template.
	OverwriteParams *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) GetOverwriteParams() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	return s.OverwriteParams
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) SetOverwriteParams(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks {
	s.OverwriteParams = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) SetTemplateId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks {
	s.TemplateId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams struct {
	// The horizontal offset of the watermark image relative to the output video.
	//
	// Default: 0
	//
	// The value can be specified in two formats:
	//
	// - An integer that specifies the offset in pixels.
	//
	//   - Range: [8, 4096]
	//
	//   - Unit: px
	//
	// - A decimal that specifies the ratio of the horizontal offset to the width of the output video.
	//
	//   - Range: (0, 1)
	//
	//   - The value can have up to four decimal places, such as 0.9999. The system automatically discards any digits beyond the fourth decimal place.
	//
	// example:
	//
	// 10
	Dx *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// The vertical offset of the watermark image relative to the output video.
	//
	// Default value: 0.
	//
	// The value can be in one of the following two formats:
	//
	// - An integer that specifies the offset in pixels.
	//
	//   - Range: [8, 4096].
	//
	//   - Unit: px.
	//
	// - A decimal that specifies the ratio of the vertical offset to the output video height.
	//
	//   - Range: (0, 1).
	//
	//   - The value supports up to four decimal places, such as 0.9999. Any additional digits are automatically discarded.
	//
	// example:
	//
	// 10
	Dy *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The image file for the watermark.
	File *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The height of the image watermark on the output video.
	//
	// The value can be specified in two ways:
	//
	// - An integer that represents the watermark height in pixels.
	//
	//   - Range: [8, 4096].
	//
	//   - Unit: px.
	//
	// - A decimal that represents the watermark height as a ratio of the output video\\"s height.
	//
	//   - Range: (0, 1).
	//
	//   - The value supports up to four decimal places, such as 0.9999. Digits beyond the fourth decimal place are automatically discarded.
	//
	// example:
	//
	// 32
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The position of the watermark.
	//
	// - Valid values: TopRight, TopLeft, BottomRight, and BottomLeft.
	//
	// - Default value: TopRight.
	//
	// example:
	//
	// TopLeft
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// The display time settings for the dynamic watermark.
	Timeline *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline `json:"Timeline,omitempty" xml:"Timeline,omitempty" type:"Struct"`
	// The width of the watermark image on the output video.
	//
	// The value can be specified in two formats:
	//
	// - An integer that specifies the width of the watermark image in pixels.
	//
	//   - Range: [8, 4096].
	//
	//   - Unit: px.
	//
	// - A decimal that represents the width of the watermark relative to the width of the output video.
	//
	//   - Range: (0, 1).
	//
	//   - The value supports up to four decimal places, such as 0.9999. Digits beyond the fourth decimal place are automatically discarded.
	//
	// example:
	//
	// 32
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GetDx() *string {
	return s.Dx
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GetDy() *string {
	return s.Dy
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GetFile() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile {
	return s.File
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GetHeight() *string {
	return s.Height
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GetReferPos() *string {
	return s.ReferPos
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GetTimeline() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline {
	return s.Timeline
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) GetWidth() *string {
	return s.Width
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) SetDx(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Dx = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) SetDy(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Dy = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) SetFile(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.File = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) SetHeight(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Height = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) SetReferPos(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.ReferPos = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) SetTimeline(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Timeline = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) SetWidth(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams {
	s.Width = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParams) Validate() error {
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

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile struct {
	// The value of the media asset:
	//
	// - If type is OSS, the value is a URL that supports the OSS and HTTP protocols.
	//
	// - If type is Media, the value is the media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The object type of the media asset.
	//
	// Valid values:
	//
	// - OSS: An OSS file.
	//
	// - Media: A media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) GetMedia() *string {
	return s.Media
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) GetType() *string {
	return s.Type
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) SetMedia(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile {
	s.Media = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) SetType(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile {
	s.Type = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline struct {
	// The duration of the watermark.
	//
	// - Valid values: [Number, ToEND]
	//
	// - Default value: ToEND
	//
	// example:
	//
	// ToEND
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The time when the watermark appears.
	//
	// - Unit: seconds
	//
	// - The value must be numeric.
	//
	// - Default value: 0
	//
	// example:
	//
	// 00:00:05
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) GetDuration() *string {
	return s.Duration
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) GetStart() *string {
	return s.Start
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) SetDuration(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline {
	s.Duration = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) SetStart(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline {
	s.Start = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles struct {
	// The overwrite parameters. If specified, these parameters overwrite the corresponding parameters in the template.
	OverwriteParams *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) GetOverwriteParams() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams {
	return s.OverwriteParams
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) SetOverwriteParams(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles {
	s.OverwriteParams = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) SetTemplateId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles {
	s.TemplateId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams struct {
	// The file encoding format.
	//
	// example:
	//
	// UTF-8
	CharEnc *string `json:"CharEnc,omitempty" xml:"CharEnc,omitempty"`
	// The subtitle file.
	File *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The caption file format.
	//
	// example:
	//
	// vtt
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) GetCharEnc() *string {
	return s.CharEnc
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) GetFile() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile {
	return s.File
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) GetFormat() *string {
	return s.Format
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) SetCharEnc(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams {
	s.CharEnc = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) SetFile(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams {
	s.File = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) SetFormat(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams {
	s.Format = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParams) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile struct {
	// The value of the media asset. If the type is OSS, the value is a URL that supports the OSS and HTTP protocols. If the type is Media, the value is the media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the media object.
	//
	// Valid values:
	//
	// - OSS: an OSS file.
	//
	// - Media: a media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) GetMedia() *string {
	return s.Media
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) GetType() *string {
	return s.Type
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) SetMedia(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile {
	s.Media = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) SetType(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile {
	s.Type = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitlesOverwriteParamsFile) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks struct {
	// The overwrite parameters. If specified, these parameters overwrite the corresponding parameters in the template.
	OverwriteParams *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) GetOverwriteParams() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	return s.OverwriteParams
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) SetOverwriteParams(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks {
	s.OverwriteParams = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) SetTemplateId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks {
	s.TemplateId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarks) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams struct {
	// Adjusts the font size based on the size of the output video.
	//
	// - Valid values: true, false
	//
	// - default: false
	//
	// example:
	//
	// false
	Adaptive *string `json:"Adaptive,omitempty" xml:"Adaptive,omitempty"`
	// The outline color.
	//
	// Default: Black
	//
	// For more values, see BorderColor.
	//
	// example:
	//
	// #006400
	BorderColor *string `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// The width of the border.
	//
	// - Default: 0
	//
	// - Range: (0, 4096]
	//
	// example:
	//
	// 0
	BorderWidth *int32 `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	// The watermark text. The string must be UTF-8 encoded. Base64 encoding is not required.
	//
	// example:
	//
	// 测试水印
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The transparency of the font.
	//
	// - Range: (0, 1].
	//
	// - Default: 1.0.
	//
	// example:
	//
	// 1.0
	FontAlpha *string `json:"FontAlpha,omitempty" xml:"FontAlpha,omitempty"`
	// The font color.
	//
	// example:
	//
	// #006400
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// The font. Default: SimSun.
	//
	// example:
	//
	// SimSun
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// The font size.
	//
	// - Default value: 16
	//
	// - Range: (4, 120)
	//
	// example:
	//
	// 16
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// The left margin of the text.
	//
	// - Default: 0
	//
	// - Range: [0,4096]
	//
	// example:
	//
	// 10
	Left *string `json:"Left,omitempty" xml:"Left,omitempty"`
	// The top margin of the text.
	//
	// - Default: 0.
	//
	// - Range: [0,4096].
	//
	// example:
	//
	// 10
	Top *string `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetAdaptive() *string {
	return s.Adaptive
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetBorderColor() *string {
	return s.BorderColor
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetBorderWidth() *int32 {
	return s.BorderWidth
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetContent() *string {
	return s.Content
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontAlpha() *string {
	return s.FontAlpha
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontColor() *string {
	return s.FontColor
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontName() *string {
	return s.FontName
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetFontSize() *int32 {
	return s.FontSize
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetLeft() *string {
	return s.Left
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) GetTop() *string {
	return s.Top
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetAdaptive(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Adaptive = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetBorderColor(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.BorderColor = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetBorderWidth(v int32) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.BorderWidth = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetContent(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Content = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontAlpha(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontAlpha = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontColor(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontColor = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontName(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontName = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetFontSize(v int32) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.FontSize = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetLeft(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Left = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) SetTop(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams {
	s.Top = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTextWatermarksOverwriteParams) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode struct {
	// The overwrite parameters. If specified, these parameters overwrite the corresponding parameters in the template.
	OverwriteParams *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) GetOverwriteParams() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams {
	return s.OverwriteParams
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) SetOverwriteParams(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode {
	s.OverwriteParams = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) SetTemplateId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode {
	s.TemplateId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscode) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams struct {
	// Audio settings.
	Audio *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// The container format settings.
	Container *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The encapsulation settings.
	MuxConfig *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	// The conditional transcoding parameters.
	TransConfig *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	// Video settings
	Video *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) GetAudio() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	return s.Audio
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) GetContainer() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer {
	return s.Container
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) GetMuxConfig() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig {
	return s.MuxConfig
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) GetTransConfig() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	return s.TransConfig
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) GetVideo() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	return s.Video
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) SetAudio(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams {
	s.Audio = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) SetContainer(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams {
	s.Container = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) SetMuxConfig(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams {
	s.MuxConfig = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) SetTransConfig(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams {
	s.TransConfig = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) SetVideo(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams {
	s.Video = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParams) Validate() error {
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

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio struct {
	// The audio bitrate of the output file.
	//
	// - Value range: [8, 1000]
	//
	// - Unit: Kbps
	//
	// - Default value: 128
	//
	// example:
	//
	// 128
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The number of sound channels.
	//
	// Default value: 2.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The audio codec format. Valid values are AAC, MP3, VORBIS, and FLAC.
	//
	// Default value: AAC
	//
	// example:
	//
	// AAC
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The audio encoding preset.
	//
	// When Codec is set to AAC, valid values are aac_low, aac_he, aac_he_v2, aac_ld, and aac_eld.
	//
	// example:
	//
	// aac_low
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Specifies whether to remove the audio stream.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The sample rate.
	//
	// - Default value: 44100. Valid values: 22050, 32000, 44100, 48000, and 96000.
	//
	// - Unit: Hz
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	// The volume control settings.
	Volume *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetChannels() *string {
	return s.Channels
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetCodec() *string {
	return s.Codec
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetProfile() *string {
	return s.Profile
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetRemove() *string {
	return s.Remove
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) GetVolume() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	return s.Volume
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetBitrate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Bitrate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetChannels(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Channels = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetCodec(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Codec = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetProfile(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Profile = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetRemove(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Remove = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetSamplerate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Samplerate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) SetVolume(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio {
	s.Volume = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio) Validate() error {
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume struct {
	// The target volume.
	//
	// example:
	//
	// -6
	IntegratedLoudnessTarget *string `json:"IntegratedLoudnessTarget,omitempty" xml:"IntegratedLoudnessTarget,omitempty"`
	// The loudness range.
	//
	// example:
	//
	// 8
	LoudnessRangeTarget *string `json:"LoudnessRangeTarget,omitempty" xml:"LoudnessRangeTarget,omitempty"`
	// The volume adjustment method.
	//
	// example:
	//
	// auto
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The true peak.
	//
	// example:
	//
	// -1
	TruePeak *string `json:"TruePeak,omitempty" xml:"TruePeak,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetIntegratedLoudnessTarget() *string {
	return s.IntegratedLoudnessTarget
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetLoudnessRangeTarget() *string {
	return s.LoudnessRangeTarget
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetMethod() *string {
	return s.Method
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) GetTruePeak() *string {
	return s.TruePeak
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetIntegratedLoudnessTarget(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.IntegratedLoudnessTarget = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetLoudnessRangeTarget(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.LoudnessRangeTarget = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetMethod(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.Method = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) SetTruePeak(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.TruePeak = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudioVolume) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer struct {
	// The container format.
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer) GetFormat() *string {
	return s.Format
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer) SetFormat(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer {
	s.Format = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig struct {
	// Segment settings.
	Segment *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) GetSegment() *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	return s.Segment
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) SetSegment(v *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig {
	s.Segment = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig) Validate() error {
	if s.Segment != nil {
		if err := s.Segment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment struct {
	// The segment length.
	//
	// example:
	//
	// 10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The time points for forced segmentation.
	//
	// example:
	//
	// 2,3
	ForceSegTime *string `json:"ForceSegTime,omitempty" xml:"ForceSegTime,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GetForceSegTime() *string {
	return s.ForceSegTime
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) SetDuration(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) SetForceSegTime(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	s.ForceSegTime = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig struct {
	// The method to adjust the resolution. This setting takes effect only when both Width and Height are specified. It can be used with LongShortMode.
	//
	// Valid values: rescale, crop, pad, and none.
	//
	// Default value: none.
	//
	// example:
	//
	// none
	AdjDarMethod *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	// Specifies whether to check the audio bitrate. You can set either this parameter or IsCheckAudioBitrateFail. IsCheckAudioBitrateFail takes precedence.
	//
	// - true: Checks the audio bitrate. If the input audio bitrate is lower than the configured output bitrate, the service uses the input audio bitrate for transcoding.
	//
	// - false: Does not check the audio bitrate.
	//
	// Default value:
	//
	// - false: If this parameter is empty and the output codec is different from the input codec.
	//
	// - true: If this parameter is empty and the output codec is the same as the input codec.
	//
	// example:
	//
	// true
	IsCheckAudioBitrate *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	// Specifies whether to check the audio bitrate. You can set either IsCheckAudioBitrate or IsCheckAudioBitrateFail. This parameter has a higher priority.
	//
	// - true: The transcoding job fails if the input audio bitrate is lower than the output bitrate setting.
	//
	// - false: The audio bitrate is not checked.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	// Specifies whether to check the video resolution. The IsCheckReso and IsCheckResoFail parameters are mutually exclusive. IsCheckResoFail takes precedence.
	//
	// - true: Enables the resolution check. If the resolution (width or height) of the input video is lower than the output resolution, the transcoding job uses the input resolution.
	//
	// - false: Disables the resolution check.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckReso *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	// Specifies whether to check the video resolution. You can use either IsCheckReso or IsCheckResoFail, but not both. This parameter has a higher priority.
	//
	// - true: Checks the resolution. The transcoding job fails if the width or height of the input video is smaller than the output resolution.
	//
	// - false: Does not check the resolution.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckResoFail *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	// Specifies whether to check the video bitrate. You can set either IsCheckVideoBitrate or IsCheckVideoBitrateFail. IsCheckVideoBitrateFail has a higher priority.
	//
	// - true: Checks the bitrate. If the input video bitrate is lower than the output bitrate, the video is transcoded at the input bitrate.
	//
	// - false: Does not check the bitrate.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckVideoBitrate *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	// Specifies whether to check the video bitrate. This parameter and IsCheckVideoBitrate are mutually exclusive. IsCheckVideoBitrateFail has a higher priority.
	//
	// - true: Enables the check. The transcoding job fails if the input video bitrate is lower than the output bitrate setting.
	//
	// - false: Disables the check.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	// The video transcoding mode. Valid values:
	//
	// - onepass: Typically used for ABR. The encoding speed is faster than twopass.
	//
	// - twopass: Typically used for VBR. The encoding speed is slower than onepass.
	//
	// - CBR: Constant Bitrate mode.
	//
	// Default value: onepass.
	//
	// example:
	//
	// onepass
	TransMode *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetAdjDarMethod() *string {
	return s.AdjDarMethod
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckAudioBitrate() *string {
	return s.IsCheckAudioBitrate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckAudioBitrateFail() *string {
	return s.IsCheckAudioBitrateFail
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckReso() *string {
	return s.IsCheckReso
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckResoFail() *string {
	return s.IsCheckResoFail
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckVideoBitrate() *string {
	return s.IsCheckVideoBitrate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckVideoBitrateFail() *string {
	return s.IsCheckVideoBitrateFail
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetAdjDarMethod(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.AdjDarMethod = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckAudioBitrate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckAudioBitrateFail(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckReso(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckResoFail(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckVideoBitrate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckVideoBitrateFail(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) SetTransMode(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig {
	s.TransMode = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo struct {
	// The maximum bitrate for adaptive bitrate (ABR) streaming. This parameter is valid only for Narrowband HD 1.0.
	//
	// - Value range: [10, 50000]
	//
	// - Unit: Kbps
	//
	// example:
	//
	// 6000
	AbrMax *string `json:"AbrMax,omitempty" xml:"AbrMax,omitempty"`
	// The average video bitrate.
	//
	// - Value range: [10, 50000].
	//
	// - Unit: Kbps.
	//
	// example:
	//
	// 3000
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The buffer size.
	//
	// - Value range: [1000, 128000]
	//
	// - Default value: 6000
	//
	// - Unit: Kb
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
	// The Constant Rate Factor (CRF).
	//
	// - The value can be from 0 to 51.
	//
	// - The default value is 23 for H264 encoding and 26 for H265 encoding.
	//
	// If Crf is set, the Bitrate setting is ignored.
	//
	// example:
	//
	// 23
	Crf *string `json:"Crf,omitempty" xml:"Crf,omitempty"`
	// Crops the video frame.
	//
	// Two methods are available.
	//
	// - To automatically detect and crop black bars, set the parameter to "border".
	//
	// - To specify a custom crop area, use the width:height:left:top format.
	//
	//   Example: 1280:800:0:140
	//
	// example:
	//
	// 1280:800:0:140
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// The frame rate.
	//
	// - Valid values: (0, 60].
	//
	// - If the input file has a frame rate greater than 60, the frame rate is capped at 60.
	//
	// - Default: The frame rate of the input file.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The maximum number of frames between two keyframes.
	//
	// - The value must be in the range of [1, 1080000].
	//
	// - The default value is 250.
	//
	// example:
	//
	// 250
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height of the video.
	//
	// - Valid values: [128, 4096].
	//
	// - Unit: px.
	//
	// - Default value: The original height of the video.
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// Specifies whether to enable automatic rotation for landscape and portrait orientations (long and short edge pattern).
	//
	// example:
	//
	// false
	LongShortMode *string `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
	// The peak video bitrate.
	//
	// - Value range: [10, 50000]
	//
	// - Unit: Kbps
	//
	// example:
	//
	// 9000
	Maxrate *string `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	// Adds black bars to the video.
	//
	// - Format: width:height:left:top
	//
	// - Example: 1280:800:0:140
	//
	// example:
	//
	// 1280:800:0:140
	Pad *string `json:"Pad,omitempty" xml:"Pad,omitempty"`
	// The pixel format of the video. Valid values include standard formats such as yuv420p and yuvj420p.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The video algorithm preset. This parameter is supported only for H.264. Supported values are veryfast, fast, medium, slow, and slower. The default value is medium.
	//
	// example:
	//
	// medium
	Preset *string `json:"Preset,omitempty" xml:"Preset,omitempty"`
	// The encoding profile.
	//
	// Supported values are baseline, main, and high.
	//
	// - baseline: For mobile devices.
	//
	// - main: For standard-resolution devices.
	//
	// - high: For high-resolution devices.
	//
	// Default value: high.
	//
	// example:
	//
	// Main
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Specifies whether to remove the video.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The scan mode. Valid values are interlaced and progressive.
	//
	// example:
	//
	// progressive
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// The width of the output video.
	//
	// - Valid values: 128 to 4096.
	//
	// - Unit: px.
	//
	// - Default value: The original width of the video.
	//
	// example:
	//
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetAbrMax() *string {
	return s.AbrMax
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetCodec() *string {
	return s.Codec
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetCrf() *string {
	return s.Crf
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetCrop() *string {
	return s.Crop
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetFps() *string {
	return s.Fps
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetGop() *string {
	return s.Gop
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetHeight() *string {
	return s.Height
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetLongShortMode() *string {
	return s.LongShortMode
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetPad() *string {
	return s.Pad
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetPreset() *string {
	return s.Preset
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetProfile() *string {
	return s.Profile
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetRemove() *string {
	return s.Remove
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) GetWidth() *string {
	return s.Width
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetAbrMax(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.AbrMax = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetBitrate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Bitrate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetBufsize(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Bufsize = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetCodec(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Codec = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetCrf(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Crf = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetCrop(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Crop = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetFps(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Fps = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetGop(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Gop = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetHeight(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Height = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetLongShortMode(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.LongShortMode = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetMaxrate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Maxrate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetPad(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Pad = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetPixFmt(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.PixFmt = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetPreset(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Preset = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetProfile(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Profile = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetRemove(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Remove = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetScanMode(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.ScanMode = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) SetWidth(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo {
	s.Width = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsVideo) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobScheduleConfig struct {
	// The pipeline ID.
	//
	// example:
	//
	// e37ebee5d98b4781897f6086e89f9c56
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the job. A larger value indicates a higher priority. The value can be an integer from 1 to 10.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) SetPipelineId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) SetPriority(v int32) *SubmitTranscodeJobResponseBodyTranscodeParentJobScheduleConfig {
	s.Priority = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobScheduleConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList struct {
	// The time the job was created.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the job finished.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The input group for the job. A single input creates a transcoding job. Multiple inputs create a media merging job.
	InputGroup []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup `json:"InputGroup,omitempty" xml:"InputGroup,omitempty" type:"Repeated"`
	// The sub-job ID.
	//
	// example:
	//
	// 7d6a7e0d4db2457a8d45ff5d43e1bf0a
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The index of the sub-job within the entire job.
	//
	// example:
	//
	// 0
	JobIndex *int32 `json:"JobIndex,omitempty" xml:"JobIndex,omitempty"`
	// The job name.
	//
	// example:
	//
	// transcode-job
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The media information of the video generated by the job.
	OutFileMeta *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta `json:"OutFileMeta,omitempty" xml:"OutFileMeta,omitempty" type:"Struct"`
	// The output media configuration.
	Output *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The parent job ID.
	//
	// example:
	//
	// 8b2198504dd340b7b3c9842a74fc9baa
	ParentJobId *string `json:"ParentJobId,omitempty" xml:"ParentJobId,omitempty"`
	// The transcoding processing configuration.
	ProcessConfig *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig `json:"ProcessConfig,omitempty" xml:"ProcessConfig,omitempty" type:"Struct"`
	// The ID of the request to submit the job.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scheduling information for the job.
	ScheduleConfig *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The status of the transcoding job.
	//
	// - Init: The job is submitted.
	//
	// - Processing: The job is being transcoded.
	//
	// - Success: The transcoding is successful.
	//
	// - Fail: The transcoding failed.
	//
	// - Deleted: The job is deleted.
	//
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The result of the job submission.
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
	// User data.
	//
	// example:
	//
	// user-data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetFinishTime() *string {
	return s.FinishTime
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetInputGroup() []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup {
	return s.InputGroup
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetJobId() *string {
	return s.JobId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetJobIndex() *int32 {
	return s.JobIndex
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetName() *string {
	return s.Name
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetOutFileMeta() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta {
	return s.OutFileMeta
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetOutput() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput {
	return s.Output
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetParentJobId() *string {
	return s.ParentJobId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetProcessConfig() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	return s.ProcessConfig
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetScheduleConfig() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig {
	return s.ScheduleConfig
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetStatus() *string {
	return s.Status
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetSubmitResultJson() map[string]interface{} {
	return s.SubmitResultJson
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) GetUserData() *string {
	return s.UserData
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetCreateTime(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.CreateTime = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetFinishTime(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.FinishTime = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetInputGroup(v []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.InputGroup = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetJobId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.JobId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetJobIndex(v int32) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.JobIndex = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetName(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.Name = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetOutFileMeta(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.OutFileMeta = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetOutput(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.Output = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetParentJobId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.ParentJobId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetProcessConfig(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.ProcessConfig = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetRequestId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.RequestId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetScheduleConfig(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.ScheduleConfig = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetStatus(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.Status = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetSubmitResultJson(v map[string]interface{}) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.SubmitResultJson = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetSubmitTime(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.SubmitTime = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) SetUserData(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList {
	s.UserData = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList) Validate() error {
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

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup struct {
	// The input stream path:
	//
	// - This parameter takes effect only when Type is Media. It lets you select a specific file from the media asset as the input.
	//
	// - The system checks whether the input URL exists in the media asset.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	InputUrl *string `json:"InputUrl,omitempty" xml:"InputUrl,omitempty"`
	// The media value:
	//
	// - If Type is OSS, this is a URL that supports the OSS or HTTP protocol.
	//
	// - If Type is Media, this is a media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The media object type. Valid values:
	//
	// - OSS: An OSS file.
	//
	// - Media: A media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) GetInputUrl() *string {
	return s.InputUrl
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) GetMedia() *string {
	return s.Media
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) GetType() *string {
	return s.Type
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) SetInputUrl(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup {
	s.InputUrl = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) SetMedia(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup {
	s.Media = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) SetType(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup {
	s.Type = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta struct {
	// The audio stream information.
	AudioStreamInfoList []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	// Basic file information.
	FileBasicInfo *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	// The video stream information.
	VideoStreamInfoList []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) GetAudioStreamInfoList() []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	return s.AudioStreamInfoList
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) GetFileBasicInfo() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	return s.FileBasicInfo
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) GetVideoStreamInfoList() []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	return s.VideoStreamInfoList
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) SetAudioStreamInfoList(v []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta {
	s.AudioStreamInfoList = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) SetFileBasicInfo(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta {
	s.FileBasicInfo = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) SetVideoStreamInfoList(v []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta {
	s.VideoStreamInfoList = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta) Validate() error {
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

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList struct {
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
	// The encoder tag name.
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
	// The duration in seconds.
	//
	// example:
	//
	// 403.039989
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The index of the stream.
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
	// The sampling format.
	//
	// example:
	//
	// fltp
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// The sample rate in Hz.
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

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetChannels() *string {
	return s.Channels
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetSampleRate() *string {
	return s.SampleRate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) GetTimebase() *string {
	return s.Timebase
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetBitrate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetChannelLayout(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetChannels(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetCodecLongName(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetCodecName(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetCodecTag(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetCodecTagString(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetCodecTimeBase(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetDuration(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetIndex(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetLang(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.Lang = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetSampleFmt(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetSampleRate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetStartTime(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) SetTimebase(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo struct {
	// The video bitrate.
	//
	// example:
	//
	// 888.563
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The duration of the video in seconds.
	//
	// example:
	//
	// 403.039999
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// file.m3u8
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The size of the file in bytes.
	//
	// example:
	//
	// 31737
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The status of the file.
	//
	// example:
	//
	// Normal
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// The file type. Valid values: source_file and transcode_file.
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
	// 73e07de0f77171eca3fc7035d0b26402
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The region where the file is located.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The width of the output file.
	//
	// example:
	//
	// 848
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetDuration() *string {
	return s.Duration
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetFileName() *string {
	return s.FileName
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetFileStatus() *string {
	return s.FileStatus
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetFileType() *string {
	return s.FileType
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetFormatName() *string {
	return s.FormatName
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetHeight() *string {
	return s.Height
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetRegion() *string {
	return s.Region
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) GetWidth() *string {
	return s.Width
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetBitrate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetDuration(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetFileName(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetFileSize(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetFileStatus(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetFileType(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetFileUrl(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetFormatName(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetHeight(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.Height = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetMediaId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.MediaId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetRegion(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.Region = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) SetWidth(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo {
	s.Width = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList struct {
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
	// The encoding format tag.
	//
	// example:
	//
	// 0x001b
	CodecTag *string `json:"Codec_tag,omitempty" xml:"Codec_tag,omitempty"`
	// The text of the encoding format tag.
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
	// The display aspect ratio.
	//
	// example:
	//
	// 16:9
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// The duration in seconds.
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
	// Indicates whether B-frames exist.
	//
	// Valid values:
	//
	// - 0: No B-frames.
	//
	// - 1: One B-frame.
	//
	// - 2: Multiple consecutive B-frames.
	//
	// example:
	//
	// 2
	HasBFrames *string `json:"Has_b_frames,omitempty" xml:"Has_b_frames,omitempty"`
	// The height of the output video stream.
	//
	// example:
	//
	// 478
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The index of the stream.
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
	// The color storage format.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The encoder preset.
	//
	// example:
	//
	// High
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The rotation angle of the video. Valid values: 0, 90, 180, and 270. The default value is 0.
	//
	// example:
	//
	// 0
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The sample aspect ratio.
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
	// The video width.
	//
	// example:
	//
	// 848
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetAvgFps() *string {
	return s.AvgFps
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetBitRate() *string {
	return s.BitRate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetCodecName() *string {
	return s.CodecName
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetCodecTag() *string {
	return s.CodecTag
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetDar() *string {
	return s.Dar
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetDuration() *string {
	return s.Duration
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetFps() *string {
	return s.Fps
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetHeight() *string {
	return s.Height
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetIndex() *string {
	return s.Index
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetLang() *string {
	return s.Lang
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetLevel() *string {
	return s.Level
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetNumFrames() *string {
	return s.NumFrames
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetPixFmt() *string {
	return s.PixFmt
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetProfile() *string {
	return s.Profile
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetRotate() *string {
	return s.Rotate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetSar() *string {
	return s.Sar
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetTimeBase() *string {
	return s.TimeBase
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) GetWidth() *string {
	return s.Width
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetAvgFps(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.AvgFps = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetBitRate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.BitRate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetCodecLongName(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetCodecName(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetCodecTag(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetCodecTagString(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetCodecTimeBase(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetDar(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetDuration(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetFps(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetHasBFrames(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetHeight(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetIndex(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetLang(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetLevel(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetNumFrames(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetPixFmt(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetProfile(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetRotate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetSar(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetStartTime(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetTimeBase(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.TimeBase = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) SetWidth(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaVideoStreamInfoList) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput struct {
	// The value of the media asset:
	//
	// - If type is OSS, the value is a URL. The OSS and HTTP protocols are supported.
	//
	// - If type is Media, the value is the media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The path of the output stream.<br>
	//
	// This parameter is valid only when \\`Type\\` is set to \\`Media\\`. It lets you select a specific file from the media asset for output.<br>
	//
	// The following placeholders are supported:<br><br>
	//
	// - {MediaId}: The ID of the media asset.
	//
	// - {JobId}: The ID of the transcoding subtask.
	//
	// - {MediaBucket}: The bucket where the media asset is stored.
	//
	// - {ExtName}: The file extension. The value is the output format specified in the transcoding template.
	//
	// - {DestMd5}: The MD5 hash of the output file.<br>
	//
	//   Note:<br>
	//
	// 1. This parameter must contain the {MediaId} and {JobId} placeholders.
	//
	// 2. The output bucket is the same as the bucket where the media asset is stored.
	//
	// example:
	//
	// oss://bucket/path/to/{MediaId}/{JobId}.mp4
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
	// The media object type.
	//
	// Valid values:
	//
	// - OSS: An OSS file.
	//
	// - Media: A media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) GetMedia() *string {
	return s.Media
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) GetType() *string {
	return s.Type
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) SetMedia(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput {
	s.Media = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) SetOutputUrl(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput {
	s.OutputUrl = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) SetType(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput {
	s.Type = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig struct {
	// The configuration for mixing multiple input streams.
	CombineConfigs []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs `json:"CombineConfigs,omitempty" xml:"CombineConfigs,omitempty" type:"Repeated"`
	// The encryption configuration.
	Encryption *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	// The image watermark configuration.
	ImageWatermarks []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks `json:"ImageWatermarks,omitempty" xml:"ImageWatermarks,omitempty" type:"Repeated"`
	// The configuration for burning in captions.
	Subtitles []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// The text watermark configuration.
	TextWatermarks []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks `json:"TextWatermarks,omitempty" xml:"TextWatermarks,omitempty" type:"Repeated"`
	// The transcoding configuration.
	Transcode *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode `json:"Transcode,omitempty" xml:"Transcode,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GetCombineConfigs() []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs {
	return s.CombineConfigs
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GetEncryption() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption {
	return s.Encryption
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GetImageWatermarks() []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks {
	return s.ImageWatermarks
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GetSubtitles() []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles {
	return s.Subtitles
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GetTextWatermarks() []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks {
	return s.TextWatermarks
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) GetTranscode() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode {
	return s.Transcode
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) SetCombineConfigs(v []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	s.CombineConfigs = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) SetEncryption(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	s.Encryption = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) SetImageWatermarks(v []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	s.ImageWatermarks = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) SetSubtitles(v []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	s.Subtitles = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) SetTextWatermarks(v []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	s.TextWatermarks = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) SetTranscode(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig {
	s.Transcode = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig) Validate() error {
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

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs struct {
	// The index of the audio stream.
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
	// The start time of the input stream. The default value is 0.
	//
	// example:
	//
	// 0.0
	Start *float64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The index of the video stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 或 exclude
	VideoIndex *string `json:"VideoIndex,omitempty" xml:"VideoIndex,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) GetAudioIndex() *string {
	return s.AudioIndex
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) GetDuration() *float64 {
	return s.Duration
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) GetStart() *float64 {
	return s.Start
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) GetVideoIndex() *string {
	return s.VideoIndex
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) SetAudioIndex(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs {
	s.AudioIndex = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) SetDuration(v float64) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs {
	s.Duration = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) SetStart(v float64) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs {
	s.Start = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) SetVideoIndex(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs {
	s.VideoIndex = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption struct {
	// The ciphertext of the key for standard encryption.
	//
	// example:
	//
	// MTYi00NDU0LTg5O****
	CipherText *string `json:"CipherText,omitempty" xml:"CipherText,omitempty"`
	// The decryption service endpoint for standard encryption.
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
	// The type of the key service. Only KMS and Base64 are supported.
	//
	// example:
	//
	// KMS
	KeyServiceType *string `json:"KeyServiceType,omitempty" xml:"KeyServiceType,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) GetCipherText() *string {
	return s.CipherText
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) GetDecryptKeyUri() *string {
	return s.DecryptKeyUri
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) GetEncryptType() *string {
	return s.EncryptType
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) GetKeyServiceType() *string {
	return s.KeyServiceType
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) SetCipherText(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption {
	s.CipherText = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) SetDecryptKeyUri(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption {
	s.DecryptKeyUri = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) SetEncryptType(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption {
	s.EncryptType = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) SetKeyServiceType(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption {
	s.KeyServiceType = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks struct {
	// Parameters to overwrite. If you specify these, they replace the corresponding parameters in the template.
	OverwriteParams *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) GetOverwriteParams() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	return s.OverwriteParams
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) SetOverwriteParams(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks {
	s.OverwriteParams = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) SetTemplateId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks {
	s.TemplateId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams struct {
	// The horizontal offset of the image watermark relative to the output video. Default value: 0.
	//
	// Values can be one of the following:
	//
	// - Integer: The offset in pixels.
	//
	//   - Valid values: [8, 4096]
	//
	//   - Unit: px
	//
	// - Decimal: The ratio of the horizontal offset to the output video width.
	//
	//   - Valid values: (0, 1)
	//
	//   - Up to four decimal places are supported, such as 0.9999. Extra digits are automatically discarded.
	//
	// example:
	//
	// 10
	Dx *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// The vertical offset of the image watermark relative to the output video. Default value: 0.
	//
	// Values can be one of the following:
	//
	// - Integer: The offset in pixels.
	//
	//   - Valid values: [8, 4096]
	//
	//   - Unit: px
	//
	// - Decimal: The ratio of the vertical offset to the output video height.
	//
	//   - Valid values: (0, 1)
	//
	//   - Up to four decimal places are supported, such as 0.9999. Extra digits are automatically discarded.
	//
	// example:
	//
	// 10
	Dy *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The watermark image file.
	File *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The height of the image watermark on the output video. Values can be one of the following:
	//
	// - Integer: The pixel height of the watermark image.
	//
	//   - Valid values: [8, 4096]
	//
	//   - Unit: px
	//
	// - Decimal: The ratio of the watermark height to the output video height.
	//
	//   - Valid values: (0, 1)
	//
	//   - Up to four decimal places are supported, such as 0.9999. Extra digits are automatically discarded.
	//
	// example:
	//
	// 32
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The position of the watermark.
	//
	// - Valid values: TopRight (top-right), TopLeft (top-left), BottomRight (bottom-right), and BottomLeft (bottom-left).
	//
	// - Default value: TopRight.
	//
	// example:
	//
	// TopLeft
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// The display time settings for a dynamic watermark.
	Timeline *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline `json:"Timeline,omitempty" xml:"Timeline,omitempty" type:"Struct"`
	// The width of the image watermark on the output video. Values can be one of the following:
	//
	// - Integer: The pixel width of the watermark image.
	//
	//   - Valid values: [8, 4096]
	//
	//   - Unit: px
	//
	// - Decimal: The ratio of the watermark width to the output video width.
	//
	//   - Valid values: (0, 1)
	//
	//   - Up to four decimal places are supported, such as 0.9999. Extra digits are automatically discarded.
	//
	// example:
	//
	// 32
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GetDx() *string {
	return s.Dx
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GetDy() *string {
	return s.Dy
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GetFile() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile {
	return s.File
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GetHeight() *string {
	return s.Height
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GetReferPos() *string {
	return s.ReferPos
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GetTimeline() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline {
	return s.Timeline
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) GetWidth() *string {
	return s.Width
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) SetDx(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	s.Dx = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) SetDy(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	s.Dy = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) SetFile(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	s.File = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) SetHeight(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	s.Height = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) SetReferPos(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	s.ReferPos = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) SetTimeline(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	s.Timeline = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) SetWidth(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams {
	s.Width = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParams) Validate() error {
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

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile struct {
	// The media value:
	//
	// - If Type is OSS, this is a URL that supports the OSS or HTTP protocol.
	//
	// - If Type is Media, this is a media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The media object type. Valid values:
	//
	// - OSS: An OSS file.
	//
	// - Media: A media asset ID.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) GetMedia() *string {
	return s.Media
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) GetType() *string {
	return s.Type
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) SetMedia(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile {
	s.Media = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) SetType(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile {
	s.Type = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline struct {
	// The duration of the watermark.
	//
	// - Valid values: [number, ToEND]
	//
	// - Default value: ToEND
	//
	// example:
	//
	// ToEND
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start time of the watermark.
	//
	// - Unit: seconds
	//
	// - Valid values: numbers
	//
	// - Default value: 0
	//
	// example:
	//
	// 00:00:05
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) GetDuration() *string {
	return s.Duration
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) GetStart() *string {
	return s.Start
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) SetDuration(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline {
	s.Duration = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) SetStart(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline {
	s.Start = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles struct {
	// Parameters to overwrite. If you specify these, they replace the corresponding parameters in the template.
	OverwriteParams *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) GetOverwriteParams() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams {
	return s.OverwriteParams
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) SetOverwriteParams(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles {
	s.OverwriteParams = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) SetTemplateId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles {
	s.TemplateId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams struct {
	// The encoding format of the file.
	//
	// example:
	//
	// UTF-8
	CharEnc *string `json:"CharEnc,omitempty" xml:"CharEnc,omitempty"`
	// The subtitle file.
	File *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The file format of the caption.
	//
	// example:
	//
	// vtt
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) GetCharEnc() *string {
	return s.CharEnc
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) GetFile() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile {
	return s.File
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) GetFormat() *string {
	return s.Format
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) SetCharEnc(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams {
	s.CharEnc = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) SetFile(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams {
	s.File = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) SetFormat(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams {
	s.Format = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParams) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile struct {
	// The value of Media:
	//
	// - If type is OSS, the value is a URL. The URL supports the OSS and HTTP protocols.
	//
	// - If type is Media, the value is the media asset ID.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the media object.
	//
	// Valid values:
	//
	// - OSS: An OSS file.
	//
	// - Media: The ID of a media asset.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) GetMedia() *string {
	return s.Media
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) GetType() *string {
	return s.Type
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) SetMedia(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile {
	s.Media = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) SetType(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile {
	s.Type = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks struct {
	// Parameters to overwrite. If you specify these, they replace the corresponding parameters in the template.
	OverwriteParams *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) GetOverwriteParams() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	return s.OverwriteParams
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) SetOverwriteParams(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks {
	s.OverwriteParams = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) SetTemplateId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks {
	s.TemplateId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarks) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams struct {
	// Adjusts the font size based on the output video size. The default is false.
	//
	// example:
	//
	// false
	Adaptive *string `json:"Adaptive,omitempty" xml:"Adaptive,omitempty"`
	// The outline color.
	//
	// Default: Black.
	//
	// For more values, see BorderColor.
	//
	// example:
	//
	// #006400
	BorderColor *string `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// The outline width.
	//
	// - Default: 0
	//
	// - Range: (0,4096]
	//
	// example:
	//
	// 0
	BorderWidth *int32 `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	// The watermark text. It does not need to be Base64 encoded. The string must be UTF-8 encoded.
	//
	// example:
	//
	// 测试水印
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The font transparency.
	//
	// - Valid values: (0, 1]
	//
	// - Default: 1.0
	//
	// example:
	//
	// 1.0
	FontAlpha *string `json:"FontAlpha,omitempty" xml:"FontAlpha,omitempty"`
	// The color.
	//
	// example:
	//
	// #006400
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// The font. Default: SimSun.
	//
	// example:
	//
	// SimSun
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// The font size.
	//
	// - Default value: 16
	//
	// - Valid values: (4, 120)
	//
	// example:
	//
	// 16
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// The left margin of the text.
	//
	// - Default: 0
	//
	// - Valid values: [0, 4096]
	//
	// example:
	//
	// 10
	Left *string `json:"Left,omitempty" xml:"Left,omitempty"`
	// The top margin of the text.
	//
	// - Default: 0
	//
	// - Valid values: [0, 4096]
	//
	// example:
	//
	// 10
	Top *string `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetAdaptive() *string {
	return s.Adaptive
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetBorderColor() *string {
	return s.BorderColor
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetBorderWidth() *int32 {
	return s.BorderWidth
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetContent() *string {
	return s.Content
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetFontAlpha() *string {
	return s.FontAlpha
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetFontColor() *string {
	return s.FontColor
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetFontName() *string {
	return s.FontName
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetFontSize() *int32 {
	return s.FontSize
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetLeft() *string {
	return s.Left
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) GetTop() *string {
	return s.Top
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetAdaptive(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.Adaptive = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetBorderColor(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.BorderColor = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetBorderWidth(v int32) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.BorderWidth = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetContent(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.Content = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetFontAlpha(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.FontAlpha = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetFontColor(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.FontColor = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetFontName(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.FontName = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetFontSize(v int32) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.FontSize = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetLeft(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.Left = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) SetTop(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams {
	s.Top = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTextWatermarksOverwriteParams) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode struct {
	// Parameters to overwrite. If you specify these, they replace the corresponding parameters in the template.
	OverwriteParams *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// 9547c6ad97cb4f2aaa29683ebd18d410
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) GetOverwriteParams() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams {
	return s.OverwriteParams
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) SetOverwriteParams(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode {
	s.OverwriteParams = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) SetTemplateId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode {
	s.TemplateId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscode) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams struct {
	// The audio settings.
	Audio *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// The container format settings.
	Container *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The multiplexing settings.
	MuxConfig *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	// The conditional transcoding parameters.
	TransConfig *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	// The video settings.
	Video *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) GetAudio() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	return s.Audio
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) GetContainer() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer {
	return s.Container
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) GetMuxConfig() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig {
	return s.MuxConfig
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) GetTransConfig() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	return s.TransConfig
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) GetVideo() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	return s.Video
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) SetAudio(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams {
	s.Audio = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) SetContainer(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams {
	s.Container = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) SetMuxConfig(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams {
	s.MuxConfig = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) SetTransConfig(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams {
	s.TransConfig = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) SetVideo(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams {
	s.Video = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParams) Validate() error {
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

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio struct {
	// The audio bitrate of the output file.
	//
	// - Valid values: [8, 1000]
	//
	// - Unit: Kbps
	//
	// - Default value: 128
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
	// The audio encoding profile. When Codec is AAC, valid values are aac_low, aac_he, aac_he_v2, aac_ld, and aac_eld.
	//
	// example:
	//
	// aac_low
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Specifies whether to delete the audio stream.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The sample rate.
	//
	// - Default value: 44100
	//
	// - Valid values: 22050, 32000, 44100, 48000, and 96000.
	//
	// - Unit: Hz
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	// The volume control.
	Volume *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GetChannels() *string {
	return s.Channels
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GetCodec() *string {
	return s.Codec
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GetProfile() *string {
	return s.Profile
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GetRemove() *string {
	return s.Remove
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GetSamplerate() *string {
	return s.Samplerate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) GetVolume() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume {
	return s.Volume
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) SetBitrate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	s.Bitrate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) SetChannels(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	s.Channels = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) SetCodec(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	s.Codec = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) SetProfile(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	s.Profile = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) SetRemove(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	s.Remove = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) SetSamplerate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	s.Samplerate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) SetVolume(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio {
	s.Volume = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudio) Validate() error {
	if s.Volume != nil {
		if err := s.Volume.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume struct {
	// The target loudness level.
	//
	// example:
	//
	// -6
	IntegratedLoudnessTarget *string `json:"IntegratedLoudnessTarget,omitempty" xml:"IntegratedLoudnessTarget,omitempty"`
	// The loudness range.
	//
	// example:
	//
	// 8
	LoudnessRangeTarget *string `json:"LoudnessRangeTarget,omitempty" xml:"LoudnessRangeTarget,omitempty"`
	// The volume adjustment method.
	//
	// example:
	//
	// auto
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The true peak level.
	//
	// example:
	//
	// -1
	TruePeak *string `json:"TruePeak,omitempty" xml:"TruePeak,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) GetIntegratedLoudnessTarget() *string {
	return s.IntegratedLoudnessTarget
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) GetLoudnessRangeTarget() *string {
	return s.LoudnessRangeTarget
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) GetMethod() *string {
	return s.Method
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) GetTruePeak() *string {
	return s.TruePeak
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) SetIntegratedLoudnessTarget(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.IntegratedLoudnessTarget = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) SetLoudnessRangeTarget(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.LoudnessRangeTarget = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) SetMethod(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.Method = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) SetTruePeak(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume {
	s.TruePeak = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsAudioVolume) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer struct {
	// The container format.
	//
	// example:
	//
	// mp4
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer) GetFormat() *string {
	return s.Format
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer) SetFormat(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer {
	s.Format = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig struct {
	// The segment settings.
	Segment *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig) GetSegment() *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	return s.Segment
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig) SetSegment(v *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig {
	s.Segment = v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig) Validate() error {
	if s.Segment != nil {
		if err := s.Segment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment struct {
	// The segment duration.
	//
	// example:
	//
	// 10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The forced segment time points.
	//
	// example:
	//
	// 2,3
	ForceSegTime *string `json:"ForceSegTime,omitempty" xml:"ForceSegTime,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GetDuration() *string {
	return s.Duration
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) GetForceSegTime() *string {
	return s.ForceSegTime
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) SetDuration(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	s.Duration = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) SetForceSegTime(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment {
	s.ForceSegTime = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfigSegment) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig struct {
	// The method used to adjust the display aspect ratio. This parameter takes effect only when both Width and Height are specified. You can use it with LongShortMode.
	//
	// Valid values: rescale, crop, pad, and none.
	//
	// Default value: none.
	//
	// example:
	//
	// none
	AdjDarMethod *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	// Specifies whether to check the audio bitrate. IsCheckAudioBitrate and IsCheckAudioBitrateFail are mutually exclusive. IsCheckAudioBitrateFail has higher priority.
	//
	// - true: Check the bitrate. If the input audio bitrate is lower than the output setting, transcode at the input bitrate.
	//
	// - false: Do not check the bitrate.
	//
	// Default value rules:
	//
	// - Empty and the codec differs from the input source: false.
	//
	// - Empty and the codec matches the input source: true.
	//
	// example:
	//
	// true
	IsCheckAudioBitrate *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	// Specifies whether to check the audio bitrate. IsCheckAudioBitrate and IsCheckAudioBitrateFail are mutually exclusive. This parameter has higher priority.
	//
	// - true: Check the bitrate. If the input audio bitrate is lower than the output setting, return a failure.
	//
	// - false (default): Do not check the bitrate.
	//
	// example:
	//
	// true
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	// Specifies whether to check the video resolution. IsCheckReso and IsCheckResoFail are mutually exclusive. IsCheckResoFail has higher priority.
	//
	// - true: Check the resolution. If the input video resolution (width or height) is smaller than the output setting, transcode at the input resolution.
	//
	// - false (default): Do not check the resolution.
	//
	// example:
	//
	// true
	IsCheckReso *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	// Specifies whether to check the video resolution. IsCheckReso and IsCheckResoFail are mutually exclusive. This parameter has higher priority.
	//
	// - true: Check the resolution. If the input video resolution (width or height) is smaller than the output setting, return a failure.
	//
	// - false (default): Do not check the resolution.
	//
	// example:
	//
	// true
	IsCheckResoFail *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	// Specifies whether to check the video bitrate. IsCheckVideoBitrate and IsCheckVideoBitrateFail are mutually exclusive. IsCheckVideoBitrateFail has higher priority.
	//
	// - true: Check the bitrate. If the input video bitrate is lower than the output setting, transcode at the input bitrate.
	//
	// - false (default): Do not check the bitrate.
	//
	// example:
	//
	// true
	IsCheckVideoBitrate *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	// Specifies whether to check the video bitrate. IsCheckVideoBitrate and IsCheckVideoBitrateFail are mutually exclusive. This parameter has higher priority.
	//
	// - true: Check the bitrate. If the input video bitrate is lower than the output setting, return a failure.
	//
	// - false (default): Do not check the bitrate.
	//
	// example:
	//
	// true
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	// The video transcoding mode. Valid values:
	//
	// - onepass (default): Used for ABR. Encoding is faster than twopass.
	//
	// - twopass: Used for VBR. Encoding is slower than onepass.
	//
	// - CBR: Constant bitrate mode.
	//
	// example:
	//
	// onepass
	TransMode *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetAdjDarMethod() *string {
	return s.AdjDarMethod
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckAudioBitrate() *string {
	return s.IsCheckAudioBitrate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckAudioBitrateFail() *string {
	return s.IsCheckAudioBitrateFail
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckReso() *string {
	return s.IsCheckReso
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckResoFail() *string {
	return s.IsCheckResoFail
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckVideoBitrate() *string {
	return s.IsCheckVideoBitrate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetIsCheckVideoBitrateFail() *string {
	return s.IsCheckVideoBitrateFail
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetAdjDarMethod(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.AdjDarMethod = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckAudioBitrate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckAudioBitrateFail(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckReso(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckResoFail(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckVideoBitrate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetIsCheckVideoBitrateFail(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) SetTransMode(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig {
	s.TransMode = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsTransConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo struct {
	// The maximum bitrate for adaptive bitrate streaming (ABR). This applies only to narrow-high 1. Valid values: [10, 50000]. Unit: Kbps.
	//
	// example:
	//
	// 6000
	AbrMax *string `json:"AbrMax,omitempty" xml:"AbrMax,omitempty"`
	// The average video bitrate.
	//
	// - Valid values: [10, 50000].
	//
	// - Unit: Kbps.
	//
	// example:
	//
	// 3000
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The buffer size.
	//
	// - Valid values: [1000, 128000]
	//
	// - Default value: 6000
	//
	// - Unit: Kb
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
	// The constant rate factor (CRF), which controls the trade-off between quality and bitrate.
	//
	// - Valid values: [0, 51].
	//
	// - Default values: 23 for H.264 and 26 for H.265.
	//
	// If you set Crf, the Bitrate setting is ignored.
	//
	// example:
	//
	// 23
	Crf *string `json:"Crf,omitempty" xml:"Crf,omitempty"`
	// The video cropping method. Two options are available.
	//
	// - Automatically detect and crop black bars. Set this to border.
	//
	// - Custom cropping. Format: width:height:left:top. Example: 1280:800:0:140
	//
	// example:
	//
	// 1280:800:0:140
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// The frame rate.
	//
	// - Valid values: (0, 60].
	//
	// - If the frame rate of the input file exceeds 60, the system uses 60.
	//
	// - Default value: The frame rate of the input file.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The maximum number of frames between keyframes.
	//
	// - Valid values: [1, 1080000].
	//
	// - Default value: 250.
	//
	// example:
	//
	// 250
	Gop *string `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height.
	//
	// - Valid values: [128, 4096].
	//
	// - Unit: px.
	//
	// - Default value: The original video height.
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// Specifies whether to enable automatic rotation for portrait or landscape videos (also known as long-side/short-side mode).
	//
	// example:
	//
	// false
	LongShortMode *string `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
	// The peak video bitrate. Valid values: [10, 50000]. Unit: Kbps.
	//
	// example:
	//
	// 9000
	Maxrate *string `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	// The padding configuration for black bars.
	//
	// - Format: width:height:left:top.
	//
	// - Example: 1280:800:0:140
	//
	// example:
	//
	// 1280:800:0:140
	Pad *string `json:"Pad,omitempty" xml:"Pad,omitempty"`
	// The video color format. Valid values include yuv420p and yuvj420p.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The video encoder preset. Only H.264 supports this parameter. Valid values: veryfast, fast, medium, slow, and slower. Default value: medium.
	//
	// example:
	//
	// medium
	Preset *string `json:"Preset,omitempty" xml:"Preset,omitempty"`
	// The encoding profile. Valid values: baseline, main, and high.
	//
	// - baseline: For mobile devices.
	//
	// - main: For standard-resolution devices.
	//
	// - high: For high-resolution devices.
	//
	// Default value: high.
	//
	// example:
	//
	// Main
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// Specifies whether to remove the video.
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
	// The width.
	//
	// - Valid values: [128, 4096].
	//
	// - Unit: px.
	//
	// - Default value: The original video width.
	//
	// example:
	//
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetAbrMax() *string {
	return s.AbrMax
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetBufsize() *string {
	return s.Bufsize
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetCodec() *string {
	return s.Codec
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetCrf() *string {
	return s.Crf
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetCrop() *string {
	return s.Crop
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetFps() *string {
	return s.Fps
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetGop() *string {
	return s.Gop
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetHeight() *string {
	return s.Height
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetLongShortMode() *string {
	return s.LongShortMode
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetMaxrate() *string {
	return s.Maxrate
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetPad() *string {
	return s.Pad
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetPixFmt() *string {
	return s.PixFmt
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetPreset() *string {
	return s.Preset
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetProfile() *string {
	return s.Profile
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetRemove() *string {
	return s.Remove
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) GetWidth() *string {
	return s.Width
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetAbrMax(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.AbrMax = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetBitrate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Bitrate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetBufsize(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Bufsize = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetCodec(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Codec = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetCrf(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Crf = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetCrop(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Crop = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetFps(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Fps = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetGop(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Gop = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetHeight(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Height = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetLongShortMode(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.LongShortMode = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetMaxrate(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Maxrate = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetPad(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Pad = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetPixFmt(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.PixFmt = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetPreset(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Preset = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetProfile(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Profile = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetRemove(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Remove = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetScanMode(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.ScanMode = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) SetWidth(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo {
	s.Width = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsVideo) Validate() error {
	return dara.Validate(s)
}

type SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig struct {
	// The ID of the pipeline.
	//
	// example:
	//
	// e37ebee5d98b4781897f6086e89f9c56
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the task. A larger value indicates a higher priority. The value can be an integer from 1 to 10.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) SetPipelineId(v string) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) SetPriority(v int32) *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig {
	s.Priority = &v
	return s
}

func (s *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig) Validate() error {
	return dara.Validate(s)
}
