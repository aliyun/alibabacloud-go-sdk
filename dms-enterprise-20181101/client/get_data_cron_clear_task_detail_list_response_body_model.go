// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCronClearTaskDetailListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataCronClearTaskDetailList(v []*GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) *GetDataCronClearTaskDetailListResponseBody
	GetDataCronClearTaskDetailList() []*GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList
	SetErrorCode(v string) *GetDataCronClearTaskDetailListResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataCronClearTaskDetailListResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataCronClearTaskDetailListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataCronClearTaskDetailListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *GetDataCronClearTaskDetailListResponseBody
	GetTotalCount() *int64
}

type GetDataCronClearTaskDetailListResponseBody struct {
	// The historical data cleansing tasks
	DataCronClearTaskDetailList []*GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList `json:"DataCronClearTaskDetailList,omitempty" xml:"DataCronClearTaskDetailList,omitempty" type:"Repeated"`
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
	// C1D39814-9808-47F8-AFE0-AF167239AC9B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of SQL tasks.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetDataCronClearTaskDetailListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataCronClearTaskDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCronClearTaskDetailListResponseBody) GetDataCronClearTaskDetailList() []*GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList {
	return s.DataCronClearTaskDetailList
}

func (s *GetDataCronClearTaskDetailListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataCronClearTaskDetailListResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataCronClearTaskDetailListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCronClearTaskDetailListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataCronClearTaskDetailListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetDataCronClearTaskDetailListResponseBody) SetDataCronClearTaskDetailList(v []*GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) *GetDataCronClearTaskDetailListResponseBody {
	s.DataCronClearTaskDetailList = v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBody) SetErrorCode(v string) *GetDataCronClearTaskDetailListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBody) SetErrorMessage(v string) *GetDataCronClearTaskDetailListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBody) SetRequestId(v string) *GetDataCronClearTaskDetailListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBody) SetSuccess(v bool) *GetDataCronClearTaskDetailListResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBody) SetTotalCount(v int64) *GetDataCronClearTaskDetailListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBody) Validate() error {
	if s.DataCronClearTaskDetailList != nil {
		for _, item := range s.DataCronClearTaskDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList struct {
	// The number of rows affected by the SQL task.
	//
	// example:
	//
	// 1
	ActualAffectRows *int64 `json:"ActualAffectRows,omitempty" xml:"ActualAffectRows,omitempty"`
	// The time when the SQL task was created.
	//
	// example:
	//
	// 2021-01-14 10:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the SQL task group.
	//
	// example:
	//
	// 432523
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

func (s GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) String() string {
	return dara.Prettify(s)
}

func (s GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) GoString() string {
	return s.String()
}

func (s *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) GetActualAffectRows() *int64 {
	return s.ActualAffectRows
}

func (s *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) GetDBTaskGroupId() *int64 {
	return s.DBTaskGroupId
}

func (s *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) SetActualAffectRows(v int64) *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList {
	s.ActualAffectRows = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) SetCreateTime(v string) *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList {
	s.CreateTime = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) SetDBTaskGroupId(v int64) *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList {
	s.DBTaskGroupId = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) SetJobStatus(v string) *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList {
	s.JobStatus = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponseBodyDataCronClearTaskDetailList) Validate() error {
	return dara.Validate(s)
}
