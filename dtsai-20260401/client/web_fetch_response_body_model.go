// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebFetchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *WebFetchResponseBody
	GetContent() *string
	SetContentFormat(v string) *WebFetchResponseBody
	GetContentFormat() *string
	SetErrorCode(v string) *WebFetchResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *WebFetchResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *WebFetchResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *WebFetchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *WebFetchResponseBody
	GetSuccess() *bool
	SetTitle(v string) *WebFetchResponseBody
	GetTitle() *string
	SetUrl(v string) *WebFetchResponseBody
	GetUrl() *string
	SetUrlType(v string) *WebFetchResponseBody
	GetUrlType() *string
}

type WebFetchResponseBody struct {
	// The content of the web page.
	//
	// example:
	//
	// Spring Boot makes it easy to create stand-alone, production-grade Spring based Applications.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The actual output format. Valid values:
	//
	// - **markdown**: Markdown format.
	//
	// - **html**: HTML format.
	//
	// - **text**: Plain text format.
	//
	// example:
	//
	// markdown
	ContentFormat *string `json:"ContentFormat,omitempty" xml:"ContentFormat,omitempty"`
	// The error code.
	//
	// example:
	//
	// ErrorMessage Code
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// ErrorMessage Example
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A30D0930-xxxx-xxxx-xxxx-C2C661CC8B58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The title of the web page.
	//
	// example:
	//
	// Spring Boot
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of the web page.
	//
	// example:
	//
	// https://xxxx/projects/spring-boot
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The URL classification result. Valid values:
	//
	// - **static_html**: Static HTML page.
	//
	// - **dynamic_js**: Single-page application (SPA) page that requires JavaScript rendering.
	//
	// - **audio**: Audio file.
	//
	// - **video**: Video file.
	//
	// - **document**: Document file (PDF, Word, PPT, or Excel).
	//
	// - **unknown**: Unrecognizable.
	//
	// example:
	//
	// static_html
	UrlType *string `json:"UrlType,omitempty" xml:"UrlType,omitempty"`
}

func (s WebFetchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WebFetchResponseBody) GoString() string {
	return s.String()
}

func (s *WebFetchResponseBody) GetContent() *string {
	return s.Content
}

func (s *WebFetchResponseBody) GetContentFormat() *string {
	return s.ContentFormat
}

func (s *WebFetchResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *WebFetchResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *WebFetchResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *WebFetchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebFetchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *WebFetchResponseBody) GetTitle() *string {
	return s.Title
}

func (s *WebFetchResponseBody) GetUrl() *string {
	return s.Url
}

func (s *WebFetchResponseBody) GetUrlType() *string {
	return s.UrlType
}

func (s *WebFetchResponseBody) SetContent(v string) *WebFetchResponseBody {
	s.Content = &v
	return s
}

func (s *WebFetchResponseBody) SetContentFormat(v string) *WebFetchResponseBody {
	s.ContentFormat = &v
	return s
}

func (s *WebFetchResponseBody) SetErrorCode(v string) *WebFetchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *WebFetchResponseBody) SetErrorMessage(v string) *WebFetchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *WebFetchResponseBody) SetHttpStatusCode(v int32) *WebFetchResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *WebFetchResponseBody) SetRequestId(v string) *WebFetchResponseBody {
	s.RequestId = &v
	return s
}

func (s *WebFetchResponseBody) SetSuccess(v bool) *WebFetchResponseBody {
	s.Success = &v
	return s
}

func (s *WebFetchResponseBody) SetTitle(v string) *WebFetchResponseBody {
	s.Title = &v
	return s
}

func (s *WebFetchResponseBody) SetUrl(v string) *WebFetchResponseBody {
	s.Url = &v
	return s
}

func (s *WebFetchResponseBody) SetUrlType(v string) *WebFetchResponseBody {
	s.UrlType = &v
	return s
}

func (s *WebFetchResponseBody) Validate() error {
	return dara.Validate(s)
}
