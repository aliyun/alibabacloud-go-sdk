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
	// The error related to the target URL.
	//
	// example:
	//
	// null
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The readable HTML of the target URL.
	//
	// example:
	//
	// <html>\\n<head><title>Example Domain</title></head>\\n<body><div>\\n<h1>Example Domain</h1>\\n<p>This domain is for use in documentation examples without needing permission. Avoid use in operations.</p>\\n<p><a href=\\"https://iana.org/domains/example\\">Learn more</a></p>\\n</div></body>\\n</html>
	Html *string `json:"html,omitempty" xml:"html,omitempty"`
	// The Markdown content of the target URL.
	//
	// example:
	//
	// # Example Domain\\nThis domain is for use in documentation examples without needing permission. Avoid use in operations.\\n[Learn more](https://iana.org/domains/example)\\n
	Markdown *string `json:"markdown,omitempty" xml:"markdown,omitempty"`
	// The raw HTML of the target URL.
	//
	// example:
	//
	// <!doctype html><html lang=\\"en\\"><head><title>Example Domain</title><meta name=\\"viewport\\" content=\\"width=device-width, initial-scale=1\\"><style>body{background:#eee;width:60vw;margin:15vh auto;font-family:system-ui,sans-serif}h1{font-size:1.5em}div{opacity:0.8}a:link,a:visited{color:#348}</style><body><div><h1>Example Domain</h1><p>This domain is for use in documentation examples without needing permission. Avoid use in operations.<p><a href=\\"https://iana.org/domains/example\\">Learn more</a></div></body></html>\\n
	RawHtml    *string `json:"rawHtml,omitempty" xml:"rawHtml,omitempty"`
	Screenshot *string `json:"screenshot,omitempty" xml:"screenshot,omitempty"`
	// 1. If the request to the target site succeeds, the HTTP status code of the target URL is returned.
	//
	// 2. If the request to the target site fails, a custom error code is returned:
	//
	//     2.1 4030: The target site has security restrictions, such as robots.txt or security policies.
	//
	//
	//
	//     2.2 4080: The request timed out.
	//
	//
	//
	//     2.3 4290: The rate limiting policy of the site was triggered.
	//
	//
	//
	//     2.4 5010: An unknown exception occurred.
	//
	// example:
	//
	// 200
	StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	// The text content of the target URL.
	//
	// example:
	//
	// # Example Domain\\nThis domain is for use in documentation examples without needing permission. Avoid use in operations.\\nLearn more\\n
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
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
