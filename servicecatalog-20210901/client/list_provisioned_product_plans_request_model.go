// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvisionedProductPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessLevelFilter(v string) *ListProvisionedProductPlansRequest
	GetAccessLevelFilter() *string
	SetApprovalFilter(v string) *ListProvisionedProductPlansRequest
	GetApprovalFilter() *string
	SetFilters(v []*ListProvisionedProductPlansRequestFilters) *ListProvisionedProductPlansRequest
	GetFilters() []*ListProvisionedProductPlansRequestFilters
	SetPageNumber(v int32) *ListProvisionedProductPlansRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProvisionedProductPlansRequest
	GetPageSize() *int32
	SetProvisionedProductId(v string) *ListProvisionedProductPlansRequest
	GetProvisionedProductId() *string
	SetSortBy(v string) *ListProvisionedProductPlansRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListProvisionedProductPlansRequest
	GetSortOrder() *string
}

type ListProvisionedProductPlansRequest struct {
	// The access filter. Valid values:
	//
	// 	- User (default): queries the plans that are created by the current requester.
	//
	// 	- Account: queries the plans that belong to the current Alibaba Cloud account.
	//
	// 	- ResourceDirectory: queries the plans that belong to the current resource directory.
	//
	// example:
	//
	// User
	AccessLevelFilter *string `json:"AccessLevelFilter,omitempty" xml:"AccessLevelFilter,omitempty"`
	// The access filter of the review dimension. Valid values:
	//
	// 	- ReceivedRequests: queries plans that are pending for review.
	//
	// 	- ApprovalHistory: queries review history.
	//
	// 	- AccountRequests: queries all plans that belong to the current Alibaba Cloud account.
	//
	// 	- AccountRequests: queries all plans that belong to the current Alibaba Cloud account.
	//
	// example:
	//
	// ReceivedRequests
	ApprovalFilter *string `json:"ApprovalFilter,omitempty" xml:"ApprovalFilter,omitempty"`
	// An array that consists of filter conditions.
	Filters []*ListProvisionedProductPlansRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Minimum value: 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the product instance.
	//
	// example:
	//
	// pp-bp1ddg1n2a****
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The information based on which you want to sort the query results.
	//
	// Set the value to CreateTime, which specifies the creation time of plans.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// 	- Asc: the ascending order
	//
	// 	- Desc (default): the descending order.
	//
	// example:
	//
	// Desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListProvisionedProductPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductPlansRequest) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlansRequest) GetAccessLevelFilter() *string {
	return s.AccessLevelFilter
}

func (s *ListProvisionedProductPlansRequest) GetApprovalFilter() *string {
	return s.ApprovalFilter
}

func (s *ListProvisionedProductPlansRequest) GetFilters() []*ListProvisionedProductPlansRequestFilters {
	return s.Filters
}

func (s *ListProvisionedProductPlansRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProvisionedProductPlansRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProvisionedProductPlansRequest) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *ListProvisionedProductPlansRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListProvisionedProductPlansRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListProvisionedProductPlansRequest) SetAccessLevelFilter(v string) *ListProvisionedProductPlansRequest {
	s.AccessLevelFilter = &v
	return s
}

func (s *ListProvisionedProductPlansRequest) SetApprovalFilter(v string) *ListProvisionedProductPlansRequest {
	s.ApprovalFilter = &v
	return s
}

func (s *ListProvisionedProductPlansRequest) SetFilters(v []*ListProvisionedProductPlansRequestFilters) *ListProvisionedProductPlansRequest {
	s.Filters = v
	return s
}

func (s *ListProvisionedProductPlansRequest) SetPageNumber(v int32) *ListProvisionedProductPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProvisionedProductPlansRequest) SetPageSize(v int32) *ListProvisionedProductPlansRequest {
	s.PageSize = &v
	return s
}

func (s *ListProvisionedProductPlansRequest) SetProvisionedProductId(v string) *ListProvisionedProductPlansRequest {
	s.ProvisionedProductId = &v
	return s
}

func (s *ListProvisionedProductPlansRequest) SetSortBy(v string) *ListProvisionedProductPlansRequest {
	s.SortBy = &v
	return s
}

func (s *ListProvisionedProductPlansRequest) SetSortOrder(v string) *ListProvisionedProductPlansRequest {
	s.SortOrder = &v
	return s
}

func (s *ListProvisionedProductPlansRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProvisionedProductPlansRequestFilters struct {
	// The name of the filter condition. Valid values:
	//
	// 	- ProvisionedProductPlanName: performs exact matches by plan name. Plan names are not case-sensitive.
	//
	// 	- ProvisionedProductPlanApprover: performs exact matches by reviewer. You must specify a reviewer in the `RamUser/RamRole:<Name of the reviewer>` format. You can specify multiple reviewers.
	//
	// 	- ProvisionedProductPlanApproverName: performs matches by reviewer name. You must specify the Resource Access Management (RAM) entity name of the reviewer. You can specify multiple reviewer names.
	//
	// 	- ProvisionedProductPlanStatus: performs matches by plan status. You must specify the state of the plan. You can specify multiple states.
	//
	// 	- ProvisionedProductPlanOwnerUid: performs exact matches by ID of Alibaba Cloud account to which a plan belongs.
	//
	// 	- FullTextSearch: performs fuzzy full-text searches by plan name.
	//
	// example:
	//
	// FullTextSearch
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	//
	// example:
	//
	// ECS
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProvisionedProductPlansRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductPlansRequestFilters) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlansRequestFilters) GetKey() *string {
	return s.Key
}

func (s *ListProvisionedProductPlansRequestFilters) GetValue() *string {
	return s.Value
}

func (s *ListProvisionedProductPlansRequestFilters) SetKey(v string) *ListProvisionedProductPlansRequestFilters {
	s.Key = &v
	return s
}

func (s *ListProvisionedProductPlansRequestFilters) SetValue(v string) *ListProvisionedProductPlansRequestFilters {
	s.Value = &v
	return s
}

func (s *ListProvisionedProductPlansRequestFilters) Validate() error {
	return dara.Validate(s)
}
