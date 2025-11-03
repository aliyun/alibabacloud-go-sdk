// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEipAddresses(v *DescribeEipAddressesResponseBodyEipAddresses) *DescribeEipAddressesResponseBody
	GetEipAddresses() *DescribeEipAddressesResponseBodyEipAddresses
	SetPageNumber(v int32) *DescribeEipAddressesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEipAddressesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEipAddressesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeEipAddressesResponseBody
	GetTotalCount() *int32
}

type DescribeEipAddressesResponseBody struct {
	// Details of the EIPs.
	EipAddresses *DescribeEipAddressesResponseBodyEipAddresses `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEipAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBody) GetEipAddresses() *DescribeEipAddressesResponseBodyEipAddresses {
	return s.EipAddresses
}

func (s *DescribeEipAddressesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEipAddressesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEipAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEipAddressesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEipAddressesResponseBody) SetEipAddresses(v *DescribeEipAddressesResponseBodyEipAddresses) *DescribeEipAddressesResponseBody {
	s.EipAddresses = v
	return s
}

func (s *DescribeEipAddressesResponseBody) SetPageNumber(v int32) *DescribeEipAddressesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEipAddressesResponseBody) SetPageSize(v int32) *DescribeEipAddressesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEipAddressesResponseBody) SetRequestId(v string) *DescribeEipAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEipAddressesResponseBody) SetTotalCount(v int32) *DescribeEipAddressesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEipAddressesResponseBody) Validate() error {
	if s.EipAddresses != nil {
		if err := s.EipAddresses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEipAddressesResponseBodyEipAddresses struct {
	EipAddress []*DescribeEipAddressesResponseBodyEipAddressesEipAddress `json:"EipAddress,omitempty" xml:"EipAddress,omitempty" type:"Repeated"`
}

func (s DescribeEipAddressesResponseBodyEipAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesResponseBodyEipAddresses) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBodyEipAddresses) GetEipAddress() []*DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	return s.EipAddress
}

func (s *DescribeEipAddressesResponseBodyEipAddresses) SetEipAddress(v []*DescribeEipAddressesResponseBodyEipAddressesEipAddress) *DescribeEipAddressesResponseBodyEipAddresses {
	s.EipAddress = v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddresses) Validate() error {
	if s.EipAddress != nil {
		for _, item := range s.EipAddress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEipAddressesResponseBodyEipAddressesEipAddress struct {
	// The ID of the EIP.
	//
	// example:
	//
	// eip-2zeerraiwb7ujcdvf****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The time when the EIP was created. The time follows the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2021-04-23T01:37:38Z
	AllocationTime *string `json:"AllocationTime,omitempty" xml:"AllocationTime,omitempty"`
	// The maximum bandwidth of the EIP. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The maximum bandwidth of the Internet Shared Bandwidth instance with which the EIP is associated. Unit: Mbit/s.
	//
	// example:
	//
	// 50
	BandwidthPackageBandwidth *string `json:"BandwidthPackageBandwidth,omitempty" xml:"BandwidthPackageBandwidth,omitempty"`
	// The ID of the Internet Shared Bandwidth instance.
	//
	// example:
	//
	// cbwp-bp1ego3i4j07ccdvf****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The type of the bandwidth. Only **CommonBandwidthPackage*	- may be returned, which indicates Internet Shared Bandwidth.
	//
	// example:
	//
	// CommonBandwidthPackage
	BandwidthPackageType *string `json:"BandwidthPackageType,omitempty" xml:"BandwidthPackageType,omitempty"`
	// The service type. Valid values:
	//
	// 	- **CloudBox*	- Only cloud box users can select this type.
	//
	// 	- **Default*	- (default)
	//
	// example:
	//
	// CloudBox
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The service status of the EIP. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **OperationLock**
	//
	// 	- **Unactivated**
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
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
	// Indicates whether deletion protection is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The description of the EIP.
	//
	// example:
	//
	// abc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum bandwidth of the EIP when it is not associated with an Internet Shared Bandwidth instance. Unit: Mbit/s.
	//
	// example:
	//
	// 101
	EipBandwidth *string `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	// The time when the EIP expires. The time follows the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2021-05-23T02:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// Indicates whether fine-grained monitoring is enabled for the EIP. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// false
	HDMonitorStatus *string `json:"HDMonitorStatus,omitempty" xml:"HDMonitorStatus,omitempty"`
	// Indicates whether renewal data is included. Valid values:
	//
	// 	- **false**
	//
	// 	- **true*	- A value of **true*	- is returned only when **IncludeReservationData*	- is set to **true*	- and some orders have not taken effect.
	//
	// example:
	//
	// false
	HasReservationData *string `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
	// The line type. Valid values:
	//
	// 	- **BGP**: BGP (Multi-ISP). The BGP (Multi-ISP) line is supported in all regions.
	//
	// 	- **BGP_PRO**: BGP (Multi-ISP) Pro lines. BGP (Multi-ISP) Pro line is supported only in the China (Hong Kong), Singapore, Japan (Tokyo), Malaysia (Kuala Lumpur), Philippines (Manila), Indonesia (Jakarta), and Thailand (Bangkok) regions.
	//
	// For more information about BGP (Multi-ISP) and BGP (Multi-ISP) Pro, see the [Line types](https://help.aliyun.com/document_detail/32321.html) section of the "What is EIP?" topic.
	//
	// If you are allowed to use single-ISP bandwidth, one of the following values may be returned:
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
	// If your services are deployed in China East 1 Finance, **BGP_FinanceCloud*	- is returned.
	//
	// example:
	//
	// BGP
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The ID of the associated instance.
	//
	// example:
	//
	// i-bp15zckdt37cdvf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the associated instance.
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	// The type of the associated instance. Valid values:
	//
	// 	- **EcsInstance**: an ECS instance in a VPC.
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
	// example:
	//
	// EcsInstance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The metering method of the EIP. Valid values:
	//
	// 	- **PayByBandwidth**
	//
	// 	- **PayByTraffic**
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The EIP.
	//
	// example:
	//
	// 47.75.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The association mode. Valid values:
	//
	// - **NAT**: NAT mode
	//
	// - **MULTI_BINDED**: multi-EIP-to-ENI mode
	//
	// - **BINDED**: cut-through mode
	//
	// example:
	//
	// NAT
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the EIP.
	//
	// example:
	//
	// EIP-01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type. Only **public*	- may be returned.
	//
	// example:
	//
	// public
	Netmode *string `json:"Netmode,omitempty" xml:"Netmode,omitempty"`
	// The details about the locked EIP.
	OperationLocks *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	// The private IP address of the secondary ENI with which the EIP is associated.
	//
	// example:
	//
	// 192.168.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The ID of the IP address pool to which the EIP belongs.
	//
	// example:
	//
	// pippool-2vc0kxcedhquybdsz****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	// The region ID of the EIP.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time when the renewal took effect. The time follows the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2021-05-23T16:00:00Z
	ReservationActiveTime *string `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	// The maximum bandwidth after the renewal takes effect. Unit: Mbit/s.
	//
	// example:
	//
	// 12
	ReservationBandwidth *string `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	// The metering method that is used after the renewal takes effect. Valid values:
	//
	// 	- **PayByBandwidth**
	//
	// 	- **PayByTraffic**
	//
	// example:
	//
	// PayByBandwidth
	ReservationInternetChargeType *string `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	// The type of the renewal order. Valid values:
	//
	// 	- **RENEWCHANGE**: renewal with an upgrade or a downgrade.
	//
	// 	- **TEMP_UPGRADE**: temporary upgrade.
	//
	// 	- **UPGRADE**: upgrade.
	//
	// example:
	//
	// RENEWCHANGE
	ReservationOrderType *string `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazcdxs****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether level-2 throttling is configured. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	SecondLimited *bool `json:"SecondLimited,omitempty" xml:"SecondLimited,omitempty"`
	// The edition of Anti-DDoS.
	//
	// 	- If an empty value is returned, it indicates that Anti-DDoS Origin Basic is used.
	//
	// 	- If **AntiDDoS_Enhanced*	- is returned, it indicates that Anti-DDoS Pro/Premium is used.
	SecurityProtectionTypes *DescribeEipAddressesResponseBodyEipAddressesEipAddressSecurityProtectionTypes `json:"SecurityProtectionTypes,omitempty" xml:"SecurityProtectionTypes,omitempty" type:"Struct"`
	// The ID of the contiguous EIP group.
	//
	// This value is returned only when you query contiguous EIPs.
	//
	// example:
	//
	// eipsg-t4nr90yik5oy38xd****
	SegmentInstanceId *string `json:"SegmentInstanceId,omitempty" xml:"SegmentInstanceId,omitempty"`
	// The ID of the service provider to which the managed instance belongs.
	//
	// > This is only valid when the ServiceManaged parameter is set to True.
	//
	// example:
	//
	// 197*************
	ServiceID *int64 `json:"ServiceID,omitempty" xml:"ServiceID,omitempty"`
	// Indicates whether the instance is managed. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 0
	ServiceManaged *int32 `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The status of the EIP. Valid values:
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
	// The tags of the EIP.
	Tags *DescribeEipAddressesResponseBodyEipAddressesEipAddressTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the VPC in which an IPv4 gateway is created and that is deployed in the same region as the EIP.
	//
	// When you associate an EIP with an IP address, the system can enable the IP address to access the Internet based on VPC route configurations.
	//
	// >  This parameter is returned if the value of **InstanceType*	- is **IpAddress**. In this case, the EIP is associated with an IP address.
	//
	// example:
	//
	// vpc-bp15zckdt37pq72zv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone of the EIP.
	//
	// This parameter is returned only when the service type is CloudBox.
	//
	// example:
	//
	// cn-hangzhou-a
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddress) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetAllocationTime() *string {
	return s.AllocationTime
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetBandwidthPackageBandwidth() *string {
	return s.BandwidthPackageBandwidth
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetBandwidthPackageType() *string {
	return s.BandwidthPackageType
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetBizType() *string {
	return s.BizType
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetDescription() *string {
	return s.Description
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetEipBandwidth() *string {
	return s.EipBandwidth
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetHDMonitorStatus() *string {
	return s.HDMonitorStatus
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetHasReservationData() *string {
	return s.HasReservationData
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetISP() *string {
	return s.ISP
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetInstanceRegionId() *string {
	return s.InstanceRegionId
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetMode() *string {
	return s.Mode
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetName() *string {
	return s.Name
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetNetmode() *string {
	return s.Netmode
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetOperationLocks() *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks {
	return s.OperationLocks
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetPublicIpAddressPoolId() *string {
	return s.PublicIpAddressPoolId
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetReservationActiveTime() *string {
	return s.ReservationActiveTime
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetReservationBandwidth() *string {
	return s.ReservationBandwidth
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetReservationInternetChargeType() *string {
	return s.ReservationInternetChargeType
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetReservationOrderType() *string {
	return s.ReservationOrderType
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetSecondLimited() *bool {
	return s.SecondLimited
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetSecurityProtectionTypes() *DescribeEipAddressesResponseBodyEipAddressesEipAddressSecurityProtectionTypes {
	return s.SecurityProtectionTypes
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetSegmentInstanceId() *string {
	return s.SegmentInstanceId
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetServiceID() *int64 {
	return s.ServiceID
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetServiceManaged() *int32 {
	return s.ServiceManaged
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetStatus() *string {
	return s.Status
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetTags() *DescribeEipAddressesResponseBodyEipAddressesEipAddressTags {
	return s.Tags
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetZone() *string {
	return s.Zone
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetAllocationId(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.AllocationId = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetAllocationTime(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.AllocationTime = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetBandwidth(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.Bandwidth = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetBandwidthPackageBandwidth(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.BandwidthPackageBandwidth = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetBandwidthPackageId(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetBandwidthPackageType(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.BandwidthPackageType = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetBizType(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.BizType = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetBusinessStatus(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetChargeType(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.ChargeType = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetDeletionProtection(v bool) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetDescription(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.Description = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetEipBandwidth(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.EipBandwidth = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetExpiredTime(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetHDMonitorStatus(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.HDMonitorStatus = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetHasReservationData(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.HasReservationData = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetISP(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.ISP = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetInstanceId(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.InstanceId = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetInstanceRegionId(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.InstanceRegionId = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetInstanceType(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.InstanceType = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetInternetChargeType(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetIpAddress(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.IpAddress = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetMode(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.Mode = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetName(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.Name = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetNetmode(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.Netmode = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetOperationLocks(v *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.OperationLocks = v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetPrivateIpAddress(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetPublicIpAddressPoolId(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.PublicIpAddressPoolId = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetRegionId(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.RegionId = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetReservationActiveTime(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.ReservationActiveTime = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetReservationBandwidth(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.ReservationBandwidth = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetReservationInternetChargeType(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.ReservationInternetChargeType = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetReservationOrderType(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.ReservationOrderType = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetResourceGroupId(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetSecondLimited(v bool) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.SecondLimited = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetSecurityProtectionTypes(v *DescribeEipAddressesResponseBodyEipAddressesEipAddressSecurityProtectionTypes) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.SecurityProtectionTypes = v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetSegmentInstanceId(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.SegmentInstanceId = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetServiceID(v int64) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.ServiceID = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetServiceManaged(v int32) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetStatus(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.Status = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetTags(v *DescribeEipAddressesResponseBodyEipAddressesEipAddressTags) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.Tags = v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetVpcId(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.VpcId = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetZone(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.Zone = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) Validate() error {
	if s.OperationLocks != nil {
		if err := s.OperationLocks.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityProtectionTypes != nil {
		if err := s.SecurityProtectionTypes.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks struct {
	LockReason []*DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) GetLockReason() []*DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason {
	return s.LockReason
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) SetLockReason(v []*DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason) *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks {
	s.LockReason = v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) Validate() error {
	if s.LockReason != nil {
		for _, item := range s.LockReason {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason struct {
	// The reason why the EIP is locked. Valid values:
	//
	// 	- **financial**: The EIP is locked due to overdue payments.
	//
	// 	- **security**: The instance is locked for security purposes.
	//
	// 	- **sharedPool**: The shared IP address pool is locked due to overdue payments.
	//
	// example:
	//
	// financial
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason) SetLockReason(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason {
	s.LockReason = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason) Validate() error {
	return dara.Validate(s)
}

type DescribeEipAddressesResponseBodyEipAddressesEipAddressSecurityProtectionTypes struct {
	SecurityProtectionType []*string `json:"SecurityProtectionType,omitempty" xml:"SecurityProtectionType,omitempty" type:"Repeated"`
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddressSecurityProtectionTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddressSecurityProtectionTypes) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressSecurityProtectionTypes) GetSecurityProtectionType() []*string {
	return s.SecurityProtectionType
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressSecurityProtectionTypes) SetSecurityProtectionType(v []*string) *DescribeEipAddressesResponseBodyEipAddressesEipAddressSecurityProtectionTypes {
	s.SecurityProtectionType = v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressSecurityProtectionTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeEipAddressesResponseBodyEipAddressesEipAddressTags struct {
	Tag []*DescribeEipAddressesResponseBodyEipAddressesEipAddressTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddressTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddressTags) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressTags) GetTag() []*DescribeEipAddressesResponseBodyEipAddressesEipAddressTagsTag {
	return s.Tag
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressTags) SetTag(v []*DescribeEipAddressesResponseBodyEipAddressesEipAddressTagsTag) *DescribeEipAddressesResponseBodyEipAddressesEipAddressTags {
	s.Tag = v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressTags) Validate() error {
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

type DescribeEipAddressesResponseBodyEipAddressesEipAddressTagsTag struct {
	// The tag key of the EIP.
	//
	// example:
	//
	// KeyTest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the EIP.
	//
	// example:
	//
	// ValueTest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddressTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddressTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressTagsTag) SetKey(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddressTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressTagsTag) SetValue(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddressTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressTagsTag) Validate() error {
	return dara.Validate(s)
}
