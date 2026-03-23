// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformDBInstancePayTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v string) *TransformDBInstancePayTypeRequest
	GetAutoRenew() *string
	SetAutoUseCoupon(v bool) *TransformDBInstancePayTypeRequest
	GetAutoUseCoupon() *bool
	SetBusinessInfo(v string) *TransformDBInstancePayTypeRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *TransformDBInstancePayTypeRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *TransformDBInstancePayTypeRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *TransformDBInstancePayTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TransformDBInstancePayTypeRequest
	GetOwnerId() *int64
	SetPayType(v string) *TransformDBInstancePayTypeRequest
	GetPayType() *string
	SetPeriod(v string) *TransformDBInstancePayTypeRequest
	GetPeriod() *string
	SetPromotionCode(v string) *TransformDBInstancePayTypeRequest
	GetPromotionCode() *string
	SetResourceOwnerAccount(v string) *TransformDBInstancePayTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TransformDBInstancePayTypeRequest
	GetResourceOwnerId() *int64
	SetUsedTime(v int32) *TransformDBInstancePayTypeRequest
	GetUsedTime() *int32
}

type TransformDBInstancePayTypeRequest struct {
	AutoRenew     *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoUseCoupon *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	BusinessInfo  *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	PromotionCode        *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UsedTime             *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s TransformDBInstancePayTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s TransformDBInstancePayTypeRequest) GoString() string {
	return s.String()
}

func (s *TransformDBInstancePayTypeRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *TransformDBInstancePayTypeRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *TransformDBInstancePayTypeRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *TransformDBInstancePayTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *TransformDBInstancePayTypeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *TransformDBInstancePayTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TransformDBInstancePayTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TransformDBInstancePayTypeRequest) GetPayType() *string {
	return s.PayType
}

func (s *TransformDBInstancePayTypeRequest) GetPeriod() *string {
	return s.Period
}

func (s *TransformDBInstancePayTypeRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *TransformDBInstancePayTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TransformDBInstancePayTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TransformDBInstancePayTypeRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *TransformDBInstancePayTypeRequest) SetAutoRenew(v string) *TransformDBInstancePayTypeRequest {
	s.AutoRenew = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetAutoUseCoupon(v bool) *TransformDBInstancePayTypeRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetBusinessInfo(v string) *TransformDBInstancePayTypeRequest {
	s.BusinessInfo = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetClientToken(v string) *TransformDBInstancePayTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetDBInstanceId(v string) *TransformDBInstancePayTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetOwnerAccount(v string) *TransformDBInstancePayTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetOwnerId(v int64) *TransformDBInstancePayTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetPayType(v string) *TransformDBInstancePayTypeRequest {
	s.PayType = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetPeriod(v string) *TransformDBInstancePayTypeRequest {
	s.Period = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetPromotionCode(v string) *TransformDBInstancePayTypeRequest {
	s.PromotionCode = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetResourceOwnerAccount(v string) *TransformDBInstancePayTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetResourceOwnerId(v int64) *TransformDBInstancePayTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) SetUsedTime(v int32) *TransformDBInstancePayTypeRequest {
	s.UsedTime = &v
	return s
}

func (s *TransformDBInstancePayTypeRequest) Validate() error {
	return dara.Validate(s)
}
