// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocTranslateTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtShrink(v string) *SubmitDocTranslateTaskShrinkRequest
	GetExtShrink() *string
	SetFormat(v string) *SubmitDocTranslateTaskShrinkRequest
	GetFormat() *string
	SetScene(v string) *SubmitDocTranslateTaskShrinkRequest
	GetScene() *string
	SetSourceLanguage(v string) *SubmitDocTranslateTaskShrinkRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *SubmitDocTranslateTaskShrinkRequest
	GetTargetLanguage() *string
	SetText(v string) *SubmitDocTranslateTaskShrinkRequest
	GetText() *string
	SetWorkspaceId(v string) *SubmitDocTranslateTaskShrinkRequest
	GetWorkspaceId() *string
}

type SubmitDocTranslateTaskShrinkRequest struct {
	// Extension parameters that control translation features.
	ExtShrink *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// The format for the translation.
	//
	// example:
	//
	// text
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// The translation model.
	//
	// This parameter is required.
	//
	// example:
	//
	// mt-turbo
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The source language code.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// The target language code.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"targetLanguage,omitempty" xml:"targetLanguage,omitempty"`
	// The URL of the document to translate.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxx-hangzhou.aliyuncs.com/docs/tmp/%E6%A0%B7%E4%BE%8B_%E6%97%A0%E5%9B%BE.pdf
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// The ID of the Model Studio workspace for the current request.
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xc
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s SubmitDocTranslateTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocTranslateTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocTranslateTaskShrinkRequest) GetExtShrink() *string {
	return s.ExtShrink
}

func (s *SubmitDocTranslateTaskShrinkRequest) GetFormat() *string {
	return s.Format
}

func (s *SubmitDocTranslateTaskShrinkRequest) GetScene() *string {
	return s.Scene
}

func (s *SubmitDocTranslateTaskShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *SubmitDocTranslateTaskShrinkRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *SubmitDocTranslateTaskShrinkRequest) GetText() *string {
	return s.Text
}

func (s *SubmitDocTranslateTaskShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitDocTranslateTaskShrinkRequest) SetExtShrink(v string) *SubmitDocTranslateTaskShrinkRequest {
	s.ExtShrink = &v
	return s
}

func (s *SubmitDocTranslateTaskShrinkRequest) SetFormat(v string) *SubmitDocTranslateTaskShrinkRequest {
	s.Format = &v
	return s
}

func (s *SubmitDocTranslateTaskShrinkRequest) SetScene(v string) *SubmitDocTranslateTaskShrinkRequest {
	s.Scene = &v
	return s
}

func (s *SubmitDocTranslateTaskShrinkRequest) SetSourceLanguage(v string) *SubmitDocTranslateTaskShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *SubmitDocTranslateTaskShrinkRequest) SetTargetLanguage(v string) *SubmitDocTranslateTaskShrinkRequest {
	s.TargetLanguage = &v
	return s
}

func (s *SubmitDocTranslateTaskShrinkRequest) SetText(v string) *SubmitDocTranslateTaskShrinkRequest {
	s.Text = &v
	return s
}

func (s *SubmitDocTranslateTaskShrinkRequest) SetWorkspaceId(v string) *SubmitDocTranslateTaskShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitDocTranslateTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
