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
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
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
	OutputShrink *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The scheduling configuration.
	ScheduleConfigShrink *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
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
