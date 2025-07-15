// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonBandwidthPackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommonBandwidthPackages(v *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackages) *DescribeCommonBandwidthPackagesResponseBody
	GetCommonBandwidthPackages() *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackages
	SetPageNumber(v int32) *DescribeCommonBandwidthPackagesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCommonBandwidthPackagesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCommonBandwidthPackagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCommonBandwidthPackagesResponseBody
	GetTotalCount() *int32
}

type DescribeCommonBandwidthPackagesResponseBody struct {
	// The details of the Internet Shared Bandwidth instance.
	CommonBandwidthPackages *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackages `json:"CommonBandwidthPackages,omitempty" xml:"CommonBandwidthPackages,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
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
	// 20E6FD1C-7321-4DAD-BDFD-EC8769E4AA33
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCommonBandwidthPackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonBandwidthPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCommonBandwidthPackagesResponseBody) GetCommonBandwidthPackages() *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackages {
	return s.CommonBandwidthPackages
}

func (s *DescribeCommonBandwidthPackagesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCommonBandwidthPackagesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCommonBandwidthPackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCommonBandwidthPackagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCommonBandwidthPackagesResponseBody) SetCommonBandwidthPackages(v *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackages) *DescribeCommonBandwidthPackagesResponseBody {
	s.CommonBandwidthPackages = v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBody) SetPageNumber(v int32) *DescribeCommonBandwidthPackagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBody) SetPageSize(v int32) *DescribeCommonBandwidthPackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBody) SetRequestId(v string) *DescribeCommonBandwidthPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBody) SetTotalCount(v int32) *DescribeCommonBandwidthPackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackages struct {
	CommonBandwidthPackage []*DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage `json:"CommonBandwidthPackage,omitempty" xml:"CommonBandwidthPackage,omitempty" type:"Repeated"`
}

func (s DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackages) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackages) GoString() string {
	return s.String()
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackages) GetCommonBandwidthPackage() []*DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	return s.CommonBandwidthPackage
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackages) SetCommonBandwidthPackage(v []*DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackages {
	s.CommonBandwidthPackage = v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackages) Validate() error {
	return dara.Validate(s)
}

type DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage struct {
	// The maximum bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s.
	//
	// example:
	//
	// 20
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the Internet Shared Bandwidth instance.
	//
	// example:
	//
	// cbwp-bp1t3sm1ffzmshdki****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The service type of the Internet Shared Bandwidth instance. Valid values:
	//
	// 	- **CloudBox*	- The cloud box. Only cloud box users can select this type.
	//
	// 	- **Default*	- (default): The general service type.
	//
	// example:
	//
	// CloudBox
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The service status of the Internet Shared Bandwidth instance. Valid values:
	//
	// 	- **Normal**: The Internet Shared Bandwidth instance runs as expected.
	//
	// 	- **FinancialLocked**: An overdue payment occurs in the Internet Shared Bandwidth instance
	//
	// 	- **Unactivated**: The Internet Shared Bandwidth instance is not activated.
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The time when the Internet Shared Bandwidth instance was created. The time is displayed in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2017-06-28T06:39:20Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
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
	// The description of the Internet Shared Bandwidth instance.
	//
	// example:
	//
	// none
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the Internet Shared Bandwidth instance expired. The time is displayed in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2019-01-15T03:08:37Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// Indicates whether the information about pending orders is returned. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// false
	HasReservationData *string `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
	// The line type. Valid values:
	//
	// 	- **BGP**: BGP (Multi-ISP) line The BGP (Multi-ISP) line is supported in all regions.
	//
	// 	- **BGP_PRO**: BGP (Multi-ISP) Pro line The BGP (Multi-ISP) Pro line is supported in the China (Hong Kong), Singapore (Singapore), Japan (Tokyo), Philippines (Manila), Malaysia (Kuala Lumpur), Indonesia (Jakarta), and Thailand (Bangkok) regions.
	//
	// If you are allowed to use single-ISP bandwidth, one of the following values is returned:
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
	// The billing method of the Internet Shared Bandwidth instance. Valid value:
	//
	// **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the Internet Shared Bandwidth instance. Valid value:
	//
	// **PayByTraffic**
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The name of the Internet Shared Bandwidth instance.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The elastic IP addresses (EIPs) that are associated with the Internet Shared Bandwidth instance.
	PublicIpAddresses *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddresses `json:"PublicIpAddresses,omitempty" xml:"PublicIpAddresses,omitempty" type:"Struct"`
	// The percentage of the minimum bandwidth commitment. Only **20*	- is returned.
	//
	// >  This parameter is supported only on the Alibaba Cloud China site.
	//
	// example:
	//
	// 20
	Ratio *int32 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The ID of the region where the Internet Shared Bandwidth instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time when the renewal took effect. The time is displayed in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2018-08-30T16:00:00Z
	ReservationActiveTime *string `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	// The new maximum bandwidth after the configurations are changed. Unit: Mbit/s.
	//
	// example:
	//
	// 1000
	ReservationBandwidth *string `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	// The metering method after the configurations are changed. Valid value:
	//
	// **PayByTraffic**
	//
	// example:
	//
	// PayByBandwidth
	ReservationInternetChargeType *string `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	// The renewal method. Valid values:
	//
	// 	- **RENEWCHANGE**: renewal with a specification change
	//
	// 	- **TEMP_UPGRADE**: renewal with a temporary upgrade
	//
	// 	- **UPGRADE**: renewal with an upgrade
	//
	// example:
	//
	// RENEWCHANGE
	ReservationOrderType *string `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The editions of Anti-DDoS.
	//
	// 	- If this parameter is empty, Anti-DDoS Origin Basic is enabled.
	//
	// 	- If **AntiDDoS_Enhanced*	- is returned, Anti-DDoS Pro/Premium is enabled.
	SecurityProtectionTypes *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageSecurityProtectionTypes `json:"SecurityProtectionTypes,omitempty" xml:"SecurityProtectionTypes,omitempty" type:"Struct"`
	// Indicates whether the resource is created by the service account. Valid values:
	//
	// 	- **0**: The resource is not created by the service account.
	//
	// 	- **1**: The resource is created by the service account.
	//
	// example:
	//
	// 1
	ServiceManaged *int32 `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The status of the Internet Shared Bandwidth instance. Valid values:
	//
	// 	- **Available**: The Internet Shared Bandwidth instance is available.
	//
	// 	- **Modifying**: The Internet Shared Bandwidth instance is being modified.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag that is added to the Internet Shared Bandwidth instance.
	Tags *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The zone of the Internet Shared Bandwidth instance. This parameter is returned only when BizType is set to CloudBox. If BizType is set to Default, an empty value is returned.
	//
	// example:
	//
	// ap-southeast-1-lzdvn-cb
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GoString() string {
	return s.String()
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetBizType() *string {
	return s.BizType
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetDescription() *string {
	return s.Description
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetHasReservationData() *string {
	return s.HasReservationData
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetISP() *string {
	return s.ISP
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetName() *string {
	return s.Name
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetPublicIpAddresses() *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddresses {
	return s.PublicIpAddresses
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetRatio() *int32 {
	return s.Ratio
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetReservationActiveTime() *string {
	return s.ReservationActiveTime
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetReservationBandwidth() *string {
	return s.ReservationBandwidth
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetReservationInternetChargeType() *string {
	return s.ReservationInternetChargeType
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetReservationOrderType() *string {
	return s.ReservationOrderType
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetSecurityProtectionTypes() *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageSecurityProtectionTypes {
	return s.SecurityProtectionTypes
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetServiceManaged() *int32 {
	return s.ServiceManaged
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetStatus() *string {
	return s.Status
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetTags() *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTags {
	return s.Tags
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) GetZone() *string {
	return s.Zone
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetBandwidth(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetBandwidthPackageId(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetBizType(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.BizType = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetBusinessStatus(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetCreationTime(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.CreationTime = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetDeletionProtection(v bool) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetDescription(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.Description = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetExpiredTime(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetHasReservationData(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.HasReservationData = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetISP(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.ISP = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetInstanceChargeType(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetInternetChargeType(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetName(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.Name = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetPublicIpAddresses(v *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddresses) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.PublicIpAddresses = v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetRatio(v int32) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.Ratio = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetRegionId(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.RegionId = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetReservationActiveTime(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.ReservationActiveTime = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetReservationBandwidth(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.ReservationBandwidth = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetReservationInternetChargeType(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.ReservationInternetChargeType = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetReservationOrderType(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.ReservationOrderType = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetResourceGroupId(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetSecurityProtectionTypes(v *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageSecurityProtectionTypes) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.SecurityProtectionTypes = v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetServiceManaged(v int32) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetStatus(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.Status = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetTags(v *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTags) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.Tags = v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) SetZone(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage {
	s.Zone = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage) Validate() error {
	return dara.Validate(s)
}

type DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddresses struct {
	PublicIpAddresse []*DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse `json:"PublicIpAddresse,omitempty" xml:"PublicIpAddresse,omitempty" type:"Repeated"`
}

func (s DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddresses) GoString() string {
	return s.String()
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddresses) GetPublicIpAddresse() []*DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse {
	return s.PublicIpAddresse
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddresses) SetPublicIpAddresse(v []*DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddresses {
	s.PublicIpAddresse = v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddresses) Validate() error {
	return dara.Validate(s)
}

type DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse struct {
	// The ID of the EIP.
	//
	// example:
	//
	// eip-bp13e9i2qst4g6jzi****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// Indicates whether the EIP is associated with the Internet Shared Bandwidth instance. Valid values:
	//
	// 	- **BINDED**
	//
	// 	- **BINDING**
	//
	// example:
	//
	// BINDED
	BandwidthPackageIpRelationStatus *string `json:"BandwidthPackageIpRelationStatus,omitempty" xml:"BandwidthPackageIpRelationStatus,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 47.95.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
}

func (s DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse) GoString() string {
	return s.String()
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse) GetBandwidthPackageIpRelationStatus() *string {
	return s.BandwidthPackageIpRelationStatus
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse) SetAllocationId(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse {
	s.AllocationId = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse) SetBandwidthPackageIpRelationStatus(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse {
	s.BandwidthPackageIpRelationStatus = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse) SetIpAddress(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse {
	s.IpAddress = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse) Validate() error {
	return dara.Validate(s)
}

type DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageSecurityProtectionTypes struct {
	SecurityProtectionType []*string `json:"SecurityProtectionType,omitempty" xml:"SecurityProtectionType,omitempty" type:"Repeated"`
}

func (s DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageSecurityProtectionTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageSecurityProtectionTypes) GoString() string {
	return s.String()
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageSecurityProtectionTypes) GetSecurityProtectionType() []*string {
	return s.SecurityProtectionType
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageSecurityProtectionTypes) SetSecurityProtectionType(v []*string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageSecurityProtectionTypes {
	s.SecurityProtectionType = v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageSecurityProtectionTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTags struct {
	Tag []*DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTags) GoString() string {
	return s.String()
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTags) GetTag() []*DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTagsTag {
	return s.Tag
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTags) SetTag(v []*DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTagsTag) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTags {
	s.Tag = v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTags) Validate() error {
	return dara.Validate(s)
}

type DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTagsTag struct {
	// The tag key that is added to the Internet Shared Bandwidth instance.
	//
	// example:
	//
	// KeyTest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that is added to the Internet Shared Bandwidth instance.
	//
	// example:
	//
	// ValueTest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTagsTag) SetKey(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTagsTag) SetValue(v string) *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTagsTag) Validate() error {
	return dara.Validate(s)
}
