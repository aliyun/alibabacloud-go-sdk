// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressionVariablePageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeExpressionVariablePageResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeExpressionVariablePageResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeExpressionVariablePageResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeExpressionVariablePageResponseBodyResultObject) *DescribeExpressionVariablePageResponseBody
	GetResultObject() []*DescribeExpressionVariablePageResponseBodyResultObject
	SetTotalItem(v int32) *DescribeExpressionVariablePageResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeExpressionVariablePageResponseBody
	GetTotalPage() *int32
}

type DescribeExpressionVariablePageResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, with a default value of 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object.
	ResultObject []*DescribeExpressionVariablePageResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items.
	//
	// example:
	//
	// 31
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 9
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeExpressionVariablePageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariablePageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariablePageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpressionVariablePageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeExpressionVariablePageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExpressionVariablePageResponseBody) GetResultObject() []*DescribeExpressionVariablePageResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeExpressionVariablePageResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeExpressionVariablePageResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeExpressionVariablePageResponseBody) SetRequestId(v string) *DescribeExpressionVariablePageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressionVariablePageResponseBody) SetCurrentPage(v int32) *DescribeExpressionVariablePageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeExpressionVariablePageResponseBody) SetPageSize(v int32) *DescribeExpressionVariablePageResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeExpressionVariablePageResponseBody) SetResultObject(v []*DescribeExpressionVariablePageResponseBodyResultObject) *DescribeExpressionVariablePageResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeExpressionVariablePageResponseBody) SetTotalItem(v int32) *DescribeExpressionVariablePageResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeExpressionVariablePageResponseBody) SetTotalPage(v int32) *DescribeExpressionVariablePageResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeExpressionVariablePageResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeExpressionVariablePageResponseBodyResultObject struct {
	// Description information.
	//
	// example:
	//
	// 变量描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Primary key of custom variable.
	//
	// example:
	//
	// 2793
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Query variable name.
	//
	// example:
	//
	// ex_OERlw0Zqfb23
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Variable return type.
	//
	// example:
	//
	// DOUBLE
	Outputs *string `json:"outputs,omitempty" xml:"outputs,omitempty"`
	// Associated Strategies.
	RuleList []*string `json:"ruleList,omitempty" xml:"ruleList,omitempty" type:"Repeated"`
	// Status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Title.
	//
	// example:
	//
	// 获取手机号前7位自定义变量
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Version information.
	//
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeExpressionVariablePageResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariablePageResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) GetOutputs() *string {
	return s.Outputs
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) GetRuleList() []*string {
	return s.RuleList
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) GetTitle() *string {
	return s.Title
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) GetVersion() *int64 {
	return s.Version
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) SetDescription(v string) *DescribeExpressionVariablePageResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) SetEventName(v string) *DescribeExpressionVariablePageResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) SetGmtModified(v int64) *DescribeExpressionVariablePageResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) SetId(v int64) *DescribeExpressionVariablePageResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) SetName(v string) *DescribeExpressionVariablePageResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) SetOutputs(v string) *DescribeExpressionVariablePageResponseBodyResultObject {
	s.Outputs = &v
	return s
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) SetRuleList(v []*string) *DescribeExpressionVariablePageResponseBodyResultObject {
	s.RuleList = v
	return s
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) SetStatus(v string) *DescribeExpressionVariablePageResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) SetTitle(v string) *DescribeExpressionVariablePageResponseBodyResultObject {
	s.Title = &v
	return s
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) SetVersion(v int64) *DescribeExpressionVariablePageResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *DescribeExpressionVariablePageResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
