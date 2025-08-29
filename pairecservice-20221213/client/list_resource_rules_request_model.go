// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *ListResourceRulesRequest
	GetAll() *bool
	SetInstanceId(v string) *ListResourceRulesRequest
	GetInstanceId() *string
	SetOrder(v string) *ListResourceRulesRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListResourceRulesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListResourceRulesRequest
	GetPageSize() *int64
	SetResourceRuleId(v string) *ListResourceRulesRequest
	GetResourceRuleId() *string
	SetResourceRuleName(v string) *ListResourceRulesRequest
	GetResourceRuleName() *string
	SetSortBy(v string) *ListResourceRulesRequest
	GetSortBy() *string
}

type ListResourceRulesRequest struct {
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required.
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Order            *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber       *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceRuleId   *string `json:"ResourceRuleId,omitempty" xml:"ResourceRuleId,omitempty"`
	ResourceRuleName *string `json:"ResourceRuleName,omitempty" xml:"ResourceRuleName,omitempty"`
	SortBy           *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListResourceRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRulesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceRulesRequest) GetAll() *bool {
	return s.All
}

func (s *ListResourceRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListResourceRulesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListResourceRulesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListResourceRulesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListResourceRulesRequest) GetResourceRuleId() *string {
	return s.ResourceRuleId
}

func (s *ListResourceRulesRequest) GetResourceRuleName() *string {
	return s.ResourceRuleName
}

func (s *ListResourceRulesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListResourceRulesRequest) SetAll(v bool) *ListResourceRulesRequest {
	s.All = &v
	return s
}

func (s *ListResourceRulesRequest) SetInstanceId(v string) *ListResourceRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListResourceRulesRequest) SetOrder(v string) *ListResourceRulesRequest {
	s.Order = &v
	return s
}

func (s *ListResourceRulesRequest) SetPageNumber(v int64) *ListResourceRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceRulesRequest) SetPageSize(v int64) *ListResourceRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceRulesRequest) SetResourceRuleId(v string) *ListResourceRulesRequest {
	s.ResourceRuleId = &v
	return s
}

func (s *ListResourceRulesRequest) SetResourceRuleName(v string) *ListResourceRulesRequest {
	s.ResourceRuleName = &v
	return s
}

func (s *ListResourceRulesRequest) SetSortBy(v string) *ListResourceRulesRequest {
	s.SortBy = &v
	return s
}

func (s *ListResourceRulesRequest) Validate() error {
	return dara.Validate(s)
}
