// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSIntelligentAclRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeHttpDDoSIntelligentAclRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHttpDDoSIntelligentAclRulesRequest
	GetPageSize() *int32
	SetRuleType(v string) *DescribeHttpDDoSIntelligentAclRulesRequest
	GetRuleType() *string
	SetSiteId(v int64) *DescribeHttpDDoSIntelligentAclRulesRequest
	GetSiteId() *int64
}

type DescribeHttpDDoSIntelligentAclRulesRequest struct {
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
	// acl
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DescribeHttpDDoSIntelligentAclRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSIntelligentAclRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSIntelligentAclRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHttpDDoSIntelligentAclRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHttpDDoSIntelligentAclRulesRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeHttpDDoSIntelligentAclRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeHttpDDoSIntelligentAclRulesRequest) SetPageNumber(v int32) *DescribeHttpDDoSIntelligentAclRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesRequest) SetPageSize(v int32) *DescribeHttpDDoSIntelligentAclRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesRequest) SetRuleType(v string) *DescribeHttpDDoSIntelligentAclRulesRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesRequest) SetSiteId(v int64) *DescribeHttpDDoSIntelligentAclRulesRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesRequest) Validate() error {
	return dara.Validate(s)
}
