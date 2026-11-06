// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemAuditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkItemAuditResponseBody
	GetCode() *string
	SetItems(v []*GetWorkItemAuditResponseBodyItems) *GetWorkItemAuditResponseBody
	GetItems() []*GetWorkItemAuditResponseBodyItems
	SetMessage(v string) *GetWorkItemAuditResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkItemAuditResponseBody
	GetRequestId() *string
}

type GetWorkItemAuditResponseBody struct {
	// 公共业务状态码
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 按事件时间倒序排列的完整工作项审计记录
	//
	// example:
	//
	// [{"eventId":"evt_example","eventType":"WORK_ITEM_UPDATED","occurredAt":"2026-11-05T08:00:00Z","status":"SUCCESS"}]
	Items []*GetWorkItemAuditResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 公共错误描述
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
}

func (s GetWorkItemAuditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemAuditResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkItemAuditResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkItemAuditResponseBody) GetItems() []*GetWorkItemAuditResponseBodyItems {
	return s.Items
}

func (s *GetWorkItemAuditResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkItemAuditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkItemAuditResponseBody) SetCode(v string) *GetWorkItemAuditResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkItemAuditResponseBody) SetItems(v []*GetWorkItemAuditResponseBodyItems) *GetWorkItemAuditResponseBody {
	s.Items = v
	return s
}

func (s *GetWorkItemAuditResponseBody) SetMessage(v string) *GetWorkItemAuditResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkItemAuditResponseBody) SetRequestId(v string) *GetWorkItemAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkItemAuditResponseBody) Validate() error {
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

type GetWorkItemAuditResponseBodyItems struct {
	// 字段变更明细列表，每项包含变更字段与前后值
	//
	// example:
	//
	// [{"field":"title","before":"旧标题","after":"新标题"}]
	Changes interface{} `json:"changes,omitempty" xml:"changes,omitempty"`
	// 事件描述
	//
	// example:
	//
	// 工作项标题调整为季度销售机会复盘
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 审计事件唯一标识
	//
	// example:
	//
	// evt_20261105_0001
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// 审计事件类型，如 WORK_ITEM_UPDATED、TODO_COMPLETED
	//
	// example:
	//
	// WORK_ITEM_UPDATED
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// 事件发生时间，ISO8601 格式
	//
	// example:
	//
	// 2026-11-05T08:00:00Z
	OccurredAt *string `json:"occurredAt,omitempty" xml:"occurredAt,omitempty"`
	// 操作人显示名
	//
	// example:
	//
	// 张三
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// 操作人用户 ID
	//
	// example:
	//
	// 10001
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// 事件涉及的待确认问题列表
	//
	// example:
	//
	// [{"question":"是否继续执行"}]
	Questions []interface{} `json:"questions,omitempty" xml:"questions,omitempty" type:"Repeated"`
	// 事件结果负载，JSON 对象；无结果时为空
	//
	// example:
	//
	// {"result":"ok"}
	ResultPayload interface{} `json:"resultPayload,omitempty" xml:"resultPayload,omitempty"`
	// 事件状态
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 事件标题
	//
	// example:
	//
	// 调查销售机会
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 关联待办 ID；无关联时为空字符串
	//
	// example:
	//
	// todo_20261105_0007
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
}

func (s GetWorkItemAuditResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemAuditResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetWorkItemAuditResponseBodyItems) GetChanges() interface{} {
	return s.Changes
}

func (s *GetWorkItemAuditResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *GetWorkItemAuditResponseBodyItems) GetEventId() *string {
	return s.EventId
}

func (s *GetWorkItemAuditResponseBodyItems) GetEventType() *string {
	return s.EventType
}

func (s *GetWorkItemAuditResponseBodyItems) GetOccurredAt() *string {
	return s.OccurredAt
}

func (s *GetWorkItemAuditResponseBodyItems) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetWorkItemAuditResponseBodyItems) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *GetWorkItemAuditResponseBodyItems) GetQuestions() []interface{} {
	return s.Questions
}

func (s *GetWorkItemAuditResponseBodyItems) GetResultPayload() interface{} {
	return s.ResultPayload
}

func (s *GetWorkItemAuditResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *GetWorkItemAuditResponseBodyItems) GetTitle() *string {
	return s.Title
}

func (s *GetWorkItemAuditResponseBodyItems) GetTodoId() *string {
	return s.TodoId
}

func (s *GetWorkItemAuditResponseBodyItems) SetChanges(v interface{}) *GetWorkItemAuditResponseBodyItems {
	s.Changes = v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetDescription(v string) *GetWorkItemAuditResponseBodyItems {
	s.Description = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetEventId(v string) *GetWorkItemAuditResponseBodyItems {
	s.EventId = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetEventType(v string) *GetWorkItemAuditResponseBodyItems {
	s.EventType = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetOccurredAt(v string) *GetWorkItemAuditResponseBodyItems {
	s.OccurredAt = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetOperatorName(v string) *GetWorkItemAuditResponseBodyItems {
	s.OperatorName = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetOperatorUserId(v string) *GetWorkItemAuditResponseBodyItems {
	s.OperatorUserId = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetQuestions(v []interface{}) *GetWorkItemAuditResponseBodyItems {
	s.Questions = v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetResultPayload(v interface{}) *GetWorkItemAuditResponseBodyItems {
	s.ResultPayload = v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetStatus(v string) *GetWorkItemAuditResponseBodyItems {
	s.Status = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetTitle(v string) *GetWorkItemAuditResponseBodyItems {
	s.Title = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) SetTodoId(v string) *GetWorkItemAuditResponseBodyItems {
	s.TodoId = &v
	return s
}

func (s *GetWorkItemAuditResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
