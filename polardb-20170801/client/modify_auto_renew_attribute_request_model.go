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
	// The cloud service provider to which the instance belongs.
	//
	// example:
	//
	// ENS
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// The IDs of the database clusters. The billing method of the clusters must be subscription. If you call this operation for a pay-as-you-go cluster, the OperationDenied.DBClusterPayType error may be returned. You can first call the TransformDBClusterPayType operation to change the billing method. Usage notes: The billing method of the cluster must be subscription. Calling this operation for a pay-as-you-go cluster returns the OperationDenied.DBClusterPayType error.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// The auto-renewal duration for the instance. Valid values:
	//
	// - If **PeriodUnit*	- is set to **Month**, valid values are `1`, `2`, `3`, `6`, and `12`.
	//
	// - If **PeriodUnit*	- is set to **Year**, valid values are `1`, `2`, and `3`.
	//
	//
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The unit of the renewal duration. Valid values:
	//
	// - **Year**: Year.
	//
	// - **Month**: Month.
	//
	//
	//
	// Default value: **Month**.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID. The value can be up to 50 characters in length.
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
	// - **AutoRenewal**: Auto-renewal.
	//
	// - **Normal**: Manual renewal.
	//
	// - **NotRenewal**: No renewal.
	//
	// Default value: **AutoRenewal**.
	//
	// > If you set this parameter to **NotRenewal**, the system no longer sends expiration reminders and only sends a non-renewal reminder three days before expiration.
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
