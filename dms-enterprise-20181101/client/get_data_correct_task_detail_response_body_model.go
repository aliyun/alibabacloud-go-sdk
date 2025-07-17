// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataCorrectTaskDetail(v *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) *GetDataCorrectTaskDetailResponseBody
	GetDataCorrectTaskDetail() *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail
	SetErrorCode(v string) *GetDataCorrectTaskDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataCorrectTaskDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataCorrectTaskDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataCorrectTaskDetailResponseBody
	GetSuccess() *bool
}

type GetDataCorrectTaskDetailResponseBody struct {
	// The details of the data change task.
	DataCorrectTaskDetail *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail `json:"DataCorrectTaskDetail,omitempty" xml:"DataCorrectTaskDetail,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B5FD0BC8-2D90-4478-B8EC-A0E92E0B1773
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataCorrectTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCorrectTaskDetailResponseBody) GetDataCorrectTaskDetail() *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail {
	return s.DataCorrectTaskDetail
}

func (s *GetDataCorrectTaskDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataCorrectTaskDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataCorrectTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCorrectTaskDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataCorrectTaskDetailResponseBody) SetDataCorrectTaskDetail(v *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) *GetDataCorrectTaskDetailResponseBody {
	s.DataCorrectTaskDetail = v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBody) SetErrorCode(v string) *GetDataCorrectTaskDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBody) SetErrorMessage(v string) *GetDataCorrectTaskDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBody) SetRequestId(v string) *GetDataCorrectTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBody) SetSuccess(v bool) *GetDataCorrectTaskDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail struct {
	// The number of rows affected by the SQL statement.
	//
	// example:
	//
	// 1
	ActualAffectRows *int64 `json:"ActualAffectRows,omitempty" xml:"ActualAffectRows,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2021-03-05 15:08:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the SQL task group.
	//
	// example:
	//
	// 1234235
	DBTaskGroupId *int64 `json:"DBTaskGroupId,omitempty" xml:"DBTaskGroupId,omitempty"`
	// The state of the SQL task. Valid values:
	//
	// 	- **INIT**: The SQL task was initialized.
	//
	// 	- **PENDING**: The SQL task waited to be run.
	//
	// 	- **BE_SCHEDULED**: The SQL task waited to be scheduled.
	//
	// 	- **FAIL**: The SQL task failed.
	//
	// 	- **SUCCESS**: The SQL task was successful.
	//
	// 	- **PAUSE**: The SQL task was paused.
	//
	// 	- **DELETE**: The SQL task was deleted.
	//
	// 	- **RUNNING**: The SQL task was being run.
	//
	// example:
	//
	// SUCCESS
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
}

func (s GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) GoString() string {
	return s.String()
}

func (s *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) GetActualAffectRows() *int64 {
	return s.ActualAffectRows
}

func (s *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) GetDBTaskGroupId() *int64 {
	return s.DBTaskGroupId
}

func (s *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) SetActualAffectRows(v int64) *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail {
	s.ActualAffectRows = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) SetCreateTime(v string) *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail {
	s.CreateTime = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) SetDBTaskGroupId(v int64) *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail {
	s.DBTaskGroupId = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) SetJobStatus(v string) *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail {
	s.JobStatus = &v
	return s
}

func (s *GetDataCorrectTaskDetailResponseBodyDataCorrectTaskDetail) Validate() error {
	return dara.Validate(s)
}
