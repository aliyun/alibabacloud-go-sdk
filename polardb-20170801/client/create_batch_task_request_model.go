// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *CreateBatchTaskRequest
	GetApplicationType() *string
	SetInstanceIds(v []*string) *CreateBatchTaskRequest
	GetInstanceIds() []*string
	SetParam(v string) *CreateBatchTaskRequest
	GetParam() *string
	SetRegionId(v string) *CreateBatchTaskRequest
	GetRegionId() *string
	SetTaskName(v string) *CreateBatchTaskRequest
	GetTaskName() *string
	SetTaskType(v string) *CreateBatchTaskRequest
	GetTaskType() *string
}

type CreateBatchTaskRequest struct {
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The instance IDs.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The task parameters.
	//
	// example:
	//
	// [{"skillName":"github","version":"1.0.0"},{"skillName":"skill-vetter","version":"1.0.1"}]
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The region ID.
	//
	// > Call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the regions of all clusters in your account.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the batch task.
	//
	// This parameter is required.
	//
	// example:
	//
	// batch_task_test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task type.
	//
	// This parameter is required.
	//
	// example:
	//
	// polarclaw_install_skills
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateBatchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *CreateBatchTaskRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateBatchTaskRequest) GetParam() *string {
	return s.Param
}

func (s *CreateBatchTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBatchTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateBatchTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateBatchTaskRequest) SetApplicationType(v string) *CreateBatchTaskRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreateBatchTaskRequest) SetInstanceIds(v []*string) *CreateBatchTaskRequest {
	s.InstanceIds = v
	return s
}

func (s *CreateBatchTaskRequest) SetParam(v string) *CreateBatchTaskRequest {
	s.Param = &v
	return s
}

func (s *CreateBatchTaskRequest) SetRegionId(v string) *CreateBatchTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBatchTaskRequest) SetTaskName(v string) *CreateBatchTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateBatchTaskRequest) SetTaskType(v string) *CreateBatchTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateBatchTaskRequest) Validate() error {
	return dara.Validate(s)
}
