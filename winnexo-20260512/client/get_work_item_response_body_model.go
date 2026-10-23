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
	// The list of input attachments.
	//
	// example:
	//
	// example
	Attachments []*GetWorkItemResponseBodyAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// The business status (seven states).
	//
	// example:
	//
	// PROCESSING
	BusinessStatus *string `json:"businessStatus,omitempty" xml:"businessStatus,omitempty"`
	// Indicates whether the caller can delete the work item.
	//
	// example:
	//
	// false
	CanDelete *bool `json:"canDelete,omitempty" xml:"canDelete,omitempty"`
	// Indicates whether the caller can manage the work item.
	//
	// example:
	//
	// false
	CanManage *bool `json:"canManage,omitempty" xml:"canManage,omitempty"`
	// Indicates whether the visibility scope can be adjusted independently.
	//
	// example:
	//
	// false
	CanManageVisibility *bool `json:"canManageVisibility,omitempty" xml:"canManageVisibility,omitempty"`
	// Indicates whether the caller can rename the work item.
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
	// The ID of the collaboration space.
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
	// The ID of the associated session.
	//
	// example:
	//
	// example
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// The creation time, in ISO 8601 format.
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
	// The creation source. Valid values: DEFAULT, TASK.
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
	// The time when the work item last entered the processing state, in ISO 8601 format.
	//
	// example:
	//
	// example
	LastProcessingAt *string `json:"lastProcessingAt,omitempty" xml:"lastProcessingAt,omitempty"`
	// Indicates whether it is a legacy free exploration topic.
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
	// The modification time, in ISO 8601 format.
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
	// The personal pin status for the caller.
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
	// The trigger time of the task, in ISO 8601 format.
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
	// The source of the visibility scope. Valid values: SELF, TASK, TEMPLATE.
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
	// The access URL for the attachment (prefixed with an OpenAPI 302 redirect).
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
