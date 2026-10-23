// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWorkItemsResponseBody
	GetCode() *string
	SetHasMore(v bool) *ListWorkItemsResponseBody
	GetHasMore() *bool
	SetItems(v []*ListWorkItemsResponseBodyItems) *ListWorkItemsResponseBody
	GetItems() []*ListWorkItemsResponseBodyItems
	SetMaxResults(v int64) *ListWorkItemsResponseBody
	GetMaxResults() *int64
	SetMessage(v string) *ListWorkItemsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListWorkItemsResponseBody
	GetNextToken() *string
	SetPage(v int64) *ListWorkItemsResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListWorkItemsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListWorkItemsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListWorkItemsResponseBody
	GetTotal() *int64
}

type ListWorkItemsResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether more data is available.
	//
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The list of work items on the current page.
	//
	// example:
	//
	// example
	Items []*ListWorkItemsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of items returned per page in the current request.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The error message.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page. An empty string indicates that no more data is available.
	//
	// example:
	//
	// example
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of items per page.
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of visible items under the filter conditions.
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListWorkItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkItemsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWorkItemsResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListWorkItemsResponseBody) GetItems() []*ListWorkItemsResponseBodyItems {
	return s.Items
}

func (s *ListWorkItemsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListWorkItemsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkItemsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkItemsResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListWorkItemsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWorkItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkItemsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListWorkItemsResponseBody) SetCode(v string) *ListWorkItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkItemsResponseBody) SetHasMore(v bool) *ListWorkItemsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListWorkItemsResponseBody) SetItems(v []*ListWorkItemsResponseBodyItems) *ListWorkItemsResponseBody {
	s.Items = v
	return s
}

func (s *ListWorkItemsResponseBody) SetMaxResults(v int64) *ListWorkItemsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkItemsResponseBody) SetMessage(v string) *ListWorkItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkItemsResponseBody) SetNextToken(v string) *ListWorkItemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkItemsResponseBody) SetPage(v int64) *ListWorkItemsResponseBody {
	s.Page = &v
	return s
}

func (s *ListWorkItemsResponseBody) SetPageSize(v int64) *ListWorkItemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListWorkItemsResponseBody) SetRequestId(v string) *ListWorkItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkItemsResponseBody) SetTotal(v int64) *ListWorkItemsResponseBody {
	s.Total = &v
	return s
}

func (s *ListWorkItemsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkItemsResponseBodyItems struct {
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
	Attachments []*ListWorkItemsResponseBodyItemsAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
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
	// The creation time in ISO 8601 format.
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
	// The version of the input snapshot.
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
	// The time when the item last entered the processing state in ISO 8601 format.
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
	// The list of user IDs of visible members.
	//
	// example:
	//
	// example
	MemberUserIds []*int64 `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
	// The modification time in ISO 8601 format.
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
	// The trigger time of the task in ISO 8601 format.
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
	// The workflow status projection.
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

func (s ListWorkItemsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListWorkItemsResponseBodyItems) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *ListWorkItemsResponseBodyItems) GetAppliedStatusVersion() *int64 {
	return s.AppliedStatusVersion
}

func (s *ListWorkItemsResponseBodyItems) GetAttachments() []*ListWorkItemsResponseBodyItemsAttachments {
	return s.Attachments
}

func (s *ListWorkItemsResponseBodyItems) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *ListWorkItemsResponseBodyItems) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *ListWorkItemsResponseBodyItems) GetCanManage() *bool {
	return s.CanManage
}

func (s *ListWorkItemsResponseBodyItems) GetCanManageVisibility() *bool {
	return s.CanManageVisibility
}

func (s *ListWorkItemsResponseBodyItems) GetCanRename() *bool {
	return s.CanRename
}

func (s *ListWorkItemsResponseBodyItems) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListWorkItemsResponseBodyItems) GetCollaboratorUserIds() []*int64 {
	return s.CollaboratorUserIds
}

func (s *ListWorkItemsResponseBodyItems) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListWorkItemsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListWorkItemsResponseBodyItems) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *ListWorkItemsResponseBodyItems) GetCreationType() *string {
	return s.CreationType
}

func (s *ListWorkItemsResponseBodyItems) GetCurrentPlanId() *string {
	return s.CurrentPlanId
}

func (s *ListWorkItemsResponseBodyItems) GetCurrentRevision() *int64 {
	return s.CurrentRevision
}

func (s *ListWorkItemsResponseBodyItems) GetCurrentRunId() *string {
	return s.CurrentRunId
}

func (s *ListWorkItemsResponseBodyItems) GetCurrentRunSeq() *int64 {
	return s.CurrentRunSeq
}

func (s *ListWorkItemsResponseBodyItems) GetCurrentStepName() *string {
	return s.CurrentStepName
}

func (s *ListWorkItemsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListWorkItemsResponseBodyItems) GetInputVersion() *int64 {
	return s.InputVersion
}

func (s *ListWorkItemsResponseBodyItems) GetItemType() *string {
	return s.ItemType
}

func (s *ListWorkItemsResponseBodyItems) GetLastProcessingAt() *string {
	return s.LastProcessingAt
}

func (s *ListWorkItemsResponseBodyItems) GetLegacyTopic() *bool {
	return s.LegacyTopic
}

func (s *ListWorkItemsResponseBodyItems) GetMemberUserIds() []*int64 {
	return s.MemberUserIds
}

func (s *ListWorkItemsResponseBodyItems) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *ListWorkItemsResponseBodyItems) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *ListWorkItemsResponseBodyItems) GetPinned() *bool {
	return s.Pinned
}

func (s *ListWorkItemsResponseBodyItems) GetPriority() *string {
	return s.Priority
}

func (s *ListWorkItemsResponseBodyItems) GetProgressCompleted() *int64 {
	return s.ProgressCompleted
}

func (s *ListWorkItemsResponseBodyItems) GetProgressTotal() *int64 {
	return s.ProgressTotal
}

func (s *ListWorkItemsResponseBodyItems) GetSourceExecutionId() *string {
	return s.SourceExecutionId
}

func (s *ListWorkItemsResponseBodyItems) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *ListWorkItemsResponseBodyItems) GetSourceTaskName() *string {
	return s.SourceTaskName
}

func (s *ListWorkItemsResponseBodyItems) GetSourceTriggerKey() *string {
	return s.SourceTriggerKey
}

func (s *ListWorkItemsResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *ListWorkItemsResponseBodyItems) GetTitle() *string {
	return s.Title
}

func (s *ListWorkItemsResponseBodyItems) GetTriggerTimezone() *string {
	return s.TriggerTimezone
}

func (s *ListWorkItemsResponseBodyItems) GetTriggeredAt() *string {
	return s.TriggeredAt
}

func (s *ListWorkItemsResponseBodyItems) GetVersion() *int64 {
	return s.Version
}

func (s *ListWorkItemsResponseBodyItems) GetVisibility() *string {
	return s.Visibility
}

func (s *ListWorkItemsResponseBodyItems) GetVisibilitySource() *string {
	return s.VisibilitySource
}

func (s *ListWorkItemsResponseBodyItems) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ListWorkItemsResponseBodyItems) GetWorkflowStatus() *string {
	return s.WorkflowStatus
}

func (s *ListWorkItemsResponseBodyItems) GetWorkflowTemplateIds() []*string {
	return s.WorkflowTemplateIds
}

func (s *ListWorkItemsResponseBodyItems) SetAcceptanceCriteria(v string) *ListWorkItemsResponseBodyItems {
	s.AcceptanceCriteria = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetAppliedStatusVersion(v int64) *ListWorkItemsResponseBodyItems {
	s.AppliedStatusVersion = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetAttachments(v []*ListWorkItemsResponseBodyItemsAttachments) *ListWorkItemsResponseBodyItems {
	s.Attachments = v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetBusinessStatus(v string) *ListWorkItemsResponseBodyItems {
	s.BusinessStatus = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetCanDelete(v bool) *ListWorkItemsResponseBodyItems {
	s.CanDelete = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetCanManage(v bool) *ListWorkItemsResponseBodyItems {
	s.CanManage = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetCanManageVisibility(v bool) *ListWorkItemsResponseBodyItems {
	s.CanManageVisibility = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetCanRename(v bool) *ListWorkItemsResponseBodyItems {
	s.CanRename = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetCollaborationGroupId(v string) *ListWorkItemsResponseBodyItems {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetCollaboratorUserIds(v []*int64) *ListWorkItemsResponseBodyItems {
	s.CollaboratorUserIds = v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetConversationId(v string) *ListWorkItemsResponseBodyItems {
	s.ConversationId = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetCreatedAt(v string) *ListWorkItemsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetCreatedByUserId(v int64) *ListWorkItemsResponseBodyItems {
	s.CreatedByUserId = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetCreationType(v string) *ListWorkItemsResponseBodyItems {
	s.CreationType = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetCurrentPlanId(v string) *ListWorkItemsResponseBodyItems {
	s.CurrentPlanId = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetCurrentRevision(v int64) *ListWorkItemsResponseBodyItems {
	s.CurrentRevision = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetCurrentRunId(v string) *ListWorkItemsResponseBodyItems {
	s.CurrentRunId = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetCurrentRunSeq(v int64) *ListWorkItemsResponseBodyItems {
	s.CurrentRunSeq = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetCurrentStepName(v string) *ListWorkItemsResponseBodyItems {
	s.CurrentStepName = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetDescription(v string) *ListWorkItemsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetInputVersion(v int64) *ListWorkItemsResponseBodyItems {
	s.InputVersion = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetItemType(v string) *ListWorkItemsResponseBodyItems {
	s.ItemType = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetLastProcessingAt(v string) *ListWorkItemsResponseBodyItems {
	s.LastProcessingAt = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetLegacyTopic(v bool) *ListWorkItemsResponseBodyItems {
	s.LegacyTopic = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetMemberUserIds(v []*int64) *ListWorkItemsResponseBodyItems {
	s.MemberUserIds = v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetModifiedAt(v string) *ListWorkItemsResponseBodyItems {
	s.ModifiedAt = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetOwnerUserId(v int64) *ListWorkItemsResponseBodyItems {
	s.OwnerUserId = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetPinned(v bool) *ListWorkItemsResponseBodyItems {
	s.Pinned = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetPriority(v string) *ListWorkItemsResponseBodyItems {
	s.Priority = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetProgressCompleted(v int64) *ListWorkItemsResponseBodyItems {
	s.ProgressCompleted = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetProgressTotal(v int64) *ListWorkItemsResponseBodyItems {
	s.ProgressTotal = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetSourceExecutionId(v string) *ListWorkItemsResponseBodyItems {
	s.SourceExecutionId = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetSourceTaskId(v string) *ListWorkItemsResponseBodyItems {
	s.SourceTaskId = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetSourceTaskName(v string) *ListWorkItemsResponseBodyItems {
	s.SourceTaskName = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetSourceTriggerKey(v string) *ListWorkItemsResponseBodyItems {
	s.SourceTriggerKey = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetSourceType(v string) *ListWorkItemsResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetTitle(v string) *ListWorkItemsResponseBodyItems {
	s.Title = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetTriggerTimezone(v string) *ListWorkItemsResponseBodyItems {
	s.TriggerTimezone = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetTriggeredAt(v string) *ListWorkItemsResponseBodyItems {
	s.TriggeredAt = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetVersion(v int64) *ListWorkItemsResponseBodyItems {
	s.Version = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetVisibility(v string) *ListWorkItemsResponseBodyItems {
	s.Visibility = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetVisibilitySource(v string) *ListWorkItemsResponseBodyItems {
	s.VisibilitySource = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetWorkItemId(v string) *ListWorkItemsResponseBodyItems {
	s.WorkItemId = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetWorkflowStatus(v string) *ListWorkItemsResponseBodyItems {
	s.WorkflowStatus = &v
	return s
}

func (s *ListWorkItemsResponseBodyItems) SetWorkflowTemplateIds(v []*string) *ListWorkItemsResponseBodyItems {
	s.WorkflowTemplateIds = v
	return s
}

func (s *ListWorkItemsResponseBodyItems) Validate() error {
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

type ListWorkItemsResponseBodyItemsAttachments struct {
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
	// The access URL of the attachment.
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

func (s ListWorkItemsResponseBodyItemsAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemsResponseBodyItemsAttachments) GoString() string {
	return s.String()
}

func (s *ListWorkItemsResponseBodyItemsAttachments) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *ListWorkItemsResponseBodyItemsAttachments) GetContentType() *string {
	return s.ContentType
}

func (s *ListWorkItemsResponseBodyItemsAttachments) GetRefUrl() *string {
	return s.RefUrl
}

func (s *ListWorkItemsResponseBodyItemsAttachments) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListWorkItemsResponseBodyItemsAttachments) GetResourceKind() *string {
	return s.ResourceKind
}

func (s *ListWorkItemsResponseBodyItemsAttachments) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *ListWorkItemsResponseBodyItemsAttachments) GetTitle() *string {
	return s.Title
}

func (s *ListWorkItemsResponseBodyItemsAttachments) SetAttachmentId(v string) *ListWorkItemsResponseBodyItemsAttachments {
	s.AttachmentId = &v
	return s
}

func (s *ListWorkItemsResponseBodyItemsAttachments) SetContentType(v string) *ListWorkItemsResponseBodyItemsAttachments {
	s.ContentType = &v
	return s
}

func (s *ListWorkItemsResponseBodyItemsAttachments) SetRefUrl(v string) *ListWorkItemsResponseBodyItemsAttachments {
	s.RefUrl = &v
	return s
}

func (s *ListWorkItemsResponseBodyItemsAttachments) SetResourceId(v string) *ListWorkItemsResponseBodyItemsAttachments {
	s.ResourceId = &v
	return s
}

func (s *ListWorkItemsResponseBodyItemsAttachments) SetResourceKind(v string) *ListWorkItemsResponseBodyItemsAttachments {
	s.ResourceKind = &v
	return s
}

func (s *ListWorkItemsResponseBodyItemsAttachments) SetSizeBytes(v int64) *ListWorkItemsResponseBodyItemsAttachments {
	s.SizeBytes = &v
	return s
}

func (s *ListWorkItemsResponseBodyItemsAttachments) SetTitle(v string) *ListWorkItemsResponseBodyItemsAttachments {
	s.Title = &v
	return s
}

func (s *ListWorkItemsResponseBodyItemsAttachments) Validate() error {
	return dara.Validate(s)
}
