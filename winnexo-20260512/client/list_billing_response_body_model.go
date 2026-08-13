// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBillingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBillingResponseBody
	GetCode() *string
	SetList(v []*ListBillingResponseBodyList) *ListBillingResponseBody
	GetList() []*ListBillingResponseBodyList
	SetMessage(v string) *ListBillingResponseBody
	GetMessage() *string
	SetPage(v int64) *ListBillingResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListBillingResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListBillingResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListBillingResponseBody
	GetTotal() *int64
}

type ListBillingResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string                        `json:"code,omitempty" xml:"code,omitempty"`
	List []*ListBillingResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 页码
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数
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
	// 总数
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListBillingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBillingResponseBody) GoString() string {
	return s.String()
}

func (s *ListBillingResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBillingResponseBody) GetList() []*ListBillingResponseBodyList {
	return s.List
}

func (s *ListBillingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBillingResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListBillingResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListBillingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBillingResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListBillingResponseBody) SetCode(v string) *ListBillingResponseBody {
	s.Code = &v
	return s
}

func (s *ListBillingResponseBody) SetList(v []*ListBillingResponseBodyList) *ListBillingResponseBody {
	s.List = v
	return s
}

func (s *ListBillingResponseBody) SetMessage(v string) *ListBillingResponseBody {
	s.Message = &v
	return s
}

func (s *ListBillingResponseBody) SetPage(v int64) *ListBillingResponseBody {
	s.Page = &v
	return s
}

func (s *ListBillingResponseBody) SetPageSize(v int64) *ListBillingResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBillingResponseBody) SetRequestId(v string) *ListBillingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBillingResponseBody) SetTotal(v int64) *ListBillingResponseBody {
	s.Total = &v
	return s
}

func (s *ListBillingResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBillingResponseBodyList struct {
	// 账单业务ID
	//
	// example:
	//
	// exampleBillingId
	BillingId *string `json:"billingId,omitempty" xml:"billingId,omitempty"`
	// 业务来源ID
	//
	// example:
	//
	// exampleBizId
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// 业务来源类型
	//
	// example:
	//
	// string_value
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// costSource
	//
	// example:
	//
	// string_value
	CostSource []*string `json:"costSource,omitempty" xml:"costSource,omitempty" type:"Repeated"`
	// costSourceDisplayName
	//
	// example:
	//
	// string_value
	CostSourceDisplayName []*string `json:"costSourceDisplayName,omitempty" xml:"costSourceDisplayName,omitempty" type:"Repeated"`
	// 结束时间
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 入口对象ID
	//
	// example:
	//
	// exampleEntryObjectId
	EntryObjectId *string `json:"entryObjectId,omitempty" xml:"entryObjectId,omitempty"`
	// 入口对象类型
	//
	// example:
	//
	// string_value
	EntryObjectType *string `json:"entryObjectType,omitempty" xml:"entryObjectType,omitempty"`
	// 是否影子账单
	//
	// example:
	//
	// true
	IsShadow *bool `json:"isShadow,omitempty" xml:"isShadow,omitempty"`
	// 操作类型
	//
	// example:
	//
	// string_value
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	// 操作类型展示名称
	//
	// example:
	//
	// string_value
	OperationDisplayName *string `json:"operationDisplayName,omitempty" xml:"operationDisplayName,omitempty"`
	// 开始时间
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 状态
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 状态展示名称
	//
	// example:
	//
	// string_value
	StatusDisplayName *string `json:"statusDisplayName,omitempty" xml:"statusDisplayName,omitempty"`
	// 租户ID
	//
	// example:
	//
	// 10000
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 汇总 credit 消耗
	//
	// example:
	//
	// string_value
	TotalCreditCost *string `json:"totalCreditCost,omitempty" xml:"totalCreditCost,omitempty"`
	// WINNEXO 平台用户ID
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s ListBillingResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListBillingResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListBillingResponseBodyList) GetBillingId() *string {
	return s.BillingId
}

func (s *ListBillingResponseBodyList) GetBizId() *string {
	return s.BizId
}

func (s *ListBillingResponseBodyList) GetBizType() *string {
	return s.BizType
}

func (s *ListBillingResponseBodyList) GetCostSource() []*string {
	return s.CostSource
}

func (s *ListBillingResponseBodyList) GetCostSourceDisplayName() []*string {
	return s.CostSourceDisplayName
}

func (s *ListBillingResponseBodyList) GetEndTime() *string {
	return s.EndTime
}

func (s *ListBillingResponseBodyList) GetEntryObjectId() *string {
	return s.EntryObjectId
}

func (s *ListBillingResponseBodyList) GetEntryObjectType() *string {
	return s.EntryObjectType
}

func (s *ListBillingResponseBodyList) GetIsShadow() *bool {
	return s.IsShadow
}

func (s *ListBillingResponseBodyList) GetOperation() *string {
	return s.Operation
}

func (s *ListBillingResponseBodyList) GetOperationDisplayName() *string {
	return s.OperationDisplayName
}

func (s *ListBillingResponseBodyList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListBillingResponseBodyList) GetStatus() *string {
	return s.Status
}

func (s *ListBillingResponseBodyList) GetStatusDisplayName() *string {
	return s.StatusDisplayName
}

func (s *ListBillingResponseBodyList) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListBillingResponseBodyList) GetTotalCreditCost() *string {
	return s.TotalCreditCost
}

func (s *ListBillingResponseBodyList) GetWnUserId() *string {
	return s.WnUserId
}

func (s *ListBillingResponseBodyList) SetBillingId(v string) *ListBillingResponseBodyList {
	s.BillingId = &v
	return s
}

func (s *ListBillingResponseBodyList) SetBizId(v string) *ListBillingResponseBodyList {
	s.BizId = &v
	return s
}

func (s *ListBillingResponseBodyList) SetBizType(v string) *ListBillingResponseBodyList {
	s.BizType = &v
	return s
}

func (s *ListBillingResponseBodyList) SetCostSource(v []*string) *ListBillingResponseBodyList {
	s.CostSource = v
	return s
}

func (s *ListBillingResponseBodyList) SetCostSourceDisplayName(v []*string) *ListBillingResponseBodyList {
	s.CostSourceDisplayName = v
	return s
}

func (s *ListBillingResponseBodyList) SetEndTime(v string) *ListBillingResponseBodyList {
	s.EndTime = &v
	return s
}

func (s *ListBillingResponseBodyList) SetEntryObjectId(v string) *ListBillingResponseBodyList {
	s.EntryObjectId = &v
	return s
}

func (s *ListBillingResponseBodyList) SetEntryObjectType(v string) *ListBillingResponseBodyList {
	s.EntryObjectType = &v
	return s
}

func (s *ListBillingResponseBodyList) SetIsShadow(v bool) *ListBillingResponseBodyList {
	s.IsShadow = &v
	return s
}

func (s *ListBillingResponseBodyList) SetOperation(v string) *ListBillingResponseBodyList {
	s.Operation = &v
	return s
}

func (s *ListBillingResponseBodyList) SetOperationDisplayName(v string) *ListBillingResponseBodyList {
	s.OperationDisplayName = &v
	return s
}

func (s *ListBillingResponseBodyList) SetStartTime(v string) *ListBillingResponseBodyList {
	s.StartTime = &v
	return s
}

func (s *ListBillingResponseBodyList) SetStatus(v string) *ListBillingResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListBillingResponseBodyList) SetStatusDisplayName(v string) *ListBillingResponseBodyList {
	s.StatusDisplayName = &v
	return s
}

func (s *ListBillingResponseBodyList) SetTenantId(v int64) *ListBillingResponseBodyList {
	s.TenantId = &v
	return s
}

func (s *ListBillingResponseBodyList) SetTotalCreditCost(v string) *ListBillingResponseBodyList {
	s.TotalCreditCost = &v
	return s
}

func (s *ListBillingResponseBodyList) SetWnUserId(v string) *ListBillingResponseBodyList {
	s.WnUserId = &v
	return s
}

func (s *ListBillingResponseBodyList) Validate() error {
	return dara.Validate(s)
}
