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
	SetDryRun(v bool) *DescribeEipAddressesRequest
	GetDryRun() *bool
	SetEipAddress(v string) *DescribeEipAddressesRequest
	GetEipAddress() *string
	SetEipName(v string) *DescribeEipAddressesRequest
	GetEipName() *string
	SetISP(v string) *DescribeEipAddressesRequest
	GetISP() *string
	SetIncludeReservationData(v bool) *DescribeEipAddressesRequest
	GetIncludeReservationData() *bool
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
	SetPublicIpAddressPoolId(v string) *DescribeEipAddressesRequest
	GetPublicIpAddressPoolId() *string
	SetRegionId(v string) *DescribeEipAddressesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeEipAddressesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeEipAddressesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeEipAddressesRequest
	GetResourceOwnerId() *int64
	SetSecurityProtectionEnabled(v bool) *DescribeEipAddressesRequest
	GetSecurityProtectionEnabled() *bool
	SetSegmentInstanceId(v string) *DescribeEipAddressesRequest
	GetSegmentInstanceId() *string
	SetServiceManaged(v bool) *DescribeEipAddressesRequest
	GetServiceManaged() *bool
	SetStatus(v string) *DescribeEipAddressesRequest
	GetStatus() *string
	SetTag(v []*DescribeEipAddressesRequestTag) *DescribeEipAddressesRequest
	GetTag() []*DescribeEipAddressesRequestTag
}

type DescribeEipAddressesRequest struct {
	Filter []*DescribeEipAddressesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The ID of the EIP that you want to query.
	//
	// You can specify up to 50 EIP IDs. Separate multiple IDs with commas (,).
	//
	// >  If both **EipAddress*	- and **AllocationId*	- are specified, you can specify up to 50 EIP IDs for **AllocationId**, and specify up to 50 EIPs for **EipAddress**.
	//
	// example:
	//
	// eip-2zeerraiwb7ujxscd****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The ID of the instance associated with the EIP.
	//
	// example:
	//
	// i-2zebb08phyccdvf****
	AssociatedInstanceId *string `json:"AssociatedInstanceId,omitempty" xml:"AssociatedInstanceId,omitempty"`
	// The type of the cloud resource with which you want to associate the EIP. Valid values:
	//
	// 	- **EcsInstance*	- (default): an Elastic Compute Service (ECS) instance in a virtual private cloud (VPC).
	//
	// 	- **SlbInstance**: a CLB instance in a VPC.
	//
	// 	- **Nat**: a NAT gateway.
	//
	// 	- **HaVip**: an HAVIP.
	//
	// 	- **NetworkInterface**: a secondary ENI.
	//
	// 	- **IpAddress**: an IP address.
	//
	// >  Each ECS instance, CLB instance, HAVIP, and IP address can be associated with only one EIP. A NAT gateway can be associated with multiple EIPs. The number of EIPs that you can associate with a secondary ENI depends on the association mode. For more information, see [Associate EIPs with and disassociate EIPs from cloud resources](https://help.aliyun.com/document_detail/72125.html).
	//
	// example:
	//
	// EcsInstance
	AssociatedInstanceType *string `json:"AssociatedInstanceType,omitempty" xml:"AssociatedInstanceType,omitempty"`
	// The billing method of the EIP. Valid values:
	//
	// 	- **PostPaid**: pay-as-you-go.
	//
	// 	- **PrePaid**: subscription.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The EIP that you want to query.
	//
	// You can specify up to 50 EIPs. Separate multiple EIPs with commas (,).
	//
	// >  If both **EipAddress*	- and **AllocationId*	- are specified, you can specify up to 50 EIPs for **EipAddress**, and specify up to 50 EIP IDs for **AllocationId**.
	//
	// example:
	//
	// 47.75.XX.XX
	EipAddress *string `json:"EipAddress,omitempty" xml:"EipAddress,omitempty"`
	// The name of the EIP.
	//
	// The name must be 1 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// EIP-01
	EipName *string `json:"EipName,omitempty" xml:"EipName,omitempty"`
	// The line type. Valid values:
	//
	// 	- **BGP*	- (default): Border Gateway Protocol (BGP) (Multi-ISP) lines. All regions support BGP (Multi-ISP) EIPs.
	//
	// 	- **BGP_PRO**: BGP (Multi-ISP) Pro lines. Only the following regions support BGP (Multi-ISP) Pro lines: China (Hong Kong), Singapore, Japan (Tokyo), Malaysia (Kuala Lumpur), Philippines (Manila), Indonesia (Jakarta), and Thailand (Bangkok).
	//
	// For more information about BGP (Multi-ISP) and BGP (Multi-ISP) Pro, see the [Line types](https://help.aliyun.com/document_detail/32321.html) section of the "What is EIP?" topic.
	//
	// If you are allowed to use single-ISP bandwidth, you can also use one of the following values:
	//
	// 	- **ChinaTelecom**
	//
	// 	- **ChinaUnicom**
	//
	// 	- **ChinaMobile**
	//
	// 	- **ChinaTelecom_L2**
	//
	// 	- **ChinaUnicom_L2**
	//
	// 	- **ChinaMobile_L2**
	//
	// If your services are deployed in China East 1 Finance, this parameter is required and you must set the value to **BGP_FinanceCloud**.
	//
	// example:
	//
	// BGP
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// Specifies whether to return information about pending orders. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	IncludeReservationData *bool `json:"IncludeReservationData,omitempty" xml:"IncludeReservationData,omitempty"`
	// The reason why the EIP is locked. Valid values:
	//
	// 	- **financial**: The EIP is locked due to overdue payments.
	//
	// 	- **security**: The EIP is locked for security reasons.
	//
	// example:
	//
	// financial
	LockReason   *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IP address pool to which the EIP that you want to query belongs.
	//
	// example:
	//
	// pippool-2vc0kxcedhquybdsz****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	// The region ID of the EIP.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the EIP belongs.
	//
	// example:
	//
	// rg-acfmxazb4pcdvf****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to activate Anti-DDoS Pro/Premium. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// false
	SecurityProtectionEnabled *bool `json:"SecurityProtectionEnabled,omitempty" xml:"SecurityProtectionEnabled,omitempty"`
	// The ID of the contiguous EIP group.
	//
	// example:
	//
	// eipsg-t4nr90yik5oy38xdy****
	SegmentInstanceId *string `json:"SegmentInstanceId,omitempty" xml:"SegmentInstanceId,omitempty"`
	// Indicates whether the instance is managed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no.
	//
	// If you do not specify this parameter, all instances are queried.
	//
	// example:
	//
	// false
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The state of the EIP. Valid values:
	//
	// 	- **Associating**
	//
	// 	- **Unassociating**
	//
	// 	- **InUse**
	//
	// 	- **Available**
	//
	// 	- **Releasing**
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags used to filter EIPs.
	Tag []*DescribeEipAddressesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *DescribeEipAddressesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeEipAddressesRequest) GetEipAddress() *string {
	return s.EipAddress
}

func (s *DescribeEipAddressesRequest) GetEipName() *string {
	return s.EipName
}

func (s *DescribeEipAddressesRequest) GetISP() *string {
	return s.ISP
}

func (s *DescribeEipAddressesRequest) GetIncludeReservationData() *bool {
	return s.IncludeReservationData
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

func (s *DescribeEipAddressesRequest) GetPublicIpAddressPoolId() *string {
	return s.PublicIpAddressPoolId
}

func (s *DescribeEipAddressesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEipAddressesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeEipAddressesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeEipAddressesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeEipAddressesRequest) GetSecurityProtectionEnabled() *bool {
	return s.SecurityProtectionEnabled
}

func (s *DescribeEipAddressesRequest) GetSegmentInstanceId() *string {
	return s.SegmentInstanceId
}

func (s *DescribeEipAddressesRequest) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeEipAddressesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeEipAddressesRequest) GetTag() []*DescribeEipAddressesRequestTag {
	return s.Tag
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

func (s *DescribeEipAddressesRequest) SetDryRun(v bool) *DescribeEipAddressesRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetEipAddress(v string) *DescribeEipAddressesRequest {
	s.EipAddress = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetEipName(v string) *DescribeEipAddressesRequest {
	s.EipName = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetISP(v string) *DescribeEipAddressesRequest {
	s.ISP = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetIncludeReservationData(v bool) *DescribeEipAddressesRequest {
	s.IncludeReservationData = &v
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

func (s *DescribeEipAddressesRequest) SetPublicIpAddressPoolId(v string) *DescribeEipAddressesRequest {
	s.PublicIpAddressPoolId = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetRegionId(v string) *DescribeEipAddressesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetResourceGroupId(v string) *DescribeEipAddressesRequest {
	s.ResourceGroupId = &v
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

func (s *DescribeEipAddressesRequest) SetSecurityProtectionEnabled(v bool) *DescribeEipAddressesRequest {
	s.SecurityProtectionEnabled = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetSegmentInstanceId(v string) *DescribeEipAddressesRequest {
	s.SegmentInstanceId = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetServiceManaged(v bool) *DescribeEipAddressesRequest {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetStatus(v string) *DescribeEipAddressesRequest {
	s.Status = &v
	return s
}

func (s *DescribeEipAddressesRequest) SetTag(v []*DescribeEipAddressesRequestTag) *DescribeEipAddressesRequest {
	s.Tag = v
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
	if s.Tag != nil {
		for _, item := range s.Tag {
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
	// The filter key used to query resources. Set the value to **CreationStartTime**, which specifies the time when the system started to create the resource.
	//
	// example:
	//
	// CreationStartTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value used to query resources. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mmZ` format. The time must be in Coordinated Universal Time (UTC).
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

type DescribeEipAddressesRequestTag struct {
	// The key of the tag. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEipAddressesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeEipAddressesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeEipAddressesRequestTag) SetKey(v string) *DescribeEipAddressesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeEipAddressesRequestTag) SetValue(v string) *DescribeEipAddressesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeEipAddressesRequestTag) Validate() error {
	return dara.Validate(s)
}
