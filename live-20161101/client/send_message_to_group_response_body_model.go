// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageToGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendMessageToGroupResponseBody
	GetRequestId() *string
	SetResult(v *SendMessageToGroupResponseBodyResult) *SendMessageToGroupResponseBody
	GetResult() *SendMessageToGroupResponseBodyResult
}

type SendMessageToGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data returned.
	Result *SendMessageToGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s SendMessageToGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendMessageToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageToGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendMessageToGroupResponseBody) GetResult() *SendMessageToGroupResponseBodyResult {
	return s.Result
}

func (s *SendMessageToGroupResponseBody) SetRequestId(v string) *SendMessageToGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageToGroupResponseBody) SetResult(v *SendMessageToGroupResponseBodyResult) *SendMessageToGroupResponseBody {
	s.Result = v
	return s
}

func (s *SendMessageToGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type SendMessageToGroupResponseBodyResult struct {
	// The ID of the message.
	//
	// example:
	//
	// qt***
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SendMessageToGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s SendMessageToGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendMessageToGroupResponseBodyResult) GetMessageId() *string {
	return s.MessageId
}

func (s *SendMessageToGroupResponseBodyResult) SetMessageId(v string) *SendMessageToGroupResponseBodyResult {
	s.MessageId = &v
	return s
}

func (s *SendMessageToGroupResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
