// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenBandwidthPackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCenBandwidthPackages(v *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) *DescribeCenBandwidthPackagesResponseBody
	GetCenBandwidthPackages() *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages
	SetPageNumber(v int32) *DescribeCenBandwidthPackagesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenBandwidthPackagesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCenBandwidthPackagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCenBandwidthPackagesResponseBody
	GetTotalCount() *int32
}

type DescribeCenBandwidthPackagesResponseBody struct {
	// The details about the bandwidth plan.
	CenBandwidthPackages *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages `json:"CenBandwidthPackages,omitempty" xml:"CenBandwidthPackages,omitempty" type:"Struct"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9D7E2400-2755-4AF5-9B73-12565E4F73A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenBandwidthPackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBody) GetCenBandwidthPackages() *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	return s.CenBandwidthPackages
}

func (s *DescribeCenBandwidthPackagesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenBandwidthPackagesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenBandwidthPackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCenBandwidthPackagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCenBandwidthPackagesResponseBody) SetCenBandwidthPackages(v *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) *DescribeCenBandwidthPackagesResponseBody {
	s.CenBandwidthPackages = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBody) SetPageNumber(v int32) *DescribeCenBandwidthPackagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBody) SetPageSize(v int32) *DescribeCenBandwidthPackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBody) SetRequestId(v string) *DescribeCenBandwidthPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBody) SetTotalCount(v int32) *DescribeCenBandwidthPackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages struct {
	CenBandwidthPackage []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage `json:"CenBandwidthPackage,omitempty" xml:"CenBandwidthPackage,omitempty" type:"Repeated"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) GetCenBandwidthPackage() []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	return s.CenBandwidthPackage
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) SetCenBandwidthPackage(v []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages {
	s.CenBandwidthPackage = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackages) Validate() error {
	return dara.Validate(s)
}

type DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage struct {
	// The maximum bandwidth of the bandwidth plan.
	//
	// example:
	//
	// 2
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The billing method of the bandwidth plan.
	//
	// example:
	//
	// PREPAY
	BandwidthPackageChargeType *string `json:"BandwidthPackageChargeType,omitempty" xml:"BandwidthPackageChargeType,omitempty"`
	// The status of the bandwidth plan. Valid values:
	//
	// 	- **Normal**: running as expected.
	//
	// 	- **FinancialLocked**: locked due to overdue payments.
	//
	// 	- **SecurityLocked**: locked due to security reasons
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The ID of the bandwidth plan.
	//
	// example:
	//
	// cenbwp-4c2zaavbvh5x****
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	// A list of CEN instances that are associated with the bandwidth plan.
	CenIds *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds `json:"CenIds,omitempty" xml:"CenIds,omitempty" type:"Struct"`
	// The time when the bandwidth plan was created. The time is displayed in the ISO8601 standard in the YYYY-MM-DDThh:mmZ format.
	//
	// example:
	//
	// 2021-02-01T11:14Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the bandwidth plan.
	//
	// example:
	//
	// cen
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the bandwidth plan expires.
	//
	// example:
	//
	// 2021-09-08T16:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the area that you want to query. Valid values:
	//
	// 	- **china**: Chinese mainland.
	//
	// 	- **asia-pacific**: Asia Pacific
	//
	// 	- **europe**: Europe
	//
	// 	- **north-america**: North America
	//
	// example:
	//
	// china
	GeographicRegionAId *string `json:"GeographicRegionAId,omitempty" xml:"GeographicRegionAId,omitempty"`
	// The ID of the other area connected by the bandwidth plan. Valid values:
	//
	// 	- **china**: Chinese mainland.
	//
	// 	- **asia-pacific**: Asia Pacific
	//
	// 	- **europe**: Europe
	//
	// 	- **north-america**: North America
	//
	// example:
	//
	// north-america
	GeographicRegionBId *string `json:"GeographicRegionBId,omitempty" xml:"GeographicRegionBId,omitempty"`
	// The ID of the connected area.
	//
	// example:
	//
	// north-america_china
	GeographicSpanId *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	// Indicates whether renewal data is included.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter returns **true*	- only when the **IncludeReservationData*	- parameter is set to **true*	- and a pending order exists.
	//
	// example:
	//
	// false
	HasReservationData *string `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
	// Indicates whether the bandwidth plan supports cross-border communication.
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// false
	IsCrossBorder *bool `json:"IsCrossBorder,omitempty" xml:"IsCrossBorder,omitempty"`
	// The name of the bandwidth plan.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The details about the connected regions.
	OrginInterRegionBandwidthLimits *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits `json:"OrginInterRegionBandwidthLimits,omitempty" xml:"OrginInterRegionBandwidthLimits,omitempty" type:"Struct"`
	// The expiration time of the temporary upgrade.
	//
	// example:
	//
	// 2021-08-30T16:00Z
	ReservationActiveTime *string `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	// The bandwidth value to which the bandwidth plan is restored when the temporary upgrade ends.
	//
	// example:
	//
	// 10
	ReservationBandwidth *string `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	// The new billing method.
	//
	// example:
	//
	// PREPAY
	ReservationInternetChargeType *string `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	// The renewal method.
	//
	// 	- **TEMP_UPGRADE**: temporary upgrade
	//
	// 	- **UPGRADE**: upgrade
	//
	// example:
	//
	// UPGRADE
	ReservationOrderType *string `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	// The ID of the resource group to which the ACL belongs.
	//
	// example:
	//
	// rg-aekzoyr5k36****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the bandwidth plan is associated with a CEN instance.
	//
	// 	- **Idle**
	//
	// 	- **InUse**
	//
	// example:
	//
	// InUse
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the bandwidth plan.
	Tags *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetBandwidthPackageChargeType() *string {
	return s.BandwidthPackageChargeType
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetCenBandwidthPackageId() *string {
	return s.CenBandwidthPackageId
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetCenIds() *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds {
	return s.CenIds
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetDescription() *string {
	return s.Description
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetGeographicRegionAId() *string {
	return s.GeographicRegionAId
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetGeographicRegionBId() *string {
	return s.GeographicRegionBId
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetGeographicSpanId() *string {
	return s.GeographicSpanId
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetHasReservationData() *string {
	return s.HasReservationData
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetIsCrossBorder() *bool {
	return s.IsCrossBorder
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetName() *string {
	return s.Name
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetOrginInterRegionBandwidthLimits() *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits {
	return s.OrginInterRegionBandwidthLimits
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetReservationActiveTime() *string {
	return s.ReservationActiveTime
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetReservationBandwidth() *string {
	return s.ReservationBandwidth
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetReservationInternetChargeType() *string {
	return s.ReservationInternetChargeType
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetReservationOrderType() *string {
	return s.ReservationOrderType
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetStatus() *string {
	return s.Status
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) GetTags() *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTags {
	return s.Tags
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetBandwidth(v int64) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetBandwidthPackageChargeType(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.BandwidthPackageChargeType = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetBusinessStatus(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetCenBandwidthPackageId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetCenIds(v *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.CenIds = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetCreationTime(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.CreationTime = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetDescription(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.Description = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetExpiredTime(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetGeographicRegionAId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.GeographicRegionAId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetGeographicRegionBId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.GeographicRegionBId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetGeographicSpanId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetHasReservationData(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.HasReservationData = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetIsCrossBorder(v bool) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.IsCrossBorder = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetName(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.Name = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetOrginInterRegionBandwidthLimits(v *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.OrginInterRegionBandwidthLimits = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetReservationActiveTime(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ReservationActiveTime = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetReservationBandwidth(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ReservationBandwidth = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetReservationInternetChargeType(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ReservationInternetChargeType = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetReservationOrderType(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ReservationOrderType = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetResourceGroupId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetStatus(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.Status = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) SetTags(v *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTags) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage {
	s.Tags = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackage) Validate() error {
	return dara.Validate(s)
}

type DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds struct {
	CenId []*string `json:"CenId,omitempty" xml:"CenId,omitempty" type:"Repeated"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds) GetCenId() []*string {
	return s.CenId
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds) SetCenId(v []*string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds {
	s.CenId = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageCenIds) Validate() error {
	return dara.Validate(s)
}

type DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits struct {
	OrginInterRegionBandwidthLimit []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit `json:"OrginInterRegionBandwidthLimit,omitempty" xml:"OrginInterRegionBandwidthLimit,omitempty" type:"Repeated"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits) GetOrginInterRegionBandwidthLimit() []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit {
	return s.OrginInterRegionBandwidthLimit
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits) SetOrginInterRegionBandwidthLimit(v []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits {
	s.OrginInterRegionBandwidthLimit = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimits) Validate() error {
	return dara.Validate(s)
}

type DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit struct {
	// The maximum bandwidth value for the inter-region connection.
	//
	// example:
	//
	// 1
	BandwidthLimit *string `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	// The connected regions.
	//
	// example:
	//
	// north-america_china
	GeographicSpanId *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	// The ID of the local region.
	//
	// example:
	//
	// cn-hangzhou
	LocalRegionId *string `json:"LocalRegionId,omitempty" xml:"LocalRegionId,omitempty"`
	// The ID of the peer region.
	//
	// example:
	//
	// us-west-1
	OppositeRegionId *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) GetBandwidthLimit() *string {
	return s.BandwidthLimit
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) GetGeographicSpanId() *string {
	return s.GeographicSpanId
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) GetLocalRegionId() *string {
	return s.LocalRegionId
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) GetOppositeRegionId() *string {
	return s.OppositeRegionId
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) SetBandwidthLimit(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit {
	s.BandwidthLimit = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) SetGeographicSpanId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) SetLocalRegionId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit {
	s.LocalRegionId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) SetOppositeRegionId(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit {
	s.OppositeRegionId = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageOrginInterRegionBandwidthLimitsOrginInterRegionBandwidthLimit) Validate() error {
	return dara.Validate(s)
}

type DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTags struct {
	Tag []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTags) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTags) GetTag() []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTagsTag {
	return s.Tag
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTags) SetTag(v []*DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTagsTag) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTags {
	s.Tag = v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTags) Validate() error {
	return dara.Validate(s)
}

type DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTagsTag) SetKey(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTagsTag) SetValue(v string) *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeCenBandwidthPackagesResponseBodyCenBandwidthPackagesCenBandwidthPackageTagsTag) Validate() error {
	return dara.Validate(s)
}
