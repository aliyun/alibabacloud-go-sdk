// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *GetTaskResponseBody
	GetConfig() *string
	SetGmtCreateTime(v string) *GetTaskResponseBody
	GetGmtCreateTime() *string
	SetGmtExecutedTime(v string) *GetTaskResponseBody
	GetGmtExecutedTime() *string
	SetGmtFinishedTime(v string) *GetTaskResponseBody
	GetGmtFinishedTime() *string
	SetGmtModifiedTime(v string) *GetTaskResponseBody
	GetGmtModifiedTime() *string
	SetObjectId(v string) *GetTaskResponseBody
	GetObjectId() *string
	SetObjectType(v string) *GetTaskResponseBody
	GetObjectType() *string
	SetProjectId(v string) *GetTaskResponseBody
	GetProjectId() *string
	SetProjectName(v string) *GetTaskResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetTaskResponseBody
	GetRequestId() *string
	SetRunningConfig(v string) *GetTaskResponseBody
	GetRunningConfig() *string
	SetStatus(v string) *GetTaskResponseBody
	GetStatus() *string
	SetType(v string) *GetTaskResponseBody
	GetType() *string
}

type GetTaskResponseBody struct {
	// The task configuration.
	//
	// example:
	//
	// {
	//
	// 	"mode": "overwrite",
	//
	// 	"partitions": {
	//
	// 		"dt": "20230820"
	//
	// 	},
	//
	// 	"event_time": "",
	//
	// 	"config": {},
	//
	// 	"offline_to_online": true
	//
	// }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The execution time.
	//
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtExecutedTime *string `json:"GmtExecutedTime,omitempty" xml:"GmtExecutedTime,omitempty"`
	// The completion time.
	//
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtFinishedTime *string `json:"GmtFinishedTime,omitempty" xml:"GmtFinishedTime,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The ID of the target object.
	//
	// example:
	//
	// 3
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The type of the target object.
	//
	// - ModelFeature: model feature
	//
	// - FeatureView: feature view
	//
	// example:
	//
	// ModelFeature
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project name.
	//
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 37D19490-AB69-567D-A852-407C94E510E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task runtime configuration.
	//
	// example:
	//
	// DROP TABLE IF EXISTS public.fsxxx
	RunningConfig *string `json:"RunningConfig,omitempty" xml:"RunningConfig,omitempty"`
	// The status of the task.
	//
	// - Initializing: The task is initializing.
	//
	// - Running: The task is running.
	//
	// - Success: The task completed successfully.
	//
	// - Failure: The task failed.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task type.
	//
	// - OfflineToOnline: offline-to-online data synchronization
	//
	// - ExportTrainingSet: training sample table export
	//
	// example:
	//
	// OfflineToOnline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetTaskResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetTaskResponseBody) GetGmtExecutedTime() *string {
	return s.GmtExecutedTime
}

func (s *GetTaskResponseBody) GetGmtFinishedTime() *string {
	return s.GmtFinishedTime
}

func (s *GetTaskResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetTaskResponseBody) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetTaskResponseBody) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetTaskResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetTaskResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskResponseBody) GetRunningConfig() *string {
	return s.RunningConfig
}

func (s *GetTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetTaskResponseBody) GetType() *string {
	return s.Type
}

func (s *GetTaskResponseBody) SetConfig(v string) *GetTaskResponseBody {
	s.Config = &v
	return s
}

func (s *GetTaskResponseBody) SetGmtCreateTime(v string) *GetTaskResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetTaskResponseBody) SetGmtExecutedTime(v string) *GetTaskResponseBody {
	s.GmtExecutedTime = &v
	return s
}

func (s *GetTaskResponseBody) SetGmtFinishedTime(v string) *GetTaskResponseBody {
	s.GmtFinishedTime = &v
	return s
}

func (s *GetTaskResponseBody) SetGmtModifiedTime(v string) *GetTaskResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetTaskResponseBody) SetObjectId(v string) *GetTaskResponseBody {
	s.ObjectId = &v
	return s
}

func (s *GetTaskResponseBody) SetObjectType(v string) *GetTaskResponseBody {
	s.ObjectType = &v
	return s
}

func (s *GetTaskResponseBody) SetProjectId(v string) *GetTaskResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetTaskResponseBody) SetProjectName(v string) *GetTaskResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetRunningConfig(v string) *GetTaskResponseBody {
	s.RunningConfig = &v
	return s
}

func (s *GetTaskResponseBody) SetStatus(v string) *GetTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBody) SetType(v string) *GetTaskResponseBody {
	s.Type = &v
	return s
}

func (s *GetTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
