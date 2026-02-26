// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundOrderCmd interface {
	dara.Model
	String() string
	GoString() string
	SetApplyReasonTextId(v int64) *RefundOrderCmd
	GetApplyReasonTextId() *int64
	SetApplyReasonTips(v string) *RefundOrderCmd
	GetApplyReasonTips() *string
	SetApplyRefundCount(v int32) *RefundOrderCmd
	GetApplyRefundCount() *int32
	SetApplyRefundFee(v int64) *RefundOrderCmd
	GetApplyRefundFee() *int64
	SetBizClaimType(v int32) *RefundOrderCmd
	GetBizClaimType() *int32
	SetGoodsStatus(v int32) *RefundOrderCmd
	GetGoodsStatus() *int32
	SetLeaveMessage(v string) *RefundOrderCmd
	GetLeaveMessage() *string
	SetLeavePictureLists(v []*LeavePictureList) *RefundOrderCmd
	GetLeavePictureLists() []*LeavePictureList
	SetOrderLineId(v string) *RefundOrderCmd
	GetOrderLineId() *string
}

type RefundOrderCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 47821
	ApplyReasonTextId *int64  `json:"applyReasonTextId,omitempty" xml:"applyReasonTextId,omitempty"`
	ApplyReasonTips   *string `json:"applyReasonTips,omitempty" xml:"applyReasonTips,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ApplyRefundCount *int32 `json:"applyRefundCount,omitempty" xml:"applyRefundCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	ApplyRefundFee *int64 `json:"applyRefundFee,omitempty" xml:"applyRefundFee,omitempty"`
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
	// 1
	GoodsStatus       *int32              `json:"goodsStatus,omitempty" xml:"goodsStatus,omitempty"`
	LeaveMessage      *string             `json:"leaveMessage,omitempty" xml:"leaveMessage,omitempty"`
	LeavePictureLists []*LeavePictureList `json:"leavePictureLists,omitempty" xml:"leavePictureLists,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 6692****5458
	OrderLineId *string `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
}

func (s RefundOrderCmd) String() string {
	return dara.Prettify(s)
}

func (s RefundOrderCmd) GoString() string {
	return s.String()
}

func (s *RefundOrderCmd) GetApplyReasonTextId() *int64 {
	return s.ApplyReasonTextId
}

func (s *RefundOrderCmd) GetApplyReasonTips() *string {
	return s.ApplyReasonTips
}

func (s *RefundOrderCmd) GetApplyRefundCount() *int32 {
	return s.ApplyRefundCount
}

func (s *RefundOrderCmd) GetApplyRefundFee() *int64 {
	return s.ApplyRefundFee
}

func (s *RefundOrderCmd) GetBizClaimType() *int32 {
	return s.BizClaimType
}

func (s *RefundOrderCmd) GetGoodsStatus() *int32 {
	return s.GoodsStatus
}

func (s *RefundOrderCmd) GetLeaveMessage() *string {
	return s.LeaveMessage
}

func (s *RefundOrderCmd) GetLeavePictureLists() []*LeavePictureList {
	return s.LeavePictureLists
}

func (s *RefundOrderCmd) GetOrderLineId() *string {
	return s.OrderLineId
}

func (s *RefundOrderCmd) SetApplyReasonTextId(v int64) *RefundOrderCmd {
	s.ApplyReasonTextId = &v
	return s
}

func (s *RefundOrderCmd) SetApplyReasonTips(v string) *RefundOrderCmd {
	s.ApplyReasonTips = &v
	return s
}

func (s *RefundOrderCmd) SetApplyRefundCount(v int32) *RefundOrderCmd {
	s.ApplyRefundCount = &v
	return s
}

func (s *RefundOrderCmd) SetApplyRefundFee(v int64) *RefundOrderCmd {
	s.ApplyRefundFee = &v
	return s
}

func (s *RefundOrderCmd) SetBizClaimType(v int32) *RefundOrderCmd {
	s.BizClaimType = &v
	return s
}

func (s *RefundOrderCmd) SetGoodsStatus(v int32) *RefundOrderCmd {
	s.GoodsStatus = &v
	return s
}

func (s *RefundOrderCmd) SetLeaveMessage(v string) *RefundOrderCmd {
	s.LeaveMessage = &v
	return s
}

func (s *RefundOrderCmd) SetLeavePictureLists(v []*LeavePictureList) *RefundOrderCmd {
	s.LeavePictureLists = v
	return s
}

func (s *RefundOrderCmd) SetOrderLineId(v string) *RefundOrderCmd {
	s.OrderLineId = &v
	return s
}

func (s *RefundOrderCmd) Validate() error {
	if s.LeavePictureLists != nil {
		for _, item := range s.LeavePictureLists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
