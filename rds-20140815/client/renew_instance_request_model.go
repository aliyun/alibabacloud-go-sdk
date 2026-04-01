// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v string) *RenewInstanceRequest
	GetAutoPay() *string
	SetAutoRenew(v string) *RenewInstanceRequest
	GetAutoRenew() *string
	SetAutoUseCoupon(v bool) *RenewInstanceRequest
	GetAutoUseCoupon() *bool
	SetClientToken(v string) *RenewInstanceRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *RenewInstanceRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *RenewInstanceRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *RenewInstanceRequest
	GetPeriod() *int32
	SetPromotionCode(v string) *RenewInstanceRequest
	GetPromotionCode() *string
	SetResourceOwnerAccount(v string) *RenewInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RenewInstanceRequest
	GetResourceOwnerId() *int64
}

type RenewInstanceRequest struct {
	// Specifies whether to enable automatic payment during the renewal. Valid values:
	//
	// 	- **True**: enables automatic payment. Make sure that your Alibaba Cloud account has adequate balance.
	//
	// 	- **False*	- (default): disables automatic payment. You have to manually pay the order in the console.
	//
	// >  For more information about how to renew the instance in the console, see the following topics:
	//
	// 	- [Manually renew an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96050.html)
	//
	// 	- [Manually renew an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96741.html)
	//
	// 	- [Manually renew an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95637.html)
	//
	// 	- [Manually renew an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97122.html)
	//
	// example:
	//
	// True
	AutoPay *string `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// 	- **true**.
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// 	- **true**: uses a coupon.
	//
	// 	- **false*	- (default): does not use a coupon.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The duration of the subscription renewal. Unit: month. Valid values:
	//
	// 	- **1~9**
	//
	// 	- **12**
	//
	// 	- **24**
	//
	// 	- **36**
	//
	// 	- **48**
	//
	// 	- **60**
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 726702810223
	PromotionCode        *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) GetAutoPay() *string {
	return s.AutoPay
}

func (s *RenewInstanceRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *RenewInstanceRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *RenewInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RenewInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RenewInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewInstanceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *RenewInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RenewInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RenewInstanceRequest) SetAutoPay(v string) *RenewInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewInstanceRequest) SetAutoRenew(v string) *RenewInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewInstanceRequest) SetAutoUseCoupon(v bool) *RenewInstanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *RenewInstanceRequest) SetClientToken(v string) *RenewInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewInstanceRequest) SetDBInstanceId(v string) *RenewInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RenewInstanceRequest) SetOwnerId(v int64) *RenewInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewInstanceRequest) SetPeriod(v int32) *RenewInstanceRequest {
	s.Period = &v
	return s
}

func (s *RenewInstanceRequest) SetPromotionCode(v string) *RenewInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *RenewInstanceRequest) SetResourceOwnerAccount(v string) *RenewInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewInstanceRequest) SetResourceOwnerId(v int64) *RenewInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenewInstanceRequest) Validate() error {
	return dara.Validate(s)
}
