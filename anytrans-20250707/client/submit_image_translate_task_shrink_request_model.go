// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageTranslateTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtShrink(v string) *SubmitImageTranslateTaskShrinkRequest
	GetExtShrink() *string
	SetFormat(v string) *SubmitImageTranslateTaskShrinkRequest
	GetFormat() *string
	SetScene(v string) *SubmitImageTranslateTaskShrinkRequest
	GetScene() *string
	SetSourceLanguage(v string) *SubmitImageTranslateTaskShrinkRequest
	GetSourceLanguage() *string
	SetTargetLanguageShrink(v string) *SubmitImageTranslateTaskShrinkRequest
	GetTargetLanguageShrink() *string
	SetText(v string) *SubmitImageTranslateTaskShrinkRequest
	GetText() *string
	SetWorkspaceId(v string) *SubmitImageTranslateTaskShrinkRequest
	GetWorkspaceId() *string
}

type SubmitImageTranslateTaskShrinkRequest struct {
	ExtShrink *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// example:
	//
	// image
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// flash
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// This parameter is required.
	TargetLanguageShrink *string `json:"targetLanguage,omitempty" xml:"targetLanguage,omitempty"`
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s SubmitImageTranslateTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageTranslateTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitImageTranslateTaskShrinkRequest) GetExtShrink() *string {
	return s.ExtShrink
}

func (s *SubmitImageTranslateTaskShrinkRequest) GetFormat() *string {
	return s.Format
}

func (s *SubmitImageTranslateTaskShrinkRequest) GetScene() *string {
	return s.Scene
}

func (s *SubmitImageTranslateTaskShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *SubmitImageTranslateTaskShrinkRequest) GetTargetLanguageShrink() *string {
	return s.TargetLanguageShrink
}

func (s *SubmitImageTranslateTaskShrinkRequest) GetText() *string {
	return s.Text
}

func (s *SubmitImageTranslateTaskShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitImageTranslateTaskShrinkRequest) SetExtShrink(v string) *SubmitImageTranslateTaskShrinkRequest {
	s.ExtShrink = &v
	return s
}

func (s *SubmitImageTranslateTaskShrinkRequest) SetFormat(v string) *SubmitImageTranslateTaskShrinkRequest {
	s.Format = &v
	return s
}

func (s *SubmitImageTranslateTaskShrinkRequest) SetScene(v string) *SubmitImageTranslateTaskShrinkRequest {
	s.Scene = &v
	return s
}

func (s *SubmitImageTranslateTaskShrinkRequest) SetSourceLanguage(v string) *SubmitImageTranslateTaskShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *SubmitImageTranslateTaskShrinkRequest) SetTargetLanguageShrink(v string) *SubmitImageTranslateTaskShrinkRequest {
	s.TargetLanguageShrink = &v
	return s
}

func (s *SubmitImageTranslateTaskShrinkRequest) SetText(v string) *SubmitImageTranslateTaskShrinkRequest {
	s.Text = &v
	return s
}

func (s *SubmitImageTranslateTaskShrinkRequest) SetWorkspaceId(v string) *SubmitImageTranslateTaskShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitImageTranslateTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
