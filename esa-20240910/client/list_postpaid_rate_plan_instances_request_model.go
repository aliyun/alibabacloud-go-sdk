// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPostpaidRatePlanInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckRemainingSiteQuota(v string) *ListPostpaidRatePlanInstancesRequest
	GetCheckRemainingSiteQuota() *string
	SetInstanceId(v string) *ListPostpaidRatePlanInstancesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListPostpaidRatePlanInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPostpaidRatePlanInstancesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListPostpaidRatePlanInstancesRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListPostpaidRatePlanInstancesRequest
	GetSortOrder() *string
	SetStatus(v string) *ListPostpaidRatePlanInstancesRequest
	GetStatus() *string
	SetUnrelatedType(v string) *ListPostpaidRatePlanInstancesRequest
	GetUnrelatedType() *string
}

type ListPostpaidRatePlanInstancesRequest struct {
	// Specifies whether to check for remaining site quota. Valid values:
	//
	// - true: Queries instances that have remaining site quota.
	//
	// - false: Does not filter by this condition.
	//
	// example:
	//
	// true
	CheckRemainingSiteQuota *string `json:"CheckRemainingSiteQuota,omitempty" xml:"CheckRemainingSiteQuota,omitempty"`
	// The instance ID. Use this parameter to query a specific instance.
	//
	// example:
	//
	// sp-dps-xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number for paging. The value must be greater than or equal to 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for paging. Valid values: 1 to 500.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field by which to sort the results. Valid values:
	//
	// - CreateTime: sorts by creation time.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order. Valid values:
	//
	//  	- asc: ascending order
	//
	//  	- desc: descending order.
	//
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The instance status. Valid values:
	//
	//  	- online: Normal.
	//
	//  	- overdue: Overdue payment.
	//
	//  	- disable: Released.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether the instance has purchased additional bot or DDoS protection.
	//
	// example:
	//
	// bot
	UnrelatedType *string `json:"UnrelatedType,omitempty" xml:"UnrelatedType,omitempty"`
}

func (s ListPostpaidRatePlanInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPostpaidRatePlanInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListPostpaidRatePlanInstancesRequest) GetCheckRemainingSiteQuota() *string {
	return s.CheckRemainingSiteQuota
}

func (s *ListPostpaidRatePlanInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPostpaidRatePlanInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPostpaidRatePlanInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPostpaidRatePlanInstancesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListPostpaidRatePlanInstancesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListPostpaidRatePlanInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListPostpaidRatePlanInstancesRequest) GetUnrelatedType() *string {
	return s.UnrelatedType
}

func (s *ListPostpaidRatePlanInstancesRequest) SetCheckRemainingSiteQuota(v string) *ListPostpaidRatePlanInstancesRequest {
	s.CheckRemainingSiteQuota = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesRequest) SetInstanceId(v string) *ListPostpaidRatePlanInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesRequest) SetPageNumber(v int32) *ListPostpaidRatePlanInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesRequest) SetPageSize(v int32) *ListPostpaidRatePlanInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesRequest) SetSortBy(v string) *ListPostpaidRatePlanInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesRequest) SetSortOrder(v string) *ListPostpaidRatePlanInstancesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesRequest) SetStatus(v string) *ListPostpaidRatePlanInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesRequest) SetUnrelatedType(v string) *ListPostpaidRatePlanInstancesRequest {
	s.UnrelatedType = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesRequest) Validate() error {
	return dara.Validate(s)
}
