// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoDetectShotTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeduplicationId(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetDeduplicationId() *string
	SetIntelliSimpPrompt(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetIntelliSimpPrompt() *string
	SetIntelliSimpPromptTemplateId(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetIntelliSimpPromptTemplateId() *string
	SetLanguage(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetLanguage() *string
	SetModelCustomPromptTemplateId(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetModelCustomPromptTemplateId() *string
	SetModelId(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetModelId() *string
	SetModelVlCustomPromptTemplateId(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetModelVlCustomPromptTemplateId() *string
	SetOptionsShrink(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetOptionsShrink() *string
	SetOriginalSessionId(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetOriginalSessionId() *string
	SetPreModelId(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetPreModelId() *string
	SetPrompt(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetPrompt() *string
	SetRecognitionOptionsShrink(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetRecognitionOptionsShrink() *string
	SetTaskId(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetTaskId() *string
	SetVideoUrl(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetVideoUrl() *string
	SetVlPrompt(v string) *SubmitVideoDetectShotTaskShrinkRequest
	GetVlPrompt() *string
}

type SubmitVideoDetectShotTaskShrinkRequest struct {
	// example:
	//
	// 1
	DeduplicationId *string `json:"deduplicationId,omitempty" xml:"deduplicationId,omitempty"`
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
	OptionsShrink                 *string `json:"options,omitempty" xml:"options,omitempty"`
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
	Prompt                   *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
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

func (s SubmitVideoDetectShotTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoDetectShotTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetDeduplicationId() *string {
	return s.DeduplicationId
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetIntelliSimpPrompt() *string {
	return s.IntelliSimpPrompt
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetIntelliSimpPromptTemplateId() *string {
	return s.IntelliSimpPromptTemplateId
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetModelCustomPromptTemplateId() *string {
	return s.ModelCustomPromptTemplateId
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetModelVlCustomPromptTemplateId() *string {
	return s.ModelVlCustomPromptTemplateId
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetOptionsShrink() *string {
	return s.OptionsShrink
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetOriginalSessionId() *string {
	return s.OriginalSessionId
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetPreModelId() *string {
	return s.PreModelId
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetRecognitionOptionsShrink() *string {
	return s.RecognitionOptionsShrink
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) GetVlPrompt() *string {
	return s.VlPrompt
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetDeduplicationId(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.DeduplicationId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetIntelliSimpPrompt(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.IntelliSimpPrompt = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetIntelliSimpPromptTemplateId(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.IntelliSimpPromptTemplateId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetLanguage(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.Language = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetModelCustomPromptTemplateId(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetModelId(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetModelVlCustomPromptTemplateId(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.ModelVlCustomPromptTemplateId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetOptionsShrink(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.OptionsShrink = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetOriginalSessionId(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetPreModelId(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.PreModelId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetPrompt(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetRecognitionOptionsShrink(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.RecognitionOptionsShrink = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetTaskId(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetVideoUrl(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.VideoUrl = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) SetVlPrompt(v string) *SubmitVideoDetectShotTaskShrinkRequest {
	s.VlPrompt = &v
	return s
}

func (s *SubmitVideoDetectShotTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
