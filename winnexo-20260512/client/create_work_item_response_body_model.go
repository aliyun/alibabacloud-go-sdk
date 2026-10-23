// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptanceCriteria(v string) *CreateWorkItemResponseBody
	GetAcceptanceCriteria() *string
	SetAppliedStatusVersion(v int64) *CreateWorkItemResponseBody
	GetAppliedStatusVersion() *int64
	SetAttachments(v []*CreateWorkItemResponseBodyAttachments) *CreateWorkItemResponseBody
	GetAttachments() []*CreateWorkItemResponseBodyAttachments
	SetBusinessStatus(v string) *CreateWorkItemResponseBody
	GetBusinessStatus() *string
	SetCanDelete(v bool) *CreateWorkItemResponseBody
	GetCanDelete() *bool
	SetCanManage(v bool) *CreateWorkItemResponseBody
	GetCanManage() *bool
	SetCanManageVisibility(v bool) *CreateWorkItemResponseBody
	GetCanManageVisibility() *bool
	SetCanRename(v bool) *CreateWorkItemResponseBody
	GetCanRename() *bool
	SetCode(v string) *CreateWorkItemResponseBody
	GetCode() *string
	SetCollaborationGroupId(v string) *CreateWorkItemResponseBody
	GetCollaborationGroupId() *string
	SetCollaboratorUserIds(v []*int64) *CreateWorkItemResponseBody
	GetCollaboratorUserIds() []*int64
	SetConversationId(v string) *CreateWorkItemResponseBody
	GetConversationId() *string
	SetCreatedAt(v string) *CreateWorkItemResponseBody
	GetCreatedAt() *string
	SetCreatedByUserId(v int64) *CreateWorkItemResponseBody
	GetCreatedByUserId() *int64
	SetCreationType(v string) *CreateWorkItemResponseBody
	GetCreationType() *string
	SetCurrentPlanId(v string) *CreateWorkItemResponseBody
	GetCurrentPlanId() *string
	SetCurrentRevision(v int64) *CreateWorkItemResponseBody
	GetCurrentRevision() *int64
	SetCurrentRunId(v string) *CreateWorkItemResponseBody
	GetCurrentRunId() *string
	SetCurrentRunSeq(v int64) *CreateWorkItemResponseBody
	GetCurrentRunSeq() *int64
	SetCurrentStepName(v string) *CreateWorkItemResponseBody
	GetCurrentStepName() *string
	SetDescription(v string) *CreateWorkItemResponseBody
	GetDescription() *string
	SetInputVersion(v int64) *CreateWorkItemResponseBody
	GetInputVersion() *int64
	SetItemType(v string) *CreateWorkItemResponseBody
	GetItemType() *string
	SetLastProcessingAt(v string) *CreateWorkItemResponseBody
	GetLastProcessingAt() *string
	SetLegacyTopic(v bool) *CreateWorkItemResponseBody
	GetLegacyTopic() *bool
	SetMemberUserIds(v []*int64) *CreateWorkItemResponseBody
	GetMemberUserIds() []*int64
	SetMessage(v string) *CreateWorkItemResponseBody
	GetMessage() *string
	SetModifiedAt(v string) *CreateWorkItemResponseBody
	GetModifiedAt() *string
	SetOwnerUserId(v int64) *CreateWorkItemResponseBody
	GetOwnerUserId() *int64
	SetPinned(v bool) *CreateWorkItemResponseBody
	GetPinned() *bool
	SetPriority(v string) *CreateWorkItemResponseBody
	GetPriority() *string
	SetProgressCompleted(v int64) *CreateWorkItemResponseBody
	GetProgressCompleted() *int64
	SetProgressTotal(v int64) *CreateWorkItemResponseBody
	GetProgressTotal() *int64
	SetRequestId(v string) *CreateWorkItemResponseBody
	GetRequestId() *string
	SetSourceExecutionId(v string) *CreateWorkItemResponseBody
	GetSourceExecutionId() *string
	SetSourceTaskId(v string) *CreateWorkItemResponseBody
	GetSourceTaskId() *string
	SetSourceTaskName(v string) *CreateWorkItemResponseBody
	GetSourceTaskName() *string
	SetSourceTriggerKey(v string) *CreateWorkItemResponseBody
	GetSourceTriggerKey() *string
	SetSourceType(v string) *CreateWorkItemResponseBody
	GetSourceType() *string
	SetTitle(v string) *CreateWorkItemResponseBody
	GetTitle() *string
	SetTriggerTimezone(v string) *CreateWorkItemResponseBody
	GetTriggerTimezone() *string
	SetTriggeredAt(v string) *CreateWorkItemResponseBody
	GetTriggeredAt() *string
	SetVersion(v int64) *CreateWorkItemResponseBody
	GetVersion() *int64
	SetVisibility(v string) *CreateWorkItemResponseBody
	GetVisibility() *string
	SetVisibilitySource(v string) *CreateWorkItemResponseBody
	GetVisibilitySource() *string
	SetWorkItemId(v string) *CreateWorkItemResponseBody
	GetWorkItemId() *string
	SetWorkflowStatus(v string) *CreateWorkItemResponseBody
	GetWorkflowStatus() *string
	SetWorkflowTemplateIds(v []*string) *CreateWorkItemResponseBody
	GetWorkflowTemplateIds() []*string
}

type CreateWorkItemResponseBody struct {
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
	Attachments []*CreateWorkItemResponseBodyAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// The business status.
	//
	// example:
	//
	// PROCESSING
	BusinessStatus *string `json:"businessStatus,omitempty" xml:"businessStatus,omitempty"`
	// Indicates whether the caller can delete the item.
	//
	// example:
	//
	// false
	CanDelete *bool `json:"canDelete,omitempty" xml:"canDelete,omitempty"`
	// Indicates whether the caller can manage the item.
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
	// Indicates whether the caller can rename the item.
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
	// The collaboration workspace ID.
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
	// The creation source. Valid values: DEFAULT and TASK.
	//
	// example:
	//
	// DEFAULT
	CreationType *string `json:"creationType,omitempty" xml:"creationType,omitempty"`
	// The current physical plan ID.
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
	// The current run ID.
	//
	// example:
	//
	// example
	CurrentRunId *string `json:"currentRunId,omitempty" xml:"currentRunId,omitempty"`
	// The current run sequence number.
	//
	// example:
	//
	// 1
	CurrentRunSeq *int64 `json:"currentRunSeq,omitempty" xml:"currentRunSeq,omitempty"`
	// The current step name.
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
	// The work item type.
	//
	// example:
	//
	// WORKFLOW
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// The time when the item last entered the processing state, in ISO 8601 format.
	//
	// example:
	//
	// example
	LastProcessingAt *string `json:"lastProcessingAt,omitempty" xml:"lastProcessingAt,omitempty"`
	// Indicates whether the topic is a free exploration topic during the compatibility period.
	//
	// example:
	//
	// false
	LegacyTopic *bool `json:"legacyTopic,omitempty" xml:"legacyTopic,omitempty"`
	// The list of user IDs for visible members.
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
	// Indicates whether the item is pinned by the caller.
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
	// The request ID.
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
	// The source of the visibility scope. Valid values: SELF, TASK, and TEMPLATE.
	//
	// example:
	//
	// SELF
	VisibilitySource *string `json:"visibilitySource,omitempty" xml:"visibilitySource,omitempty"`
	// The work item ID.
	//
	// example:
	//
	// example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
	// The four-state workflow status projection.
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

func (s CreateWorkItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkItemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkItemResponseBody) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *CreateWorkItemResponseBody) GetAppliedStatusVersion() *int64 {
	return s.AppliedStatusVersion
}

func (s *CreateWorkItemResponseBody) GetAttachments() []*CreateWorkItemResponseBodyAttachments {
	return s.Attachments
}

func (s *CreateWorkItemResponseBody) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *CreateWorkItemResponseBody) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *CreateWorkItemResponseBody) GetCanManage() *bool {
	return s.CanManage
}

func (s *CreateWorkItemResponseBody) GetCanManageVisibility() *bool {
	return s.CanManageVisibility
}

func (s *CreateWorkItemResponseBody) GetCanRename() *bool {
	return s.CanRename
}

func (s *CreateWorkItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWorkItemResponseBody) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateWorkItemResponseBody) GetCollaboratorUserIds() []*int64 {
	return s.CollaboratorUserIds
}

func (s *CreateWorkItemResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *CreateWorkItemResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateWorkItemResponseBody) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *CreateWorkItemResponseBody) GetCreationType() *string {
	return s.CreationType
}

func (s *CreateWorkItemResponseBody) GetCurrentPlanId() *string {
	return s.CurrentPlanId
}

func (s *CreateWorkItemResponseBody) GetCurrentRevision() *int64 {
	return s.CurrentRevision
}

func (s *CreateWorkItemResponseBody) GetCurrentRunId() *string {
	return s.CurrentRunId
}

func (s *CreateWorkItemResponseBody) GetCurrentRunSeq() *int64 {
	return s.CurrentRunSeq
}

func (s *CreateWorkItemResponseBody) GetCurrentStepName() *string {
	return s.CurrentStepName
}

func (s *CreateWorkItemResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkItemResponseBody) GetInputVersion() *int64 {
	return s.InputVersion
}

func (s *CreateWorkItemResponseBody) GetItemType() *string {
	return s.ItemType
}

func (s *CreateWorkItemResponseBody) GetLastProcessingAt() *string {
	return s.LastProcessingAt
}

func (s *CreateWorkItemResponseBody) GetLegacyTopic() *bool {
	return s.LegacyTopic
}

func (s *CreateWorkItemResponseBody) GetMemberUserIds() []*int64 {
	return s.MemberUserIds
}

func (s *CreateWorkItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWorkItemResponseBody) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *CreateWorkItemResponseBody) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *CreateWorkItemResponseBody) GetPinned() *bool {
	return s.Pinned
}

func (s *CreateWorkItemResponseBody) GetPriority() *string {
	return s.Priority
}

func (s *CreateWorkItemResponseBody) GetProgressCompleted() *int64 {
	return s.ProgressCompleted
}

func (s *CreateWorkItemResponseBody) GetProgressTotal() *int64 {
	return s.ProgressTotal
}

func (s *CreateWorkItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkItemResponseBody) GetSourceExecutionId() *string {
	return s.SourceExecutionId
}

func (s *CreateWorkItemResponseBody) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *CreateWorkItemResponseBody) GetSourceTaskName() *string {
	return s.SourceTaskName
}

func (s *CreateWorkItemResponseBody) GetSourceTriggerKey() *string {
	return s.SourceTriggerKey
}

func (s *CreateWorkItemResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateWorkItemResponseBody) GetTitle() *string {
	return s.Title
}

func (s *CreateWorkItemResponseBody) GetTriggerTimezone() *string {
	return s.TriggerTimezone
}

func (s *CreateWorkItemResponseBody) GetTriggeredAt() *string {
	return s.TriggeredAt
}

func (s *CreateWorkItemResponseBody) GetVersion() *int64 {
	return s.Version
}

func (s *CreateWorkItemResponseBody) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateWorkItemResponseBody) GetVisibilitySource() *string {
	return s.VisibilitySource
}

func (s *CreateWorkItemResponseBody) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *CreateWorkItemResponseBody) GetWorkflowStatus() *string {
	return s.WorkflowStatus
}

func (s *CreateWorkItemResponseBody) GetWorkflowTemplateIds() []*string {
	return s.WorkflowTemplateIds
}

func (s *CreateWorkItemResponseBody) SetAcceptanceCriteria(v string) *CreateWorkItemResponseBody {
	s.AcceptanceCriteria = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetAppliedStatusVersion(v int64) *CreateWorkItemResponseBody {
	s.AppliedStatusVersion = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetAttachments(v []*CreateWorkItemResponseBodyAttachments) *CreateWorkItemResponseBody {
	s.Attachments = v
	return s
}

func (s *CreateWorkItemResponseBody) SetBusinessStatus(v string) *CreateWorkItemResponseBody {
	s.BusinessStatus = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCanDelete(v bool) *CreateWorkItemResponseBody {
	s.CanDelete = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCanManage(v bool) *CreateWorkItemResponseBody {
	s.CanManage = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCanManageVisibility(v bool) *CreateWorkItemResponseBody {
	s.CanManageVisibility = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCanRename(v bool) *CreateWorkItemResponseBody {
	s.CanRename = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCode(v string) *CreateWorkItemResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCollaborationGroupId(v string) *CreateWorkItemResponseBody {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCollaboratorUserIds(v []*int64) *CreateWorkItemResponseBody {
	s.CollaboratorUserIds = v
	return s
}

func (s *CreateWorkItemResponseBody) SetConversationId(v string) *CreateWorkItemResponseBody {
	s.ConversationId = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCreatedAt(v string) *CreateWorkItemResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCreatedByUserId(v int64) *CreateWorkItemResponseBody {
	s.CreatedByUserId = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCreationType(v string) *CreateWorkItemResponseBody {
	s.CreationType = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCurrentPlanId(v string) *CreateWorkItemResponseBody {
	s.CurrentPlanId = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCurrentRevision(v int64) *CreateWorkItemResponseBody {
	s.CurrentRevision = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCurrentRunId(v string) *CreateWorkItemResponseBody {
	s.CurrentRunId = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCurrentRunSeq(v int64) *CreateWorkItemResponseBody {
	s.CurrentRunSeq = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetCurrentStepName(v string) *CreateWorkItemResponseBody {
	s.CurrentStepName = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetDescription(v string) *CreateWorkItemResponseBody {
	s.Description = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetInputVersion(v int64) *CreateWorkItemResponseBody {
	s.InputVersion = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetItemType(v string) *CreateWorkItemResponseBody {
	s.ItemType = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetLastProcessingAt(v string) *CreateWorkItemResponseBody {
	s.LastProcessingAt = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetLegacyTopic(v bool) *CreateWorkItemResponseBody {
	s.LegacyTopic = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetMemberUserIds(v []*int64) *CreateWorkItemResponseBody {
	s.MemberUserIds = v
	return s
}

func (s *CreateWorkItemResponseBody) SetMessage(v string) *CreateWorkItemResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetModifiedAt(v string) *CreateWorkItemResponseBody {
	s.ModifiedAt = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetOwnerUserId(v int64) *CreateWorkItemResponseBody {
	s.OwnerUserId = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetPinned(v bool) *CreateWorkItemResponseBody {
	s.Pinned = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetPriority(v string) *CreateWorkItemResponseBody {
	s.Priority = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetProgressCompleted(v int64) *CreateWorkItemResponseBody {
	s.ProgressCompleted = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetProgressTotal(v int64) *CreateWorkItemResponseBody {
	s.ProgressTotal = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetRequestId(v string) *CreateWorkItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetSourceExecutionId(v string) *CreateWorkItemResponseBody {
	s.SourceExecutionId = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetSourceTaskId(v string) *CreateWorkItemResponseBody {
	s.SourceTaskId = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetSourceTaskName(v string) *CreateWorkItemResponseBody {
	s.SourceTaskName = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetSourceTriggerKey(v string) *CreateWorkItemResponseBody {
	s.SourceTriggerKey = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetSourceType(v string) *CreateWorkItemResponseBody {
	s.SourceType = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetTitle(v string) *CreateWorkItemResponseBody {
	s.Title = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetTriggerTimezone(v string) *CreateWorkItemResponseBody {
	s.TriggerTimezone = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetTriggeredAt(v string) *CreateWorkItemResponseBody {
	s.TriggeredAt = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetVersion(v int64) *CreateWorkItemResponseBody {
	s.Version = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetVisibility(v string) *CreateWorkItemResponseBody {
	s.Visibility = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetVisibilitySource(v string) *CreateWorkItemResponseBody {
	s.VisibilitySource = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetWorkItemId(v string) *CreateWorkItemResponseBody {
	s.WorkItemId = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetWorkflowStatus(v string) *CreateWorkItemResponseBody {
	s.WorkflowStatus = &v
	return s
}

func (s *CreateWorkItemResponseBody) SetWorkflowTemplateIds(v []*string) *CreateWorkItemResponseBody {
	s.WorkflowTemplateIds = v
	return s
}

func (s *CreateWorkItemResponseBody) Validate() error {
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

type CreateWorkItemResponseBodyAttachments struct {
	// The attachment ID.
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
	// The access URL of the attachment, prefixed with an OpenAPI 302 redirect.
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

func (s CreateWorkItemResponseBodyAttachments) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkItemResponseBodyAttachments) GoString() string {
	return s.String()
}

func (s *CreateWorkItemResponseBodyAttachments) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *CreateWorkItemResponseBodyAttachments) GetContentType() *string {
	return s.ContentType
}

func (s *CreateWorkItemResponseBodyAttachments) GetRefUrl() *string {
	return s.RefUrl
}

func (s *CreateWorkItemResponseBodyAttachments) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateWorkItemResponseBodyAttachments) GetResourceKind() *string {
	return s.ResourceKind
}

func (s *CreateWorkItemResponseBodyAttachments) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *CreateWorkItemResponseBodyAttachments) GetTitle() *string {
	return s.Title
}

func (s *CreateWorkItemResponseBodyAttachments) SetAttachmentId(v string) *CreateWorkItemResponseBodyAttachments {
	s.AttachmentId = &v
	return s
}

func (s *CreateWorkItemResponseBodyAttachments) SetContentType(v string) *CreateWorkItemResponseBodyAttachments {
	s.ContentType = &v
	return s
}

func (s *CreateWorkItemResponseBodyAttachments) SetRefUrl(v string) *CreateWorkItemResponseBodyAttachments {
	s.RefUrl = &v
	return s
}

func (s *CreateWorkItemResponseBodyAttachments) SetResourceId(v string) *CreateWorkItemResponseBodyAttachments {
	s.ResourceId = &v
	return s
}

func (s *CreateWorkItemResponseBodyAttachments) SetResourceKind(v string) *CreateWorkItemResponseBodyAttachments {
	s.ResourceKind = &v
	return s
}

func (s *CreateWorkItemResponseBodyAttachments) SetSizeBytes(v int64) *CreateWorkItemResponseBodyAttachments {
	s.SizeBytes = &v
	return s
}

func (s *CreateWorkItemResponseBodyAttachments) SetTitle(v string) *CreateWorkItemResponseBodyAttachments {
	s.Title = &v
	return s
}

func (s *CreateWorkItemResponseBodyAttachments) Validate() error {
	return dara.Validate(s)
}
