// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *CreateBatchTaskShrinkRequest
	GetInstanceIdsShrink() *string
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
	// This parameter is required.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// [{"skillName":"github","version":"1.0.0"},{"skillName":"skill-vetter","version":"1.0.1"}]
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// batch_task_test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
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

func (s *CreateBatchTaskShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
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

func (s *CreateBatchTaskShrinkRequest) SetInstanceIdsShrink(v string) *CreateBatchTaskShrinkRequest {
	s.InstanceIdsShrink = &v
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
