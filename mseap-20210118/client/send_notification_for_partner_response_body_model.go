// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNotificationForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMsg(v string) *SendNotificationForPartnerResponseBody
	GetErrorMsg() *string
	SetMsgId(v string) *SendNotificationForPartnerResponseBody
	GetMsgId() *string
	SetRequestId(v string) *SendNotificationForPartnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendNotificationForPartnerResponseBody
	GetSuccess() *bool
}

type SendNotificationForPartnerResponseBody struct {
	// example:
	//
	// 11111111111111111111111
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 0A011920166449C2FAAE8D179E1704C5
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// example:
	//
	// 1940A84F-6D90-5764-9119-6279970C6FF5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendNotificationForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *SendNotificationForPartnerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SendNotificationForPartnerResponseBody) GetMsgId() *string {
	return s.MsgId
}

func (s *SendNotificationForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendNotificationForPartnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendNotificationForPartnerResponseBody) SetErrorMsg(v string) *SendNotificationForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SendNotificationForPartnerResponseBody) SetMsgId(v string) *SendNotificationForPartnerResponseBody {
	s.MsgId = &v
	return s
}

func (s *SendNotificationForPartnerResponseBody) SetRequestId(v string) *SendNotificationForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendNotificationForPartnerResponseBody) SetSuccess(v bool) *SendNotificationForPartnerResponseBody {
	s.Success = &v
	return s
}

func (s *SendNotificationForPartnerResponseBody) Validate() error {
	return dara.Validate(s)
}
