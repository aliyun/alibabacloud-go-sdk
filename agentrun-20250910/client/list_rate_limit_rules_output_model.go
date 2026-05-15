// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRateLimitRulesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*RateLimitRule) *ListRateLimitRulesOutput
	GetItems() []*RateLimitRule
	SetPageNumber(v int) *ListRateLimitRulesOutput
	GetPageNumber() *int
	SetPageSize(v int) *ListRateLimitRulesOutput
	GetPageSize() *int
	SetTotal(v int) *ListRateLimitRulesOutput
	GetTotal() *int
}

type ListRateLimitRulesOutput struct {
	// 限流规则列表
	Items []*RateLimitRule `json:"items" xml:"items" type:"Repeated"`
	// 当前页码
	//
	// example:
	//
	// 1
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页记录数
	//
	// example:
	//
	// 20
	PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 符合条件的限流规则总数
	//
	// example:
	//
	// 10
	Total *int `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListRateLimitRulesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListRateLimitRulesOutput) GoString() string {
	return s.String()
}

func (s *ListRateLimitRulesOutput) GetItems() []*RateLimitRule {
	return s.Items
}

func (s *ListRateLimitRulesOutput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListRateLimitRulesOutput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListRateLimitRulesOutput) GetTotal() *int {
	return s.Total
}

func (s *ListRateLimitRulesOutput) SetItems(v []*RateLimitRule) *ListRateLimitRulesOutput {
	s.Items = v
	return s
}

func (s *ListRateLimitRulesOutput) SetPageNumber(v int) *ListRateLimitRulesOutput {
	s.PageNumber = &v
	return s
}

func (s *ListRateLimitRulesOutput) SetPageSize(v int) *ListRateLimitRulesOutput {
	s.PageSize = &v
	return s
}

func (s *ListRateLimitRulesOutput) SetTotal(v int) *ListRateLimitRulesOutput {
	s.Total = &v
	return s
}

func (s *ListRateLimitRulesOutput) Validate() error {
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
