// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchTranslateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *BatchTranslateShrinkRequest
	GetAppName() *string
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
	// The name of the calling application.
	//
	// example:
	//
	// baidufanyi
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// The extended parameters that control translation features.
	ExtShrink *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// The translation format.
	//
	// example:
	//
	// text
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// The translation model.
	//
	// example:
	//
	// mt-turbo
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The source language.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// The target language.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"targetLanguage,omitempty" xml:"targetLanguage,omitempty"`
	// A map of texts to translate, in which the key is a custom identifier and the value is the source text.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"0":"明天天气怎么样？","1":"你中午吃饭了吗"}
	TextShrink *string `json:"text,omitempty" xml:"text,omitempty"`
	// The ID of the Model Studio workspace used for this request.
	//
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

func (s *BatchTranslateShrinkRequest) GetAppName() *string {
	return s.AppName
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

func (s *BatchTranslateShrinkRequest) SetAppName(v string) *BatchTranslateShrinkRequest {
	s.AppName = &v
	return s
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
