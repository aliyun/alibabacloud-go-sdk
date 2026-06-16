// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDmsGatewayOrder interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DmsGatewayOrder
	GetBizType() *string
	SetChargeType(v string) *DmsGatewayOrder
	GetChargeType() *string
	SetCommodityCode(v string) *DmsGatewayOrder
	GetCommodityCode() *string
	SetExpireTime(v string) *DmsGatewayOrder
	GetExpireTime() *string
	SetInstanceId(v string) *DmsGatewayOrder
	GetInstanceId() *string
	SetInstanceType(v string) *DmsGatewayOrder
	GetInstanceType() *string
	SetOrderId(v int64) *DmsGatewayOrder
	GetOrderId() *int64
	SetPayNum(v int32) *DmsGatewayOrder
	GetPayNum() *int32
	SetRegion(v string) *DmsGatewayOrder
	GetRegion() *string
	SetState(v string) *DmsGatewayOrder
	GetState() *string
}

type DmsGatewayOrder struct {
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ChargeType    *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ExpireTime    *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType  *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OrderId       *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PayNum        *int32  `json:"PayNum,omitempty" xml:"PayNum,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DmsGatewayOrder) String() string {
	return dara.Prettify(s)
}

func (s DmsGatewayOrder) GoString() string {
	return s.String()
}

func (s *DmsGatewayOrder) GetBizType() *string {
	return s.BizType
}

func (s *DmsGatewayOrder) GetChargeType() *string {
	return s.ChargeType
}

func (s *DmsGatewayOrder) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DmsGatewayOrder) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DmsGatewayOrder) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DmsGatewayOrder) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DmsGatewayOrder) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DmsGatewayOrder) GetPayNum() *int32 {
	return s.PayNum
}

func (s *DmsGatewayOrder) GetRegion() *string {
	return s.Region
}

func (s *DmsGatewayOrder) GetState() *string {
	return s.State
}

func (s *DmsGatewayOrder) SetBizType(v string) *DmsGatewayOrder {
	s.BizType = &v
	return s
}

func (s *DmsGatewayOrder) SetChargeType(v string) *DmsGatewayOrder {
	s.ChargeType = &v
	return s
}

func (s *DmsGatewayOrder) SetCommodityCode(v string) *DmsGatewayOrder {
	s.CommodityCode = &v
	return s
}

func (s *DmsGatewayOrder) SetExpireTime(v string) *DmsGatewayOrder {
	s.ExpireTime = &v
	return s
}

func (s *DmsGatewayOrder) SetInstanceId(v string) *DmsGatewayOrder {
	s.InstanceId = &v
	return s
}

func (s *DmsGatewayOrder) SetInstanceType(v string) *DmsGatewayOrder {
	s.InstanceType = &v
	return s
}

func (s *DmsGatewayOrder) SetOrderId(v int64) *DmsGatewayOrder {
	s.OrderId = &v
	return s
}

func (s *DmsGatewayOrder) SetPayNum(v int32) *DmsGatewayOrder {
	s.PayNum = &v
	return s
}

func (s *DmsGatewayOrder) SetRegion(v string) *DmsGatewayOrder {
	s.Region = &v
	return s
}

func (s *DmsGatewayOrder) SetState(v string) *DmsGatewayOrder {
	s.State = &v
	return s
}

func (s *DmsGatewayOrder) Validate() error {
	return dara.Validate(s)
}
