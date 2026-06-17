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
	// The cloud provider of the cluster.
	//
	// example:
	//
	// ENS
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// The ID of the database cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// The auto-renewal period. Valid values:
	//
	// - When **PeriodUnit*	- is **Month**, the valid values are `[1,2,3,6,12]`.
	//
	// - When **PeriodUnit*	- is **Year**, the valid values are `[1-3]`.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The unit of the auto-renewal period. Valid values:
	//
	// - **Year**
	//
	// - **Month**
	//
	// Default value: **Month**.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID. The ID can be up to 50 characters in length.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2319134.html) operation to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The auto-renewal status. Valid values:
	//
	// - **AutoRenewal**: Enables auto-renewal.
	//
	// - **Normal**: Enables manual renewal.
	//
	// - **NotRenewal**: Disables renewal.
	//
	// Default value: **AutoRenewal**.
	//
	// > When you set this parameter to **NotRenewal**, the system stops sending expiration reminders and sends only a non-renewal notification three days before the expiration date.
	//
	// example:
	//
	// AutoRenewal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	// The resource group ID.
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
