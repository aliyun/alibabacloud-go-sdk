// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoTranslationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SubmitVideoTranslationJobRequest
	GetClientToken() *string
	SetDescription(v string) *SubmitVideoTranslationJobRequest
	GetDescription() *string
	SetInput(v string) *SubmitVideoTranslationJobRequest
	GetInput() *string
	SetJobParameters(v string) *SubmitVideoTranslationJobRequest
	GetJobParameters() *string
	SetJobType(v string) *SubmitVideoTranslationJobRequest
	GetJobType() *string
	SetOutput(v string) *SubmitVideoTranslationJobRequest
	GetOutput() *string
	SetTitle(v string) *SubmitVideoTranslationJobRequest
	GetTitle() *string
	SetUserData(v string) *SubmitVideoTranslationJobRequest
	GetUserData() *string
}

type SubmitVideoTranslationJobRequest struct {
	// example:
	//
	// ****3e761e9d11edba640c42a1b7****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"Video":"oss://bucket/path/input.mp4"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"SourceLanguage":"zh","TargetLanguage":"en","NeedDetext":false,"NeedVisualTranslate":false}
	JobParameters *string `json:"JobParameters,omitempty" xml:"JobParameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VoiceTranslate
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"OssUri":"oss://bucket/output/"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// {}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitVideoTranslationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoTranslationJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoTranslationJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitVideoTranslationJobRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitVideoTranslationJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitVideoTranslationJobRequest) GetJobParameters() *string {
	return s.JobParameters
}

func (s *SubmitVideoTranslationJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *SubmitVideoTranslationJobRequest) GetOutput() *string {
	return s.Output
}

func (s *SubmitVideoTranslationJobRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitVideoTranslationJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitVideoTranslationJobRequest) SetClientToken(v string) *SubmitVideoTranslationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetDescription(v string) *SubmitVideoTranslationJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetInput(v string) *SubmitVideoTranslationJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetJobParameters(v string) *SubmitVideoTranslationJobRequest {
	s.JobParameters = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetJobType(v string) *SubmitVideoTranslationJobRequest {
	s.JobType = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetOutput(v string) *SubmitVideoTranslationJobRequest {
	s.Output = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetTitle(v string) *SubmitVideoTranslationJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetUserData(v string) *SubmitVideoTranslationJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) Validate() error {
	return dara.Validate(s)
}
