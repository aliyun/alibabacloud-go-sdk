// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeEipAddressesRequestFilter) *DescribeEipAddressesRequest
	GetFilter() []*DescribeEipAddressesRequestFilter
	SetAllocationId(v string) *DescribeEipAddressesRequest
	GetAllocationId() *string
	SetAssociatedInstanceId(v string) *DescribeEipAddressesRequest
	GetAssociatedInstanceId() *string
	SetAssociatedInstanceType(v string) *DescribeEipAddressesRequest
	GetAssociatedInstanceType() *string
	SetChargeType(v string) *DescribeEipAddressesRequest
	GetChargeType() *string
	SetEipAddress(v string) *DescribeEipAddressesRequest
	GetEipAddress() *string
	SetISP(v string) *DescribeEipAddressesRequest
	GetISP() *string
	SetLockReason(v string) *DescribeEipAddressesRequest
	GetLockReason() *string
	SetOwnerAccount(v string) *DescribeEipAddressesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeEipAddressesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeEipAddressesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEipAddressesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeEipAddressesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeEipAddressesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeEipAddressesRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeEipAddressesRequest
	GetStatus() *string
}

type DescribeEipAddressesRequest struct {
	Filter []*DescribeEipAddressesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The ID of the EIP to query. You can enter up to 50 IDs. Separate multiple IDs with commas.
	//
	// example:
	//
	// eip-2zeerraiwb7ujxscd****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The ID of the associated cloud service instance.
	//
	// example:
	//
	// i-2zebb08phyccdvf****
	AssociatedInstanceId *string `json:"AssociatedInstanceId,omitempty" xml:"AssociatedInstanceId,omitempty"`
	// The type of the associated cloud service instance. Valid values: EcsInstance, SlbInstance, Nat, HaVip, NetworkInterface, and IpAddress.
	//
	// example:
	//
	// EcsInstance
	AssociatedInstanceType *string `json:"AssociatedInstanceType,omitempty" xml:"AssociatedInstanceType,omitempty"`
	// The billing method of the EIP. Valid values: PostPaid (pay-as-you-go) and PrePaid (subscription).
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The EIP address to query. You can enter up to 50 addresses. Separate multiple addresses with commas.
	//
	// example:
	//
	// 47.75.XX.XX
	EipAddress *string `json:"EipAddress,omitempty" xml:"EipAddress,omitempty"`
	// The line type. Valid values include BGP and BGP_PRO.
	//
	// example:
	//
	// BGP
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The lock type. Valid values: financial (locked due to overdue payment) and security (locked due to security reasons).
	//
	// example:
	//
	// financial
	LockReason   *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the EIP resides. You can call the DescribeRegions operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the EIP. Valid values: Associating, Unassociating, InUse, Available, and Releasing.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEipAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesRequest) GetFilter() []*DescribeEipAddressesRequestFilter {
	return s.Filter
}

func (s *DescribeEipAddressesRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeEipAddressesRequest) GetAssociatedInstanceId() *string {
	return s.AssociatedInstanceId
}

func (s *DescribeEipAddressesRequest) GetAssociatedInstanceType() *string {
	return s.AssociatedInstanceType
}

func (s *DescribeEipAddressesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeEipAddressesRequest) GetEipAddress() *string {
	return s.EipAddress
}

func (s *DescribeEipAddressesRequest) GetISP() *string {
	return s.ISP
}

func (s *DescribeEipAddressesRequest) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeEipAddressesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeEipAddressesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeEipAddressesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEipAddressesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEipAddressesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEipAddressesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeEipAddressesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeEipAddressesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeEipAddressesRequest) SetFilter(v []*DescribeEipAddressesRequestFilter) *DescribeEipAddressesRequest {
	s.Filter = v
	return s
}

func (s *DescribeEipAddressesRequest) SetAllocationId(v string) *DescribeEipAddressesRequest {
	s.AllocationId = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetAssociatedInstanceId(v string) *DescribeEipAddressesRequest {
	s.AssociatedInstanceId = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetAssociatedInstanceType(v string) *DescribeEipAddressesRequest {
	s.AssociatedInstanceType = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetChargeType(v string) *DescribeEipAddressesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetEipAddress(v string) *DescribeEipAddressesRequest {
	s.EipAddress = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetISP(v string) *DescribeEipAddressesRequest {
	s.ISP = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetLockReason(v string) *DescribeEipAddressesRequest {
	s.LockReason = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetOwnerAccount(v string) *DescribeEipAddressesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetOwnerId(v int64) *DescribeEipAddressesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetPageNumber(v int32) *DescribeEipAddressesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetPageSize(v int32) *DescribeEipAddressesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetRegionId(v string) *DescribeEipAddressesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetResourceOwnerAccount(v string) *DescribeEipAddressesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetResourceOwnerId(v int64) *DescribeEipAddressesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetStatus(v string) *DescribeEipAddressesRequest {
	s.Status = &v
	return s
}

func (s *DescribeEipAddressesRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEipAddressesRequestFilter struct {
	// The first filter key. Set the value to CreationStartTime, which specifies the start time of resource creation.
	//
	// example:
	//
	// CreationStartTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The first filter value. Specify the time in UTC in the YYYY-MM-DDThh:mmZ format.
	//
	// example:
	//
	// 2023-01-01T01:00Z
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEipAddressesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeEipAddressesRequestFilter) GetValue() *string {
	return s.Value
}

func (s *DescribeEipAddressesRequestFilter) SetKey(v string) *DescribeEipAddressesRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeEipAddressesRequestFilter) SetValue(v string) *DescribeEipAddressesRequestFilter {
	s.Value = &v
	return s
}

func (s *DescribeEipAddressesRequestFilter) Validate() error {
	return dara.Validate(s)
}
