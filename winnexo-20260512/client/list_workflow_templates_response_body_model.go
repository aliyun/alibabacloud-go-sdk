// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWorkflowTemplatesResponseBody
	GetCode() *string
	SetHasMore(v bool) *ListWorkflowTemplatesResponseBody
	GetHasMore() *bool
	SetItems(v []*ListWorkflowTemplatesResponseBodyItems) *ListWorkflowTemplatesResponseBody
	GetItems() []*ListWorkflowTemplatesResponseBodyItems
	SetMaxResults(v int64) *ListWorkflowTemplatesResponseBody
	GetMaxResults() *int64
	SetMessage(v string) *ListWorkflowTemplatesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListWorkflowTemplatesResponseBody
	GetNextToken() *string
	SetPage(v int64) *ListWorkflowTemplatesResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListWorkflowTemplatesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListWorkflowTemplatesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListWorkflowTemplatesResponseBody
	GetTotal() *int64
}

type ListWorkflowTemplatesResponseBody struct {
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
	// 当前页模板摘要列表
	//
	// example:
	//
	// example
	Items []*ListWorkflowTemplatesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 本次生效的单页最大返回数量
	//
	// example:
	//
	// 1
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
	// 1
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

func (s ListWorkflowTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWorkflowTemplatesResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListWorkflowTemplatesResponseBody) GetItems() []*ListWorkflowTemplatesResponseBodyItems {
	return s.Items
}

func (s *ListWorkflowTemplatesResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListWorkflowTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkflowTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowTemplatesResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListWorkflowTemplatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWorkflowTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowTemplatesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListWorkflowTemplatesResponseBody) SetCode(v string) *ListWorkflowTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetHasMore(v bool) *ListWorkflowTemplatesResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetItems(v []*ListWorkflowTemplatesResponseBodyItems) *ListWorkflowTemplatesResponseBody {
	s.Items = v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetMaxResults(v int64) *ListWorkflowTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetMessage(v string) *ListWorkflowTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetNextToken(v string) *ListWorkflowTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetPage(v int64) *ListWorkflowTemplatesResponseBody {
	s.Page = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetPageSize(v int64) *ListWorkflowTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetRequestId(v string) *ListWorkflowTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetTotal(v int64) *ListWorkflowTemplatesResponseBody {
	s.Total = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) Validate() error {
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

type ListWorkflowTemplatesResponseBodyItems struct {
	// 模板描述
	//
	// example:
	//
	// example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 最近修改时间(ISO8601)
	//
	// example:
	//
	// example
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 模板名称
	//
	// example:
	//
	// example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 去重后的负责人数量
	//
	// example:
	//
	// 1
	OwnerCount *int64 `json:"ownerCount,omitempty" xml:"ownerCount,omitempty"`
	// 配置了审核人的步骤数
	//
	// example:
	//
	// 1
	ReviewStepCount *int64 `json:"reviewStepCount,omitempty" xml:"reviewStepCount,omitempty"`
	// 步骤数
	//
	// example:
	//
	// 1
	StepCount *int64 `json:"stepCount,omitempty" xml:"stepCount,omitempty"`
	// 模板ID
	//
	// example:
	//
	// example
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s ListWorkflowTemplatesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTemplatesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetOwnerCount() *int64 {
	return s.OwnerCount
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetReviewStepCount() *int64 {
	return s.ReviewStepCount
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetStepCount() *int64 {
	return s.StepCount
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetDescription(v string) *ListWorkflowTemplatesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetGmtModified(v string) *ListWorkflowTemplatesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetName(v string) *ListWorkflowTemplatesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetOwnerCount(v int64) *ListWorkflowTemplatesResponseBodyItems {
	s.OwnerCount = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetReviewStepCount(v int64) *ListWorkflowTemplatesResponseBodyItems {
	s.ReviewStepCount = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetStepCount(v int64) *ListWorkflowTemplatesResponseBodyItems {
	s.StepCount = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetTemplateId(v string) *ListWorkflowTemplatesResponseBodyItems {
	s.TemplateId = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
