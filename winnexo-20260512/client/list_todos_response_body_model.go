// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTodosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTodosResponseBody
	GetCode() *string
	SetHasMore(v bool) *ListTodosResponseBody
	GetHasMore() *bool
	SetItems(v []*ListTodosResponseBodyItems) *ListTodosResponseBody
	GetItems() []*ListTodosResponseBodyItems
	SetMaxResults(v int64) *ListTodosResponseBody
	GetMaxResults() *int64
	SetMessage(v string) *ListTodosResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTodosResponseBody
	GetNextToken() *string
	SetPage(v int64) *ListTodosResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListTodosResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListTodosResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListTodosResponseBody
	GetTotal() *int64
}

type ListTodosResponseBody struct {
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
	// The list of to-do items on the current page.
	Items []*ListTodosResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The effective maximum number of entries to return per page for the current request.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page. This field is an empty string if no more data is available.
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
	// The effective number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request tracking ID.
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

func (s ListTodosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTodosResponseBody) GoString() string {
	return s.String()
}

func (s *ListTodosResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTodosResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListTodosResponseBody) GetItems() []*ListTodosResponseBodyItems {
	return s.Items
}

func (s *ListTodosResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListTodosResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTodosResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTodosResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListTodosResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTodosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTodosResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListTodosResponseBody) SetCode(v string) *ListTodosResponseBody {
	s.Code = &v
	return s
}

func (s *ListTodosResponseBody) SetHasMore(v bool) *ListTodosResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListTodosResponseBody) SetItems(v []*ListTodosResponseBodyItems) *ListTodosResponseBody {
	s.Items = v
	return s
}

func (s *ListTodosResponseBody) SetMaxResults(v int64) *ListTodosResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTodosResponseBody) SetMessage(v string) *ListTodosResponseBody {
	s.Message = &v
	return s
}

func (s *ListTodosResponseBody) SetNextToken(v string) *ListTodosResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTodosResponseBody) SetPage(v int64) *ListTodosResponseBody {
	s.Page = &v
	return s
}

func (s *ListTodosResponseBody) SetPageSize(v int64) *ListTodosResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTodosResponseBody) SetRequestId(v string) *ListTodosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTodosResponseBody) SetTotal(v int64) *ListTodosResponseBody {
	s.Total = &v
	return s
}

func (s *ListTodosResponseBody) Validate() error {
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

type ListTodosResponseBodyItems struct {
	// The number of assignees.
	//
	// example:
	//
	// 2
	AssigneeCount *int64 `json:"assigneeCount,omitempty" xml:"assigneeCount,omitempty"`
	// The completion time in the YYYY-MM-DD HH:mm:ss format (UTC+8).
	//
	// example:
	//
	// 2026-10-12 10:00:00
	CompletedAt *string `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// The user ID of the user who completed the item. This field is empty if the item is not completed.
	//
	// example:
	//
	// 42
	CompletedByUserId *string `json:"completedByUserId,omitempty" xml:"completedByUserId,omitempty"`
	// The name of the user who completed the item. This field is empty if the item is not completed.
	//
	// example:
	//
	// John Doe
	CompletedByUserName *string `json:"completedByUserName,omitempty" xml:"completedByUserName,omitempty"`
	// The creation time in the YYYY-MM-DD HH:mm:ss format (UTC+8).
	//
	// example:
	//
	// 2026-10-12 10:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time in the YYYY-MM-DD HH:mm:ss format (UTC+8).
	//
	// example:
	//
	// 2026-10-12 10:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Indicates whether the caller has read the item.
	//
	// example:
	//
	// false
	IsRead *bool `json:"isRead,omitempty" xml:"isRead,omitempty"`
	// The message ID of the assistant that sent the to-do item. This field is empty if there is no conversation source.
	//
	// example:
	//
	// msg-1
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The action status of the caller as an assignee.
	//
	// example:
	//
	// PENDING
	MyAssigneeStatus *string `json:"myAssigneeStatus,omitempty" xml:"myAssigneeStatus,omitempty"`
	// The list display status for the caller.
	//
	// example:
	//
	// TODO
	MyDisplayState *string `json:"myDisplayState,omitempty" xml:"myDisplayState,omitempty"`
	// The priority.
	//
	// example:
	//
	// HIGH
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// The number of questions.
	//
	// example:
	//
	// 1
	QuestionCount *int64 `json:"questionCount,omitempty" xml:"questionCount,omitempty"`
	// The list of question summaries.
	QuestionSummaries []*string `json:"questionSummaries,omitempty" xml:"questionSummaries,omitempty" type:"Repeated"`
	// The source context. This context is passed through as-is during creation.
	//
	// example:
	//
	// source=openapi
	SourceContext map[string]interface{} `json:"sourceContext,omitempty" xml:"sourceContext,omitempty"`
	// The source type of the to-do item.
	//
	// example:
	//
	// WORKFLOW
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The primary status of the to-do item.
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The to-do title.
	//
	// example:
	//
	// Governance approval
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The to-do ID.
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
}

func (s ListTodosResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListTodosResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListTodosResponseBodyItems) GetAssigneeCount() *int64 {
	return s.AssigneeCount
}

func (s *ListTodosResponseBodyItems) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *ListTodosResponseBodyItems) GetCompletedByUserId() *string {
	return s.CompletedByUserId
}

func (s *ListTodosResponseBodyItems) GetCompletedByUserName() *string {
	return s.CompletedByUserName
}

func (s *ListTodosResponseBodyItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListTodosResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListTodosResponseBodyItems) GetIsRead() *bool {
	return s.IsRead
}

func (s *ListTodosResponseBodyItems) GetMessageId() *string {
	return s.MessageId
}

func (s *ListTodosResponseBodyItems) GetMyAssigneeStatus() *string {
	return s.MyAssigneeStatus
}

func (s *ListTodosResponseBodyItems) GetMyDisplayState() *string {
	return s.MyDisplayState
}

func (s *ListTodosResponseBodyItems) GetPriority() *string {
	return s.Priority
}

func (s *ListTodosResponseBodyItems) GetQuestionCount() *int64 {
	return s.QuestionCount
}

func (s *ListTodosResponseBodyItems) GetQuestionSummaries() []*string {
	return s.QuestionSummaries
}

func (s *ListTodosResponseBodyItems) GetSourceContext() map[string]interface{} {
	return s.SourceContext
}

func (s *ListTodosResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *ListTodosResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListTodosResponseBodyItems) GetTitle() *string {
	return s.Title
}

func (s *ListTodosResponseBodyItems) GetTodoId() *string {
	return s.TodoId
}

func (s *ListTodosResponseBodyItems) SetAssigneeCount(v int64) *ListTodosResponseBodyItems {
	s.AssigneeCount = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetCompletedAt(v string) *ListTodosResponseBodyItems {
	s.CompletedAt = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetCompletedByUserId(v string) *ListTodosResponseBodyItems {
	s.CompletedByUserId = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetCompletedByUserName(v string) *ListTodosResponseBodyItems {
	s.CompletedByUserName = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetGmtCreate(v string) *ListTodosResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetGmtModified(v string) *ListTodosResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetIsRead(v bool) *ListTodosResponseBodyItems {
	s.IsRead = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetMessageId(v string) *ListTodosResponseBodyItems {
	s.MessageId = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetMyAssigneeStatus(v string) *ListTodosResponseBodyItems {
	s.MyAssigneeStatus = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetMyDisplayState(v string) *ListTodosResponseBodyItems {
	s.MyDisplayState = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetPriority(v string) *ListTodosResponseBodyItems {
	s.Priority = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetQuestionCount(v int64) *ListTodosResponseBodyItems {
	s.QuestionCount = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetQuestionSummaries(v []*string) *ListTodosResponseBodyItems {
	s.QuestionSummaries = v
	return s
}

func (s *ListTodosResponseBodyItems) SetSourceContext(v map[string]interface{}) *ListTodosResponseBodyItems {
	s.SourceContext = v
	return s
}

func (s *ListTodosResponseBodyItems) SetSourceType(v string) *ListTodosResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetStatus(v string) *ListTodosResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetTitle(v string) *ListTodosResponseBodyItems {
	s.Title = &v
	return s
}

func (s *ListTodosResponseBodyItems) SetTodoId(v string) *ListTodosResponseBodyItems {
	s.TodoId = &v
	return s
}

func (s *ListTodosResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
