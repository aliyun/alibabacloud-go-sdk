// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTasksResponseBody
	GetRequestId() *string
	SetTasks(v []*ListTasksResponseBodyTasks) *ListTasksResponseBody
	GetTasks() []*ListTasksResponseBodyTasks
	SetTotalCount(v int32) *ListTasksResponseBody
	GetTotalCount() *int32
}

type ListTasksResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C33E160C-BFCA-5719-B958-942850E949F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of tasks.
	Tasks []*ListTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of tasks.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTasksResponseBody) GetTasks() []*ListTasksResponseBodyTasks {
	return s.Tasks
}

func (s *ListTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTasks(v []*ListTasksResponseBodyTasks) *ListTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTasksResponseBody) SetTotalCount(v int32) *ListTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTasksResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTasksResponseBodyTasks struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the task was executed.
	//
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtExecutedTime *string `json:"GmtExecutedTime,omitempty" xml:"GmtExecutedTime,omitempty"`
	// The time when the task was completed.
	//
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtFinishedTime *string `json:"GmtFinishedTime,omitempty" xml:"GmtFinishedTime,omitempty"`
	// The ID of the object.
	//
	// example:
	//
	// 5
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The type of the object. Valid values:
	//
	// ● ModelFeature: a model feature.
	//
	// ● FeatureView: a feature view.
	//
	// example:
	//
	// ModelFeature
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 4
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project name.
	//
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The status of the task. Valid values:
	//
	// ● Initializing: The task is being initialized.
	//
	// ● Running: The task is in progress.
	//
	// ● Success: The task is successful.
	//
	// ● Failure: The task failed.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 4
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. Valid values:
	//
	// ● OfflineToOnline: The task synchronizes data from offline to online.
	//
	// ● ExportTrainingSet: The task exports a training set.
	//
	// example:
	//
	// OfflineToOnline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasks) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListTasksResponseBodyTasks) GetGmtExecutedTime() *string {
	return s.GmtExecutedTime
}

func (s *ListTasksResponseBodyTasks) GetGmtFinishedTime() *string {
	return s.GmtFinishedTime
}

func (s *ListTasksResponseBodyTasks) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListTasksResponseBodyTasks) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListTasksResponseBodyTasks) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListTasksResponseBodyTasks) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListTasksResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *ListTasksResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTasksResponseBodyTasks) GetType() *string {
	return s.Type
}

func (s *ListTasksResponseBodyTasks) SetGmtCreateTime(v string) *ListTasksResponseBodyTasks {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetGmtExecutedTime(v string) *ListTasksResponseBodyTasks {
	s.GmtExecutedTime = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetGmtFinishedTime(v string) *ListTasksResponseBodyTasks {
	s.GmtFinishedTime = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetObjectId(v string) *ListTasksResponseBodyTasks {
	s.ObjectId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetObjectType(v string) *ListTasksResponseBodyTasks {
	s.ObjectType = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetProjectId(v string) *ListTasksResponseBodyTasks {
	s.ProjectId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetProjectName(v string) *ListTasksResponseBodyTasks {
	s.ProjectName = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetStatus(v string) *ListTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTaskId(v string) *ListTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetType(v string) *ListTasksResponseBodyTasks {
	s.Type = &v
	return s
}

func (s *ListTasksResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
