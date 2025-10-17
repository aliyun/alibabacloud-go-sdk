// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudProvider(v string) *ModifyAutoRenewAttributeRequest
	GetCloudProvider() *string
	SetDBClusterIds(v string) *ModifyAutoRenewAttributeRequest
	GetDBClusterIds() *string
	SetDuration(v string) *ModifyAutoRenewAttributeRequest
	GetDuration() *string
	SetOwnerAccount(v string) *ModifyAutoRenewAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAutoRenewAttributeRequest
	GetOwnerId() *int64
	SetPeriodUnit(v string) *ModifyAutoRenewAttributeRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *ModifyAutoRenewAttributeRequest
	GetRegionId() *string
	SetRenewalStatus(v string) *ModifyAutoRenewAttributeRequest
	GetRenewalStatus() *string
	SetResourceGroupId(v string) *ModifyAutoRenewAttributeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyAutoRenewAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAutoRenewAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyAutoRenewAttributeRequest struct {
	// example:
	//
	// ENS
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// The cluster ID. If you need to specify multiple cluster IDs, separate the cluster IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// The automatic renewal period.
	//
	// 	- Valid values when you set the **PeriodUnit*	- parameter to **Month**: `1, 2, 3, 6, and 12`.
	//
	// 	- Valid values when you set the **PeriodUnit*	- parameter to **Year**: `1, 2, and 3`.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The unit of the renewal period. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// Default value: **Month**.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the region. The region ID can be up to 50 characters in length.
	//
	// cn-hangzhou
	//
	//
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The auto-renewal status of the cluster. Valid values:
	//
	// 	- **AutoRenewal:*	- The cluster is automatically renewed.
	//
	// 	- **Normal**: The cluster is manually renewed.
	//
	// 	- **NotRenewal:*	- The cluster is not renewed after expiration.
	//
	// Default value: **AutoRenewal**.
	//
	// >  If you set this parameter to **NotRenewal**, the system sends a notification that indicates the cluster is not renewed three days before the cluster expires. After the cluster expires, the system no longer sends a notification.
	//
	// example:
	//
	// AutoRenewal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeRequest) GetCloudProvider() *string {
	return s.CloudProvider
}

func (s *ModifyAutoRenewAttributeRequest) GetDBClusterIds() *string {
	return s.DBClusterIds
}

func (s *ModifyAutoRenewAttributeRequest) GetDuration() *string {
	return s.Duration
}

func (s *ModifyAutoRenewAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAutoRenewAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAutoRenewAttributeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyAutoRenewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAutoRenewAttributeRequest) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *ModifyAutoRenewAttributeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyAutoRenewAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAutoRenewAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAutoRenewAttributeRequest) SetCloudProvider(v string) *ModifyAutoRenewAttributeRequest {
	s.CloudProvider = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetDBClusterIds(v string) *ModifyAutoRenewAttributeRequest {
	s.DBClusterIds = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetDuration(v string) *ModifyAutoRenewAttributeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetOwnerAccount(v string) *ModifyAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetOwnerId(v int64) *ModifyAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetPeriodUnit(v string) *ModifyAutoRenewAttributeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetRegionId(v string) *ModifyAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetRenewalStatus(v string) *ModifyAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetResourceGroupId(v string) *ModifyAutoRenewAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *ModifyAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *ModifyAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) Validate() error {
	return dara.Validate(s)
}
