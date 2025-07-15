// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchNewsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchNewsResponseBody
	GetCode() *string
	SetCurrent(v int32) *SearchNewsResponseBody
	GetCurrent() *int32
	SetData(v []*SearchNewsResponseBodyData) *SearchNewsResponseBody
	GetData() []*SearchNewsResponseBodyData
	SetHttpStatusCode(v int32) *SearchNewsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SearchNewsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SearchNewsResponseBody
	GetRequestId() *string
	SetSize(v int32) *SearchNewsResponseBody
	GetSize() *int32
	SetSuccess(v bool) *SearchNewsResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *SearchNewsResponseBody
	GetTotal() *int32
}

type SearchNewsResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                        `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*SearchNewsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchNewsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchNewsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchNewsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchNewsResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *SearchNewsResponseBody) GetData() []*SearchNewsResponseBodyData {
	return s.Data
}

func (s *SearchNewsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SearchNewsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchNewsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchNewsResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *SearchNewsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchNewsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *SearchNewsResponseBody) SetCode(v string) *SearchNewsResponseBody {
	s.Code = &v
	return s
}

func (s *SearchNewsResponseBody) SetCurrent(v int32) *SearchNewsResponseBody {
	s.Current = &v
	return s
}

func (s *SearchNewsResponseBody) SetData(v []*SearchNewsResponseBodyData) *SearchNewsResponseBody {
	s.Data = v
	return s
}

func (s *SearchNewsResponseBody) SetHttpStatusCode(v int32) *SearchNewsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SearchNewsResponseBody) SetMessage(v string) *SearchNewsResponseBody {
	s.Message = &v
	return s
}

func (s *SearchNewsResponseBody) SetRequestId(v string) *SearchNewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchNewsResponseBody) SetSize(v int32) *SearchNewsResponseBody {
	s.Size = &v
	return s
}

func (s *SearchNewsResponseBody) SetSuccess(v bool) *SearchNewsResponseBody {
	s.Success = &v
	return s
}

func (s *SearchNewsResponseBody) SetTotal(v int32) *SearchNewsResponseBody {
	s.Total = &v
	return s
}

func (s *SearchNewsResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchNewsResponseBodyData struct {
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
	// 9a598b44c6444da5907b8ea68a5f82c4
	DocUuid   *string   `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	ImageUrls []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-01-18 06:46:22
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// QuarkCommonNews
	SearchSource *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	// example:
	//
	// 夸克检索
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
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
	// 2024-01-18 06:46:22
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 文章URL
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SearchNewsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchNewsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchNewsResponseBodyData) GetAuthor() *string {
	return s.Author
}

func (s *SearchNewsResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *SearchNewsResponseBodyData) GetDocUuid() *string {
	return s.DocUuid
}

func (s *SearchNewsResponseBodyData) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *SearchNewsResponseBodyData) GetPubTime() *string {
	return s.PubTime
}

func (s *SearchNewsResponseBodyData) GetSearchSource() *string {
	return s.SearchSource
}

func (s *SearchNewsResponseBodyData) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *SearchNewsResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *SearchNewsResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *SearchNewsResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *SearchNewsResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *SearchNewsResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SearchNewsResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *SearchNewsResponseBodyData) SetAuthor(v string) *SearchNewsResponseBodyData {
	s.Author = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetContent(v string) *SearchNewsResponseBodyData {
	s.Content = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetDocUuid(v string) *SearchNewsResponseBodyData {
	s.DocUuid = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetImageUrls(v []*string) *SearchNewsResponseBodyData {
	s.ImageUrls = v
	return s
}

func (s *SearchNewsResponseBodyData) SetPubTime(v string) *SearchNewsResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetSearchSource(v string) *SearchNewsResponseBodyData {
	s.SearchSource = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetSearchSourceName(v string) *SearchNewsResponseBodyData {
	s.SearchSourceName = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetSource(v string) *SearchNewsResponseBodyData {
	s.Source = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetSummary(v string) *SearchNewsResponseBodyData {
	s.Summary = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetTag(v string) *SearchNewsResponseBodyData {
	s.Tag = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetTitle(v string) *SearchNewsResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetUpdateTime(v string) *SearchNewsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetUrl(v string) *SearchNewsResponseBodyData {
	s.Url = &v
	return s
}

func (s *SearchNewsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
