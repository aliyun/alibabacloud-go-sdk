// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLiveMessageGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMsgTid(v string) *SendLiveMessageGroupResponseBody
	GetMsgTid() *string
	SetRequestId(v string) *SendLiveMessageGroupResponseBody
	GetRequestId() *string
}

type SendLiveMessageGroupResponseBody struct {
	// The unique identifier of the message, which can be used to delete the message. The ID can contain only letters and digits and can be up to 64 bytes in length.
	//
	// example:
	//
	// 169830****
	MsgTid *string `json:"MsgTid,omitempty" xml:"MsgTid,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E4C1245F-597B-1BD1-B9BB-9D220E99****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendLiveMessageGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendLiveMessageGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SendLiveMessageGroupResponseBody) GetMsgTid() *string {
	return s.MsgTid
}

func (s *SendLiveMessageGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendLiveMessageGroupResponseBody) SetMsgTid(v string) *SendLiveMessageGroupResponseBody {
	s.MsgTid = &v
	return s
}

func (s *SendLiveMessageGroupResponseBody) SetRequestId(v string) *SendLiveMessageGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendLiveMessageGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
