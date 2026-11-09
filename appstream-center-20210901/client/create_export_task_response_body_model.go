// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExportTaskModel(v *CreateExportTaskResponseBodyExportTaskModel) *CreateExportTaskResponseBody
	GetExportTaskModel() *CreateExportTaskResponseBodyExportTaskModel
	SetRequestId(v string) *CreateExportTaskResponseBody
	GetRequestId() *string
}

type CreateExportTaskResponseBody struct {
	// The information about the newly created export task. If the task is created, Status is INIT, Progress is 0, and the download URL and error message are empty.
	ExportTaskModel *CreateExportTaskResponseBodyExportTaskModel `json:"ExportTaskModel,omitempty" xml:"ExportTaskModel,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate call records and troubleshoot issues. The example value is partially desensitized.
	//
	// example:
	//
	// 7D30C5D6-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExportTaskResponseBody) GetExportTaskModel() *CreateExportTaskResponseBodyExportTaskModel {
	return s.ExportTaskModel
}

func (s *CreateExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExportTaskResponseBody) SetExportTaskModel(v *CreateExportTaskResponseBodyExportTaskModel) *CreateExportTaskResponseBody {
	s.ExportTaskModel = v
	return s
}

func (s *CreateExportTaskResponseBody) SetRequestId(v string) *CreateExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExportTaskResponseBody) Validate() error {
	if s.ExportTaskModel != nil {
		if err := s.ExportTaskModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateExportTaskResponseBodyExportTaskModel struct {
	// The time when the task was created. The time is in the yyyy-MM-dd HH:mm:ss format, accurate to seconds. The returned string does not contain a time zone offset.
	//
	// example:
	//
	// 2026-07-27 19:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The download URL of the exported file. When the task is created by calling this operation, the file is not generated and this field is empty. The null in the example represents a JSON null value, not the string "null". You can subsequently call DescribeExportTasks to query the latest download URL of the completed task and use it within the validity period.
	//
	// example:
	//
	// null
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The error message when the task execution fails. If the task is created by calling this operation, this field is empty. The null in the example represents a JSON null value, not the string "null". When you query the task status and the status is FAILED, you can use this field to identify the cause of the failure.
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The progress of the export task, in percentage. Valid values: 0 to 100. If the task is created by calling this operation, 0 is returned. The value is updated based on the execution phase and reaches 100 when the task is completed.
	//
	// example:
	//
	// 0
	Progress *float32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of the task. If the task is created by calling this operation, INIT (initialized and waiting for execution) is returned. You can subsequently call DescribeExportTasks to query RUNNING (executing), FINISHED (completed), or FAILED (failed). Do not treat INIT as export completion.
	//
	// example:
	//
	// INIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the newly created export task, which is a 32-character UUID string without hyphens. You can use this value as the TaskId parameter in DescribeExportTasks to query the task.
	//
	// example:
	//
	// 8119078c29e948dba2e0a9917995332a
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the export task, which is the same as the TaskType in the request. For the meaning of each type, see the TaskType request parameter.
	//
	// example:
	//
	// WUYING_SERVER_LIST
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateExportTaskResponseBodyExportTaskModel) String() string {
	return dara.Prettify(s)
}

func (s CreateExportTaskResponseBodyExportTaskModel) GoString() string {
	return s.String()
}

func (s *CreateExportTaskResponseBodyExportTaskModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateExportTaskResponseBodyExportTaskModel) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *CreateExportTaskResponseBodyExportTaskModel) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateExportTaskResponseBodyExportTaskModel) GetProgress() *float32 {
	return s.Progress
}

func (s *CreateExportTaskResponseBodyExportTaskModel) GetStatus() *string {
	return s.Status
}

func (s *CreateExportTaskResponseBodyExportTaskModel) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateExportTaskResponseBodyExportTaskModel) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateExportTaskResponseBodyExportTaskModel) SetCreateTime(v string) *CreateExportTaskResponseBodyExportTaskModel {
	s.CreateTime = &v
	return s
}

func (s *CreateExportTaskResponseBodyExportTaskModel) SetDownloadUrl(v string) *CreateExportTaskResponseBodyExportTaskModel {
	s.DownloadUrl = &v
	return s
}

func (s *CreateExportTaskResponseBodyExportTaskModel) SetErrorMsg(v string) *CreateExportTaskResponseBodyExportTaskModel {
	s.ErrorMsg = &v
	return s
}

func (s *CreateExportTaskResponseBodyExportTaskModel) SetProgress(v float32) *CreateExportTaskResponseBodyExportTaskModel {
	s.Progress = &v
	return s
}

func (s *CreateExportTaskResponseBodyExportTaskModel) SetStatus(v string) *CreateExportTaskResponseBodyExportTaskModel {
	s.Status = &v
	return s
}

func (s *CreateExportTaskResponseBodyExportTaskModel) SetTaskId(v string) *CreateExportTaskResponseBodyExportTaskModel {
	s.TaskId = &v
	return s
}

func (s *CreateExportTaskResponseBodyExportTaskModel) SetTaskType(v string) *CreateExportTaskResponseBodyExportTaskModel {
	s.TaskType = &v
	return s
}

func (s *CreateExportTaskResponseBodyExportTaskModel) Validate() error {
	return dara.Validate(s)
}
