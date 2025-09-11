// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTranslateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtShrink(v string) *TextTranslateShrinkRequest
	GetExtShrink() *string
	SetFormat(v string) *TextTranslateShrinkRequest
	GetFormat() *string
	SetScene(v string) *TextTranslateShrinkRequest
	GetScene() *string
	SetSourceLanguage(v string) *TextTranslateShrinkRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *TextTranslateShrinkRequest
	GetTargetLanguage() *string
	SetText(v string) *TextTranslateShrinkRequest
	GetText() *string
	SetWorkspaceId(v string) *TextTranslateShrinkRequest
	GetWorkspaceId() *string
}

type TextTranslateShrinkRequest struct {
	ExtShrink *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// example:
	//
	// text
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// mt-turbo
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"targetLanguage,omitempty" xml:"targetLanguage,omitempty"`
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s TextTranslateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateShrinkRequest) GoString() string {
	return s.String()
}

func (s *TextTranslateShrinkRequest) GetExtShrink() *string {
	return s.ExtShrink
}

func (s *TextTranslateShrinkRequest) GetFormat() *string {
	return s.Format
}

func (s *TextTranslateShrinkRequest) GetScene() *string {
	return s.Scene
}

func (s *TextTranslateShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TextTranslateShrinkRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TextTranslateShrinkRequest) GetText() *string {
	return s.Text
}

func (s *TextTranslateShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *TextTranslateShrinkRequest) SetExtShrink(v string) *TextTranslateShrinkRequest {
	s.ExtShrink = &v
	return s
}

func (s *TextTranslateShrinkRequest) SetFormat(v string) *TextTranslateShrinkRequest {
	s.Format = &v
	return s
}

func (s *TextTranslateShrinkRequest) SetScene(v string) *TextTranslateShrinkRequest {
	s.Scene = &v
	return s
}

func (s *TextTranslateShrinkRequest) SetSourceLanguage(v string) *TextTranslateShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TextTranslateShrinkRequest) SetTargetLanguage(v string) *TextTranslateShrinkRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TextTranslateShrinkRequest) SetText(v string) *TextTranslateShrinkRequest {
	s.Text = &v
	return s
}

func (s *TextTranslateShrinkRequest) SetWorkspaceId(v string) *TextTranslateShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *TextTranslateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
