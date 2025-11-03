// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEcdReportTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExportTaskList(v []*DescribeEcdReportTasksResponseBodyExportTaskList) *DescribeEcdReportTasksResponseBody
	GetExportTaskList() []*DescribeEcdReportTasksResponseBodyExportTaskList
	SetRequestId(v string) *DescribeEcdReportTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeEcdReportTasksResponseBody
	GetTotalCount() *int64
}

type DescribeEcdReportTasksResponseBody struct {
	// The report export tasks.
	ExportTaskList []*DescribeEcdReportTasksResponseBodyExportTaskList `json:"ExportTaskList,omitempty" xml:"ExportTaskList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEcdReportTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEcdReportTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEcdReportTasksResponseBody) GetExportTaskList() []*DescribeEcdReportTasksResponseBodyExportTaskList {
	return s.ExportTaskList
}

func (s *DescribeEcdReportTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEcdReportTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeEcdReportTasksResponseBody) SetExportTaskList(v []*DescribeEcdReportTasksResponseBodyExportTaskList) *DescribeEcdReportTasksResponseBody {
	s.ExportTaskList = v
	return s
}

func (s *DescribeEcdReportTasksResponseBody) SetRequestId(v string) *DescribeEcdReportTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEcdReportTasksResponseBody) SetTotalCount(v int64) *DescribeEcdReportTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEcdReportTasksResponseBody) Validate() error {
	if s.ExportTaskList != nil {
		for _, item := range s.ExportTaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEcdReportTasksResponseBodyExportTaskList struct {
	// The download URL of the report file.
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// No Data.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2025-07-14T07:46:49.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the task was last modified.
	//
	// example:
	//
	// 2025-07-14T07:46:49.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The progress of the report export task. Unit: %.
	//
	// example:
	//
	// 80
	Progress *float32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The name of the report file.
	//
	// example:
	//
	// TestFileName
	ReportFileName *string `json:"ReportFileName,omitempty" xml:"ReportFileName,omitempty"`
	// The task status.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sub-type of the report export task.
	//
	// Valid value:
	//
	// 	- DESKTOP: cloud computer
	//
	// example:
	//
	// DESKTOP
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The ID of the report export task.
	//
	// example:
	//
	// ret-asdfkjg*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the report.
	//
	// Valid value:
	//
	// 	- RESOURCE_REPORT
	//
	// example:
	//
	// RESOURCE_REPORT
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeEcdReportTasksResponseBodyExportTaskList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEcdReportTasksResponseBodyExportTaskList) GoString() string {
	return s.String()
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) GetProgress() *float32 {
	return s.Progress
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) GetReportFileName() *string {
	return s.ReportFileName
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) GetStatus() *string {
	return s.Status
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) GetSubType() *string {
	return s.SubType
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) SetDownloadUrl(v string) *DescribeEcdReportTasksResponseBodyExportTaskList {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) SetErrorCode(v string) *DescribeEcdReportTasksResponseBodyExportTaskList {
	s.ErrorCode = &v
	return s
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) SetErrorMsg(v string) *DescribeEcdReportTasksResponseBodyExportTaskList {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) SetGmtCreate(v string) *DescribeEcdReportTasksResponseBodyExportTaskList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) SetGmtModified(v string) *DescribeEcdReportTasksResponseBodyExportTaskList {
	s.GmtModified = &v
	return s
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) SetProgress(v float32) *DescribeEcdReportTasksResponseBodyExportTaskList {
	s.Progress = &v
	return s
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) SetReportFileName(v string) *DescribeEcdReportTasksResponseBodyExportTaskList {
	s.ReportFileName = &v
	return s
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) SetStatus(v string) *DescribeEcdReportTasksResponseBodyExportTaskList {
	s.Status = &v
	return s
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) SetSubType(v string) *DescribeEcdReportTasksResponseBodyExportTaskList {
	s.SubType = &v
	return s
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) SetTaskId(v string) *DescribeEcdReportTasksResponseBodyExportTaskList {
	s.TaskId = &v
	return s
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) SetTaskType(v string) *DescribeEcdReportTasksResponseBodyExportTaskList {
	s.TaskType = &v
	return s
}

func (s *DescribeEcdReportTasksResponseBodyExportTaskList) Validate() error {
	return dara.Validate(s)
}
