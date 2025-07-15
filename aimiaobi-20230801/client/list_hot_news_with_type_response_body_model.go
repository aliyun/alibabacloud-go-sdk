// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotNewsWithTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHotNewsWithTypeResponseBody
	GetCode() *string
	SetData(v []*ListHotNewsWithTypeResponseBodyData) *ListHotNewsWithTypeResponseBody
	GetData() []*ListHotNewsWithTypeResponseBodyData
	SetHttpStatusCode(v int32) *ListHotNewsWithTypeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListHotNewsWithTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHotNewsWithTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHotNewsWithTypeResponseBody
	GetSuccess() *bool
}

type ListHotNewsWithTypeResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListHotNewsWithTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListHotNewsWithTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotNewsWithTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHotNewsWithTypeResponseBody) GetData() []*ListHotNewsWithTypeResponseBodyData {
	return s.Data
}

func (s *ListHotNewsWithTypeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHotNewsWithTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotNewsWithTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotNewsWithTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHotNewsWithTypeResponseBody) SetCode(v string) *ListHotNewsWithTypeResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetData(v []*ListHotNewsWithTypeResponseBodyData) *ListHotNewsWithTypeResponseBody {
	s.Data = v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetHttpStatusCode(v int32) *ListHotNewsWithTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetMessage(v string) *ListHotNewsWithTypeResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetRequestId(v string) *ListHotNewsWithTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetSuccess(v bool) *ListHotNewsWithTypeResponseBody {
	s.Success = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHotNewsWithTypeResponseBodyData struct {
	News []*ListHotNewsWithTypeResponseBodyDataNews `json:"News,omitempty" xml:"News,omitempty" type:"Repeated"`
	// example:
	//
	// society
	NewsType     *string `json:"NewsType,omitempty" xml:"NewsType,omitempty"`
	NewsTypeName *string `json:"NewsTypeName,omitempty" xml:"NewsTypeName,omitempty"`
	// example:
	//
	// 77
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListHotNewsWithTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHotNewsWithTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeResponseBodyData) GetNews() []*ListHotNewsWithTypeResponseBodyDataNews {
	return s.News
}

func (s *ListHotNewsWithTypeResponseBodyData) GetNewsType() *string {
	return s.NewsType
}

func (s *ListHotNewsWithTypeResponseBodyData) GetNewsTypeName() *string {
	return s.NewsTypeName
}

func (s *ListHotNewsWithTypeResponseBodyData) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListHotNewsWithTypeResponseBodyData) SetNews(v []*ListHotNewsWithTypeResponseBodyDataNews) *ListHotNewsWithTypeResponseBodyData {
	s.News = v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyData) SetNewsType(v string) *ListHotNewsWithTypeResponseBodyData {
	s.NewsType = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyData) SetNewsTypeName(v string) *ListHotNewsWithTypeResponseBodyData {
	s.NewsTypeName = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyData) SetTotalPages(v int32) *ListHotNewsWithTypeResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListHotNewsWithTypeResponseBodyDataNews struct {
	Author  *string `json:"Author,omitempty" xml:"Author,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	DocUuid   *string   `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	ImageUrls []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-04-11 06:14:07
	PubTime          *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	SearchSource     *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	Source           *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Summary          *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Tag              *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Title            *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2023-10-14 14:30:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// http://xxxxx/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListHotNewsWithTypeResponseBodyDataNews) String() string {
	return dara.Prettify(s)
}

func (s ListHotNewsWithTypeResponseBodyDataNews) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) GetAuthor() *string {
	return s.Author
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) GetContent() *string {
	return s.Content
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) GetDocUuid() *string {
	return s.DocUuid
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) GetPubTime() *string {
	return s.PubTime
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) GetSearchSource() *string {
	return s.SearchSource
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) GetSource() *string {
	return s.Source
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) GetSummary() *string {
	return s.Summary
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) GetTag() *string {
	return s.Tag
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) GetTitle() *string {
	return s.Title
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) GetUrl() *string {
	return s.Url
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetAuthor(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Author = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetContent(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Content = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetDocUuid(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.DocUuid = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetImageUrls(v []*string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.ImageUrls = v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetPubTime(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.PubTime = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetSearchSource(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.SearchSource = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetSearchSourceName(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.SearchSourceName = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetSource(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Source = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetSummary(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Summary = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetTag(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Tag = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetTitle(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Title = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetUpdateTime(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.UpdateTime = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetUrl(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Url = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) Validate() error {
	return dara.Validate(s)
}
