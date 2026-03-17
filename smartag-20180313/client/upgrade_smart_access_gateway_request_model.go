// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeSmartAccessGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *UpgradeSmartAccessGatewayRequest
	GetAutoPay() *bool
	SetBandWidthSpec(v int64) *UpgradeSmartAccessGatewayRequest
	GetBandWidthSpec() *int64
	SetOwnerAccount(v string) *UpgradeSmartAccessGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpgradeSmartAccessGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpgradeSmartAccessGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpgradeSmartAccessGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeSmartAccessGatewayRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *UpgradeSmartAccessGatewayRequest
	GetSmartAGId() *string
}

type UpgradeSmartAccessGatewayRequest struct {
	// Indicates whether to automatically pay the bill for a subscription instance.
	//
	// Valid values: **true | false**. Default value: false
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The bandwidth of the SAG instance.
	//
	// 	- Value range for an SAG-100WM device: 2 to 50. Unit: Mbit/s
	//
	// 	- Value range for an SAG-1000 device: 10 to 500. Unit: Mbit/s
	//
	// example:
	//
	// 3
	BandWidthSpec *int64  `json:"BandWidthSpec,omitempty" xml:"BandWidthSpec,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the SAG instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-d3m51apgw4po******
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s UpgradeSmartAccessGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeSmartAccessGatewayRequest) GoString() string {
	return s.String()
}

func (s *UpgradeSmartAccessGatewayRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *UpgradeSmartAccessGatewayRequest) GetBandWidthSpec() *int64 {
	return s.BandWidthSpec
}

func (s *UpgradeSmartAccessGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpgradeSmartAccessGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeSmartAccessGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeSmartAccessGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeSmartAccessGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeSmartAccessGatewayRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *UpgradeSmartAccessGatewayRequest) SetAutoPay(v bool) *UpgradeSmartAccessGatewayRequest {
	s.AutoPay = &v
	return s
}

func (s *UpgradeSmartAccessGatewayRequest) SetBandWidthSpec(v int64) *UpgradeSmartAccessGatewayRequest {
	s.BandWidthSpec = &v
	return s
}

func (s *UpgradeSmartAccessGatewayRequest) SetOwnerAccount(v string) *UpgradeSmartAccessGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeSmartAccessGatewayRequest) SetOwnerId(v int64) *UpgradeSmartAccessGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeSmartAccessGatewayRequest) SetRegionId(v string) *UpgradeSmartAccessGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeSmartAccessGatewayRequest) SetResourceOwnerAccount(v string) *UpgradeSmartAccessGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeSmartAccessGatewayRequest) SetResourceOwnerId(v int64) *UpgradeSmartAccessGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeSmartAccessGatewayRequest) SetSmartAGId(v string) *UpgradeSmartAccessGatewayRequest {
	s.SmartAGId = &v
	return s
}

func (s *UpgradeSmartAccessGatewayRequest) Validate() error {
	return dara.Validate(s)
}
