// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnlineDDLProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetOnlineDDLProgressResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetOnlineDDLProgressResponseBody
	GetErrorMessage() *string
	SetOnlineDDLTaskDetail(v *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) *GetOnlineDDLProgressResponseBody
	GetOnlineDDLTaskDetail() *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail
	SetRequestId(v string) *GetOnlineDDLProgressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOnlineDDLProgressResponseBody
	GetSuccess() *bool
}

type GetOnlineDDLProgressResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The details of the task.
	OnlineDDLTaskDetail *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail `json:"OnlineDDLTaskDetail,omitempty" xml:"OnlineDDLTaskDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 34E01EDD-6A16-4CF0-9541-C644D1BE01AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOnlineDDLProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineDDLProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetOnlineDDLProgressResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetOnlineDDLProgressResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetOnlineDDLProgressResponseBody) GetOnlineDDLTaskDetail() *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail {
	return s.OnlineDDLTaskDetail
}

func (s *GetOnlineDDLProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOnlineDDLProgressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOnlineDDLProgressResponseBody) SetErrorCode(v string) *GetOnlineDDLProgressResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBody) SetErrorMessage(v string) *GetOnlineDDLProgressResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBody) SetOnlineDDLTaskDetail(v *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) *GetOnlineDDLProgressResponseBody {
	s.OnlineDDLTaskDetail = v
	return s
}

func (s *GetOnlineDDLProgressResponseBody) SetRequestId(v string) *GetOnlineDDLProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBody) SetSuccess(v bool) *GetOnlineDDLProgressResponseBody {
	s.Success = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBody) Validate() error {
	if s.OnlineDDLTaskDetail != nil {
		if err := s.OnlineDDLTaskDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail struct {
	// The cleanup policy of the original table after the cut-over. Valid values:
	//
	// 	- **DROP**: Invalid original tables are deleted.
	//
	// 	- **MOVE**: Invalid original tables are moved to the test database. You can delete the tables manually.
	//
	// 	- **NOTHING**: Invalid original tables are retained in the original database. You can delete the tables manually.
	//
	// example:
	//
	// DROP
	CleanStrategy *string `json:"CleanStrategy,omitempty" xml:"CleanStrategy,omitempty"`
	// The policy of full replication. Valid values:
	//
	// 	- **AUTO**: DMS dynamically adjusts the chunk size based on the performance of the database. Tables are locked for less than 1.5 seconds during a single replication operation.
	//
	// 	- **RUNNING**: DMS uses the specified value of the CopyChunkSize parameter. The valid value of the CopyChunkSize parameter ranges from 1 to 60000. If you set this parameter to RUNNING, you must specify the CopyChunkSize parameter.
	//
	// example:
	//
	// AUTO
	CopyChunkMode *string `json:"CopyChunkMode,omitempty" xml:"CopyChunkMode,omitempty"`
	// The size of each chunk that is used to replicate data. This parameter is used to specify the size of each chunk. A larger chunk size increases the replication efficiency and decreases the business performance.
	//
	// > During full replication, the original table is divided into N small chunks and each chunk is replicated to the temporary table one by one. By default, DMS dynamically adjusts the size of each chunk.
	//
	// example:
	//
	// 1000
	CopyChunkSize *int64 `json:"CopyChunkSize,omitempty" xml:"CopyChunkSize,omitempty"`
	// The actual amount of data replicated from the original table in the lock-free change operation.
	//
	// example:
	//
	// 9
	CopyCount *int64 `json:"CopyCount,omitempty" xml:"CopyCount,omitempty"`
	// The estimated total number of rows of the data. The value is obtained from the statistical data in the information_schema database. In most cases, the estimated total number of rows is smaller than the actual number of rows in a table.
	//
	// example:
	//
	// 10
	CopyTotal *int64 `json:"CopyTotal,omitempty" xml:"CopyTotal,omitempty"`
	// The number of retries when the cut-over fails.
	//
	// example:
	//
	// 3
	CutoverFailRetryTimes *int64 `json:"CutoverFailRetryTimes,omitempty" xml:"CutoverFailRetryTimes,omitempty"`
	// The maximum period of time that a table can be locked during cut-over. Unit: seconds.
	//
	// example:
	//
	// 2
	CutoverLockTimeSeconds *int64 `json:"CutoverLockTimeSeconds,omitempty" xml:"CutoverLockTimeSeconds,omitempty"`
	// The end of the time window of the cut-over operation. This value is at least 30 minutes later than the CutoverWindowStartTime parameter. Default value: 23:59:59
	//
	// example:
	//
	// 13:00:00
	CutoverWindowEndTime *string `json:"CutoverWindowEndTime,omitempty" xml:"CutoverWindowEndTime,omitempty"`
	// The beginning of the time window of the cut-over operation. Default value: 00:00:00. This parameter controls the time window of the cut-over. Cut-over can be performed only when the cut-over conditions are met and the time is within the specified time window. If the time is not within the time window, the cut-over operation is not performed until the time reaches the beginning of the time window.
	//
	// example:
	//
	// 12:00:00
	CutoverWindowStartTime *string `json:"CutoverWindowStartTime,omitempty" xml:"CutoverWindowStartTime,omitempty"`
	// The replay latency of DMS. Unit: seconds. The replay latency is the period of time that is taken to replay the binary logs of the table to the temporary table. The latency does not indicate the data migration latency between a primary database and a secondary database.
	//
	// example:
	//
	// 0
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The state of the task. Valid values:
	//
	// 	- **INIT**: The task is being initialized.
	//
	// 	- **SUCCESS**: The task is complete.
	//
	// 	- **RUNNING**: The task is being executed.
	//
	// 	- **WAITING_CUTOVER**: The task is waiting for cut-over.
	//
	// 	- **RESTARTING**: The task is restarting.
	//
	// 	- **PAUSE**: The task is suspended.
	//
	// 	- **UNSUPPORTED**: The task is not supported.
	//
	// 	- **CANCELED**: The task is canceled.
	//
	// 	- **FAIL**: The task failed.
	//
	// 	- **INTERRUPT**: The task is interrupted.
	//
	// example:
	//
	// SUCCESS
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The estimated execution progress. The actual progress is subject to the task status.
	//
	// example:
	//
	// 90%
	ProgressRatio *string `json:"ProgressRatio,omitempty" xml:"ProgressRatio,omitempty"`
	// The description of the task status.
	//
	// example:
	//
	// Success
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
}

func (s GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) GoString() string {
	return s.String()
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) GetCleanStrategy() *string {
	return s.CleanStrategy
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) GetCopyChunkMode() *string {
	return s.CopyChunkMode
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) GetCopyChunkSize() *int64 {
	return s.CopyChunkSize
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) GetCopyCount() *int64 {
	return s.CopyCount
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) GetCopyTotal() *int64 {
	return s.CopyTotal
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) GetCutoverFailRetryTimes() *int64 {
	return s.CutoverFailRetryTimes
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) GetCutoverLockTimeSeconds() *int64 {
	return s.CutoverLockTimeSeconds
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) GetCutoverWindowEndTime() *string {
	return s.CutoverWindowEndTime
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) GetCutoverWindowStartTime() *string {
	return s.CutoverWindowStartTime
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) GetDelaySeconds() *int64 {
	return s.DelaySeconds
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) GetProgressRatio() *string {
	return s.ProgressRatio
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) SetCleanStrategy(v string) *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail {
	s.CleanStrategy = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) SetCopyChunkMode(v string) *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail {
	s.CopyChunkMode = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) SetCopyChunkSize(v int64) *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail {
	s.CopyChunkSize = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) SetCopyCount(v int64) *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail {
	s.CopyCount = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) SetCopyTotal(v int64) *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail {
	s.CopyTotal = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) SetCutoverFailRetryTimes(v int64) *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail {
	s.CutoverFailRetryTimes = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) SetCutoverLockTimeSeconds(v int64) *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail {
	s.CutoverLockTimeSeconds = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) SetCutoverWindowEndTime(v string) *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail {
	s.CutoverWindowEndTime = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) SetCutoverWindowStartTime(v string) *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail {
	s.CutoverWindowStartTime = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) SetDelaySeconds(v int64) *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail {
	s.DelaySeconds = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) SetJobStatus(v string) *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail {
	s.JobStatus = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) SetProgressRatio(v string) *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail {
	s.ProgressRatio = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) SetStatusDesc(v string) *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail {
	s.StatusDesc = &v
	return s
}

func (s *GetOnlineDDLProgressResponseBodyOnlineDDLTaskDetail) Validate() error {
	return dara.Validate(s)
}
