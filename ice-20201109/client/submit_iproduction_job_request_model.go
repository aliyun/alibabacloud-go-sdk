// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIProductionJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *SubmitIProductionJobRequest
	GetFunctionName() *string
	SetInput(v *SubmitIProductionJobRequestInput) *SubmitIProductionJobRequest
	GetInput() *SubmitIProductionJobRequestInput
	SetJobParams(v string) *SubmitIProductionJobRequest
	GetJobParams() *string
	SetModelId(v string) *SubmitIProductionJobRequest
	GetModelId() *string
	SetName(v string) *SubmitIProductionJobRequest
	GetName() *string
	SetOutput(v *SubmitIProductionJobRequestOutput) *SubmitIProductionJobRequest
	GetOutput() *SubmitIProductionJobRequestOutput
	SetScheduleConfig(v *SubmitIProductionJobRequestScheduleConfig) *SubmitIProductionJobRequest
	GetScheduleConfig() *SubmitIProductionJobRequestScheduleConfig
	SetTemplateId(v string) *SubmitIProductionJobRequest
	GetTemplateId() *string
	SetUserData(v string) *SubmitIProductionJobRequest
	GetUserData() *string
}

type SubmitIProductionJobRequest struct {
	// The name of the algorithm function to use. Valid values:
	//
	// - **Cover**: intelligent cover
	//
	// - **VideoClip**: video synopsis
	//
	// - **VideoDelogo**: video logo removal
	//
	// - **VideoDetext**: video subtitle removal
	//
	// - **CaptionExtraction**: caption extraction
	//
	// - **VideoGreenScreenMatting**: image matting
	//
	// - **FaceBeauty**: video face beautification
	//
	// - **VideoH2V**: intelligent landscape-to-portrait
	//
	// - **MusicSegmentDetect**: chorus detection
	//
	// - **AudioBeatDetection**: beat detection
	//
	// - **AudioQualityAssessment**: audio quality assessment
	//
	// - **SpeechDenoise**: speech denoising
	//
	// - **AudioMixing**: audio mixing
	//
	// - **MusicDemix**: vocal and accompaniment separation
	//
	// This parameter is required.
	//
	// example:
	//
	// Cover
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The input media. Object Storage Service (OSS) paths and media asset IDs are supported.
	//
	// Different algorithm functions have different input file requirements. For more information, see the supplementary description below.
	//
	// This parameter is required.
	Input *SubmitIProductionJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The algorithm job parameters. This is a JSON object. The parameters vary depending on the algorithm. For more information, see the supplementary description.
	//
	// example:
	//
	// {"Model":"gif"}
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// The algorithm model ID. If this parameter is left empty, the default model for the corresponding function is used. In most cases, leave this parameter empty to use the default model.
	//
	// The following algorithm functions have non-default models available:
	//
	// 	- VideoDetext
	//
	//   	- ModelId = algo-video-detext-new: a subtitle removal algorithm with better results but slower speed and higher cost than the default algorithm.
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The job name. The name can be up to 100 characters in length.
	//
	// example:
	//
	// Test task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output media. OSS paths and media asset IDs are supported.
	//
	// Different algorithm functions produce different output files. For more information, see the supplementary description below.
	//
	// This parameter is required.
	Output *SubmitIProductionJobRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The job scheduling configuration.
	ScheduleConfig *SubmitIProductionJobRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The custom user data, which is returned as-is when you retrieve the result. The value can be up to 256 characters in length.
	//
	// example:
	//
	// {"test":1}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitIProductionJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitIProductionJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitIProductionJobRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *SubmitIProductionJobRequest) GetInput() *SubmitIProductionJobRequestInput {
	return s.Input
}

func (s *SubmitIProductionJobRequest) GetJobParams() *string {
	return s.JobParams
}

func (s *SubmitIProductionJobRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitIProductionJobRequest) GetName() *string {
	return s.Name
}

func (s *SubmitIProductionJobRequest) GetOutput() *SubmitIProductionJobRequestOutput {
	return s.Output
}

func (s *SubmitIProductionJobRequest) GetScheduleConfig() *SubmitIProductionJobRequestScheduleConfig {
	return s.ScheduleConfig
}

func (s *SubmitIProductionJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitIProductionJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitIProductionJobRequest) SetFunctionName(v string) *SubmitIProductionJobRequest {
	s.FunctionName = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetInput(v *SubmitIProductionJobRequestInput) *SubmitIProductionJobRequest {
	s.Input = v
	return s
}

func (s *SubmitIProductionJobRequest) SetJobParams(v string) *SubmitIProductionJobRequest {
	s.JobParams = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetModelId(v string) *SubmitIProductionJobRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetName(v string) *SubmitIProductionJobRequest {
	s.Name = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetOutput(v *SubmitIProductionJobRequestOutput) *SubmitIProductionJobRequest {
	s.Output = v
	return s
}

func (s *SubmitIProductionJobRequest) SetScheduleConfig(v *SubmitIProductionJobRequestScheduleConfig) *SubmitIProductionJobRequest {
	s.ScheduleConfig = v
	return s
}

func (s *SubmitIProductionJobRequest) SetTemplateId(v string) *SubmitIProductionJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitIProductionJobRequest) SetUserData(v string) *SubmitIProductionJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitIProductionJobRequest) Validate() error {
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
	return nil
}

type SubmitIProductionJobRequestInput struct {
	// The input media. OSS paths and media asset IDs are supported.
	//
	// OSS path rules (use either format):
	//
	// 1. oss://bucket/object
	//
	// 2. http(s)://bucket.oss-[regionId].aliyuncs.com/object
	//
	// where bucket is the name of an OSS bucket in the same region as the current project, and object is the file path.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The media type. Valid values:
	//
	// - OSS: an OSS path
	//
	// - Media: a media asset ID
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitIProductionJobRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitIProductionJobRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitIProductionJobRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitIProductionJobRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitIProductionJobRequestInput) SetMedia(v string) *SubmitIProductionJobRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitIProductionJobRequestInput) SetType(v string) *SubmitIProductionJobRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitIProductionJobRequestInput) Validate() error {
	return dara.Validate(s)
}

type SubmitIProductionJobRequestOutput struct {
	// The business type to which the media asset belongs.
	//
	// example:
	//
	// IMS
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The OSS path of the output file when Type is set to Media. The bucket must be registered in IMS or VOD.
	//
	// example:
	//
	// http(s)://bucket.oss-[RegionId].aliyuncs.com/object
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
	// The media type. Valid values:
	//
	// - OSS: an OSS path
	//
	// - Media: a media asset ID
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitIProductionJobRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitIProductionJobRequestOutput) GoString() string {
	return s.String()
}

func (s *SubmitIProductionJobRequestOutput) GetBiz() *string {
	return s.Biz
}

func (s *SubmitIProductionJobRequestOutput) GetMedia() *string {
	return s.Media
}

func (s *SubmitIProductionJobRequestOutput) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *SubmitIProductionJobRequestOutput) GetType() *string {
	return s.Type
}

func (s *SubmitIProductionJobRequestOutput) SetBiz(v string) *SubmitIProductionJobRequestOutput {
	s.Biz = &v
	return s
}

func (s *SubmitIProductionJobRequestOutput) SetMedia(v string) *SubmitIProductionJobRequestOutput {
	s.Media = &v
	return s
}

func (s *SubmitIProductionJobRequestOutput) SetOutputUrl(v string) *SubmitIProductionJobRequestOutput {
	s.OutputUrl = &v
	return s
}

func (s *SubmitIProductionJobRequestOutput) SetType(v string) *SubmitIProductionJobRequestOutput {
	s.Type = &v
	return s
}

func (s *SubmitIProductionJobRequestOutput) Validate() error {
	return dara.Validate(s)
}

type SubmitIProductionJobRequestScheduleConfig struct {
	// The pipeline ID.
	//
	// example:
	//
	// 5246b8d12a62433ab77845074039c3dc
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority. Valid values: 1 to 10. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 6
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s SubmitIProductionJobRequestScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitIProductionJobRequestScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitIProductionJobRequestScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitIProductionJobRequestScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *SubmitIProductionJobRequestScheduleConfig) SetPipelineId(v string) *SubmitIProductionJobRequestScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *SubmitIProductionJobRequestScheduleConfig) SetPriority(v int32) *SubmitIProductionJobRequestScheduleConfig {
	s.Priority = &v
	return s
}

func (s *SubmitIProductionJobRequestScheduleConfig) Validate() error {
	return dara.Validate(s)
}
