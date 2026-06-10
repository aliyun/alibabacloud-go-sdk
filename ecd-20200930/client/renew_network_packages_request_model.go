// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewNetworkPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewNetworkPackagesRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RenewNetworkPackagesRequest
	GetAutoRenew() *bool
	SetNetworkPackageId(v []*string) *RenewNetworkPackagesRequest
	GetNetworkPackageId() []*string
	SetPeriod(v int32) *RenewNetworkPackagesRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewNetworkPackagesRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *RenewNetworkPackagesRequest
	GetPromotionId() *string
	SetRegionId(v string) *RenewNetworkPackagesRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *RenewNetworkPackagesRequest
	GetResellerOwnerUid() *int64
}

type RenewNetworkPackagesRequest struct {
	// Whether to enable automatic payment.
	//
	// example:
	//
	// true
	AutoPay   *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// A list of premium public bandwidth IDs. You can specify 1 to 100 IDs.
	//
	// This parameter is required.
	NetworkPackageId []*string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" type:"Repeated"`
	// The renewal duration. Valid values depend on the value of `PeriodUnit`.
	//
	// - If `PeriodUnit` is `Week`, valid values are: 1.
	//
	// - If `PeriodUnit` is `Month`, valid values are: 1, 2, 3, or 6.
	//
	// - If `PeriodUnit` is `Year`, valid values are: 1, 2, or 3.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit for the renewal duration.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 500038360030606
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. Call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to list regions that support WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s RenewNetworkPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewNetworkPackagesRequest) GoString() string {
	return s.String()
}

func (s *RenewNetworkPackagesRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewNetworkPackagesRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewNetworkPackagesRequest) GetNetworkPackageId() []*string {
	return s.NetworkPackageId
}

func (s *RenewNetworkPackagesRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewNetworkPackagesRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewNetworkPackagesRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *RenewNetworkPackagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RenewNetworkPackagesRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *RenewNetworkPackagesRequest) SetAutoPay(v bool) *RenewNetworkPackagesRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetAutoRenew(v bool) *RenewNetworkPackagesRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetNetworkPackageId(v []*string) *RenewNetworkPackagesRequest {
	s.NetworkPackageId = v
	return s
}

func (s *RenewNetworkPackagesRequest) SetPeriod(v int32) *RenewNetworkPackagesRequest {
	s.Period = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetPeriodUnit(v string) *RenewNetworkPackagesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetPromotionId(v string) *RenewNetworkPackagesRequest {
	s.PromotionId = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetRegionId(v string) *RenewNetworkPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *RenewNetworkPackagesRequest) SetResellerOwnerUid(v int64) *RenewNetworkPackagesRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *RenewNetworkPackagesRequest) Validate() error {
	return dara.Validate(s)
}
