// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptanceCriteria(v string) *UpdateWorkItemResponseBody
	GetAcceptanceCriteria() *string
	SetAppliedStatusVersion(v int64) *UpdateWorkItemResponseBody
	GetAppliedStatusVersion() *int64
	SetAttachments(v []*UpdateWorkItemResponseBodyAttachments) *UpdateWorkItemResponseBody
	GetAttachments() []*UpdateWorkItemResponseBodyAttachments
	SetBusinessStatus(v string) *UpdateWorkItemResponseBody
	GetBusinessStatus() *string
	SetCanDelete(v bool) *UpdateWorkItemResponseBody
	GetCanDelete() *bool
	SetCanManage(v bool) *UpdateWorkItemResponseBody
	GetCanManage() *bool
	SetCanManageVisibility(v bool) *UpdateWorkItemResponseBody
	GetCanManageVisibility() *bool
	SetCanRename(v bool) *UpdateWorkItemResponseBody
	GetCanRename() *bool
	SetCode(v string) *UpdateWorkItemResponseBody
	GetCode() *string
	SetCollaborationGroupId(v string) *UpdateWorkItemResponseBody
	GetCollaborationGroupId() *string
	SetCollaboratorUserIds(v []*int64) *UpdateWorkItemResponseBody
	GetCollaboratorUserIds() []*int64
	SetConversationId(v string) *UpdateWorkItemResponseBody
	GetConversationId() *string
	SetCreatedAt(v string) *UpdateWorkItemResponseBody
	GetCreatedAt() *string
	SetCreatedByUserId(v int64) *UpdateWorkItemResponseBody
	GetCreatedByUserId() *int64
	SetCreationType(v string) *UpdateWorkItemResponseBody
	GetCreationType() *string
	SetCurrentPlanId(v string) *UpdateWorkItemResponseBody
	GetCurrentPlanId() *string
	SetCurrentRevision(v int64) *UpdateWorkItemResponseBody
	GetCurrentRevision() *int64
	SetCurrentRunId(v string) *UpdateWorkItemResponseBody
	GetCurrentRunId() *string
	SetCurrentRunSeq(v int64) *UpdateWorkItemResponseBody
	GetCurrentRunSeq() *int64
	SetCurrentStepName(v string) *UpdateWorkItemResponseBody
	GetCurrentStepName() *string
	SetDescription(v string) *UpdateWorkItemResponseBody
	GetDescription() *string
	SetInputVersion(v int64) *UpdateWorkItemResponseBody
	GetInputVersion() *int64
	SetItemType(v string) *UpdateWorkItemResponseBody
	GetItemType() *string
	SetLastProcessingAt(v string) *UpdateWorkItemResponseBody
	GetLastProcessingAt() *string
	SetLegacyTopic(v bool) *UpdateWorkItemResponseBody
	GetLegacyTopic() *bool
	SetMemberUserIds(v []*int64) *UpdateWorkItemResponseBody
	GetMemberUserIds() []*int64
	SetMessage(v string) *UpdateWorkItemResponseBody
	GetMessage() *string
	SetModifiedAt(v string) *UpdateWorkItemResponseBody
	GetModifiedAt() *string
	SetOwnerUserId(v int64) *UpdateWorkItemResponseBody
	GetOwnerUserId() *int64
	SetPinned(v bool) *UpdateWorkItemResponseBody
	GetPinned() *bool
	SetPriority(v string) *UpdateWorkItemResponseBody
	GetPriority() *string
	SetProgressCompleted(v int64) *UpdateWorkItemResponseBody
	GetProgressCompleted() *int64
	SetProgressTotal(v int64) *UpdateWorkItemResponseBody
	GetProgressTotal() *int64
	SetRequestId(v string) *UpdateWorkItemResponseBody
	GetRequestId() *string
	SetSourceExecutionId(v string) *UpdateWorkItemResponseBody
	GetSourceExecutionId() *string
	SetSourceTaskId(v string) *UpdateWorkItemResponseBody
	GetSourceTaskId() *string
	SetSourceTaskName(v string) *UpdateWorkItemResponseBody
	GetSourceTaskName() *string
	SetSourceTriggerKey(v string) *UpdateWorkItemResponseBody
	GetSourceTriggerKey() *string
	SetSourceType(v string) *UpdateWorkItemResponseBody
	GetSourceType() *string
	SetTitle(v string) *UpdateWorkItemResponseBody
	GetTitle() *string
	SetTriggerTimezone(v string) *UpdateWorkItemResponseBody
	GetTriggerTimezone() *string
	SetTriggeredAt(v string) *UpdateWorkItemResponseBody
	GetTriggeredAt() *string
	SetVersion(v int64) *UpdateWorkItemResponseBody
	GetVersion() *int64
	SetVisibility(v string) *UpdateWorkItemResponseBody
	GetVisibility() *string
	SetVisibilitySource(v string) *UpdateWorkItemResponseBody
	GetVisibilitySource() *string
	SetWorkItemId(v string) *UpdateWorkItemResponseBody
	GetWorkItemId() *string
	SetWorkflowStatus(v string) *UpdateWorkItemResponseBody
	GetWorkflowStatus() *string
	SetWorkflowTemplateIds(v []*string) *UpdateWorkItemResponseBody
	GetWorkflowTemplateIds() []*string
}

type UpdateWorkItemResponseBody struct {
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
	Attachments []*UpdateWorkItemResponseBodyAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// The business status (seven states).
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
	// The creation time (ISO 8601).
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
	// The time when the item last entered the processing state (ISO 8601).
	//
	// example:
	//
	// example
	LastProcessingAt *string `json:"lastProcessingAt,omitempty" xml:"lastProcessingAt,omitempty"`
	// Indicates whether it is a free exploration topic during the compatibility period.
	//
	// example:
	//
	// false
	LegacyTopic *bool `json:"legacyTopic,omitempty" xml:"legacyTopic,omitempty"`
	// The list of user IDs of visible members.
	//
	// example:
	//
	// example
	MemberUserIds []*int64 `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The modification time (ISO 8601).
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
	// The task execution ID.
	//
	// example:
	//
	// example
	SourceExecutionId *string `json:"sourceExecutionId,omitempty" xml:"sourceExecutionId,omitempty"`
	// The task source ID.
	//
	// example:
	//
	// example
	SourceTaskId *string `json:"sourceTaskId,omitempty" xml:"sourceTaskId,omitempty"`
	// The task source name.
	//
	// example:
	//
	// example
	SourceTaskName *string `json:"sourceTaskName,omitempty" xml:"sourceTaskName,omitempty"`
	// The task trigger key.
	//
	// example:
	//
	// example
	SourceTriggerKey *string `json:"sourceTriggerKey,omitempty" xml:"sourceTriggerKey,omitempty"`
	// The task source type.
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
	// The task trigger time zone.
	//
	// example:
	//
	// example
	TriggerTimezone *string `json:"triggerTimezone,omitempty" xml:"triggerTimezone,omitempty"`
	// The task trigger time (ISO 8601).
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
	// The workflow status projection (four states).
	//
	// example:
	//
	// example
	WorkflowStatus *string `json:"workflowStatus,omitempty" xml:"workflowStatus,omitempty"`
	// The list of IDs for the bound workflow templates.
	//
	// example:
	//
	// example
	WorkflowTemplateIds []*string `json:"workflowTemplateIds,omitempty" xml:"workflowTemplateIds,omitempty" type:"Repeated"`
}

func (s UpdateWorkItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkItemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkItemResponseBody) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *UpdateWorkItemResponseBody) GetAppliedStatusVersion() *int64 {
	return s.AppliedStatusVersion
}

func (s *UpdateWorkItemResponseBody) GetAttachments() []*UpdateWorkItemResponseBodyAttachments {
	return s.Attachments
}

func (s *UpdateWorkItemResponseBody) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *UpdateWorkItemResponseBody) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *UpdateWorkItemResponseBody) GetCanManage() *bool {
	return s.CanManage
}

func (s *UpdateWorkItemResponseBody) GetCanManageVisibility() *bool {
	return s.CanManageVisibility
}

func (s *UpdateWorkItemResponseBody) GetCanRename() *bool {
	return s.CanRename
}

func (s *UpdateWorkItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateWorkItemResponseBody) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *UpdateWorkItemResponseBody) GetCollaboratorUserIds() []*int64 {
	return s.CollaboratorUserIds
}

func (s *UpdateWorkItemResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *UpdateWorkItemResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateWorkItemResponseBody) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *UpdateWorkItemResponseBody) GetCreationType() *string {
	return s.CreationType
}

func (s *UpdateWorkItemResponseBody) GetCurrentPlanId() *string {
	return s.CurrentPlanId
}

func (s *UpdateWorkItemResponseBody) GetCurrentRevision() *int64 {
	return s.CurrentRevision
}

func (s *UpdateWorkItemResponseBody) GetCurrentRunId() *string {
	return s.CurrentRunId
}

func (s *UpdateWorkItemResponseBody) GetCurrentRunSeq() *int64 {
	return s.CurrentRunSeq
}

func (s *UpdateWorkItemResponseBody) GetCurrentStepName() *string {
	return s.CurrentStepName
}

func (s *UpdateWorkItemResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkItemResponseBody) GetInputVersion() *int64 {
	return s.InputVersion
}

func (s *UpdateWorkItemResponseBody) GetItemType() *string {
	return s.ItemType
}

func (s *UpdateWorkItemResponseBody) GetLastProcessingAt() *string {
	return s.LastProcessingAt
}

func (s *UpdateWorkItemResponseBody) GetLegacyTopic() *bool {
	return s.LegacyTopic
}

func (s *UpdateWorkItemResponseBody) GetMemberUserIds() []*int64 {
	return s.MemberUserIds
}

func (s *UpdateWorkItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWorkItemResponseBody) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *UpdateWorkItemResponseBody) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *UpdateWorkItemResponseBody) GetPinned() *bool {
	return s.Pinned
}

func (s *UpdateWorkItemResponseBody) GetPriority() *string {
	return s.Priority
}

func (s *UpdateWorkItemResponseBody) GetProgressCompleted() *int64 {
	return s.ProgressCompleted
}

func (s *UpdateWorkItemResponseBody) GetProgressTotal() *int64 {
	return s.ProgressTotal
}

func (s *UpdateWorkItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkItemResponseBody) GetSourceExecutionId() *string {
	return s.SourceExecutionId
}

func (s *UpdateWorkItemResponseBody) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *UpdateWorkItemResponseBody) GetSourceTaskName() *string {
	return s.SourceTaskName
}

func (s *UpdateWorkItemResponseBody) GetSourceTriggerKey() *string {
	return s.SourceTriggerKey
}

func (s *UpdateWorkItemResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateWorkItemResponseBody) GetTitle() *string {
	return s.Title
}

func (s *UpdateWorkItemResponseBody) GetTriggerTimezone() *string {
	return s.TriggerTimezone
}

func (s *UpdateWorkItemResponseBody) GetTriggeredAt() *string {
	return s.TriggeredAt
}

func (s *UpdateWorkItemResponseBody) GetVersion() *int64 {
	return s.Version
}

func (s *UpdateWorkItemResponseBody) GetVisibility() *string {
	return s.Visibility
}

func (s *UpdateWorkItemResponseBody) GetVisibilitySource() *string {
	return s.VisibilitySource
}

func (s *UpdateWorkItemResponseBody) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *UpdateWorkItemResponseBody) GetWorkflowStatus() *string {
	return s.WorkflowStatus
}

func (s *UpdateWorkItemResponseBody) GetWorkflowTemplateIds() []*string {
	return s.WorkflowTemplateIds
}

func (s *UpdateWorkItemResponseBody) SetAcceptanceCriteria(v string) *UpdateWorkItemResponseBody {
	s.AcceptanceCriteria = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetAppliedStatusVersion(v int64) *UpdateWorkItemResponseBody {
	s.AppliedStatusVersion = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetAttachments(v []*UpdateWorkItemResponseBodyAttachments) *UpdateWorkItemResponseBody {
	s.Attachments = v
	return s
}

func (s *UpdateWorkItemResponseBody) SetBusinessStatus(v string) *UpdateWorkItemResponseBody {
	s.BusinessStatus = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCanDelete(v bool) *UpdateWorkItemResponseBody {
	s.CanDelete = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCanManage(v bool) *UpdateWorkItemResponseBody {
	s.CanManage = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCanManageVisibility(v bool) *UpdateWorkItemResponseBody {
	s.CanManageVisibility = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCanRename(v bool) *UpdateWorkItemResponseBody {
	s.CanRename = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCode(v string) *UpdateWorkItemResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCollaborationGroupId(v string) *UpdateWorkItemResponseBody {
	s.CollaborationGroupId = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCollaboratorUserIds(v []*int64) *UpdateWorkItemResponseBody {
	s.CollaboratorUserIds = v
	return s
}

func (s *UpdateWorkItemResponseBody) SetConversationId(v string) *UpdateWorkItemResponseBody {
	s.ConversationId = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCreatedAt(v string) *UpdateWorkItemResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCreatedByUserId(v int64) *UpdateWorkItemResponseBody {
	s.CreatedByUserId = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCreationType(v string) *UpdateWorkItemResponseBody {
	s.CreationType = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCurrentPlanId(v string) *UpdateWorkItemResponseBody {
	s.CurrentPlanId = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCurrentRevision(v int64) *UpdateWorkItemResponseBody {
	s.CurrentRevision = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCurrentRunId(v string) *UpdateWorkItemResponseBody {
	s.CurrentRunId = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCurrentRunSeq(v int64) *UpdateWorkItemResponseBody {
	s.CurrentRunSeq = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetCurrentStepName(v string) *UpdateWorkItemResponseBody {
	s.CurrentStepName = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetDescription(v string) *UpdateWorkItemResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetInputVersion(v int64) *UpdateWorkItemResponseBody {
	s.InputVersion = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetItemType(v string) *UpdateWorkItemResponseBody {
	s.ItemType = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetLastProcessingAt(v string) *UpdateWorkItemResponseBody {
	s.LastProcessingAt = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetLegacyTopic(v bool) *UpdateWorkItemResponseBody {
	s.LegacyTopic = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetMemberUserIds(v []*int64) *UpdateWorkItemResponseBody {
	s.MemberUserIds = v
	return s
}

func (s *UpdateWorkItemResponseBody) SetMessage(v string) *UpdateWorkItemResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetModifiedAt(v string) *UpdateWorkItemResponseBody {
	s.ModifiedAt = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetOwnerUserId(v int64) *UpdateWorkItemResponseBody {
	s.OwnerUserId = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetPinned(v bool) *UpdateWorkItemResponseBody {
	s.Pinned = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetPriority(v string) *UpdateWorkItemResponseBody {
	s.Priority = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetProgressCompleted(v int64) *UpdateWorkItemResponseBody {
	s.ProgressCompleted = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetProgressTotal(v int64) *UpdateWorkItemResponseBody {
	s.ProgressTotal = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetRequestId(v string) *UpdateWorkItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetSourceExecutionId(v string) *UpdateWorkItemResponseBody {
	s.SourceExecutionId = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetSourceTaskId(v string) *UpdateWorkItemResponseBody {
	s.SourceTaskId = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetSourceTaskName(v string) *UpdateWorkItemResponseBody {
	s.SourceTaskName = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetSourceTriggerKey(v string) *UpdateWorkItemResponseBody {
	s.SourceTriggerKey = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetSourceType(v string) *UpdateWorkItemResponseBody {
	s.SourceType = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetTitle(v string) *UpdateWorkItemResponseBody {
	s.Title = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetTriggerTimezone(v string) *UpdateWorkItemResponseBody {
	s.TriggerTimezone = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetTriggeredAt(v string) *UpdateWorkItemResponseBody {
	s.TriggeredAt = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetVersion(v int64) *UpdateWorkItemResponseBody {
	s.Version = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetVisibility(v string) *UpdateWorkItemResponseBody {
	s.Visibility = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetVisibilitySource(v string) *UpdateWorkItemResponseBody {
	s.VisibilitySource = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetWorkItemId(v string) *UpdateWorkItemResponseBody {
	s.WorkItemId = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetWorkflowStatus(v string) *UpdateWorkItemResponseBody {
	s.WorkflowStatus = &v
	return s
}

func (s *UpdateWorkItemResponseBody) SetWorkflowTemplateIds(v []*string) *UpdateWorkItemResponseBody {
	s.WorkflowTemplateIds = v
	return s
}

func (s *UpdateWorkItemResponseBody) Validate() error {
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

type UpdateWorkItemResponseBodyAttachments struct {
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
	// The access URL of the attachment (prefixed with an OpenAPI 302 redirect).
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

func (s UpdateWorkItemResponseBodyAttachments) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkItemResponseBodyAttachments) GoString() string {
	return s.String()
}

func (s *UpdateWorkItemResponseBodyAttachments) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *UpdateWorkItemResponseBodyAttachments) GetContentType() *string {
	return s.ContentType
}

func (s *UpdateWorkItemResponseBodyAttachments) GetRefUrl() *string {
	return s.RefUrl
}

func (s *UpdateWorkItemResponseBodyAttachments) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateWorkItemResponseBodyAttachments) GetResourceKind() *string {
	return s.ResourceKind
}

func (s *UpdateWorkItemResponseBodyAttachments) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *UpdateWorkItemResponseBodyAttachments) GetTitle() *string {
	return s.Title
}

func (s *UpdateWorkItemResponseBodyAttachments) SetAttachmentId(v string) *UpdateWorkItemResponseBodyAttachments {
	s.AttachmentId = &v
	return s
}

func (s *UpdateWorkItemResponseBodyAttachments) SetContentType(v string) *UpdateWorkItemResponseBodyAttachments {
	s.ContentType = &v
	return s
}

func (s *UpdateWorkItemResponseBodyAttachments) SetRefUrl(v string) *UpdateWorkItemResponseBodyAttachments {
	s.RefUrl = &v
	return s
}

func (s *UpdateWorkItemResponseBodyAttachments) SetResourceId(v string) *UpdateWorkItemResponseBodyAttachments {
	s.ResourceId = &v
	return s
}

func (s *UpdateWorkItemResponseBodyAttachments) SetResourceKind(v string) *UpdateWorkItemResponseBodyAttachments {
	s.ResourceKind = &v
	return s
}

func (s *UpdateWorkItemResponseBodyAttachments) SetSizeBytes(v int64) *UpdateWorkItemResponseBodyAttachments {
	s.SizeBytes = &v
	return s
}

func (s *UpdateWorkItemResponseBodyAttachments) SetTitle(v string) *UpdateWorkItemResponseBodyAttachments {
	s.Title = &v
	return s
}

func (s *UpdateWorkItemResponseBodyAttachments) Validate() error {
	return dara.Validate(s)
}
