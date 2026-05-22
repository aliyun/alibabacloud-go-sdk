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
	CheckRemainingSiteQuota *string `json:"CheckRemainingSiteQuota,omitempty" xml:"CheckRemainingSiteQuota,omitempty"`
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber              *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlanNameEn              *string `json:"PlanNameEn,omitempty" xml:"PlanNameEn,omitempty"`
	PlanType                *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	RemainingExpireDays     *int32  `json:"RemainingExpireDays,omitempty" xml:"RemainingExpireDays,omitempty"`
	SortBy                  *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortOrder               *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// if can be null:
	// false
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
