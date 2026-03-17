// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeSmartAccessGatewaySoftwareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *UpgradeSmartAccessGatewaySoftwareRequest
	GetAutoPay() *bool
	SetDataPlan(v int64) *UpgradeSmartAccessGatewaySoftwareRequest
	GetDataPlan() *int64
	SetOwnerAccount(v string) *UpgradeSmartAccessGatewaySoftwareRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpgradeSmartAccessGatewaySoftwareRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpgradeSmartAccessGatewaySoftwareRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpgradeSmartAccessGatewaySoftwareRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeSmartAccessGatewaySoftwareRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *UpgradeSmartAccessGatewaySoftwareRequest
	GetSmartAGId() *string
	SetUserCount(v int32) *UpgradeSmartAccessGatewaySoftwareRequest
	GetUserCount() *int32
}

type UpgradeSmartAccessGatewaySoftwareRequest struct {
	// Specifies whether to enable auto-payment for the instance.
	//
	// 	- **false**: no
	//
	// 	- **true**: yes
	//
	// >  If the parameter is set to false, you must complete the payment in the SAG console after you call this operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The data transfer plan for each client account. Unit: GB.
	//
	// >  Each client account has a data transfer plan of 5 GB that is free of charge each month.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	DataPlan     *int64  `json:"DataPlan,omitempty" xml:"DataPlan,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG app instance is created.
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
	// sag-8biez7habqwmx6****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The maximum number of client accounts supported by the SAG app instance.
	//
	// After you complete the payment, you can create a client account for each employee who needs to use the SAG app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s UpgradeSmartAccessGatewaySoftwareRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeSmartAccessGatewaySoftwareRequest) GoString() string {
	return s.String()
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) GetDataPlan() *int64 {
	return s.DataPlan
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) GetUserCount() *int32 {
	return s.UserCount
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) SetAutoPay(v bool) *UpgradeSmartAccessGatewaySoftwareRequest {
	s.AutoPay = &v
	return s
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) SetDataPlan(v int64) *UpgradeSmartAccessGatewaySoftwareRequest {
	s.DataPlan = &v
	return s
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) SetOwnerAccount(v string) *UpgradeSmartAccessGatewaySoftwareRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) SetOwnerId(v int64) *UpgradeSmartAccessGatewaySoftwareRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) SetRegionId(v string) *UpgradeSmartAccessGatewaySoftwareRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) SetResourceOwnerAccount(v string) *UpgradeSmartAccessGatewaySoftwareRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) SetResourceOwnerId(v int64) *UpgradeSmartAccessGatewaySoftwareRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) SetSmartAGId(v string) *UpgradeSmartAccessGatewaySoftwareRequest {
	s.SmartAGId = &v
	return s
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) SetUserCount(v int32) *UpgradeSmartAccessGatewaySoftwareRequest {
	s.UserCount = &v
	return s
}

func (s *UpgradeSmartAccessGatewaySoftwareRequest) Validate() error {
	return dara.Validate(s)
}
