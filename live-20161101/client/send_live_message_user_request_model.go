// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLiveMessageUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SendLiveMessageUserRequest
	GetAppId() *string
	SetBody(v string) *SendLiveMessageUserRequest
	GetBody() *string
	SetDataCenter(v string) *SendLiveMessageUserRequest
	GetDataCenter() *string
	SetHighReliability(v bool) *SendLiveMessageUserRequest
	GetHighReliability() *bool
	SetMsgTid(v string) *SendLiveMessageUserRequest
	GetMsgTid() *string
	SetMsgType(v int64) *SendLiveMessageUserRequest
	GetMsgType() *int64
	SetReceiverId(v string) *SendLiveMessageUserRequest
	GetReceiverId() *string
	SetSenderId(v string) *SendLiveMessageUserRequest
	GetSenderId() *string
	SetSenderInfo(v string) *SendLiveMessageUserRequest
	GetSenderInfo() *string
	SetStorage(v bool) *SendLiveMessageUserRequest
	GetStorage() *bool
}

type SendLiveMessageUserRequest struct {
	// The ID of the interactive messaging application in which the message is sent.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The message body. It can be up to 15 KB in length.
	//
	// example:
	//
	// hello, user
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// Specifies whether to set the message as a highly reliable message. A highly reliable message ensures that success is returned only after the receiver has received the message and responded. If the receiver does not respond within 3 seconds, failure is returned.
	//
	// 	- true: sets the message as a highly reliable message.
	//
	// 	- false (default): does not set the message as a highly reliable message.
	//
	// >  This parameter is supported only by the client SDK V1.5.1 and later. When you send a message to a client with an earlier SDK version, the message can be successfully sent without waiting for an acknowledgement (ACK) response.
	//
	// example:
	//
	// false
	HighReliability *bool `json:"HighReliability,omitempty" xml:"HighReliability,omitempty"`
	// The ID of the message, which is a unique identifier that can be used to delete the message. The ID can be up to 64 bytes in length and can contain letters and digits.
	//
	// example:
	//
	// 169830****
	MsgTid *string `json:"MsgTid,omitempty" xml:"MsgTid,omitempty"`
	// The message type.
	//
	// example:
	//
	// 2
	MsgType *int64 `json:"MsgType,omitempty" xml:"MsgType,omitempty"`
	// The ID of the user who receives the message. The ID can be up to 64 bytes in length and can contain letters and digits.
	//
	// >  Make sure that the user who receives the message is online. You can call the CheckLiveMessageUsersOnline operation to query whether the user is online.
	//
	// This parameter is required.
	//
	// example:
	//
	// uid2
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	// The ID of the user who sends the message. The ID can be up to 64 bytes in length and can contain letters and digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// uid1
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// The additional information about the user who sends the message. It can be up to 512 bytes in length.
	//
	// example:
	//
	// uid1meta1
	SenderInfo *string `json:"SenderInfo,omitempty" xml:"SenderInfo,omitempty"`
	// Specifies whether to store the message.
	//
	// 	- true: stores the message.
	//
	// 	- false (default): does not store the message.
	//
	// example:
	//
	// false
	Storage *bool `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s SendLiveMessageUserRequest) String() string {
	return dara.Prettify(s)
}

func (s SendLiveMessageUserRequest) GoString() string {
	return s.String()
}

func (s *SendLiveMessageUserRequest) GetAppId() *string {
	return s.AppId
}

func (s *SendLiveMessageUserRequest) GetBody() *string {
	return s.Body
}

func (s *SendLiveMessageUserRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *SendLiveMessageUserRequest) GetHighReliability() *bool {
	return s.HighReliability
}

func (s *SendLiveMessageUserRequest) GetMsgTid() *string {
	return s.MsgTid
}

func (s *SendLiveMessageUserRequest) GetMsgType() *int64 {
	return s.MsgType
}

func (s *SendLiveMessageUserRequest) GetReceiverId() *string {
	return s.ReceiverId
}

func (s *SendLiveMessageUserRequest) GetSenderId() *string {
	return s.SenderId
}

func (s *SendLiveMessageUserRequest) GetSenderInfo() *string {
	return s.SenderInfo
}

func (s *SendLiveMessageUserRequest) GetStorage() *bool {
	return s.Storage
}

func (s *SendLiveMessageUserRequest) SetAppId(v string) *SendLiveMessageUserRequest {
	s.AppId = &v
	return s
}

func (s *SendLiveMessageUserRequest) SetBody(v string) *SendLiveMessageUserRequest {
	s.Body = &v
	return s
}

func (s *SendLiveMessageUserRequest) SetDataCenter(v string) *SendLiveMessageUserRequest {
	s.DataCenter = &v
	return s
}

func (s *SendLiveMessageUserRequest) SetHighReliability(v bool) *SendLiveMessageUserRequest {
	s.HighReliability = &v
	return s
}

func (s *SendLiveMessageUserRequest) SetMsgTid(v string) *SendLiveMessageUserRequest {
	s.MsgTid = &v
	return s
}

func (s *SendLiveMessageUserRequest) SetMsgType(v int64) *SendLiveMessageUserRequest {
	s.MsgType = &v
	return s
}

func (s *SendLiveMessageUserRequest) SetReceiverId(v string) *SendLiveMessageUserRequest {
	s.ReceiverId = &v
	return s
}

func (s *SendLiveMessageUserRequest) SetSenderId(v string) *SendLiveMessageUserRequest {
	s.SenderId = &v
	return s
}

func (s *SendLiveMessageUserRequest) SetSenderInfo(v string) *SendLiveMessageUserRequest {
	s.SenderInfo = &v
	return s
}

func (s *SendLiveMessageUserRequest) SetStorage(v bool) *SendLiveMessageUserRequest {
	s.Storage = &v
	return s
}

func (s *SendLiveMessageUserRequest) Validate() error {
	return dara.Validate(s)
}
