// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescriptionShrink(v string) *UpdateScheduledTaskShrinkRequest
	GetDescriptionShrink() *string
	SetDigitalEmployeeNameShrink(v string) *UpdateScheduledTaskShrinkRequest
	GetDigitalEmployeeNameShrink() *string
	SetIsOpen(v bool) *UpdateScheduledTaskShrinkRequest
	GetIsOpen() *bool
	SetModel(v string) *UpdateScheduledTaskShrinkRequest
	GetModel() *string
	SetName(v string) *UpdateScheduledTaskShrinkRequest
	GetName() *string
	SetSegmentsShrink(v string) *UpdateScheduledTaskShrinkRequest
	GetSegmentsShrink() *string
	SetTaskDetailShrink(v string) *UpdateScheduledTaskShrinkRequest
	GetTaskDetailShrink() *string
	SetTaskId(v string) *UpdateScheduledTaskShrinkRequest
	GetTaskId() *string
	SetTenantId(v string) *UpdateScheduledTaskShrinkRequest
	GetTenantId() *string
	SetTriggerConfigShrink(v string) *UpdateScheduledTaskShrinkRequest
	GetTriggerConfigShrink() *string
}

type UpdateScheduledTaskShrinkRequest struct {
	DescriptionShrink *string `json:"description,omitempty" xml:"description,omitempty"`
	// 数字员工名称列表
	//
	// example:
	//
	// string_value
	DigitalEmployeeNameShrink *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// 是否公开访问
	//
	// example:
	//
	// true
	IsOpen *bool `json:"isOpen,omitempty" xml:"isOpen,omitempty"`
	// 执行模型档位；不传则不更新
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	SegmentsShrink   *string `json:"segments,omitempty" xml:"segments,omitempty"`
	TaskDetailShrink *string `json:"taskDetail,omitempty" xml:"taskDetail,omitempty"`
	// 任务 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId            *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	TriggerConfigShrink *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
}

func (s UpdateScheduledTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduledTaskShrinkRequest) GetDescriptionShrink() *string {
	return s.DescriptionShrink
}

func (s *UpdateScheduledTaskShrinkRequest) GetDigitalEmployeeNameShrink() *string {
	return s.DigitalEmployeeNameShrink
}

func (s *UpdateScheduledTaskShrinkRequest) GetIsOpen() *bool {
	return s.IsOpen
}

func (s *UpdateScheduledTaskShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *UpdateScheduledTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateScheduledTaskShrinkRequest) GetSegmentsShrink() *string {
	return s.SegmentsShrink
}

func (s *UpdateScheduledTaskShrinkRequest) GetTaskDetailShrink() *string {
	return s.TaskDetailShrink
}

func (s *UpdateScheduledTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateScheduledTaskShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateScheduledTaskShrinkRequest) GetTriggerConfigShrink() *string {
	return s.TriggerConfigShrink
}

func (s *UpdateScheduledTaskShrinkRequest) SetDescriptionShrink(v string) *UpdateScheduledTaskShrinkRequest {
	s.DescriptionShrink = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetDigitalEmployeeNameShrink(v string) *UpdateScheduledTaskShrinkRequest {
	s.DigitalEmployeeNameShrink = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetIsOpen(v bool) *UpdateScheduledTaskShrinkRequest {
	s.IsOpen = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetModel(v string) *UpdateScheduledTaskShrinkRequest {
	s.Model = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetName(v string) *UpdateScheduledTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetSegmentsShrink(v string) *UpdateScheduledTaskShrinkRequest {
	s.SegmentsShrink = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetTaskDetailShrink(v string) *UpdateScheduledTaskShrinkRequest {
	s.TaskDetailShrink = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetTaskId(v string) *UpdateScheduledTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetTenantId(v string) *UpdateScheduledTaskShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetTriggerConfigShrink(v string) *UpdateScheduledTaskShrinkRequest {
	s.TriggerConfigShrink = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
