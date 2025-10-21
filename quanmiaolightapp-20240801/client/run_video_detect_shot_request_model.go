// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunVideoDetectShotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIntelliSimpPrompt(v string) *RunVideoDetectShotRequest
	GetIntelliSimpPrompt() *string
	SetIntelliSimpPromptTemplateId(v string) *RunVideoDetectShotRequest
	GetIntelliSimpPromptTemplateId() *string
	SetLanguage(v string) *RunVideoDetectShotRequest
	GetLanguage() *string
	SetModelCustomPromptTemplateId(v string) *RunVideoDetectShotRequest
	GetModelCustomPromptTemplateId() *string
	SetModelId(v string) *RunVideoDetectShotRequest
	GetModelId() *string
	SetModelVlCustomPromptTemplateId(v string) *RunVideoDetectShotRequest
	GetModelVlCustomPromptTemplateId() *string
	SetOptions(v []*string) *RunVideoDetectShotRequest
	GetOptions() []*string
	SetOriginalSessionId(v string) *RunVideoDetectShotRequest
	GetOriginalSessionId() *string
	SetPreModelId(v string) *RunVideoDetectShotRequest
	GetPreModelId() *string
	SetPrompt(v string) *RunVideoDetectShotRequest
	GetPrompt() *string
	SetRecognitionOptions(v []*string) *RunVideoDetectShotRequest
	GetRecognitionOptions() []*string
	SetTaskId(v string) *RunVideoDetectShotRequest
	GetTaskId() *string
	SetVideoUrl(v string) *RunVideoDetectShotRequest
	GetVideoUrl() *string
	SetVlPrompt(v string) *RunVideoDetectShotRequest
	GetVlPrompt() *string
}

type RunVideoDetectShotRequest struct {
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
	Options []*string `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
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
	RecognitionOptions []*string `json:"recognitionOptions,omitempty" xml:"recognitionOptions,omitempty" type:"Repeated"`
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

func (s RunVideoDetectShotRequest) String() string {
	return dara.Prettify(s)
}

func (s RunVideoDetectShotRequest) GoString() string {
	return s.String()
}

func (s *RunVideoDetectShotRequest) GetIntelliSimpPrompt() *string {
	return s.IntelliSimpPrompt
}

func (s *RunVideoDetectShotRequest) GetIntelliSimpPromptTemplateId() *string {
	return s.IntelliSimpPromptTemplateId
}

func (s *RunVideoDetectShotRequest) GetLanguage() *string {
	return s.Language
}

func (s *RunVideoDetectShotRequest) GetModelCustomPromptTemplateId() *string {
	return s.ModelCustomPromptTemplateId
}

func (s *RunVideoDetectShotRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunVideoDetectShotRequest) GetModelVlCustomPromptTemplateId() *string {
	return s.ModelVlCustomPromptTemplateId
}

func (s *RunVideoDetectShotRequest) GetOptions() []*string {
	return s.Options
}

func (s *RunVideoDetectShotRequest) GetOriginalSessionId() *string {
	return s.OriginalSessionId
}

func (s *RunVideoDetectShotRequest) GetPreModelId() *string {
	return s.PreModelId
}

func (s *RunVideoDetectShotRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunVideoDetectShotRequest) GetRecognitionOptions() []*string {
	return s.RecognitionOptions
}

func (s *RunVideoDetectShotRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunVideoDetectShotRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *RunVideoDetectShotRequest) GetVlPrompt() *string {
	return s.VlPrompt
}

func (s *RunVideoDetectShotRequest) SetIntelliSimpPrompt(v string) *RunVideoDetectShotRequest {
	s.IntelliSimpPrompt = &v
	return s
}

func (s *RunVideoDetectShotRequest) SetIntelliSimpPromptTemplateId(v string) *RunVideoDetectShotRequest {
	s.IntelliSimpPromptTemplateId = &v
	return s
}

func (s *RunVideoDetectShotRequest) SetLanguage(v string) *RunVideoDetectShotRequest {
	s.Language = &v
	return s
}

func (s *RunVideoDetectShotRequest) SetModelCustomPromptTemplateId(v string) *RunVideoDetectShotRequest {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *RunVideoDetectShotRequest) SetModelId(v string) *RunVideoDetectShotRequest {
	s.ModelId = &v
	return s
}

func (s *RunVideoDetectShotRequest) SetModelVlCustomPromptTemplateId(v string) *RunVideoDetectShotRequest {
	s.ModelVlCustomPromptTemplateId = &v
	return s
}

func (s *RunVideoDetectShotRequest) SetOptions(v []*string) *RunVideoDetectShotRequest {
	s.Options = v
	return s
}

func (s *RunVideoDetectShotRequest) SetOriginalSessionId(v string) *RunVideoDetectShotRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *RunVideoDetectShotRequest) SetPreModelId(v string) *RunVideoDetectShotRequest {
	s.PreModelId = &v
	return s
}

func (s *RunVideoDetectShotRequest) SetPrompt(v string) *RunVideoDetectShotRequest {
	s.Prompt = &v
	return s
}

func (s *RunVideoDetectShotRequest) SetRecognitionOptions(v []*string) *RunVideoDetectShotRequest {
	s.RecognitionOptions = v
	return s
}

func (s *RunVideoDetectShotRequest) SetTaskId(v string) *RunVideoDetectShotRequest {
	s.TaskId = &v
	return s
}

func (s *RunVideoDetectShotRequest) SetVideoUrl(v string) *RunVideoDetectShotRequest {
	s.VideoUrl = &v
	return s
}

func (s *RunVideoDetectShotRequest) SetVlPrompt(v string) *RunVideoDetectShotRequest {
	s.VlPrompt = &v
	return s
}

func (s *RunVideoDetectShotRequest) Validate() error {
	return dara.Validate(s)
}
