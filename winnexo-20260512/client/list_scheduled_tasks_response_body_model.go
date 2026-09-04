// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScheduledTasksResponseBody
	GetCode() *string
	SetHasMore(v bool) *ListScheduledTasksResponseBody
	GetHasMore() *bool
	SetItems(v []*ListScheduledTasksResponseBodyItems) *ListScheduledTasksResponseBody
	GetItems() []*ListScheduledTasksResponseBodyItems
	SetMaxResults(v int32) *ListScheduledTasksResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListScheduledTasksResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListScheduledTasksResponseBody
	GetNextToken() *string
	SetPage(v int64) *ListScheduledTasksResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListScheduledTasksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListScheduledTasksResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListScheduledTasksResponseBody
	GetTotal() *int64
}

type ListScheduledTasksResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether more data is available. Valid values:
	//
	// - true: More data is available.
	//
	// - false: No more data is available.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The list of skill cards.
	Items []*ListScheduledTasksResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// 1763604514518000_531300
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListScheduledTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScheduledTasksResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListScheduledTasksResponseBody) GetItems() []*ListScheduledTasksResponseBodyItems {
	return s.Items
}

func (s *ListScheduledTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListScheduledTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScheduledTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListScheduledTasksResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListScheduledTasksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListScheduledTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScheduledTasksResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListScheduledTasksResponseBody) SetCode(v string) *ListScheduledTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetHasMore(v bool) *ListScheduledTasksResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetItems(v []*ListScheduledTasksResponseBodyItems) *ListScheduledTasksResponseBody {
	s.Items = v
	return s
}

func (s *ListScheduledTasksResponseBody) SetMaxResults(v int32) *ListScheduledTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetMessage(v string) *ListScheduledTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetNextToken(v string) *ListScheduledTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetPage(v int64) *ListScheduledTasksResponseBody {
	s.Page = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetPageSize(v int64) *ListScheduledTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetRequestId(v string) *ListScheduledTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetTotal(v int64) *ListScheduledTasksResponseBody {
	s.Total = &v
	return s
}

func (s *ListScheduledTasksResponseBody) Validate() error {
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

type ListScheduledTasksResponseBodyItems struct {
	// The reason for the exception. This field has a value only when status is abnormal.
	//
	// example:
	//
	// string_value
	AbnormalReason *string `json:"abnormalReason,omitempty" xml:"abnormalReason,omitempty"`
	// Indicates whether the current caller can delete the task (only the task creator and group owner can do so). Always returns true for personal tasks.
	//
	// example:
	//
	// true
	CanDelete *bool `json:"canDelete,omitempty" xml:"canDelete,omitempty"`
	// Indicates whether the task can be edited or deleted.
	//
	// example:
	//
	// true
	CanEdit *bool `json:"canEdit,omitempty" xml:"canEdit,omitempty"`
	// Indicates whether the current caller can immediately execute the task (anyone with visibility can operate. Returns false for abnormal tasks). Always returns true for personal tasks.
	//
	// example:
	//
	// true
	CanExecute *bool `json:"canExecute,omitempty" xml:"canExecute,omitempty"`
	// Indicates whether the current caller can start or stop the task (only the task creator and group owner can do so. Returns false for abnormal tasks). Always returns true for personal tasks.
	//
	// example:
	//
	// true
	CanToggle *bool `json:"canToggle,omitempty" xml:"canToggle,omitempty"`
	// The ID of the collaboration group (such as cg_101). If specified, a group task is created (the caller must be a valid group member). If left empty, a personal task is created.
	//
	// example:
	//
	// exampleCollaborationGroupId
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The creator.
	//
	// example:
	//
	// string_value
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The creator.
	//
	// example:
	//
	// string_value
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// The cron expression.
	//
	// example:
	//
	// string_value
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// The description of the to-do card type.
	//
	// example:
	//
	// SampleDescription
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The list of digital employee names.
	//
	// example:
	//
	// string_value
	DigitalEmployeeName []*string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty" type:"Repeated"`
	// The total number of executions.
	//
	// example:
	//
	// 1
	ExecutionCount *int64 `json:"executionCount,omitempty" xml:"executionCount,omitempty"`
	// The creation time.
	//
	// example:
	//
	// string_value
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// string_value
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Indicates whether public access is enabled.
	//
	// example:
	//
	// true
	IsOpen *bool `json:"isOpen,omitempty" xml:"isOpen,omitempty"`
	// The execution model tier. If not specified, the value is not updated.
	//
	// example:
	//
	// standard
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The task status. Running is returned upon submission.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The trigger type.
	//
	// example:
	//
	// string_value
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
	// The visibility scope of the group task. Valid values:
	//
	// - PRIVATE: visible only to the creator and group owner.
	//
	// - COLLABORATIVE: visible to specified collaborators.
	//
	// - PUBLIC: visible to all group members.
	//
	// For group tasks, the default value is PRIVATE if not specified. This field is ignored for personal tasks.
	//
	// example:
	//
	// PRIVATE
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
	// The list of collaborators (excluding the task creator and group creator, who are covered by the authentication layer). This field is returned only for group tasks. An empty list is returned for PRIVATE or PUBLIC visibility.
	//
	// example:
	//
	// string_value
	VisibleMemberUserIds []*string `json:"visibleMemberUserIds,omitempty" xml:"visibleMemberUserIds,omitempty" type:"Repeated"`
}

func (s ListScheduledTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksResponseBodyItems) GetAbnormalReason() *string {
	return s.AbnormalReason
}

func (s *ListScheduledTasksResponseBodyItems) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *ListScheduledTasksResponseBodyItems) GetCanEdit() *bool {
	return s.CanEdit
}

func (s *ListScheduledTasksResponseBodyItems) GetCanExecute() *bool {
	return s.CanExecute
}

func (s *ListScheduledTasksResponseBodyItems) GetCanToggle() *bool {
	return s.CanToggle
}

func (s *ListScheduledTasksResponseBodyItems) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListScheduledTasksResponseBodyItems) GetCreator() *string {
	return s.Creator
}

func (s *ListScheduledTasksResponseBodyItems) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListScheduledTasksResponseBodyItems) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ListScheduledTasksResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListScheduledTasksResponseBodyItems) GetDigitalEmployeeName() []*string {
	return s.DigitalEmployeeName
}

func (s *ListScheduledTasksResponseBodyItems) GetExecutionCount() *int64 {
	return s.ExecutionCount
}

func (s *ListScheduledTasksResponseBodyItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListScheduledTasksResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListScheduledTasksResponseBodyItems) GetIsOpen() *bool {
	return s.IsOpen
}

func (s *ListScheduledTasksResponseBodyItems) GetModel() *string {
	return s.Model
}

func (s *ListScheduledTasksResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListScheduledTasksResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListScheduledTasksResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *ListScheduledTasksResponseBodyItems) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListScheduledTasksResponseBodyItems) GetVisibility() *string {
	return s.Visibility
}

func (s *ListScheduledTasksResponseBodyItems) GetVisibleMemberUserIds() []*string {
	return s.VisibleMemberUserIds
}

func (s *ListScheduledTasksResponseBodyItems) SetAbnormalReason(v string) *ListScheduledTasksResponseBodyItems {
	s.AbnormalReason = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetCanDelete(v bool) *ListScheduledTasksResponseBodyItems {
	s.CanDelete = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetCanEdit(v bool) *ListScheduledTasksResponseBodyItems {
	s.CanEdit = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetCanExecute(v bool) *ListScheduledTasksResponseBodyItems {
	s.CanExecute = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetCanToggle(v bool) *ListScheduledTasksResponseBodyItems {
	s.CanToggle = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetCollaborationGroupId(v string) *ListScheduledTasksResponseBodyItems {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetCreator(v string) *ListScheduledTasksResponseBodyItems {
	s.Creator = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetCreatorName(v string) *ListScheduledTasksResponseBodyItems {
	s.CreatorName = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetCronExpression(v string) *ListScheduledTasksResponseBodyItems {
	s.CronExpression = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetDescription(v string) *ListScheduledTasksResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetDigitalEmployeeName(v []*string) *ListScheduledTasksResponseBodyItems {
	s.DigitalEmployeeName = v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetExecutionCount(v int64) *ListScheduledTasksResponseBodyItems {
	s.ExecutionCount = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetGmtCreate(v string) *ListScheduledTasksResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetGmtModified(v string) *ListScheduledTasksResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetIsOpen(v bool) *ListScheduledTasksResponseBodyItems {
	s.IsOpen = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetModel(v string) *ListScheduledTasksResponseBodyItems {
	s.Model = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetName(v string) *ListScheduledTasksResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetStatus(v string) *ListScheduledTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetTaskId(v string) *ListScheduledTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetTriggerType(v string) *ListScheduledTasksResponseBodyItems {
	s.TriggerType = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetVisibility(v string) *ListScheduledTasksResponseBodyItems {
	s.Visibility = &v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) SetVisibleMemberUserIds(v []*string) *ListScheduledTasksResponseBodyItems {
	s.VisibleMemberUserIds = v
	return s
}

func (s *ListScheduledTasksResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
