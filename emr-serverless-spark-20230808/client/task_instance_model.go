// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskInstance interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *TaskInstance
	GetBizId() *string
	SetCreator(v int64) *TaskInstance
	GetCreator() *int64
	SetFenixRunId(v string) *TaskInstance
	GetFenixRunId() *string
	SetGmtCreated(v string) *TaskInstance
	GetGmtCreated() *string
	SetTaskBizId(v string) *TaskInstance
	GetTaskBizId() *string
	SetTaskInfo(v *Task) *TaskInstance
	GetTaskInfo() *Task
	SetTaskStatus(v string) *TaskInstance
	GetTaskStatus() *string
	SetWorkspaceBizId(v string) *TaskInstance
	GetWorkspaceBizId() *string
}

type TaskInstance struct {
	// The ID of the folder.
	//
	// example:
	//
	// w-d8********
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// The ID of the user who creates the job.
	//
	// example:
	//
	// 150978934701****
	Creator *int64 `json:"creator,omitempty" xml:"creator,omitempty"`
	// The job run ID.
	//
	// example:
	//
	// jr-93d98d2f7061****
	FenixRunId *string `json:"fenixRunId,omitempty" xml:"fenixRunId,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2024-09-05T02:03:19Z
	GmtCreated *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	// The ID of the data development job.
	//
	// example:
	//
	// TSK-d87******************
	TaskBizId *string `json:"taskBizId,omitempty" xml:"taskBizId,omitempty"`
	// The details of the job.
	TaskInfo *Task `json:"taskInfo,omitempty" xml:"taskInfo,omitempty"`
	// The job run ID.
	//
	// example:
	//
	// Running
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// w-d2d82aa09151****
	WorkspaceBizId *string `json:"workspaceBizId,omitempty" xml:"workspaceBizId,omitempty"`
}

func (s TaskInstance) String() string {
	return dara.Prettify(s)
}

func (s TaskInstance) GoString() string {
	return s.String()
}

func (s *TaskInstance) GetBizId() *string {
	return s.BizId
}

func (s *TaskInstance) GetCreator() *int64 {
	return s.Creator
}

func (s *TaskInstance) GetFenixRunId() *string {
	return s.FenixRunId
}

func (s *TaskInstance) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *TaskInstance) GetTaskBizId() *string {
	return s.TaskBizId
}

func (s *TaskInstance) GetTaskInfo() *Task {
	return s.TaskInfo
}

func (s *TaskInstance) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *TaskInstance) GetWorkspaceBizId() *string {
	return s.WorkspaceBizId
}

func (s *TaskInstance) SetBizId(v string) *TaskInstance {
	s.BizId = &v
	return s
}

func (s *TaskInstance) SetCreator(v int64) *TaskInstance {
	s.Creator = &v
	return s
}

func (s *TaskInstance) SetFenixRunId(v string) *TaskInstance {
	s.FenixRunId = &v
	return s
}

func (s *TaskInstance) SetGmtCreated(v string) *TaskInstance {
	s.GmtCreated = &v
	return s
}

func (s *TaskInstance) SetTaskBizId(v string) *TaskInstance {
	s.TaskBizId = &v
	return s
}

func (s *TaskInstance) SetTaskInfo(v *Task) *TaskInstance {
	s.TaskInfo = v
	return s
}

func (s *TaskInstance) SetTaskStatus(v string) *TaskInstance {
	s.TaskStatus = &v
	return s
}

func (s *TaskInstance) SetWorkspaceBizId(v string) *TaskInstance {
	s.WorkspaceBizId = &v
	return s
}

func (s *TaskInstance) Validate() error {
	if s.TaskInfo != nil {
		if err := s.TaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
