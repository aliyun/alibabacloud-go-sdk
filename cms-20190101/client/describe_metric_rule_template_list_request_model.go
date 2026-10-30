// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleTemplateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHistory(v bool) *DescribeMetricRuleTemplateListRequest
	GetHistory() *bool
	SetKeyword(v string) *DescribeMetricRuleTemplateListRequest
	GetKeyword() *string
	SetName(v string) *DescribeMetricRuleTemplateListRequest
	GetName() *string
	SetOrder(v bool) *DescribeMetricRuleTemplateListRequest
	GetOrder() *bool
	SetOrderBy(v string) *DescribeMetricRuleTemplateListRequest
	GetOrderBy() *string
	SetPageNumber(v int64) *DescribeMetricRuleTemplateListRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeMetricRuleTemplateListRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeMetricRuleTemplateListRequest
	GetRegionId() *string
	SetTemplateId(v int64) *DescribeMetricRuleTemplateListRequest
	GetTemplateId() *int64
}

type DescribeMetricRuleTemplateListRequest struct {
	// Specifies whether to display the history of applying the alert template to application groups. Valid values:
	//
	// - true: Displays the history.
	//
	// - false: Does not display the history. This is the default value.
	//
	// example:
	//
	// false
	History *bool `json:"History,omitempty" xml:"History,omitempty"`
	// The keyword in the alert template name.
	//
	// example:
	//
	// ECS
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The name of the alert template.
	//
	// example:
	//
	// ECS_Template1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort order. Valid values:
	//
	// - true: Ascending order. This is the default value.
	//
	// - false: Descending order.
	//
	// example:
	//
	// true
	Order *bool `json:"Order,omitempty" xml:"Order,omitempty"`
	// The sort field. Valid values:
	//
	// - gmtMotified: Modification date.
	//
	// - gmtCreate: Creation date. This is the default value.
	//
	// example:
	//
	// gmtCreate
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the alert template.
	//
	// example:
	//
	// 70****
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeMetricRuleTemplateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateListRequest) GetHistory() *bool {
	return s.History
}

func (s *DescribeMetricRuleTemplateListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeMetricRuleTemplateListRequest) GetName() *string {
	return s.Name
}

func (s *DescribeMetricRuleTemplateListRequest) GetOrder() *bool {
	return s.Order
}

func (s *DescribeMetricRuleTemplateListRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeMetricRuleTemplateListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeMetricRuleTemplateListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeMetricRuleTemplateListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMetricRuleTemplateListRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeMetricRuleTemplateListRequest) SetHistory(v bool) *DescribeMetricRuleTemplateListRequest {
	s.History = &v
	return s
}

func (s *DescribeMetricRuleTemplateListRequest) SetKeyword(v string) *DescribeMetricRuleTemplateListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeMetricRuleTemplateListRequest) SetName(v string) *DescribeMetricRuleTemplateListRequest {
	s.Name = &v
	return s
}

func (s *DescribeMetricRuleTemplateListRequest) SetOrder(v bool) *DescribeMetricRuleTemplateListRequest {
	s.Order = &v
	return s
}

func (s *DescribeMetricRuleTemplateListRequest) SetOrderBy(v string) *DescribeMetricRuleTemplateListRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeMetricRuleTemplateListRequest) SetPageNumber(v int64) *DescribeMetricRuleTemplateListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMetricRuleTemplateListRequest) SetPageSize(v int64) *DescribeMetricRuleTemplateListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMetricRuleTemplateListRequest) SetRegionId(v string) *DescribeMetricRuleTemplateListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetricRuleTemplateListRequest) SetTemplateId(v int64) *DescribeMetricRuleTemplateListRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeMetricRuleTemplateListRequest) Validate() error {
	return dara.Validate(s)
}
