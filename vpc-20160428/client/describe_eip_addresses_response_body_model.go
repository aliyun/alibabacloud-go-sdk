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
	AllocationId                  *string                                                                        `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	AllocationTime                *string                                                                        `json:"AllocationTime,omitempty" xml:"AllocationTime,omitempty"`
	Bandwidth                     *string                                                                        `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthPackageBandwidth     *string                                                                        `json:"BandwidthPackageBandwidth,omitempty" xml:"BandwidthPackageBandwidth,omitempty"`
	BandwidthPackageId            *string                                                                        `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BandwidthPackageType          *string                                                                        `json:"BandwidthPackageType,omitempty" xml:"BandwidthPackageType,omitempty"`
	BizType                       *string                                                                        `json:"BizType,omitempty" xml:"BizType,omitempty"`
	BusinessStatus                *string                                                                        `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	ChargeType                    *string                                                                        `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DeletionProtection            *bool                                                                          `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	Description                   *string                                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	EipBandwidth                  *string                                                                        `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	ExpiredTime                   *string                                                                        `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	HDMonitorStatus               *string                                                                        `json:"HDMonitorStatus,omitempty" xml:"HDMonitorStatus,omitempty"`
	HasReservationData            *string                                                                        `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
	ISP                           *string                                                                        `json:"ISP,omitempty" xml:"ISP,omitempty"`
	InstanceId                    *string                                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceRegionId              *string                                                                        `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	InstanceType                  *string                                                                        `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetChargeType            *string                                                                        `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	IpAddress                     *string                                                                        `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Mode                          *string                                                                        `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Name                          *string                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Netmode                       *string                                                                        `json:"Netmode,omitempty" xml:"Netmode,omitempty"`
	OperationLocks                *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks          `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	PrivateIpAddress              *string                                                                        `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	PublicIpAddressPoolId         *string                                                                        `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	RegionId                      *string                                                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReservationActiveTime         *string                                                                        `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	ReservationBandwidth          *string                                                                        `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	ReservationInternetChargeType *string                                                                        `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	ReservationOrderType          *string                                                                        `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	ResourceGroupId               *string                                                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecondLimited                 *bool                                                                          `json:"SecondLimited,omitempty" xml:"SecondLimited,omitempty"`
	SecurityProtectionTypes       *DescribeEipAddressesResponseBodyEipAddressesEipAddressSecurityProtectionTypes `json:"SecurityProtectionTypes,omitempty" xml:"SecurityProtectionTypes,omitempty" type:"Struct"`
	SegmentInstanceId             *string                                                                        `json:"SegmentInstanceId,omitempty" xml:"SegmentInstanceId,omitempty"`
	ServiceID                     *int64                                                                         `json:"ServiceID,omitempty" xml:"ServiceID,omitempty"`
	ServiceManaged                *int32                                                                         `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	Status                        *string                                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                          *DescribeEipAddressesResponseBodyEipAddressesEipAddressTags                    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VpcId                         *string                                                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Zone                          *string                                                                        `json:"Zone,omitempty" xml:"Zone,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
