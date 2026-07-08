// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHtmlTranslateTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtShrink(v string) *SubmitHtmlTranslateTaskShrinkRequest
	GetExtShrink() *string
	SetFormat(v string) *SubmitHtmlTranslateTaskShrinkRequest
	GetFormat() *string
	SetScene(v string) *SubmitHtmlTranslateTaskShrinkRequest
	GetScene() *string
	SetSourceLanguage(v string) *SubmitHtmlTranslateTaskShrinkRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *SubmitHtmlTranslateTaskShrinkRequest
	GetTargetLanguage() *string
	SetText(v string) *SubmitHtmlTranslateTaskShrinkRequest
	GetText() *string
	SetWorkspaceId(v string) *SubmitHtmlTranslateTaskShrinkRequest
	GetWorkspaceId() *string
}

type SubmitHtmlTranslateTaskShrinkRequest struct {
	// Extended parameters to control translation behavior.
	ExtShrink *string `json:"ext,omitempty" xml:"ext,omitempty"`
	// The format of the input text.
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
	// example:
	//
	// zh
	SourceLanguage *string `json:"sourceLanguage,omitempty" xml:"sourceLanguage,omitempty"`
	// The target language.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"targetLanguage,omitempty" xml:"targetLanguage,omitempty"`
	// The HTML content to be translated.
	//
	// example:
	//
	// <!DOCTYPE html>
	//
	// <html lang="zh-CN">
	//
	// <head>
	//
	//     <meta charset="UTF-8">
	//
	//     <meta name="viewport" content="width=device-width, initial-scale=1.0">
	//
	//     <title>我的第一个网页</title>
	//
	// </head>
	//
	// <body>
	//
	//     <h1>欢迎来到我的网页！</h1>
	//
	//     <p>这是一个简单的 HTML 页面。</p>
	//
	//     <p>学习 HTML 是进入网页开发的第一步。</p>
	//
	//     <a href="https://www.example.com">点击这里访问示例网站</a>
	//
	// </body>
	//
	// </html>
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// The ID of the Model Studio workspace used for this request.
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-kqtrcpdee4xm29xx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s SubmitHtmlTranslateTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitHtmlTranslateTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) GetExtShrink() *string {
	return s.ExtShrink
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) GetFormat() *string {
	return s.Format
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) GetScene() *string {
	return s.Scene
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) GetText() *string {
	return s.Text
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) SetExtShrink(v string) *SubmitHtmlTranslateTaskShrinkRequest {
	s.ExtShrink = &v
	return s
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) SetFormat(v string) *SubmitHtmlTranslateTaskShrinkRequest {
	s.Format = &v
	return s
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) SetScene(v string) *SubmitHtmlTranslateTaskShrinkRequest {
	s.Scene = &v
	return s
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) SetSourceLanguage(v string) *SubmitHtmlTranslateTaskShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) SetTargetLanguage(v string) *SubmitHtmlTranslateTaskShrinkRequest {
	s.TargetLanguage = &v
	return s
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) SetText(v string) *SubmitHtmlTranslateTaskShrinkRequest {
	s.Text = &v
	return s
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) SetWorkspaceId(v string) *SubmitHtmlTranslateTaskShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitHtmlTranslateTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
