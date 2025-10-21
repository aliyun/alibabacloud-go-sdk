// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunVideoDetectShotShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIntelliSimpPrompt(v string) *RunVideoDetectShotShrinkRequest
	GetIntelliSimpPrompt() *string
	SetIntelliSimpPromptTemplateId(v string) *RunVideoDetectShotShrinkRequest
	GetIntelliSimpPromptTemplateId() *string
	SetLanguage(v string) *RunVideoDetectShotShrinkRequest
	GetLanguage() *string
	SetModelCustomPromptTemplateId(v string) *RunVideoDetectShotShrinkRequest
	GetModelCustomPromptTemplateId() *string
	SetModelId(v string) *RunVideoDetectShotShrinkRequest
	GetModelId() *string
	SetModelVlCustomPromptTemplateId(v string) *RunVideoDetectShotShrinkRequest
	GetModelVlCustomPromptTemplateId() *string
	SetOptionsShrink(v string) *RunVideoDetectShotShrinkRequest
	GetOptionsShrink() *string
	SetOriginalSessionId(v string) *RunVideoDetectShotShrinkRequest
	GetOriginalSessionId() *string
	SetPreModelId(v string) *RunVideoDetectShotShrinkRequest
	GetPreModelId() *string
	SetPrompt(v string) *RunVideoDetectShotShrinkRequest
	GetPrompt() *string
	SetRecognitionOptionsShrink(v string) *RunVideoDetectShotShrinkRequest
	GetRecognitionOptionsShrink() *string
	SetTaskId(v string) *RunVideoDetectShotShrinkRequest
	GetTaskId() *string
	SetVideoUrl(v string) *RunVideoDetectShotShrinkRequest
	GetVideoUrl() *string
	SetVlPrompt(v string) *RunVideoDetectShotShrinkRequest
	GetVlPrompt() *string
}

type RunVideoDetectShotShrinkRequest struct {
	// example:
	//
	// xxx
	IntelliSimpPrompt *string `json:"intelliSimpPrompt,omitempty" xml:"intelliSimpPrompt,omitempty"`
	// example:
	//
	// intelliSimpShowPrompt
	IntelliSimpPromptTemplateId *string `json:"intelliSimpPromptTemplateId,omitempty" xml:"intelliSimpPromptTemplateId,omitempty"`
	// example:
	//
	// chinese
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// videoDetectShotShowPrompt
	ModelCustomPromptTemplateId *string `json:"modelCustomPromptTemplateId,omitempty" xml:"modelCustomPromptTemplateId,omitempty"`
	// example:
	//
	// deepseek-r1
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// videoDetectShotVlShowPrompt
	ModelVlCustomPromptTemplateId *string `json:"modelVlCustomPromptTemplateId,omitempty" xml:"modelVlCustomPromptTemplateId,omitempty"`
	// This parameter is required.
	OptionsShrink *string `json:"options,omitempty" xml:"options,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5ax
	OriginalSessionId *string `json:"originalSessionId,omitempty" xml:"originalSessionId,omitempty"`
	// example:
	//
	// deepseek-v3.1
	PreModelId *string `json:"preModelId,omitempty" xml:"preModelId,omitempty"`
	// example:
	//
	// xxx
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// This parameter is required.
	RecognitionOptionsShrink *string `json:"recognitionOptions,omitempty" xml:"recognitionOptions,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxx.mp4
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
	// example:
	//
	// xxx
	VlPrompt *string `json:"vlPrompt,omitempty" xml:"vlPrompt,omitempty"`
}

func (s RunVideoDetectShotShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunVideoDetectShotShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunVideoDetectShotShrinkRequest) GetIntelliSimpPrompt() *string {
	return s.IntelliSimpPrompt
}

func (s *RunVideoDetectShotShrinkRequest) GetIntelliSimpPromptTemplateId() *string {
	return s.IntelliSimpPromptTemplateId
}

func (s *RunVideoDetectShotShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *RunVideoDetectShotShrinkRequest) GetModelCustomPromptTemplateId() *string {
	return s.ModelCustomPromptTemplateId
}

func (s *RunVideoDetectShotShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunVideoDetectShotShrinkRequest) GetModelVlCustomPromptTemplateId() *string {
	return s.ModelVlCustomPromptTemplateId
}

func (s *RunVideoDetectShotShrinkRequest) GetOptionsShrink() *string {
	return s.OptionsShrink
}

func (s *RunVideoDetectShotShrinkRequest) GetOriginalSessionId() *string {
	return s.OriginalSessionId
}

func (s *RunVideoDetectShotShrinkRequest) GetPreModelId() *string {
	return s.PreModelId
}

func (s *RunVideoDetectShotShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunVideoDetectShotShrinkRequest) GetRecognitionOptionsShrink() *string {
	return s.RecognitionOptionsShrink
}

func (s *RunVideoDetectShotShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunVideoDetectShotShrinkRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *RunVideoDetectShotShrinkRequest) GetVlPrompt() *string {
	return s.VlPrompt
}

func (s *RunVideoDetectShotShrinkRequest) SetIntelliSimpPrompt(v string) *RunVideoDetectShotShrinkRequest {
	s.IntelliSimpPrompt = &v
	return s
}

func (s *RunVideoDetectShotShrinkRequest) SetIntelliSimpPromptTemplateId(v string) *RunVideoDetectShotShrinkRequest {
	s.IntelliSimpPromptTemplateId = &v
	return s
}

func (s *RunVideoDetectShotShrinkRequest) SetLanguage(v string) *RunVideoDetectShotShrinkRequest {
	s.Language = &v
	return s
}

func (s *RunVideoDetectShotShrinkRequest) SetModelCustomPromptTemplateId(v string) *RunVideoDetectShotShrinkRequest {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *RunVideoDetectShotShrinkRequest) SetModelId(v string) *RunVideoDetectShotShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunVideoDetectShotShrinkRequest) SetModelVlCustomPromptTemplateId(v string) *RunVideoDetectShotShrinkRequest {
	s.ModelVlCustomPromptTemplateId = &v
	return s
}

func (s *RunVideoDetectShotShrinkRequest) SetOptionsShrink(v string) *RunVideoDetectShotShrinkRequest {
	s.OptionsShrink = &v
	return s
}

func (s *RunVideoDetectShotShrinkRequest) SetOriginalSessionId(v string) *RunVideoDetectShotShrinkRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *RunVideoDetectShotShrinkRequest) SetPreModelId(v string) *RunVideoDetectShotShrinkRequest {
	s.PreModelId = &v
	return s
}

func (s *RunVideoDetectShotShrinkRequest) SetPrompt(v string) *RunVideoDetectShotShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunVideoDetectShotShrinkRequest) SetRecognitionOptionsShrink(v string) *RunVideoDetectShotShrinkRequest {
	s.RecognitionOptionsShrink = &v
	return s
}

func (s *RunVideoDetectShotShrinkRequest) SetTaskId(v string) *RunVideoDetectShotShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunVideoDetectShotShrinkRequest) SetVideoUrl(v string) *RunVideoDetectShotShrinkRequest {
	s.VideoUrl = &v
	return s
}

func (s *RunVideoDetectShotShrinkRequest) SetVlPrompt(v string) *RunVideoDetectShotShrinkRequest {
	s.VlPrompt = &v
	return s
}

func (s *RunVideoDetectShotShrinkRequest) Validate() error {
	return dara.Validate(s)
}
