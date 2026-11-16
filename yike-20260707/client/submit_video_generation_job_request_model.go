// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoGenerationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAspectRatio(v string) *SubmitVideoGenerationJobRequest
	GetAspectRatio() *string
	SetClientToken(v string) *SubmitVideoGenerationJobRequest
	GetClientToken() *string
	SetDuration(v string) *SubmitVideoGenerationJobRequest
	GetDuration() *string
	SetInput(v string) *SubmitVideoGenerationJobRequest
	GetInput() *string
	SetJobParameters(v string) *SubmitVideoGenerationJobRequest
	GetJobParameters() *string
	SetJobType(v string) *SubmitVideoGenerationJobRequest
	GetJobType() *string
	SetModel(v string) *SubmitVideoGenerationJobRequest
	GetModel() *string
	SetN(v int32) *SubmitVideoGenerationJobRequest
	GetN() *int32
	SetOutput(v string) *SubmitVideoGenerationJobRequest
	GetOutput() *string
	SetResolution(v string) *SubmitVideoGenerationJobRequest
	GetResolution() *string
	SetScene(v string) *SubmitVideoGenerationJobRequest
	GetScene() *string
	SetUserData(v string) *SubmitVideoGenerationJobRequest
	GetUserData() *string
}

type SubmitVideoGenerationJobRequest struct {
	// The aspect ratio. Valid values: 16:9 (default), 9:16, 4:3, 3:4, 1:1, and adaptive (valid only for wan3.0-video and wan3.0-video-prime).
	//
	// example:
	//
	// 9:16
	AspectRatio *string `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	// The idempotency parameter. A unique, case-sensitive, 32-character random string. This parameter ensures that the request is completed no more than once, preventing duplicate operations caused by multiple retries.
	//
	// example:
	//
	// ****3e761e9d11edba640c42a1b7****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The output duration. Valid values: 4 to 15 seconds. Default value: 5 seconds.
	//
	// - For wan3.0-video and wan3.0-video-prime, the maximum duration is 30 seconds.
	//
	// example:
	//
	// 5
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The task input. This parameter is required. The value is a JSON string that contains the following fields:
	//
	// - Prompt: the prompt. This field is required. Data type: String.
	//
	// - Medias: the media list.
	//
	//   - If JobType is set to image_to_video, this field is required and only one media file is required.
	//
	//   - If JobType is set to first_last_frame, this field is required and exactly two media files are required.
	//
	//   - If JobType is set to reference_to_video, this field is required and up to 9 media files are allowed. For wan3.0-video and wan3.0-video-prime, up to 20 media files are allowed, including up to 10 images, 5 videos, and 5 audio files. The total duration of audio and video files cannot exceed 15 seconds.
	//
	// > The Media structure contains the following fields: Type, the media type. Data type: String. Valid values: image, video, and audio. URL, the download URL of the media file. Data type: String. MediaId, the media asset ID. Data type: String.
	//
	// >
	//
	// example:
	//
	// {"Prompt":"Image 1 is on the basketball court, performing a slam dunk with Image 2","Medias":[{"Type":"image","Url":"https://xxx/xxx.jpg"},{"Type":"image","Url":"https://xxx/xxx.jpg"}]}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The task feature parameters. The value is a JSON string that contains the following fields:
	//
	// - EnableAudio: specifies whether to output audio. Data type: Boolean. This field is optional. Valid values:
	//
	//   - true: outputs audio.
	//
	//   - false: does not output audio.
	//
	// - Watermark: specifies whether to include a watermark. Data type: Boolean. This field is optional. Valid values:
	//
	//   - true: adds an "AI-generated" watermark to the lower-right corner of the video.
	//
	//   - false: does not add a watermark to the generated video.
	//
	// - PromptExtend: specifies whether to enable intelligent prompt rewriting. Data type: Boolean. This field is optional. This field is valid only for wan3.0-video and wan3.0-video-prime. Valid values:
	//
	//   - true: enables intelligent prompt rewriting. This is the default value.
	//
	//   - false: disables intelligent prompt rewriting.
	//
	// - PostProcessing: the post-processing feature switch. Data type: Object. This field contains the EnableSubtitleRemoval field. Data type: Boolean. Valid values:
	//
	//   - true: enables subtitle removal.
	//
	//   - false: disables subtitle removal.
	//
	// example:
	//
	// {}
	JobParameters *string `json:"JobParameters,omitempty" xml:"JobParameters,omitempty"`
	// The task type. This parameter is required. Valid values:
	//
	// - text_to_video: text-to-video
	//
	// - image_to_video: image-to-video
	//
	// - first_last_frame: first and last frames to video
	//
	// - reference_to_video: reference to video
	//
	// example:
	//
	// text_to_video
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The model name. This parameter is required. Valid values:
	//
	// - Wonder-Ultra
	//
	// - Wonder-Pro
	//
	// - Wonder-Standard
	//
	// - wan3.0-video
	//
	// - wan3.0-video-prime
	//
	// - happyhorse-1.1
	//
	// - happyhorse-1.0
	//
	// - wan2.7
	//
	// example:
	//
	// happyhorse-1.1
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The number of outputs. Valid values: 1 to 4. Default value: 1.
	//
	// example:
	//
	// 1
	N *int32 `json:"N,omitempty" xml:"N,omitempty"`
	// The output configuration. The value is a JSON string. OssUri is the optional customer OSS output directory. If this parameter is not specified, a signed URL for the service-generated output is returned.
	//
	// example:
	//
	// {"OssUri":"oss://example-bucket/video-translation/output/"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The resolution. Valid values:
	//
	// - 1080P
	//
	// - 720P: the default value
	//
	// - 480P: valid only for wan3.0-video and wan3.0-video-prime
	//
	// example:
	//
	// 720P
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The scene type. Valid values: general.
	//
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The user-defined parameters. The value is a JSON string. The callback result returns this parameter as is. The system reserved field NotifyAddress specifies the callback URL. The system calls this URL after the task is completed.
	//
	// example:
	//
	// {"NotifyAddress": "http://xxx.callback.url"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitVideoGenerationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoGenerationJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoGenerationJobRequest) GetAspectRatio() *string {
	return s.AspectRatio
}

func (s *SubmitVideoGenerationJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitVideoGenerationJobRequest) GetDuration() *string {
	return s.Duration
}

func (s *SubmitVideoGenerationJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitVideoGenerationJobRequest) GetJobParameters() *string {
	return s.JobParameters
}

func (s *SubmitVideoGenerationJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *SubmitVideoGenerationJobRequest) GetModel() *string {
	return s.Model
}

func (s *SubmitVideoGenerationJobRequest) GetN() *int32 {
	return s.N
}

func (s *SubmitVideoGenerationJobRequest) GetOutput() *string {
	return s.Output
}

func (s *SubmitVideoGenerationJobRequest) GetResolution() *string {
	return s.Resolution
}

func (s *SubmitVideoGenerationJobRequest) GetScene() *string {
	return s.Scene
}

func (s *SubmitVideoGenerationJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitVideoGenerationJobRequest) SetAspectRatio(v string) *SubmitVideoGenerationJobRequest {
	s.AspectRatio = &v
	return s
}

func (s *SubmitVideoGenerationJobRequest) SetClientToken(v string) *SubmitVideoGenerationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitVideoGenerationJobRequest) SetDuration(v string) *SubmitVideoGenerationJobRequest {
	s.Duration = &v
	return s
}

func (s *SubmitVideoGenerationJobRequest) SetInput(v string) *SubmitVideoGenerationJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitVideoGenerationJobRequest) SetJobParameters(v string) *SubmitVideoGenerationJobRequest {
	s.JobParameters = &v
	return s
}

func (s *SubmitVideoGenerationJobRequest) SetJobType(v string) *SubmitVideoGenerationJobRequest {
	s.JobType = &v
	return s
}

func (s *SubmitVideoGenerationJobRequest) SetModel(v string) *SubmitVideoGenerationJobRequest {
	s.Model = &v
	return s
}

func (s *SubmitVideoGenerationJobRequest) SetN(v int32) *SubmitVideoGenerationJobRequest {
	s.N = &v
	return s
}

func (s *SubmitVideoGenerationJobRequest) SetOutput(v string) *SubmitVideoGenerationJobRequest {
	s.Output = &v
	return s
}

func (s *SubmitVideoGenerationJobRequest) SetResolution(v string) *SubmitVideoGenerationJobRequest {
	s.Resolution = &v
	return s
}

func (s *SubmitVideoGenerationJobRequest) SetScene(v string) *SubmitVideoGenerationJobRequest {
	s.Scene = &v
	return s
}

func (s *SubmitVideoGenerationJobRequest) SetUserData(v string) *SubmitVideoGenerationJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitVideoGenerationJobRequest) Validate() error {
	return dara.Validate(s)
}
