// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDynamicImageJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *SubmitDynamicImageJobRequestInput) *SubmitDynamicImageJobRequest
	GetInput() *SubmitDynamicImageJobRequestInput
	SetName(v string) *SubmitDynamicImageJobRequest
	GetName() *string
	SetOutput(v *SubmitDynamicImageJobRequestOutput) *SubmitDynamicImageJobRequest
	GetOutput() *SubmitDynamicImageJobRequestOutput
	SetScheduleConfig(v *SubmitDynamicImageJobRequestScheduleConfig) *SubmitDynamicImageJobRequest
	GetScheduleConfig() *SubmitDynamicImageJobRequestScheduleConfig
	SetTemplateConfig(v *SubmitDynamicImageJobRequestTemplateConfig) *SubmitDynamicImageJobRequest
	GetTemplateConfig() *SubmitDynamicImageJobRequestTemplateConfig
	SetUserData(v string) *SubmitDynamicImageJobRequest
	GetUserData() *string
}

type SubmitDynamicImageJobRequest struct {
	// The job input.
	//
	// This parameter is required.
	Input *SubmitDynamicImageJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The job name.
	//
	// example:
	//
	// SampleJob
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The job output.
	//
	// This parameter is required.
	Output *SubmitDynamicImageJobRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The scheduling configuration.
	ScheduleConfig *SubmitDynamicImageJobRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The snapshot template configuration.
	//
	// This parameter is required.
	TemplateConfig *SubmitDynamicImageJobRequestTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty" type:"Struct"`
	// The user-defined data.
	//
	// example:
	//
	// {"SampleKey": "SampleValue"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitDynamicImageJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicImageJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobRequest) GetInput() *SubmitDynamicImageJobRequestInput {
	return s.Input
}

func (s *SubmitDynamicImageJobRequest) GetName() *string {
	return s.Name
}

func (s *SubmitDynamicImageJobRequest) GetOutput() *SubmitDynamicImageJobRequestOutput {
	return s.Output
}

func (s *SubmitDynamicImageJobRequest) GetScheduleConfig() *SubmitDynamicImageJobRequestScheduleConfig {
	return s.ScheduleConfig
}

func (s *SubmitDynamicImageJobRequest) GetTemplateConfig() *SubmitDynamicImageJobRequestTemplateConfig {
	return s.TemplateConfig
}

func (s *SubmitDynamicImageJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitDynamicImageJobRequest) SetInput(v *SubmitDynamicImageJobRequestInput) *SubmitDynamicImageJobRequest {
	s.Input = v
	return s
}

func (s *SubmitDynamicImageJobRequest) SetName(v string) *SubmitDynamicImageJobRequest {
	s.Name = &v
	return s
}

func (s *SubmitDynamicImageJobRequest) SetOutput(v *SubmitDynamicImageJobRequestOutput) *SubmitDynamicImageJobRequest {
	s.Output = v
	return s
}

func (s *SubmitDynamicImageJobRequest) SetScheduleConfig(v *SubmitDynamicImageJobRequestScheduleConfig) *SubmitDynamicImageJobRequest {
	s.ScheduleConfig = v
	return s
}

func (s *SubmitDynamicImageJobRequest) SetTemplateConfig(v *SubmitDynamicImageJobRequestTemplateConfig) *SubmitDynamicImageJobRequest {
	s.TemplateConfig = v
	return s
}

func (s *SubmitDynamicImageJobRequest) SetUserData(v string) *SubmitDynamicImageJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitDynamicImageJobRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleConfig != nil {
		if err := s.ScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TemplateConfig != nil {
		if err := s.TemplateConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitDynamicImageJobRequestInput struct {
	// The input media resource.
	//
	// - If `Type` is set to `OSS`, specify the OSS URL of the input file.
	//
	// - If `Type` is set to `Media`, specify the media asset ID.
	//
	// An OSS URL must be in one of the following formats:
	//
	// 1. `oss://bucket/object`
	//
	// 2. `http(s)://bucket.oss-[RegionId].aliyuncs.com/object`
	//
	// In these formats, `bucket` is the name of an OSS bucket in the same region as the current project, and `object` is the file path.
	//
	// > The specified OSS bucket must be registered in IMS [storage management](https://help.aliyun.com/document_detail/609918.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the job input. Valid values:
	//
	// - `OSS`: An Object Storage Service (OSS) file URL.
	//
	// - `Media`: A media asset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitDynamicImageJobRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicImageJobRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitDynamicImageJobRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitDynamicImageJobRequestInput) SetMedia(v string) *SubmitDynamicImageJobRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitDynamicImageJobRequestInput) SetType(v string) *SubmitDynamicImageJobRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitDynamicImageJobRequestInput) Validate() error {
	return dara.Validate(s)
}

type SubmitDynamicImageJobRequestOutput struct {
	// The destination OSS URL for the output file. This parameter is required when `Type` is set to `OSS`. The URL must be in one of the following formats:
	//
	// - `oss://bucket/object`
	//
	// - `http(s)://bucket.oss-[regionId].aliyuncs.com/object`
	//
	// In these formats, `bucket` is the name of an OSS bucket in the same region as the current project, and `object` is the file path.
	//
	// > The specified OSS bucket must be registered in IMS [storage management](https://help.aliyun.com/document_detail/609918.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the job output. Valid values:
	//
	// - `OSS`: The output is an OSS file.
	//
	// - `Media`: The output is a new media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// Media
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitDynamicImageJobRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicImageJobRequestOutput) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobRequestOutput) GetMedia() *string {
	return s.Media
}

func (s *SubmitDynamicImageJobRequestOutput) GetType() *string {
	return s.Type
}

func (s *SubmitDynamicImageJobRequestOutput) SetMedia(v string) *SubmitDynamicImageJobRequestOutput {
	s.Media = &v
	return s
}

func (s *SubmitDynamicImageJobRequestOutput) SetType(v string) *SubmitDynamicImageJobRequestOutput {
	s.Type = &v
	return s
}

func (s *SubmitDynamicImageJobRequestOutput) Validate() error {
	return dara.Validate(s)
}

type SubmitDynamicImageJobRequestScheduleConfig struct {
	// The pipeline ID.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the job. Valid range: [1, 10]. A higher value indicates a higher priority. Default value: 6.
	//
	// example:
	//
	// 6
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s SubmitDynamicImageJobRequestScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicImageJobRequestScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobRequestScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitDynamicImageJobRequestScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *SubmitDynamicImageJobRequestScheduleConfig) SetPipelineId(v string) *SubmitDynamicImageJobRequestScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *SubmitDynamicImageJobRequestScheduleConfig) SetPriority(v int32) *SubmitDynamicImageJobRequestScheduleConfig {
	s.Priority = &v
	return s
}

func (s *SubmitDynamicImageJobRequestScheduleConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitDynamicImageJobRequestTemplateConfig struct {
	// The overwrite parameters.
	OverwriteParams *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitDynamicImageJobRequestTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicImageJobRequestTemplateConfig) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobRequestTemplateConfig) GetOverwriteParams() *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams {
	return s.OverwriteParams
}

func (s *SubmitDynamicImageJobRequestTemplateConfig) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitDynamicImageJobRequestTemplateConfig) SetOverwriteParams(v *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) *SubmitDynamicImageJobRequestTemplateConfig {
	s.OverwriteParams = v
	return s
}

func (s *SubmitDynamicImageJobRequestTemplateConfig) SetTemplateId(v string) *SubmitDynamicImageJobRequestTemplateConfig {
	s.TemplateId = &v
	return s
}

func (s *SubmitDynamicImageJobRequestTemplateConfig) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitDynamicImageJobRequestTemplateConfigOverwriteParams struct {
	// The animated image format. Valid values:
	//
	// - `gif`
	//
	// - `webp`
	//
	// example:
	//
	// gif
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The frame rate. Valid range: [1, 60].
	//
	// example:
	//
	// 15
	Fps *int32 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The height of the output animated image. Valid range: [128, 4096].
	//
	// example:
	//
	// 720
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Specifies whether to enable adaptive orientation based on the long and short edges of the video. Valid values:
	//
	// - **true**: Enables adaptive orientation.
	//
	// - **false**: Disables adaptive orientation.
	//
	// Default value: **true**.
	//
	// > When enabled, this mode sets the output width to the source video\\"s long edge and the output height to its short edge. For a portrait video, its height is treated as the long edge and its width as the short edge.
	//
	// example:
	//
	// false
	LongShortMode *bool `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
	// The scan mode. Valid values:
	//
	// - **interlaced**: Interlaced scanning.
	//
	// - **progressive**: Progressive scanning. This is the default value.
	//
	// example:
	//
	// progressive
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// Specifies the time range of the video to process for the animated image.
	TimeSpan *SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan `json:"TimeSpan,omitempty" xml:"TimeSpan,omitempty" type:"Struct"`
	// The width of the output animated image. Valid range: [128, 4096].
	//
	// example:
	//
	// 1024
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) GetFormat() *string {
	return s.Format
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) GetFps() *int32 {
	return s.Fps
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) GetHeight() *int32 {
	return s.Height
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) GetLongShortMode() *bool {
	return s.LongShortMode
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) GetScanMode() *string {
	return s.ScanMode
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) GetTimeSpan() *SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan {
	return s.TimeSpan
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) GetWidth() *int32 {
	return s.Width
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) SetFormat(v string) *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams {
	s.Format = &v
	return s
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) SetFps(v int32) *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams {
	s.Fps = &v
	return s
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) SetHeight(v int32) *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams {
	s.Height = &v
	return s
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) SetLongShortMode(v bool) *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams {
	s.LongShortMode = &v
	return s
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) SetScanMode(v string) *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams {
	s.ScanMode = &v
	return s
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) SetTimeSpan(v *SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan) *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams {
	s.TimeSpan = v
	return s
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) SetWidth(v int32) *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams {
	s.Width = &v
	return s
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParams) Validate() error {
	if s.TimeSpan != nil {
		if err := s.TimeSpan.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan struct {
	// The duration of the video segment to be processed.
	//
	// - Format: `hh:mm:ss[.SSS]` or `sssss[.SSS]`.
	//
	// - Valid range: `[00:00:00.000,23:59:59.999]` or `[0.000,86399.999]`.
	//
	// example:
	//
	// 01:59:59.999 or 32000.23
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The end time of the video segment to be processed. If this parameter is set, the `Duration` parameter is ignored.
	//
	// - Format: `hh:mm:ss[.SSS]` or `sssss[.SSS]`.
	//
	// - Valid range: `[00:00:00.000,23:59:59.999]` or `[0.000,86399.999]`.
	//
	// example:
	//
	// 01:59:59.999 or 32000.23
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The start time of the video segment to be processed.
	//
	// - Format: `hh:mm:ss[.SSS]` or `sssss[.SSS]`.
	//
	// - Valid range: `[00:00:00.000,23:59:59.999]` or `[0.000,86399.999]`.
	//
	// example:
	//
	// 01:59:59.999 or 32000.23
	Seek *string `json:"Seek,omitempty" xml:"Seek,omitempty"`
}

func (s SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan) GetDuration() *string {
	return s.Duration
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan) GetEnd() *string {
	return s.End
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan) GetSeek() *string {
	return s.Seek
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan) SetDuration(v string) *SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan {
	s.Duration = &v
	return s
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan) SetEnd(v string) *SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan {
	s.End = &v
	return s
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan) SetSeek(v string) *SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan {
	s.Seek = &v
	return s
}

func (s *SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan) Validate() error {
	return dara.Validate(s)
}
