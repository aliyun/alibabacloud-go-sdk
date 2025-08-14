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
	// The input of the job.
	//
	// This parameter is required.
	Input *SubmitDynamicImageJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The name of the job.
	//
	// example:
	//
	// SampleJob
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output of the job.
	//
	// This parameter is required.
	Output *SubmitDynamicImageJobRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The scheduling settings.
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
	return dara.Validate(s)
}

type SubmitDynamicImageJobRequestInput struct {
	// The input file. If Type is set to OSS, set this parameter to the URL of an OSS object. If Type is set to Media, set this parameter to the ID of a media asset. The URL of an OSS object can be in one of the following formats:
	//
	// 1.  oss://bucket/object
	//
	// 2.  http(s)://bucket.oss-[RegionId].aliyuncs.com/object
	//
	// In the URL, bucket specifies an OSS bucket that resides in the same region as the job, and object specifies the object URL in OSS.
	//
	// >  Before you use the OSS bucket in the URL, you must add the bucket on the [Storage Management](https://help.aliyun.com/document_detail/609918.html) page of the Intelligent Media Services (IMS) console.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the input file. Valid values:
	//
	// 1.  OSS: an Object Storage Service (OSS) object.
	//
	// 2.  Media: a media asset.
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
	// The output file. The file can be an OSS object or a media asset. The URL of an OSS object can be in one of the following formats:
	//
	// 	- oss://bucket/object
	//
	// 	- http(s)://bucket.oss-[regionId].aliyuncs.com/object
	//
	// In the URL, bucket specifies an OSS bucket that resides in the same region as the job, and object specifies the object URL in OSS.
	//
	// >  Before you use the OSS bucket in the URL, you must add the bucket on the [Storage Management](https://help.aliyun.com/document_detail/609918.html) page of the IMS console.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the output file. Valid values:
	//
	// 1.  OSS: an OSS object.
	//
	// 2.  Media: a media asset.
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
	// The ID of the MPS queue to which the job was submitted.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority. Valid values: 1 to 10. Default value: 6. A greater value specifies a higher priority.
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
	// The parameters that are used to overwrite the corresponding parameters.
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
	return dara.Validate(s)
}

type SubmitDynamicImageJobRequestTemplateConfigOverwriteParams struct {
	// The format of the animated image. Valid values:
	//
	// 	- **gif**
	//
	// 	- **webp**
	//
	// example:
	//
	// gif
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The frame rate. Valid values: [1,60].
	//
	// example:
	//
	// 15
	Fps *int32 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The height of the animated image. Valid values: [128,4096].
	//
	// example:
	//
	// 720
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Specifies whether to enable the auto-rotate screen feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// Default value: **true**.
	//
	// >  If this feature is enabled, the width of the output video corresponds to the long side of the input video, which is the height of the input video in portrait mode. The height of the output video corresponds to the short side of the input video, which is the width of the input video in portrait mode.
	//
	// example:
	//
	// false
	LongShortMode *bool `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
	// The scan mode. Valid values:
	//
	// 	- **interlaced**
	//
	// 	- **progressive*	- This is the default value.
	//
	// example:
	//
	// progressive
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// The timeline parameters.
	TimeSpan *SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan `json:"TimeSpan,omitempty" xml:"TimeSpan,omitempty" type:"Struct"`
	// The width of the animated image. Valid values: [128,4096].
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
	return dara.Validate(s)
}

type SubmitDynamicImageJobRequestTemplateConfigOverwriteParamsTimeSpan struct {
	// The length of the clip.
	//
	// 	- Format: `hh:mm:ss[.SSS]` or `sssss[.SSS]`.
	//
	// 	- Valid values: `[00:00:00.000,23:59:59.999]` or `[0.000,86399.999]`.
	//
	// example:
	//
	// 01:59:59.999 or 32000.23
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The length of the ending part of the original clip to be cropped out. If you specify this parameter, the Duration parameter becomes invalid.
	//
	// 	- Format: `hh:mm:ss[.SSS]` or `sssss[.SSS]`.
	//
	// 	- Valid values: `[00:00:00.000,23:59:59.999]` or `[0.000,86399.999]`.
	//
	// example:
	//
	// 01:59:59.999 or 32000.23
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The start point of the clip.
	//
	// 	- Format: `hh:mm:ss[.SSS]` or `sssss[.SSS]`.
	//
	// 	- Valid values: `[00:00:00.000,23:59:59.999]` or `[0.000,86399.999]`.
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
