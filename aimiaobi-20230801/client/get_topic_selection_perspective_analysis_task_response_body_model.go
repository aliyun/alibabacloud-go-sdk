// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicSelectionPerspectiveAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetCode() *string
	SetData(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetData() *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData
	SetHttpStatusCode(v int32) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetSuccess() *bool
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) GetData() *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	return s.Data
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) SetCode(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) SetData(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) SetMessage(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) SetRequestId(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) SetSuccess(v bool) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData struct {
	// example:
	//
	// 错误信息
	ErrorMessage          *string                                                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FreshViewPointsResult *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult `json:"FreshViewPointsResult,omitempty" xml:"FreshViewPointsResult,omitempty" type:"Struct"`
	HotViewPointsResult   *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult   `json:"HotViewPointsResult,omitempty" xml:"HotViewPointsResult,omitempty" type:"Struct"`
	// example:
	//
	// SUSPENDED
	Status                *string                                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	TimedViewPointsResult *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult `json:"TimedViewPointsResult,omitempty" xml:"TimedViewPointsResult,omitempty" type:"Struct"`
	// example:
	//
	// 热点主题事件
	Topic                 *string                                                                        `json:"Topic,omitempty" xml:"Topic,omitempty"`
	TopicSummaryResult    *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult    `json:"TopicSummaryResult,omitempty" xml:"TopicSummaryResult,omitempty" type:"Struct"`
	WebReviewPointsResult *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult `json:"WebReviewPointsResult,omitempty" xml:"WebReviewPointsResult,omitempty" type:"Struct"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GetFreshViewPointsResult() *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult {
	return s.FreshViewPointsResult
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GetHotViewPointsResult() *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult {
	return s.HotViewPointsResult
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GetTimedViewPointsResult() *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult {
	return s.TimedViewPointsResult
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GetTopicSummaryResult() *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult {
	return s.TopicSummaryResult
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GetWebReviewPointsResult() *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult {
	return s.WebReviewPointsResult
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetErrorMessage(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetFreshViewPointsResult(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.FreshViewPointsResult = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetHotViewPointsResult(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.HotViewPointsResult = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetStatus(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetTimedViewPointsResult(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.TimedViewPointsResult = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetTopic(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.Topic = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetTopicSummaryResult(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.TopicSummaryResult = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetWebReviewPointsResult(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.WebReviewPointsResult = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) Validate() error {
	if s.FreshViewPointsResult != nil {
		if err := s.FreshViewPointsResult.Validate(); err != nil {
			return err
		}
	}
	if s.HotViewPointsResult != nil {
		if err := s.HotViewPointsResult.Validate(); err != nil {
			return err
		}
	}
	if s.TimedViewPointsResult != nil {
		if err := s.TimedViewPointsResult.Validate(); err != nil {
			return err
		}
	}
	if s.TopicSummaryResult != nil {
		if err := s.TopicSummaryResult.Validate(); err != nil {
			return err
		}
	}
	if s.WebReviewPointsResult != nil {
		if err := s.WebReviewPointsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult struct {
	Attitudes []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes `json:"Attitudes,omitempty" xml:"Attitudes,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult) GetAttitudes() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes {
	return s.Attitudes
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult) SetAttitudes(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult {
	s.Attitudes = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult) Validate() error {
	if s.Attitudes != nil {
		for _, item := range s.Attitudes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes struct {
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
	// 当前观点占比
	Ratio      *string                                                                                             `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ViewPoints []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) GetAttitude() *string {
	return s.Attitude
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) GetAttitudeType() *string {
	return s.AttitudeType
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) GetRatio() *string {
	return s.Ratio
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) GetViewPoints() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints {
	return s.ViewPoints
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) SetAttitude(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes {
	s.Attitude = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) SetAttitudeType(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes {
	s.AttitudeType = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) SetRatio(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes {
	s.Ratio = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) SetViewPoints(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes {
	s.ViewPoints = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) Validate() error {
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

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints struct {
	Outlines []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) GetOutlines() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines {
	return s.Outlines
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) GetPoint() *string {
	return s.Point
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) GetSummary() *string {
	return s.Summary
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) SetOutlines(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints {
	s.Outlines = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) SetPoint(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints {
	s.Point = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints {
	s.Summary = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) Validate() error {
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

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines) GetOutline() *string {
	return s.Outline
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines) GetSummary() *string {
	return s.Summary
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines) SetOutline(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines {
	s.Summary = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines) Validate() error {
	return dara.Validate(s)
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult struct {
	Attitudes []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes `json:"Attitudes,omitempty" xml:"Attitudes,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult) GetAttitudes() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes {
	return s.Attitudes
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult) SetAttitudes(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult {
	s.Attitudes = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult) Validate() error {
	if s.Attitudes != nil {
		for _, item := range s.Attitudes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes struct {
	// example:
	//
	// 当前观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点类型
	AttitudeType *string                                                                                     `json:"AttitudeType,omitempty" xml:"AttitudeType,omitempty"`
	News         []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews `json:"News,omitempty" xml:"News,omitempty" type:"Repeated"`
	// example:
	//
	// 当前观点占比
	Ratio      *string                                                                                           `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ViewPoints []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) GetAttitude() *string {
	return s.Attitude
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) GetAttitudeType() *string {
	return s.AttitudeType
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) GetNews() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	return s.News
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) GetRatio() *string {
	return s.Ratio
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) GetViewPoints() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints {
	return s.ViewPoints
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) SetAttitude(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes {
	s.Attitude = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) SetAttitudeType(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes {
	s.AttitudeType = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) SetNews(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes {
	s.News = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) SetRatio(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes {
	s.Ratio = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) SetViewPoints(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes {
	s.ViewPoints = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) Validate() error {
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

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews struct {
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 9957175DEDCF49C5ACF7A956B4FD67B2
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 123456
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.png
	ImageUrls []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	PubTime   *string   `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	Source    *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	Summary   *string   `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// ["标签1","标签2"]
	Tags  []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Title *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 文章主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Url   *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) GetContent() *string {
	return s.Content
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) GetDocId() *string {
	return s.DocId
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) GetDocUuid() *string {
	return s.DocUuid
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) GetPubTime() *string {
	return s.PubTime
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) GetSource() *string {
	return s.Source
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) GetSummary() *string {
	return s.Summary
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) GetTags() []*string {
	return s.Tags
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) GetTitle() *string {
	return s.Title
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) GetTopic() *string {
	return s.Topic
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) GetUrl() *string {
	return s.Url
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetContent(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.Content = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetCreateTime(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.CreateTime = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetDocId(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.DocId = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetDocUuid(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.DocUuid = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetImageUrls(v []*string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.ImageUrls = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetPubTime(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.PubTime = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetSource(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.Source = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.Summary = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetTags(v []*string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.Tags = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetTitle(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.Title = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetTopic(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.Topic = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetUrl(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.Url = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) Validate() error {
	return dara.Validate(s)
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints struct {
	Outlines []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) GetOutlines() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines {
	return s.Outlines
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) GetPoint() *string {
	return s.Point
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) GetSummary() *string {
	return s.Summary
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) SetOutlines(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints {
	s.Outlines = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) SetPoint(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints {
	s.Point = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints {
	s.Summary = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) Validate() error {
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

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines) GetOutline() *string {
	return s.Outline
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines) GetSummary() *string {
	return s.Summary
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines) SetOutline(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines {
	s.Summary = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines) Validate() error {
	return dara.Validate(s)
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult struct {
	Attitudes []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes `json:"Attitudes,omitempty" xml:"Attitudes,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult) GetAttitudes() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	return s.Attitudes
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult) SetAttitudes(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult {
	s.Attitudes = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult) Validate() error {
	if s.Attitudes != nil {
		for _, item := range s.Attitudes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes struct {
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
	Url        *string                                                                                             `json:"Url,omitempty" xml:"Url,omitempty"`
	ViewPoints []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) GetAttitude() *string {
	return s.Attitude
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) GetAttitudeType() *string {
	return s.AttitudeType
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) GetPubTime() *string {
	return s.PubTime
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) GetRatio() *string {
	return s.Ratio
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) GetSource() *string {
	return s.Source
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) GetTitle() *string {
	return s.Title
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) GetUrl() *string {
	return s.Url
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) GetViewPoints() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints {
	return s.ViewPoints
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetAttitude(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.Attitude = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetAttitudeType(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.AttitudeType = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetPubTime(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.PubTime = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetRatio(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.Ratio = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetSource(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.Source = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetTitle(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.Title = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetUrl(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.Url = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetViewPoints(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.ViewPoints = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) Validate() error {
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

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints struct {
	Outlines []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) GetOutlines() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines {
	return s.Outlines
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) GetPoint() *string {
	return s.Point
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) GetSummary() *string {
	return s.Summary
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) SetOutlines(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints {
	s.Outlines = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) SetPoint(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints {
	s.Point = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints {
	s.Summary = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) Validate() error {
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

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines) GetOutline() *string {
	return s.Outline
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines) GetSummary() *string {
	return s.Summary
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines) SetOutline(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines {
	s.Summary = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines) Validate() error {
	return dara.Validate(s)
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult struct {
	Summaries []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries `json:"Summaries,omitempty" xml:"Summaries,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult) GetSummaries() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries {
	return s.Summaries
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult) SetSummaries(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult {
	s.Summaries = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult) Validate() error {
	if s.Summaries != nil {
		for _, item := range s.Summaries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries struct {
	DocList []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList `json:"DocList,omitempty" xml:"DocList,omitempty" type:"Repeated"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) GetDocList() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList {
	return s.DocList
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) GetSummary() *string {
	return s.Summary
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) GetTitle() *string {
	return s.Title
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) SetDocList(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries {
	s.DocList = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries {
	s.Summary = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) SetTitle(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries {
	s.Title = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) Validate() error {
	if s.DocList != nil {
		for _, item := range s.DocList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList struct {
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) GetSource() *string {
	return s.Source
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) GetTitle() *string {
	return s.Title
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) GetUrl() *string {
	return s.Url
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) SetSource(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList {
	s.Source = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) SetTitle(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList {
	s.Title = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) SetUrl(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList {
	s.Url = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) Validate() error {
	return dara.Validate(s)
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult struct {
	Attitudes []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes `json:"Attitudes,omitempty" xml:"Attitudes,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult) GetAttitudes() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes {
	return s.Attitudes
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult) SetAttitudes(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult {
	s.Attitudes = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult) Validate() error {
	if s.Attitudes != nil {
		for _, item := range s.Attitudes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes struct {
	// example:
	//
	// 当前观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点类型
	AttitudeType *string                                                                                           `json:"AttitudeType,omitempty" xml:"AttitudeType,omitempty"`
	Comments     []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments `json:"Comments,omitempty" xml:"Comments,omitempty" type:"Repeated"`
	// example:
	//
	// 当前观点占比
	Ratio      *string                                                                                             `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ViewPoints []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) GetAttitude() *string {
	return s.Attitude
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) GetAttitudeType() *string {
	return s.AttitudeType
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) GetComments() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments {
	return s.Comments
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) GetRatio() *string {
	return s.Ratio
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) GetViewPoints() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints {
	return s.ViewPoints
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) SetAttitude(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes {
	s.Attitude = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) SetAttitudeType(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes {
	s.AttitudeType = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) SetComments(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes {
	s.Comments = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) SetRatio(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes {
	s.Ratio = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) SetViewPoints(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes {
	s.ViewPoints = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) Validate() error {
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

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments struct {
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

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) GetSource() *string {
	return s.Source
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) GetText() *string {
	return s.Text
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) GetTitle() *string {
	return s.Title
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) GetUrl() *string {
	return s.Url
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) GetUsername() *string {
	return s.Username
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) SetSource(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments {
	s.Source = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) SetText(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments {
	s.Text = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) SetTitle(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments {
	s.Title = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) SetUrl(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments {
	s.Url = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) SetUsername(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments {
	s.Username = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) Validate() error {
	return dara.Validate(s)
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints struct {
	Outlines []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) GetOutlines() []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines {
	return s.Outlines
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) GetPoint() *string {
	return s.Point
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) GetSummary() *string {
	return s.Summary
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) SetOutlines(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints {
	s.Outlines = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) SetPoint(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints {
	s.Point = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints {
	s.Summary = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) Validate() error {
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

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines) GetOutline() *string {
	return s.Outline
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines) GetSummary() *string {
	return s.Summary
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines) SetOutline(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines {
	s.Summary = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines) Validate() error {
	return dara.Validate(s)
}
