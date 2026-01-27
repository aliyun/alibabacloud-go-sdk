// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCDiskChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v string) *ModifyRCDiskChargeTypeRequest
	GetAutoRenew() *string
	SetAutoUseCoupon(v bool) *ModifyRCDiskChargeTypeRequest
	GetAutoUseCoupon() *bool
	SetBusinessInfo(v string) *ModifyRCDiskChargeTypeRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *ModifyRCDiskChargeTypeRequest
	GetClientToken() *string
	SetInstanceId(v string) *ModifyRCDiskChargeTypeRequest
	GetInstanceId() *string
	SetPayType(v string) *ModifyRCDiskChargeTypeRequest
	GetPayType() *string
	SetPromotionCode(v string) *ModifyRCDiskChargeTypeRequest
	GetPromotionCode() *string
	SetRegionId(v string) *ModifyRCDiskChargeTypeRequest
	GetRegionId() *string
}

type ModifyRCDiskChargeTypeRequest struct {
	// example:
	//
	// true
	AutoRenew     *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoUseCoupon *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// example:
	//
	// None
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// example:
	//
	// 2F20251224*
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rcd-pq2091s13go9bkb04*
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 72802442****
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyRCDiskChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCDiskChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCDiskChargeTypeRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *ModifyRCDiskChargeTypeRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyRCDiskChargeTypeRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *ModifyRCDiskChargeTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyRCDiskChargeTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyRCDiskChargeTypeRequest) GetPayType() *string {
	return s.PayType
}

func (s *ModifyRCDiskChargeTypeRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyRCDiskChargeTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCDiskChargeTypeRequest) SetAutoRenew(v string) *ModifyRCDiskChargeTypeRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyRCDiskChargeTypeRequest) SetAutoUseCoupon(v bool) *ModifyRCDiskChargeTypeRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyRCDiskChargeTypeRequest) SetBusinessInfo(v string) *ModifyRCDiskChargeTypeRequest {
	s.BusinessInfo = &v
	return s
}

func (s *ModifyRCDiskChargeTypeRequest) SetClientToken(v string) *ModifyRCDiskChargeTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyRCDiskChargeTypeRequest) SetInstanceId(v string) *ModifyRCDiskChargeTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRCDiskChargeTypeRequest) SetPayType(v string) *ModifyRCDiskChargeTypeRequest {
	s.PayType = &v
	return s
}

func (s *ModifyRCDiskChargeTypeRequest) SetPromotionCode(v string) *ModifyRCDiskChargeTypeRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyRCDiskChargeTypeRequest) SetRegionId(v string) *ModifyRCDiskChargeTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCDiskChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
