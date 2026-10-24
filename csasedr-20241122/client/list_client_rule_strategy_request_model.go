// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientRuleStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListClientRuleStrategyRequest
	GetCurrentPage() *int32
	SetName(v string) *ListClientRuleStrategyRequest
	GetName() *string
	SetPageSize(v int32) *ListClientRuleStrategyRequest
	GetPageSize() *int32
	SetPolicyType(v int32) *ListClientRuleStrategyRequest
	GetPolicyType() *int32
}

type ListClientRuleStrategyRequest struct {
	// The current page number for the paged query with paging. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The policy name for substring matching. If this parameter is not specified, policies are not filtered by name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page for the paged query with paging. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The policy type. Valid values:
	//
	// - **0**: Software hardening.
	//
	// - **1**: Firewall.
	//
	// - **2**: Proactive defense.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	PolicyType *int32 `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListClientRuleStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClientRuleStrategyRequest) GoString() string {
	return s.String()
}

func (s *ListClientRuleStrategyRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListClientRuleStrategyRequest) GetName() *string {
	return s.Name
}

func (s *ListClientRuleStrategyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClientRuleStrategyRequest) GetPolicyType() *int32 {
	return s.PolicyType
}

func (s *ListClientRuleStrategyRequest) SetCurrentPage(v int32) *ListClientRuleStrategyRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListClientRuleStrategyRequest) SetName(v string) *ListClientRuleStrategyRequest {
	s.Name = &v
	return s
}

func (s *ListClientRuleStrategyRequest) SetPageSize(v int32) *ListClientRuleStrategyRequest {
	s.PageSize = &v
	return s
}

func (s *ListClientRuleStrategyRequest) SetPolicyType(v int32) *ListClientRuleStrategyRequest {
	s.PolicyType = &v
	return s
}

func (s *ListClientRuleStrategyRequest) Validate() error {
	return dara.Validate(s)
}
