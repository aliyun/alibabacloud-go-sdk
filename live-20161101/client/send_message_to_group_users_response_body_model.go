// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageToGroupUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendMessageToGroupUsersResponseBody
	GetRequestId() *string
	SetResult(v *SendMessageToGroupUsersResponseBodyResult) *SendMessageToGroupUsersResponseBody
	GetResult() *SendMessageToGroupUsersResponseBodyResult
}

type SendMessageToGroupUsersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *SendMessageToGroupUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s SendMessageToGroupUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendMessageToGroupUsersResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageToGroupUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendMessageToGroupUsersResponseBody) GetResult() *SendMessageToGroupUsersResponseBodyResult {
	return s.Result
}

func (s *SendMessageToGroupUsersResponseBody) SetRequestId(v string) *SendMessageToGroupUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageToGroupUsersResponseBody) SetResult(v *SendMessageToGroupUsersResponseBodyResult) *SendMessageToGroupUsersResponseBody {
	s.Result = v
	return s
}

func (s *SendMessageToGroupUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type SendMessageToGroupUsersResponseBodyResult struct {
	// The ID of the message.
	//
	// example:
	//
	// hp***
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SendMessageToGroupUsersResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s SendMessageToGroupUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendMessageToGroupUsersResponseBodyResult) GetMessageId() *string {
	return s.MessageId
}

func (s *SendMessageToGroupUsersResponseBodyResult) SetMessageId(v string) *SendMessageToGroupUsersResponseBodyResult {
	s.MessageId = &v
	return s
}

func (s *SendMessageToGroupUsersResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
