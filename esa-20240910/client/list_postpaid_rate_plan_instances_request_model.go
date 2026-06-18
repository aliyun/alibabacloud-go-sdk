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
	// Specifies whether to return only instances that have remaining site quota. Valid values:
	//
	// - `true`: Returns only instances with remaining site quota.
	//
	// - `false`: Returns all instances, regardless of site quota.
	//
	// example:
	//
	// true
	CheckRemainingSiteQuota *string `json:"CheckRemainingSiteQuota,omitempty" xml:"CheckRemainingSiteQuota,omitempty"`
	// The ID of the instance to query.
	//
	// example:
	//
	// sp-dps-xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. The value must be greater than or equal to 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The maximum value is 500.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field to sort the results by. Valid value:
	//
	// - `CreateTime`: Sorts by creation time.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order. Valid values:
	//
	// - `asc`: ascending
	//
	// - `desc`: descending
	//
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The instance status. Valid values:
	//
	// - `online`: The instance is running.
	//
	// - `overdue`: The payment for the instance is overdue.
	//
	// - `disable`: The instance is released.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of add-on service to filter by, such as `bot` or `ddos`.
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
