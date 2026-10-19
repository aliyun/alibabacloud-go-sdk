// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseWorkItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptanceCriteria(v string) *PauseWorkItemResponseBody
	GetAcceptanceCriteria() *string
	SetAppliedStatusVersion(v int64) *PauseWorkItemResponseBody
	GetAppliedStatusVersion() *int64
	SetAttachments(v []*PauseWorkItemResponseBodyAttachments) *PauseWorkItemResponseBody
	GetAttachments() []*PauseWorkItemResponseBodyAttachments
	SetBusinessStatus(v string) *PauseWorkItemResponseBody
	GetBusinessStatus() *string
	SetCanDelete(v bool) *PauseWorkItemResponseBody
	GetCanDelete() *bool
	SetCanManage(v bool) *PauseWorkItemResponseBody
	GetCanManage() *bool
	SetCanManageVisibility(v bool) *PauseWorkItemResponseBody
	GetCanManageVisibility() *bool
	SetCanRename(v bool) *PauseWorkItemResponseBody
	GetCanRename() *bool
	SetCode(v string) *PauseWorkItemResponseBody
	GetCode() *string
	SetCollaborationGroupId(v string) *PauseWorkItemResponseBody
	GetCollaborationGroupId() *string
	SetCollaboratorUserIds(v []*int64) *PauseWorkItemResponseBody
	GetCollaboratorUserIds() []*int64
	SetConversationId(v string) *PauseWorkItemResponseBody
	GetConversationId() *string
	SetCreatedAt(v string) *PauseWorkItemResponseBody
	GetCreatedAt() *string
	SetCreatedByUserId(v int64) *PauseWorkItemResponseBody
	GetCreatedByUserId() *int64
	SetCreationType(v string) *PauseWorkItemResponseBody
	GetCreationType() *string
	SetCurrentPlanId(v string) *PauseWorkItemResponseBody
	GetCurrentPlanId() *string
	SetCurrentRevision(v int64) *PauseWorkItemResponseBody
	GetCurrentRevision() *int64
	SetCurrentRunId(v string) *PauseWorkItemResponseBody
	GetCurrentRunId() *string
	SetCurrentRunSeq(v int64) *PauseWorkItemResponseBody
	GetCurrentRunSeq() *int64
	SetCurrentStepName(v string) *PauseWorkItemResponseBody
	GetCurrentStepName() *string
	SetDescription(v string) *PauseWorkItemResponseBody
	GetDescription() *string
	SetInputVersion(v int64) *PauseWorkItemResponseBody
	GetInputVersion() *int64
	SetItemType(v string) *PauseWorkItemResponseBody
	GetItemType() *string
	SetLastProcessingAt(v string) *PauseWorkItemResponseBody
	GetLastProcessingAt() *string
	SetLegacyTopic(v bool) *PauseWorkItemResponseBody
	GetLegacyTopic() *bool
	SetMemberUserIds(v []*int64) *PauseWorkItemResponseBody
	GetMemberUserIds() []*int64
	SetMessage(v string) *PauseWorkItemResponseBody
	GetMessage() *string
	SetModifiedAt(v string) *PauseWorkItemResponseBody
	GetModifiedAt() *string
	SetOwnerUserId(v int64) *PauseWorkItemResponseBody
	GetOwnerUserId() *int64
	SetPinned(v bool) *PauseWorkItemResponseBody
	GetPinned() *bool
	SetPriority(v string) *PauseWorkItemResponseBody
	GetPriority() *string
	SetProgressCompleted(v int64) *PauseWorkItemResponseBody
	GetProgressCompleted() *int64
	SetProgressTotal(v int64) *PauseWorkItemResponseBody
	GetProgressTotal() *int64
	SetRequestId(v string) *PauseWorkItemResponseBody
	GetRequestId() *string
	SetSourceExecutionId(v string) *PauseWorkItemResponseBody
	GetSourceExecutionId() *string
	SetSourceTaskId(v string) *PauseWorkItemResponseBody
	GetSourceTaskId() *string
	SetSourceTaskName(v string) *PauseWorkItemResponseBody
	GetSourceTaskName() *string
	SetSourceTriggerKey(v string) *PauseWorkItemResponseBody
	GetSourceTriggerKey() *string
	SetSourceType(v string) *PauseWorkItemResponseBody
	GetSourceType() *string
	SetTitle(v string) *PauseWorkItemResponseBody
	GetTitle() *string
	SetTriggerTimezone(v string) *PauseWorkItemResponseBody
	GetTriggerTimezone() *string
	SetTriggeredAt(v string) *PauseWorkItemResponseBody
	GetTriggeredAt() *string
	SetVersion(v int64) *PauseWorkItemResponseBody
	GetVersion() *int64
	SetVisibility(v string) *PauseWorkItemResponseBody
	GetVisibility() *string
	SetVisibilitySource(v string) *PauseWorkItemResponseBody
	GetVisibilitySource() *string
	SetWorkItemId(v string) *PauseWorkItemResponseBody
	GetWorkItemId() *string
	SetWorkflowStatus(v string) *PauseWorkItemResponseBody
	GetWorkflowStatus() *string
	SetWorkflowTemplateIds(v []*string) *PauseWorkItemResponseBody
	GetWorkflowTemplateIds() []*string
}

type PauseWorkItemResponseBody struct {
	// 验收标准
	//
	// example:
	//
	// example
	AcceptanceCriteria *string `json:"acceptanceCriteria,omitempty" xml:"acceptanceCriteria,omitempty"`
	// 已应用的状态版本
	//
	// example:
	//
	// 1
	AppliedStatusVersion *int64 `json:"appliedStatusVersion,omitempty" xml:"appliedStatusVersion,omitempty"`
	// 输入附件列表
	//
	// example:
	//
	// example
	Attachments []*PauseWorkItemResponseBodyAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// 业务状态（七态）
	//
	// example:
	//
	// PROCESSING
	BusinessStatus *string `json:"businessStatus,omitempty" xml:"businessStatus,omitempty"`
	// 调用者是否可删除
	//
	// example:
	//
	// false
	CanDelete *bool `json:"canDelete,omitempty" xml:"canDelete,omitempty"`
	// 调用者是否可管理
	//
	// example:
	//
	// false
	CanManage *bool `json:"canManage,omitempty" xml:"canManage,omitempty"`
	// 是否可独立调整可见范围
	//
	// example:
	//
	// false
	CanManageVisibility *bool `json:"canManageVisibility,omitempty" xml:"canManageVisibility,omitempty"`
	// 调用者是否可重命名
	//
	// example:
	//
	// false
	CanRename *bool `json:"canRename,omitempty" xml:"canRename,omitempty"`
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 协作空间ID
	//
	// example:
	//
	// example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 协作人用户ID列表
	//
	// example:
	//
	// example
	CollaboratorUserIds []*int64 `json:"collaboratorUserIds,omitempty" xml:"collaboratorUserIds,omitempty" type:"Repeated"`
	// 关联会话ID
	//
	// example:
	//
	// example
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// 创建时间(ISO8601)
	//
	// example:
	//
	// example
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 创建人用户ID
	//
	// example:
	//
	// 1
	CreatedByUserId *int64 `json:"createdByUserId,omitempty" xml:"createdByUserId,omitempty"`
	// 创建来源：DEFAULT/TASK
	//
	// example:
	//
	// DEFAULT
	CreationType *string `json:"creationType,omitempty" xml:"creationType,omitempty"`
	// 当前物理计划ID
	//
	// example:
	//
	// example
	CurrentPlanId *string `json:"currentPlanId,omitempty" xml:"currentPlanId,omitempty"`
	// 当前修订号
	//
	// example:
	//
	// 1
	CurrentRevision *int64 `json:"currentRevision,omitempty" xml:"currentRevision,omitempty"`
	// 当前运行ID
	//
	// example:
	//
	// example
	CurrentRunId *string `json:"currentRunId,omitempty" xml:"currentRunId,omitempty"`
	// 当前运行序号
	//
	// example:
	//
	// 1
	CurrentRunSeq *int64 `json:"currentRunSeq,omitempty" xml:"currentRunSeq,omitempty"`
	// 当前步骤名称
	//
	// example:
	//
	// example
	CurrentStepName *string `json:"currentStepName,omitempty" xml:"currentStepName,omitempty"`
	// 描述
	//
	// example:
	//
	// example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 输入快照版本
	//
	// example:
	//
	// 1
	InputVersion *int64 `json:"inputVersion,omitempty" xml:"inputVersion,omitempty"`
	// 工作项类型
	//
	// example:
	//
	// WORKFLOW
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// 最近进入处理中的时间(ISO8601)
	//
	// example:
	//
	// example
	LastProcessingAt *string `json:"lastProcessingAt,omitempty" xml:"lastProcessingAt,omitempty"`
	// 是否为兼容期自由探索话题
	//
	// example:
	//
	// false
	LegacyTopic *bool `json:"legacyTopic,omitempty" xml:"legacyTopic,omitempty"`
	// 可见成员用户ID列表
	//
	// example:
	//
	// example
	MemberUserIds []*int64 `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 修改时间(ISO8601)
	//
	// example:
	//
	// example
	ModifiedAt *string `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// 负责人用户ID
	//
	// example:
	//
	// 1
	OwnerUserId *int64 `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// 调用者个人置顶状态
	//
	// example:
	//
	// false
	Pinned *bool `json:"pinned,omitempty" xml:"pinned,omitempty"`
	// 优先级
	//
	// example:
	//
	// HIGH
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// 已完成进度
	//
	// example:
	//
	// 1
	ProgressCompleted *int64 `json:"progressCompleted,omitempty" xml:"progressCompleted,omitempty"`
	// 总进度
	//
	// example:
	//
	// 1
	ProgressTotal *int64 `json:"progressTotal,omitempty" xml:"progressTotal,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 任务执行ID
	//
	// example:
	//
	// example
	SourceExecutionId *string `json:"sourceExecutionId,omitempty" xml:"sourceExecutionId,omitempty"`
	// 任务来源ID
	//
	// example:
	//
	// example
	SourceTaskId *string `json:"sourceTaskId,omitempty" xml:"sourceTaskId,omitempty"`
	// 任务来源名称
	//
	// example:
	//
	// example
	SourceTaskName *string `json:"sourceTaskName,omitempty" xml:"sourceTaskName,omitempty"`
	// 任务触发键
	//
	// example:
	//
	// example
	SourceTriggerKey *string `json:"sourceTriggerKey,omitempty" xml:"sourceTriggerKey,omitempty"`
	// 任务来源类型
	//
	// example:
	//
	// example
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// 标题
	//
	// example:
	//
	// example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 任务触发时区
	//
	// example:
	//
	// example
	TriggerTimezone *string `json:"triggerTimezone,omitempty" xml:"triggerTimezone,omitempty"`
	// 任务触发时间(ISO8601)
	//
	// example:
	//
	// example
	TriggeredAt *string `json:"triggeredAt,omitempty" xml:"triggeredAt,omitempty"`
	// 乐观版本号
	//
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 可见范围
	//
	// example:
	//
	// example
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
	// 可见范围来源：SELF/TASK/TEMPLATE
	//
	// example:
	//
	// SELF
	VisibilitySource *string `json:"visibilitySource,omitempty" xml:"visibilitySource,omitempty"`
	// 工作项ID
	//
	// example:
	//
	// example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
	// Workflow四态投影
	//
	// example:
	//
	// example
	WorkflowStatus *string `json:"workflowStatus,omitempty" xml:"workflowStatus,omitempty"`
	// 绑定的工作流模板ID列表
	//
	// example:
	//
	// example
	WorkflowTemplateIds []*string `json:"workflowTemplateIds,omitempty" xml:"workflowTemplateIds,omitempty" type:"Repeated"`
}

func (s PauseWorkItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseWorkItemResponseBody) GoString() string {
	return s.String()
}

func (s *PauseWorkItemResponseBody) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *PauseWorkItemResponseBody) GetAppliedStatusVersion() *int64 {
	return s.AppliedStatusVersion
}

func (s *PauseWorkItemResponseBody) GetAttachments() []*PauseWorkItemResponseBodyAttachments {
	return s.Attachments
}

func (s *PauseWorkItemResponseBody) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *PauseWorkItemResponseBody) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *PauseWorkItemResponseBody) GetCanManage() *bool {
	return s.CanManage
}

func (s *PauseWorkItemResponseBody) GetCanManageVisibility() *bool {
	return s.CanManageVisibility
}

func (s *PauseWorkItemResponseBody) GetCanRename() *bool {
	return s.CanRename
}

func (s *PauseWorkItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *PauseWorkItemResponseBody) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *PauseWorkItemResponseBody) GetCollaboratorUserIds() []*int64 {
	return s.CollaboratorUserIds
}

func (s *PauseWorkItemResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *PauseWorkItemResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *PauseWorkItemResponseBody) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *PauseWorkItemResponseBody) GetCreationType() *string {
	return s.CreationType
}

func (s *PauseWorkItemResponseBody) GetCurrentPlanId() *string {
	return s.CurrentPlanId
}

func (s *PauseWorkItemResponseBody) GetCurrentRevision() *int64 {
	return s.CurrentRevision
}

func (s *PauseWorkItemResponseBody) GetCurrentRunId() *string {
	return s.CurrentRunId
}

func (s *PauseWorkItemResponseBody) GetCurrentRunSeq() *int64 {
	return s.CurrentRunSeq
}

func (s *PauseWorkItemResponseBody) GetCurrentStepName() *string {
	return s.CurrentStepName
}

func (s *PauseWorkItemResponseBody) GetDescription() *string {
	return s.Description
}

func (s *PauseWorkItemResponseBody) GetInputVersion() *int64 {
	return s.InputVersion
}

func (s *PauseWorkItemResponseBody) GetItemType() *string {
	return s.ItemType
}

func (s *PauseWorkItemResponseBody) GetLastProcessingAt() *string {
	return s.LastProcessingAt
}

func (s *PauseWorkItemResponseBody) GetLegacyTopic() *bool {
	return s.LegacyTopic
}

func (s *PauseWorkItemResponseBody) GetMemberUserIds() []*int64 {
	return s.MemberUserIds
}

func (s *PauseWorkItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PauseWorkItemResponseBody) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *PauseWorkItemResponseBody) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *PauseWorkItemResponseBody) GetPinned() *bool {
	return s.Pinned
}

func (s *PauseWorkItemResponseBody) GetPriority() *string {
	return s.Priority
}

func (s *PauseWorkItemResponseBody) GetProgressCompleted() *int64 {
	return s.ProgressCompleted
}

func (s *PauseWorkItemResponseBody) GetProgressTotal() *int64 {
	return s.ProgressTotal
}

func (s *PauseWorkItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseWorkItemResponseBody) GetSourceExecutionId() *string {
	return s.SourceExecutionId
}

func (s *PauseWorkItemResponseBody) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *PauseWorkItemResponseBody) GetSourceTaskName() *string {
	return s.SourceTaskName
}

func (s *PauseWorkItemResponseBody) GetSourceTriggerKey() *string {
	return s.SourceTriggerKey
}

func (s *PauseWorkItemResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *PauseWorkItemResponseBody) GetTitle() *string {
	return s.Title
}

func (s *PauseWorkItemResponseBody) GetTriggerTimezone() *string {
	return s.TriggerTimezone
}

func (s *PauseWorkItemResponseBody) GetTriggeredAt() *string {
	return s.TriggeredAt
}

func (s *PauseWorkItemResponseBody) GetVersion() *int64 {
	return s.Version
}

func (s *PauseWorkItemResponseBody) GetVisibility() *string {
	return s.Visibility
}

func (s *PauseWorkItemResponseBody) GetVisibilitySource() *string {
	return s.VisibilitySource
}

func (s *PauseWorkItemResponseBody) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *PauseWorkItemResponseBody) GetWorkflowStatus() *string {
	return s.WorkflowStatus
}

func (s *PauseWorkItemResponseBody) GetWorkflowTemplateIds() []*string {
	return s.WorkflowTemplateIds
}

func (s *PauseWorkItemResponseBody) SetAcceptanceCriteria(v string) *PauseWorkItemResponseBody {
	s.AcceptanceCriteria = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetAppliedStatusVersion(v int64) *PauseWorkItemResponseBody {
	s.AppliedStatusVersion = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetAttachments(v []*PauseWorkItemResponseBodyAttachments) *PauseWorkItemResponseBody {
	s.Attachments = v
	return s
}

func (s *PauseWorkItemResponseBody) SetBusinessStatus(v string) *PauseWorkItemResponseBody {
	s.BusinessStatus = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCanDelete(v bool) *PauseWorkItemResponseBody {
	s.CanDelete = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCanManage(v bool) *PauseWorkItemResponseBody {
	s.CanManage = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCanManageVisibility(v bool) *PauseWorkItemResponseBody {
	s.CanManageVisibility = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCanRename(v bool) *PauseWorkItemResponseBody {
	s.CanRename = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCode(v string) *PauseWorkItemResponseBody {
	s.Code = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCollaborationGroupId(v string) *PauseWorkItemResponseBody {
	s.CollaborationGroupId = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCollaboratorUserIds(v []*int64) *PauseWorkItemResponseBody {
	s.CollaboratorUserIds = v
	return s
}

func (s *PauseWorkItemResponseBody) SetConversationId(v string) *PauseWorkItemResponseBody {
	s.ConversationId = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCreatedAt(v string) *PauseWorkItemResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCreatedByUserId(v int64) *PauseWorkItemResponseBody {
	s.CreatedByUserId = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCreationType(v string) *PauseWorkItemResponseBody {
	s.CreationType = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCurrentPlanId(v string) *PauseWorkItemResponseBody {
	s.CurrentPlanId = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCurrentRevision(v int64) *PauseWorkItemResponseBody {
	s.CurrentRevision = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCurrentRunId(v string) *PauseWorkItemResponseBody {
	s.CurrentRunId = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCurrentRunSeq(v int64) *PauseWorkItemResponseBody {
	s.CurrentRunSeq = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetCurrentStepName(v string) *PauseWorkItemResponseBody {
	s.CurrentStepName = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetDescription(v string) *PauseWorkItemResponseBody {
	s.Description = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetInputVersion(v int64) *PauseWorkItemResponseBody {
	s.InputVersion = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetItemType(v string) *PauseWorkItemResponseBody {
	s.ItemType = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetLastProcessingAt(v string) *PauseWorkItemResponseBody {
	s.LastProcessingAt = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetLegacyTopic(v bool) *PauseWorkItemResponseBody {
	s.LegacyTopic = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetMemberUserIds(v []*int64) *PauseWorkItemResponseBody {
	s.MemberUserIds = v
	return s
}

func (s *PauseWorkItemResponseBody) SetMessage(v string) *PauseWorkItemResponseBody {
	s.Message = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetModifiedAt(v string) *PauseWorkItemResponseBody {
	s.ModifiedAt = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetOwnerUserId(v int64) *PauseWorkItemResponseBody {
	s.OwnerUserId = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetPinned(v bool) *PauseWorkItemResponseBody {
	s.Pinned = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetPriority(v string) *PauseWorkItemResponseBody {
	s.Priority = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetProgressCompleted(v int64) *PauseWorkItemResponseBody {
	s.ProgressCompleted = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetProgressTotal(v int64) *PauseWorkItemResponseBody {
	s.ProgressTotal = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetRequestId(v string) *PauseWorkItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetSourceExecutionId(v string) *PauseWorkItemResponseBody {
	s.SourceExecutionId = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetSourceTaskId(v string) *PauseWorkItemResponseBody {
	s.SourceTaskId = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetSourceTaskName(v string) *PauseWorkItemResponseBody {
	s.SourceTaskName = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetSourceTriggerKey(v string) *PauseWorkItemResponseBody {
	s.SourceTriggerKey = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetSourceType(v string) *PauseWorkItemResponseBody {
	s.SourceType = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetTitle(v string) *PauseWorkItemResponseBody {
	s.Title = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetTriggerTimezone(v string) *PauseWorkItemResponseBody {
	s.TriggerTimezone = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetTriggeredAt(v string) *PauseWorkItemResponseBody {
	s.TriggeredAt = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetVersion(v int64) *PauseWorkItemResponseBody {
	s.Version = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetVisibility(v string) *PauseWorkItemResponseBody {
	s.Visibility = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetVisibilitySource(v string) *PauseWorkItemResponseBody {
	s.VisibilitySource = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetWorkItemId(v string) *PauseWorkItemResponseBody {
	s.WorkItemId = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetWorkflowStatus(v string) *PauseWorkItemResponseBody {
	s.WorkflowStatus = &v
	return s
}

func (s *PauseWorkItemResponseBody) SetWorkflowTemplateIds(v []*string) *PauseWorkItemResponseBody {
	s.WorkflowTemplateIds = v
	return s
}

func (s *PauseWorkItemResponseBody) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PauseWorkItemResponseBodyAttachments struct {
	// 附件ID
	//
	// example:
	//
	// example
	AttachmentId *string `json:"attachmentId,omitempty" xml:"attachmentId,omitempty"`
	// 内容类型
	//
	// example:
	//
	// example
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 附件访问入口（OpenAPI 302 前置）
	//
	// example:
	//
	// example
	RefUrl *string `json:"refUrl,omitempty" xml:"refUrl,omitempty"`
	// 附件资源ID
	//
	// example:
	//
	// example
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// 附件资源类型
	//
	// example:
	//
	// example
	ResourceKind *string `json:"resourceKind,omitempty" xml:"resourceKind,omitempty"`
	// 附件大小(字节)
	//
	// example:
	//
	// 1
	SizeBytes *int64 `json:"sizeBytes,omitempty" xml:"sizeBytes,omitempty"`
	// 附件标题
	//
	// example:
	//
	// example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s PauseWorkItemResponseBodyAttachments) String() string {
	return dara.Prettify(s)
}

func (s PauseWorkItemResponseBodyAttachments) GoString() string {
	return s.String()
}

func (s *PauseWorkItemResponseBodyAttachments) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *PauseWorkItemResponseBodyAttachments) GetContentType() *string {
	return s.ContentType
}

func (s *PauseWorkItemResponseBodyAttachments) GetRefUrl() *string {
	return s.RefUrl
}

func (s *PauseWorkItemResponseBodyAttachments) GetResourceId() *string {
	return s.ResourceId
}

func (s *PauseWorkItemResponseBodyAttachments) GetResourceKind() *string {
	return s.ResourceKind
}

func (s *PauseWorkItemResponseBodyAttachments) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *PauseWorkItemResponseBodyAttachments) GetTitle() *string {
	return s.Title
}

func (s *PauseWorkItemResponseBodyAttachments) SetAttachmentId(v string) *PauseWorkItemResponseBodyAttachments {
	s.AttachmentId = &v
	return s
}

func (s *PauseWorkItemResponseBodyAttachments) SetContentType(v string) *PauseWorkItemResponseBodyAttachments {
	s.ContentType = &v
	return s
}

func (s *PauseWorkItemResponseBodyAttachments) SetRefUrl(v string) *PauseWorkItemResponseBodyAttachments {
	s.RefUrl = &v
	return s
}

func (s *PauseWorkItemResponseBodyAttachments) SetResourceId(v string) *PauseWorkItemResponseBodyAttachments {
	s.ResourceId = &v
	return s
}

func (s *PauseWorkItemResponseBodyAttachments) SetResourceKind(v string) *PauseWorkItemResponseBodyAttachments {
	s.ResourceKind = &v
	return s
}

func (s *PauseWorkItemResponseBodyAttachments) SetSizeBytes(v int64) *PauseWorkItemResponseBodyAttachments {
	s.SizeBytes = &v
	return s
}

func (s *PauseWorkItemResponseBodyAttachments) SetTitle(v string) *PauseWorkItemResponseBodyAttachments {
	s.Title = &v
	return s
}

func (s *PauseWorkItemResponseBodyAttachments) Validate() error {
	return dara.Validate(s)
}
