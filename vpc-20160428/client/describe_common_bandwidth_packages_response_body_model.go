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
	if s.CommonBandwidthPackages != nil {
		if err := s.CommonBandwidthPackages.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.CommonBandwidthPackage != nil {
		for _, item := range s.CommonBandwidthPackage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackage struct {
	Bandwidth                     *string                                                                                                          `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthPackageId            *string                                                                                                          `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BizType                       *string                                                                                                          `json:"BizType,omitempty" xml:"BizType,omitempty"`
	BusinessStatus                *string                                                                                                          `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	CreationTime                  *string                                                                                                          `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DeletionProtection            *bool                                                                                                            `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	Description                   *string                                                                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime                   *string                                                                                                          `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	HasReservationData            *string                                                                                                          `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
	ISP                           *string                                                                                                          `json:"ISP,omitempty" xml:"ISP,omitempty"`
	InstanceChargeType            *string                                                                                                          `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InternetChargeType            *string                                                                                                          `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	Name                          *string                                                                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	PublicIpAddresses             *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddresses       `json:"PublicIpAddresses,omitempty" xml:"PublicIpAddresses,omitempty" type:"Struct"`
	Ratio                         *int32                                                                                                           `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	RegionId                      *string                                                                                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReservationActiveTime         *string                                                                                                          `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	ReservationBandwidth          *string                                                                                                          `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	ReservationInternetChargeType *string                                                                                                          `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	ReservationOrderType          *string                                                                                                          `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	ResourceGroupId               *string                                                                                                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityProtectionTypes       *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageSecurityProtectionTypes `json:"SecurityProtectionTypes,omitempty" xml:"SecurityProtectionTypes,omitempty" type:"Struct"`
	ServiceManaged                *int32                                                                                                           `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	Status                        *string                                                                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                          *DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTags                    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Zone                          *string                                                                                                          `json:"Zone,omitempty" xml:"Zone,omitempty"`
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
	if s.PublicIpAddresses != nil {
		if err := s.PublicIpAddresses.Validate(); err != nil {
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
	if s.PublicIpAddresse != nil {
		for _, item := range s.PublicIpAddresse {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackagePublicIpAddressesPublicIpAddresse struct {
	AllocationId                     *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	BandwidthPackageIpRelationStatus *string `json:"BandwidthPackageIpRelationStatus,omitempty" xml:"BandwidthPackageIpRelationStatus,omitempty"`
	IpAddress                        *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
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

type DescribeCommonBandwidthPackagesResponseBodyCommonBandwidthPackagesCommonBandwidthPackageTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
