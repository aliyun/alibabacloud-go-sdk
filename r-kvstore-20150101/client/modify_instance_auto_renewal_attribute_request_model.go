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
	SetProduct(v string) *ModifyInstanceAutoRenewalAttributeRequest
	GetProduct() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceAutoRenewalAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceAutoRenewalAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyInstanceAutoRenewalAttributeRequest
	GetSecurityToken() *string
}

type ModifyInstanceAutoRenewalAttributeRequest struct {
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  The default value is **false**.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the instance. Separate multiple instance IDs with commas (,).
	//
	// > You can specify up to 30 instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The auto-renewal period. Valid values: **1*	- to **12**. Unit: months. When the instance is about to expire, the instance is automatically renewed based on the number of months specified by this parameter.
	//
	// > This parameter is available and required only if the **AutoRenew*	- parameter is set to **true**.
	//
	// example:
	//
	// 3
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The service. Set the value to kvstore.
	//
	// example:
	//
	// kvstore
	Product              *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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

func (s *ModifyInstanceAutoRenewalAttributeRequest) GetProduct() *string {
	return s.Product
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) GetSecurityToken() *string {
	return s.SecurityToken
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

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetProduct(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.Product = &v
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

func (s *ModifyInstanceAutoRenewalAttributeRequest) SetSecurityToken(v string) *ModifyInstanceAutoRenewalAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeRequest) Validate() error {
	return dara.Validate(s)
}
