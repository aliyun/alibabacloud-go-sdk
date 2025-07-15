// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudDriveServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateCloudDriveServiceRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateCloudDriveServiceRequest
	GetAutoRenew() *bool
	SetBizType(v int32) *CreateCloudDriveServiceRequest
	GetBizType() *int32
	SetCdsChargeType(v string) *CreateCloudDriveServiceRequest
	GetCdsChargeType() *string
	SetCenId(v string) *CreateCloudDriveServiceRequest
	GetCenId() *string
	SetDomainName(v string) *CreateCloudDriveServiceRequest
	GetDomainName() *string
	SetEndUserId(v []*string) *CreateCloudDriveServiceRequest
	GetEndUserId() []*string
	SetMaxSize(v int64) *CreateCloudDriveServiceRequest
	GetMaxSize() *int64
	SetName(v string) *CreateCloudDriveServiceRequest
	GetName() *string
	SetOfficeSiteId(v string) *CreateCloudDriveServiceRequest
	GetOfficeSiteId() *string
	SetOfficeSiteType(v string) *CreateCloudDriveServiceRequest
	GetOfficeSiteType() *string
	SetPeriod(v int64) *CreateCloudDriveServiceRequest
	GetPeriod() *int64
	SetPeriodUnit(v string) *CreateCloudDriveServiceRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *CreateCloudDriveServiceRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *CreateCloudDriveServiceRequest
	GetResellerOwnerUid() *int64
	SetSolutionId(v string) *CreateCloudDriveServiceRequest
	GetSolutionId() *string
	SetUserCount(v int64) *CreateCloudDriveServiceRequest
	GetUserCount() *int64
	SetUserMaxSize(v int64) *CreateCloudDriveServiceRequest
	GetUserMaxSize() *int64
}

type CreateCloudDriveServiceRequest struct {
	// Specifies whether to enable the auto-payment feature.
	//
	// Valid values:
	//
	// 	- true: enables the auto-payment feature. Ensure your Alibaba Cloud account has sufficient balance. Insufficient balance may result in abnormal orders.
	//
	// 	- false (default): disables the auto-payment feature. The order is generated, but payment must be made manually. You can log on to the Alibaba Cloud Management Console and complete the payment based on the order ID on the Orders page.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Optional. Specifies whether to enable the auto-renewal feature. This parameter takes effect only if you set CdsChargeType to `Prepaid`.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 3
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The billing method of the enterprise drive.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// 	- PrePaid: subscription.
	//
	// example:
	//
	// PostPaid
	CdsChargeType *string `json:"CdsChargeType,omitempty" xml:"CdsChargeType,omitempty"`
	// The ID of the Cloud Enterprise Network (CEN) instance. This parameter takes effect only if you set `OfficeSiteType` to `AD_CONNECTOR`. If you have configured `OfficeSiteId`, you can leave this parameter empty.
	//
	// example:
	//
	// cen-g4ba1mkji8nj6****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The domain name of the enterprise AD office network. This parameter takes effect only if you set `OfficeSiteType` to `AD_CONNECTOR`. If you have configured `OfficeSiteId`, you can leave this parameter empty.
	//
	// example:
	//
	// test.local
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The user IDs.
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	// The maximum storage capacity of the enterprise drive.
	//
	// 	- For a pay-as-you-go enterprise drive, the unit is bytes.
	//
	// 	- For a subscription enterprise drive, the unit is GiB. For example, to create a 500 GiB subscription drive, set the value to 500 GiB. To create a 2 TiB subscription drive, set the value to 2048 GiB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 536870912000
	MaxSize *int64 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The name of the enterprise drive
	//
	// example:
	//
	// wuying-pds
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the office network. This parameter takes effect only if you set OfficeSiteType to `AD_CONNECTOR`.
	//
	// example:
	//
	// cn-hangzhou+dir-400695****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The type of the office network.
	//
	// Valid values:
	//
	// 	- SIMPLE: convenience office network.
	//
	// 	- AD_CONNECTOR: enterprise Active Directory (AD) office network.
	//
	// example:
	//
	// SIMPLE
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	// The subscription duration. The unit is specified by `PeriodUnit`. This parameter takes effect only if you set `CdsChargeType` to `PrePaid`.
	//
	// Valid values:
	//
	// 	- 1
	//
	// 	- 2
	//
	// 	- 3
	//
	// example:
	//
	// 1
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// Required. The unit of the subscription duration. This parameter takes effect only if you set `CdsChargeType` to `PrePaid`.
	//
	// Valid value:
	//
	// 	- Year
	//
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the list of regions where Enterprise Drive Service is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// co-0esnf8kb8zpfbqmvt
	SolutionId *string `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
	// Required. The maximum number of users allowed on the enterprise drive. This parameter takes effect only if you set `CdsChargeType` to `PrePaid`.
	//
	// Valid values:
	//
	// 	- 5 when the value of MaxSize is 500 GiB.
	//
	// 	- 20 when the value of MaxSize is 2048 GiB.
	//
	// 	- 50 when the value of MaxSize is 5120 GiB.
	//
	// example:
	//
	// 5
	UserCount *int64 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	// The maximum storage capacity of the user\\"s personal disk when allocated. Unit: bytes.
	//
	// example:
	//
	// 1024000
	UserMaxSize *int64 `json:"UserMaxSize,omitempty" xml:"UserMaxSize,omitempty"`
}

func (s CreateCloudDriveServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudDriveServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudDriveServiceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateCloudDriveServiceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateCloudDriveServiceRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *CreateCloudDriveServiceRequest) GetCdsChargeType() *string {
	return s.CdsChargeType
}

func (s *CreateCloudDriveServiceRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateCloudDriveServiceRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateCloudDriveServiceRequest) GetEndUserId() []*string {
	return s.EndUserId
}

func (s *CreateCloudDriveServiceRequest) GetMaxSize() *int64 {
	return s.MaxSize
}

func (s *CreateCloudDriveServiceRequest) GetName() *string {
	return s.Name
}

func (s *CreateCloudDriveServiceRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateCloudDriveServiceRequest) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *CreateCloudDriveServiceRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *CreateCloudDriveServiceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateCloudDriveServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCloudDriveServiceRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *CreateCloudDriveServiceRequest) GetSolutionId() *string {
	return s.SolutionId
}

func (s *CreateCloudDriveServiceRequest) GetUserCount() *int64 {
	return s.UserCount
}

func (s *CreateCloudDriveServiceRequest) GetUserMaxSize() *int64 {
	return s.UserMaxSize
}

func (s *CreateCloudDriveServiceRequest) SetAutoPay(v bool) *CreateCloudDriveServiceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetAutoRenew(v bool) *CreateCloudDriveServiceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetBizType(v int32) *CreateCloudDriveServiceRequest {
	s.BizType = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetCdsChargeType(v string) *CreateCloudDriveServiceRequest {
	s.CdsChargeType = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetCenId(v string) *CreateCloudDriveServiceRequest {
	s.CenId = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetDomainName(v string) *CreateCloudDriveServiceRequest {
	s.DomainName = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetEndUserId(v []*string) *CreateCloudDriveServiceRequest {
	s.EndUserId = v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetMaxSize(v int64) *CreateCloudDriveServiceRequest {
	s.MaxSize = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetName(v string) *CreateCloudDriveServiceRequest {
	s.Name = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetOfficeSiteId(v string) *CreateCloudDriveServiceRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetOfficeSiteType(v string) *CreateCloudDriveServiceRequest {
	s.OfficeSiteType = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetPeriod(v int64) *CreateCloudDriveServiceRequest {
	s.Period = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetPeriodUnit(v string) *CreateCloudDriveServiceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetRegionId(v string) *CreateCloudDriveServiceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetResellerOwnerUid(v int64) *CreateCloudDriveServiceRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetSolutionId(v string) *CreateCloudDriveServiceRequest {
	s.SolutionId = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetUserCount(v int64) *CreateCloudDriveServiceRequest {
	s.UserCount = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) SetUserMaxSize(v int64) *CreateCloudDriveServiceRequest {
	s.UserMaxSize = &v
	return s
}

func (s *CreateCloudDriveServiceRequest) Validate() error {
	return dara.Validate(s)
}
