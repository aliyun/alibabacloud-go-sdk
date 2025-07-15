// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkPackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkPackages(v []*DescribeNetworkPackagesResponseBodyNetworkPackages) *DescribeNetworkPackagesResponseBody
	GetNetworkPackages() []*DescribeNetworkPackagesResponseBodyNetworkPackages
	SetNextToken(v string) *DescribeNetworkPackagesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeNetworkPackagesResponseBody
	GetRequestId() *string
}

type DescribeNetworkPackagesResponseBody struct {
	// The premium bandwidth plans.
	NetworkPackages []*DescribeNetworkPackagesResponseBodyNetworkPackages `json:"NetworkPackages,omitempty" xml:"NetworkPackages,omitempty" type:"Repeated"`
	// The token that is used to start the next query. If the value of this parameter is empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNetworkPackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesResponseBody) GetNetworkPackages() []*DescribeNetworkPackagesResponseBodyNetworkPackages {
	return s.NetworkPackages
}

func (s *DescribeNetworkPackagesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNetworkPackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkPackagesResponseBody) SetNetworkPackages(v []*DescribeNetworkPackagesResponseBodyNetworkPackages) *DescribeNetworkPackagesResponseBody {
	s.NetworkPackages = v
	return s
}

func (s *DescribeNetworkPackagesResponseBody) SetNextToken(v string) *DescribeNetworkPackagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBody) SetRequestId(v string) *DescribeNetworkPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkPackagesResponseBodyNetworkPackages struct {
	// The bandwidth provided by the premium bandwidth plan. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The business status.
	//
	// Valid values:
	//
	// 	- Expired
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Normal
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The time when the premium bandwidth plan was created.
	//
	// example:
	//
	// 2021-05-10T02:35:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The public egress IP address of the premium bandwidth plan.
	EipAddresses []*string `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty" type:"Repeated"`
	// The time when the premium bandwidth plan expires.
	//
	// 	- If the plan is a subscription one, the time when the plan expires is returned.
	//
	// 	- If the plan is a pay-as-you-go one, `2099-12-31T15:59:59Z` is returned.
	//
	// example:
	//
	// 2099-12-31T15:59:59Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The charge type of the premium bandwidth plan.
	//
	// 	- Valid value when the `PayType` parameter is set to `PrePaid`:
	//
	//     	- PayByBandwidth: charges by fixed bandwidth.
	//
	// 	- Valid values when the `PayType` parameter is set to `PostPaid`:
	//
	//     	- PayByTraffic: charges by data transfer.
	//
	//     	- PayByBandwidth: charges by fixed bandwidth.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The ID of the premium bandwidth plan.
	//
	// example:
	//
	// np-amtp8e8q1o9e4****
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	// The status of the premium bandwidth plan.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Released
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- InUse
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Releasing
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// InUse
	NetworkPackageStatus *string `json:"NetworkPackageStatus,omitempty" xml:"NetworkPackageStatus,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The office network name.
	//
	// example:
	//
	// test
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The type of the office network.
	//
	// Valid values:
	//
	// 	- standard: advanced office network
	//
	// 	- customized: custom office network
	//
	// 	- basic: basic office network
	//
	// example:
	//
	// basic
	OfficeSiteVpcType *string `json:"OfficeSiteVpcType,omitempty" xml:"OfficeSiteVpcType,omitempty"`
	// The billing method of the premium bandwidth plan.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go
	//
	// 	- PrePaid: subscription
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The time when the reserved network bandwidth took effect.
	//
	// example:
	//
	// 2021-07-10T00:00:00Z
	ReservationActiveTime *string `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	// The peak bandwidth that is reserved for the premium bandwidth plan. Unit: Mbit/s.
	//
	// example:
	//
	// 20
	ReservationBandwidth *int32 `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	// The billing method of the reserved network bandwidth.
	//
	// Valid values:
	//
	// 	- PayByTraffic: charges by data transfer.
	//
	// 	- PayByBandwidth: charges by fixed bandwidth.
	//
	// example:
	//
	// PayByBandwidth
	ReservationInternetChargeType *string `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
}

func (s DescribeNetworkPackagesResponseBodyNetworkPackages) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkPackagesResponseBodyNetworkPackages) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetEipAddresses() []*string {
	return s.EipAddresses
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetNetworkPackageId() *string {
	return s.NetworkPackageId
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetNetworkPackageStatus() *string {
	return s.NetworkPackageStatus
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetOfficeSiteVpcType() *string {
	return s.OfficeSiteVpcType
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetPayType() *string {
	return s.PayType
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetReservationActiveTime() *string {
	return s.ReservationActiveTime
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetReservationBandwidth() *int32 {
	return s.ReservationBandwidth
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetReservationInternetChargeType() *string {
	return s.ReservationInternetChargeType
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetBandwidth(v int32) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.Bandwidth = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetBusinessStatus(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetCreateTime(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.CreateTime = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetEipAddresses(v []*string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.EipAddresses = v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetExpiredTime(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetInternetChargeType(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetNetworkPackageId(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.NetworkPackageId = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetNetworkPackageStatus(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.NetworkPackageStatus = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetOfficeSiteId(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetOfficeSiteName(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetOfficeSiteVpcType(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.OfficeSiteVpcType = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetPayType(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.PayType = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetReservationActiveTime(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.ReservationActiveTime = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetReservationBandwidth(v int32) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.ReservationBandwidth = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetReservationInternetChargeType(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.ReservationInternetChargeType = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) Validate() error {
	return dara.Validate(s)
}
