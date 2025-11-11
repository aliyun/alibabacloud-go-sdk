// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBiddingDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBiddingDocResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListBiddingDocResponseBody
	GetCurrent() *int32
	SetData(v []*ListBiddingDocResponseBodyData) *ListBiddingDocResponseBody
	GetData() []*ListBiddingDocResponseBodyData
	SetHttpStatusCode(v int32) *ListBiddingDocResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListBiddingDocResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListBiddingDocResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListBiddingDocResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListBiddingDocResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListBiddingDocResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListBiddingDocResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListBiddingDocResponseBody
	GetTotal() *int32
	SetTotalCount(v int32) *ListBiddingDocResponseBody
	GetTotalCount() *int32
}

type ListBiddingDocResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                            `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*ListBiddingDocResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// null
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
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
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// null
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBiddingDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBiddingDocResponseBody) GoString() string {
	return s.String()
}

func (s *ListBiddingDocResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBiddingDocResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListBiddingDocResponseBody) GetData() []*ListBiddingDocResponseBodyData {
	return s.Data
}

func (s *ListBiddingDocResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBiddingDocResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBiddingDocResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBiddingDocResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBiddingDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBiddingDocResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListBiddingDocResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBiddingDocResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListBiddingDocResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBiddingDocResponseBody) SetCode(v string) *ListBiddingDocResponseBody {
	s.Code = &v
	return s
}

func (s *ListBiddingDocResponseBody) SetCurrent(v int32) *ListBiddingDocResponseBody {
	s.Current = &v
	return s
}

func (s *ListBiddingDocResponseBody) SetData(v []*ListBiddingDocResponseBodyData) *ListBiddingDocResponseBody {
	s.Data = v
	return s
}

func (s *ListBiddingDocResponseBody) SetHttpStatusCode(v int32) *ListBiddingDocResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBiddingDocResponseBody) SetMaxResults(v int32) *ListBiddingDocResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBiddingDocResponseBody) SetMessage(v string) *ListBiddingDocResponseBody {
	s.Message = &v
	return s
}

func (s *ListBiddingDocResponseBody) SetNextToken(v string) *ListBiddingDocResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBiddingDocResponseBody) SetRequestId(v string) *ListBiddingDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBiddingDocResponseBody) SetSize(v int32) *ListBiddingDocResponseBody {
	s.Size = &v
	return s
}

func (s *ListBiddingDocResponseBody) SetSuccess(v bool) *ListBiddingDocResponseBody {
	s.Success = &v
	return s
}

func (s *ListBiddingDocResponseBody) SetTotal(v int32) *ListBiddingDocResponseBody {
	s.Total = &v
	return s
}

func (s *ListBiddingDocResponseBody) SetTotalCount(v int32) *ListBiddingDocResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBiddingDocResponseBody) Validate() error {
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

type ListBiddingDocResponseBodyData struct {
	// example:
	//
	// 2023-03-18 02:00:00
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// analysis
	//
	// writing
	TaskStep *string `json:"TaskStep,omitempty" xml:"TaskStep,omitempty"`
}

func (s ListBiddingDocResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBiddingDocResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBiddingDocResponseBodyData) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *ListBiddingDocResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListBiddingDocResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *ListBiddingDocResponseBodyData) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ListBiddingDocResponseBodyData) GetTaskStep() *string {
	return s.TaskStep
}

func (s *ListBiddingDocResponseBodyData) SetCreateTimeStart(v string) *ListBiddingDocResponseBodyData {
	s.CreateTimeStart = &v
	return s
}

func (s *ListBiddingDocResponseBodyData) SetTaskId(v string) *ListBiddingDocResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListBiddingDocResponseBodyData) SetTaskName(v string) *ListBiddingDocResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *ListBiddingDocResponseBodyData) SetTaskStatus(v int32) *ListBiddingDocResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *ListBiddingDocResponseBodyData) SetTaskStep(v string) *ListBiddingDocResponseBodyData {
	s.TaskStep = &v
	return s
}

func (s *ListBiddingDocResponseBodyData) Validate() error {
	return dara.Validate(s)
}
