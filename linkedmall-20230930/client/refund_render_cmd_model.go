// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundRenderCmd interface {
	dara.Model
	String() string
	GoString() string
	SetBizClaimType(v int32) *RefundRenderCmd
	GetBizClaimType() *int32
	SetGoodsStatus(v int32) *RefundRenderCmd
	GetGoodsStatus() *int32
	SetOrderLineId(v string) *RefundRenderCmd
	GetOrderLineId() *string
}

type RefundRenderCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	BizClaimType *int32 `json:"bizClaimType,omitempty" xml:"bizClaimType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	GoodsStatus *int32 `json:"goodsStatus,omitempty" xml:"goodsStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6692****5458
	OrderLineId *string `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
}

func (s RefundRenderCmd) String() string {
	return dara.Prettify(s)
}

func (s RefundRenderCmd) GoString() string {
	return s.String()
}

func (s *RefundRenderCmd) GetBizClaimType() *int32 {
	return s.BizClaimType
}

func (s *RefundRenderCmd) GetGoodsStatus() *int32 {
	return s.GoodsStatus
}

func (s *RefundRenderCmd) GetOrderLineId() *string {
	return s.OrderLineId
}

func (s *RefundRenderCmd) SetBizClaimType(v int32) *RefundRenderCmd {
	s.BizClaimType = &v
	return s
}

func (s *RefundRenderCmd) SetGoodsStatus(v int32) *RefundRenderCmd {
	s.GoodsStatus = &v
	return s
}

func (s *RefundRenderCmd) SetOrderLineId(v string) *RefundRenderCmd {
	s.OrderLineId = &v
	return s
}

func (s *RefundRenderCmd) Validate() error {
	return dara.Validate(s)
}
