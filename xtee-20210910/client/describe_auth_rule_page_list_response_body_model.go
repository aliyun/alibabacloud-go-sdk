// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthRulePageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAuthRulePageListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int64) *DescribeAuthRulePageListResponseBody
	GetCurrentPage() *int64
	SetPageSize(v int64) *DescribeAuthRulePageListResponseBody
	GetPageSize() *int64
	SetResultObject(v []*DescribeAuthRulePageListResponseBodyResultObject) *DescribeAuthRulePageListResponseBody
	GetResultObject() []*DescribeAuthRulePageListResponseBodyResultObject
	SetTotalItem(v int64) *DescribeAuthRulePageListResponseBody
	GetTotalItem() *int64
	SetTotalPage(v int64) *DescribeAuthRulePageListResponseBody
	GetTotalPage() *int64
}

type DescribeAuthRulePageListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The returned object.
	ResultObject []*DescribeAuthRulePageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 6
	TotalItem *int64 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 9
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeAuthRulePageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthRulePageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuthRulePageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuthRulePageListResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeAuthRulePageListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAuthRulePageListResponseBody) GetResultObject() []*DescribeAuthRulePageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeAuthRulePageListResponseBody) GetTotalItem() *int64 {
	return s.TotalItem
}

func (s *DescribeAuthRulePageListResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeAuthRulePageListResponseBody) SetRequestId(v string) *DescribeAuthRulePageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBody) SetCurrentPage(v int64) *DescribeAuthRulePageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBody) SetPageSize(v int64) *DescribeAuthRulePageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBody) SetResultObject(v []*DescribeAuthRulePageListResponseBodyResultObject) *DescribeAuthRulePageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeAuthRulePageListResponseBody) SetTotalItem(v int64) *DescribeAuthRulePageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBody) SetTotalPage(v int64) *DescribeAuthRulePageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBody) Validate() error {
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

type DescribeAuthRulePageListResponseBodyResultObject struct {
	// The console rule ID.
	//
	// example:
	//
	// 6715
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// The creation type.
	//
	// example:
	//
	// MORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The primary key ID of the policy.
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The remarks.
	//
	// example:
	//
	// 分析中心事件测试_策略01
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// 102059
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The policy name.
	//
	// example:
	//
	// 营销风险识别
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// The primary key ID of the policy version.
	//
	// example:
	//
	// 3823
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
	// The status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The version number.
	//
	// example:
	//
	// 8
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeAuthRulePageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthRulePageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) GetMemo() *string {
	return s.Memo
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) GetVersion() *int64 {
	return s.Version
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) SetConsoleRuleId(v int64) *DescribeAuthRulePageListResponseBodyResultObject {
	s.ConsoleRuleId = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) SetCreateType(v string) *DescribeAuthRulePageListResponseBodyResultObject {
	s.CreateType = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) SetGmtModified(v int64) *DescribeAuthRulePageListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) SetId(v int64) *DescribeAuthRulePageListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) SetMemo(v string) *DescribeAuthRulePageListResponseBodyResultObject {
	s.Memo = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) SetRuleId(v string) *DescribeAuthRulePageListResponseBodyResultObject {
	s.RuleId = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) SetRuleName(v string) *DescribeAuthRulePageListResponseBodyResultObject {
	s.RuleName = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) SetRuleVersionId(v int64) *DescribeAuthRulePageListResponseBodyResultObject {
	s.RuleVersionId = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) SetStatus(v string) *DescribeAuthRulePageListResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) SetVersion(v int64) *DescribeAuthRulePageListResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *DescribeAuthRulePageListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
