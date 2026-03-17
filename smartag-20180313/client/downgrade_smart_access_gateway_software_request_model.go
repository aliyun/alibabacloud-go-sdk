// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradeSmartAccessGatewaySoftwareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *DowngradeSmartAccessGatewaySoftwareRequest
	GetAutoPay() *bool
	SetDataPlan(v int64) *DowngradeSmartAccessGatewaySoftwareRequest
	GetDataPlan() *int64
	SetOwnerAccount(v string) *DowngradeSmartAccessGatewaySoftwareRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DowngradeSmartAccessGatewaySoftwareRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DowngradeSmartAccessGatewaySoftwareRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DowngradeSmartAccessGatewaySoftwareRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DowngradeSmartAccessGatewaySoftwareRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DowngradeSmartAccessGatewaySoftwareRequest
	GetSmartAGId() *string
	SetUserCount(v int32) *DowngradeSmartAccessGatewaySoftwareRequest
	GetUserCount() *int32
}

type DowngradeSmartAccessGatewaySoftwareRequest struct {
	// Specify whether the bill belongs to a subscription instance that has auto renewal enabled. Valid values:
	//
	// 	- **false*	- (default): no
	//
	// 	- **true**: yes
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The amount of free data transfer allocated to each client account per month, which is 5 GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	DataPlan     *int64  `json:"DataPlan,omitempty" xml:"DataPlan,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG app instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-n2uym2h45lnd31****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The quota of client accounts that can be connected to an SAG app instance. Typically, you need to create an account for each user that needs to log on to the SAG app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s DowngradeSmartAccessGatewaySoftwareRequest) String() string {
	return dara.Prettify(s)
}

func (s DowngradeSmartAccessGatewaySoftwareRequest) GoString() string {
	return s.String()
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) GetDataPlan() *int64 {
	return s.DataPlan
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) GetUserCount() *int32 {
	return s.UserCount
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) SetAutoPay(v bool) *DowngradeSmartAccessGatewaySoftwareRequest {
	s.AutoPay = &v
	return s
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) SetDataPlan(v int64) *DowngradeSmartAccessGatewaySoftwareRequest {
	s.DataPlan = &v
	return s
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) SetOwnerAccount(v string) *DowngradeSmartAccessGatewaySoftwareRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) SetOwnerId(v int64) *DowngradeSmartAccessGatewaySoftwareRequest {
	s.OwnerId = &v
	return s
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) SetRegionId(v string) *DowngradeSmartAccessGatewaySoftwareRequest {
	s.RegionId = &v
	return s
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) SetResourceOwnerAccount(v string) *DowngradeSmartAccessGatewaySoftwareRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) SetResourceOwnerId(v int64) *DowngradeSmartAccessGatewaySoftwareRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) SetSmartAGId(v string) *DowngradeSmartAccessGatewaySoftwareRequest {
	s.SmartAGId = &v
	return s
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) SetUserCount(v int32) *DowngradeSmartAccessGatewaySoftwareRequest {
	s.UserCount = &v
	return s
}

func (s *DowngradeSmartAccessGatewaySoftwareRequest) Validate() error {
	return dara.Validate(s)
}
