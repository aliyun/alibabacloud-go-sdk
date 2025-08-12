// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribeEventRuleListRequest
	GetGroupId() *string
	SetIsEnable(v bool) *DescribeEventRuleListRequest
	GetIsEnable() *bool
	SetNamePrefix(v string) *DescribeEventRuleListRequest
	GetNamePrefix() *string
	SetPageNumber(v string) *DescribeEventRuleListRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeEventRuleListRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeEventRuleListRequest
	GetRegionId() *string
}

type DescribeEventRuleListRequest struct {
	// The ID of the application group.
	//
	// example:
	//
	// 7378****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Specifies whether to enable the event-triggered alert rule. Valid values:
	//
	// - true (default)
	//
	// - false
	//
	// example:
	//
	// true
	IsEnable *bool `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	// The prefix in the name of the event-triggered alert rule.
	//
	// example:
	//
	// test
	NamePrefix *string `json:"NamePrefix,omitempty" xml:"NamePrefix,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Pages start from page 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEventRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeEventRuleListRequest) GetIsEnable() *bool {
	return s.IsEnable
}

func (s *DescribeEventRuleListRequest) GetNamePrefix() *string {
	return s.NamePrefix
}

func (s *DescribeEventRuleListRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeEventRuleListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeEventRuleListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEventRuleListRequest) SetGroupId(v string) *DescribeEventRuleListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeEventRuleListRequest) SetIsEnable(v bool) *DescribeEventRuleListRequest {
	s.IsEnable = &v
	return s
}

func (s *DescribeEventRuleListRequest) SetNamePrefix(v string) *DescribeEventRuleListRequest {
	s.NamePrefix = &v
	return s
}

func (s *DescribeEventRuleListRequest) SetPageNumber(v string) *DescribeEventRuleListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventRuleListRequest) SetPageSize(v string) *DescribeEventRuleListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventRuleListRequest) SetRegionId(v string) *DescribeEventRuleListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEventRuleListRequest) Validate() error {
	return dara.Validate(s)
}
