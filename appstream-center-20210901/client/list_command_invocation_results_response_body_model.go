// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommandInvocationResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommandInvocationResultModel(v []*ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) *ListCommandInvocationResultsResponseBody
	GetCommandInvocationResultModel() []*ListCommandInvocationResultsResponseBodyCommandInvocationResultModel
	SetMaxResults(v int32) *ListCommandInvocationResultsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListCommandInvocationResultsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListCommandInvocationResultsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCommandInvocationResultsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCommandInvocationResultsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCommandInvocationResultsResponseBody
	GetTotalCount() *int32
}

type ListCommandInvocationResultsResponseBody struct {
	// The list of command execution results. Each element corresponds to an execution record on a persistent session. If no records meet the conditions, an empty array is returned.
	CommandInvocationResultModel []*ListCommandInvocationResultsResponseBodyCommandInvocationResultModel `json:"CommandInvocationResultModel,omitempty" xml:"CommandInvocationResultModel,omitempty" type:"Repeated"`
	// This field is not returned in the current version. Use PageNumber, PageSize, and TotalCount for pagination.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// This field is not returned in the current version. Use PageNumber, PageSize, and TotalCount for pagination.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kDw2Ml****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number of the currently displayed query results. In the current version, this field always returns 0 and does not echo echo request parameter PageNumber. Refer to echo request parameter for the actual value.
	//
	// example:
	//
	// 0
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of query results per page. In the current version, this field always returns 0 and does not echo echo request parameter PageSize. Refer to echo request parameter for the actual value.
	//
	// example:
	//
	// 0
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of query results. If no records meet the conditions, 0 is returned.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCommandInvocationResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCommandInvocationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommandInvocationResultsResponseBody) GetCommandInvocationResultModel() []*ListCommandInvocationResultsResponseBodyCommandInvocationResultModel {
	return s.CommandInvocationResultModel
}

func (s *ListCommandInvocationResultsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCommandInvocationResultsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCommandInvocationResultsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCommandInvocationResultsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCommandInvocationResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCommandInvocationResultsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCommandInvocationResultsResponseBody) SetCommandInvocationResultModel(v []*ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) *ListCommandInvocationResultsResponseBody {
	s.CommandInvocationResultModel = v
	return s
}

func (s *ListCommandInvocationResultsResponseBody) SetMaxResults(v int32) *ListCommandInvocationResultsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCommandInvocationResultsResponseBody) SetNextToken(v string) *ListCommandInvocationResultsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCommandInvocationResultsResponseBody) SetPageNumber(v int32) *ListCommandInvocationResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCommandInvocationResultsResponseBody) SetPageSize(v int32) *ListCommandInvocationResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCommandInvocationResultsResponseBody) SetRequestId(v string) *ListCommandInvocationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommandInvocationResultsResponseBody) SetTotalCount(v int32) *ListCommandInvocationResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCommandInvocationResultsResponseBody) Validate() error {
	if s.CommandInvocationResultModel != nil {
		for _, item := range s.CommandInvocationResultModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCommandInvocationResultsResponseBodyCommandInvocationResultModel struct {
	// The ID of the delivery group to which the persistent session that executed the command belongs.
	//
	// example:
	//
	// ws-0bxls9m9arax6****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The instance ID of the app that executed the command.
	//
	// example:
	//
	// ai-azn3kmwruh1vl****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The ID of the persistent session that executed the command.
	//
	// example:
	//
	// p-0cc7s3mw2fg4j****
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// The time when the command execution finished. The time is in UTC and follows the RFC 3339 standard in the yyyy-MM-ddTHH:mm:ssZ format. It is an empty string if the command execution has not finished.
	//
	// example:
	//
	// 2026-11-10T08:49:30Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The command execution ID, which is the same as the request parameter InvocationId.
	//
	// example:
	//
	// t-0bxls9m9ara****
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
	// The execution status of the command on the persistent session, such as Pending (waiting to be executed), Running (executing), Success (executed successfully), and Timeout (execution timed out).
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The output content generated by the command execution, which is raw text. It is an empty string if the command has not finished or has no output.
	//
	// example:
	//
	// hello world
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The time when the command started to execute. The time is in UTC and follows the RFC 3339 standard in the yyyy-MM-ddTHH:mm:ssZ format. It is an empty string if the command has not started to execute.
	//
	// example:
	//
	// 2026-11-10T08:49:24Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) String() string {
	return dara.Prettify(s)
}

func (s ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) GoString() string {
	return s.String()
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) GetAppInstancePersistentId() *string {
	return s.AppInstancePersistentId
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) GetInvocationId() *string {
	return s.InvocationId
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) GetOutput() *string {
	return s.Output
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) GetStartTime() *string {
	return s.StartTime
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) SetAppInstanceGroupId(v string) *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) SetAppInstanceId(v string) *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel {
	s.AppInstanceId = &v
	return s
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) SetAppInstancePersistentId(v string) *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel {
	s.AppInstancePersistentId = &v
	return s
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) SetFinishTime(v string) *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel {
	s.FinishTime = &v
	return s
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) SetInvocationId(v string) *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel {
	s.InvocationId = &v
	return s
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) SetInvocationStatus(v string) *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel {
	s.InvocationStatus = &v
	return s
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) SetOutput(v string) *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel {
	s.Output = &v
	return s
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) SetStartTime(v string) *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel {
	s.StartTime = &v
	return s
}

func (s *ListCommandInvocationResultsResponseBodyCommandInvocationResultModel) Validate() error {
	return dara.Validate(s)
}
