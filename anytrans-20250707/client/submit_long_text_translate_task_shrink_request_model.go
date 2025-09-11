// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLongTextTranslateTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtShrink(v string) *SubmitLongTextTranslateTaskShrinkRequest
	GetExtShrink() *string
	SetFormat(v string) *SubmitLongTextTranslateTaskShrinkRequest
	GetFormat() *string
	SetScene(v string) *SubmitLongTextTranslateTaskShrinkRequest
	GetScene() *string
	SetSourceLanguage(v string) *SubmitLongTextTranslateTaskShrinkRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *SubmitLongTextTranslateTaskShrinkRequest
	GetTargetLanguage() *string
	SetText(v string) *SubmitLongTextTranslateTaskShrinkRequest
	GetText() *string
	SetWorkspaceId(v string) *SubmitLongTextTranslateTaskShrinkRequest
	GetWorkspaceId() *string
}

type SubmitLongTextTranslateTaskShrinkRequest struct {
	ExtShrink *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// example:
	//
	// text
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// mt-turbo
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// zh
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// example:
	//
	// en
	TargetLanguage *string `json:"targetLanguage,omitempty" xml:"targetLanguage,omitempty"`
	Text           *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s SubmitLongTextTranslateTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitLongTextTranslateTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) GetExtShrink() *string {
	return s.ExtShrink
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) GetFormat() *string {
	return s.Format
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) GetScene() *string {
	return s.Scene
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) GetText() *string {
	return s.Text
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) SetExtShrink(v string) *SubmitLongTextTranslateTaskShrinkRequest {
	s.ExtShrink = &v
	return s
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) SetFormat(v string) *SubmitLongTextTranslateTaskShrinkRequest {
	s.Format = &v
	return s
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) SetScene(v string) *SubmitLongTextTranslateTaskShrinkRequest {
	s.Scene = &v
	return s
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) SetSourceLanguage(v string) *SubmitLongTextTranslateTaskShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) SetTargetLanguage(v string) *SubmitLongTextTranslateTaskShrinkRequest {
	s.TargetLanguage = &v
	return s
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) SetText(v string) *SubmitLongTextTranslateTaskShrinkRequest {
	s.Text = &v
	return s
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) SetWorkspaceId(v string) *SubmitLongTextTranslateTaskShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitLongTextTranslateTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
