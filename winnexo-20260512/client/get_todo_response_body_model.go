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
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// todoId 查询时的单条结果（其余查询为空）；字段集同 getTodoDetail 平铺结构
	//
	// example:
	//
	// 示例待办键值对
	Todo map[string]interface{} `json:"todo,omitempty" xml:"todo,omitempty"`
	// workItemId/messageId 查询时的多条结果（todoId 查询为空）
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
	// 全部处理人
	Assignees []*GetTodoResponseBodyTodosAssignees `json:"assignees,omitempty" xml:"assignees,omitempty" type:"Repeated"`
	// 完成时间（北京时间）
	//
	// example:
	//
	// 2026-10-12 10:00:00
	CompletedAt *string `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// 完成人用户ID
	//
	// example:
	//
	// 42
	CompletedByUserId *string `json:"completedByUserId,omitempty" xml:"completedByUserId,omitempty"`
	// 创建时间（北京时间）
	//
	// example:
	//
	// 2026-10-12 10:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 最后修改时间（北京时间）
	//
	// example:
	//
	// 2026-10-12 10:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// assistant 消息ID；无会话来源为空
	//
	// example:
	//
	// msg-1
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// 优先级
	//
	// example:
	//
	// MEDIUM
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// 问题清单（JSONB 原样）
	//
	// example:
	//
	// q1=示例问题
	Questions map[string]interface{} `json:"questions,omitempty" xml:"questions,omitempty"`
	// 作答内容（JSONB 原样）
	//
	// example:
	//
	// result=同意
	ResultPayload map[string]interface{} `json:"resultPayload,omitempty" xml:"resultPayload,omitempty"`
	// 来源上下文（原样透传）
	//
	// example:
	//
	// source=openapi
	SourceContext map[string]interface{} `json:"sourceContext,omitempty" xml:"sourceContext,omitempty"`
	// 待办主状态
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 待办摘要
	//
	// example:
	//
	// 待办事项摘要
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 待办标题
	//
	// example:
	//
	// 治理审批
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 待办ID
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
	// 处理时间（北京时间）；未处理为空
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	ActedAt *string `json:"actedAt,omitempty" xml:"actedAt,omitempty"`
	// 处理人来源：ORIGINAL/TRANSFERRED
	//
	// example:
	//
	// ORIGINAL
	AssigneeSource *string `json:"assigneeSource,omitempty" xml:"assigneeSource,omitempty"`
	// 处理人用户ID
	//
	// example:
	//
	// 42
	AssigneeUserId *string `json:"assigneeUserId,omitempty" xml:"assigneeUserId,omitempty"`
	// 已读时间（北京时间）；未读为空
	//
	// example:
	//
	// 2026-10-12 10:01:00
	ReadAt *string `json:"readAt,omitempty" xml:"readAt,omitempty"`
	// 处理人行状态
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 转出人用户ID；仅转办产生
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
