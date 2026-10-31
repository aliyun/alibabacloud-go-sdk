// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRectifyFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRectifyFlowResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListRectifyFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListRectifyFlowResponseBody
	GetMessage() *string
	SetPageResult(v *ListRectifyFlowResponseBodyPageResult) *ListRectifyFlowResponseBody
	GetPageResult() *ListRectifyFlowResponseBodyPageResult
	SetRequestId(v string) *ListRectifyFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRectifyFlowResponseBody
	GetSuccess() *bool
}

type ListRectifyFlowResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination result.
	PageResult *ListRectifyFlowResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRectifyFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRectifyFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListRectifyFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRectifyFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRectifyFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRectifyFlowResponseBody) GetPageResult() *ListRectifyFlowResponseBodyPageResult {
	return s.PageResult
}

func (s *ListRectifyFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRectifyFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRectifyFlowResponseBody) SetCode(v string) *ListRectifyFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ListRectifyFlowResponseBody) SetHttpStatusCode(v int32) *ListRectifyFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRectifyFlowResponseBody) SetMessage(v string) *ListRectifyFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ListRectifyFlowResponseBody) SetPageResult(v *ListRectifyFlowResponseBodyPageResult) *ListRectifyFlowResponseBody {
	s.PageResult = v
	return s
}

func (s *ListRectifyFlowResponseBody) SetRequestId(v string) *ListRectifyFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRectifyFlowResponseBody) SetSuccess(v bool) *ListRectifyFlowResponseBody {
	s.Success = &v
	return s
}

func (s *ListRectifyFlowResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRectifyFlowResponseBodyPageResult struct {
	// The processes on the current page.
	Data []*ListRectifyFlowResponseBodyPageResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of processes.
	//
	// example:
	//
	// 699386170553
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRectifyFlowResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListRectifyFlowResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListRectifyFlowResponseBodyPageResult) GetData() []*ListRectifyFlowResponseBodyPageResultData {
	return s.Data
}

func (s *ListRectifyFlowResponseBodyPageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRectifyFlowResponseBodyPageResult) SetData(v []*ListRectifyFlowResponseBodyPageResultData) *ListRectifyFlowResponseBodyPageResult {
	s.Data = v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResult) SetTotalCount(v int64) *ListRectifyFlowResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResult) Validate() error {
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

type ListRectifyFlowResponseBodyPageResultData struct {
	// The completion time. The value is a UNIX timestamp in milliseconds. If the process is not completed, this parameter is empty.
	//
	// example:
	//
	// 648218233037
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The total number of issues in the current process.
	//
	// example:
	//
	// 264461775547
	ProblemCount *int64 `json:"ProblemCount,omitempty" xml:"ProblemCount,omitempty"`
	// The process ID.
	//
	// example:
	//
	// 474448887101
	RectifyId *int64 `json:"RectifyId,omitempty" xml:"RectifyId,omitempty"`
	// The process name.
	//
	// example:
	//
	// test
	RectifyName *string `json:"RectifyName,omitempty" xml:"RectifyName,omitempty"`
	// The priority.
	//
	// example:
	//
	// 1
	RectifyPriority *string `json:"RectifyPriority,omitempty" xml:"RectifyPriority,omitempty"`
	// The rectifier ID.
	//
	// example:
	//
	// 3001
	RectifyUser *string `json:"RectifyUser,omitempty" xml:"RectifyUser,omitempty"`
	// The rectifier name.
	//
	// example:
	//
	// test
	RectifyUserName *string `json:"RectifyUserName,omitempty" xml:"RectifyUserName,omitempty"`
	// The start time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 356278124692
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The process status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The initiator ID.
	//
	// example:
	//
	// 3001
	SubmitUser *string `json:"SubmitUser,omitempty" xml:"SubmitUser,omitempty"`
	// The initiator name.
	//
	// example:
	//
	// test
	SubmitUserName *string `json:"SubmitUserName,omitempty" xml:"SubmitUserName,omitempty"`
	// The verifier ID.
	//
	// example:
	//
	// 3001
	VerifyUser *string `json:"VerifyUser,omitempty" xml:"VerifyUser,omitempty"`
	// The verifier name.
	//
	// example:
	//
	// test
	VerifyUserName *string `json:"VerifyUserName,omitempty" xml:"VerifyUserName,omitempty"`
}

func (s ListRectifyFlowResponseBodyPageResultData) String() string {
	return dara.Prettify(s)
}

func (s ListRectifyFlowResponseBodyPageResultData) GoString() string {
	return s.String()
}

func (s *ListRectifyFlowResponseBodyPageResultData) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *ListRectifyFlowResponseBodyPageResultData) GetProblemCount() *int64 {
	return s.ProblemCount
}

func (s *ListRectifyFlowResponseBodyPageResultData) GetRectifyId() *int64 {
	return s.RectifyId
}

func (s *ListRectifyFlowResponseBodyPageResultData) GetRectifyName() *string {
	return s.RectifyName
}

func (s *ListRectifyFlowResponseBodyPageResultData) GetRectifyPriority() *string {
	return s.RectifyPriority
}

func (s *ListRectifyFlowResponseBodyPageResultData) GetRectifyUser() *string {
	return s.RectifyUser
}

func (s *ListRectifyFlowResponseBodyPageResultData) GetRectifyUserName() *string {
	return s.RectifyUserName
}

func (s *ListRectifyFlowResponseBodyPageResultData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListRectifyFlowResponseBodyPageResultData) GetStatus() *string {
	return s.Status
}

func (s *ListRectifyFlowResponseBodyPageResultData) GetSubmitUser() *string {
	return s.SubmitUser
}

func (s *ListRectifyFlowResponseBodyPageResultData) GetSubmitUserName() *string {
	return s.SubmitUserName
}

func (s *ListRectifyFlowResponseBodyPageResultData) GetVerifyUser() *string {
	return s.VerifyUser
}

func (s *ListRectifyFlowResponseBodyPageResultData) GetVerifyUserName() *string {
	return s.VerifyUserName
}

func (s *ListRectifyFlowResponseBodyPageResultData) SetFinishTime(v int64) *ListRectifyFlowResponseBodyPageResultData {
	s.FinishTime = &v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResultData) SetProblemCount(v int64) *ListRectifyFlowResponseBodyPageResultData {
	s.ProblemCount = &v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResultData) SetRectifyId(v int64) *ListRectifyFlowResponseBodyPageResultData {
	s.RectifyId = &v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResultData) SetRectifyName(v string) *ListRectifyFlowResponseBodyPageResultData {
	s.RectifyName = &v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResultData) SetRectifyPriority(v string) *ListRectifyFlowResponseBodyPageResultData {
	s.RectifyPriority = &v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResultData) SetRectifyUser(v string) *ListRectifyFlowResponseBodyPageResultData {
	s.RectifyUser = &v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResultData) SetRectifyUserName(v string) *ListRectifyFlowResponseBodyPageResultData {
	s.RectifyUserName = &v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResultData) SetStartTime(v int64) *ListRectifyFlowResponseBodyPageResultData {
	s.StartTime = &v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResultData) SetStatus(v string) *ListRectifyFlowResponseBodyPageResultData {
	s.Status = &v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResultData) SetSubmitUser(v string) *ListRectifyFlowResponseBodyPageResultData {
	s.SubmitUser = &v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResultData) SetSubmitUserName(v string) *ListRectifyFlowResponseBodyPageResultData {
	s.SubmitUserName = &v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResultData) SetVerifyUser(v string) *ListRectifyFlowResponseBodyPageResultData {
	s.VerifyUser = &v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResultData) SetVerifyUserName(v string) *ListRectifyFlowResponseBodyPageResultData {
	s.VerifyUserName = &v
	return s
}

func (s *ListRectifyFlowResponseBodyPageResultData) Validate() error {
	return dara.Validate(s)
}
