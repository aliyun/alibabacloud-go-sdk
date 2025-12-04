// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpdGrantRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnablePage(v bool) *ListVpdGrantRulesRequest
	GetEnablePage() *bool
	SetErId(v string) *ListVpdGrantRulesRequest
	GetErId() *string
	SetForSelect(v bool) *ListVpdGrantRulesRequest
	GetForSelect() *bool
	SetGrantRuleId(v string) *ListVpdGrantRulesRequest
	GetGrantRuleId() *string
	SetGrantTenantId(v string) *ListVpdGrantRulesRequest
	GetGrantTenantId() *string
	SetInstanceId(v string) *ListVpdGrantRulesRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ListVpdGrantRulesRequest
	GetInstanceName() *string
	SetPageNumber(v int32) *ListVpdGrantRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVpdGrantRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListVpdGrantRulesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListVpdGrantRulesRequest
	GetResourceGroupId() *string
}

type ListVpdGrantRulesRequest struct {
	// Specifies whether to enable pagination query.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Lingjun HUB Instance ID
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
	// grant-rule-8rgvqazb
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Authorized Tenant ID
	//
	// example:
	//
	// 1620939556166279
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// The ID of the network instance that you want to query.
	//
	// example:
	//
	// vpd-8rgvqazb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instance name
	//
	// example:
	//
	// vpd-1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
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
	// rg-aek2l4sq6l7unhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListVpdGrantRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpdGrantRulesRequest) GoString() string {
	return s.String()
}

func (s *ListVpdGrantRulesRequest) GetEnablePage() *bool {
	return s.EnablePage
}

func (s *ListVpdGrantRulesRequest) GetErId() *string {
	return s.ErId
}

func (s *ListVpdGrantRulesRequest) GetForSelect() *bool {
	return s.ForSelect
}

func (s *ListVpdGrantRulesRequest) GetGrantRuleId() *string {
	return s.GrantRuleId
}

func (s *ListVpdGrantRulesRequest) GetGrantTenantId() *string {
	return s.GrantTenantId
}

func (s *ListVpdGrantRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVpdGrantRulesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListVpdGrantRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVpdGrantRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVpdGrantRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpdGrantRulesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpdGrantRulesRequest) SetEnablePage(v bool) *ListVpdGrantRulesRequest {
	s.EnablePage = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetErId(v string) *ListVpdGrantRulesRequest {
	s.ErId = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetForSelect(v bool) *ListVpdGrantRulesRequest {
	s.ForSelect = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetGrantRuleId(v string) *ListVpdGrantRulesRequest {
	s.GrantRuleId = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetGrantTenantId(v string) *ListVpdGrantRulesRequest {
	s.GrantTenantId = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetInstanceId(v string) *ListVpdGrantRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetInstanceName(v string) *ListVpdGrantRulesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetPageNumber(v int32) *ListVpdGrantRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetPageSize(v int32) *ListVpdGrantRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetRegionId(v string) *ListVpdGrantRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpdGrantRulesRequest) SetResourceGroupId(v string) *ListVpdGrantRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpdGrantRulesRequest) Validate() error {
	return dara.Validate(s)
}
