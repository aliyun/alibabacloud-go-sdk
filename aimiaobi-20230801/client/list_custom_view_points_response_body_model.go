// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomViewPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCustomViewPointsResponseBody
	GetCode() *string
	SetData(v []*ListCustomViewPointsResponseBodyData) *ListCustomViewPointsResponseBody
	GetData() []*ListCustomViewPointsResponseBodyData
	SetHttpStatusCode(v int32) *ListCustomViewPointsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListCustomViewPointsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListCustomViewPointsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListCustomViewPointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCustomViewPointsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCustomViewPointsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListCustomViewPointsResponseBody
	GetTotalCount() *int32
}

type ListCustomViewPointsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data []*ListCustomViewPointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The maximum number of results returned.
	//
	// example:
	//
	// 60
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The error description.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token for the next page of results.
	//
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The unique identifier of the request.
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. \\`true\\`: The request was successful. \\`false\\`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 73
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomViewPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomViewPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomViewPointsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCustomViewPointsResponseBody) GetData() []*ListCustomViewPointsResponseBodyData {
	return s.Data
}

func (s *ListCustomViewPointsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCustomViewPointsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCustomViewPointsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCustomViewPointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomViewPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomViewPointsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCustomViewPointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomViewPointsResponseBody) SetCode(v string) *ListCustomViewPointsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetData(v []*ListCustomViewPointsResponseBodyData) *ListCustomViewPointsResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetHttpStatusCode(v int32) *ListCustomViewPointsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetMaxResults(v int32) *ListCustomViewPointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetMessage(v string) *ListCustomViewPointsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetNextToken(v string) *ListCustomViewPointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetRequestId(v string) *ListCustomViewPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetSuccess(v bool) *ListCustomViewPointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetTotalCount(v int32) *ListCustomViewPointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) Validate() error {
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

type ListCustomViewPointsResponseBodyData struct {
	// The ID of the asynchronous task that analyzes the custom viewpoint.
	//
	// example:
	//
	// 2323ac73e174428a98c91097a59c67e0
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// The viewpoint.
	//
	// example:
	//
	// 观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// The time when the viewpoint was created.
	//
	// example:
	//
	// 2024-08-15 16:18:59
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who created the viewpoint.
	//
	// example:
	//
	// 1
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The ID of the custom viewpoint.
	//
	// example:
	//
	// 709806dd051042d5ab9de8bdbb3a64ca
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The task execution status. Valid values: PENDING, RUNNING, SUCCEEDED, SUSPENDED, FAILED, and CANCELED.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The error message for the task execution.
	//
	// example:
	//
	// 参数校验失败
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// The task execution status. 0: PENDING, 1: RUNNING, 2: SUCCEEDED, 3: PAUSED, 4: FAILED (retriable), 5: FAILED (non-retriable), 6: CANCELED.
	//
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// A list of topic selection viewpoints.
	ViewPoints []*ListCustomViewPointsResponseBodyDataViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s ListCustomViewPointsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCustomViewPointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomViewPointsResponseBodyData) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *ListCustomViewPointsResponseBodyData) GetAttitude() *string {
	return s.Attitude
}

func (s *ListCustomViewPointsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCustomViewPointsResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListCustomViewPointsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListCustomViewPointsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListCustomViewPointsResponseBodyData) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *ListCustomViewPointsResponseBodyData) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ListCustomViewPointsResponseBodyData) GetViewPoints() []*ListCustomViewPointsResponseBodyDataViewPoints {
	return s.ViewPoints
}

func (s *ListCustomViewPointsResponseBodyData) SetAsyncTaskId(v string) *ListCustomViewPointsResponseBodyData {
	s.AsyncTaskId = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetAttitude(v string) *ListCustomViewPointsResponseBodyData {
	s.Attitude = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetCreateTime(v string) *ListCustomViewPointsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetCreateUser(v string) *ListCustomViewPointsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetId(v string) *ListCustomViewPointsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetStatus(v string) *ListCustomViewPointsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetTaskErrorMessage(v string) *ListCustomViewPointsResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetTaskStatus(v int32) *ListCustomViewPointsResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetViewPoints(v []*ListCustomViewPointsResponseBodyDataViewPoints) *ListCustomViewPointsResponseBodyData {
	s.ViewPoints = v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) Validate() error {
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

type ListCustomViewPointsResponseBodyDataViewPoints struct {
	// The outline.
	Outlines []*ListCustomViewPointsResponseBodyDataViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// The generated viewpoint.
	//
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// The summary.
	//
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListCustomViewPointsResponseBodyDataViewPoints) String() string {
	return dara.Prettify(s)
}

func (s ListCustomViewPointsResponseBodyDataViewPoints) GoString() string {
	return s.String()
}

func (s *ListCustomViewPointsResponseBodyDataViewPoints) GetOutlines() []*ListCustomViewPointsResponseBodyDataViewPointsOutlines {
	return s.Outlines
}

func (s *ListCustomViewPointsResponseBodyDataViewPoints) GetPoint() *string {
	return s.Point
}

func (s *ListCustomViewPointsResponseBodyDataViewPoints) GetSummary() *string {
	return s.Summary
}

func (s *ListCustomViewPointsResponseBodyDataViewPoints) SetOutlines(v []*ListCustomViewPointsResponseBodyDataViewPointsOutlines) *ListCustomViewPointsResponseBodyDataViewPoints {
	s.Outlines = v
	return s
}

func (s *ListCustomViewPointsResponseBodyDataViewPoints) SetPoint(v string) *ListCustomViewPointsResponseBodyDataViewPoints {
	s.Point = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyDataViewPoints) SetSummary(v string) *ListCustomViewPointsResponseBodyDataViewPoints {
	s.Summary = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyDataViewPoints) Validate() error {
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

type ListCustomViewPointsResponseBodyDataViewPointsOutlines struct {
	// The outline.
	//
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// The summary of the outline.
	//
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListCustomViewPointsResponseBodyDataViewPointsOutlines) String() string {
	return dara.Prettify(s)
}

func (s ListCustomViewPointsResponseBodyDataViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *ListCustomViewPointsResponseBodyDataViewPointsOutlines) GetOutline() *string {
	return s.Outline
}

func (s *ListCustomViewPointsResponseBodyDataViewPointsOutlines) GetSummary() *string {
	return s.Summary
}

func (s *ListCustomViewPointsResponseBodyDataViewPointsOutlines) SetOutline(v string) *ListCustomViewPointsResponseBodyDataViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyDataViewPointsOutlines) SetSummary(v string) *ListCustomViewPointsResponseBodyDataViewPointsOutlines {
	s.Summary = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyDataViewPointsOutlines) Validate() error {
	return dara.Validate(s)
}
