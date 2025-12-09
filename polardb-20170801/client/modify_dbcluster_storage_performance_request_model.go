// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterStoragePerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUseCoupon(v bool) *ModifyDBClusterStoragePerformanceRequest
	GetAutoUseCoupon() *bool
	SetBurstingEnabled(v string) *ModifyDBClusterStoragePerformanceRequest
	GetBurstingEnabled() *string
	SetClientToken(v string) *ModifyDBClusterStoragePerformanceRequest
	GetClientToken() *string
	SetDBClusterId(v string) *ModifyDBClusterStoragePerformanceRequest
	GetDBClusterId() *string
	SetModifyType(v string) *ModifyDBClusterStoragePerformanceRequest
	GetModifyType() *string
	SetPromotionCode(v string) *ModifyDBClusterStoragePerformanceRequest
	GetPromotionCode() *string
	SetProvisionedIops(v int32) *ModifyDBClusterStoragePerformanceRequest
	GetProvisionedIops() *int32
	SetResourceOwnerId(v int64) *ModifyDBClusterStoragePerformanceRequest
	GetResourceOwnerId() *int64
	SetStorageType(v string) *ModifyDBClusterStoragePerformanceRequest
	GetStorageType() *string
}

type ModifyDBClusterStoragePerformanceRequest struct {
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// Specifies whether to enable the I/O Burst feature for the ESSD AutoPL disk. Valid value:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >  This parameter is available only when the StorageType parameter is set to ESSDAUTOPL.
	//
	// example:
	//
	// false
	BurstingEnabled *string `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// example:
	//
	// 6000170000591aed949d0f******************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// Upgrade
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// example:
	//
	// 727xxxxxx934
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// example:
	//
	// 2500
	ProvisionedIops *int32 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// ESSDAUTOPL
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s ModifyDBClusterStoragePerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterStoragePerformanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterStoragePerformanceRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyDBClusterStoragePerformanceRequest) GetBurstingEnabled() *string {
	return s.BurstingEnabled
}

func (s *ModifyDBClusterStoragePerformanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBClusterStoragePerformanceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterStoragePerformanceRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *ModifyDBClusterStoragePerformanceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyDBClusterStoragePerformanceRequest) GetProvisionedIops() *int32 {
	return s.ProvisionedIops
}

func (s *ModifyDBClusterStoragePerformanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterStoragePerformanceRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *ModifyDBClusterStoragePerformanceRequest) SetAutoUseCoupon(v bool) *ModifyDBClusterStoragePerformanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyDBClusterStoragePerformanceRequest) SetBurstingEnabled(v string) *ModifyDBClusterStoragePerformanceRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyDBClusterStoragePerformanceRequest) SetClientToken(v string) *ModifyDBClusterStoragePerformanceRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBClusterStoragePerformanceRequest) SetDBClusterId(v string) *ModifyDBClusterStoragePerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterStoragePerformanceRequest) SetModifyType(v string) *ModifyDBClusterStoragePerformanceRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyDBClusterStoragePerformanceRequest) SetPromotionCode(v string) *ModifyDBClusterStoragePerformanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyDBClusterStoragePerformanceRequest) SetProvisionedIops(v int32) *ModifyDBClusterStoragePerformanceRequest {
	s.ProvisionedIops = &v
	return s
}

func (s *ModifyDBClusterStoragePerformanceRequest) SetResourceOwnerId(v int64) *ModifyDBClusterStoragePerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterStoragePerformanceRequest) SetStorageType(v string) *ModifyDBClusterStoragePerformanceRequest {
	s.StorageType = &v
	return s
}

func (s *ModifyDBClusterStoragePerformanceRequest) Validate() error {
	return dara.Validate(s)
}
