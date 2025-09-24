// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetJob interface {
	dara.Model
	String() string
	GoString() string
	SetCompletedFileCount(v int64) *DatasetJob
	GetCompletedFileCount() *int64
	SetCreateTime(v string) *DatasetJob
	GetCreateTime() *string
	SetDatasetJobId(v string) *DatasetJob
	GetDatasetJobId() *string
	SetDatasetVersion(v string) *DatasetJob
	GetDatasetVersion() *string
	SetDescription(v string) *DatasetJob
	GetDescription() *string
	SetFailedFileCount(v int64) *DatasetJob
	GetFailedFileCount() *int64
	SetFinishTime(v string) *DatasetJob
	GetFinishTime() *string
	SetJobAction(v string) *DatasetJob
	GetJobAction() *string
	SetJobMode(v string) *DatasetJob
	GetJobMode() *string
	SetJobSpec(v string) *DatasetJob
	GetJobSpec() *string
	SetLogs(v []*string) *DatasetJob
	GetLogs() []*string
	SetStatus(v string) *DatasetJob
	GetStatus() *string
	SetTotalFileCount(v int64) *DatasetJob
	GetTotalFileCount() *int64
	SetWorkspaceId(v string) *DatasetJob
	GetWorkspaceId() *string
}

type DatasetJob struct {
	CompletedFileCount *int64    `json:"CompletedFileCount,omitempty" xml:"CompletedFileCount,omitempty"`
	CreateTime         *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatasetJobId       *string   `json:"DatasetJobId,omitempty" xml:"DatasetJobId,omitempty"`
	DatasetVersion     *string   `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	Description        *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	FailedFileCount    *int64    `json:"FailedFileCount,omitempty" xml:"FailedFileCount,omitempty"`
	FinishTime         *string   `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	JobAction          *string   `json:"JobAction,omitempty" xml:"JobAction,omitempty"`
	JobMode            *string   `json:"JobMode,omitempty" xml:"JobMode,omitempty"`
	JobSpec            *string   `json:"JobSpec,omitempty" xml:"JobSpec,omitempty"`
	Logs               []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	Status             *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalFileCount     *int64    `json:"TotalFileCount,omitempty" xml:"TotalFileCount,omitempty"`
	WorkspaceId        *string   `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DatasetJob) String() string {
	return dara.Prettify(s)
}

func (s DatasetJob) GoString() string {
	return s.String()
}

func (s *DatasetJob) GetCompletedFileCount() *int64 {
	return s.CompletedFileCount
}

func (s *DatasetJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DatasetJob) GetDatasetJobId() *string {
	return s.DatasetJobId
}

func (s *DatasetJob) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *DatasetJob) GetDescription() *string {
	return s.Description
}

func (s *DatasetJob) GetFailedFileCount() *int64 {
	return s.FailedFileCount
}

func (s *DatasetJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DatasetJob) GetJobAction() *string {
	return s.JobAction
}

func (s *DatasetJob) GetJobMode() *string {
	return s.JobMode
}

func (s *DatasetJob) GetJobSpec() *string {
	return s.JobSpec
}

func (s *DatasetJob) GetLogs() []*string {
	return s.Logs
}

func (s *DatasetJob) GetStatus() *string {
	return s.Status
}

func (s *DatasetJob) GetTotalFileCount() *int64 {
	return s.TotalFileCount
}

func (s *DatasetJob) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DatasetJob) SetCompletedFileCount(v int64) *DatasetJob {
	s.CompletedFileCount = &v
	return s
}

func (s *DatasetJob) SetCreateTime(v string) *DatasetJob {
	s.CreateTime = &v
	return s
}

func (s *DatasetJob) SetDatasetJobId(v string) *DatasetJob {
	s.DatasetJobId = &v
	return s
}

func (s *DatasetJob) SetDatasetVersion(v string) *DatasetJob {
	s.DatasetVersion = &v
	return s
}

func (s *DatasetJob) SetDescription(v string) *DatasetJob {
	s.Description = &v
	return s
}

func (s *DatasetJob) SetFailedFileCount(v int64) *DatasetJob {
	s.FailedFileCount = &v
	return s
}

func (s *DatasetJob) SetFinishTime(v string) *DatasetJob {
	s.FinishTime = &v
	return s
}

func (s *DatasetJob) SetJobAction(v string) *DatasetJob {
	s.JobAction = &v
	return s
}

func (s *DatasetJob) SetJobMode(v string) *DatasetJob {
	s.JobMode = &v
	return s
}

func (s *DatasetJob) SetJobSpec(v string) *DatasetJob {
	s.JobSpec = &v
	return s
}

func (s *DatasetJob) SetLogs(v []*string) *DatasetJob {
	s.Logs = v
	return s
}

func (s *DatasetJob) SetStatus(v string) *DatasetJob {
	s.Status = &v
	return s
}

func (s *DatasetJob) SetTotalFileCount(v int64) *DatasetJob {
	s.TotalFileCount = &v
	return s
}

func (s *DatasetJob) SetWorkspaceId(v string) *DatasetJob {
	s.WorkspaceId = &v
	return s
}

func (s *DatasetJob) Validate() error {
	return dara.Validate(s)
}
