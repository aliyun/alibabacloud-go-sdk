// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallBackThirdRightSendPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizGroup(v string) *CallBackThirdRightSendPlanRequest
	GetBizGroup() *string
	SetBizType(v string) *CallBackThirdRightSendPlanRequest
	GetBizType() *string
	SetCardType(v int32) *CallBackThirdRightSendPlanRequest
	GetCardType() *int32
	SetErrorMsg(v string) *CallBackThirdRightSendPlanRequest
	GetErrorMsg() *string
	SetExtendInfo(v map[string]interface{}) *CallBackThirdRightSendPlanRequest
	GetExtendInfo() map[string]interface{}
	SetGenieOpenId(v string) *CallBackThirdRightSendPlanRequest
	GetGenieOpenId() *string
	SetReceiveStatus(v int32) *CallBackThirdRightSendPlanRequest
	GetReceiveStatus() *int32
	SetSn(v string) *CallBackThirdRightSendPlanRequest
	GetSn() *string
	SetSupplierId(v int64) *CallBackThirdRightSendPlanRequest
	GetSupplierId() *int64
}

type CallBackThirdRightSendPlanRequest struct {
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
	ExtendInfo map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
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

func (s CallBackThirdRightSendPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CallBackThirdRightSendPlanRequest) GoString() string {
	return s.String()
}

func (s *CallBackThirdRightSendPlanRequest) GetBizGroup() *string {
	return s.BizGroup
}

func (s *CallBackThirdRightSendPlanRequest) GetBizType() *string {
	return s.BizType
}

func (s *CallBackThirdRightSendPlanRequest) GetCardType() *int32 {
	return s.CardType
}

func (s *CallBackThirdRightSendPlanRequest) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CallBackThirdRightSendPlanRequest) GetExtendInfo() map[string]interface{} {
	return s.ExtendInfo
}

func (s *CallBackThirdRightSendPlanRequest) GetGenieOpenId() *string {
	return s.GenieOpenId
}

func (s *CallBackThirdRightSendPlanRequest) GetReceiveStatus() *int32 {
	return s.ReceiveStatus
}

func (s *CallBackThirdRightSendPlanRequest) GetSn() *string {
	return s.Sn
}

func (s *CallBackThirdRightSendPlanRequest) GetSupplierId() *int64 {
	return s.SupplierId
}

func (s *CallBackThirdRightSendPlanRequest) SetBizGroup(v string) *CallBackThirdRightSendPlanRequest {
	s.BizGroup = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetBizType(v string) *CallBackThirdRightSendPlanRequest {
	s.BizType = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetCardType(v int32) *CallBackThirdRightSendPlanRequest {
	s.CardType = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetErrorMsg(v string) *CallBackThirdRightSendPlanRequest {
	s.ErrorMsg = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetExtendInfo(v map[string]interface{}) *CallBackThirdRightSendPlanRequest {
	s.ExtendInfo = v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetGenieOpenId(v string) *CallBackThirdRightSendPlanRequest {
	s.GenieOpenId = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetReceiveStatus(v int32) *CallBackThirdRightSendPlanRequest {
	s.ReceiveStatus = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetSn(v string) *CallBackThirdRightSendPlanRequest {
	s.Sn = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetSupplierId(v int64) *CallBackThirdRightSendPlanRequest {
	s.SupplierId = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) Validate() error {
	return dara.Validate(s)
}
