// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoDetectShotTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeduplicationId(v string) *SubmitVideoDetectShotTaskRequest
	GetDeduplicationId() *string
	SetIntelliSimpPrompt(v string) *SubmitVideoDetectShotTaskRequest
	GetIntelliSimpPrompt() *string
	SetIntelliSimpPromptTemplateId(v string) *SubmitVideoDetectShotTaskRequest
	GetIntelliSimpPromptTemplateId() *string
	SetLanguage(v string) *SubmitVideoDetectShotTaskRequest
	GetLanguage() *string
	SetModelCustomPromptTemplateId(v string) *SubmitVideoDetectShotTaskRequest
	GetModelCustomPromptTemplateId() *string
	SetModelId(v string) *SubmitVideoDetectShotTaskRequest
	GetModelId() *string
	SetModelVlCustomPromptTemplateId(v string) *SubmitVideoDetectShotTaskRequest
	GetModelVlCustomPromptTemplateId() *string
	SetOptions(v []*string) *SubmitVideoDetectShotTaskRequest
	GetOptions() []*string
	SetOriginalSessionId(v string) *SubmitVideoDetectShotTaskRequest
	GetOriginalSessionId() *string
	SetPreModelId(v string) *SubmitVideoDetectShotTaskRequest
	GetPreModelId() *string
	SetPrompt(v string) *SubmitVideoDetectShotTaskRequest
	GetPrompt() *string
	SetRecognitionOptions(v []*string) *SubmitVideoDetectShotTaskRequest
	GetRecognitionOptions() []*string
	SetTaskId(v string) *SubmitVideoDetectShotTaskRequest
	GetTaskId() *string
	SetVideoUrl(v string) *SubmitVideoDetectShotTaskRequest
	GetVideoUrl() *string
	SetVlPrompt(v string) *SubmitVideoDetectShotTaskRequest
	GetVlPrompt() *string
}

type SubmitVideoDetectShotTaskRequest struct {
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
	ModelVlCustomPromptTemplateId *string   `json:"modelVlCustomPromptTemplateId,omitempty" xml:"modelVlCustomPromptTemplateId,omitempty"`
	Options                       []*string `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
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
	Prompt             *string   `json:"prompt,omitempty" xml:"prompt,omitempty"`
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

func (s SubmitVideoDetectShotTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoDetectShotTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoDetectShotTaskRequest) GetDeduplicationId() *string {
	return s.DeduplicationId
}

func (s *SubmitVideoDetectShotTaskRequest) GetIntelliSimpPrompt() *string {
	return s.IntelliSimpPrompt
}

func (s *SubmitVideoDetectShotTaskRequest) GetIntelliSimpPromptTemplateId() *string {
	return s.IntelliSimpPromptTemplateId
}

func (s *SubmitVideoDetectShotTaskRequest) GetLanguage() *string {
	return s.Language
}

func (s *SubmitVideoDetectShotTaskRequest) GetModelCustomPromptTemplateId() *string {
	return s.ModelCustomPromptTemplateId
}

func (s *SubmitVideoDetectShotTaskRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitVideoDetectShotTaskRequest) GetModelVlCustomPromptTemplateId() *string {
	return s.ModelVlCustomPromptTemplateId
}

func (s *SubmitVideoDetectShotTaskRequest) GetOptions() []*string {
	return s.Options
}

func (s *SubmitVideoDetectShotTaskRequest) GetOriginalSessionId() *string {
	return s.OriginalSessionId
}

func (s *SubmitVideoDetectShotTaskRequest) GetPreModelId() *string {
	return s.PreModelId
}

func (s *SubmitVideoDetectShotTaskRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *SubmitVideoDetectShotTaskRequest) GetRecognitionOptions() []*string {
	return s.RecognitionOptions
}

func (s *SubmitVideoDetectShotTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitVideoDetectShotTaskRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *SubmitVideoDetectShotTaskRequest) GetVlPrompt() *string {
	return s.VlPrompt
}

func (s *SubmitVideoDetectShotTaskRequest) SetDeduplicationId(v string) *SubmitVideoDetectShotTaskRequest {
	s.DeduplicationId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) SetIntelliSimpPrompt(v string) *SubmitVideoDetectShotTaskRequest {
	s.IntelliSimpPrompt = &v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) SetIntelliSimpPromptTemplateId(v string) *SubmitVideoDetectShotTaskRequest {
	s.IntelliSimpPromptTemplateId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) SetLanguage(v string) *SubmitVideoDetectShotTaskRequest {
	s.Language = &v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) SetModelCustomPromptTemplateId(v string) *SubmitVideoDetectShotTaskRequest {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) SetModelId(v string) *SubmitVideoDetectShotTaskRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) SetModelVlCustomPromptTemplateId(v string) *SubmitVideoDetectShotTaskRequest {
	s.ModelVlCustomPromptTemplateId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) SetOptions(v []*string) *SubmitVideoDetectShotTaskRequest {
	s.Options = v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) SetOriginalSessionId(v string) *SubmitVideoDetectShotTaskRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) SetPreModelId(v string) *SubmitVideoDetectShotTaskRequest {
	s.PreModelId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) SetPrompt(v string) *SubmitVideoDetectShotTaskRequest {
	s.Prompt = &v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) SetRecognitionOptions(v []*string) *SubmitVideoDetectShotTaskRequest {
	s.RecognitionOptions = v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) SetTaskId(v string) *SubmitVideoDetectShotTaskRequest {
	s.TaskId = &v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) SetVideoUrl(v string) *SubmitVideoDetectShotTaskRequest {
	s.VideoUrl = &v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) SetVlPrompt(v string) *SubmitVideoDetectShotTaskRequest {
	s.VlPrompt = &v
	return s
}

func (s *SubmitVideoDetectShotTaskRequest) Validate() error {
	return dara.Validate(s)
}
