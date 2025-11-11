// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotViewPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHotViewPointsResponseBody
	GetCode() *string
	SetData(v []*ListHotViewPointsResponseBodyData) *ListHotViewPointsResponseBody
	GetData() []*ListHotViewPointsResponseBodyData
	SetHttpStatusCode(v int32) *ListHotViewPointsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListHotViewPointsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListHotViewPointsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListHotViewPointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListHotViewPointsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHotViewPointsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListHotViewPointsResponseBody
	GetTotalCount() *int32
}

type ListHotViewPointsResponseBody struct {
	// example:
	//
	// NoData
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListHotViewPointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 67
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 70
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHotViewPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotViewPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotViewPointsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHotViewPointsResponseBody) GetData() []*ListHotViewPointsResponseBodyData {
	return s.Data
}

func (s *ListHotViewPointsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHotViewPointsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHotViewPointsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotViewPointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHotViewPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotViewPointsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHotViewPointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHotViewPointsResponseBody) SetCode(v string) *ListHotViewPointsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotViewPointsResponseBody) SetData(v []*ListHotViewPointsResponseBodyData) *ListHotViewPointsResponseBody {
	s.Data = v
	return s
}

func (s *ListHotViewPointsResponseBody) SetHttpStatusCode(v int32) *ListHotViewPointsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHotViewPointsResponseBody) SetMaxResults(v int32) *ListHotViewPointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHotViewPointsResponseBody) SetMessage(v string) *ListHotViewPointsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotViewPointsResponseBody) SetNextToken(v string) *ListHotViewPointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHotViewPointsResponseBody) SetRequestId(v string) *ListHotViewPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotViewPointsResponseBody) SetSuccess(v bool) *ListHotViewPointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListHotViewPointsResponseBody) SetTotalCount(v int32) *ListHotViewPointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHotViewPointsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotViewPointsResponseBodyData struct {
	// example:
	//
	// 当前观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点类型
	AttitudeType *string                                  `json:"AttitudeType,omitempty" xml:"AttitudeType,omitempty"`
	News         []*ListHotViewPointsResponseBodyDataNews `json:"News,omitempty" xml:"News,omitempty" type:"Repeated"`
	// example:
	//
	// 当前观点占比
	Ratio      *string                                        `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ViewPoints []*ListHotViewPointsResponseBodyDataViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s ListHotViewPointsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHotViewPointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotViewPointsResponseBodyData) GetAttitude() *string {
	return s.Attitude
}

func (s *ListHotViewPointsResponseBodyData) GetAttitudeType() *string {
	return s.AttitudeType
}

func (s *ListHotViewPointsResponseBodyData) GetNews() []*ListHotViewPointsResponseBodyDataNews {
	return s.News
}

func (s *ListHotViewPointsResponseBodyData) GetRatio() *string {
	return s.Ratio
}

func (s *ListHotViewPointsResponseBodyData) GetViewPoints() []*ListHotViewPointsResponseBodyDataViewPoints {
	return s.ViewPoints
}

func (s *ListHotViewPointsResponseBodyData) SetAttitude(v string) *ListHotViewPointsResponseBodyData {
	s.Attitude = &v
	return s
}

func (s *ListHotViewPointsResponseBodyData) SetAttitudeType(v string) *ListHotViewPointsResponseBodyData {
	s.AttitudeType = &v
	return s
}

func (s *ListHotViewPointsResponseBodyData) SetNews(v []*ListHotViewPointsResponseBodyDataNews) *ListHotViewPointsResponseBodyData {
	s.News = v
	return s
}

func (s *ListHotViewPointsResponseBodyData) SetRatio(v string) *ListHotViewPointsResponseBodyData {
	s.Ratio = &v
	return s
}

func (s *ListHotViewPointsResponseBodyData) SetViewPoints(v []*ListHotViewPointsResponseBodyDataViewPoints) *ListHotViewPointsResponseBodyData {
	s.ViewPoints = v
	return s
}

func (s *ListHotViewPointsResponseBodyData) Validate() error {
	if s.News != nil {
		for _, item := range s.News {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ViewPoints != nil {
		for _, item := range s.ViewPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotViewPointsResponseBodyDataNews struct {
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
	// xxxxx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 123456
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.png
	ImageUrls []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-01-22 10:29:00
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 新浪
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// ["标签1","标签2"]
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 文章主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListHotViewPointsResponseBodyDataNews) String() string {
	return dara.Prettify(s)
}

func (s ListHotViewPointsResponseBodyDataNews) GoString() string {
	return s.String()
}

func (s *ListHotViewPointsResponseBodyDataNews) GetAuthor() *string {
	return s.Author
}

func (s *ListHotViewPointsResponseBodyDataNews) GetContent() *string {
	return s.Content
}

func (s *ListHotViewPointsResponseBodyDataNews) GetDocId() *string {
	return s.DocId
}

func (s *ListHotViewPointsResponseBodyDataNews) GetDocUuid() *string {
	return s.DocUuid
}

func (s *ListHotViewPointsResponseBodyDataNews) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *ListHotViewPointsResponseBodyDataNews) GetPubTime() *string {
	return s.PubTime
}

func (s *ListHotViewPointsResponseBodyDataNews) GetSource() *string {
	return s.Source
}

func (s *ListHotViewPointsResponseBodyDataNews) GetSummary() *string {
	return s.Summary
}

func (s *ListHotViewPointsResponseBodyDataNews) GetTags() []*string {
	return s.Tags
}

func (s *ListHotViewPointsResponseBodyDataNews) GetTitle() *string {
	return s.Title
}

func (s *ListHotViewPointsResponseBodyDataNews) GetTopic() *string {
	return s.Topic
}

func (s *ListHotViewPointsResponseBodyDataNews) GetUrl() *string {
	return s.Url
}

func (s *ListHotViewPointsResponseBodyDataNews) SetAuthor(v string) *ListHotViewPointsResponseBodyDataNews {
	s.Author = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetContent(v string) *ListHotViewPointsResponseBodyDataNews {
	s.Content = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetDocId(v string) *ListHotViewPointsResponseBodyDataNews {
	s.DocId = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetDocUuid(v string) *ListHotViewPointsResponseBodyDataNews {
	s.DocUuid = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetImageUrls(v []*string) *ListHotViewPointsResponseBodyDataNews {
	s.ImageUrls = v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetPubTime(v string) *ListHotViewPointsResponseBodyDataNews {
	s.PubTime = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetSource(v string) *ListHotViewPointsResponseBodyDataNews {
	s.Source = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetSummary(v string) *ListHotViewPointsResponseBodyDataNews {
	s.Summary = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetTags(v []*string) *ListHotViewPointsResponseBodyDataNews {
	s.Tags = v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetTitle(v string) *ListHotViewPointsResponseBodyDataNews {
	s.Title = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetTopic(v string) *ListHotViewPointsResponseBodyDataNews {
	s.Topic = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetUrl(v string) *ListHotViewPointsResponseBodyDataNews {
	s.Url = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) Validate() error {
	return dara.Validate(s)
}

type ListHotViewPointsResponseBodyDataViewPoints struct {
	Outlines []*ListHotViewPointsResponseBodyDataViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListHotViewPointsResponseBodyDataViewPoints) String() string {
	return dara.Prettify(s)
}

func (s ListHotViewPointsResponseBodyDataViewPoints) GoString() string {
	return s.String()
}

func (s *ListHotViewPointsResponseBodyDataViewPoints) GetOutlines() []*ListHotViewPointsResponseBodyDataViewPointsOutlines {
	return s.Outlines
}

func (s *ListHotViewPointsResponseBodyDataViewPoints) GetPoint() *string {
	return s.Point
}

func (s *ListHotViewPointsResponseBodyDataViewPoints) GetSummary() *string {
	return s.Summary
}

func (s *ListHotViewPointsResponseBodyDataViewPoints) SetOutlines(v []*ListHotViewPointsResponseBodyDataViewPointsOutlines) *ListHotViewPointsResponseBodyDataViewPoints {
	s.Outlines = v
	return s
}

func (s *ListHotViewPointsResponseBodyDataViewPoints) SetPoint(v string) *ListHotViewPointsResponseBodyDataViewPoints {
	s.Point = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataViewPoints) SetSummary(v string) *ListHotViewPointsResponseBodyDataViewPoints {
	s.Summary = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataViewPoints) Validate() error {
	if s.Outlines != nil {
		for _, item := range s.Outlines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotViewPointsResponseBodyDataViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListHotViewPointsResponseBodyDataViewPointsOutlines) String() string {
	return dara.Prettify(s)
}

func (s ListHotViewPointsResponseBodyDataViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *ListHotViewPointsResponseBodyDataViewPointsOutlines) GetOutline() *string {
	return s.Outline
}

func (s *ListHotViewPointsResponseBodyDataViewPointsOutlines) GetSummary() *string {
	return s.Summary
}

func (s *ListHotViewPointsResponseBodyDataViewPointsOutlines) SetOutline(v string) *ListHotViewPointsResponseBodyDataViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataViewPointsOutlines) SetSummary(v string) *ListHotViewPointsResponseBodyDataViewPointsOutlines {
	s.Summary = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataViewPointsOutlines) Validate() error {
	return dara.Validate(s)
}
