// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAutoRenewalAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v string) *ModifyInstanceAutoRenewalAttributeRequest
	GetAutoRenew() *string
	SetDBInstanceId(v string) *ModifyInstanceAutoRenewalAttributeRequest
	GetDBInstanceId() *string
	SetDuration(v string) *ModifyInstanceAutoRenewalAttributeRequest
	GetDuration() *string
	SetOwnerAccount(v string) *ModifyInstanceAutoRenewalAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceAutoRenewalAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyInstanceAutoRenewalAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceAutoRenewalAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceAutoRenewalAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyInstanceAutoRenewalAttributeRequest struct {
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  If this parameter is set to **true**, you must set the **Duration*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp15da1923e3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The auto-renewal period. Valid values: **1*	- to **12**. Unit: month.
	//
	// >  This parameter is valid only when **AutoRenew*	- is set to **true**.
	//
	// example:
	//
	// 1
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to query the region ID of the instance.
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

func (s ModifyInstanceAutoRenewalAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAutoRenewalAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) GetDuration() *string {
	return s.Duration
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetAutoRenew(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetDBInstanceId(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetDuration(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetOwnerAccount(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetOwnerId(v int64) *ModifyInstanceAutoRenewalAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetRegionId(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetResourceOwnerId(v int64) *ModifyInstanceAutoRenewalAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) Validate() error {
	return dara.Validate(s)
}
