// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeWorkItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptanceCriteria(v string) *ResumeWorkItemResponseBody
	GetAcceptanceCriteria() *string
	SetAppliedStatusVersion(v int64) *ResumeWorkItemResponseBody
	GetAppliedStatusVersion() *int64
	SetAttachments(v []*ResumeWorkItemResponseBodyAttachments) *ResumeWorkItemResponseBody
	GetAttachments() []*ResumeWorkItemResponseBodyAttachments
	SetBusinessStatus(v string) *ResumeWorkItemResponseBody
	GetBusinessStatus() *string
	SetCanDelete(v bool) *ResumeWorkItemResponseBody
	GetCanDelete() *bool
	SetCanManage(v bool) *ResumeWorkItemResponseBody
	GetCanManage() *bool
	SetCanManageVisibility(v bool) *ResumeWorkItemResponseBody
	GetCanManageVisibility() *bool
	SetCanRename(v bool) *ResumeWorkItemResponseBody
	GetCanRename() *bool
	SetCode(v string) *ResumeWorkItemResponseBody
	GetCode() *string
	SetCollaborationGroupId(v string) *ResumeWorkItemResponseBody
	GetCollaborationGroupId() *string
	SetCollaboratorUserIds(v []*int64) *ResumeWorkItemResponseBody
	GetCollaboratorUserIds() []*int64
	SetConversationId(v string) *ResumeWorkItemResponseBody
	GetConversationId() *string
	SetCreatedAt(v string) *ResumeWorkItemResponseBody
	GetCreatedAt() *string
	SetCreatedByUserId(v int64) *ResumeWorkItemResponseBody
	GetCreatedByUserId() *int64
	SetCreationType(v string) *ResumeWorkItemResponseBody
	GetCreationType() *string
	SetCurrentPlanId(v string) *ResumeWorkItemResponseBody
	GetCurrentPlanId() *string
	SetCurrentRevision(v int64) *ResumeWorkItemResponseBody
	GetCurrentRevision() *int64
	SetCurrentRunId(v string) *ResumeWorkItemResponseBody
	GetCurrentRunId() *string
	SetCurrentRunSeq(v int64) *ResumeWorkItemResponseBody
	GetCurrentRunSeq() *int64
	SetCurrentStepName(v string) *ResumeWorkItemResponseBody
	GetCurrentStepName() *string
	SetDescription(v string) *ResumeWorkItemResponseBody
	GetDescription() *string
	SetInputVersion(v int64) *ResumeWorkItemResponseBody
	GetInputVersion() *int64
	SetItemType(v string) *ResumeWorkItemResponseBody
	GetItemType() *string
	SetLastProcessingAt(v string) *ResumeWorkItemResponseBody
	GetLastProcessingAt() *string
	SetLegacyTopic(v bool) *ResumeWorkItemResponseBody
	GetLegacyTopic() *bool
	SetMemberUserIds(v []*int64) *ResumeWorkItemResponseBody
	GetMemberUserIds() []*int64
	SetMessage(v string) *ResumeWorkItemResponseBody
	GetMessage() *string
	SetModifiedAt(v string) *ResumeWorkItemResponseBody
	GetModifiedAt() *string
	SetOwnerUserId(v int64) *ResumeWorkItemResponseBody
	GetOwnerUserId() *int64
	SetPinned(v bool) *ResumeWorkItemResponseBody
	GetPinned() *bool
	SetPriority(v string) *ResumeWorkItemResponseBody
	GetPriority() *string
	SetProgressCompleted(v int64) *ResumeWorkItemResponseBody
	GetProgressCompleted() *int64
	SetProgressTotal(v int64) *ResumeWorkItemResponseBody
	GetProgressTotal() *int64
	SetRequestId(v string) *ResumeWorkItemResponseBody
	GetRequestId() *string
	SetSourceExecutionId(v string) *ResumeWorkItemResponseBody
	GetSourceExecutionId() *string
	SetSourceTaskId(v string) *ResumeWorkItemResponseBody
	GetSourceTaskId() *string
	SetSourceTaskName(v string) *ResumeWorkItemResponseBody
	GetSourceTaskName() *string
	SetSourceTriggerKey(v string) *ResumeWorkItemResponseBody
	GetSourceTriggerKey() *string
	SetSourceType(v string) *ResumeWorkItemResponseBody
	GetSourceType() *string
	SetTitle(v string) *ResumeWorkItemResponseBody
	GetTitle() *string
	SetTriggerTimezone(v string) *ResumeWorkItemResponseBody
	GetTriggerTimezone() *string
	SetTriggeredAt(v string) *ResumeWorkItemResponseBody
	GetTriggeredAt() *string
	SetVersion(v int64) *ResumeWorkItemResponseBody
	GetVersion() *int64
	SetVisibility(v string) *ResumeWorkItemResponseBody
	GetVisibility() *string
	SetVisibilitySource(v string) *ResumeWorkItemResponseBody
	GetVisibilitySource() *string
	SetWorkItemId(v string) *ResumeWorkItemResponseBody
	GetWorkItemId() *string
	SetWorkflowStatus(v string) *ResumeWorkItemResponseBody
	GetWorkflowStatus() *string
	SetWorkflowTemplateIds(v []*string) *ResumeWorkItemResponseBody
	GetWorkflowTemplateIds() []*string
}

type ResumeWorkItemResponseBody struct {
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
	Attachments []*ResumeWorkItemResponseBodyAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
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

func (s ResumeWorkItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeWorkItemResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeWorkItemResponseBody) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *ResumeWorkItemResponseBody) GetAppliedStatusVersion() *int64 {
	return s.AppliedStatusVersion
}

func (s *ResumeWorkItemResponseBody) GetAttachments() []*ResumeWorkItemResponseBodyAttachments {
	return s.Attachments
}

func (s *ResumeWorkItemResponseBody) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *ResumeWorkItemResponseBody) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *ResumeWorkItemResponseBody) GetCanManage() *bool {
	return s.CanManage
}

func (s *ResumeWorkItemResponseBody) GetCanManageVisibility() *bool {
	return s.CanManageVisibility
}

func (s *ResumeWorkItemResponseBody) GetCanRename() *bool {
	return s.CanRename
}

func (s *ResumeWorkItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResumeWorkItemResponseBody) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ResumeWorkItemResponseBody) GetCollaboratorUserIds() []*int64 {
	return s.CollaboratorUserIds
}

func (s *ResumeWorkItemResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *ResumeWorkItemResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ResumeWorkItemResponseBody) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *ResumeWorkItemResponseBody) GetCreationType() *string {
	return s.CreationType
}

func (s *ResumeWorkItemResponseBody) GetCurrentPlanId() *string {
	return s.CurrentPlanId
}

func (s *ResumeWorkItemResponseBody) GetCurrentRevision() *int64 {
	return s.CurrentRevision
}

func (s *ResumeWorkItemResponseBody) GetCurrentRunId() *string {
	return s.CurrentRunId
}

func (s *ResumeWorkItemResponseBody) GetCurrentRunSeq() *int64 {
	return s.CurrentRunSeq
}

func (s *ResumeWorkItemResponseBody) GetCurrentStepName() *string {
	return s.CurrentStepName
}

func (s *ResumeWorkItemResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ResumeWorkItemResponseBody) GetInputVersion() *int64 {
	return s.InputVersion
}

func (s *ResumeWorkItemResponseBody) GetItemType() *string {
	return s.ItemType
}

func (s *ResumeWorkItemResponseBody) GetLastProcessingAt() *string {
	return s.LastProcessingAt
}

func (s *ResumeWorkItemResponseBody) GetLegacyTopic() *bool {
	return s.LegacyTopic
}

func (s *ResumeWorkItemResponseBody) GetMemberUserIds() []*int64 {
	return s.MemberUserIds
}

func (s *ResumeWorkItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResumeWorkItemResponseBody) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *ResumeWorkItemResponseBody) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *ResumeWorkItemResponseBody) GetPinned() *bool {
	return s.Pinned
}

func (s *ResumeWorkItemResponseBody) GetPriority() *string {
	return s.Priority
}

func (s *ResumeWorkItemResponseBody) GetProgressCompleted() *int64 {
	return s.ProgressCompleted
}

func (s *ResumeWorkItemResponseBody) GetProgressTotal() *int64 {
	return s.ProgressTotal
}

func (s *ResumeWorkItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeWorkItemResponseBody) GetSourceExecutionId() *string {
	return s.SourceExecutionId
}

func (s *ResumeWorkItemResponseBody) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *ResumeWorkItemResponseBody) GetSourceTaskName() *string {
	return s.SourceTaskName
}

func (s *ResumeWorkItemResponseBody) GetSourceTriggerKey() *string {
	return s.SourceTriggerKey
}

func (s *ResumeWorkItemResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *ResumeWorkItemResponseBody) GetTitle() *string {
	return s.Title
}

func (s *ResumeWorkItemResponseBody) GetTriggerTimezone() *string {
	return s.TriggerTimezone
}

func (s *ResumeWorkItemResponseBody) GetTriggeredAt() *string {
	return s.TriggeredAt
}

func (s *ResumeWorkItemResponseBody) GetVersion() *int64 {
	return s.Version
}

func (s *ResumeWorkItemResponseBody) GetVisibility() *string {
	return s.Visibility
}

func (s *ResumeWorkItemResponseBody) GetVisibilitySource() *string {
	return s.VisibilitySource
}

func (s *ResumeWorkItemResponseBody) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ResumeWorkItemResponseBody) GetWorkflowStatus() *string {
	return s.WorkflowStatus
}

func (s *ResumeWorkItemResponseBody) GetWorkflowTemplateIds() []*string {
	return s.WorkflowTemplateIds
}

func (s *ResumeWorkItemResponseBody) SetAcceptanceCriteria(v string) *ResumeWorkItemResponseBody {
	s.AcceptanceCriteria = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetAppliedStatusVersion(v int64) *ResumeWorkItemResponseBody {
	s.AppliedStatusVersion = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetAttachments(v []*ResumeWorkItemResponseBodyAttachments) *ResumeWorkItemResponseBody {
	s.Attachments = v
	return s
}

func (s *ResumeWorkItemResponseBody) SetBusinessStatus(v string) *ResumeWorkItemResponseBody {
	s.BusinessStatus = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCanDelete(v bool) *ResumeWorkItemResponseBody {
	s.CanDelete = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCanManage(v bool) *ResumeWorkItemResponseBody {
	s.CanManage = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCanManageVisibility(v bool) *ResumeWorkItemResponseBody {
	s.CanManageVisibility = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCanRename(v bool) *ResumeWorkItemResponseBody {
	s.CanRename = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCode(v string) *ResumeWorkItemResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCollaborationGroupId(v string) *ResumeWorkItemResponseBody {
	s.CollaborationGroupId = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCollaboratorUserIds(v []*int64) *ResumeWorkItemResponseBody {
	s.CollaboratorUserIds = v
	return s
}

func (s *ResumeWorkItemResponseBody) SetConversationId(v string) *ResumeWorkItemResponseBody {
	s.ConversationId = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCreatedAt(v string) *ResumeWorkItemResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCreatedByUserId(v int64) *ResumeWorkItemResponseBody {
	s.CreatedByUserId = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCreationType(v string) *ResumeWorkItemResponseBody {
	s.CreationType = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCurrentPlanId(v string) *ResumeWorkItemResponseBody {
	s.CurrentPlanId = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCurrentRevision(v int64) *ResumeWorkItemResponseBody {
	s.CurrentRevision = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCurrentRunId(v string) *ResumeWorkItemResponseBody {
	s.CurrentRunId = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCurrentRunSeq(v int64) *ResumeWorkItemResponseBody {
	s.CurrentRunSeq = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetCurrentStepName(v string) *ResumeWorkItemResponseBody {
	s.CurrentStepName = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetDescription(v string) *ResumeWorkItemResponseBody {
	s.Description = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetInputVersion(v int64) *ResumeWorkItemResponseBody {
	s.InputVersion = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetItemType(v string) *ResumeWorkItemResponseBody {
	s.ItemType = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetLastProcessingAt(v string) *ResumeWorkItemResponseBody {
	s.LastProcessingAt = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetLegacyTopic(v bool) *ResumeWorkItemResponseBody {
	s.LegacyTopic = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetMemberUserIds(v []*int64) *ResumeWorkItemResponseBody {
	s.MemberUserIds = v
	return s
}

func (s *ResumeWorkItemResponseBody) SetMessage(v string) *ResumeWorkItemResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetModifiedAt(v string) *ResumeWorkItemResponseBody {
	s.ModifiedAt = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetOwnerUserId(v int64) *ResumeWorkItemResponseBody {
	s.OwnerUserId = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetPinned(v bool) *ResumeWorkItemResponseBody {
	s.Pinned = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetPriority(v string) *ResumeWorkItemResponseBody {
	s.Priority = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetProgressCompleted(v int64) *ResumeWorkItemResponseBody {
	s.ProgressCompleted = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetProgressTotal(v int64) *ResumeWorkItemResponseBody {
	s.ProgressTotal = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetRequestId(v string) *ResumeWorkItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetSourceExecutionId(v string) *ResumeWorkItemResponseBody {
	s.SourceExecutionId = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetSourceTaskId(v string) *ResumeWorkItemResponseBody {
	s.SourceTaskId = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetSourceTaskName(v string) *ResumeWorkItemResponseBody {
	s.SourceTaskName = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetSourceTriggerKey(v string) *ResumeWorkItemResponseBody {
	s.SourceTriggerKey = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetSourceType(v string) *ResumeWorkItemResponseBody {
	s.SourceType = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetTitle(v string) *ResumeWorkItemResponseBody {
	s.Title = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetTriggerTimezone(v string) *ResumeWorkItemResponseBody {
	s.TriggerTimezone = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetTriggeredAt(v string) *ResumeWorkItemResponseBody {
	s.TriggeredAt = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetVersion(v int64) *ResumeWorkItemResponseBody {
	s.Version = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetVisibility(v string) *ResumeWorkItemResponseBody {
	s.Visibility = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetVisibilitySource(v string) *ResumeWorkItemResponseBody {
	s.VisibilitySource = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetWorkItemId(v string) *ResumeWorkItemResponseBody {
	s.WorkItemId = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetWorkflowStatus(v string) *ResumeWorkItemResponseBody {
	s.WorkflowStatus = &v
	return s
}

func (s *ResumeWorkItemResponseBody) SetWorkflowTemplateIds(v []*string) *ResumeWorkItemResponseBody {
	s.WorkflowTemplateIds = v
	return s
}

func (s *ResumeWorkItemResponseBody) Validate() error {
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

type ResumeWorkItemResponseBodyAttachments struct {
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

func (s ResumeWorkItemResponseBodyAttachments) String() string {
	return dara.Prettify(s)
}

func (s ResumeWorkItemResponseBodyAttachments) GoString() string {
	return s.String()
}

func (s *ResumeWorkItemResponseBodyAttachments) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *ResumeWorkItemResponseBodyAttachments) GetContentType() *string {
	return s.ContentType
}

func (s *ResumeWorkItemResponseBodyAttachments) GetRefUrl() *string {
	return s.RefUrl
}

func (s *ResumeWorkItemResponseBodyAttachments) GetResourceId() *string {
	return s.ResourceId
}

func (s *ResumeWorkItemResponseBodyAttachments) GetResourceKind() *string {
	return s.ResourceKind
}

func (s *ResumeWorkItemResponseBodyAttachments) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *ResumeWorkItemResponseBodyAttachments) GetTitle() *string {
	return s.Title
}

func (s *ResumeWorkItemResponseBodyAttachments) SetAttachmentId(v string) *ResumeWorkItemResponseBodyAttachments {
	s.AttachmentId = &v
	return s
}

func (s *ResumeWorkItemResponseBodyAttachments) SetContentType(v string) *ResumeWorkItemResponseBodyAttachments {
	s.ContentType = &v
	return s
}

func (s *ResumeWorkItemResponseBodyAttachments) SetRefUrl(v string) *ResumeWorkItemResponseBodyAttachments {
	s.RefUrl = &v
	return s
}

func (s *ResumeWorkItemResponseBodyAttachments) SetResourceId(v string) *ResumeWorkItemResponseBodyAttachments {
	s.ResourceId = &v
	return s
}

func (s *ResumeWorkItemResponseBodyAttachments) SetResourceKind(v string) *ResumeWorkItemResponseBodyAttachments {
	s.ResourceKind = &v
	return s
}

func (s *ResumeWorkItemResponseBodyAttachments) SetSizeBytes(v int64) *ResumeWorkItemResponseBodyAttachments {
	s.SizeBytes = &v
	return s
}

func (s *ResumeWorkItemResponseBodyAttachments) SetTitle(v string) *ResumeWorkItemResponseBodyAttachments {
	s.Title = &v
	return s
}

func (s *ResumeWorkItemResponseBodyAttachments) Validate() error {
	return dara.Validate(s)
}
