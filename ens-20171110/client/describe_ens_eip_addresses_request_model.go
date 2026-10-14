// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsEipAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *DescribeEnsEipAddressesRequest
	GetAllocationId() *string
	SetAssociatedInstanceId(v string) *DescribeEnsEipAddressesRequest
	GetAssociatedInstanceId() *string
	SetAssociatedInstanceType(v string) *DescribeEnsEipAddressesRequest
	GetAssociatedInstanceType() *string
	SetEipAddress(v string) *DescribeEnsEipAddressesRequest
	GetEipAddress() *string
	SetEipName(v string) *DescribeEnsEipAddressesRequest
	GetEipName() *string
	SetEnsRegionId(v string) *DescribeEnsEipAddressesRequest
	GetEnsRegionId() *string
	SetEnsRegionIds(v []*string) *DescribeEnsEipAddressesRequest
	GetEnsRegionIds() []*string
	SetIcmpReplyEnabled(v bool) *DescribeEnsEipAddressesRequest
	GetIcmpReplyEnabled() *bool
	SetLockReason(v string) *DescribeEnsEipAddressesRequest
	GetLockReason() *string
	SetPageNumber(v int32) *DescribeEnsEipAddressesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEnsEipAddressesRequest
	GetPageSize() *int32
	SetStandby(v string) *DescribeEnsEipAddressesRequest
	GetStandby() *string
	SetStatus(v string) *DescribeEnsEipAddressesRequest
	GetStatus() *string
}

type DescribeEnsEipAddressesRequest struct {
	// The ID of the EIP instance to query. You can specify up to 50 EIP instance IDs separated by commas (,).
	//
	// example:
	//
	// eip-5q9uwkd9bznjpxz8hr6cirnjk
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The instance ID of the cloud service.
	//
	// example:
	//
	// lb-5t18quoohsrc3xkf86spmnu77
	AssociatedInstanceId *string `json:"AssociatedInstanceId,omitempty" xml:"AssociatedInstanceId,omitempty"`
	// The type of the cloud service instance to associate. Valid values:
	//
	// - **EnsInstance**: a VPC-type ENS instance.
	//
	// - **SlbInstance**: a load balancing instance.
	//
	// example:
	//
	// SlbInstance
	AssociatedInstanceType *string `json:"AssociatedInstanceType,omitempty" xml:"AssociatedInstanceType,omitempty"`
	// The IP address of the EIP to query. You can specify up to 50 EIP addresses separated by commas (,).
	//
	// example:
	//
	// 120.XXX.XXX.4
	EipAddress *string `json:"EipAddress,omitempty" xml:"EipAddress,omitempty"`
	// The name of the EIP.
	//
	// example:
	//
	// test
	EipName *string `json:"EipName,omitempty" xml:"EipName,omitempty"`
	// The ENS edge zone ID.
	//
	// example:
	//
	// cn-chengdu-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The array of ENS edge zone IDs. Array length: 1 to 100.
	EnsRegionIds []*string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" type:"Repeated"`
	// Specifies whether ICMP echo reply is enabled. Valid values:
	//
	// - **true*	- (default): enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// false
	IcmpReplyEnabled *bool `json:"IcmpReplyEnabled,omitempty" xml:"IcmpReplyEnabled,omitempty"`
	// The lock type.
	//
	// - financial: locked due to overdue payment.
	//
	// example:
	//
	// financial
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The page number of the list. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether the EIP is a standby EIP. Valid values:
	//
	// - **true**: standby.
	//
	// - **false**: not standby.
	//
	// example:
	//
	// true
	Standby *string `json:"Standby,omitempty" xml:"Standby,omitempty"`
	// The status of the EIP. Valid values: Associating: being associated. Unassociating: being disassociated. InUse: allocated. Available: available.
	//
	// example:
	//
	// InUse
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEnsEipAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsEipAddressesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsEipAddressesRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeEnsEipAddressesRequest) GetAssociatedInstanceId() *string {
	return s.AssociatedInstanceId
}

func (s *DescribeEnsEipAddressesRequest) GetAssociatedInstanceType() *string {
	return s.AssociatedInstanceType
}

func (s *DescribeEnsEipAddressesRequest) GetEipAddress() *string {
	return s.EipAddress
}

func (s *DescribeEnsEipAddressesRequest) GetEipName() *string {
	return s.EipName
}

func (s *DescribeEnsEipAddressesRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEnsEipAddressesRequest) GetEnsRegionIds() []*string {
	return s.EnsRegionIds
}

func (s *DescribeEnsEipAddressesRequest) GetIcmpReplyEnabled() *bool {
	return s.IcmpReplyEnabled
}

func (s *DescribeEnsEipAddressesRequest) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeEnsEipAddressesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEnsEipAddressesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEnsEipAddressesRequest) GetStandby() *string {
	return s.Standby
}

func (s *DescribeEnsEipAddressesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeEnsEipAddressesRequest) SetAllocationId(v string) *DescribeEnsEipAddressesRequest {
	s.AllocationId = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetAssociatedInstanceId(v string) *DescribeEnsEipAddressesRequest {
	s.AssociatedInstanceId = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetAssociatedInstanceType(v string) *DescribeEnsEipAddressesRequest {
	s.AssociatedInstanceType = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetEipAddress(v string) *DescribeEnsEipAddressesRequest {
	s.EipAddress = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetEipName(v string) *DescribeEnsEipAddressesRequest {
	s.EipName = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetEnsRegionId(v string) *DescribeEnsEipAddressesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetEnsRegionIds(v []*string) *DescribeEnsEipAddressesRequest {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetIcmpReplyEnabled(v bool) *DescribeEnsEipAddressesRequest {
	s.IcmpReplyEnabled = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetLockReason(v string) *DescribeEnsEipAddressesRequest {
	s.LockReason = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetPageNumber(v int32) *DescribeEnsEipAddressesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetPageSize(v int32) *DescribeEnsEipAddressesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetStandby(v string) *DescribeEnsEipAddressesRequest {
	s.Standby = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) SetStatus(v string) *DescribeEnsEipAddressesRequest {
	s.Status = &v
	return s
}

func (s *DescribeEnsEipAddressesRequest) Validate() error {
	return dara.Validate(s)
}
