// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeExportTasksResponseBody
	GetCount() *int32
	SetExportTaskModel(v []*DescribeExportTasksResponseBodyExportTaskModel) *DescribeExportTasksResponseBody
	GetExportTaskModel() []*DescribeExportTasksResponseBodyExportTaskModel
	SetPageNumber(v int32) *DescribeExportTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeExportTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeExportTasksResponseBody
	GetRequestId() *string
}

type DescribeExportTasksResponseBody struct {
	// The total number of tasks that meet the query conditions, regardless of the number of tasks returned on the current page. You can use this value and PageSize to calculate the total number of pages. If no tasks meet the conditions, 0 is returned.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of export tasks, sorted in descending order of task creation time, with the most recently created tasks listed first. If no tasks meet the conditions, an empty array is returned.
	ExportTaskModel []*DescribeExportTasksResponseBodyExportTaskModel `json:"ExportTaskModel,omitempty" xml:"ExportTaskModel,omitempty" type:"Repeated"`
	// The actual page number used in this query, which is the echoed value of the PageNumber request parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The actual number of tasks per page used in this query, which is the echoed value of the PageSize request parameter.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate call records and troubleshoot issues. The example value is partially desensitized.
	//
	// example:
	//
	// D4A13D91-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExportTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExportTasksResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeExportTasksResponseBody) GetExportTaskModel() []*DescribeExportTasksResponseBodyExportTaskModel {
	return s.ExportTaskModel
}

func (s *DescribeExportTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeExportTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExportTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExportTasksResponseBody) SetCount(v int32) *DescribeExportTasksResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeExportTasksResponseBody) SetExportTaskModel(v []*DescribeExportTasksResponseBodyExportTaskModel) *DescribeExportTasksResponseBody {
	s.ExportTaskModel = v
	return s
}

func (s *DescribeExportTasksResponseBody) SetPageNumber(v int32) *DescribeExportTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeExportTasksResponseBody) SetPageSize(v int32) *DescribeExportTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeExportTasksResponseBody) SetRequestId(v string) *DescribeExportTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExportTasksResponseBody) Validate() error {
	if s.ExportTaskModel != nil {
		for _, item := range s.ExportTaskModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeExportTasksResponseBodyExportTaskModel struct {
	// The creation time of the task, in yyyy-MM-dd HH:mm:ss format, accurate to seconds. The returned string does not include the time zone offset.
	//
	// example:
	//
	// 2026-07-27 19:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The OSS pre-signed download URL of the export file, which is returned only for tasks whose Status is FINISHED. The default maximum validity period of the signature is 1 hour and does not exceed the 24-hour validity period after the task is created. When the remaining validity period is less than 10 minutes, the server automatically refreshes and returns a new URL. If the task is not completed, the task has expired, or the URL refresh fails, this field is empty. The domain name and path in the example URL are for demonstration purposes. The actual returned value is a complete URL with signature parameters. Use the URL returned by the most recent query for downloading.
	//
	// example:
	//
	// https://example.oss-cn-hangzhou.aliyuncs.com/...
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The cause of the task execution failure, which is returned only when Status is FAILED. The content is the English error description recorded by the server and can be directly displayed to users or used for troubleshooting. This field is empty in other states.
	//
	// example:
	//
	// The export sub type is not supported.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The name of the export file, which consists of the task type, task creation time (in yyyyMMddHHmmss format), and the first 8 characters of the task ID, connected by underscores. The file extension is .xlsx. If the file is not generated, this field is empty.
	//
	// example:
	//
	// WUYING_SERVER_LIST_20260727190000_8119078c.xlsx
	ExportFileName *string `json:"ExportFileName,omitempty" xml:"ExportFileName,omitempty"`
	// The progress of the export task, in percentage. Valid values: 0 to 100. The value is updated based on the execution stage and reaches 100 when the task is completed.
	//
	// example:
	//
	// 100
	Progress *float32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of the export task. This operation returns the following valid values:
	//
	// - INIT: The task is created and waiting to be executed. Continue polling.
	//
	// - RUNNING: The export file is being generated. You can use Progress to display the progress and continue polling.
	//
	// - FINISHED: The export file is generated. Stop polling and use DownloadUrl to download the file.
	//
	// - FAILED: The task execution failed. Stop polling and check ErrorMsg.
	//
	// The server also has an internal EXPIRED state, but this operation does not return tasks in this state.
	//
	// example:
	//
	// FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the export task, which is a 32-character UUID string without hyphens. You can use this value as the TaskId parameter to poll this task.
	//
	// example:
	//
	// 8119078c29e948dba2e0a9917995332a
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the export task. For the meaning of each value, see the TaskType request parameter.
	//
	// example:
	//
	// WUYING_SERVER_LIST
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeExportTasksResponseBodyExportTaskModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportTasksResponseBodyExportTaskModel) GoString() string {
	return s.String()
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) GetExportFileName() *string {
	return s.ExportFileName
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) GetProgress() *float32 {
	return s.Progress
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) GetStatus() *string {
	return s.Status
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) SetCreateTime(v string) *DescribeExportTasksResponseBodyExportTaskModel {
	s.CreateTime = &v
	return s
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) SetDownloadUrl(v string) *DescribeExportTasksResponseBodyExportTaskModel {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) SetErrorMsg(v string) *DescribeExportTasksResponseBodyExportTaskModel {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) SetExportFileName(v string) *DescribeExportTasksResponseBodyExportTaskModel {
	s.ExportFileName = &v
	return s
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) SetProgress(v float32) *DescribeExportTasksResponseBodyExportTaskModel {
	s.Progress = &v
	return s
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) SetStatus(v string) *DescribeExportTasksResponseBodyExportTaskModel {
	s.Status = &v
	return s
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) SetTaskId(v string) *DescribeExportTasksResponseBodyExportTaskModel {
	s.TaskId = &v
	return s
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) SetTaskType(v string) *DescribeExportTasksResponseBodyExportTaskModel {
	s.TaskType = &v
	return s
}

func (s *DescribeExportTasksResponseBodyExportTaskModel) Validate() error {
	return dara.Validate(s)
}
