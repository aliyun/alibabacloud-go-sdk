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
	// The aspect ratio. Valid values: 16:9 (default), 9:16, 4:3, 3:4, and 1:1.
	//
	// example:
	//
	// 4:3
	AspectRatio *string `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	// The idempotency parameter.
	//
	// example:
	//
	// ****3e761e9d11edba640c42a1b7****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The task input. This is a JSON string that contains the following fields:
	//
	// - Prompt: String. Required. The prompt.
	//
	// - Medias: the media list. Required when the task type is `image_to_image`. A maximum of 9 items are supported.
	//
	// > The Media struct contains: Type, the media type, String, valid value: image. URL, the media download URL, String.
	//
	// >
	//
	// example:
	//
	// {\\"Bucket\\":\\"dbj-app-prod\\",\\"Location\\":\\"oss-cn-hangzhou\\",\\"Object\\":\\"classpal/1767838045280pzmgnvwe.mp4\\"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The task feature parameters. This is a JSON string. No configuration is required at this time.
	//
	// example:
	//
	// {}
	JobParameters *string `json:"JobParameters,omitempty" xml:"JobParameters,omitempty"`
	// The type of the generation task. Valid values:
	//
	// - text_to_image: text-to-image.
	//
	// - image_to_image: image-to-image.
	//
	// example:
	//
	// text_to_image
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The model name.
	//
	// example:
	//
	// wan2.7-image
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The number of images. Valid values: 1 to 4. Default value: 1.
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
	// The scenario. This is an enumeration type. Currently, only `general` is supported.
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
