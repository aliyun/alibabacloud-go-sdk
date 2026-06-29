// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleTask interface {
	dara.Model
	String() string
	GoString() string
	SetArchived(v bool) *SimpleTask
	GetArchived() *bool
	SetArchivedInfos(v string) *SimpleTask
	GetArchivedInfos() *string
	SetCreator(v *SimpleUser) *SimpleTask
	GetCreator() *SimpleUser
	SetGmtCreateTime(v string) *SimpleTask
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *SimpleTask
	GetGmtModifiedTime() *string
	SetLabelStyle(v string) *SimpleTask
	GetLabelStyle() *string
	SetModifier(v *SimpleUser) *SimpleTask
	GetModifier() *SimpleUser
	SetRefTaskId(v string) *SimpleTask
	GetRefTaskId() *string
	SetRemark(v string) *SimpleTask
	GetRemark() *string
	SetStage(v string) *SimpleTask
	GetStage() *string
	SetStatus(v string) *SimpleTask
	GetStatus() *string
	SetTags(v []*string) *SimpleTask
	GetTags() []*string
	SetTaskId(v string) *SimpleTask
	GetTaskId() *string
	SetTaskName(v string) *SimpleTask
	GetTaskName() *string
	SetTaskType(v string) *SimpleTask
	GetTaskType() *string
	SetTemplateId(v string) *SimpleTask
	GetTemplateId() *string
	SetTenantId(v string) *SimpleTask
	GetTenantId() *string
	SetUUID(v string) *SimpleTask
	GetUUID() *string
	SetWorkflowNodes(v []*string) *SimpleTask
	GetWorkflowNodes() []*string
}

type SimpleTask struct {
	// Indicates whether the job is archived. Possible values:
	//
	// - false: Not archived.
	//
	// - true: Archived.
	//
	// example:
	//
	// true
	Archived *bool `json:"Archived,omitempty" xml:"Archived,omitempty"`
	// Data archiving information.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// null
	ArchivedInfos *string `json:"ArchivedInfos,omitempty" xml:"ArchivedInfos,omitempty"`
	// Creator information.
	Creator *SimpleUser `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// UTC creation time.
	//
	// example:
	//
	// 2021-07-07 16:09:20
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// UTC time of the last update.
	//
	// example:
	//
	// 2021-07-07 16:09:20
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Label style.
	//
	// example:
	//
	// IMG
	LabelStyle *string `json:"LabelStyle,omitempty" xml:"LabelStyle,omitempty"`
	// Updated By information.
	Modifier *SimpleUser `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// Auto triggered task ID.
	//
	// example:
	//
	// 123***5124
	RefTaskId *string `json:"RefTaskId,omitempty" xml:"RefTaskId,omitempty"`
	// Remarks.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 备注
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Current streaming node. Possible values:
	//
	// - MARK: Annotating.
	//
	// - CHECK: Checking.
	//
	// - FINISHED: Completed.
	//
	// example:
	//
	// MARK
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// Task Status. Possible values:
	//
	// - CREATED
	//
	// - SUCCESS
	//
	// example:
	//
	// CREATED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// List of job labels.
	//
	// if can be null:
	// true
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Job ID.
	//
	// example:
	//
	// 15438***8306500608
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Job name.
	//
	// example:
	//
	// demo
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// Task Type.
	//
	// example:
	//
	// NORMAL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// Template ID.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1529***348342353920
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Tenant ID.
	//
	// example:
	//
	// GA***W134
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// UUID.
	//
	// example:
	//
	// paiworkspace-****
	UUID *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
	// Pipeline nodes.
	WorkflowNodes []*string `json:"WorkflowNodes,omitempty" xml:"WorkflowNodes,omitempty" type:"Repeated"`
}

func (s SimpleTask) String() string {
	return dara.Prettify(s)
}

func (s SimpleTask) GoString() string {
	return s.String()
}

func (s *SimpleTask) GetArchived() *bool {
	return s.Archived
}

func (s *SimpleTask) GetArchivedInfos() *string {
	return s.ArchivedInfos
}

func (s *SimpleTask) GetCreator() *SimpleUser {
	return s.Creator
}

func (s *SimpleTask) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *SimpleTask) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *SimpleTask) GetLabelStyle() *string {
	return s.LabelStyle
}

func (s *SimpleTask) GetModifier() *SimpleUser {
	return s.Modifier
}

func (s *SimpleTask) GetRefTaskId() *string {
	return s.RefTaskId
}

func (s *SimpleTask) GetRemark() *string {
	return s.Remark
}

func (s *SimpleTask) GetStage() *string {
	return s.Stage
}

func (s *SimpleTask) GetStatus() *string {
	return s.Status
}

func (s *SimpleTask) GetTags() []*string {
	return s.Tags
}

func (s *SimpleTask) GetTaskId() *string {
	return s.TaskId
}

func (s *SimpleTask) GetTaskName() *string {
	return s.TaskName
}

func (s *SimpleTask) GetTaskType() *string {
	return s.TaskType
}

func (s *SimpleTask) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SimpleTask) GetTenantId() *string {
	return s.TenantId
}

func (s *SimpleTask) GetUUID() *string {
	return s.UUID
}

func (s *SimpleTask) GetWorkflowNodes() []*string {
	return s.WorkflowNodes
}

func (s *SimpleTask) SetArchived(v bool) *SimpleTask {
	s.Archived = &v
	return s
}

func (s *SimpleTask) SetArchivedInfos(v string) *SimpleTask {
	s.ArchivedInfos = &v
	return s
}

func (s *SimpleTask) SetCreator(v *SimpleUser) *SimpleTask {
	s.Creator = v
	return s
}

func (s *SimpleTask) SetGmtCreateTime(v string) *SimpleTask {
	s.GmtCreateTime = &v
	return s
}

func (s *SimpleTask) SetGmtModifiedTime(v string) *SimpleTask {
	s.GmtModifiedTime = &v
	return s
}

func (s *SimpleTask) SetLabelStyle(v string) *SimpleTask {
	s.LabelStyle = &v
	return s
}

func (s *SimpleTask) SetModifier(v *SimpleUser) *SimpleTask {
	s.Modifier = v
	return s
}

func (s *SimpleTask) SetRefTaskId(v string) *SimpleTask {
	s.RefTaskId = &v
	return s
}

func (s *SimpleTask) SetRemark(v string) *SimpleTask {
	s.Remark = &v
	return s
}

func (s *SimpleTask) SetStage(v string) *SimpleTask {
	s.Stage = &v
	return s
}

func (s *SimpleTask) SetStatus(v string) *SimpleTask {
	s.Status = &v
	return s
}

func (s *SimpleTask) SetTags(v []*string) *SimpleTask {
	s.Tags = v
	return s
}

func (s *SimpleTask) SetTaskId(v string) *SimpleTask {
	s.TaskId = &v
	return s
}

func (s *SimpleTask) SetTaskName(v string) *SimpleTask {
	s.TaskName = &v
	return s
}

func (s *SimpleTask) SetTaskType(v string) *SimpleTask {
	s.TaskType = &v
	return s
}

func (s *SimpleTask) SetTemplateId(v string) *SimpleTask {
	s.TemplateId = &v
	return s
}

func (s *SimpleTask) SetTenantId(v string) *SimpleTask {
	s.TenantId = &v
	return s
}

func (s *SimpleTask) SetUUID(v string) *SimpleTask {
	s.UUID = &v
	return s
}

func (s *SimpleTask) SetWorkflowNodes(v []*string) *SimpleTask {
	s.WorkflowNodes = v
	return s
}

func (s *SimpleTask) Validate() error {
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.Modifier != nil {
		if err := s.Modifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}
