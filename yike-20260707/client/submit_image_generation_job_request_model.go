// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageGenerationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAspectRatio(v string) *SubmitImageGenerationJobRequest
	GetAspectRatio() *string
	SetClientToken(v string) *SubmitImageGenerationJobRequest
	GetClientToken() *string
	SetInput(v string) *SubmitImageGenerationJobRequest
	GetInput() *string
	SetJobParameters(v string) *SubmitImageGenerationJobRequest
	GetJobParameters() *string
	SetJobType(v string) *SubmitImageGenerationJobRequest
	GetJobType() *string
	SetModel(v string) *SubmitImageGenerationJobRequest
	GetModel() *string
	SetN(v string) *SubmitImageGenerationJobRequest
	GetN() *string
	SetResolution(v string) *SubmitImageGenerationJobRequest
	GetResolution() *string
	SetScene(v string) *SubmitImageGenerationJobRequest
	GetScene() *string
	SetUserData(v string) *SubmitImageGenerationJobRequest
	GetUserData() *string
}

type SubmitImageGenerationJobRequest struct {
	// The aspect ratio. Valid values: 16:9 (default), 9:16, 4:3, 3:4, 1:1, and 21:9.
	//
	// - The qwen-image-3.0 model also supports the following values: 4:5, 5:4, 2:3, and 3:2.
	//
	// example:
	//
	// 4:3
	AspectRatio *string `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	// The idempotency parameter. A unique, case-sensitive, 32-character random string used to ensure that the request is completed only once. This prevents duplicate operations caused by multiple retries.
	//
	// example:
	//
	// ****3e761e9d11edba640c42a1b7****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The task input. Required. A JSON string that contains the following fields:
	//
	// - Prompt: the prompt. String. Required.
	//
	// - Medias: the list of media files. Required when the task type is `image_to_image`. Up to 9 files.
	//
	// > The Media structure contains the following fields: Type, the media type. String. Valid value: image. URL, the download URL of the media file. String. MediaId, the media asset ID. String.
	//
	// >
	//
	// example:
	//
	// {"Prompt":"xxx","Medias":[{"Type":"image","URL":"http://test.oss-cn-shanghai.aliyuncs.com/test.png"}]}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The task feature parameters. A JSON string that contains the following fields:
	//
	// - PromptExtend: specifies whether to optimize the prompt. Boolean. Valid values: `true` and `false`.
	//
	// - PromptExtendMode: the prompt optimization mode. Valid values: `direct` (default) and `agent` (supported only for text-to-image). Only the qwen-image-3.0 model supports this parameter.
	//
	// example:
	//
	// {}
	JobParameters *string `json:"JobParameters,omitempty" xml:"JobParameters,omitempty"`
	// The type of the generation task. Required. Valid values:
	//
	// - text_to_image: text-to-image
	//
	// - image_to_image: image-to-image
	//
	// example:
	//
	// text_to_image
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The name of the model. Required. Valid values:
	//
	// - Wonder-Image-2.5-Pro
	//
	// - Wonder-Image-2.5-Fast
	//
	// - Wonder-Image-2
	//
	// - qwen-image-3.0
	//
	// - qwen-image-2.0-pro
	//
	// - qwen-image-2.0
	//
	// - wan2.7-image
	//
	// example:
	//
	// wan2.7-image
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The number of images to generate. Valid values: 1 to 4. Default value: 1.
	//
	// example:
	//
	// 1
	N *string `json:"N,omitempty" xml:"N,omitempty"`
	// The resolution. Valid values: 1K (default), 2K, and 4K.
	//
	// example:
	//
	// 720P
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The scene. Valid values: `general`.
	//
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The user business data in JSON format.
	//
	// example:
	//
	// {"env":"prd"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitImageGenerationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageGenerationJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitImageGenerationJobRequest) GetAspectRatio() *string {
	return s.AspectRatio
}

func (s *SubmitImageGenerationJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitImageGenerationJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitImageGenerationJobRequest) GetJobParameters() *string {
	return s.JobParameters
}

func (s *SubmitImageGenerationJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *SubmitImageGenerationJobRequest) GetModel() *string {
	return s.Model
}

func (s *SubmitImageGenerationJobRequest) GetN() *string {
	return s.N
}

func (s *SubmitImageGenerationJobRequest) GetResolution() *string {
	return s.Resolution
}

func (s *SubmitImageGenerationJobRequest) GetScene() *string {
	return s.Scene
}

func (s *SubmitImageGenerationJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitImageGenerationJobRequest) SetAspectRatio(v string) *SubmitImageGenerationJobRequest {
	s.AspectRatio = &v
	return s
}

func (s *SubmitImageGenerationJobRequest) SetClientToken(v string) *SubmitImageGenerationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitImageGenerationJobRequest) SetInput(v string) *SubmitImageGenerationJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitImageGenerationJobRequest) SetJobParameters(v string) *SubmitImageGenerationJobRequest {
	s.JobParameters = &v
	return s
}

func (s *SubmitImageGenerationJobRequest) SetJobType(v string) *SubmitImageGenerationJobRequest {
	s.JobType = &v
	return s
}

func (s *SubmitImageGenerationJobRequest) SetModel(v string) *SubmitImageGenerationJobRequest {
	s.Model = &v
	return s
}

func (s *SubmitImageGenerationJobRequest) SetN(v string) *SubmitImageGenerationJobRequest {
	s.N = &v
	return s
}

func (s *SubmitImageGenerationJobRequest) SetResolution(v string) *SubmitImageGenerationJobRequest {
	s.Resolution = &v
	return s
}

func (s *SubmitImageGenerationJobRequest) SetScene(v string) *SubmitImageGenerationJobRequest {
	s.Scene = &v
	return s
}

func (s *SubmitImageGenerationJobRequest) SetUserData(v string) *SubmitImageGenerationJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitImageGenerationJobRequest) Validate() error {
	return dara.Validate(s)
}
