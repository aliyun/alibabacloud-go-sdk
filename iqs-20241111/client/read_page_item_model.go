// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadPageItem interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *ReadPageItem
	GetErrorMessage() *string
	SetHtml(v string) *ReadPageItem
	GetHtml() *string
	SetMarkdown(v string) *ReadPageItem
	GetMarkdown() *string
	SetRawHtml(v string) *ReadPageItem
	GetRawHtml() *string
	SetScreenshot(v string) *ReadPageItem
	GetScreenshot() *string
	SetStatusCode(v int32) *ReadPageItem
	GetStatusCode() *int32
	SetText(v string) *ReadPageItem
	GetText() *string
}

type ReadPageItem struct {
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Html         *string `json:"html,omitempty" xml:"html,omitempty"`
	Markdown     *string `json:"markdown,omitempty" xml:"markdown,omitempty"`
	RawHtml      *string `json:"rawHtml,omitempty" xml:"rawHtml,omitempty"`
	Screenshot   *string `json:"screenshot,omitempty" xml:"screenshot,omitempty"`
	StatusCode   *int32  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Text         *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ReadPageItem) String() string {
	return dara.Prettify(s)
}

func (s ReadPageItem) GoString() string {
	return s.String()
}

func (s *ReadPageItem) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ReadPageItem) GetHtml() *string {
	return s.Html
}

func (s *ReadPageItem) GetMarkdown() *string {
	return s.Markdown
}

func (s *ReadPageItem) GetRawHtml() *string {
	return s.RawHtml
}

func (s *ReadPageItem) GetScreenshot() *string {
	return s.Screenshot
}

func (s *ReadPageItem) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadPageItem) GetText() *string {
	return s.Text
}

func (s *ReadPageItem) SetErrorMessage(v string) *ReadPageItem {
	s.ErrorMessage = &v
	return s
}

func (s *ReadPageItem) SetHtml(v string) *ReadPageItem {
	s.Html = &v
	return s
}

func (s *ReadPageItem) SetMarkdown(v string) *ReadPageItem {
	s.Markdown = &v
	return s
}

func (s *ReadPageItem) SetRawHtml(v string) *ReadPageItem {
	s.RawHtml = &v
	return s
}

func (s *ReadPageItem) SetScreenshot(v string) *ReadPageItem {
	s.Screenshot = &v
	return s
}

func (s *ReadPageItem) SetStatusCode(v int32) *ReadPageItem {
	s.StatusCode = &v
	return s
}

func (s *ReadPageItem) SetText(v string) *ReadPageItem {
	s.Text = &v
	return s
}

func (s *ReadPageItem) Validate() error {
	return dara.Validate(s)
}
