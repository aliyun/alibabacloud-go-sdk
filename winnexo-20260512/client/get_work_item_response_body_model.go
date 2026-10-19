// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptanceCriteria(v string) *GetWorkItemResponseBody
	GetAcceptanceCriteria() *string
	SetAppliedStatusVersion(v int64) *GetWorkItemResponseBody
	GetAppliedStatusVersion() *int64
	SetAttachments(v []*GetWorkItemResponseBodyAttachments) *GetWorkItemResponseBody
	GetAttachments() []*GetWorkItemResponseBodyAttachments
	SetBusinessStatus(v string) *GetWorkItemResponseBody
	GetBusinessStatus() *string
	SetCanDelete(v bool) *GetWorkItemResponseBody
	GetCanDelete() *bool
	SetCanManage(v bool) *GetWorkItemResponseBody
	GetCanManage() *bool
	SetCanManageVisibility(v bool) *GetWorkItemResponseBody
	GetCanManageVisibility() *bool
	SetCanRename(v bool) *GetWorkItemResponseBody
	GetCanRename() *bool
	SetCode(v string) *GetWorkItemResponseBody
	GetCode() *string
	SetCollaborationGroupId(v string) *GetWorkItemResponseBody
	GetCollaborationGroupId() *string
	SetCollaboratorUserIds(v []*int64) *GetWorkItemResponseBody
	GetCollaboratorUserIds() []*int64
	SetConversationId(v string) *GetWorkItemResponseBody
	GetConversationId() *string
	SetCreatedAt(v string) *GetWorkItemResponseBody
	GetCreatedAt() *string
	SetCreatedByUserId(v int64) *GetWorkItemResponseBody
	GetCreatedByUserId() *int64
	SetCreationType(v string) *GetWorkItemResponseBody
	GetCreationType() *string
	SetCurrentPlanId(v string) *GetWorkItemResponseBody
	GetCurrentPlanId() *string
	SetCurrentRevision(v int64) *GetWorkItemResponseBody
	GetCurrentRevision() *int64
	SetCurrentRunId(v string) *GetWorkItemResponseBody
	GetCurrentRunId() *string
	SetCurrentRunSeq(v int64) *GetWorkItemResponseBody
	GetCurrentRunSeq() *int64
	SetCurrentStepName(v string) *GetWorkItemResponseBody
	GetCurrentStepName() *string
	SetDescription(v string) *GetWorkItemResponseBody
	GetDescription() *string
	SetInputVersion(v int64) *GetWorkItemResponseBody
	GetInputVersion() *int64
	SetItemType(v string) *GetWorkItemResponseBody
	GetItemType() *string
	SetLastProcessingAt(v string) *GetWorkItemResponseBody
	GetLastProcessingAt() *string
	SetLegacyTopic(v bool) *GetWorkItemResponseBody
	GetLegacyTopic() *bool
	SetMemberUserIds(v []*int64) *GetWorkItemResponseBody
	GetMemberUserIds() []*int64
	SetMessage(v string) *GetWorkItemResponseBody
	GetMessage() *string
	SetModifiedAt(v string) *GetWorkItemResponseBody
	GetModifiedAt() *string
	SetOwnerUserId(v int64) *GetWorkItemResponseBody
	GetOwnerUserId() *int64
	SetPinned(v bool) *GetWorkItemResponseBody
	GetPinned() *bool
	SetPriority(v string) *GetWorkItemResponseBody
	GetPriority() *string
	SetProgressCompleted(v int64) *GetWorkItemResponseBody
	GetProgressCompleted() *int64
	SetProgressTotal(v int64) *GetWorkItemResponseBody
	GetProgressTotal() *int64
	SetRequestId(v string) *GetWorkItemResponseBody
	GetRequestId() *string
	SetSourceExecutionId(v string) *GetWorkItemResponseBody
	GetSourceExecutionId() *string
	SetSourceTaskId(v string) *GetWorkItemResponseBody
	GetSourceTaskId() *string
	SetSourceTaskName(v string) *GetWorkItemResponseBody
	GetSourceTaskName() *string
	SetSourceTriggerKey(v string) *GetWorkItemResponseBody
	GetSourceTriggerKey() *string
	SetSourceType(v string) *GetWorkItemResponseBody
	GetSourceType() *string
	SetTitle(v string) *GetWorkItemResponseBody
	GetTitle() *string
	SetTriggerTimezone(v string) *GetWorkItemResponseBody
	GetTriggerTimezone() *string
	SetTriggeredAt(v string) *GetWorkItemResponseBody
	GetTriggeredAt() *string
	SetVersion(v int64) *GetWorkItemResponseBody
	GetVersion() *int64
	SetVisibility(v string) *GetWorkItemResponseBody
	GetVisibility() *string
	SetVisibilitySource(v string) *GetWorkItemResponseBody
	GetVisibilitySource() *string
	SetWorkItemId(v string) *GetWorkItemResponseBody
	GetWorkItemId() *string
	SetWorkflowStatus(v string) *GetWorkItemResponseBody
	GetWorkflowStatus() *string
	SetWorkflowTemplateIds(v []*string) *GetWorkItemResponseBody
	GetWorkflowTemplateIds() []*string
}

type GetWorkItemResponseBody struct {
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
	Attachments []*GetWorkItemResponseBodyAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
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

func (s GetWorkItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkItemResponseBody) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *GetWorkItemResponseBody) GetAppliedStatusVersion() *int64 {
	return s.AppliedStatusVersion
}

func (s *GetWorkItemResponseBody) GetAttachments() []*GetWorkItemResponseBodyAttachments {
	return s.Attachments
}

func (s *GetWorkItemResponseBody) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *GetWorkItemResponseBody) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *GetWorkItemResponseBody) GetCanManage() *bool {
	return s.CanManage
}

func (s *GetWorkItemResponseBody) GetCanManageVisibility() *bool {
	return s.CanManageVisibility
}

func (s *GetWorkItemResponseBody) GetCanRename() *bool {
	return s.CanRename
}

func (s *GetWorkItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkItemResponseBody) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkItemResponseBody) GetCollaboratorUserIds() []*int64 {
	return s.CollaboratorUserIds
}

func (s *GetWorkItemResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetWorkItemResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetWorkItemResponseBody) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *GetWorkItemResponseBody) GetCreationType() *string {
	return s.CreationType
}

func (s *GetWorkItemResponseBody) GetCurrentPlanId() *string {
	return s.CurrentPlanId
}

func (s *GetWorkItemResponseBody) GetCurrentRevision() *int64 {
	return s.CurrentRevision
}

func (s *GetWorkItemResponseBody) GetCurrentRunId() *string {
	return s.CurrentRunId
}

func (s *GetWorkItemResponseBody) GetCurrentRunSeq() *int64 {
	return s.CurrentRunSeq
}

func (s *GetWorkItemResponseBody) GetCurrentStepName() *string {
	return s.CurrentStepName
}

func (s *GetWorkItemResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetWorkItemResponseBody) GetInputVersion() *int64 {
	return s.InputVersion
}

func (s *GetWorkItemResponseBody) GetItemType() *string {
	return s.ItemType
}

func (s *GetWorkItemResponseBody) GetLastProcessingAt() *string {
	return s.LastProcessingAt
}

func (s *GetWorkItemResponseBody) GetLegacyTopic() *bool {
	return s.LegacyTopic
}

func (s *GetWorkItemResponseBody) GetMemberUserIds() []*int64 {
	return s.MemberUserIds
}

func (s *GetWorkItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkItemResponseBody) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *GetWorkItemResponseBody) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *GetWorkItemResponseBody) GetPinned() *bool {
	return s.Pinned
}

func (s *GetWorkItemResponseBody) GetPriority() *string {
	return s.Priority
}

func (s *GetWorkItemResponseBody) GetProgressCompleted() *int64 {
	return s.ProgressCompleted
}

func (s *GetWorkItemResponseBody) GetProgressTotal() *int64 {
	return s.ProgressTotal
}

func (s *GetWorkItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkItemResponseBody) GetSourceExecutionId() *string {
	return s.SourceExecutionId
}

func (s *GetWorkItemResponseBody) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *GetWorkItemResponseBody) GetSourceTaskName() *string {
	return s.SourceTaskName
}

func (s *GetWorkItemResponseBody) GetSourceTriggerKey() *string {
	return s.SourceTriggerKey
}

func (s *GetWorkItemResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetWorkItemResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetWorkItemResponseBody) GetTriggerTimezone() *string {
	return s.TriggerTimezone
}

func (s *GetWorkItemResponseBody) GetTriggeredAt() *string {
	return s.TriggeredAt
}

func (s *GetWorkItemResponseBody) GetVersion() *int64 {
	return s.Version
}

func (s *GetWorkItemResponseBody) GetVisibility() *string {
	return s.Visibility
}

func (s *GetWorkItemResponseBody) GetVisibilitySource() *string {
	return s.VisibilitySource
}

func (s *GetWorkItemResponseBody) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetWorkItemResponseBody) GetWorkflowStatus() *string {
	return s.WorkflowStatus
}

func (s *GetWorkItemResponseBody) GetWorkflowTemplateIds() []*string {
	return s.WorkflowTemplateIds
}

func (s *GetWorkItemResponseBody) SetAcceptanceCriteria(v string) *GetWorkItemResponseBody {
	s.AcceptanceCriteria = &v
	return s
}

func (s *GetWorkItemResponseBody) SetAppliedStatusVersion(v int64) *GetWorkItemResponseBody {
	s.AppliedStatusVersion = &v
	return s
}

func (s *GetWorkItemResponseBody) SetAttachments(v []*GetWorkItemResponseBodyAttachments) *GetWorkItemResponseBody {
	s.Attachments = v
	return s
}

func (s *GetWorkItemResponseBody) SetBusinessStatus(v string) *GetWorkItemResponseBody {
	s.BusinessStatus = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCanDelete(v bool) *GetWorkItemResponseBody {
	s.CanDelete = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCanManage(v bool) *GetWorkItemResponseBody {
	s.CanManage = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCanManageVisibility(v bool) *GetWorkItemResponseBody {
	s.CanManageVisibility = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCanRename(v bool) *GetWorkItemResponseBody {
	s.CanRename = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCode(v string) *GetWorkItemResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCollaborationGroupId(v string) *GetWorkItemResponseBody {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCollaboratorUserIds(v []*int64) *GetWorkItemResponseBody {
	s.CollaboratorUserIds = v
	return s
}

func (s *GetWorkItemResponseBody) SetConversationId(v string) *GetWorkItemResponseBody {
	s.ConversationId = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCreatedAt(v string) *GetWorkItemResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCreatedByUserId(v int64) *GetWorkItemResponseBody {
	s.CreatedByUserId = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCreationType(v string) *GetWorkItemResponseBody {
	s.CreationType = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCurrentPlanId(v string) *GetWorkItemResponseBody {
	s.CurrentPlanId = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCurrentRevision(v int64) *GetWorkItemResponseBody {
	s.CurrentRevision = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCurrentRunId(v string) *GetWorkItemResponseBody {
	s.CurrentRunId = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCurrentRunSeq(v int64) *GetWorkItemResponseBody {
	s.CurrentRunSeq = &v
	return s
}

func (s *GetWorkItemResponseBody) SetCurrentStepName(v string) *GetWorkItemResponseBody {
	s.CurrentStepName = &v
	return s
}

func (s *GetWorkItemResponseBody) SetDescription(v string) *GetWorkItemResponseBody {
	s.Description = &v
	return s
}

func (s *GetWorkItemResponseBody) SetInputVersion(v int64) *GetWorkItemResponseBody {
	s.InputVersion = &v
	return s
}

func (s *GetWorkItemResponseBody) SetItemType(v string) *GetWorkItemResponseBody {
	s.ItemType = &v
	return s
}

func (s *GetWorkItemResponseBody) SetLastProcessingAt(v string) *GetWorkItemResponseBody {
	s.LastProcessingAt = &v
	return s
}

func (s *GetWorkItemResponseBody) SetLegacyTopic(v bool) *GetWorkItemResponseBody {
	s.LegacyTopic = &v
	return s
}

func (s *GetWorkItemResponseBody) SetMemberUserIds(v []*int64) *GetWorkItemResponseBody {
	s.MemberUserIds = v
	return s
}

func (s *GetWorkItemResponseBody) SetMessage(v string) *GetWorkItemResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkItemResponseBody) SetModifiedAt(v string) *GetWorkItemResponseBody {
	s.ModifiedAt = &v
	return s
}

func (s *GetWorkItemResponseBody) SetOwnerUserId(v int64) *GetWorkItemResponseBody {
	s.OwnerUserId = &v
	return s
}

func (s *GetWorkItemResponseBody) SetPinned(v bool) *GetWorkItemResponseBody {
	s.Pinned = &v
	return s
}

func (s *GetWorkItemResponseBody) SetPriority(v string) *GetWorkItemResponseBody {
	s.Priority = &v
	return s
}

func (s *GetWorkItemResponseBody) SetProgressCompleted(v int64) *GetWorkItemResponseBody {
	s.ProgressCompleted = &v
	return s
}

func (s *GetWorkItemResponseBody) SetProgressTotal(v int64) *GetWorkItemResponseBody {
	s.ProgressTotal = &v
	return s
}

func (s *GetWorkItemResponseBody) SetRequestId(v string) *GetWorkItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkItemResponseBody) SetSourceExecutionId(v string) *GetWorkItemResponseBody {
	s.SourceExecutionId = &v
	return s
}

func (s *GetWorkItemResponseBody) SetSourceTaskId(v string) *GetWorkItemResponseBody {
	s.SourceTaskId = &v
	return s
}

func (s *GetWorkItemResponseBody) SetSourceTaskName(v string) *GetWorkItemResponseBody {
	s.SourceTaskName = &v
	return s
}

func (s *GetWorkItemResponseBody) SetSourceTriggerKey(v string) *GetWorkItemResponseBody {
	s.SourceTriggerKey = &v
	return s
}

func (s *GetWorkItemResponseBody) SetSourceType(v string) *GetWorkItemResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetWorkItemResponseBody) SetTitle(v string) *GetWorkItemResponseBody {
	s.Title = &v
	return s
}

func (s *GetWorkItemResponseBody) SetTriggerTimezone(v string) *GetWorkItemResponseBody {
	s.TriggerTimezone = &v
	return s
}

func (s *GetWorkItemResponseBody) SetTriggeredAt(v string) *GetWorkItemResponseBody {
	s.TriggeredAt = &v
	return s
}

func (s *GetWorkItemResponseBody) SetVersion(v int64) *GetWorkItemResponseBody {
	s.Version = &v
	return s
}

func (s *GetWorkItemResponseBody) SetVisibility(v string) *GetWorkItemResponseBody {
	s.Visibility = &v
	return s
}

func (s *GetWorkItemResponseBody) SetVisibilitySource(v string) *GetWorkItemResponseBody {
	s.VisibilitySource = &v
	return s
}

func (s *GetWorkItemResponseBody) SetWorkItemId(v string) *GetWorkItemResponseBody {
	s.WorkItemId = &v
	return s
}

func (s *GetWorkItemResponseBody) SetWorkflowStatus(v string) *GetWorkItemResponseBody {
	s.WorkflowStatus = &v
	return s
}

func (s *GetWorkItemResponseBody) SetWorkflowTemplateIds(v []*string) *GetWorkItemResponseBody {
	s.WorkflowTemplateIds = v
	return s
}

func (s *GetWorkItemResponseBody) Validate() error {
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

type GetWorkItemResponseBodyAttachments struct {
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

func (s GetWorkItemResponseBodyAttachments) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemResponseBodyAttachments) GoString() string {
	return s.String()
}

func (s *GetWorkItemResponseBodyAttachments) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *GetWorkItemResponseBodyAttachments) GetContentType() *string {
	return s.ContentType
}

func (s *GetWorkItemResponseBodyAttachments) GetRefUrl() *string {
	return s.RefUrl
}

func (s *GetWorkItemResponseBodyAttachments) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetWorkItemResponseBodyAttachments) GetResourceKind() *string {
	return s.ResourceKind
}

func (s *GetWorkItemResponseBodyAttachments) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *GetWorkItemResponseBodyAttachments) GetTitle() *string {
	return s.Title
}

func (s *GetWorkItemResponseBodyAttachments) SetAttachmentId(v string) *GetWorkItemResponseBodyAttachments {
	s.AttachmentId = &v
	return s
}

func (s *GetWorkItemResponseBodyAttachments) SetContentType(v string) *GetWorkItemResponseBodyAttachments {
	s.ContentType = &v
	return s
}

func (s *GetWorkItemResponseBodyAttachments) SetRefUrl(v string) *GetWorkItemResponseBodyAttachments {
	s.RefUrl = &v
	return s
}

func (s *GetWorkItemResponseBodyAttachments) SetResourceId(v string) *GetWorkItemResponseBodyAttachments {
	s.ResourceId = &v
	return s
}

func (s *GetWorkItemResponseBodyAttachments) SetResourceKind(v string) *GetWorkItemResponseBodyAttachments {
	s.ResourceKind = &v
	return s
}

func (s *GetWorkItemResponseBodyAttachments) SetSizeBytes(v int64) *GetWorkItemResponseBodyAttachments {
	s.SizeBytes = &v
	return s
}

func (s *GetWorkItemResponseBodyAttachments) SetTitle(v string) *GetWorkItemResponseBodyAttachments {
	s.Title = &v
	return s
}

func (s *GetWorkItemResponseBodyAttachments) Validate() error {
	return dara.Validate(s)
}
