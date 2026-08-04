// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupMachineGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorID(v string) *ListResourceGroupMachineGroupsRequest
	GetCreatorID() *string
	SetDiskPL(v string) *ListResourceGroupMachineGroupsRequest
	GetDiskPL() *string
	SetEcsSpec(v string) *ListResourceGroupMachineGroupsRequest
	GetEcsSpec() *string
	SetMachineGroupIDs(v string) *ListResourceGroupMachineGroupsRequest
	GetMachineGroupIDs() *string
	SetName(v string) *ListResourceGroupMachineGroupsRequest
	GetName() *string
	SetOrder(v string) *ListResourceGroupMachineGroupsRequest
	GetOrder() *string
	SetOrderInstanceId(v string) *ListResourceGroupMachineGroupsRequest
	GetOrderInstanceId() *string
	SetPageNumber(v int32) *ListResourceGroupMachineGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceGroupMachineGroupsRequest
	GetPageSize() *int32
	SetPaymentDuration(v string) *ListResourceGroupMachineGroupsRequest
	GetPaymentDuration() *string
	SetPaymentDurationUnit(v string) *ListResourceGroupMachineGroupsRequest
	GetPaymentDurationUnit() *string
	SetPaymentType(v string) *ListResourceGroupMachineGroupsRequest
	GetPaymentType() *string
	SetSortBy(v string) *ListResourceGroupMachineGroupsRequest
	GetSortBy() *string
	SetStatus(v string) *ListResourceGroupMachineGroupsRequest
	GetStatus() *string
}

type ListResourceGroupMachineGroupsRequest struct {
	// The ID of the user who created the machine group.
	//
	// example:
	//
	// 1612285282502326
	CreatorID *string `json:"CreatorID,omitempty" xml:"CreatorID,omitempty"`
	DiskPL    *string `json:"DiskPL,omitempty" xml:"DiskPL,omitempty"`
	// The ECS instance type.
	//
	// example:
	//
	// ecs.c6.large
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// The machine group IDs. Separate multiple IDs with a comma (,).
	//
	// example:
	//
	// mg105ecqwfe49hwb
	MachineGroupIDs *string `json:"MachineGroupIDs,omitempty" xml:"MachineGroupIDs,omitempty"`
	// The name of the machine group.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort order.
	//
	// - `Asc`: Sorts the results in ascending order.
	//
	// - `Desc`: Sorts the results in descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The ID of the purchase order.
	//
	// example:
	//
	// 236553689400333
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The payment duration. If `PaymentDurationUnit` is set to `Month`, valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	PaymentDuration *string `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	// The billing duration unit of the machine group.
	//
	// example:
	//
	// Month
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	// The payment type.
	//
	// example:
	//
	// PREPAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The field to sort by.
	//
	// example:
	//
	// GmtCreatedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The instance status. `Creating`: The instance is being created. `Ready`: The instance is running. `Expiring`: The instance is expiring. `Expired`: The instance has expired. `Stopping`: The instance is being stopped. `Stopped`: The instance is stopped.
	//
	// example:
	//
	// Ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListResourceGroupMachineGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupMachineGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupMachineGroupsRequest) GetCreatorID() *string {
	return s.CreatorID
}

func (s *ListResourceGroupMachineGroupsRequest) GetDiskPL() *string {
	return s.DiskPL
}

func (s *ListResourceGroupMachineGroupsRequest) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *ListResourceGroupMachineGroupsRequest) GetMachineGroupIDs() *string {
	return s.MachineGroupIDs
}

func (s *ListResourceGroupMachineGroupsRequest) GetName() *string {
	return s.Name
}

func (s *ListResourceGroupMachineGroupsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListResourceGroupMachineGroupsRequest) GetOrderInstanceId() *string {
	return s.OrderInstanceId
}

func (s *ListResourceGroupMachineGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceGroupMachineGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceGroupMachineGroupsRequest) GetPaymentDuration() *string {
	return s.PaymentDuration
}

func (s *ListResourceGroupMachineGroupsRequest) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *ListResourceGroupMachineGroupsRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListResourceGroupMachineGroupsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListResourceGroupMachineGroupsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListResourceGroupMachineGroupsRequest) SetCreatorID(v string) *ListResourceGroupMachineGroupsRequest {
	s.CreatorID = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetDiskPL(v string) *ListResourceGroupMachineGroupsRequest {
	s.DiskPL = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetEcsSpec(v string) *ListResourceGroupMachineGroupsRequest {
	s.EcsSpec = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetMachineGroupIDs(v string) *ListResourceGroupMachineGroupsRequest {
	s.MachineGroupIDs = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetName(v string) *ListResourceGroupMachineGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetOrder(v string) *ListResourceGroupMachineGroupsRequest {
	s.Order = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetOrderInstanceId(v string) *ListResourceGroupMachineGroupsRequest {
	s.OrderInstanceId = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetPageNumber(v int32) *ListResourceGroupMachineGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetPageSize(v int32) *ListResourceGroupMachineGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetPaymentDuration(v string) *ListResourceGroupMachineGroupsRequest {
	s.PaymentDuration = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetPaymentDurationUnit(v string) *ListResourceGroupMachineGroupsRequest {
	s.PaymentDurationUnit = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetPaymentType(v string) *ListResourceGroupMachineGroupsRequest {
	s.PaymentType = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetSortBy(v string) *ListResourceGroupMachineGroupsRequest {
	s.SortBy = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) SetStatus(v string) *ListResourceGroupMachineGroupsRequest {
	s.Status = &v
	return s
}

func (s *ListResourceGroupMachineGroupsRequest) Validate() error {
	return dara.Validate(s)
}
