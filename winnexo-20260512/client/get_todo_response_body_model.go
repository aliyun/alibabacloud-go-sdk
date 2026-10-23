// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTodoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTodoResponseBody
	GetCode() *string
	SetMessage(v string) *GetTodoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTodoResponseBody
	GetRequestId() *string
	SetTodo(v map[string]interface{}) *GetTodoResponseBody
	GetTodo() map[string]interface{}
	SetTodos(v []*GetTodoResponseBodyTodos) *GetTodoResponseBody
	GetTodos() []*GetTodoResponseBodyTodos
}

type GetTodoResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The single result returned when you query by todoId. This field is empty for other query types. The field set is the same as the flattened structure of getTodoDetail.
	//
	// example:
	//
	// Sample to-do key-value pair
	Todo map[string]interface{} `json:"todo,omitempty" xml:"todo,omitempty"`
	// The multiple results returned when you query by workItemId or messageId. This field is empty for todoId queries.
	Todos []*GetTodoResponseBodyTodos `json:"todos,omitempty" xml:"todos,omitempty" type:"Repeated"`
}

func (s GetTodoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTodoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTodoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTodoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTodoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTodoResponseBody) GetTodo() map[string]interface{} {
	return s.Todo
}

func (s *GetTodoResponseBody) GetTodos() []*GetTodoResponseBodyTodos {
	return s.Todos
}

func (s *GetTodoResponseBody) SetCode(v string) *GetTodoResponseBody {
	s.Code = &v
	return s
}

func (s *GetTodoResponseBody) SetMessage(v string) *GetTodoResponseBody {
	s.Message = &v
	return s
}

func (s *GetTodoResponseBody) SetRequestId(v string) *GetTodoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTodoResponseBody) SetTodo(v map[string]interface{}) *GetTodoResponseBody {
	s.Todo = v
	return s
}

func (s *GetTodoResponseBody) SetTodos(v []*GetTodoResponseBodyTodos) *GetTodoResponseBody {
	s.Todos = v
	return s
}

func (s *GetTodoResponseBody) Validate() error {
	if s.Todos != nil {
		for _, item := range s.Todos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTodoResponseBodyTodos struct {
	// The assignees.
	Assignees []*GetTodoResponseBodyTodosAssignees `json:"assignees,omitempty" xml:"assignees,omitempty" type:"Repeated"`
	// The completion time in UTC+8.
	//
	// example:
	//
	// 2026-10-12 10:00:00
	CompletedAt *string `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// The user ID of the person who completed the to-do item.
	//
	// example:
	//
	// 42
	CompletedByUserId *string `json:"completedByUserId,omitempty" xml:"completedByUserId,omitempty"`
	// The creation time in UTC+8.
	//
	// example:
	//
	// 2026-10-12 10:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time in UTC+8.
	//
	// example:
	//
	// 2026-10-12 10:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The ID of the assistant message. This field is empty if there is no conversation source.
	//
	// example:
	//
	// msg-1
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The priority.
	//
	// example:
	//
	// MEDIUM
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// The list of questions in the original JSONB format.
	//
	// example:
	//
	// q1=Sample question
	Questions map[string]interface{} `json:"questions,omitempty" xml:"questions,omitempty"`
	// The response content in the original JSONB format.
	//
	// example:
	//
	// result=Approved
	ResultPayload map[string]interface{} `json:"resultPayload,omitempty" xml:"resultPayload,omitempty"`
	// The source context passed through in its original format.
	//
	// example:
	//
	// source=openapi
	SourceContext map[string]interface{} `json:"sourceContext,omitempty" xml:"sourceContext,omitempty"`
	// The primary status of the to-do item.
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The summary of the to-do item.
	//
	// example:
	//
	// To-do item summary
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// The title of the to-do item.
	//
	// example:
	//
	// Governance approval
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The ID of the to-do item.
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
}

func (s GetTodoResponseBodyTodos) String() string {
	return dara.Prettify(s)
}

func (s GetTodoResponseBodyTodos) GoString() string {
	return s.String()
}

func (s *GetTodoResponseBodyTodos) GetAssignees() []*GetTodoResponseBodyTodosAssignees {
	return s.Assignees
}

func (s *GetTodoResponseBodyTodos) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *GetTodoResponseBodyTodos) GetCompletedByUserId() *string {
	return s.CompletedByUserId
}

func (s *GetTodoResponseBodyTodos) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetTodoResponseBodyTodos) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetTodoResponseBodyTodos) GetMessageId() *string {
	return s.MessageId
}

func (s *GetTodoResponseBodyTodos) GetPriority() *string {
	return s.Priority
}

func (s *GetTodoResponseBodyTodos) GetQuestions() map[string]interface{} {
	return s.Questions
}

func (s *GetTodoResponseBodyTodos) GetResultPayload() map[string]interface{} {
	return s.ResultPayload
}

func (s *GetTodoResponseBodyTodos) GetSourceContext() map[string]interface{} {
	return s.SourceContext
}

func (s *GetTodoResponseBodyTodos) GetStatus() *string {
	return s.Status
}

func (s *GetTodoResponseBodyTodos) GetSummary() *string {
	return s.Summary
}

func (s *GetTodoResponseBodyTodos) GetTitle() *string {
	return s.Title
}

func (s *GetTodoResponseBodyTodos) GetTodoId() *string {
	return s.TodoId
}

func (s *GetTodoResponseBodyTodos) SetAssignees(v []*GetTodoResponseBodyTodosAssignees) *GetTodoResponseBodyTodos {
	s.Assignees = v
	return s
}

func (s *GetTodoResponseBodyTodos) SetCompletedAt(v string) *GetTodoResponseBodyTodos {
	s.CompletedAt = &v
	return s
}

func (s *GetTodoResponseBodyTodos) SetCompletedByUserId(v string) *GetTodoResponseBodyTodos {
	s.CompletedByUserId = &v
	return s
}

func (s *GetTodoResponseBodyTodos) SetGmtCreate(v string) *GetTodoResponseBodyTodos {
	s.GmtCreate = &v
	return s
}

func (s *GetTodoResponseBodyTodos) SetGmtModified(v string) *GetTodoResponseBodyTodos {
	s.GmtModified = &v
	return s
}

func (s *GetTodoResponseBodyTodos) SetMessageId(v string) *GetTodoResponseBodyTodos {
	s.MessageId = &v
	return s
}

func (s *GetTodoResponseBodyTodos) SetPriority(v string) *GetTodoResponseBodyTodos {
	s.Priority = &v
	return s
}

func (s *GetTodoResponseBodyTodos) SetQuestions(v map[string]interface{}) *GetTodoResponseBodyTodos {
	s.Questions = v
	return s
}

func (s *GetTodoResponseBodyTodos) SetResultPayload(v map[string]interface{}) *GetTodoResponseBodyTodos {
	s.ResultPayload = v
	return s
}

func (s *GetTodoResponseBodyTodos) SetSourceContext(v map[string]interface{}) *GetTodoResponseBodyTodos {
	s.SourceContext = v
	return s
}

func (s *GetTodoResponseBodyTodos) SetStatus(v string) *GetTodoResponseBodyTodos {
	s.Status = &v
	return s
}

func (s *GetTodoResponseBodyTodos) SetSummary(v string) *GetTodoResponseBodyTodos {
	s.Summary = &v
	return s
}

func (s *GetTodoResponseBodyTodos) SetTitle(v string) *GetTodoResponseBodyTodos {
	s.Title = &v
	return s
}

func (s *GetTodoResponseBodyTodos) SetTodoId(v string) *GetTodoResponseBodyTodos {
	s.TodoId = &v
	return s
}

func (s *GetTodoResponseBodyTodos) Validate() error {
	if s.Assignees != nil {
		for _, item := range s.Assignees {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTodoResponseBodyTodosAssignees struct {
	// The processing time in UTC+8. This field is empty if the item is unprocessed.
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	ActedAt *string `json:"actedAt,omitempty" xml:"actedAt,omitempty"`
	// The source of the assignee. Valid values: ORIGINAL and TRANSFERRED.
	//
	// example:
	//
	// ORIGINAL
	AssigneeSource *string `json:"assigneeSource,omitempty" xml:"assigneeSource,omitempty"`
	// The user ID of the assignee.
	//
	// example:
	//
	// 42
	AssigneeUserId *string `json:"assigneeUserId,omitempty" xml:"assigneeUserId,omitempty"`
	// The read time in UTC+8. This field is empty if the item is unread.
	//
	// example:
	//
	// 2026-10-12 10:01:00
	ReadAt *string `json:"readAt,omitempty" xml:"readAt,omitempty"`
	// The action status of the assignee.
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The user ID of the person who transferred the task. This field is populated only when the task is transferred.
	//
	// example:
	//
	// 43
	TransferredFromUserId *string `json:"transferredFromUserId,omitempty" xml:"transferredFromUserId,omitempty"`
}

func (s GetTodoResponseBodyTodosAssignees) String() string {
	return dara.Prettify(s)
}

func (s GetTodoResponseBodyTodosAssignees) GoString() string {
	return s.String()
}

func (s *GetTodoResponseBodyTodosAssignees) GetActedAt() *string {
	return s.ActedAt
}

func (s *GetTodoResponseBodyTodosAssignees) GetAssigneeSource() *string {
	return s.AssigneeSource
}

func (s *GetTodoResponseBodyTodosAssignees) GetAssigneeUserId() *string {
	return s.AssigneeUserId
}

func (s *GetTodoResponseBodyTodosAssignees) GetReadAt() *string {
	return s.ReadAt
}

func (s *GetTodoResponseBodyTodosAssignees) GetStatus() *string {
	return s.Status
}

func (s *GetTodoResponseBodyTodosAssignees) GetTransferredFromUserId() *string {
	return s.TransferredFromUserId
}

func (s *GetTodoResponseBodyTodosAssignees) SetActedAt(v string) *GetTodoResponseBodyTodosAssignees {
	s.ActedAt = &v
	return s
}

func (s *GetTodoResponseBodyTodosAssignees) SetAssigneeSource(v string) *GetTodoResponseBodyTodosAssignees {
	s.AssigneeSource = &v
	return s
}

func (s *GetTodoResponseBodyTodosAssignees) SetAssigneeUserId(v string) *GetTodoResponseBodyTodosAssignees {
	s.AssigneeUserId = &v
	return s
}

func (s *GetTodoResponseBodyTodosAssignees) SetReadAt(v string) *GetTodoResponseBodyTodosAssignees {
	s.ReadAt = &v
	return s
}

func (s *GetTodoResponseBodyTodosAssignees) SetStatus(v string) *GetTodoResponseBodyTodosAssignees {
	s.Status = &v
	return s
}

func (s *GetTodoResponseBodyTodosAssignees) SetTransferredFromUserId(v string) *GetTodoResponseBodyTodosAssignees {
	s.TransferredFromUserId = &v
	return s
}

func (s *GetTodoResponseBodyTodosAssignees) Validate() error {
	return dara.Validate(s)
}
