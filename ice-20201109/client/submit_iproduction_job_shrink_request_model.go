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
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
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
	OutputShrink *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The configuration for job scheduling.
	ScheduleConfigShrink *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
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
