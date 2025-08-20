// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallBackThirdRightSendPlanShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizGroup(v string) *CallBackThirdRightSendPlanShrinkRequest
	GetBizGroup() *string
	SetBizType(v string) *CallBackThirdRightSendPlanShrinkRequest
	GetBizType() *string
	SetCardType(v int32) *CallBackThirdRightSendPlanShrinkRequest
	GetCardType() *int32
	SetErrorMsg(v string) *CallBackThirdRightSendPlanShrinkRequest
	GetErrorMsg() *string
	SetExtendInfoShrink(v string) *CallBackThirdRightSendPlanShrinkRequest
	GetExtendInfoShrink() *string
	SetGenieOpenId(v string) *CallBackThirdRightSendPlanShrinkRequest
	GetGenieOpenId() *string
	SetReceiveStatus(v int32) *CallBackThirdRightSendPlanShrinkRequest
	GetReceiveStatus() *int32
	SetSn(v string) *CallBackThirdRightSendPlanShrinkRequest
	GetSn() *string
	SetSupplierId(v int64) *CallBackThirdRightSendPlanShrinkRequest
	GetSupplierId() *int64
}

type CallBackThirdRightSendPlanShrinkRequest struct {
	// example:
	//
	// cc
	BizGroup *string `json:"BizGroup,omitempty" xml:"BizGroup,omitempty"`
	// example:
	//
	// ailabs
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 1001
	CardType *int32 `json:"CardType,omitempty" xml:"CardType,omitempty"`
	// example:
	//
	// 领取异常
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// {}
	ExtendInfoShrink *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// example:
	//
	// 1dsds2FzCXFGVA1ADS
	GenieOpenId *string `json:"GenieOpenId,omitempty" xml:"GenieOpenId,omitempty"`
	// example:
	//
	// 1
	ReceiveStatus *int32 `json:"ReceiveStatus,omitempty" xml:"ReceiveStatus,omitempty"`
	// example:
	//
	// 01000019100307010000600
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// example:
	//
	// 1
	SupplierId *int64 `json:"SupplierId,omitempty" xml:"SupplierId,omitempty"`
}

func (s CallBackThirdRightSendPlanShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CallBackThirdRightSendPlanShrinkRequest) GoString() string {
	return s.String()
}

func (s *CallBackThirdRightSendPlanShrinkRequest) GetBizGroup() *string {
	return s.BizGroup
}

func (s *CallBackThirdRightSendPlanShrinkRequest) GetBizType() *string {
	return s.BizType
}

func (s *CallBackThirdRightSendPlanShrinkRequest) GetCardType() *int32 {
	return s.CardType
}

func (s *CallBackThirdRightSendPlanShrinkRequest) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CallBackThirdRightSendPlanShrinkRequest) GetExtendInfoShrink() *string {
	return s.ExtendInfoShrink
}

func (s *CallBackThirdRightSendPlanShrinkRequest) GetGenieOpenId() *string {
	return s.GenieOpenId
}

func (s *CallBackThirdRightSendPlanShrinkRequest) GetReceiveStatus() *int32 {
	return s.ReceiveStatus
}

func (s *CallBackThirdRightSendPlanShrinkRequest) GetSn() *string {
	return s.Sn
}

func (s *CallBackThirdRightSendPlanShrinkRequest) GetSupplierId() *int64 {
	return s.SupplierId
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetBizGroup(v string) *CallBackThirdRightSendPlanShrinkRequest {
	s.BizGroup = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetBizType(v string) *CallBackThirdRightSendPlanShrinkRequest {
	s.BizType = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetCardType(v int32) *CallBackThirdRightSendPlanShrinkRequest {
	s.CardType = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetErrorMsg(v string) *CallBackThirdRightSendPlanShrinkRequest {
	s.ErrorMsg = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetExtendInfoShrink(v string) *CallBackThirdRightSendPlanShrinkRequest {
	s.ExtendInfoShrink = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetGenieOpenId(v string) *CallBackThirdRightSendPlanShrinkRequest {
	s.GenieOpenId = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetReceiveStatus(v int32) *CallBackThirdRightSendPlanShrinkRequest {
	s.ReceiveStatus = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetSn(v string) *CallBackThirdRightSendPlanShrinkRequest {
	s.Sn = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetSupplierId(v int64) *CallBackThirdRightSendPlanShrinkRequest {
	s.SupplierId = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) Validate() error {
	return dara.Validate(s)
}
