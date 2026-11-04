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
	SetInstanceParams(v string) *CreateBatchTaskRequest
	GetInstanceParams() *string
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
	// The target application type of the batch task.
	//
	// example:
	//
	// polarclaw
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The machine instance IDs.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The per-instance parameter overrides (JSON object: instanceId -> param). This parameter is used in scenarios where multiple instances have different targets, such as different agent workspaces.
	//
	// example:
	//
	// {"pa-xxx1":{"agentId":"test1"},"pa-xxx2":{"scope":"global"}}
	InstanceParams *string `json:"InstanceParams,omitempty" xml:"InstanceParams,omitempty"`
	// The archive parameters.
	//
	// example:
	//
	// [{"skillName":"github","version":"1.0.0"},{"skillName":"skill-vetter","version":"1.0.1"}]
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to view the regions of all clusters under the target account.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the model service.
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

func (s *CreateBatchTaskRequest) GetInstanceParams() *string {
	return s.InstanceParams
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

func (s *CreateBatchTaskRequest) SetInstanceParams(v string) *CreateBatchTaskRequest {
	s.InstanceParams = &v
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
