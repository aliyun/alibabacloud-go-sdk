// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchTranslateForHtmlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *BatchTranslateForHtmlShrinkRequest
	GetAppName() *string
	SetExtShrink(v string) *BatchTranslateForHtmlShrinkRequest
	GetExtShrink() *string
	SetFormat(v string) *BatchTranslateForHtmlShrinkRequest
	GetFormat() *string
	SetScene(v string) *BatchTranslateForHtmlShrinkRequest
	GetScene() *string
	SetSourceLanguage(v string) *BatchTranslateForHtmlShrinkRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *BatchTranslateForHtmlShrinkRequest
	GetTargetLanguage() *string
	SetTextShrink(v string) *BatchTranslateForHtmlShrinkRequest
	GetTextShrink() *string
	SetWorkspaceId(v string) *BatchTranslateForHtmlShrinkRequest
	GetWorkspaceId() *string
}

type BatchTranslateForHtmlShrinkRequest struct {
	// example:
	//
	// baidufanyi
	AppName   *string `json:"appName,omitempty" xml:"appName,omitempty"`
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

func (s BatchTranslateForHtmlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateForHtmlShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchTranslateForHtmlShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *BatchTranslateForHtmlShrinkRequest) GetExtShrink() *string {
	return s.ExtShrink
}

func (s *BatchTranslateForHtmlShrinkRequest) GetFormat() *string {
	return s.Format
}

func (s *BatchTranslateForHtmlShrinkRequest) GetScene() *string {
	return s.Scene
}

func (s *BatchTranslateForHtmlShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *BatchTranslateForHtmlShrinkRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *BatchTranslateForHtmlShrinkRequest) GetTextShrink() *string {
	return s.TextShrink
}

func (s *BatchTranslateForHtmlShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *BatchTranslateForHtmlShrinkRequest) SetAppName(v string) *BatchTranslateForHtmlShrinkRequest {
	s.AppName = &v
	return s
}

func (s *BatchTranslateForHtmlShrinkRequest) SetExtShrink(v string) *BatchTranslateForHtmlShrinkRequest {
	s.ExtShrink = &v
	return s
}

func (s *BatchTranslateForHtmlShrinkRequest) SetFormat(v string) *BatchTranslateForHtmlShrinkRequest {
	s.Format = &v
	return s
}

func (s *BatchTranslateForHtmlShrinkRequest) SetScene(v string) *BatchTranslateForHtmlShrinkRequest {
	s.Scene = &v
	return s
}

func (s *BatchTranslateForHtmlShrinkRequest) SetSourceLanguage(v string) *BatchTranslateForHtmlShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *BatchTranslateForHtmlShrinkRequest) SetTargetLanguage(v string) *BatchTranslateForHtmlShrinkRequest {
	s.TargetLanguage = &v
	return s
}

func (s *BatchTranslateForHtmlShrinkRequest) SetTextShrink(v string) *BatchTranslateForHtmlShrinkRequest {
	s.TextShrink = &v
	return s
}

func (s *BatchTranslateForHtmlShrinkRequest) SetWorkspaceId(v string) *BatchTranslateForHtmlShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchTranslateForHtmlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
