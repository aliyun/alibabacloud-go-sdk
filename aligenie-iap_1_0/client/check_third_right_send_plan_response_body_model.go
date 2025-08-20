// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckThirdRightSendPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRetCode(v int32) *CheckThirdRightSendPlanResponseBody
	GetRetCode() *int32
	SetRetMsg(v string) *CheckThirdRightSendPlanResponseBody
	GetRetMsg() *string
	SetRetValue(v *CheckThirdRightSendPlanResponseBodyRetValue) *CheckThirdRightSendPlanResponseBody
	GetRetValue() *CheckThirdRightSendPlanResponseBodyRetValue
}

type CheckThirdRightSendPlanResponseBody struct {
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	// example:
	//
	// 系统异常
	RetMsg   *string                                      `json:"RetMsg,omitempty" xml:"RetMsg,omitempty"`
	RetValue *CheckThirdRightSendPlanResponseBodyRetValue `json:"RetValue,omitempty" xml:"RetValue,omitempty" type:"Struct"`
}

func (s CheckThirdRightSendPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckThirdRightSendPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CheckThirdRightSendPlanResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *CheckThirdRightSendPlanResponseBody) GetRetMsg() *string {
	return s.RetMsg
}

func (s *CheckThirdRightSendPlanResponseBody) GetRetValue() *CheckThirdRightSendPlanResponseBodyRetValue {
	return s.RetValue
}

func (s *CheckThirdRightSendPlanResponseBody) SetRetCode(v int32) *CheckThirdRightSendPlanResponseBody {
	s.RetCode = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBody) SetRetMsg(v string) *CheckThirdRightSendPlanResponseBody {
	s.RetMsg = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBody) SetRetValue(v *CheckThirdRightSendPlanResponseBodyRetValue) *CheckThirdRightSendPlanResponseBody {
	s.RetValue = v
	return s
}

func (s *CheckThirdRightSendPlanResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckThirdRightSendPlanResponseBodyRetValue struct {
	// example:
	//
	// "1713262192005"
	ActivateDate *string `json:"ActivateDate,omitempty" xml:"ActivateDate,omitempty"`
	// example:
	//
	// 1001 日卡 1002 月卡 1003 季卡 1004 年卡
	CardType *int32 `json:"CardType,omitempty" xml:"CardType,omitempty"`
	// example:
	//
	// TB
	ChannelCode *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	// example:
	//
	// 淘宝
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// {}
	ExtendInfo map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// example:
	//
	// 908FA068-529C-0C20-8DB5-63B0EF7CFF1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// "1713262192005"
	RightsExpiredDate *string `json:"RightsExpiredDate,omitempty" xml:"RightsExpiredDate,omitempty"`
}

func (s CheckThirdRightSendPlanResponseBodyRetValue) String() string {
	return dara.Prettify(s)
}

func (s CheckThirdRightSendPlanResponseBodyRetValue) GoString() string {
	return s.String()
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) GetActivateDate() *string {
	return s.ActivateDate
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) GetCardType() *int32 {
	return s.CardType
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) GetChannelCode() *string {
	return s.ChannelCode
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) GetChannelName() *string {
	return s.ChannelName
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) GetExtendInfo() map[string]interface{} {
	return s.ExtendInfo
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) GetRightsExpiredDate() *string {
	return s.RightsExpiredDate
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) SetActivateDate(v string) *CheckThirdRightSendPlanResponseBodyRetValue {
	s.ActivateDate = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) SetCardType(v int32) *CheckThirdRightSendPlanResponseBodyRetValue {
	s.CardType = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) SetChannelCode(v string) *CheckThirdRightSendPlanResponseBodyRetValue {
	s.ChannelCode = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) SetChannelName(v string) *CheckThirdRightSendPlanResponseBodyRetValue {
	s.ChannelName = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) SetExtendInfo(v map[string]interface{}) *CheckThirdRightSendPlanResponseBodyRetValue {
	s.ExtendInfo = v
	return s
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) SetRequestId(v string) *CheckThirdRightSendPlanResponseBodyRetValue {
	s.RequestId = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) SetRightsExpiredDate(v string) *CheckThirdRightSendPlanResponseBodyRetValue {
	s.RightsExpiredDate = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) Validate() error {
	return dara.Validate(s)
}
