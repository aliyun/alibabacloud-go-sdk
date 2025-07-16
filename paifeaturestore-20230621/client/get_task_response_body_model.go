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
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtExecutedTime *string `json:"GmtExecutedTime,omitempty" xml:"GmtExecutedTime,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtFinishedTime *string `json:"GmtFinishedTime,omitempty" xml:"GmtFinishedTime,omitempty"`
	// example:
	//
	// 2023-07-04T11:26:09.036+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 3
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// example:
	//
	// ModelFeature
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 37D19490-AB69-567D-A852-407C94E510E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// DROP TABLE IF EXISTS public.fsxxx
	RunningConfig *string `json:"RunningConfig,omitempty" xml:"RunningConfig,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
