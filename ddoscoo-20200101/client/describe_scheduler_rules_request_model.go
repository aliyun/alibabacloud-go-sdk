// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSchedulerRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSchedulerRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSchedulerRulesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeSchedulerRulesRequest
	GetResourceGroupId() *string
	SetRuleName(v string) *DescribeSchedulerRulesRequest
	GetRuleName() *string
}

type DescribeSchedulerRulesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// testrule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeSchedulerRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSchedulerRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSchedulerRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSchedulerRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSchedulerRulesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSchedulerRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeSchedulerRulesRequest) SetPageNumber(v int32) *DescribeSchedulerRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSchedulerRulesRequest) SetPageSize(v int32) *DescribeSchedulerRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSchedulerRulesRequest) SetResourceGroupId(v string) *DescribeSchedulerRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSchedulerRulesRequest) SetRuleName(v string) *DescribeSchedulerRulesRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeSchedulerRulesRequest) Validate() error {
	return dara.Validate(s)
}
