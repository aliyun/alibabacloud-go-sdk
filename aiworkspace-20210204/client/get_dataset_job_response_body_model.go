// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompletedFileCount(v int64) *GetDatasetJobResponseBody
	GetCompletedFileCount() *int64
	SetCreateTime(v string) *GetDatasetJobResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetDatasetJobResponseBody
	GetDescription() *string
	SetFailedFileCount(v int64) *GetDatasetJobResponseBody
	GetFailedFileCount() *int64
	SetFinishTime(v string) *GetDatasetJobResponseBody
	GetFinishTime() *string
	SetJobAction(v string) *GetDatasetJobResponseBody
	GetJobAction() *string
	SetJobMode(v string) *GetDatasetJobResponseBody
	GetJobMode() *string
	SetJobSpec(v string) *GetDatasetJobResponseBody
	GetJobSpec() *string
	SetLogs(v []*string) *GetDatasetJobResponseBody
	GetLogs() []*string
	SetRequestId(v string) *GetDatasetJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetDatasetJobResponseBody
	GetStatus() *string
	SetTotalFileCount(v int64) *GetDatasetJobResponseBody
	GetTotalFileCount() *int64
}

type GetDatasetJobResponseBody struct {
	// The total number of completed files.
	//
	// example:
	//
	// 990
	CompletedFileCount *int64 `json:"CompletedFileCount,omitempty" xml:"CompletedFileCount,omitempty"`
	// The time when the job is started.
	//
	// example:
	//
	// 2024-11-15T07:06:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The job description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The total number of failed files.
	//
	// example:
	//
	// 10
	FailedFileCount *int64 `json:"FailedFileCount,omitempty" xml:"FailedFileCount,omitempty"`
	// The time when the job ends.
	//
	// example:
	//
	// 2024-07-16T02:03:23Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The action that is performed on the job.
	//
	// Valid values:
	//
	// 	- SemanticIndex: semantic indexing
	//
	// 	- IntelligentTag: smart labeling
	//
	// 	- FileMetaExport: metadata export
	//
	// example:
	//
	// SemanticIndex
	JobAction *string `json:"JobAction,omitempty" xml:"JobAction,omitempty"`
	// The job mode.
	//
	// Valid value:
	//
	// 	- Full: full data mode.
	//
	// example:
	//
	// Full
	JobMode *string `json:"JobMode,omitempty" xml:"JobMode,omitempty"`
	// The job details.
	//
	// example:
	//
	// {\\"modelId\\":\\"xxx\\"}
	JobSpec *string `json:"JobSpec,omitempty" xml:"JobSpec,omitempty"`
	// The job logs.
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 64B50C1D-D4C2-560C-86A3-A6ED****16D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job state.
	//
	// Valid values:
	//
	// 	- Succeeded
	//
	// 	- Failed
	//
	// 	- Running
	//
	// 	- Pending
	//
	// 	- PartialFailed
	//
	// 	- Deleting
	//
	// 	- ManuallyStop
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of job files.
	//
	// example:
	//
	// 1000
	TotalFileCount *int64 `json:"TotalFileCount,omitempty" xml:"TotalFileCount,omitempty"`
}

func (s GetDatasetJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetJobResponseBody) GetCompletedFileCount() *int64 {
	return s.CompletedFileCount
}

func (s *GetDatasetJobResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDatasetJobResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetDatasetJobResponseBody) GetFailedFileCount() *int64 {
	return s.FailedFileCount
}

func (s *GetDatasetJobResponseBody) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetDatasetJobResponseBody) GetJobAction() *string {
	return s.JobAction
}

func (s *GetDatasetJobResponseBody) GetJobMode() *string {
	return s.JobMode
}

func (s *GetDatasetJobResponseBody) GetJobSpec() *string {
	return s.JobSpec
}

func (s *GetDatasetJobResponseBody) GetLogs() []*string {
	return s.Logs
}

func (s *GetDatasetJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasetJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetDatasetJobResponseBody) GetTotalFileCount() *int64 {
	return s.TotalFileCount
}

func (s *GetDatasetJobResponseBody) SetCompletedFileCount(v int64) *GetDatasetJobResponseBody {
	s.CompletedFileCount = &v
	return s
}

func (s *GetDatasetJobResponseBody) SetCreateTime(v string) *GetDatasetJobResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetDatasetJobResponseBody) SetDescription(v string) *GetDatasetJobResponseBody {
	s.Description = &v
	return s
}

func (s *GetDatasetJobResponseBody) SetFailedFileCount(v int64) *GetDatasetJobResponseBody {
	s.FailedFileCount = &v
	return s
}

func (s *GetDatasetJobResponseBody) SetFinishTime(v string) *GetDatasetJobResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetDatasetJobResponseBody) SetJobAction(v string) *GetDatasetJobResponseBody {
	s.JobAction = &v
	return s
}

func (s *GetDatasetJobResponseBody) SetJobMode(v string) *GetDatasetJobResponseBody {
	s.JobMode = &v
	return s
}

func (s *GetDatasetJobResponseBody) SetJobSpec(v string) *GetDatasetJobResponseBody {
	s.JobSpec = &v
	return s
}

func (s *GetDatasetJobResponseBody) SetLogs(v []*string) *GetDatasetJobResponseBody {
	s.Logs = v
	return s
}

func (s *GetDatasetJobResponseBody) SetRequestId(v string) *GetDatasetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetJobResponseBody) SetStatus(v string) *GetDatasetJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetDatasetJobResponseBody) SetTotalFileCount(v int64) *GetDatasetJobResponseBody {
	s.TotalFileCount = &v
	return s
}

func (s *GetDatasetJobResponseBody) Validate() error {
	return dara.Validate(s)
}
