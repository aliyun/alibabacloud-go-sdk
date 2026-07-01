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
	// The name of the algorithm function. Valid values:
	//
	// - **Cover**: Generates a smart cover.
	//
	// - **VideoClip**: Creates a video summary.
	//
	// - **VideoDelogo**: Removes logos from a video.
	//
	// - **VideoDetext**: Removes text from a video.
	//
	// - **CaptionExtraction**: Extracts captions from a video.
	//
	// - **VideoGreenScreenMatting**: Performs green screen keying for a video.
	//
	// - **FaceBeauty**: Applies beauty filters to faces in a video.
	//
	// - **VideoH2V**: Converts a horizontal video to a vertical video.
	//
	// - **MusicSegmentDetect**: Detects chorus segments in music.
	//
	// - **AudioBeatDetection**: Detects the beat of an audio track.
	//
	// - **AudioQualityAssessment**: Assesses audio quality.
	//
	// - **SpeechDenoise**: Reduces noise in speech audio.
	//
	// - **AudioMixing**: Mixes audio tracks.
	//
	// - **MusicDemix**: Separates vocals from accompaniment in music.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cover
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The input media asset. You can specify an OSS file or a media asset ID.
	//
	// The requirements for input files vary by algorithm function. For more information, see the supplementary instructions.
	//
	// This parameter is required.
	Input *SubmitIProductionJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The algorithm job parameters, specified as a JSON-formatted string. The content of the JSON object varies by algorithm function. For more information, see the supplementary instructions.
	//
	// example:
	//
	// {"Model":"gif"}
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// The ID of the algorithm model. If you do not specify this parameter, the system uses the default model for the selected function. We recommend leaving this parameter empty unless you need to use a specific alternative model.
	//
	// The following function offers an alternative model:
	//
	// - `VideoDetext`
	//
	//   - Set `ModelId` to `algo-video-detext-new` to use an advanced subtitle removal algorithm. This model provides higher quality results but is slower and more expensive than the default model.
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The name of the job, which can be up to 100 characters long.
	//
	// example:
	//
	// Test task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output destination. You can specify an OSS file path or a media asset ID.
	//
	// The output files vary by algorithm function. For more information, see the supplementary instructions.
	//
	// This parameter is required.
	Output *SubmitIProductionJobRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The configuration for job scheduling.
	ScheduleConfig *SubmitIProductionJobRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The ID of the template.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Custom user data. The system passes this data through and returns it as-is in the callback or response. The length cannot exceed 256 characters.
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
	// The OSS URL of the input file or the ID of the input media asset.
	//
	// The OSS URL can be in one of the following formats:
	//
	// 1. `oss://<bucket>/<object>`
	//
	// 2. `http(s)://<bucket>.oss-<regionId>.aliyuncs.com/<object>`
	//
	//    In these formats, `<bucket>` is the name of an OSS bucket in the same region as your project, and `<object>` is the file path.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of input media. Valid values:
	//
	// - `OSS`: An OSS file path.
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
	// The service to which the media asset belongs.
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
	// If `Type` is set to `Media`, you can use this parameter to specify the OSS URL for the output file. The bucket must be registered in either IMS or VOD.
	//
	// example:
	//
	// http(s)://bucket.oss-[RegionId].aliyuncs.com/object
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
	// The type of the output media. Valid values:
	//
	// - `OSS`: An OSS file path.
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
	// The ID of the pipeline.
	//
	// example:
	//
	// 5246b8d12a62433ab77845074039c3dc
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The job priority, which can be an integer from 1 to 10. A smaller value indicates a higher priority.
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
