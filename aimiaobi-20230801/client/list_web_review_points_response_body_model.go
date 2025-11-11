// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebReviewPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWebReviewPointsResponseBody
	GetCode() *string
	SetData(v []*ListWebReviewPointsResponseBodyData) *ListWebReviewPointsResponseBody
	GetData() []*ListWebReviewPointsResponseBodyData
	SetHttpStatusCode(v int32) *ListWebReviewPointsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListWebReviewPointsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListWebReviewPointsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListWebReviewPointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWebReviewPointsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWebReviewPointsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListWebReviewPointsResponseBody
	GetTotalCount() *int32
}

type ListWebReviewPointsResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListWebReviewPointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 79
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
	// 32
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWebReviewPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWebReviewPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWebReviewPointsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWebReviewPointsResponseBody) GetData() []*ListWebReviewPointsResponseBodyData {
	return s.Data
}

func (s *ListWebReviewPointsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListWebReviewPointsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWebReviewPointsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWebReviewPointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWebReviewPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWebReviewPointsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWebReviewPointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWebReviewPointsResponseBody) SetCode(v string) *ListWebReviewPointsResponseBody {
	s.Code = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetData(v []*ListWebReviewPointsResponseBodyData) *ListWebReviewPointsResponseBody {
	s.Data = v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetHttpStatusCode(v int32) *ListWebReviewPointsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetMaxResults(v int32) *ListWebReviewPointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetMessage(v string) *ListWebReviewPointsResponseBody {
	s.Message = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetNextToken(v string) *ListWebReviewPointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetRequestId(v string) *ListWebReviewPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetSuccess(v bool) *ListWebReviewPointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetTotalCount(v int32) *ListWebReviewPointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) Validate() error {
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

type ListWebReviewPointsResponseBodyData struct {
	// example:
	//
	// 当前观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点类型
	AttitudeType *string                                        `json:"AttitudeType,omitempty" xml:"AttitudeType,omitempty"`
	Comments     []*ListWebReviewPointsResponseBodyDataComments `json:"Comments,omitempty" xml:"Comments,omitempty" type:"Repeated"`
	// example:
	//
	// 当前观点占比
	Ratio      *string                                          `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ViewPoints []*ListWebReviewPointsResponseBodyDataViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s ListWebReviewPointsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWebReviewPointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWebReviewPointsResponseBodyData) GetAttitude() *string {
	return s.Attitude
}

func (s *ListWebReviewPointsResponseBodyData) GetAttitudeType() *string {
	return s.AttitudeType
}

func (s *ListWebReviewPointsResponseBodyData) GetComments() []*ListWebReviewPointsResponseBodyDataComments {
	return s.Comments
}

func (s *ListWebReviewPointsResponseBodyData) GetRatio() *string {
	return s.Ratio
}

func (s *ListWebReviewPointsResponseBodyData) GetViewPoints() []*ListWebReviewPointsResponseBodyDataViewPoints {
	return s.ViewPoints
}

func (s *ListWebReviewPointsResponseBodyData) SetAttitude(v string) *ListWebReviewPointsResponseBodyData {
	s.Attitude = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyData) SetAttitudeType(v string) *ListWebReviewPointsResponseBodyData {
	s.AttitudeType = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyData) SetComments(v []*ListWebReviewPointsResponseBodyDataComments) *ListWebReviewPointsResponseBodyData {
	s.Comments = v
	return s
}

func (s *ListWebReviewPointsResponseBodyData) SetRatio(v string) *ListWebReviewPointsResponseBodyData {
	s.Ratio = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyData) SetViewPoints(v []*ListWebReviewPointsResponseBodyDataViewPoints) *ListWebReviewPointsResponseBodyData {
	s.ViewPoints = v
	return s
}

func (s *ListWebReviewPointsResponseBodyData) Validate() error {
	if s.Comments != nil {
		for _, item := range s.Comments {
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

type ListWebReviewPointsResponseBodyDataComments struct {
	// example:
	//
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 评论内容
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 当前评论所属的URL
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// 评论用户名
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListWebReviewPointsResponseBodyDataComments) String() string {
	return dara.Prettify(s)
}

func (s ListWebReviewPointsResponseBodyDataComments) GoString() string {
	return s.String()
}

func (s *ListWebReviewPointsResponseBodyDataComments) GetSource() *string {
	return s.Source
}

func (s *ListWebReviewPointsResponseBodyDataComments) GetText() *string {
	return s.Text
}

func (s *ListWebReviewPointsResponseBodyDataComments) GetTitle() *string {
	return s.Title
}

func (s *ListWebReviewPointsResponseBodyDataComments) GetUrl() *string {
	return s.Url
}

func (s *ListWebReviewPointsResponseBodyDataComments) GetUsername() *string {
	return s.Username
}

func (s *ListWebReviewPointsResponseBodyDataComments) SetSource(v string) *ListWebReviewPointsResponseBodyDataComments {
	s.Source = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataComments) SetText(v string) *ListWebReviewPointsResponseBodyDataComments {
	s.Text = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataComments) SetTitle(v string) *ListWebReviewPointsResponseBodyDataComments {
	s.Title = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataComments) SetUrl(v string) *ListWebReviewPointsResponseBodyDataComments {
	s.Url = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataComments) SetUsername(v string) *ListWebReviewPointsResponseBodyDataComments {
	s.Username = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataComments) Validate() error {
	return dara.Validate(s)
}

type ListWebReviewPointsResponseBodyDataViewPoints struct {
	Outlines []*ListWebReviewPointsResponseBodyDataViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListWebReviewPointsResponseBodyDataViewPoints) String() string {
	return dara.Prettify(s)
}

func (s ListWebReviewPointsResponseBodyDataViewPoints) GoString() string {
	return s.String()
}

func (s *ListWebReviewPointsResponseBodyDataViewPoints) GetOutlines() []*ListWebReviewPointsResponseBodyDataViewPointsOutlines {
	return s.Outlines
}

func (s *ListWebReviewPointsResponseBodyDataViewPoints) GetPoint() *string {
	return s.Point
}

func (s *ListWebReviewPointsResponseBodyDataViewPoints) GetSummary() *string {
	return s.Summary
}

func (s *ListWebReviewPointsResponseBodyDataViewPoints) SetOutlines(v []*ListWebReviewPointsResponseBodyDataViewPointsOutlines) *ListWebReviewPointsResponseBodyDataViewPoints {
	s.Outlines = v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataViewPoints) SetPoint(v string) *ListWebReviewPointsResponseBodyDataViewPoints {
	s.Point = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataViewPoints) SetSummary(v string) *ListWebReviewPointsResponseBodyDataViewPoints {
	s.Summary = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataViewPoints) Validate() error {
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

type ListWebReviewPointsResponseBodyDataViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListWebReviewPointsResponseBodyDataViewPointsOutlines) String() string {
	return dara.Prettify(s)
}

func (s ListWebReviewPointsResponseBodyDataViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *ListWebReviewPointsResponseBodyDataViewPointsOutlines) GetOutline() *string {
	return s.Outline
}

func (s *ListWebReviewPointsResponseBodyDataViewPointsOutlines) GetSummary() *string {
	return s.Summary
}

func (s *ListWebReviewPointsResponseBodyDataViewPointsOutlines) SetOutline(v string) *ListWebReviewPointsResponseBodyDataViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataViewPointsOutlines) SetSummary(v string) *ListWebReviewPointsResponseBodyDataViewPointsOutlines {
	s.Summary = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataViewPointsOutlines) Validate() error {
	return dara.Validate(s)
}
