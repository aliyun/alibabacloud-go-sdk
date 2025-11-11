// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTimedViewAttitudeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTimedViewAttitudeResponseBody
	GetCode() *string
	SetData(v []*ListTimedViewAttitudeResponseBodyData) *ListTimedViewAttitudeResponseBody
	GetData() []*ListTimedViewAttitudeResponseBodyData
	SetHttpStatusCode(v int32) *ListTimedViewAttitudeResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListTimedViewAttitudeResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListTimedViewAttitudeResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTimedViewAttitudeResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTimedViewAttitudeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTimedViewAttitudeResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListTimedViewAttitudeResponseBody
	GetTotalCount() *int32
}

type ListTimedViewAttitudeResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListTimedViewAttitudeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 15
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
	// 58
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTimedViewAttitudeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTimedViewAttitudeResponseBody) GoString() string {
	return s.String()
}

func (s *ListTimedViewAttitudeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTimedViewAttitudeResponseBody) GetData() []*ListTimedViewAttitudeResponseBodyData {
	return s.Data
}

func (s *ListTimedViewAttitudeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTimedViewAttitudeResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTimedViewAttitudeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTimedViewAttitudeResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTimedViewAttitudeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTimedViewAttitudeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTimedViewAttitudeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTimedViewAttitudeResponseBody) SetCode(v string) *ListTimedViewAttitudeResponseBody {
	s.Code = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetData(v []*ListTimedViewAttitudeResponseBodyData) *ListTimedViewAttitudeResponseBody {
	s.Data = v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetHttpStatusCode(v int32) *ListTimedViewAttitudeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetMaxResults(v int32) *ListTimedViewAttitudeResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetMessage(v string) *ListTimedViewAttitudeResponseBody {
	s.Message = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetNextToken(v string) *ListTimedViewAttitudeResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetRequestId(v string) *ListTimedViewAttitudeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetSuccess(v bool) *ListTimedViewAttitudeResponseBody {
	s.Success = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetTotalCount(v int32) *ListTimedViewAttitudeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) Validate() error {
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

type ListTimedViewAttitudeResponseBodyData struct {
	// example:
	//
	// 当前观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点类型
	AttitudeType *string `json:"AttitudeType,omitempty" xml:"AttitudeType,omitempty"`
	// example:
	//
	// 2024-01-22 10:29
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 当前观点占比
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// example:
	//
	// 新浪
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://www.example.com/news/1.html
	Url        *string                                            `json:"Url,omitempty" xml:"Url,omitempty"`
	ViewPoints []*ListTimedViewAttitudeResponseBodyDataViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s ListTimedViewAttitudeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTimedViewAttitudeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTimedViewAttitudeResponseBodyData) GetAttitude() *string {
	return s.Attitude
}

func (s *ListTimedViewAttitudeResponseBodyData) GetAttitudeType() *string {
	return s.AttitudeType
}

func (s *ListTimedViewAttitudeResponseBodyData) GetPubTime() *string {
	return s.PubTime
}

func (s *ListTimedViewAttitudeResponseBodyData) GetRatio() *string {
	return s.Ratio
}

func (s *ListTimedViewAttitudeResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *ListTimedViewAttitudeResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListTimedViewAttitudeResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *ListTimedViewAttitudeResponseBodyData) GetViewPoints() []*ListTimedViewAttitudeResponseBodyDataViewPoints {
	return s.ViewPoints
}

func (s *ListTimedViewAttitudeResponseBodyData) SetAttitude(v string) *ListTimedViewAttitudeResponseBodyData {
	s.Attitude = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) SetAttitudeType(v string) *ListTimedViewAttitudeResponseBodyData {
	s.AttitudeType = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) SetPubTime(v string) *ListTimedViewAttitudeResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) SetRatio(v string) *ListTimedViewAttitudeResponseBodyData {
	s.Ratio = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) SetSource(v string) *ListTimedViewAttitudeResponseBodyData {
	s.Source = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) SetTitle(v string) *ListTimedViewAttitudeResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) SetUrl(v string) *ListTimedViewAttitudeResponseBodyData {
	s.Url = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) SetViewPoints(v []*ListTimedViewAttitudeResponseBodyDataViewPoints) *ListTimedViewAttitudeResponseBodyData {
	s.ViewPoints = v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) Validate() error {
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

type ListTimedViewAttitudeResponseBodyDataViewPoints struct {
	Outlines []*ListTimedViewAttitudeResponseBodyDataViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListTimedViewAttitudeResponseBodyDataViewPoints) String() string {
	return dara.Prettify(s)
}

func (s ListTimedViewAttitudeResponseBodyDataViewPoints) GoString() string {
	return s.String()
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPoints) GetOutlines() []*ListTimedViewAttitudeResponseBodyDataViewPointsOutlines {
	return s.Outlines
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPoints) GetPoint() *string {
	return s.Point
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPoints) GetSummary() *string {
	return s.Summary
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPoints) SetOutlines(v []*ListTimedViewAttitudeResponseBodyDataViewPointsOutlines) *ListTimedViewAttitudeResponseBodyDataViewPoints {
	s.Outlines = v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPoints) SetPoint(v string) *ListTimedViewAttitudeResponseBodyDataViewPoints {
	s.Point = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPoints) SetSummary(v string) *ListTimedViewAttitudeResponseBodyDataViewPoints {
	s.Summary = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPoints) Validate() error {
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

type ListTimedViewAttitudeResponseBodyDataViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListTimedViewAttitudeResponseBodyDataViewPointsOutlines) String() string {
	return dara.Prettify(s)
}

func (s ListTimedViewAttitudeResponseBodyDataViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPointsOutlines) GetOutline() *string {
	return s.Outline
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPointsOutlines) GetSummary() *string {
	return s.Summary
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPointsOutlines) SetOutline(v string) *ListTimedViewAttitudeResponseBodyDataViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPointsOutlines) SetSummary(v string) *ListTimedViewAttitudeResponseBodyDataViewPointsOutlines {
	s.Summary = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPointsOutlines) Validate() error {
	return dara.Validate(s)
}
