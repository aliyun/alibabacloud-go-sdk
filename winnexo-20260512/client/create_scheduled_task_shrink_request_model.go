// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *CreateScheduledTaskShrinkRequest
	GetCollaborationGroupId() *string
	SetDescriptionShrink(v string) *CreateScheduledTaskShrinkRequest
	GetDescriptionShrink() *string
	SetDigitalEmployeeNameShrink(v string) *CreateScheduledTaskShrinkRequest
	GetDigitalEmployeeNameShrink() *string
	SetIsOpen(v bool) *CreateScheduledTaskShrinkRequest
	GetIsOpen() *bool
	SetModel(v string) *CreateScheduledTaskShrinkRequest
	GetModel() *string
	SetName(v string) *CreateScheduledTaskShrinkRequest
	GetName() *string
	SetSegmentsShrink(v string) *CreateScheduledTaskShrinkRequest
	GetSegmentsShrink() *string
	SetTaskDetailShrink(v string) *CreateScheduledTaskShrinkRequest
	GetTaskDetailShrink() *string
	SetTenantId(v string) *CreateScheduledTaskShrinkRequest
	GetTenantId() *string
	SetTriggerConfigShrink(v string) *CreateScheduledTaskShrinkRequest
	GetTriggerConfigShrink() *string
}

type CreateScheduledTaskShrinkRequest struct {
	// 所属协作群组 ID（如 cg_101）；传入时创建群空间任务（调用者需为有效群成员），为空创建个人任务
	//
	// example:
	//
	// exampleCollaborationGroupId
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	DescriptionShrink    *string `json:"description,omitempty" xml:"description,omitempty"`
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
	// 执行模型档位，不传默认 standard
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// 文件名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例名称.pdf
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	SegmentsShrink   *string `json:"segments,omitempty" xml:"segments,omitempty"`
	TaskDetailShrink *string `json:"taskDetail,omitempty" xml:"taskDetail,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId            *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	TriggerConfigShrink *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
}

func (s CreateScheduledTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateScheduledTaskShrinkRequest) GetDescriptionShrink() *string {
	return s.DescriptionShrink
}

func (s *CreateScheduledTaskShrinkRequest) GetDigitalEmployeeNameShrink() *string {
	return s.DigitalEmployeeNameShrink
}

func (s *CreateScheduledTaskShrinkRequest) GetIsOpen() *bool {
	return s.IsOpen
}

func (s *CreateScheduledTaskShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *CreateScheduledTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateScheduledTaskShrinkRequest) GetSegmentsShrink() *string {
	return s.SegmentsShrink
}

func (s *CreateScheduledTaskShrinkRequest) GetTaskDetailShrink() *string {
	return s.TaskDetailShrink
}

func (s *CreateScheduledTaskShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateScheduledTaskShrinkRequest) GetTriggerConfigShrink() *string {
	return s.TriggerConfigShrink
}

func (s *CreateScheduledTaskShrinkRequest) SetCollaborationGroupId(v string) *CreateScheduledTaskShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetDescriptionShrink(v string) *CreateScheduledTaskShrinkRequest {
	s.DescriptionShrink = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetDigitalEmployeeNameShrink(v string) *CreateScheduledTaskShrinkRequest {
	s.DigitalEmployeeNameShrink = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetIsOpen(v bool) *CreateScheduledTaskShrinkRequest {
	s.IsOpen = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetModel(v string) *CreateScheduledTaskShrinkRequest {
	s.Model = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetName(v string) *CreateScheduledTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetSegmentsShrink(v string) *CreateScheduledTaskShrinkRequest {
	s.SegmentsShrink = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetTaskDetailShrink(v string) *CreateScheduledTaskShrinkRequest {
	s.TaskDetailShrink = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetTenantId(v string) *CreateScheduledTaskShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetTriggerConfigShrink(v string) *CreateScheduledTaskShrinkRequest {
	s.TriggerConfigShrink = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
