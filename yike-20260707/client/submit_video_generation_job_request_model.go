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
	SetResolution(v string) *SubmitVideoGenerationJobRequest
	GetResolution() *string
	SetScene(v string) *SubmitVideoGenerationJobRequest
	GetScene() *string
	SetUserData(v string) *SubmitVideoGenerationJobRequest
	GetUserData() *string
}

type SubmitVideoGenerationJobRequest struct {
	// example:
	//
	// 9:16
	AspectRatio *string `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	// example:
	//
	// ****3e761e9d11edba640c42a1b7****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 5
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// {"Prompt":"图1在篮球场上，用图2来了个灌篮","Medias":[{"Type":"image","Url":"https://xxx/xxx.jpg"},{"Type":"image","Url":"https://xxx/xxx.jpg"}]}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// {}
	JobParameters *string `json:"JobParameters,omitempty" xml:"JobParameters,omitempty"`
	// example:
	//
	// text_to_video
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// happyhorse-1.1
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 1
	N *int32 `json:"N,omitempty" xml:"N,omitempty"`
	// example:
	//
	// 720P
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// {"env":"prd"}
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
