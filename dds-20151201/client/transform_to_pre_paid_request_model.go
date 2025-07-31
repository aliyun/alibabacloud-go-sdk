// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformToPrePaidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *TransformToPrePaidRequest
	GetAutoPay() *bool
	SetAutoRenew(v string) *TransformToPrePaidRequest
	GetAutoRenew() *string
	SetBusinessInfo(v string) *TransformToPrePaidRequest
	GetBusinessInfo() *string
	SetCouponNo(v string) *TransformToPrePaidRequest
	GetCouponNo() *string
	SetInstanceId(v string) *TransformToPrePaidRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *TransformToPrePaidRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TransformToPrePaidRequest
	GetOwnerId() *int64
	SetPeriod(v int64) *TransformToPrePaidRequest
	GetPeriod() *int64
	SetResourceOwnerAccount(v string) *TransformToPrePaidRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TransformToPrePaidRequest
	GetResourceOwnerId() *int64
}

type TransformToPrePaidRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true**: enables automatic payment.
	//
	// 	- **false**: disables automatic payment. For more information, see [Renew an ApsaraDB for MongoDB subscription instance](https://help.aliyun.com/document_detail/85052.html).
	//
	// >  Default value: **true**.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  Default value: **false**.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The business information. This is an additional parameter.
	//
	// example:
	//
	// {“ActivityId":"000000000"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The coupon code. Default value: `youhuiquan_promotion_option_id_for_blank`.
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
	// dds-bp1366caac83****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration of the instance. Unit: months. Valid values: **1**, **2**, **3**, **4**, **5**, **6**, **7**, **8**, **9**, **12**, **24**, and **36**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period               *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s TransformToPrePaidRequest) String() string {
	return dara.Prettify(s)
}

func (s TransformToPrePaidRequest) GoString() string {
	return s.String()
}

func (s *TransformToPrePaidRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *TransformToPrePaidRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *TransformToPrePaidRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *TransformToPrePaidRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *TransformToPrePaidRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TransformToPrePaidRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TransformToPrePaidRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TransformToPrePaidRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *TransformToPrePaidRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TransformToPrePaidRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TransformToPrePaidRequest) SetAutoPay(v bool) *TransformToPrePaidRequest {
	s.AutoPay = &v
	return s
}

func (s *TransformToPrePaidRequest) SetAutoRenew(v string) *TransformToPrePaidRequest {
	s.AutoRenew = &v
	return s
}

func (s *TransformToPrePaidRequest) SetBusinessInfo(v string) *TransformToPrePaidRequest {
	s.BusinessInfo = &v
	return s
}

func (s *TransformToPrePaidRequest) SetCouponNo(v string) *TransformToPrePaidRequest {
	s.CouponNo = &v
	return s
}

func (s *TransformToPrePaidRequest) SetInstanceId(v string) *TransformToPrePaidRequest {
	s.InstanceId = &v
	return s
}

func (s *TransformToPrePaidRequest) SetOwnerAccount(v string) *TransformToPrePaidRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TransformToPrePaidRequest) SetOwnerId(v int64) *TransformToPrePaidRequest {
	s.OwnerId = &v
	return s
}

func (s *TransformToPrePaidRequest) SetPeriod(v int64) *TransformToPrePaidRequest {
	s.Period = &v
	return s
}

func (s *TransformToPrePaidRequest) SetResourceOwnerAccount(v string) *TransformToPrePaidRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TransformToPrePaidRequest) SetResourceOwnerId(v int64) *TransformToPrePaidRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TransformToPrePaidRequest) Validate() error {
	return dara.Validate(s)
}
