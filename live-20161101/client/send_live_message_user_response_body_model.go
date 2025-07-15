// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLiveMessageUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMsgTid(v string) *SendLiveMessageUserResponseBody
	GetMsgTid() *string
	SetRequestId(v string) *SendLiveMessageUserResponseBody
	GetRequestId() *string
}

type SendLiveMessageUserResponseBody struct {
	// The ID of the message, which is a unique identifier that can be used to delete the message. The ID can be up to 64 bytes in length and can contain letters and digits.
	//
	// example:
	//
	// 169830****
	MsgTid *string `json:"MsgTid,omitempty" xml:"MsgTid,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6CFDE7AB-571A-14EA-B072-989FF753****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendLiveMessageUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendLiveMessageUserResponseBody) GoString() string {
	return s.String()
}

func (s *SendLiveMessageUserResponseBody) GetMsgTid() *string {
	return s.MsgTid
}

func (s *SendLiveMessageUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendLiveMessageUserResponseBody) SetMsgTid(v string) *SendLiveMessageUserResponseBody {
	s.MsgTid = &v
	return s
}

func (s *SendLiveMessageUserResponseBody) SetRequestId(v string) *SendLiveMessageUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendLiveMessageUserResponseBody) Validate() error {
	return dara.Validate(s)
}
