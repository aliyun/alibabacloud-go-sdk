// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenewalPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessInfo(v string) *DescribeRenewalPriceRequest
	GetBusinessInfo() *string
	SetCouponNo(v string) *DescribeRenewalPriceRequest
	GetCouponNo() *string
	SetDBInstanceId(v string) *DescribeRenewalPriceRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeRenewalPriceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRenewalPriceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeRenewalPriceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRenewalPriceRequest
	GetResourceOwnerId() *int64
}

type DescribeRenewalPriceRequest struct {
	// The business information. This is an additional parameter.
	//
	// example:
	//
	// {“ActivityId":"000000000"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The coupon code. Default value: **youhuiquan_promotion_option_id_for_blank**.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp12c5b040dc****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRenewalPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *DescribeRenewalPriceRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *DescribeRenewalPriceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeRenewalPriceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRenewalPriceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRenewalPriceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRenewalPriceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRenewalPriceRequest) SetBusinessInfo(v string) *DescribeRenewalPriceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetCouponNo(v string) *DescribeRenewalPriceRequest {
	s.CouponNo = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetDBInstanceId(v string) *DescribeRenewalPriceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetOwnerAccount(v string) *DescribeRenewalPriceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetOwnerId(v int64) *DescribeRenewalPriceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetResourceOwnerAccount(v string) *DescribeRenewalPriceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetResourceOwnerId(v int64) *DescribeRenewalPriceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) Validate() error {
	return dara.Validate(s)
}
