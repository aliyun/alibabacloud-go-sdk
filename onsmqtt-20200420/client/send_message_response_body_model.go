// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMsgId(v string) *SendMessageResponseBody
	GetMsgId() *string
	SetRequestId(v string) *SendMessageResponseBody
	GetRequestId() *string
}

type SendMessageResponseBody struct {
	// The unique message ID that is returned by the ApsaraMQ for MQTT broker after the message is sent.
	//
	// example:
	//
	// 0B736D997B7F45FF54E61C1C1B58****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The unique ID that the system generates for the request. This parameter is a common parameter.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) GetMsgId() *string {
	return s.MsgId
}

func (s *SendMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendMessageResponseBody) SetMsgId(v string) *SendMessageResponseBody {
	s.MsgId = &v
	return s
}

func (s *SendMessageResponseBody) SetRequestId(v string) *SendMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
