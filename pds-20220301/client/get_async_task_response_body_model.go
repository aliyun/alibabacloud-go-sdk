// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *GetAsyncTaskResponseBody
	GetAsyncTaskId() *string
	SetCategory(v string) *GetAsyncTaskResponseBody
	GetCategory() *string
	SetConsumedProcess(v int64) *GetAsyncTaskResponseBody
	GetConsumedProcess() *int64
	SetCreatedAt(v string) *GetAsyncTaskResponseBody
	GetCreatedAt() *string
	SetErrCode(v int64) *GetAsyncTaskResponseBody
	GetErrCode() *int64
	SetErrorCode(v string) *GetAsyncTaskResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAsyncTaskResponseBody
	GetErrorMessage() *string
	SetFailedProcess(v int64) *GetAsyncTaskResponseBody
	GetFailedProcess() *int64
	SetFinishedAt(v string) *GetAsyncTaskResponseBody
	GetFinishedAt() *string
	SetMessage(v string) *GetAsyncTaskResponseBody
	GetMessage() *string
	SetSkippedProcess(v int64) *GetAsyncTaskResponseBody
	GetSkippedProcess() *int64
	SetStartedAt(v string) *GetAsyncTaskResponseBody
	GetStartedAt() *string
	SetState(v string) *GetAsyncTaskResponseBody
	GetState() *string
	SetStatus(v string) *GetAsyncTaskResponseBody
	GetStatus() *string
	SetTaskType(v string) *GetAsyncTaskResponseBody
	GetTaskType() *string
	SetTotalProcess(v int64) *GetAsyncTaskResponseBody
	GetTotalProcess() *int64
	SetUncompressFileList(v []*UncompressedFileInfo) *GetAsyncTaskResponseBody
	GetUncompressFileList() []*UncompressedFileInfo
	SetUrl(v string) *GetAsyncTaskResponseBody
	GetUrl() *string
}

type GetAsyncTaskResponseBody struct {
	// The ID of the asynchronous task.
	//
	// example:
	//
	// 000e89fb-cf8f-11e9-8ab4-b6e980803a3b
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	// The custom category of the task.
	//
	// example:
	//
	// album
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The total amount of work that is done in the asynchronous task, such as the number of files that are packaged for package download on the server.
	//
	// example:
	//
	// 100
	ConsumedProcess *int64 `json:"consumed_process,omitempty" xml:"consumed_process,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. Example: 2019-03-28T13:03:29.298Z.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// <warning>This parameter is no longer used. We recommend that you use error_code instead.</warning>
	//
	// The error code returned if the asynchronous task failed.
	//
	// example:
	//
	// InternalError
	ErrCode *int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// The error code returned if the asynchronous task failed.
	//
	// example:
	//
	// InternalError
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// The error message returned if the asynchronous task failed.
	//
	// example:
	//
	// The request has been failed due to some unknown error. Please try again later.
	ErrorMessage  *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	FailedProcess *int64  `json:"failed_process,omitempty" xml:"failed_process,omitempty"`
	// The time when the task was complete. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. Example: 2019-03-28T13:03:29.298Z.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	FinishedAt *string `json:"finished_at,omitempty" xml:"finished_at,omitempty"`
	// <warning>This parameter is no longer used. We recommend that you use error_message instead.</warning>
	//
	// The error message returned if the asynchronous task failed.
	//
	// example:
	//
	// The request has been failed due to some unknown error. Please try again later.
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	SkippedProcess *int64  `json:"skipped_process,omitempty" xml:"skipped_process,omitempty"`
	// The time when the task was started. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. Example: 2019-03-28T13:03:29.298Z.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	StartedAt *string `json:"started_at,omitempty" xml:"started_at,omitempty"`
	// The state of the task. Valid values:
	//
	// 	- Failed
	//
	// 	- Running
	//
	// 	- PartialSucceed
	//
	// 	- Succeed
	//
	// example:
	//
	// Succeed
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// <warning>This parameter is no longer used. We recommend that you use state instead.</warning>
	//
	// The state of the task. Valid values:
	//
	// 	- Failed
	//
	// 	- Running
	//
	// 	- PartialSucceed
	//
	// 	- Succeed
	//
	// example:
	//
	// Succeed
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskType *string `json:"task_type,omitempty" xml:"task_type,omitempty"`
	// The total amount of work to be done in the asynchronous task, such as the number of files to be packaged for package download on the server.
	//
	// example:
	//
	// 1000
	TotalProcess *int64 `json:"total_process,omitempty" xml:"total_process,omitempty"`
	// The extracted files.
	UncompressFileList []*UncompressedFileInfo `json:"uncompress_file_list,omitempty" xml:"uncompress_file_list,omitempty" type:"Repeated"`
	// The download URL of the data generated by the asynchronous task, such as the download URL of the packaged files generated by the task of package download on the server.
	//
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d5b986facbec311ef844c25954f96821497b383%2F5d5b986f955410dd991646bb87c6b4e899eff525?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAsyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskResponseBody) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *GetAsyncTaskResponseBody) GetCategory() *string {
	return s.Category
}

func (s *GetAsyncTaskResponseBody) GetConsumedProcess() *int64 {
	return s.ConsumedProcess
}

func (s *GetAsyncTaskResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetAsyncTaskResponseBody) GetErrCode() *int64 {
	return s.ErrCode
}

func (s *GetAsyncTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAsyncTaskResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAsyncTaskResponseBody) GetFailedProcess() *int64 {
	return s.FailedProcess
}

func (s *GetAsyncTaskResponseBody) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetAsyncTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAsyncTaskResponseBody) GetSkippedProcess() *int64 {
	return s.SkippedProcess
}

func (s *GetAsyncTaskResponseBody) GetStartedAt() *string {
	return s.StartedAt
}

func (s *GetAsyncTaskResponseBody) GetState() *string {
	return s.State
}

func (s *GetAsyncTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetAsyncTaskResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *GetAsyncTaskResponseBody) GetTotalProcess() *int64 {
	return s.TotalProcess
}

func (s *GetAsyncTaskResponseBody) GetUncompressFileList() []*UncompressedFileInfo {
	return s.UncompressFileList
}

func (s *GetAsyncTaskResponseBody) GetUrl() *string {
	return s.Url
}

func (s *GetAsyncTaskResponseBody) SetAsyncTaskId(v string) *GetAsyncTaskResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetCategory(v string) *GetAsyncTaskResponseBody {
	s.Category = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetConsumedProcess(v int64) *GetAsyncTaskResponseBody {
	s.ConsumedProcess = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetCreatedAt(v string) *GetAsyncTaskResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetErrCode(v int64) *GetAsyncTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetErrorCode(v string) *GetAsyncTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetErrorMessage(v string) *GetAsyncTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetFailedProcess(v int64) *GetAsyncTaskResponseBody {
	s.FailedProcess = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetFinishedAt(v string) *GetAsyncTaskResponseBody {
	s.FinishedAt = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetMessage(v string) *GetAsyncTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetSkippedProcess(v int64) *GetAsyncTaskResponseBody {
	s.SkippedProcess = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetStartedAt(v string) *GetAsyncTaskResponseBody {
	s.StartedAt = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetState(v string) *GetAsyncTaskResponseBody {
	s.State = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetStatus(v string) *GetAsyncTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetTaskType(v string) *GetAsyncTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetTotalProcess(v int64) *GetAsyncTaskResponseBody {
	s.TotalProcess = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetUncompressFileList(v []*UncompressedFileInfo) *GetAsyncTaskResponseBody {
	s.UncompressFileList = v
	return s
}

func (s *GetAsyncTaskResponseBody) SetUrl(v string) *GetAsyncTaskResponseBody {
	s.Url = &v
	return s
}

func (s *GetAsyncTaskResponseBody) Validate() error {
	if s.UncompressFileList != nil {
		for _, item := range s.UncompressFileList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
