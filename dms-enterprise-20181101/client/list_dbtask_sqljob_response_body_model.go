// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDBTaskSQLJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBTaskSQLJobList(v []*ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) *ListDBTaskSQLJobResponseBody
	GetDBTaskSQLJobList() []*ListDBTaskSQLJobResponseBodyDBTaskSQLJobList
	SetErrorCode(v string) *ListDBTaskSQLJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDBTaskSQLJobResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDBTaskSQLJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDBTaskSQLJobResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListDBTaskSQLJobResponseBody
	GetTotalCount() *int64
}

type ListDBTaskSQLJobResponseBody struct {
	// The list of the SQL tasks.
	DBTaskSQLJobList []*ListDBTaskSQLJobResponseBodyDBTaskSQLJobList `json:"DBTaskSQLJobList,omitempty" xml:"DBTaskSQLJobList,omitempty" type:"Repeated"`
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
	// F6C47680-8D2D-43A4-8902-F2740D71A398
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of the SQL tasks.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDBTaskSQLJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDBTaskSQLJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobResponseBody) GetDBTaskSQLJobList() []*ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	return s.DBTaskSQLJobList
}

func (s *ListDBTaskSQLJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDBTaskSQLJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDBTaskSQLJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDBTaskSQLJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDBTaskSQLJobResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDBTaskSQLJobResponseBody) SetDBTaskSQLJobList(v []*ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) *ListDBTaskSQLJobResponseBody {
	s.DBTaskSQLJobList = v
	return s
}

func (s *ListDBTaskSQLJobResponseBody) SetErrorCode(v string) *ListDBTaskSQLJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBody) SetErrorMessage(v string) *ListDBTaskSQLJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBody) SetRequestId(v string) *ListDBTaskSQLJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBody) SetSuccess(v bool) *ListDBTaskSQLJobResponseBody {
	s.Success = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBody) SetTotalCount(v int64) *ListDBTaskSQLJobResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBody) Validate() error {
	if s.DBTaskSQLJobList != nil {
		for _, item := range s.DBTaskSQLJobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDBTaskSQLJobResponseBodyDBTaskSQLJobList struct {
	// The description of the SQL task.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the SQL task was created.
	//
	// example:
	//
	// 2021-02-18 17:49:20
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the database.
	//
	// example:
	//
	// 43214523
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The name that is used to search for the database.
	//
	// example:
	//
	// test@xxx:3306【test】
	DbSearchName *string `json:"DbSearchName,omitempty" xml:"DbSearchName,omitempty"`
	// The ID of the SQL task group.
	//
	// example:
	//
	// 4324132
	DbTaskGroupId *int64 `json:"DbTaskGroupId,omitempty" xml:"DbTaskGroupId,omitempty"`
	// The ID of the SQL task.
	//
	// example:
	//
	// 123435
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The type of the SQL task.
	//
	// example:
	//
	// STRUCT_SYNC
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The time when the SQL task was last executed.
	//
	// example:
	//
	// 2021-02-18 17:49:31
	LastExecTime *string `json:"LastExecTime,omitempty" xml:"LastExecTime,omitempty"`
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is a logical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
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
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the SQL task is executed as a transaction. Valid values:
	//
	// 	- **true**: The SQL task is executed as a transaction.
	//
	// 	- **false**: The SQL task is not executed as a transaction.
	//
	// example:
	//
	// false
	Transactional *bool `json:"Transactional,omitempty" xml:"Transactional,omitempty"`
}

func (s ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) String() string {
	return dara.Prettify(s)
}

func (s ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) GetComment() *string {
	return s.Comment
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) GetDbId() *int64 {
	return s.DbId
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) GetDbSearchName() *string {
	return s.DbSearchName
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) GetDbTaskGroupId() *int64 {
	return s.DbTaskGroupId
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) GetJobId() *int64 {
	return s.JobId
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) GetJobType() *string {
	return s.JobType
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) GetLastExecTime() *string {
	return s.LastExecTime
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) GetLogic() *bool {
	return s.Logic
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) GetStatus() *string {
	return s.Status
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) GetTransactional() *bool {
	return s.Transactional
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetComment(v string) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.Comment = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetCreateTime(v string) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.CreateTime = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetDbId(v int64) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.DbId = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetDbSearchName(v string) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.DbSearchName = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetDbTaskGroupId(v int64) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.DbTaskGroupId = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetJobId(v int64) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.JobId = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetJobType(v string) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.JobType = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetLastExecTime(v string) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.LastExecTime = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetLogic(v bool) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.Logic = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetStatus(v string) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.Status = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) SetTransactional(v bool) *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList {
	s.Transactional = &v
	return s
}

func (s *ListDBTaskSQLJobResponseBodyDBTaskSQLJobList) Validate() error {
	return dara.Validate(s)
}
