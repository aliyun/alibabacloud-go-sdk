// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFormationTaskByIDRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *QueryFormationTaskByIDRequest
	GetDBClusterId() *string
	SetRegionId(v string) *QueryFormationTaskByIDRequest
	GetRegionId() *string
	SetTaskId(v int64) *QueryFormationTaskByIDRequest
	GetTaskId() *int64
	SetTaskType(v string) *QueryFormationTaskByIDRequest
	GetTaskType() *string
}

type QueryFormationTaskByIDRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1ub9grke1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 36
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["CRAWLER"]
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s QueryFormationTaskByIDRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFormationTaskByIDRequest) GoString() string {
	return s.String()
}

func (s *QueryFormationTaskByIDRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *QueryFormationTaskByIDRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryFormationTaskByIDRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QueryFormationTaskByIDRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryFormationTaskByIDRequest) SetDBClusterId(v string) *QueryFormationTaskByIDRequest {
	s.DBClusterId = &v
	return s
}

func (s *QueryFormationTaskByIDRequest) SetRegionId(v string) *QueryFormationTaskByIDRequest {
	s.RegionId = &v
	return s
}

func (s *QueryFormationTaskByIDRequest) SetTaskId(v int64) *QueryFormationTaskByIDRequest {
	s.TaskId = &v
	return s
}

func (s *QueryFormationTaskByIDRequest) SetTaskType(v string) *QueryFormationTaskByIDRequest {
	s.TaskType = &v
	return s
}

func (s *QueryFormationTaskByIDRequest) Validate() error {
	return dara.Validate(s)
}
