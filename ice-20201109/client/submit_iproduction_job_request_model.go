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
	// The name of the algorithm that you want to use for the job. Valid values:
	//
	// 	- **Cover**: This algorithm intelligently generates a thumbnail image for a video.
	//
	// 	- **VideoClip**: This algorithm intelligently generates a summary for a video.
	//
	// 	- **VideoDelogo**: This algorithm removes logos from a video.
	//
	// 	- **VideoDetext**: This algorithm removes captions from a video.
	//
	// 	- **CaptionExtraction**: This algorithm extracts captions from a video and generates the caption file.
	//
	// 	- **VideoGreenScreenMatting**: This algorithm performs green-screen image matting on a video and generates a new video.
	//
	// 	- **FaceBeauty**: This algorithm performs video retouching.
	//
	// 	- **VideoH2V**: This algorithm transforms a video from the landscape mode to the portrait mode.
	//
	// 	- **MusicSegmentDetect**: This algorithm detects the chorus of a song.
	//
	// 	- **AudioBeatDetection**: This algorithm detects rhythms.
	//
	// 	- **AudioQualityAssessment**: This algorithm assesses the audio quality.
	//
	// 	- **SpeechDenoise**: This algorithm performs noise reduction.
	//
	// 	- **AudioMixing**: This algorithm mixes audio streams.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cover
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The input file. The file can be an Object Storage Service (OSS) object or a media asset.
	//
	// This parameter is required.
	Input *SubmitIProductionJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The algorithm-specific parameters. The parameters are specified as JSON objects and vary based on the algorithm. For more information, see the "Parameters of JobParams" section of this topic.
	//
	// example:
	//
	// {"Model":"gif"}
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	ModelId   *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The name of the intelligent production job. The name can be up to 100 characters in length.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output file. The file can be an OSS object or a media asset.
	//
	// This parameter is required.
	Output *SubmitIProductionJobRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The scheduling configuration.
	ScheduleConfig *SubmitIProductionJobRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The template ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The user-defined data that is returned in the response. The value can be up to 1,024 bytes in length.
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
	return dara.Validate(s)
}

type SubmitIProductionJobRequestInput struct {
	// The input file. The file can be an OSS object or a media asset. You can specify the path of an OSS object in one of the following formats:
	//
	// 1.  oss://bucket/object
	//
	// 2.  http(s)://bucket.oss-[regionId].aliyuncs.com/object bucket in the path specifies an OSS bucket that resides in the same region as the intelligent production job. object in the path specifies the object path in OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The media type. Valid values:
	//
	// 	- OSS: OSS object
	//
	// 	- Media: media asset
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
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The output file. If Type is set to OSS, set this parameter to the path of an OSS object. If Type is set to Media, set this parameter to the ID of a media asset. You can specify the path of an OSS object in one of the following formats:
	//
	// 1.  oss://bucket/object
	//
	// 2.  http(s)://bucket.oss-[RegionId].aliyuncs.com/object bucket in the path specifies an OSS bucket that resides in the same region as the intelligent production job. object in the path specifies the object path in OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object
	Media     *string `json:"Media,omitempty" xml:"Media,omitempty"`
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
	// The media type. Valid values:
	//
	// 	- OSS: OSS object
	//
	// 	- Media: media asset
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
	// The ID of the ApsaraVideo Media Processing (MPS) queue.
	//
	// example:
	//
	// 5246b8d12a62433ab77845074039c3dc
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the job. Valid values: 1 to 10. A smaller value indicates a higher priority.
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
