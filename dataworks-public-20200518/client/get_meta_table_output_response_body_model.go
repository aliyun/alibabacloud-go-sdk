// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaTableOutputResponseBodyData) *GetMetaTableOutputResponseBody
	GetData() *GetMetaTableOutputResponseBodyData
	SetErrorCode(v string) *GetMetaTableOutputResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaTableOutputResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMetaTableOutputResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMetaTableOutputResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaTableOutputResponseBody
	GetSuccess() *bool
}

type GetMetaTableOutputResponseBody struct {
	// The business data.
	Data *GetMetaTableOutputResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// SUCCESS
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0000-0000-0000-000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetaTableOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableOutputResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaTableOutputResponseBody) GetData() *GetMetaTableOutputResponseBodyData {
	return s.Data
}

func (s *GetMetaTableOutputResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaTableOutputResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaTableOutputResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMetaTableOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaTableOutputResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaTableOutputResponseBody) SetData(v *GetMetaTableOutputResponseBodyData) *GetMetaTableOutputResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaTableOutputResponseBody) SetErrorCode(v string) *GetMetaTableOutputResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaTableOutputResponseBody) SetErrorMessage(v string) *GetMetaTableOutputResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaTableOutputResponseBody) SetHttpStatusCode(v int32) *GetMetaTableOutputResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMetaTableOutputResponseBody) SetRequestId(v string) *GetMetaTableOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTableOutputResponseBody) SetSuccess(v bool) *GetMetaTableOutputResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaTableOutputResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableOutputResponseBodyData struct {
	// The partitions.
	DataEntityList []*GetMetaTableOutputResponseBodyDataDataEntityList `json:"DataEntityList,omitempty" xml:"DataEntityList,omitempty" type:"Repeated"`
	// The page number. Valid values: 1 to 30. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 128
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetMetaTableOutputResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableOutputResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaTableOutputResponseBodyData) GetDataEntityList() []*GetMetaTableOutputResponseBodyDataDataEntityList {
	return s.DataEntityList
}

func (s *GetMetaTableOutputResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetMetaTableOutputResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaTableOutputResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetMetaTableOutputResponseBodyData) SetDataEntityList(v []*GetMetaTableOutputResponseBodyDataDataEntityList) *GetMetaTableOutputResponseBodyData {
	s.DataEntityList = v
	return s
}

func (s *GetMetaTableOutputResponseBodyData) SetPageNumber(v int32) *GetMetaTableOutputResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetMetaTableOutputResponseBodyData) SetPageSize(v int32) *GetMetaTableOutputResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetMetaTableOutputResponseBodyData) SetTotalCount(v int64) *GetMetaTableOutputResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetMetaTableOutputResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableOutputResponseBodyDataDataEntityList struct {
	// The end time.
	//
	// example:
	//
	// 2022-02-12 0:32:12
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 128
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2022-02-12 0:34:13
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The GUID of the MaxCompute table.
	//
	// example:
	//
	// odps.sample_project.sample_table
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1048576
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 987654321
	TaskInstanceId *int64 `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	// The waiting time.
	//
	// example:
	//
	// 3
	WaitTime *string `json:"WaitTime,omitempty" xml:"WaitTime,omitempty"`
}

func (s GetMetaTableOutputResponseBodyDataDataEntityList) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableOutputResponseBodyDataDataEntityList) GoString() string {
	return s.String()
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) GetEndTime() *string {
	return s.EndTime
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) GetTaskId() *string {
	return s.TaskId
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) GetTaskInstanceId() *int64 {
	return s.TaskInstanceId
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) GetWaitTime() *string {
	return s.WaitTime
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) SetEndTime(v string) *GetMetaTableOutputResponseBodyDataDataEntityList {
	s.EndTime = &v
	return s
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) SetProjectId(v int64) *GetMetaTableOutputResponseBodyDataDataEntityList {
	s.ProjectId = &v
	return s
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) SetStartTime(v string) *GetMetaTableOutputResponseBodyDataDataEntityList {
	s.StartTime = &v
	return s
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) SetTableGuid(v string) *GetMetaTableOutputResponseBodyDataDataEntityList {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) SetTaskId(v string) *GetMetaTableOutputResponseBodyDataDataEntityList {
	s.TaskId = &v
	return s
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) SetTaskInstanceId(v int64) *GetMetaTableOutputResponseBodyDataDataEntityList {
	s.TaskInstanceId = &v
	return s
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) SetWaitTime(v string) *GetMetaTableOutputResponseBodyDataDataEntityList {
	s.WaitTime = &v
	return s
}

func (s *GetMetaTableOutputResponseBodyDataDataEntityList) Validate() error {
	return dara.Validate(s)
}
