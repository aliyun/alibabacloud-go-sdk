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
	// The idempotency token. A unique, case-sensitive string of up to 32 characters. This token ensures that the request is completed no more than once, preventing duplicate operations caused by multiple retries.
	//
	// example:
	//
	// ****3e761e9d11edba640c42a1b7****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The output duration. Valid values: 4 to 15 seconds. Default value: 5 seconds.
	//
	// - For wan3.0-video and wan3.0-video-prime, the maximum value is 30 seconds.
	//
	// example:
	//
	// 5
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The task input. This parameter is required. The value is a JSON string that contains the following fields:
	//
	// - Prompt: string. Required. The prompt.
	//
	// - Medias: the list of media items.
	//
	//   - If JobType is set to image_to_video, this field is required and only 1 media item is needed.
	//
	//   - If JobType is set to first_last_frame, this field is required and exactly 2 media items are needed.
	//
	//   - If JobType is set to reference_to_video, this field is required and up to 9 media items are allowed. For wan3.0-video and wan3.0-video-prime, up to 20 media items are allowed, including up to 10 images, 5 videos, and 5 audio files. The total duration of audio and video files cannot exceed 15 seconds.
	//
	// > The Media structure contains the following fields: Type, the media type (string). Valid values: `image`, `video`, and `audio`. URL, the media download URL (string). MediaId, the media asset ID (string).
	//
	// >
	//
	// example:
	//
	// {"Prompt":"Person 1 dunks a basketball on the court using the move shown in image 2","Medias":[{"Type":"image","Url":"https://xxx/xxx.jpg"},{"Type":"image","Url":"https://xxx/xxx.jpg"}]}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The task parameters as a JSON string that contains the following fields:
	//
	// - EnableAudio: boolean. Optional. Specifies whether to include audio in the output. Valid values: true and false.
	//
	// - Watermark: boolean. Optional. Specifies whether to include a watermark. Valid values: true (an "AI-generated" watermark is added to the lower-right corner of the video) and false (no watermark is added).
	//
	// - PromptExtend: boolean. Optional. Specifies whether to enable intelligent prompt rewriting. This parameter is valid only for wan3.0-video and wan3.0-video-prime. Valid values: true (enabled, default) and false (disabled).
	//
	// example:
	//
	// {}
	JobParameters *string `json:"JobParameters,omitempty" xml:"JobParameters,omitempty"`
	// The task type. This parameter is required. Valid values:
	//
	// - text_to_video: text-to-video.
	//
	// - image_to_video: image-to-video.
	//
	// - first_last_frame: first and last frame to video.
	//
	// - reference_to_video: reference-to-video.
	//
	// example:
	//
	// text_to_video
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The model name. This parameter is required. Valid values:
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
	// The output configuration as a JSON string. OssUri is an optional OSS output directory. If not specified, a signed URL for the service-generated output is returned.
	//
	// example:
	//
	// {"OssUri":"oss://example-bucket/video-translation/output/"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The resolution. Valid values:
	//
	// - 1080P
	//
	// - 720P: default value.
	//
	// - 480P: valid only for wan3.0-video and wan3.0-video-prime.
	//
	// example:
	//
	// 720P
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The scene type. Currently, only `general` is supported.
	//
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The custom user parameters as a JSON string. These parameters are returned as-is in the callback result. The system reserved field NotifyAddress specifies the callback URL. The system sends a callback to this URL when the task is complete.
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
