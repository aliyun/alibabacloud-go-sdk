// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserRatePlanInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckRemainingSiteQuota(v string) *ListUserRatePlanInstancesRequest
	GetCheckRemainingSiteQuota() *string
	SetInstanceId(v string) *ListUserRatePlanInstancesRequest
	GetInstanceId() *string
	SetIsShared(v bool) *ListUserRatePlanInstancesRequest
	GetIsShared() *bool
	SetPageNumber(v int32) *ListUserRatePlanInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserRatePlanInstancesRequest
	GetPageSize() *int32
	SetPlanNameEn(v string) *ListUserRatePlanInstancesRequest
	GetPlanNameEn() *string
	SetPlanType(v string) *ListUserRatePlanInstancesRequest
	GetPlanType() *string
	SetRemainingExpireDays(v int32) *ListUserRatePlanInstancesRequest
	GetRemainingExpireDays() *int32
	SetResourceOwner(v int64) *ListUserRatePlanInstancesRequest
	GetResourceOwner() *int64
	SetSortBy(v string) *ListUserRatePlanInstancesRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListUserRatePlanInstancesRequest
	GetSortOrder() *string
	SetStatus(v string) *ListUserRatePlanInstancesRequest
	GetStatus() *string
	SetSubscribeType(v string) *ListUserRatePlanInstancesRequest
	GetSubscribeType() *string
}

type ListUserRatePlanInstancesRequest struct {
	// Specifies whether to filter plan instances that have remaining site quotas. Valid values:
	//
	// - **true**: Filters plan instances that have remaining site quotas.
	//
	// - **false**: Queries all plan instances under the user.
	//
	// example:
	//
	// true
	CheckRemainingSiteQuota *string `json:"CheckRemainingSiteQuota,omitempty" xml:"CheckRemainingSiteQuota,omitempty"`
	// Instance ID of the plan instance. You can obtain instance ID by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// example:
	//
	// sp-xcdn-96wblslz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsShared   *bool   `json:"IsShared,omitempty" xml:"IsShared,omitempty"`
	// The page number Settings for a paging query. Default value: **1**. Valid values: **1*	- to **100000**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paging query. Valid values: 1 to 500.
	//
	// example:
	//
	// 500
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the plan in English.
	//
	// example:
	//
	// entranceplan
	PlanNameEn *string `json:"PlanNameEn,omitempty" xml:"PlanNameEn,omitempty"`
	// The type of the plan. Valid values:
	//
	// - normal: Fixed plan.
	//
	// - enterprise: Enterprise plan.
	//
	// example:
	//
	// enterprise
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	// Queries plan instances whose remaining validity period is less than or equal to the specified number of days. The value must be a positive integer. Unit: days.
	//
	// example:
	//
	// 30
	RemainingExpireDays *int32 `json:"RemainingExpireDays,omitempty" xml:"RemainingExpireDays,omitempty"`
	ResourceOwner       *int64 `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The field by which to sort the results. By default, the results are sorted by creation time. Valid values:
	//
	// - **CreateTime**: Sorts by creation time.
	//
	// - **ExpireTime**: Sorts by expiration time.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order. By default, the results are sorted in descending order. Valid values:
	//
	// - **asc**: Ascending order.
	//
	// - **desc**: Descending order.
	//
	// example:
	//
	// asc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The instance status. Valid values:
	//
	// - **online**: The instance is in normal service status.
	//
	// - **offline**: The instance has expired but is not overdue, and is in an inactive state.
	//
	// - **disable**: The instance is released.
	//
	// - **overdue**: The instance has an overdue payment.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subscription type of the plan. Valid values:
	//
	// - entranceplan: Free Edition (Chinese mainland)
	//
	// - entranceplan_intl: Free Edition (International)
	//
	// - basicplan: Basic Edition
	//
	// - standardplan: Standard Edition
	//
	// - advancedplan: Advanced Edition
	//
	// - enterpriseplan: Enterprise Edition
	//
	// example:
	//
	// basicplan
	SubscribeType *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
}

func (s ListUserRatePlanInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserRatePlanInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListUserRatePlanInstancesRequest) GetCheckRemainingSiteQuota() *string {
	return s.CheckRemainingSiteQuota
}

func (s *ListUserRatePlanInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserRatePlanInstancesRequest) GetIsShared() *bool {
	return s.IsShared
}

func (s *ListUserRatePlanInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserRatePlanInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserRatePlanInstancesRequest) GetPlanNameEn() *string {
	return s.PlanNameEn
}

func (s *ListUserRatePlanInstancesRequest) GetPlanType() *string {
	return s.PlanType
}

func (s *ListUserRatePlanInstancesRequest) GetRemainingExpireDays() *int32 {
	return s.RemainingExpireDays
}

func (s *ListUserRatePlanInstancesRequest) GetResourceOwner() *int64 {
	return s.ResourceOwner
}

func (s *ListUserRatePlanInstancesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListUserRatePlanInstancesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListUserRatePlanInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListUserRatePlanInstancesRequest) GetSubscribeType() *string {
	return s.SubscribeType
}

func (s *ListUserRatePlanInstancesRequest) SetCheckRemainingSiteQuota(v string) *ListUserRatePlanInstancesRequest {
	s.CheckRemainingSiteQuota = &v
	return s
}

func (s *ListUserRatePlanInstancesRequest) SetInstanceId(v string) *ListUserRatePlanInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserRatePlanInstancesRequest) SetIsShared(v bool) *ListUserRatePlanInstancesRequest {
	s.IsShared = &v
	return s
}

func (s *ListUserRatePlanInstancesRequest) SetPageNumber(v int32) *ListUserRatePlanInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserRatePlanInstancesRequest) SetPageSize(v int32) *ListUserRatePlanInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserRatePlanInstancesRequest) SetPlanNameEn(v string) *ListUserRatePlanInstancesRequest {
	s.PlanNameEn = &v
	return s
}

func (s *ListUserRatePlanInstancesRequest) SetPlanType(v string) *ListUserRatePlanInstancesRequest {
	s.PlanType = &v
	return s
}

func (s *ListUserRatePlanInstancesRequest) SetRemainingExpireDays(v int32) *ListUserRatePlanInstancesRequest {
	s.RemainingExpireDays = &v
	return s
}

func (s *ListUserRatePlanInstancesRequest) SetResourceOwner(v int64) *ListUserRatePlanInstancesRequest {
	s.ResourceOwner = &v
	return s
}

func (s *ListUserRatePlanInstancesRequest) SetSortBy(v string) *ListUserRatePlanInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListUserRatePlanInstancesRequest) SetSortOrder(v string) *ListUserRatePlanInstancesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListUserRatePlanInstancesRequest) SetStatus(v string) *ListUserRatePlanInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListUserRatePlanInstancesRequest) SetSubscribeType(v string) *ListUserRatePlanInstancesRequest {
	s.SubscribeType = &v
	return s
}

func (s *ListUserRatePlanInstancesRequest) Validate() error {
	return dara.Validate(s)
}
