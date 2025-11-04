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
	InputGroup []*SubmitTranscodeJobResponseBodyTranscodeParentJobInputGroup `json:"InputGroup,omitempty" xml:"InputGroup,omitempty" type:"Repeated"`
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
	OutputGroup []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroup `json:"OutputGroup,omitempty" xml:"OutputGroup,omitempty" type:"Repeated"`
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
	ScheduleConfig *SubmitTranscodeJobResponseBodyTranscodeParentJobScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The state of the job. Success: At least one of the subjobs is successful. Fail: All subjobs failed.
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
	// The list of subjobs.
	TranscodeJobList []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobList `json:"TranscodeJobList,omitempty" xml:"TranscodeJobList,omitempty" type:"Repeated"`
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
	// The media object.
	//
	// 	- If Type is set to OSS, set this parameter to the URL of an OSS object. Both the OSS and HTTP protocols are supported.
	//
	// 	- If Type is set to Media, set this parameter to the ID of a media asset.
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
	// The output file configuration.
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
	// The media object.
	//
	// 	- If Type is set to OSS, set this parameter to the URL of an OSS object. Both the OSS and HTTP protocols are supported.
	//
	// 	- If Type is set to Media, set this parameter to the ID of a media asset.
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
	// The multi-input stream merge configuration.
	CombineConfigs []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigCombineConfigs `json:"CombineConfigs,omitempty" xml:"CombineConfigs,omitempty" type:"Repeated"`
	// The encryption settings.
	Encryption *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigEncryption `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	// The watermark configuration of an image.
	ImageWatermarks []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarks `json:"ImageWatermarks,omitempty" xml:"ImageWatermarks,omitempty" type:"Repeated"`
	// The subtitle configuration.
	Subtitles []*SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigSubtitles `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// The configurations of the text watermark.
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
	// The ciphertext of HLS encryption.
	//
	// example:
	//
	// MTYi00NDU0LTg5O****
	CipherText *string `json:"CipherText,omitempty" xml:"CipherText,omitempty"`
	// The address of the decryption service for HLS encryption.
	//
	// example:
	//
	// https://sample.com/path?CipherText=MTYi00NDU0LTg5O****
	DecryptKeyUri *string `json:"DecryptKeyUri,omitempty" xml:"DecryptKeyUri,omitempty"`
	// Specifies the encryption type.
	//
	// example:
	//
	// PrivateEncryption
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The type of the key service. Valid values: KMS and Base64.
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
	// The parameters that are used to overwrite the corresponding parameters of the template.
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
	// The horizontal offset of the watermark relative to the output video. Default value: 0.
	//
	// The following value types are supported:
	//
	// 	- Integer: the pixel value of the horizontal offset.
	//
	//     	- Valid values: [8,4096].
	//
	//     	- Unit: pixels.
	//
	// 	- Decimal: the ratio of the horizontal offset to the width of the output video.
	//
	//     	- Valid values: (0,1).
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
	//
	// example:
	//
	// 10
	Dx *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// The vertical offset of the watermark relative to the output video. Default value: 0.
	//
	// The following value types are supported:
	//
	// 	- Integer: the pixel value of the horizontal offset.
	//
	//     	- Valid values: [8,4096].
	//
	//     	- Unit: pixels.
	//
	// 	- Decimal: the ratio of the vertical offset to the height of the output video.
	//
	//     	- Valid values: (0,1).
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
	//
	// example:
	//
	// 10
	Dy *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The watermark image file.
	File *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The height of the watermark image in the output video. The following value types are supported:
	//
	// 	- Integer: the pixel value of the watermark height.
	//
	//     	- Valid values: [8,4096].
	//
	//     	- Unit: pixels.
	//
	// 	- Decimal: the ratio of the watermark height to the height of the output video.
	//
	//     	- Valid values: (0,1).
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
	//
	// example:
	//
	// 32
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The position of the watermark.
	//
	// 	- Valid values: TopRight, TopLeft, BottomRight, and BottomLeft.
	//
	// 	- Default value: TopRight.
	//
	// example:
	//
	// TopLeft
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// The time settings of the dynamic watermark.
	Timeline *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigImageWatermarksOverwriteParamsTimeline `json:"Timeline,omitempty" xml:"Timeline,omitempty" type:"Struct"`
	// The width of the watermark in the output video. The following value types are supported:
	//
	// 	- Integer: the pixel value of the watermark width.
	//
	//     	- Valid values: [8,4096].
	//
	//     	- Unit: pixels.
	//
	// 	- Decimal: the ratio of the watermark width to the width of the output video.
	//
	//     	- Valid values: (0,1).
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
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
	// The media object.
	//
	// 	- If Type is set to OSS, set this parameter to the URL of an OSS object. Both the OSS and HTTP protocols are supported.
	//
	// 	- If Type is set to Media, set this parameter to the ID of a media asset.
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
	// The time range in which the watermark is displayed.
	//
	// 	- Valid values: integers and ToEND.
	//
	// 	- Default value: ToEND.
	//
	// example:
	//
	// ToEND
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The beginning of the time range in which the watermark is displayed.
	//
	// 	- Unit: seconds.
	//
	// 	- Value values: integers.
	//
	// 	- Default value: 0.
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
	// The parameters that are used to overwrite the corresponding parameters of the template.
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
	// The format of the subtitle file.
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
	// The parameters that are used to overwrite the corresponding parameters of the template.
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
	// Specifies whether to the font size based on the output video dimensions.
	//
	// 	- true: false
	//
	// 	- default: false
	//
	// example:
	//
	// false
	Adaptive *string `json:"Adaptive,omitempty" xml:"Adaptive,omitempty"`
	// The outline color of the text watermark. Default value: black. For more information, see BorderColor.
	//
	// example:
	//
	// #006400
	BorderColor *string `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// The outline width of the text watermark.
	//
	// 	- Default value: 0.
	//
	// 	- Valid values: (0,4096].
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
	// The transparency of the text.
	//
	// 	- Valid values: (0,1].
	//
	// 	- Default value: 1.
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
	// The font of the text. Default value: SimSun.
	//
	// example:
	//
	// SimSun
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// The size of the text.
	//
	// 	- Default value: 16.
	//
	// 	- Valid values: (4,120).
	//
	// example:
	//
	// 16
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// The left margin of the text watermark.
	//
	// 	- Default value: 0.
	//
	// 	- Valid values: [0,4096].
	//
	// example:
	//
	// 10
	Left *string `json:"Left,omitempty" xml:"Left,omitempty"`
	// The top margin of the text.
	//
	// 	- Default value: 0.
	//
	// 	- Valid values: [0,4096].
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
	// The parameters that are used to overwrite the corresponding parameters of the template.
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
	// The audio settings.
	Audio *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	// The encapsulation format settings.
	Container *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The encapsulation settings.
	MuxConfig *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	// The conditional transcoding configurations.
	TransConfig *SubmitTranscodeJobResponseBodyTranscodeParentJobOutputGroupProcessConfigTranscodeOverwriteParamsTransConfig `json:"TransConfig,omitempty" xml:"TransConfig,omitempty" type:"Struct"`
	// The video settings.
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
	// Specifies whether to delete the audio stream.
	//
	// example:
	//
	// false
	Remove *string `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The sampling rate.
	//
	// 	- Valid values: 22050, 32000, 44100, 48000, and 96000. Default value: 44100.
	//
	// 	- Unit: Hz.
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	// The volume configurations.
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
	// The segment settings.
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
	// The forced segmentation point in time.
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
	// The method that is used to adjust the resolution. This parameter takes effect only if both the Width and Height parameters are specified. You can use this parameter together with the LongShortMode parameter.
	//
	// Valid values: rescale, crop, pad, and none.
	//
	// Default value: none.
	//
	// example:
	//
	// none
	AdjDarMethod *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	// Specifies whether to check the audio bitrate. You can specify only one of the IsCheckAudioBitrate and IsCheckAudioBitrateFail parameters. The priority of the IsCheckAudioBitrateFail parameter is higher. Valid values:
	//
	// 	- true: checks the video resolution. If the bitrate of the input audio is less than that of the output audio, the bitrate of the input audio is used for transcoding.
	//
	// 	- false: does not check the video resolution.
	//
	// Default value:
	//
	// 	- If this parameter is not specified and the codec of the output audio is different from that of the input audio, the default value is false.
	//
	// 	- If this parameter is not specified and the codec of the output audio is the same as that of the input audio, the default value is true.
	//
	// example:
	//
	// true
	IsCheckAudioBitrate *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	// Specifies whether to check the audio bitrate. You can specify only one of the IsCheckAudioBitrate and IsCheckAudioBitrateFail parameters. The priority of the IsCheckAudioBitrateFail parameter is higher. Valid values:
	//
	// 	- true: checks the video resolution. If the bitrate of the input audio is less than that of the output audio, the transcoding job fails.
	//
	// 	- false: does not check the video resolution.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	// Specifies whether to check the video resolution. You can specify only one of the IsCheckReso and IsCheckResoFail parameters. The priority of the IsCheckResoFail parameter is higher. Valid values:
	//
	// 	- true: checks the video resolution. If the width or height of the input video is less than that of the output video, the resolution of the input video is used for transcoding.
	//
	// 	- false: does not check the video resolution.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckReso *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	// Specifies whether to check the video resolution. You can specify only one of the IsCheckReso and IsCheckResoFail parameters. The priority of the IsCheckResoFail parameter is higher. Valid values:
	//
	// 	- true: checks the video resolution. If the width or height of the input video is less than that of the output video, the transcoding job fails.
	//
	// 	- false: does not check the video resolution.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckResoFail *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	// Specifies whether to check the video bitrate. You can specify only one of the IsCheckVideoBitrate and IsCheckVideoBitrateFail parameters. The priority of the IsCheckVideoBitrateFail parameter is higher. Valid values:
	//
	// 	- true: checks the video resolution. If the bitrate of the input video is less than that of the output video, the bitrate of the input video is used for transcoding.
	//
	// 	- false: does not check the video resolution.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckVideoBitrate *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	// Specifies whether to check the video bitrate. You can specify only one of the IsCheckVideoBitrate and IsCheckVideoBitrateFail parameters. The priority of the IsCheckVideoBitrateFail parameter is higher. Valid values:
	//
	// 	- true: checks the video resolution. If the bitrate of the input video is less than that of the output video, the transcoding job fails.
	//
	// 	- false: does not check the video resolution.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	// The video transcoding mode. Valid values:
	//
	// 	- onepass: You can set this parameter to onepass if the Bitrate parameter is set to ABR. The encoding speed of this mode is faster than that of the twopass mode.
	//
	// 	- twopass: You can set this parameter to twopass if the Bitrate parameter is set to VBR. The encoding speed of this mode is slower than that of the onepass mode.
	//
	// 	- CBR: the constant bitrate mode.
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
	// The maximum ABR. This parameter takes effect only for Narrowband HD 1.0.
	//
	// 	- Valid values: [10,50000].
	//
	// 	- Unit: Kbit/s.
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
	// 	- Default value: 23 if the encoding format is H.264, or Default value when the Codec parameter is set to H.265: 26.
	//
	// If this parameter is specified, the value of Bitrate becomes invalid.
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
	// Specifies whether to enable the auto-rotate screen feature.
	//
	// example:
	//
	// false
	LongShortMode *string `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
	// The maximum bitrate of the output video.
	//
	// 	- Valid values: [10,50000].
	//
	// 	- Unit: Kbit/s.
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
	// The time when the job was created.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the job was complete.
	//
	// example:
	//
	// 2022-01-12T08:49:41Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The input group of the job. An input of a single file indicates a transcoding job. An input of multiple files indicates an audio and video stream merge job.
	InputGroup []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListInputGroup `json:"InputGroup,omitempty" xml:"InputGroup,omitempty" type:"Repeated"`
	// The subjob ID.
	//
	// example:
	//
	// 7d6a7e0d4db2457a8d45ff5d43e1bf0a
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The index number of the subjob in the entire job.
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
	// The media information about the video generated by the job.
	OutFileMeta *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMeta `json:"OutFileMeta,omitempty" xml:"OutFileMeta,omitempty" type:"Struct"`
	// The output file configuration.
	Output *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The main job ID.
	//
	// example:
	//
	// 8b2198504dd340b7b3c9842a74fc9baa
	ParentJobId *string `json:"ParentJobId,omitempty" xml:"ParentJobId,omitempty"`
	// The transcoding configuration.
	ProcessConfig *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfig `json:"ProcessConfig,omitempty" xml:"ProcessConfig,omitempty" type:"Struct"`
	// The ID of the request that submitted the job.
	//
	// example:
	//
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scheduling information about the job.
	ScheduleConfig *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The state of the transcoding job. Valid values:
	//
	// 	- Init: The job is submitted.
	//
	// 	- Processing: The job is in progress.
	//
	// 	- Success: The job is successful.
	//
	// 	- Fail: The job failed.
	//
	// 	- Deleted: The job is deleted.
	//
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The job submission result.
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
	// The user data.
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
	// The URL of the input stream:
	//
	// 	- This parameter takes effect only when Type is set to Media. You can select a specific file within the media asset as an input.
	//
	// 	- The system checks whether the input URL exists within the media asset.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	InputUrl *string `json:"InputUrl,omitempty" xml:"InputUrl,omitempty"`
	// The media object.
	//
	// 	- If Type is set to OSS, set this parameter to the URL of an OSS object. Both the OSS and HTTP protocols are supported.
	//
	// 	- If Type is set to Media, set this parameter to the ID of a media asset.
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
	// The information about the audio stream.
	AudioStreamInfoList []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaAudioStreamInfoList `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	// The basic file information.
	FileBasicInfo *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListOutFileMetaFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	// The information about the video stream.
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
	// cn
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
	// The state of the file.
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
	// 73e07de0f77171eca3fc7035d0b26402
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
	// The duration of the stream. Unit: seconds.
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
	// cn
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
	// The rotation angle of the video image. Valid values: 0, 90, 180, and 270. Default value: 0.
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
	// The media object.
	//
	// 	- If Type is set to OSS, set this parameter to the URL of an OSS object. Both the OSS and HTTP protocols are supported.
	//
	// 	- If Type is set to Media, set this parameter to the ID of a media asset.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The URL of the output stream.\\
	//
	// This parameter takes effect only when Type is set to Media. You can select a specific file within the media asset as an output.\\
	//
	// Supported placeholders:
	//
	// 	- {MediaId}: the ID of the media asset.
	//
	// 	- {JobId}: the ID of the transcoding subjob.
	//
	// 	- {MediaBucket}: the bucket to which the media asset belongs.
	//
	// 	- {ExtName}: the file suffix, which uses the output format of the transcoding template.
	//
	// 	- {DestMd5}: the MD5 value of the transcoded output file.\\
	//
	//     Notes:
	//
	// 1.  This parameter must contain the {MediaId} and {JobId} placeholders.
	//
	// 2.  The output bucket is the same as the bucket to which the media asset belongs.
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
	// The multi-input stream merge configuration.
	CombineConfigs []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigCombineConfigs `json:"CombineConfigs,omitempty" xml:"CombineConfigs,omitempty" type:"Repeated"`
	// The encryption settings.
	Encryption *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigEncryption `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	// The watermark configuration of an image.
	ImageWatermarks []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarks `json:"ImageWatermarks,omitempty" xml:"ImageWatermarks,omitempty" type:"Repeated"`
	// The subtitle configuration.
	Subtitles []*SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitles `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// The configurations of the text watermark.
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
	// The ciphertext of HLS encryption.
	//
	// example:
	//
	// MTYi00NDU0LTg5O****
	CipherText *string `json:"CipherText,omitempty" xml:"CipherText,omitempty"`
	// The address of the decryption service for HLS encryption.
	//
	// example:
	//
	// https://sample.com/path?CipherText=MTYi00NDU0LTg5O****
	DecryptKeyUri *string `json:"DecryptKeyUri,omitempty" xml:"DecryptKeyUri,omitempty"`
	// Specifies the encryption type.
	//
	// example:
	//
	// PrivateEncryption
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The type of the key service. Valid values: KMS and Base64.
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
	// The parameters that are used to overwrite the corresponding parameters of the template.
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
	// The horizontal offset of the watermark relative to the output video. Default value: 0.
	//
	// The following value types are supported:
	//
	// 	- Integer: the pixel value of the horizontal offset.
	//
	//     	- Valid values: [8,4096].
	//
	//     	- Unit: pixels.
	//
	// 	- Decimal: the ratio of the horizontal offset to the width of the output video.
	//
	//     	- Valid values: (0,1).
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
	//
	// example:
	//
	// 10
	Dx *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// The vertical offset of the watermark relative to the output video. Default value: 0.
	//
	// The following value types are supported:
	//
	// 	- Integer: the pixel value of the horizontal offset.
	//
	//     	- Valid values: [8,4096].
	//
	//     	- Unit: pixels.
	//
	// 	- Decimal: the ratio of the vertical offset to the height of the output video.
	//
	//     	- Valid values: (0,1).
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
	//
	// example:
	//
	// 10
	Dy *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The watermark image file.
	File *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The height of the watermark image in the output video. The following value types are supported:
	//
	// 	- Integer: the pixel value of the watermark height.
	//
	//     	- Valid values: [8,4096].
	//
	//     	- Unit: pixels.
	//
	// 	- Decimal: the ratio of the watermark height to the height of the output video.
	//
	//     	- Valid values: (0,1).
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
	//
	// example:
	//
	// 32
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The position of the watermark.
	//
	// 	- Valid values: TopRight, TopLeft, BottomRight, and BottomLeft.
	//
	// 	- Default value: TopRight.
	//
	// example:
	//
	// TopLeft
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// The time settings of the dynamic watermark.
	Timeline *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigImageWatermarksOverwriteParamsTimeline `json:"Timeline,omitempty" xml:"Timeline,omitempty" type:"Struct"`
	// The width of the watermark in the output video. The following value types are supported:
	//
	// 	- Integer: the pixel value of the watermark width.
	//
	//     	- Valid values: [8,4096].
	//
	//     	- Unit: pixels.
	//
	// 	- Decimal: the ratio of the watermark width to the width of the output video.
	//
	//     	- Valid values: (0,1).
	//
	//     	- The decimal number can be accurate to four decimal places, such as 0.9999. Excessive digits are automatically discarded.
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
	// The media object.
	//
	// 	- If Type is set to OSS, set this parameter to the URL of an OSS object. Both the OSS and HTTP protocols are supported.
	//
	// 	- If Type is set to Media, set this parameter to the ID of a media asset.
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
	// The time range in which the watermark is displayed.
	//
	// 	- Valid values: integers and ToEND.
	//
	// 	- Default value: ToEND.
	//
	// example:
	//
	// ToEND
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The beginning of the time range in which the watermark is displayed.
	//
	// 	- Unit: seconds.
	//
	// 	- Value values: integers.
	//
	// 	- Default value: 0.
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
	// The parameters that are used to overwrite the corresponding parameters of the template.
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
	// The file encoding format.
	//
	// example:
	//
	// UTF-8
	CharEnc *string `json:"CharEnc,omitempty" xml:"CharEnc,omitempty"`
	// The subtitle file.
	File *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigSubtitlesOverwriteParamsFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The format of the subtitle file.
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
	// The media object.
	//
	// 	- If Type is set to OSS, set this parameter to the URL of an OSS object. Both the OSS and HTTP protocols are supported.
	//
	// 	- If Type is set to Media, set this parameter to the ID of a media asset.
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
	// The parameters that are used to overwrite the corresponding parameters of the template.
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
	// Specifies whether to the font size based on the output video dimensions. true / false, default: false
	//
	// example:
	//
	// false
	Adaptive *string `json:"Adaptive,omitempty" xml:"Adaptive,omitempty"`
	// The outline color of the text watermark. Default value: black. For more information, see BorderColor.
	//
	// example:
	//
	// #006400
	BorderColor *string `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// The outline width of the text watermark.
	//
	// 	- Default value: 0.
	//
	// 	- Valid values: (0,4096].
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
	// The transparency of the text.
	//
	// 	- Valid values: (0,1].
	//
	// 	- Default value: 1.
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
	// The font of the text. Default value: SimSun.
	//
	// example:
	//
	// SimSun
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// The size of the text.
	//
	// 	- Default value: 16.
	//
	// 	- Valid values: (4,120).
	//
	// example:
	//
	// 16
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// The left margin of the text watermark.
	//
	// 	- Default value: 0.
	//
	// 	- Valid values: [0,4096].
	//
	// example:
	//
	// 10
	Left *string `json:"Left,omitempty" xml:"Left,omitempty"`
	// The top margin of the text.
	//
	// 	- Default value: 0.
	//
	// 	- Valid values: [0,4096].
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
	// The parameters that are used to overwrite the corresponding parameters of the template.
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
	// The encapsulation format settings.
	Container *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The encapsulation settings.
	MuxConfig *SubmitTranscodeJobResponseBodyTranscodeParentJobTranscodeJobListProcessConfigTranscodeOverwriteParamsMuxConfig `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty" type:"Struct"`
	// The conditional transcoding configurations.
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
	// Specifies whether to delete the audio stream.
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
	// The method that is used to adjust the resolution. This parameter takes effect only if both the Width and Height parameters are specified. You can use this parameter together with the LongShortMode parameter.
	//
	// Valid values: rescale, crop, pad, and none.
	//
	// Default value: none.
	//
	// example:
	//
	// none
	AdjDarMethod *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	// Specifies whether to check the audio bitrate. You can specify only one of the IsCheckAudioBitrate and IsCheckAudioBitrateFail parameters. The priority of the IsCheckAudioBitrateFail parameter is higher. Valid values:
	//
	// 	- true: checks the video resolution. If the bitrate of the input audio is less than that of the output audio, the bitrate of the input audio is used for transcoding.
	//
	// 	- false: does not check the video resolution.
	//
	// Default values:
	//
	// 	- If this parameter is not specified and the codec of the output audio is different from that of the input audio, the default value is false.
	//
	// 	- If this parameter is not specified and the codec of the output audio is the same as that of the input audio, the default value is true.
	//
	// example:
	//
	// true
	IsCheckAudioBitrate *string `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	// Specifies whether to check the audio bitrate. You can specify only one of the IsCheckAudioBitrate and IsCheckAudioBitrateFail parameters. The priority of the IsCheckAudioBitrateFail parameter is higher. Valid values:
	//
	// 	- true: checks the video resolution. If the bitrate of the input audio is less than that of the output audio, the transcoding job fails.
	//
	// 	- false: does not check the video resolution. This is the default value.
	//
	// example:
	//
	// true
	IsCheckAudioBitrateFail *string `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	// Specifies whether to check the video resolution. You can specify only one of the IsCheckReso and IsCheckResoFail parameters. The priority of the IsCheckResoFail parameter is higher. Valid values:
	//
	// 	- true: checks the video resolution. If the width or height of the input video is less than that of the output video, the resolution of the input video is used for transcoding.
	//
	// 	- false: does not check the video resolution. This is the default value.
	//
	// example:
	//
	// true
	IsCheckReso *string `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	// Specifies whether to check the video resolution. You can specify only one of the IsCheckReso and IsCheckResoFail parameters. The priority of the IsCheckResoFail parameter is higher. Valid values:
	//
	// 	- true: checks the video resolution. If the width or height of the input video is less than that of the output video, the transcoding job fails.
	//
	// 	- false: does not check the video resolution. This is the default value.
	//
	// example:
	//
	// true
	IsCheckResoFail *string `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	// Specifies whether to check the video bitrate. You can specify only one of the IsCheckVideoBitrate and IsCheckVideoBitrateFail parameters. The priority of the IsCheckVideoBitrateFail parameter is higher. Valid values:
	//
	// 	- true: checks the video resolution. If the bitrate of the input video is less than that of the output video, the bitrate of the input video is used for transcoding.
	//
	// 	- false: does not check the video resolution. This is the default value.
	//
	// example:
	//
	// true
	IsCheckVideoBitrate *string `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	// Specifies whether to check the video bitrate. You can specify only one of the IsCheckVideoBitrate and IsCheckVideoBitrateFail parameters. The priority of the IsCheckVideoBitrateFail parameter is higher. Valid values:
	//
	// 	- true: checks the video resolution. If the bitrate of the input video is less than that of the output video, the transcoding job fails.
	//
	// 	- false: does not check the video resolution. This is the default value.
	//
	// example:
	//
	// true
	IsCheckVideoBitrateFail *string `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	// The video transcoding mode. Valid values:
	//
	// 	- onepass: You can set this parameter to onepass if the Bitrate parameter is set to ABR. This is the default value. The encoding speed of this mode is faster than that of the twopass mode.
	//
	// 	- twopass: You can set this parameter to twopass if the Bitrate parameter is set to VBR. The encoding speed of this mode is slower than that of the onepass mode.
	//
	// 	- CBR: the constant bitrate mode.
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
	// The maximum ABR. This parameter takes effect only for Narrowband HD 1.0. Valid values: [10,50000]. Unit: Kbit/s.
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
	// 	- Default value: 23 if the encoding format is H.264, or Default value when the Codec parameter is set to H.265: 26.
	//
	// If this parameter is specified, the value of Bitrate becomes invalid.
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
	// Specifies whether to enable the auto-rotate screen feature.
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
