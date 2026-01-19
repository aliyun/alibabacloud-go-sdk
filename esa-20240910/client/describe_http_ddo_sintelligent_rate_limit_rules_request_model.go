// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSIntelligentRateLimitRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeHttpDDoSIntelligentRateLimitRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHttpDDoSIntelligentRateLimitRulesRequest
	GetPageSize() *int32
	SetRuleType(v string) *DescribeHttpDDoSIntelligentRateLimitRulesRequest
	GetRuleType() *string
	SetSiteId(v int64) *DescribeHttpDDoSIntelligentRateLimitRulesRequest
	GetSiteId() *int64
}

type DescribeHttpDDoSIntelligentRateLimitRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cc
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DescribeHttpDDoSIntelligentRateLimitRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSIntelligentRateLimitRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesRequest) SetPageNumber(v int32) *DescribeHttpDDoSIntelligentRateLimitRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesRequest) SetPageSize(v int32) *DescribeHttpDDoSIntelligentRateLimitRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesRequest) SetRuleType(v string) *DescribeHttpDDoSIntelligentRateLimitRulesRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesRequest) SetSiteId(v int64) *DescribeHttpDDoSIntelligentRateLimitRulesRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesRequest) Validate() error {
	return dara.Validate(s)
}
