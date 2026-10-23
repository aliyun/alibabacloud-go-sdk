// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkItemPinnedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptanceCriteria(v string) *SetWorkItemPinnedResponseBody
	GetAcceptanceCriteria() *string
	SetAppliedStatusVersion(v int64) *SetWorkItemPinnedResponseBody
	GetAppliedStatusVersion() *int64
	SetAttachments(v []*SetWorkItemPinnedResponseBodyAttachments) *SetWorkItemPinnedResponseBody
	GetAttachments() []*SetWorkItemPinnedResponseBodyAttachments
	SetBusinessStatus(v string) *SetWorkItemPinnedResponseBody
	GetBusinessStatus() *string
	SetCanDelete(v bool) *SetWorkItemPinnedResponseBody
	GetCanDelete() *bool
	SetCanManage(v bool) *SetWorkItemPinnedResponseBody
	GetCanManage() *bool
	SetCanManageVisibility(v bool) *SetWorkItemPinnedResponseBody
	GetCanManageVisibility() *bool
	SetCanRename(v bool) *SetWorkItemPinnedResponseBody
	GetCanRename() *bool
	SetCode(v string) *SetWorkItemPinnedResponseBody
	GetCode() *string
	SetCollaborationGroupId(v string) *SetWorkItemPinnedResponseBody
	GetCollaborationGroupId() *string
	SetCollaboratorUserIds(v []*int64) *SetWorkItemPinnedResponseBody
	GetCollaboratorUserIds() []*int64
	SetConversationId(v string) *SetWorkItemPinnedResponseBody
	GetConversationId() *string
	SetCreatedAt(v string) *SetWorkItemPinnedResponseBody
	GetCreatedAt() *string
	SetCreatedByUserId(v int64) *SetWorkItemPinnedResponseBody
	GetCreatedByUserId() *int64
	SetCreationType(v string) *SetWorkItemPinnedResponseBody
	GetCreationType() *string
	SetCurrentPlanId(v string) *SetWorkItemPinnedResponseBody
	GetCurrentPlanId() *string
	SetCurrentRevision(v int64) *SetWorkItemPinnedResponseBody
	GetCurrentRevision() *int64
	SetCurrentRunId(v string) *SetWorkItemPinnedResponseBody
	GetCurrentRunId() *string
	SetCurrentRunSeq(v int64) *SetWorkItemPinnedResponseBody
	GetCurrentRunSeq() *int64
	SetCurrentStepName(v string) *SetWorkItemPinnedResponseBody
	GetCurrentStepName() *string
	SetDescription(v string) *SetWorkItemPinnedResponseBody
	GetDescription() *string
	SetInputVersion(v int64) *SetWorkItemPinnedResponseBody
	GetInputVersion() *int64
	SetItemType(v string) *SetWorkItemPinnedResponseBody
	GetItemType() *string
	SetLastProcessingAt(v string) *SetWorkItemPinnedResponseBody
	GetLastProcessingAt() *string
	SetLegacyTopic(v bool) *SetWorkItemPinnedResponseBody
	GetLegacyTopic() *bool
	SetMemberUserIds(v []*int64) *SetWorkItemPinnedResponseBody
	GetMemberUserIds() []*int64
	SetMessage(v string) *SetWorkItemPinnedResponseBody
	GetMessage() *string
	SetModifiedAt(v string) *SetWorkItemPinnedResponseBody
	GetModifiedAt() *string
	SetOwnerUserId(v int64) *SetWorkItemPinnedResponseBody
	GetOwnerUserId() *int64
	SetPinned(v bool) *SetWorkItemPinnedResponseBody
	GetPinned() *bool
	SetPriority(v string) *SetWorkItemPinnedResponseBody
	GetPriority() *string
	SetProgressCompleted(v int64) *SetWorkItemPinnedResponseBody
	GetProgressCompleted() *int64
	SetProgressTotal(v int64) *SetWorkItemPinnedResponseBody
	GetProgressTotal() *int64
	SetRequestId(v string) *SetWorkItemPinnedResponseBody
	GetRequestId() *string
	SetSourceExecutionId(v string) *SetWorkItemPinnedResponseBody
	GetSourceExecutionId() *string
	SetSourceTaskId(v string) *SetWorkItemPinnedResponseBody
	GetSourceTaskId() *string
	SetSourceTaskName(v string) *SetWorkItemPinnedResponseBody
	GetSourceTaskName() *string
	SetSourceTriggerKey(v string) *SetWorkItemPinnedResponseBody
	GetSourceTriggerKey() *string
	SetSourceType(v string) *SetWorkItemPinnedResponseBody
	GetSourceType() *string
	SetTitle(v string) *SetWorkItemPinnedResponseBody
	GetTitle() *string
	SetTriggerTimezone(v string) *SetWorkItemPinnedResponseBody
	GetTriggerTimezone() *string
	SetTriggeredAt(v string) *SetWorkItemPinnedResponseBody
	GetTriggeredAt() *string
	SetVersion(v int64) *SetWorkItemPinnedResponseBody
	GetVersion() *int64
	SetVisibility(v string) *SetWorkItemPinnedResponseBody
	GetVisibility() *string
	SetVisibilitySource(v string) *SetWorkItemPinnedResponseBody
	GetVisibilitySource() *string
	SetWorkItemId(v string) *SetWorkItemPinnedResponseBody
	GetWorkItemId() *string
	SetWorkflowStatus(v string) *SetWorkItemPinnedResponseBody
	GetWorkflowStatus() *string
	SetWorkflowTemplateIds(v []*string) *SetWorkItemPinnedResponseBody
	GetWorkflowTemplateIds() []*string
}

type SetWorkItemPinnedResponseBody struct {
	// The acceptance criteria.
	//
	// example:
	//
	// example
	AcceptanceCriteria *string `json:"acceptanceCriteria,omitempty" xml:"acceptanceCriteria,omitempty"`
	// The applied status version.
	//
	// example:
	//
	// 1
	AppliedStatusVersion *int64 `json:"appliedStatusVersion,omitempty" xml:"appliedStatusVersion,omitempty"`
	// The list of attachments.
	//
	// example:
	//
	// example
	Attachments []*SetWorkItemPinnedResponseBodyAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// The business status (seven states).
	//
	// example:
	//
	// PROCESSING
	BusinessStatus *string `json:"businessStatus,omitempty" xml:"businessStatus,omitempty"`
	// Specifies whether the caller can delete the work item.
	//
	// example:
	//
	// false
	CanDelete *bool `json:"canDelete,omitempty" xml:"canDelete,omitempty"`
	// Specifies whether the caller can manage the work item.
	//
	// example:
	//
	// false
	CanManage *bool `json:"canManage,omitempty" xml:"canManage,omitempty"`
	// Specifies whether the visibility scope can be adjusted independently.
	//
	// example:
	//
	// false
	CanManageVisibility *bool `json:"canManageVisibility,omitempty" xml:"canManageVisibility,omitempty"`
	// Specifies whether the caller can rename the work item.
	//
	// example:
	//
	// false
	CanRename *bool `json:"canRename,omitempty" xml:"canRename,omitempty"`
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The ID of the collaboration workspace.
	//
	// example:
	//
	// example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The list of user IDs of the collaborators.
	//
	// example:
	//
	// example
	CollaboratorUserIds []*int64 `json:"collaboratorUserIds,omitempty" xml:"collaboratorUserIds,omitempty" type:"Repeated"`
	// The ID of the associated conversation.
	//
	// example:
	//
	// example
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// The creation time (ISO 8601 format).
	//
	// example:
	//
	// example
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The user ID of the creator.
	//
	// example:
	//
	// 1
	CreatedByUserId *int64 `json:"createdByUserId,omitempty" xml:"createdByUserId,omitempty"`
	// The creation source. Valid values: DEFAULT and TASK.
	//
	// example:
	//
	// DEFAULT
	CreationType *string `json:"creationType,omitempty" xml:"creationType,omitempty"`
	// The ID of the current physical plan.
	//
	// example:
	//
	// example
	CurrentPlanId *string `json:"currentPlanId,omitempty" xml:"currentPlanId,omitempty"`
	// The current revision number.
	//
	// example:
	//
	// 1
	CurrentRevision *int64 `json:"currentRevision,omitempty" xml:"currentRevision,omitempty"`
	// The ID of the current run.
	//
	// example:
	//
	// example
	CurrentRunId *string `json:"currentRunId,omitempty" xml:"currentRunId,omitempty"`
	// The sequence number of the current run.
	//
	// example:
	//
	// 1
	CurrentRunSeq *int64 `json:"currentRunSeq,omitempty" xml:"currentRunSeq,omitempty"`
	// The name of the current step.
	//
	// example:
	//
	// example
	CurrentStepName *string `json:"currentStepName,omitempty" xml:"currentStepName,omitempty"`
	// The description.
	//
	// example:
	//
	// example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The input snapshot version.
	//
	// example:
	//
	// 1
	InputVersion *int64 `json:"inputVersion,omitempty" xml:"inputVersion,omitempty"`
	// The type of the work item.
	//
	// example:
	//
	// WORKFLOW
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// The time when the work item last entered the processing state (ISO 8601 format).
	//
	// example:
	//
	// example
	LastProcessingAt *string `json:"lastProcessingAt,omitempty" xml:"lastProcessingAt,omitempty"`
	// Specifies whether the topic is a free exploration topic during the compatibility period.
	//
	// example:
	//
	// false
	LegacyTopic *bool `json:"legacyTopic,omitempty" xml:"legacyTopic,omitempty"`
	// The list of user IDs of the visible members.
	//
	// example:
	//
	// example
	MemberUserIds []*int64 `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The modification time (ISO 8601 format).
	//
	// example:
	//
	// example
	ModifiedAt *string `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// The user ID of the owner.
	//
	// example:
	//
	// 1
	OwnerUserId *int64 `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// The personal pin status of the caller.
	//
	// example:
	//
	// false
	Pinned *bool `json:"pinned,omitempty" xml:"pinned,omitempty"`
	// The priority.
	//
	// example:
	//
	// HIGH
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// The completed progress.
	//
	// example:
	//
	// 1
	ProgressCompleted *int64 `json:"progressCompleted,omitempty" xml:"progressCompleted,omitempty"`
	// The total progress.
	//
	// example:
	//
	// 1
	ProgressTotal *int64 `json:"progressTotal,omitempty" xml:"progressTotal,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The execution ID of the task.
	//
	// example:
	//
	// example
	SourceExecutionId *string `json:"sourceExecutionId,omitempty" xml:"sourceExecutionId,omitempty"`
	// The source ID of the task.
	//
	// example:
	//
	// example
	SourceTaskId *string `json:"sourceTaskId,omitempty" xml:"sourceTaskId,omitempty"`
	// The source name of the task.
	//
	// example:
	//
	// example
	SourceTaskName *string `json:"sourceTaskName,omitempty" xml:"sourceTaskName,omitempty"`
	// The trigger key of the task.
	//
	// example:
	//
	// example
	SourceTriggerKey *string `json:"sourceTriggerKey,omitempty" xml:"sourceTriggerKey,omitempty"`
	// The source type of the task.
	//
	// example:
	//
	// example
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The title.
	//
	// example:
	//
	// example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The time zone when the task is triggered.
	//
	// example:
	//
	// example
	TriggerTimezone *string `json:"triggerTimezone,omitempty" xml:"triggerTimezone,omitempty"`
	// The trigger time of the task (ISO 8601 format).
	//
	// example:
	//
	// example
	TriggeredAt *string `json:"triggeredAt,omitempty" xml:"triggeredAt,omitempty"`
	// The optimistic version number.
	//
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
	// The visibility scope.
	//
	// example:
	//
	// example
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
	// The source of the visibility scope. Valid values: SELF, TASK, and TEMPLATE.
	//
	// example:
	//
	// SELF
	VisibilitySource *string `json:"visibilitySource,omitempty" xml:"visibilitySource,omitempty"`
	// The ID of the work item.
	//
	// example:
	//
	// example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
	// The four-state projection of the workflow status.
	//
	// example:
	//
	// example
	WorkflowStatus *string `json:"workflowStatus,omitempty" xml:"workflowStatus,omitempty"`
	// The list of IDs of the bound workflow templates.
	//
	// example:
	//
	// example
	WorkflowTemplateIds []*string `json:"workflowTemplateIds,omitempty" xml:"workflowTemplateIds,omitempty" type:"Repeated"`
}

func (s SetWorkItemPinnedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetWorkItemPinnedResponseBody) GoString() string {
	return s.String()
}

func (s *SetWorkItemPinnedResponseBody) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *SetWorkItemPinnedResponseBody) GetAppliedStatusVersion() *int64 {
	return s.AppliedStatusVersion
}

func (s *SetWorkItemPinnedResponseBody) GetAttachments() []*SetWorkItemPinnedResponseBodyAttachments {
	return s.Attachments
}

func (s *SetWorkItemPinnedResponseBody) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *SetWorkItemPinnedResponseBody) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *SetWorkItemPinnedResponseBody) GetCanManage() *bool {
	return s.CanManage
}

func (s *SetWorkItemPinnedResponseBody) GetCanManageVisibility() *bool {
	return s.CanManageVisibility
}

func (s *SetWorkItemPinnedResponseBody) GetCanRename() *bool {
	return s.CanRename
}

func (s *SetWorkItemPinnedResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetWorkItemPinnedResponseBody) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *SetWorkItemPinnedResponseBody) GetCollaboratorUserIds() []*int64 {
	return s.CollaboratorUserIds
}

func (s *SetWorkItemPinnedResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *SetWorkItemPinnedResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *SetWorkItemPinnedResponseBody) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *SetWorkItemPinnedResponseBody) GetCreationType() *string {
	return s.CreationType
}

func (s *SetWorkItemPinnedResponseBody) GetCurrentPlanId() *string {
	return s.CurrentPlanId
}

func (s *SetWorkItemPinnedResponseBody) GetCurrentRevision() *int64 {
	return s.CurrentRevision
}

func (s *SetWorkItemPinnedResponseBody) GetCurrentRunId() *string {
	return s.CurrentRunId
}

func (s *SetWorkItemPinnedResponseBody) GetCurrentRunSeq() *int64 {
	return s.CurrentRunSeq
}

func (s *SetWorkItemPinnedResponseBody) GetCurrentStepName() *string {
	return s.CurrentStepName
}

func (s *SetWorkItemPinnedResponseBody) GetDescription() *string {
	return s.Description
}

func (s *SetWorkItemPinnedResponseBody) GetInputVersion() *int64 {
	return s.InputVersion
}

func (s *SetWorkItemPinnedResponseBody) GetItemType() *string {
	return s.ItemType
}

func (s *SetWorkItemPinnedResponseBody) GetLastProcessingAt() *string {
	return s.LastProcessingAt
}

func (s *SetWorkItemPinnedResponseBody) GetLegacyTopic() *bool {
	return s.LegacyTopic
}

func (s *SetWorkItemPinnedResponseBody) GetMemberUserIds() []*int64 {
	return s.MemberUserIds
}

func (s *SetWorkItemPinnedResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetWorkItemPinnedResponseBody) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *SetWorkItemPinnedResponseBody) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *SetWorkItemPinnedResponseBody) GetPinned() *bool {
	return s.Pinned
}

func (s *SetWorkItemPinnedResponseBody) GetPriority() *string {
	return s.Priority
}

func (s *SetWorkItemPinnedResponseBody) GetProgressCompleted() *int64 {
	return s.ProgressCompleted
}

func (s *SetWorkItemPinnedResponseBody) GetProgressTotal() *int64 {
	return s.ProgressTotal
}

func (s *SetWorkItemPinnedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetWorkItemPinnedResponseBody) GetSourceExecutionId() *string {
	return s.SourceExecutionId
}

func (s *SetWorkItemPinnedResponseBody) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *SetWorkItemPinnedResponseBody) GetSourceTaskName() *string {
	return s.SourceTaskName
}

func (s *SetWorkItemPinnedResponseBody) GetSourceTriggerKey() *string {
	return s.SourceTriggerKey
}

func (s *SetWorkItemPinnedResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *SetWorkItemPinnedResponseBody) GetTitle() *string {
	return s.Title
}

func (s *SetWorkItemPinnedResponseBody) GetTriggerTimezone() *string {
	return s.TriggerTimezone
}

func (s *SetWorkItemPinnedResponseBody) GetTriggeredAt() *string {
	return s.TriggeredAt
}

func (s *SetWorkItemPinnedResponseBody) GetVersion() *int64 {
	return s.Version
}

func (s *SetWorkItemPinnedResponseBody) GetVisibility() *string {
	return s.Visibility
}

func (s *SetWorkItemPinnedResponseBody) GetVisibilitySource() *string {
	return s.VisibilitySource
}

func (s *SetWorkItemPinnedResponseBody) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *SetWorkItemPinnedResponseBody) GetWorkflowStatus() *string {
	return s.WorkflowStatus
}

func (s *SetWorkItemPinnedResponseBody) GetWorkflowTemplateIds() []*string {
	return s.WorkflowTemplateIds
}

func (s *SetWorkItemPinnedResponseBody) SetAcceptanceCriteria(v string) *SetWorkItemPinnedResponseBody {
	s.AcceptanceCriteria = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetAppliedStatusVersion(v int64) *SetWorkItemPinnedResponseBody {
	s.AppliedStatusVersion = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetAttachments(v []*SetWorkItemPinnedResponseBodyAttachments) *SetWorkItemPinnedResponseBody {
	s.Attachments = v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetBusinessStatus(v string) *SetWorkItemPinnedResponseBody {
	s.BusinessStatus = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCanDelete(v bool) *SetWorkItemPinnedResponseBody {
	s.CanDelete = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCanManage(v bool) *SetWorkItemPinnedResponseBody {
	s.CanManage = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCanManageVisibility(v bool) *SetWorkItemPinnedResponseBody {
	s.CanManageVisibility = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCanRename(v bool) *SetWorkItemPinnedResponseBody {
	s.CanRename = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCode(v string) *SetWorkItemPinnedResponseBody {
	s.Code = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCollaborationGroupId(v string) *SetWorkItemPinnedResponseBody {
	s.CollaborationGroupId = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCollaboratorUserIds(v []*int64) *SetWorkItemPinnedResponseBody {
	s.CollaboratorUserIds = v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetConversationId(v string) *SetWorkItemPinnedResponseBody {
	s.ConversationId = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCreatedAt(v string) *SetWorkItemPinnedResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCreatedByUserId(v int64) *SetWorkItemPinnedResponseBody {
	s.CreatedByUserId = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCreationType(v string) *SetWorkItemPinnedResponseBody {
	s.CreationType = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCurrentPlanId(v string) *SetWorkItemPinnedResponseBody {
	s.CurrentPlanId = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCurrentRevision(v int64) *SetWorkItemPinnedResponseBody {
	s.CurrentRevision = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCurrentRunId(v string) *SetWorkItemPinnedResponseBody {
	s.CurrentRunId = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCurrentRunSeq(v int64) *SetWorkItemPinnedResponseBody {
	s.CurrentRunSeq = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetCurrentStepName(v string) *SetWorkItemPinnedResponseBody {
	s.CurrentStepName = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetDescription(v string) *SetWorkItemPinnedResponseBody {
	s.Description = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetInputVersion(v int64) *SetWorkItemPinnedResponseBody {
	s.InputVersion = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetItemType(v string) *SetWorkItemPinnedResponseBody {
	s.ItemType = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetLastProcessingAt(v string) *SetWorkItemPinnedResponseBody {
	s.LastProcessingAt = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetLegacyTopic(v bool) *SetWorkItemPinnedResponseBody {
	s.LegacyTopic = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetMemberUserIds(v []*int64) *SetWorkItemPinnedResponseBody {
	s.MemberUserIds = v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetMessage(v string) *SetWorkItemPinnedResponseBody {
	s.Message = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetModifiedAt(v string) *SetWorkItemPinnedResponseBody {
	s.ModifiedAt = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetOwnerUserId(v int64) *SetWorkItemPinnedResponseBody {
	s.OwnerUserId = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetPinned(v bool) *SetWorkItemPinnedResponseBody {
	s.Pinned = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetPriority(v string) *SetWorkItemPinnedResponseBody {
	s.Priority = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetProgressCompleted(v int64) *SetWorkItemPinnedResponseBody {
	s.ProgressCompleted = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetProgressTotal(v int64) *SetWorkItemPinnedResponseBody {
	s.ProgressTotal = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetRequestId(v string) *SetWorkItemPinnedResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetSourceExecutionId(v string) *SetWorkItemPinnedResponseBody {
	s.SourceExecutionId = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetSourceTaskId(v string) *SetWorkItemPinnedResponseBody {
	s.SourceTaskId = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetSourceTaskName(v string) *SetWorkItemPinnedResponseBody {
	s.SourceTaskName = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetSourceTriggerKey(v string) *SetWorkItemPinnedResponseBody {
	s.SourceTriggerKey = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetSourceType(v string) *SetWorkItemPinnedResponseBody {
	s.SourceType = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetTitle(v string) *SetWorkItemPinnedResponseBody {
	s.Title = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetTriggerTimezone(v string) *SetWorkItemPinnedResponseBody {
	s.TriggerTimezone = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetTriggeredAt(v string) *SetWorkItemPinnedResponseBody {
	s.TriggeredAt = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetVersion(v int64) *SetWorkItemPinnedResponseBody {
	s.Version = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetVisibility(v string) *SetWorkItemPinnedResponseBody {
	s.Visibility = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetVisibilitySource(v string) *SetWorkItemPinnedResponseBody {
	s.VisibilitySource = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetWorkItemId(v string) *SetWorkItemPinnedResponseBody {
	s.WorkItemId = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetWorkflowStatus(v string) *SetWorkItemPinnedResponseBody {
	s.WorkflowStatus = &v
	return s
}

func (s *SetWorkItemPinnedResponseBody) SetWorkflowTemplateIds(v []*string) *SetWorkItemPinnedResponseBody {
	s.WorkflowTemplateIds = v
	return s
}

func (s *SetWorkItemPinnedResponseBody) Validate() error {
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

type SetWorkItemPinnedResponseBodyAttachments struct {
	// The ID of the attachment.
	//
	// example:
	//
	// example
	AttachmentId *string `json:"attachmentId,omitempty" xml:"attachmentId,omitempty"`
	// The content type.
	//
	// example:
	//
	// example
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The access URL of the attachment (OpenAPI 302 redirect).
	//
	// example:
	//
	// example
	RefUrl *string `json:"refUrl,omitempty" xml:"refUrl,omitempty"`
	// The resource ID of the attachment.
	//
	// example:
	//
	// example
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type of the attachment.
	//
	// example:
	//
	// example
	ResourceKind *string `json:"resourceKind,omitempty" xml:"resourceKind,omitempty"`
	// The size of the attachment in bytes.
	//
	// example:
	//
	// 1
	SizeBytes *int64 `json:"sizeBytes,omitempty" xml:"sizeBytes,omitempty"`
	// The title of the attachment.
	//
	// example:
	//
	// example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SetWorkItemPinnedResponseBodyAttachments) String() string {
	return dara.Prettify(s)
}

func (s SetWorkItemPinnedResponseBodyAttachments) GoString() string {
	return s.String()
}

func (s *SetWorkItemPinnedResponseBodyAttachments) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *SetWorkItemPinnedResponseBodyAttachments) GetContentType() *string {
	return s.ContentType
}

func (s *SetWorkItemPinnedResponseBodyAttachments) GetRefUrl() *string {
	return s.RefUrl
}

func (s *SetWorkItemPinnedResponseBodyAttachments) GetResourceId() *string {
	return s.ResourceId
}

func (s *SetWorkItemPinnedResponseBodyAttachments) GetResourceKind() *string {
	return s.ResourceKind
}

func (s *SetWorkItemPinnedResponseBodyAttachments) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *SetWorkItemPinnedResponseBodyAttachments) GetTitle() *string {
	return s.Title
}

func (s *SetWorkItemPinnedResponseBodyAttachments) SetAttachmentId(v string) *SetWorkItemPinnedResponseBodyAttachments {
	s.AttachmentId = &v
	return s
}

func (s *SetWorkItemPinnedResponseBodyAttachments) SetContentType(v string) *SetWorkItemPinnedResponseBodyAttachments {
	s.ContentType = &v
	return s
}

func (s *SetWorkItemPinnedResponseBodyAttachments) SetRefUrl(v string) *SetWorkItemPinnedResponseBodyAttachments {
	s.RefUrl = &v
	return s
}

func (s *SetWorkItemPinnedResponseBodyAttachments) SetResourceId(v string) *SetWorkItemPinnedResponseBodyAttachments {
	s.ResourceId = &v
	return s
}

func (s *SetWorkItemPinnedResponseBodyAttachments) SetResourceKind(v string) *SetWorkItemPinnedResponseBodyAttachments {
	s.ResourceKind = &v
	return s
}

func (s *SetWorkItemPinnedResponseBodyAttachments) SetSizeBytes(v int64) *SetWorkItemPinnedResponseBodyAttachments {
	s.SizeBytes = &v
	return s
}

func (s *SetWorkItemPinnedResponseBodyAttachments) SetTitle(v string) *SetWorkItemPinnedResponseBodyAttachments {
	s.Title = &v
	return s
}

func (s *SetWorkItemPinnedResponseBodyAttachments) Validate() error {
	return dara.Validate(s)
}
