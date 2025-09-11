// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchTranslateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtShrink(v string) *BatchTranslateShrinkRequest
	GetExtShrink() *string
	SetFormat(v string) *BatchTranslateShrinkRequest
	GetFormat() *string
	SetScene(v string) *BatchTranslateShrinkRequest
	GetScene() *string
	SetSourceLanguage(v string) *BatchTranslateShrinkRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *BatchTranslateShrinkRequest
	GetTargetLanguage() *string
	SetTextShrink(v string) *BatchTranslateShrinkRequest
	GetTextShrink() *string
	SetWorkspaceId(v string) *BatchTranslateShrinkRequest
	GetWorkspaceId() *string
}

type BatchTranslateShrinkRequest struct {
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
	TextShrink *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s BatchTranslateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchTranslateShrinkRequest) GetExtShrink() *string {
	return s.ExtShrink
}

func (s *BatchTranslateShrinkRequest) GetFormat() *string {
	return s.Format
}

func (s *BatchTranslateShrinkRequest) GetScene() *string {
	return s.Scene
}

func (s *BatchTranslateShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *BatchTranslateShrinkRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *BatchTranslateShrinkRequest) GetTextShrink() *string {
	return s.TextShrink
}

func (s *BatchTranslateShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *BatchTranslateShrinkRequest) SetExtShrink(v string) *BatchTranslateShrinkRequest {
	s.ExtShrink = &v
	return s
}

func (s *BatchTranslateShrinkRequest) SetFormat(v string) *BatchTranslateShrinkRequest {
	s.Format = &v
	return s
}

func (s *BatchTranslateShrinkRequest) SetScene(v string) *BatchTranslateShrinkRequest {
	s.Scene = &v
	return s
}

func (s *BatchTranslateShrinkRequest) SetSourceLanguage(v string) *BatchTranslateShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *BatchTranslateShrinkRequest) SetTargetLanguage(v string) *BatchTranslateShrinkRequest {
	s.TargetLanguage = &v
	return s
}

func (s *BatchTranslateShrinkRequest) SetTextShrink(v string) *BatchTranslateShrinkRequest {
	s.TextShrink = &v
	return s
}

func (s *BatchTranslateShrinkRequest) SetWorkspaceId(v string) *BatchTranslateShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchTranslateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
