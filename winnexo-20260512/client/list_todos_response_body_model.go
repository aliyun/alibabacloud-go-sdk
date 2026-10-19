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
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否还有更多数据
	//
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 当前页待办列表
	Items []*ListTodosResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 本次生效的单页最大返回数量
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 下一页翻页令牌；无更多数据时为空字符串
	//
	// example:
	//
	// example
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 当前页码
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 生效单页条数
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 筛选条件下可见总数
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
	// 处理人数量
	//
	// example:
	//
	// 2
	AssigneeCount *int64 `json:"assigneeCount,omitempty" xml:"assigneeCount,omitempty"`
	// 完成时间（北京时间 YYYY-MM-DD HH:mm:ss）
	//
	// example:
	//
	// 2026-10-12 10:00:00
	CompletedAt *string `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// 完成人用户ID；未完成为空
	//
	// example:
	//
	// 42
	CompletedByUserId *string `json:"completedByUserId,omitempty" xml:"completedByUserId,omitempty"`
	// 完成人姓名；未完成为空
	//
	// example:
	//
	// 张三
	CompletedByUserName *string `json:"completedByUserName,omitempty" xml:"completedByUserName,omitempty"`
	// 创建时间（北京时间 YYYY-MM-DD HH:mm:ss）
	//
	// example:
	//
	// 2026-10-12 10:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 最后修改时间（北京时间 YYYY-MM-DD HH:mm:ss）
	//
	// example:
	//
	// 2026-10-12 10:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 调用者是否已读
	//
	// example:
	//
	// false
	IsRead *bool `json:"isRead,omitempty" xml:"isRead,omitempty"`
	// 发出该待办的 assistant 消息ID；无会话来源为空
	//
	// example:
	//
	// msg-1
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// 调用者的处理人行状态
	//
	// example:
	//
	// PENDING
	MyAssigneeStatus *string `json:"myAssigneeStatus,omitempty" xml:"myAssigneeStatus,omitempty"`
	// 调用者的列表展示状态
	//
	// example:
	//
	// TODO
	MyDisplayState *string `json:"myDisplayState,omitempty" xml:"myDisplayState,omitempty"`
	// 优先级
	//
	// example:
	//
	// HIGH
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// 问题数量
	//
	// example:
	//
	// 1
	QuestionCount *int64 `json:"questionCount,omitempty" xml:"questionCount,omitempty"`
	// 问题摘要列表
	QuestionSummaries []*string `json:"questionSummaries,omitempty" xml:"questionSummaries,omitempty" type:"Repeated"`
	// 来源上下文（创建时原样透传）
	//
	// example:
	//
	// source=openapi
	SourceContext map[string]interface{} `json:"sourceContext,omitempty" xml:"sourceContext,omitempty"`
	// 待办来源类型
	//
	// example:
	//
	// WORKFLOW
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// 待办主状态
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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
