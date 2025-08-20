// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallBackThirdRightSendPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRetCode(v string) *CallBackThirdRightSendPlanResponseBody
	GetRetCode() *string
	SetRetMsg(v string) *CallBackThirdRightSendPlanResponseBody
	GetRetMsg() *string
	SetRetValue(v bool) *CallBackThirdRightSendPlanResponseBody
	GetRetValue() *bool
	SetRequestId(v string) *CallBackThirdRightSendPlanResponseBody
	GetRequestId() *string
}

type CallBackThirdRightSendPlanResponseBody struct {
	// example:
	//
	// 400
	RetCode *string `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	// example:
	//
	// 系统异常
	RetMsg   *string `json:"RetMsg,omitempty" xml:"RetMsg,omitempty"`
	RetValue *bool   `json:"RetValue,omitempty" xml:"RetValue,omitempty"`
	// example:
	//
	// 908FA068-529C-0C20-8DB5-63B0EF7CFF1F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CallBackThirdRightSendPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CallBackThirdRightSendPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CallBackThirdRightSendPlanResponseBody) GetRetCode() *string {
	return s.RetCode
}

func (s *CallBackThirdRightSendPlanResponseBody) GetRetMsg() *string {
	return s.RetMsg
}

func (s *CallBackThirdRightSendPlanResponseBody) GetRetValue() *bool {
	return s.RetValue
}

func (s *CallBackThirdRightSendPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CallBackThirdRightSendPlanResponseBody) SetRetCode(v string) *CallBackThirdRightSendPlanResponseBody {
	s.RetCode = &v
	return s
}

func (s *CallBackThirdRightSendPlanResponseBody) SetRetMsg(v string) *CallBackThirdRightSendPlanResponseBody {
	s.RetMsg = &v
	return s
}

func (s *CallBackThirdRightSendPlanResponseBody) SetRetValue(v bool) *CallBackThirdRightSendPlanResponseBody {
	s.RetValue = &v
	return s
}

func (s *CallBackThirdRightSendPlanResponseBody) SetRequestId(v string) *CallBackThirdRightSendPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CallBackThirdRightSendPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
