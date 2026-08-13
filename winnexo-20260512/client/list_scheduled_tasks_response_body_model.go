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
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否有更多数据
	//
	// example:
	//
	// true
	HasMore *bool                                  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Items   []*ListScheduledTasksResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 本次实际生效的单页最大返回数量
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 下一页翻页令牌，原样回传即可取下一页；无更多数据时为空字符串
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 当前页码（实际生效值）
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数（实际生效值）
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 满足条件的总数
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
	// 创建人
	//
	// example:
	//
	// string_value
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// Cron 表达式
	//
	// example:
	//
	// string_value
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// 任务简述
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 累计执行次数
	//
	// example:
	//
	// 1
	ExecutionCount *int64 `json:"executionCount,omitempty" xml:"executionCount,omitempty"`
	// 创建时间 ISO8601
	//
	// example:
	//
	// string_value
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间 ISO8601
	//
	// example:
	//
	// string_value
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 是否公开
	//
	// example:
	//
	// true
	IsOpen *bool `json:"isOpen,omitempty" xml:"isOpen,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 任务 ID
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 触发类型（manual/cron/event）
	//
	// example:
	//
	// string_value
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s ListScheduledTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksResponseBodyItems) GetCreator() *string {
	return s.Creator
}

func (s *ListScheduledTasksResponseBodyItems) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ListScheduledTasksResponseBodyItems) GetDescription() *string {
	return s.Description
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

func (s *ListScheduledTasksResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListScheduledTasksResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *ListScheduledTasksResponseBodyItems) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListScheduledTasksResponseBodyItems) SetCreator(v string) *ListScheduledTasksResponseBodyItems {
	s.Creator = &v
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

func (s *ListScheduledTasksResponseBodyItems) SetName(v string) *ListScheduledTasksResponseBodyItems {
	s.Name = &v
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

func (s *ListScheduledTasksResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
