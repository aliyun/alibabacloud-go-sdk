// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIProductionJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *SubmitIProductionJobShrinkRequest
	GetFunctionName() *string
	SetInputShrink(v string) *SubmitIProductionJobShrinkRequest
	GetInputShrink() *string
	SetJobParams(v string) *SubmitIProductionJobShrinkRequest
	GetJobParams() *string
	SetModelId(v string) *SubmitIProductionJobShrinkRequest
	GetModelId() *string
	SetName(v string) *SubmitIProductionJobShrinkRequest
	GetName() *string
	SetOutputShrink(v string) *SubmitIProductionJobShrinkRequest
	GetOutputShrink() *string
	SetScheduleConfigShrink(v string) *SubmitIProductionJobShrinkRequest
	GetScheduleConfigShrink() *string
	SetTemplateId(v string) *SubmitIProductionJobShrinkRequest
	GetTemplateId() *string
	SetUserData(v string) *SubmitIProductionJobShrinkRequest
	GetUserData() *string
}

type SubmitIProductionJobShrinkRequest struct {
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
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
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
	OutputShrink *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The job scheduling configuration.
	ScheduleConfigShrink *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
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

func (s SubmitIProductionJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitIProductionJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitIProductionJobShrinkRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *SubmitIProductionJobShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SubmitIProductionJobShrinkRequest) GetJobParams() *string {
	return s.JobParams
}

func (s *SubmitIProductionJobShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitIProductionJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *SubmitIProductionJobShrinkRequest) GetOutputShrink() *string {
	return s.OutputShrink
}

func (s *SubmitIProductionJobShrinkRequest) GetScheduleConfigShrink() *string {
	return s.ScheduleConfigShrink
}

func (s *SubmitIProductionJobShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitIProductionJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitIProductionJobShrinkRequest) SetFunctionName(v string) *SubmitIProductionJobShrinkRequest {
	s.FunctionName = &v
	return s
}

func (s *SubmitIProductionJobShrinkRequest) SetInputShrink(v string) *SubmitIProductionJobShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SubmitIProductionJobShrinkRequest) SetJobParams(v string) *SubmitIProductionJobShrinkRequest {
	s.JobParams = &v
	return s
}

func (s *SubmitIProductionJobShrinkRequest) SetModelId(v string) *SubmitIProductionJobShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitIProductionJobShrinkRequest) SetName(v string) *SubmitIProductionJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *SubmitIProductionJobShrinkRequest) SetOutputShrink(v string) *SubmitIProductionJobShrinkRequest {
	s.OutputShrink = &v
	return s
}

func (s *SubmitIProductionJobShrinkRequest) SetScheduleConfigShrink(v string) *SubmitIProductionJobShrinkRequest {
	s.ScheduleConfigShrink = &v
	return s
}

func (s *SubmitIProductionJobShrinkRequest) SetTemplateId(v string) *SubmitIProductionJobShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitIProductionJobShrinkRequest) SetUserData(v string) *SubmitIProductionJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitIProductionJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
