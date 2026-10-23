// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorkItemFeedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryWorkItemFeedResponseBody
	GetCode() *string
	SetEntries(v []*QueryWorkItemFeedResponseBodyEntries) *QueryWorkItemFeedResponseBody
	GetEntries() []*QueryWorkItemFeedResponseBodyEntries
	SetHasMore(v bool) *QueryWorkItemFeedResponseBody
	GetHasMore() *bool
	SetMessage(v string) *QueryWorkItemFeedResponseBody
	GetMessage() *string
	SetNextCursor(v string) *QueryWorkItemFeedResponseBody
	GetNextCursor() *string
	SetRequestId(v string) *QueryWorkItemFeedResponseBody
	GetRequestId() *string
	SetScopeCounts(v map[string]interface{}) *QueryWorkItemFeedResponseBody
	GetScopeCounts() map[string]interface{}
	SetStatusCounts(v map[string]interface{}) *QueryWorkItemFeedResponseBody
	GetStatusCounts() map[string]interface{}
	SetTotalEntryCount(v int64) *QueryWorkItemFeedResponseBody
	GetTotalEntryCount() *int64
	SetTotalWorkItemCount(v int64) *QueryWorkItemFeedResponseBody
	GetTotalWorkItemCount() *int64
}

type QueryWorkItemFeedResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The entries on the current page. The entryType determines the effective fields: WORK_ITEM maps to item, and TASK_GROUP maps to aggregation fields.
	//
	// example:
	//
	// example
	Entries []*QueryWorkItemFeedResponseBodyEntries `json:"entries,omitempty" xml:"entries,omitempty" type:"Repeated"`
	// Indicates whether more data is available.
	//
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The cursor for the next page.
	//
	// example:
	//
	// example
	NextCursor *string `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The number of items in the visible pool by personal scope.
	//
	// example:
	//
	// example
	ScopeCounts map[string]interface{} `json:"scopeCounts,omitempty" xml:"scopeCounts,omitempty"`
	// The number of items in the visible pool by business status.
	//
	// example:
	//
	// example
	StatusCounts map[string]interface{} `json:"statusCounts,omitempty" xml:"statusCounts,omitempty"`
	// The total number of entries in the visible pool.
	//
	// example:
	//
	// 1
	TotalEntryCount *int64 `json:"totalEntryCount,omitempty" xml:"totalEntryCount,omitempty"`
	// The total number of work items in the visible pool.
	//
	// example:
	//
	// 1
	TotalWorkItemCount *int64 `json:"totalWorkItemCount,omitempty" xml:"totalWorkItemCount,omitempty"`
}

func (s QueryWorkItemFeedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkItemFeedResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWorkItemFeedResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryWorkItemFeedResponseBody) GetEntries() []*QueryWorkItemFeedResponseBodyEntries {
	return s.Entries
}

func (s *QueryWorkItemFeedResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *QueryWorkItemFeedResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryWorkItemFeedResponseBody) GetNextCursor() *string {
	return s.NextCursor
}

func (s *QueryWorkItemFeedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryWorkItemFeedResponseBody) GetScopeCounts() map[string]interface{} {
	return s.ScopeCounts
}

func (s *QueryWorkItemFeedResponseBody) GetStatusCounts() map[string]interface{} {
	return s.StatusCounts
}

func (s *QueryWorkItemFeedResponseBody) GetTotalEntryCount() *int64 {
	return s.TotalEntryCount
}

func (s *QueryWorkItemFeedResponseBody) GetTotalWorkItemCount() *int64 {
	return s.TotalWorkItemCount
}

func (s *QueryWorkItemFeedResponseBody) SetCode(v string) *QueryWorkItemFeedResponseBody {
	s.Code = &v
	return s
}

func (s *QueryWorkItemFeedResponseBody) SetEntries(v []*QueryWorkItemFeedResponseBodyEntries) *QueryWorkItemFeedResponseBody {
	s.Entries = v
	return s
}

func (s *QueryWorkItemFeedResponseBody) SetHasMore(v bool) *QueryWorkItemFeedResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryWorkItemFeedResponseBody) SetMessage(v string) *QueryWorkItemFeedResponseBody {
	s.Message = &v
	return s
}

func (s *QueryWorkItemFeedResponseBody) SetNextCursor(v string) *QueryWorkItemFeedResponseBody {
	s.NextCursor = &v
	return s
}

func (s *QueryWorkItemFeedResponseBody) SetRequestId(v string) *QueryWorkItemFeedResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBody) SetScopeCounts(v map[string]interface{}) *QueryWorkItemFeedResponseBody {
	s.ScopeCounts = v
	return s
}

func (s *QueryWorkItemFeedResponseBody) SetStatusCounts(v map[string]interface{}) *QueryWorkItemFeedResponseBody {
	s.StatusCounts = v
	return s
}

func (s *QueryWorkItemFeedResponseBody) SetTotalEntryCount(v int64) *QueryWorkItemFeedResponseBody {
	s.TotalEntryCount = &v
	return s
}

func (s *QueryWorkItemFeedResponseBody) SetTotalWorkItemCount(v int64) *QueryWorkItemFeedResponseBody {
	s.TotalWorkItemCount = &v
	return s
}

func (s *QueryWorkItemFeedResponseBody) Validate() error {
	if s.Entries != nil {
		for _, item := range s.Entries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryWorkItemFeedResponseBodyEntries struct {
	// The entry type.
	//
	// example:
	//
	// WORK_ITEM
	EntryType *string `json:"entryType,omitempty" xml:"entryType,omitempty"`
	// Indicates whether more child items are available.
	//
	// example:
	//
	// false
	HasMoreChildren *bool `json:"hasMoreChildren,omitempty" xml:"hasMoreChildren,omitempty"`
	// The content of a regular work item row.
	//
	// example:
	//
	// example
	Item *QueryWorkItemFeedResponseBodyEntriesItem `json:"item,omitempty" xml:"item,omitempty" type:"Struct"`
	// The latest five work items in the task aggregation row.
	//
	// example:
	//
	// example
	Items []*QueryWorkItemFeedResponseBodyEntriesItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The time of the latest work item, in ISO 8601 format.
	//
	// example:
	//
	// example
	LatestWorkItemAt *string `json:"latestWorkItemAt,omitempty" xml:"latestWorkItemAt,omitempty"`
	// The total number of matched work items.
	//
	// example:
	//
	// 1
	MatchedWorkItemCount *int64 `json:"matchedWorkItemCount,omitempty" xml:"matchedWorkItemCount,omitempty"`
	// The source task ID.
	//
	// example:
	//
	// example
	SourceTaskId *string `json:"sourceTaskId,omitempty" xml:"sourceTaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// example
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s QueryWorkItemFeedResponseBodyEntries) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkItemFeedResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *QueryWorkItemFeedResponseBodyEntries) GetEntryType() *string {
	return s.EntryType
}

func (s *QueryWorkItemFeedResponseBodyEntries) GetHasMoreChildren() *bool {
	return s.HasMoreChildren
}

func (s *QueryWorkItemFeedResponseBodyEntries) GetItem() *QueryWorkItemFeedResponseBodyEntriesItem {
	return s.Item
}

func (s *QueryWorkItemFeedResponseBodyEntries) GetItems() []*QueryWorkItemFeedResponseBodyEntriesItems {
	return s.Items
}

func (s *QueryWorkItemFeedResponseBodyEntries) GetLatestWorkItemAt() *string {
	return s.LatestWorkItemAt
}

func (s *QueryWorkItemFeedResponseBodyEntries) GetMatchedWorkItemCount() *int64 {
	return s.MatchedWorkItemCount
}

func (s *QueryWorkItemFeedResponseBodyEntries) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *QueryWorkItemFeedResponseBodyEntries) GetTaskName() *string {
	return s.TaskName
}

func (s *QueryWorkItemFeedResponseBodyEntries) SetEntryType(v string) *QueryWorkItemFeedResponseBodyEntries {
	s.EntryType = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntries) SetHasMoreChildren(v bool) *QueryWorkItemFeedResponseBodyEntries {
	s.HasMoreChildren = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntries) SetItem(v *QueryWorkItemFeedResponseBodyEntriesItem) *QueryWorkItemFeedResponseBodyEntries {
	s.Item = v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntries) SetItems(v []*QueryWorkItemFeedResponseBodyEntriesItems) *QueryWorkItemFeedResponseBodyEntries {
	s.Items = v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntries) SetLatestWorkItemAt(v string) *QueryWorkItemFeedResponseBodyEntries {
	s.LatestWorkItemAt = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntries) SetMatchedWorkItemCount(v int64) *QueryWorkItemFeedResponseBodyEntries {
	s.MatchedWorkItemCount = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntries) SetSourceTaskId(v string) *QueryWorkItemFeedResponseBodyEntries {
	s.SourceTaskId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntries) SetTaskName(v string) *QueryWorkItemFeedResponseBodyEntries {
	s.TaskName = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntries) Validate() error {
	if s.Item != nil {
		if err := s.Item.Validate(); err != nil {
			return err
		}
	}
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

type QueryWorkItemFeedResponseBodyEntriesItem struct {
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
	Attachments []*QueryWorkItemFeedResponseBodyEntriesItemAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
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
	// The associated conversation ID.
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
	// The list of user IDs of visible members.
	//
	// example:
	//
	// example
	MemberUserIds []*int64 `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
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
	// The pinned status for the caller.
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
	// The time zone when the task is triggered.
	//
	// example:
	//
	// example
	TriggerTimezone *string `json:"triggerTimezone,omitempty" xml:"triggerTimezone,omitempty"`
	// The task trigger time, in ISO 8601 format.
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
	// The list of bound workflow template IDs.
	//
	// example:
	//
	// example
	WorkflowTemplateIds []*string `json:"workflowTemplateIds,omitempty" xml:"workflowTemplateIds,omitempty" type:"Repeated"`
}

func (s QueryWorkItemFeedResponseBodyEntriesItem) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkItemFeedResponseBodyEntriesItem) GoString() string {
	return s.String()
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetAppliedStatusVersion() *int64 {
	return s.AppliedStatusVersion
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetAttachments() []*QueryWorkItemFeedResponseBodyEntriesItemAttachments {
	return s.Attachments
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetCanManage() *bool {
	return s.CanManage
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetCanManageVisibility() *bool {
	return s.CanManageVisibility
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetCanRename() *bool {
	return s.CanRename
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetCollaboratorUserIds() []*int64 {
	return s.CollaboratorUserIds
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetConversationId() *string {
	return s.ConversationId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetCreationType() *string {
	return s.CreationType
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetCurrentPlanId() *string {
	return s.CurrentPlanId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetCurrentRevision() *int64 {
	return s.CurrentRevision
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetCurrentRunId() *string {
	return s.CurrentRunId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetCurrentRunSeq() *int64 {
	return s.CurrentRunSeq
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetCurrentStepName() *string {
	return s.CurrentStepName
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetDescription() *string {
	return s.Description
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetInputVersion() *int64 {
	return s.InputVersion
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetItemType() *string {
	return s.ItemType
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetLastProcessingAt() *string {
	return s.LastProcessingAt
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetLegacyTopic() *bool {
	return s.LegacyTopic
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetMemberUserIds() []*int64 {
	return s.MemberUserIds
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetPinned() *bool {
	return s.Pinned
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetPriority() *string {
	return s.Priority
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetProgressCompleted() *int64 {
	return s.ProgressCompleted
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetProgressTotal() *int64 {
	return s.ProgressTotal
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetSourceExecutionId() *string {
	return s.SourceExecutionId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetSourceTaskName() *string {
	return s.SourceTaskName
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetSourceTriggerKey() *string {
	return s.SourceTriggerKey
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetSourceType() *string {
	return s.SourceType
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetTitle() *string {
	return s.Title
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetTriggerTimezone() *string {
	return s.TriggerTimezone
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetTriggeredAt() *string {
	return s.TriggeredAt
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetVersion() *int64 {
	return s.Version
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetVisibility() *string {
	return s.Visibility
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetVisibilitySource() *string {
	return s.VisibilitySource
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetWorkflowStatus() *string {
	return s.WorkflowStatus
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) GetWorkflowTemplateIds() []*string {
	return s.WorkflowTemplateIds
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetAcceptanceCriteria(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.AcceptanceCriteria = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetAppliedStatusVersion(v int64) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.AppliedStatusVersion = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetAttachments(v []*QueryWorkItemFeedResponseBodyEntriesItemAttachments) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.Attachments = v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetBusinessStatus(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.BusinessStatus = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetCanDelete(v bool) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.CanDelete = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetCanManage(v bool) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.CanManage = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetCanManageVisibility(v bool) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.CanManageVisibility = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetCanRename(v bool) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.CanRename = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetCollaborationGroupId(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.CollaborationGroupId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetCollaboratorUserIds(v []*int64) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.CollaboratorUserIds = v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetConversationId(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.ConversationId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetCreatedAt(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.CreatedAt = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetCreatedByUserId(v int64) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.CreatedByUserId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetCreationType(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.CreationType = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetCurrentPlanId(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.CurrentPlanId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetCurrentRevision(v int64) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.CurrentRevision = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetCurrentRunId(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.CurrentRunId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetCurrentRunSeq(v int64) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.CurrentRunSeq = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetCurrentStepName(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.CurrentStepName = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetDescription(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.Description = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetInputVersion(v int64) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.InputVersion = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetItemType(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.ItemType = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetLastProcessingAt(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.LastProcessingAt = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetLegacyTopic(v bool) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.LegacyTopic = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetMemberUserIds(v []*int64) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.MemberUserIds = v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetModifiedAt(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.ModifiedAt = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetOwnerUserId(v int64) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.OwnerUserId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetPinned(v bool) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.Pinned = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetPriority(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.Priority = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetProgressCompleted(v int64) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.ProgressCompleted = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetProgressTotal(v int64) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.ProgressTotal = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetSourceExecutionId(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.SourceExecutionId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetSourceTaskId(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.SourceTaskId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetSourceTaskName(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.SourceTaskName = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetSourceTriggerKey(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.SourceTriggerKey = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetSourceType(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.SourceType = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetTitle(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.Title = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetTriggerTimezone(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.TriggerTimezone = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetTriggeredAt(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.TriggeredAt = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetVersion(v int64) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.Version = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetVisibility(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.Visibility = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetVisibilitySource(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.VisibilitySource = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetWorkItemId(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.WorkItemId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetWorkflowStatus(v string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.WorkflowStatus = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) SetWorkflowTemplateIds(v []*string) *QueryWorkItemFeedResponseBodyEntriesItem {
	s.WorkflowTemplateIds = v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItem) Validate() error {
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

type QueryWorkItemFeedResponseBodyEntriesItemAttachments struct {
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
	// The attachment resource ID.
	//
	// example:
	//
	// example
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The attachment resource type.
	//
	// example:
	//
	// example
	ResourceKind *string `json:"resourceKind,omitempty" xml:"resourceKind,omitempty"`
	// The attachment size in bytes.
	//
	// example:
	//
	// 1
	SizeBytes *int64 `json:"sizeBytes,omitempty" xml:"sizeBytes,omitempty"`
	// The attachment title.
	//
	// example:
	//
	// example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryWorkItemFeedResponseBodyEntriesItemAttachments) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkItemFeedResponseBodyEntriesItemAttachments) GoString() string {
	return s.String()
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) GetContentType() *string {
	return s.ContentType
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) GetRefUrl() *string {
	return s.RefUrl
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) GetResourceKind() *string {
	return s.ResourceKind
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) GetTitle() *string {
	return s.Title
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) SetAttachmentId(v string) *QueryWorkItemFeedResponseBodyEntriesItemAttachments {
	s.AttachmentId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) SetContentType(v string) *QueryWorkItemFeedResponseBodyEntriesItemAttachments {
	s.ContentType = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) SetRefUrl(v string) *QueryWorkItemFeedResponseBodyEntriesItemAttachments {
	s.RefUrl = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) SetResourceId(v string) *QueryWorkItemFeedResponseBodyEntriesItemAttachments {
	s.ResourceId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) SetResourceKind(v string) *QueryWorkItemFeedResponseBodyEntriesItemAttachments {
	s.ResourceKind = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) SetSizeBytes(v int64) *QueryWorkItemFeedResponseBodyEntriesItemAttachments {
	s.SizeBytes = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) SetTitle(v string) *QueryWorkItemFeedResponseBodyEntriesItemAttachments {
	s.Title = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemAttachments) Validate() error {
	return dara.Validate(s)
}

type QueryWorkItemFeedResponseBodyEntriesItems struct {
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
	Attachments []*QueryWorkItemFeedResponseBodyEntriesItemsAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
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
	// The associated conversation ID.
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
	// The list of user IDs of visible members.
	//
	// example:
	//
	// example
	MemberUserIds []*int64 `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
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
	// The pinned status for the caller.
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
	// The time zone when the task is triggered.
	//
	// example:
	//
	// example
	TriggerTimezone *string `json:"triggerTimezone,omitempty" xml:"triggerTimezone,omitempty"`
	// The task trigger time, in ISO 8601 format.
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
	// The list of bound workflow template IDs.
	//
	// example:
	//
	// example
	WorkflowTemplateIds []*string `json:"workflowTemplateIds,omitempty" xml:"workflowTemplateIds,omitempty" type:"Repeated"`
}

func (s QueryWorkItemFeedResponseBodyEntriesItems) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkItemFeedResponseBodyEntriesItems) GoString() string {
	return s.String()
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetAppliedStatusVersion() *int64 {
	return s.AppliedStatusVersion
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetAttachments() []*QueryWorkItemFeedResponseBodyEntriesItemsAttachments {
	return s.Attachments
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetCanManage() *bool {
	return s.CanManage
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetCanManageVisibility() *bool {
	return s.CanManageVisibility
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetCanRename() *bool {
	return s.CanRename
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetCollaboratorUserIds() []*int64 {
	return s.CollaboratorUserIds
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetConversationId() *string {
	return s.ConversationId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetCreationType() *string {
	return s.CreationType
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetCurrentPlanId() *string {
	return s.CurrentPlanId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetCurrentRevision() *int64 {
	return s.CurrentRevision
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetCurrentRunId() *string {
	return s.CurrentRunId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetCurrentRunSeq() *int64 {
	return s.CurrentRunSeq
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetCurrentStepName() *string {
	return s.CurrentStepName
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetDescription() *string {
	return s.Description
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetInputVersion() *int64 {
	return s.InputVersion
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetItemType() *string {
	return s.ItemType
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetLastProcessingAt() *string {
	return s.LastProcessingAt
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetLegacyTopic() *bool {
	return s.LegacyTopic
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetMemberUserIds() []*int64 {
	return s.MemberUserIds
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetPinned() *bool {
	return s.Pinned
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetPriority() *string {
	return s.Priority
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetProgressCompleted() *int64 {
	return s.ProgressCompleted
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetProgressTotal() *int64 {
	return s.ProgressTotal
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetSourceExecutionId() *string {
	return s.SourceExecutionId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetSourceTaskId() *string {
	return s.SourceTaskId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetSourceTaskName() *string {
	return s.SourceTaskName
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetSourceTriggerKey() *string {
	return s.SourceTriggerKey
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetSourceType() *string {
	return s.SourceType
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetTitle() *string {
	return s.Title
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetTriggerTimezone() *string {
	return s.TriggerTimezone
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetTriggeredAt() *string {
	return s.TriggeredAt
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetVersion() *int64 {
	return s.Version
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetVisibility() *string {
	return s.Visibility
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetVisibilitySource() *string {
	return s.VisibilitySource
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetWorkflowStatus() *string {
	return s.WorkflowStatus
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) GetWorkflowTemplateIds() []*string {
	return s.WorkflowTemplateIds
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetAcceptanceCriteria(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.AcceptanceCriteria = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetAppliedStatusVersion(v int64) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.AppliedStatusVersion = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetAttachments(v []*QueryWorkItemFeedResponseBodyEntriesItemsAttachments) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.Attachments = v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetBusinessStatus(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.BusinessStatus = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetCanDelete(v bool) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.CanDelete = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetCanManage(v bool) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.CanManage = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetCanManageVisibility(v bool) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.CanManageVisibility = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetCanRename(v bool) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.CanRename = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetCollaborationGroupId(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.CollaborationGroupId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetCollaboratorUserIds(v []*int64) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.CollaboratorUserIds = v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetConversationId(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.ConversationId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetCreatedAt(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.CreatedAt = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetCreatedByUserId(v int64) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.CreatedByUserId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetCreationType(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.CreationType = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetCurrentPlanId(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.CurrentPlanId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetCurrentRevision(v int64) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.CurrentRevision = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetCurrentRunId(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.CurrentRunId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetCurrentRunSeq(v int64) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.CurrentRunSeq = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetCurrentStepName(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.CurrentStepName = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetDescription(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.Description = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetInputVersion(v int64) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.InputVersion = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetItemType(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.ItemType = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetLastProcessingAt(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.LastProcessingAt = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetLegacyTopic(v bool) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.LegacyTopic = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetMemberUserIds(v []*int64) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.MemberUserIds = v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetModifiedAt(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.ModifiedAt = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetOwnerUserId(v int64) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.OwnerUserId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetPinned(v bool) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.Pinned = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetPriority(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.Priority = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetProgressCompleted(v int64) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.ProgressCompleted = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetProgressTotal(v int64) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.ProgressTotal = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetSourceExecutionId(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.SourceExecutionId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetSourceTaskId(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.SourceTaskId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetSourceTaskName(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.SourceTaskName = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetSourceTriggerKey(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.SourceTriggerKey = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetSourceType(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.SourceType = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetTitle(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.Title = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetTriggerTimezone(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.TriggerTimezone = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetTriggeredAt(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.TriggeredAt = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetVersion(v int64) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.Version = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetVisibility(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.Visibility = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetVisibilitySource(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.VisibilitySource = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetWorkItemId(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.WorkItemId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetWorkflowStatus(v string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.WorkflowStatus = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) SetWorkflowTemplateIds(v []*string) *QueryWorkItemFeedResponseBodyEntriesItems {
	s.WorkflowTemplateIds = v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItems) Validate() error {
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

type QueryWorkItemFeedResponseBodyEntriesItemsAttachments struct {
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
	// The attachment resource ID.
	//
	// example:
	//
	// example
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The attachment resource type.
	//
	// example:
	//
	// example
	ResourceKind *string `json:"resourceKind,omitempty" xml:"resourceKind,omitempty"`
	// The attachment size in bytes.
	//
	// example:
	//
	// 1
	SizeBytes *int64 `json:"sizeBytes,omitempty" xml:"sizeBytes,omitempty"`
	// The attachment title.
	//
	// example:
	//
	// example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryWorkItemFeedResponseBodyEntriesItemsAttachments) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkItemFeedResponseBodyEntriesItemsAttachments) GoString() string {
	return s.String()
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) GetContentType() *string {
	return s.ContentType
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) GetRefUrl() *string {
	return s.RefUrl
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) GetResourceKind() *string {
	return s.ResourceKind
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) GetTitle() *string {
	return s.Title
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) SetAttachmentId(v string) *QueryWorkItemFeedResponseBodyEntriesItemsAttachments {
	s.AttachmentId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) SetContentType(v string) *QueryWorkItemFeedResponseBodyEntriesItemsAttachments {
	s.ContentType = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) SetRefUrl(v string) *QueryWorkItemFeedResponseBodyEntriesItemsAttachments {
	s.RefUrl = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) SetResourceId(v string) *QueryWorkItemFeedResponseBodyEntriesItemsAttachments {
	s.ResourceId = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) SetResourceKind(v string) *QueryWorkItemFeedResponseBodyEntriesItemsAttachments {
	s.ResourceKind = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) SetSizeBytes(v int64) *QueryWorkItemFeedResponseBodyEntriesItemsAttachments {
	s.SizeBytes = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) SetTitle(v string) *QueryWorkItemFeedResponseBodyEntriesItemsAttachments {
	s.Title = &v
	return s
}

func (s *QueryWorkItemFeedResponseBodyEntriesItemsAttachments) Validate() error {
	return dara.Validate(s)
}
