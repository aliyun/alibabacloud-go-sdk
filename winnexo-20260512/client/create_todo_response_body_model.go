// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTodoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssignees(v []*CreateTodoResponseBodyAssignees) *CreateTodoResponseBody
	GetAssignees() []*CreateTodoResponseBodyAssignees
	SetCode(v string) *CreateTodoResponseBody
	GetCode() *string
	SetGmtCreate(v string) *CreateTodoResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreateTodoResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTodoResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateTodoResponseBody
	GetStatus() *string
	SetTodoId(v string) *CreateTodoResponseBody
	GetTodoId() *string
}

type CreateTodoResponseBody struct {
	// 处理人列表
	Assignees []*CreateTodoResponseBodyAssignees `json:"assignees,omitempty" xml:"assignees,omitempty" type:"Repeated"`
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 创建时间（ISO8601 带时区；与列表/详情的北京时间格式串不同）
	//
	// example:
	//
	// 2026-10-12T15:00:00+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
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
	// 待办状态
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 新建待办ID
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
}

func (s CreateTodoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTodoResponseBody) GetAssignees() []*CreateTodoResponseBodyAssignees {
	return s.Assignees
}

func (s *CreateTodoResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTodoResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateTodoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTodoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTodoResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateTodoResponseBody) GetTodoId() *string {
	return s.TodoId
}

func (s *CreateTodoResponseBody) SetAssignees(v []*CreateTodoResponseBodyAssignees) *CreateTodoResponseBody {
	s.Assignees = v
	return s
}

func (s *CreateTodoResponseBody) SetCode(v string) *CreateTodoResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTodoResponseBody) SetGmtCreate(v string) *CreateTodoResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateTodoResponseBody) SetMessage(v string) *CreateTodoResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTodoResponseBody) SetRequestId(v string) *CreateTodoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTodoResponseBody) SetStatus(v string) *CreateTodoResponseBody {
	s.Status = &v
	return s
}

func (s *CreateTodoResponseBody) SetTodoId(v string) *CreateTodoResponseBody {
	s.TodoId = &v
	return s
}

func (s *CreateTodoResponseBody) Validate() error {
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

type CreateTodoResponseBodyAssignees struct {
	// 处理人用户ID
	//
	// example:
	//
	// 43
	AssigneeUserId *string `json:"assigneeUserId,omitempty" xml:"assigneeUserId,omitempty"`
	// 处理人行状态
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateTodoResponseBodyAssignees) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoResponseBodyAssignees) GoString() string {
	return s.String()
}

func (s *CreateTodoResponseBodyAssignees) GetAssigneeUserId() *string {
	return s.AssigneeUserId
}

func (s *CreateTodoResponseBodyAssignees) GetStatus() *string {
	return s.Status
}

func (s *CreateTodoResponseBodyAssignees) SetAssigneeUserId(v string) *CreateTodoResponseBodyAssignees {
	s.AssigneeUserId = &v
	return s
}

func (s *CreateTodoResponseBodyAssignees) SetStatus(v string) *CreateTodoResponseBodyAssignees {
	s.Status = &v
	return s
}

func (s *CreateTodoResponseBodyAssignees) Validate() error {
	return dara.Validate(s)
}
