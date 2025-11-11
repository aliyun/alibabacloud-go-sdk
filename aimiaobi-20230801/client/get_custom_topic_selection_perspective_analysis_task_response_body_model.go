// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetCode() *string
	SetData(v *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetData() *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData
	SetHttpStatusCode(v int32) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody
	GetSuccess() *bool
}

type GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GetData() *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	return s.Data
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetCode(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetData(v *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetMessage(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetRequestId(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetSuccess(v bool) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData struct {
	CustomViewPointsResult *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult `json:"CustomViewPointsResult,omitempty" xml:"CustomViewPointsResult,omitempty" type:"Struct"`
	// example:
	//
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// FAILED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GetCustomViewPointsResult() *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult {
	return s.CustomViewPointsResult
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetCustomViewPointsResult(v *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.CustomViewPointsResult = v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetErrorMessage(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetStatus(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) Validate() error {
	if s.CustomViewPointsResult != nil {
		if err := s.CustomViewPointsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult struct {
	Attitudes []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes `json:"Attitudes,omitempty" xml:"Attitudes,omitempty" type:"Repeated"`
	// example:
	//
	// 热点主题事件
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult) GetAttitudes() []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes {
	return s.Attitudes
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult) GetTopic() *string {
	return s.Topic
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult) SetAttitudes(v []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult {
	s.Attitudes = v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult) SetTopic(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult {
	s.Topic = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult) Validate() error {
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

type GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes struct {
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
	Ratio      *string                                                                                                    `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ViewPoints []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) GetAttitude() *string {
	return s.Attitude
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) GetAttitudeType() *string {
	return s.AttitudeType
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) GetRatio() *string {
	return s.Ratio
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) GetViewPoints() []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints {
	return s.ViewPoints
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) SetAttitude(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes {
	s.Attitude = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) SetAttitudeType(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes {
	s.AttitudeType = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) SetRatio(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes {
	s.Ratio = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) SetViewPoints(v []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes {
	s.ViewPoints = v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) Validate() error {
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

type GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints struct {
	Outlines []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) GetOutlines() []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines {
	return s.Outlines
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) GetPoint() *string {
	return s.Point
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) GetSummary() *string {
	return s.Summary
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) SetOutlines(v []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints {
	s.Outlines = v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) SetPoint(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints {
	s.Point = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) SetSummary(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints {
	s.Summary = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) Validate() error {
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

type GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines) GetOutline() *string {
	return s.Outline
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines) GetSummary() *string {
	return s.Summary
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines) SetOutline(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines) SetSummary(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines {
	s.Summary = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines) Validate() error {
	return dara.Validate(s)
}
