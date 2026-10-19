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
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 当前页条目；entryType 决定生效字段（WORK_ITEM→item / TASK_GROUP→聚合字段）
	//
	// example:
	//
	// example
	Entries []*QueryWorkItemFeedResponseBodyEntries `json:"entries,omitempty" xml:"entries,omitempty" type:"Repeated"`
	// 是否还有更多数据
	//
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 下一页游标
	//
	// example:
	//
	// example
	NextCursor *string `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 可见池按个人范围计数
	//
	// example:
	//
	// example
	ScopeCounts map[string]interface{} `json:"scopeCounts,omitempty" xml:"scopeCounts,omitempty"`
	// 可见池按业务状态计数
	//
	// example:
	//
	// example
	StatusCounts map[string]interface{} `json:"statusCounts,omitempty" xml:"statusCounts,omitempty"`
	// 可见池条目总数
	//
	// example:
	//
	// 1
	TotalEntryCount *int64 `json:"totalEntryCount,omitempty" xml:"totalEntryCount,omitempty"`
	// 可见池工作项总数
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
	// 条目类型
	//
	// example:
	//
	// WORK_ITEM
	EntryType *string `json:"entryType,omitempty" xml:"entryType,omitempty"`
	// 是否还有更多未展示的子项
	//
	// example:
	//
	// false
	HasMoreChildren *bool `json:"hasMoreChildren,omitempty" xml:"hasMoreChildren,omitempty"`
	// 普通工作项行内容
	//
	// example:
	//
	// example
	Item *QueryWorkItemFeedResponseBodyEntriesItem `json:"item,omitempty" xml:"item,omitempty" type:"Struct"`
	// 任务聚合行内最新五条工作项
	//
	// example:
	//
	// example
	Items []*QueryWorkItemFeedResponseBodyEntriesItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 最近工作项时间(ISO8601)
	//
	// example:
	//
	// example
	LatestWorkItemAt *string `json:"latestWorkItemAt,omitempty" xml:"latestWorkItemAt,omitempty"`
	// 命中的工作项总数
	//
	// example:
	//
	// 1
	MatchedWorkItemCount *int64 `json:"matchedWorkItemCount,omitempty" xml:"matchedWorkItemCount,omitempty"`
	// 来源任务ID
	//
	// example:
	//
	// example
	SourceTaskId *string `json:"sourceTaskId,omitempty" xml:"sourceTaskId,omitempty"`
	// 任务名称
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
	Attachments []*QueryWorkItemFeedResponseBodyEntriesItemAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
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
	Attachments []*QueryWorkItemFeedResponseBodyEntriesItemsAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
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
