// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoRenewalAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenewalPeriod(v string) *ModifyAutoRenewalAttributeRequest
	GetAutoRenewalPeriod() *string
	SetAutoRenewalPeriodUnit(v string) *ModifyAutoRenewalAttributeRequest
	GetAutoRenewalPeriodUnit() *string
	SetAutoRenewalStatus(v string) *ModifyAutoRenewalAttributeRequest
	GetAutoRenewalStatus() *string
	SetDBClusterId(v string) *ModifyAutoRenewalAttributeRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyAutoRenewalAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAutoRenewalAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyAutoRenewalAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyAutoRenewalAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAutoRenewalAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyAutoRenewalAttributeRequest struct {
	// The duration of the auto-renewal. Default value: 1. Valid values:
	//
	// 	- When **AutoRenewalPeriod*	- is set to **Month**, the value ranges from 1 to 11 (integer).
	//
	// 	- When **AutoRenewalPeriod*	- is set to **Month**, the valid values are 1, 2, 3, and 5 (integer).
	//
	// >  Longer renewal periods offer better pricing. Renewing for 1 year is more cost-effective than renewing for 10 or 11 months.
	//
	// example:
	//
	// 1
	AutoRenewalPeriod *string `json:"AutoRenewalPeriod,omitempty" xml:"AutoRenewalPeriod,omitempty"`
	// Auto-renewal duration. Valid values:
	//
	// 	- Year.
	//
	// 	- Month.
	//
	// example:
	//
	// Year
	AutoRenewalPeriodUnit *string `json:"AutoRenewalPeriodUnit,omitempty" xml:"AutoRenewalPeriodUnit,omitempty"`
	// The renewal method. Valid values:
	//
	// 	- **AutoRenewal**: The cluster is automatically renewed.
	//
	// 	- **Normal**: The cluster is manually renewed. Before the cluster expires, the system sends you a reminder by SMS message.
	//
	// 	- **NotRenewal**: The cluster is not renewed. Reminders are only sent three days before cluster expiration.
	//
	// example:
	//
	// AutoRenewal
	AutoRenewalStatus *string `json:"AutoRenewalStatus,omitempty" xml:"AutoRenewalStatus,omitempty"`
	// The ID of cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-uf6pl56w1j8h****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
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

func (s ModifyAutoRenewalAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRenewalAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewalAttributeRequest) GetAutoRenewalPeriod() *string {
	return s.AutoRenewalPeriod
}

func (s *ModifyAutoRenewalAttributeRequest) GetAutoRenewalPeriodUnit() *string {
	return s.AutoRenewalPeriodUnit
}

func (s *ModifyAutoRenewalAttributeRequest) GetAutoRenewalStatus() *string {
	return s.AutoRenewalStatus
}

func (s *ModifyAutoRenewalAttributeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAutoRenewalAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAutoRenewalAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAutoRenewalAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAutoRenewalAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAutoRenewalAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAutoRenewalAttributeRequest) SetAutoRenewalPeriod(v string) *ModifyAutoRenewalAttributeRequest {
	s.AutoRenewalPeriod = &v
	return s
}

func (s *ModifyAutoRenewalAttributeRequest) SetAutoRenewalPeriodUnit(v string) *ModifyAutoRenewalAttributeRequest {
	s.AutoRenewalPeriodUnit = &v
	return s
}

func (s *ModifyAutoRenewalAttributeRequest) SetAutoRenewalStatus(v string) *ModifyAutoRenewalAttributeRequest {
	s.AutoRenewalStatus = &v
	return s
}

func (s *ModifyAutoRenewalAttributeRequest) SetDBClusterId(v string) *ModifyAutoRenewalAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAutoRenewalAttributeRequest) SetOwnerAccount(v string) *ModifyAutoRenewalAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAutoRenewalAttributeRequest) SetOwnerId(v int64) *ModifyAutoRenewalAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAutoRenewalAttributeRequest) SetRegionId(v string) *ModifyAutoRenewalAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAutoRenewalAttributeRequest) SetResourceOwnerAccount(v string) *ModifyAutoRenewalAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAutoRenewalAttributeRequest) SetResourceOwnerId(v int64) *ModifyAutoRenewalAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAutoRenewalAttributeRequest) Validate() error {
	return dara.Validate(s)
}
