// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVccGrantRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnablePage(v bool) *ListVccGrantRulesRequest
	GetEnablePage() *bool
	SetErId(v string) *ListVccGrantRulesRequest
	GetErId() *string
	SetForSelect(v bool) *ListVccGrantRulesRequest
	GetForSelect() *bool
	SetGrantRuleId(v string) *ListVccGrantRulesRequest
	GetGrantRuleId() *string
	SetGrantTenantId(v string) *ListVccGrantRulesRequest
	GetGrantTenantId() *string
	SetInstanceId(v string) *ListVccGrantRulesRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ListVccGrantRulesRequest
	GetInstanceName() *string
	SetPageNumber(v int32) *ListVccGrantRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVccGrantRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListVccGrantRulesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListVccGrantRulesRequest
	GetResourceGroupId() *string
}

type ListVccGrantRulesRequest struct {
	// Specifies whether to enable paged query. Optional values:
	//
	// 	- **true**: Enable pagination query
	//
	// 	- **false**: Pagination query is disabled
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Lingjun HUB ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Use the drop-down box
	//
	// example:
	//
	// true
	ForSelect *bool `json:"ForSelect,omitempty" xml:"ForSelect,omitempty"`
	// Authorization Entry ID
	//
	// example:
	//
	// grant-rule-jaj33d1b804
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Authorized Tenant ID
	//
	// example:
	//
	// 1620939556166277
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Network Instance ID
	//
	// example:
	//
	// vcc-cn-jaj33d1b804
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instance name
	//
	// example:
	//
	// vcc-1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-aekzb3n5lgk2ieq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListVccGrantRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVccGrantRulesRequest) GoString() string {
	return s.String()
}

func (s *ListVccGrantRulesRequest) GetEnablePage() *bool {
	return s.EnablePage
}

func (s *ListVccGrantRulesRequest) GetErId() *string {
	return s.ErId
}

func (s *ListVccGrantRulesRequest) GetForSelect() *bool {
	return s.ForSelect
}

func (s *ListVccGrantRulesRequest) GetGrantRuleId() *string {
	return s.GrantRuleId
}

func (s *ListVccGrantRulesRequest) GetGrantTenantId() *string {
	return s.GrantTenantId
}

func (s *ListVccGrantRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVccGrantRulesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListVccGrantRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVccGrantRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVccGrantRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVccGrantRulesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVccGrantRulesRequest) SetEnablePage(v bool) *ListVccGrantRulesRequest {
	s.EnablePage = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetErId(v string) *ListVccGrantRulesRequest {
	s.ErId = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetForSelect(v bool) *ListVccGrantRulesRequest {
	s.ForSelect = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetGrantRuleId(v string) *ListVccGrantRulesRequest {
	s.GrantRuleId = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetGrantTenantId(v string) *ListVccGrantRulesRequest {
	s.GrantTenantId = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetInstanceId(v string) *ListVccGrantRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetInstanceName(v string) *ListVccGrantRulesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetPageNumber(v int32) *ListVccGrantRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetPageSize(v int32) *ListVccGrantRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetRegionId(v string) *ListVccGrantRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListVccGrantRulesRequest) SetResourceGroupId(v string) *ListVccGrantRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVccGrantRulesRequest) Validate() error {
	return dara.Validate(s)
}
