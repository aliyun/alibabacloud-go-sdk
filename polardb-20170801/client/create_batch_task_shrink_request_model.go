// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *CreateBatchTaskShrinkRequest
	GetApplicationType() *string
	SetInstanceIdsShrink(v string) *CreateBatchTaskShrinkRequest
	GetInstanceIdsShrink() *string
	SetInstanceParams(v string) *CreateBatchTaskShrinkRequest
	GetInstanceParams() *string
	SetParam(v string) *CreateBatchTaskShrinkRequest
	GetParam() *string
	SetRegionId(v string) *CreateBatchTaskShrinkRequest
	GetRegionId() *string
	SetTaskName(v string) *CreateBatchTaskShrinkRequest
	GetTaskName() *string
	SetTaskType(v string) *CreateBatchTaskShrinkRequest
	GetTaskType() *string
}

type CreateBatchTaskShrinkRequest struct {
	// The target application type of the batch task.
	//
	// example:
	//
	// polarclaw
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The machine instance IDs.
	//
	// This parameter is required.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The per-instance task parameters in JSON object format (instanceId -> param). This parameter is used in scenarios where multiple instances have different targets, such as different agent workspaces.
	//
	// example:
	//
	// {"pa-xxx1":{"agentId":"test1"},"pa-xxx2":{"scope":"global"}}
	InstanceParams *string `json:"InstanceParams,omitempty" xml:"InstanceParams,omitempty"`
	// The task parameters.
	//
	// example:
	//
	// [{"skillName":"github","version":"1.0.0"},{"skillName":"skill-vetter","version":"1.0.1"}]
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to view the region information of all clusters under the account.
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

func (s CreateBatchTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskShrinkRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *CreateBatchTaskShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *CreateBatchTaskShrinkRequest) GetInstanceParams() *string {
	return s.InstanceParams
}

func (s *CreateBatchTaskShrinkRequest) GetParam() *string {
	return s.Param
}

func (s *CreateBatchTaskShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBatchTaskShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateBatchTaskShrinkRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateBatchTaskShrinkRequest) SetApplicationType(v string) *CreateBatchTaskShrinkRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreateBatchTaskShrinkRequest) SetInstanceIdsShrink(v string) *CreateBatchTaskShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *CreateBatchTaskShrinkRequest) SetInstanceParams(v string) *CreateBatchTaskShrinkRequest {
	s.InstanceParams = &v
	return s
}

func (s *CreateBatchTaskShrinkRequest) SetParam(v string) *CreateBatchTaskShrinkRequest {
	s.Param = &v
	return s
}

func (s *CreateBatchTaskShrinkRequest) SetRegionId(v string) *CreateBatchTaskShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBatchTaskShrinkRequest) SetTaskName(v string) *CreateBatchTaskShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *CreateBatchTaskShrinkRequest) SetTaskType(v string) *CreateBatchTaskShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *CreateBatchTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
