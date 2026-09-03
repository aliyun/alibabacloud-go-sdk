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
	// The list of premium Internet bandwidth plans.
	NetworkPackages []*DescribeNetworkPackagesResponseBodyNetworkPackages `json:"NetworkPackages,omitempty" xml:"NetworkPackages,omitempty" type:"Repeated"`
	// The pagination token for the next query. If NextToken is empty, no more pages exist.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
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
	if s.NetworkPackages != nil {
		for _, item := range s.NetworkPackages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkPackagesResponseBodyNetworkPackages struct {
	// The bandwidth of the premium Internet bandwidth plan. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The business status.
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The time when the plan was created. The time is in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format (UTC).
	//
	// example:
	//
	// 2021-05-10T02:35:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The public egress IP address of the premium Internet bandwidth plan.
	EipAddresses []*string `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty" type:"Repeated"`
	// The expiration time of the premium Internet bandwidth plan.
	//
	// - If the plan uses the subscription billing method, the actual expiration time is returned.
	//
	// - If the plan uses the pay-as-you-go billing method, `2099-12-31T15:59:59Z` is returned.
	//
	// The time is in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format (UTC).
	//
	// example:
	//
	// 2099-12-31T15:59:59Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The billing method of the premium Internet bandwidth plan.
	//
	// - If the parameter `PayType` is set to `PrePaid`, the valid value is:
	//
	//     - PayByBandwidth: billing by fixed bandwidth.
	//
	// - If the parameter `PayType` is set to `PostPaid`, valid values are:
	//
	//     - PayByTraffic: billing by data transfer.
	//
	//     - PayByBandwidth: billing by fixed bandwidth.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The ID of the premium Internet bandwidth plan.
	//
	// example:
	//
	// np-amtp8e8q1o9e4****
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	// The status of the premium Internet bandwidth plan.
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
	// default
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The office network type.
	//
	// example:
	//
	// basic
	OfficeSiteVpcType *string `json:"OfficeSiteVpcType,omitempty" xml:"OfficeSiteVpcType,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The effective period of the reserved network bandwidth. The time is in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format (UTC).
	//
	// example:
	//
	// 2021-07-10T00:00:00Z
	ReservationActiveTime *string `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	// The peak reserved network bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 20
	ReservationBandwidth *int32 `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	// The billing method of the reserved network bandwidth.
	//
	// example:
	//
	// PayByBandwidth
	ReservationInternetChargeType *string `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	// The tags.
	Tags []*DescribeNetworkPackagesResponseBodyNetworkPackagesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) GetTags() []*DescribeNetworkPackagesResponseBodyNetworkPackagesTags {
	return s.Tags
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

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetTags(v []*DescribeNetworkPackagesResponseBodyNetworkPackagesTags) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.Tags = v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkPackagesResponseBodyNetworkPackagesTags struct {
	// The tag key. If you specify this parameter, the value cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNetworkPackagesResponseBodyNetworkPackagesTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkPackagesResponseBodyNetworkPackagesTags) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackagesTags) GetKey() *string {
	return s.Key
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackagesTags) GetValue() *string {
	return s.Value
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackagesTags) SetKey(v string) *DescribeNetworkPackagesResponseBodyNetworkPackagesTags {
	s.Key = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackagesTags) SetValue(v string) *DescribeNetworkPackagesResponseBodyNetworkPackagesTags {
	s.Value = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackagesTags) Validate() error {
	return dara.Validate(s)
}
