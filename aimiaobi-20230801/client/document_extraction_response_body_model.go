// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentExtractionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DocumentExtractionResponseBody
	GetCode() *string
	SetData(v []*DocumentExtractionResponseBodyData) *DocumentExtractionResponseBody
	GetData() []*DocumentExtractionResponseBodyData
	SetHttpStatusCode(v int32) *DocumentExtractionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DocumentExtractionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DocumentExtractionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DocumentExtractionResponseBody
	GetSuccess() *bool
}

type DocumentExtractionResponseBody struct {
	// example:
	//
	// NoData
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*DocumentExtractionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DocumentExtractionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DocumentExtractionResponseBody) GoString() string {
	return s.String()
}

func (s *DocumentExtractionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DocumentExtractionResponseBody) GetData() []*DocumentExtractionResponseBodyData {
	return s.Data
}

func (s *DocumentExtractionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DocumentExtractionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DocumentExtractionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DocumentExtractionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DocumentExtractionResponseBody) SetCode(v string) *DocumentExtractionResponseBody {
	s.Code = &v
	return s
}

func (s *DocumentExtractionResponseBody) SetData(v []*DocumentExtractionResponseBodyData) *DocumentExtractionResponseBody {
	s.Data = v
	return s
}

func (s *DocumentExtractionResponseBody) SetHttpStatusCode(v int32) *DocumentExtractionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DocumentExtractionResponseBody) SetMessage(v string) *DocumentExtractionResponseBody {
	s.Message = &v
	return s
}

func (s *DocumentExtractionResponseBody) SetRequestId(v string) *DocumentExtractionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DocumentExtractionResponseBody) SetSuccess(v bool) *DocumentExtractionResponseBody {
	s.Success = &v
	return s
}

func (s *DocumentExtractionResponseBody) Validate() error {
	return dara.Validate(s)
}

type DocumentExtractionResponseBodyData struct {
	// example:
	//
	// 作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// example:
	//
	// 文章内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 文档-自定义的唯一ID
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 8df2d69d63a247b6b52ff455b2d426b6
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 2024-05-14 08:54:33
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 央视网
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 文章标签
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DocumentExtractionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DocumentExtractionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DocumentExtractionResponseBodyData) GetAuthor() *string {
	return s.Author
}

func (s *DocumentExtractionResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *DocumentExtractionResponseBodyData) GetDocId() *string {
	return s.DocId
}

func (s *DocumentExtractionResponseBodyData) GetDocUuid() *string {
	return s.DocUuid
}

func (s *DocumentExtractionResponseBodyData) GetPubTime() *string {
	return s.PubTime
}

func (s *DocumentExtractionResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *DocumentExtractionResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *DocumentExtractionResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *DocumentExtractionResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *DocumentExtractionResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *DocumentExtractionResponseBodyData) SetAuthor(v string) *DocumentExtractionResponseBodyData {
	s.Author = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetContent(v string) *DocumentExtractionResponseBodyData {
	s.Content = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetDocId(v string) *DocumentExtractionResponseBodyData {
	s.DocId = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetDocUuid(v string) *DocumentExtractionResponseBodyData {
	s.DocUuid = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetPubTime(v string) *DocumentExtractionResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetSource(v string) *DocumentExtractionResponseBodyData {
	s.Source = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetSummary(v string) *DocumentExtractionResponseBodyData {
	s.Summary = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetTag(v string) *DocumentExtractionResponseBodyData {
	s.Tag = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetTitle(v string) *DocumentExtractionResponseBodyData {
	s.Title = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetUrl(v string) *DocumentExtractionResponseBodyData {
	s.Url = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
