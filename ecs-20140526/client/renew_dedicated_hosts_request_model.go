// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDedicatedHostsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RenewDedicatedHostsRequest
	GetClientToken() *string
	SetDedicatedHostIds(v string) *RenewDedicatedHostsRequest
	GetDedicatedHostIds() *string
	SetOwnerAccount(v string) *RenewDedicatedHostsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RenewDedicatedHostsRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *RenewDedicatedHostsRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewDedicatedHostsRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *RenewDedicatedHostsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RenewDedicatedHostsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RenewDedicatedHostsRequest
	GetResourceOwnerId() *int64
}

type RenewDedicatedHostsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of dedicated host IDs. You can specify up to 100 subscription dedicated host IDs. Specify multiple dedicated host IDs in a JSON array in the format of `["dh-xxxxxxxxx", "dh-yyyyyyyyy", … "dh-zzzzzzzzz"]`. Separate IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// dh-bp199lyny9b3****
	DedicatedHostIds *string `json:"DedicatedHostIds,omitempty" xml:"DedicatedHostIds,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The renewal period. Valid values:
	//
	// <props="china">
	//
	// - If PeriodUnit is set to Week: 1, 2, 3, and 4.
	//
	// - If PeriodUnit is set to Month: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// - If PeriodUnit is set to Year: 1, 2, 3, 4, and 5.
	//
	//
	//
	// <props="intl">
	//
	// - If PeriodUnit is set to Month: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// - If PeriodUnit is set to Year: 1, 2, 3, 4, and 5.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal period. Valid values:
	//
	// <props="china">
	//
	// - Week
	//
	// - Month
	//
	// - Year
	//
	//
	//
	// <props="intl">
	//
	// - Month
	//
	// - Year
	//
	//
	//
	// Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID of the dedicated host. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RenewDedicatedHostsRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewDedicatedHostsRequest) GoString() string {
	return s.String()
}

func (s *RenewDedicatedHostsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewDedicatedHostsRequest) GetDedicatedHostIds() *string {
	return s.DedicatedHostIds
}

func (s *RenewDedicatedHostsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RenewDedicatedHostsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewDedicatedHostsRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewDedicatedHostsRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewDedicatedHostsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RenewDedicatedHostsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RenewDedicatedHostsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RenewDedicatedHostsRequest) SetClientToken(v string) *RenewDedicatedHostsRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewDedicatedHostsRequest) SetDedicatedHostIds(v string) *RenewDedicatedHostsRequest {
	s.DedicatedHostIds = &v
	return s
}

func (s *RenewDedicatedHostsRequest) SetOwnerAccount(v string) *RenewDedicatedHostsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RenewDedicatedHostsRequest) SetOwnerId(v int64) *RenewDedicatedHostsRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewDedicatedHostsRequest) SetPeriod(v int32) *RenewDedicatedHostsRequest {
	s.Period = &v
	return s
}

func (s *RenewDedicatedHostsRequest) SetPeriodUnit(v string) *RenewDedicatedHostsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewDedicatedHostsRequest) SetRegionId(v string) *RenewDedicatedHostsRequest {
	s.RegionId = &v
	return s
}

func (s *RenewDedicatedHostsRequest) SetResourceOwnerAccount(v string) *RenewDedicatedHostsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewDedicatedHostsRequest) SetResourceOwnerId(v int64) *RenewDedicatedHostsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenewDedicatedHostsRequest) Validate() error {
	return dara.Validate(s)
}
