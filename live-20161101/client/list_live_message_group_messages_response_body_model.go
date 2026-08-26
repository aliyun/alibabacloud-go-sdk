// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageGroupMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListLiveMessageGroupMessagesResponseBody
	GetGroupId() *string
	SetHasmore(v bool) *ListLiveMessageGroupMessagesResponseBody
	GetHasmore() *bool
	SetMessageList(v []*ListLiveMessageGroupMessagesResponseBodyMessageList) *ListLiveMessageGroupMessagesResponseBody
	GetMessageList() []*ListLiveMessageGroupMessagesResponseBodyMessageList
	SetNextPageToken(v int64) *ListLiveMessageGroupMessagesResponseBody
	GetNextPageToken() *int64
	SetRequestId(v string) *ListLiveMessageGroupMessagesResponseBody
	GetRequestId() *string
}

type ListLiveMessageGroupMessagesResponseBody struct {
	// The group ID of the group to query.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Indicates whether there is a next page.
	//
	// example:
	//
	// false
	Hasmore *bool `json:"Hasmore,omitempty" xml:"Hasmore,omitempty"`
	// The message list.
	MessageList []*ListLiveMessageGroupMessagesResponseBodyMessageList `json:"MessageList,omitempty" xml:"MessageList,omitempty" type:"Repeated"`
	// The start position of the next page. The value is 0 if there is no next page.
	//
	// example:
	//
	// 0
	NextPageToken *int64 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1668FDC3-63D7-102F-B5D4-3D2F91D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLiveMessageGroupMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupMessagesResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *ListLiveMessageGroupMessagesResponseBody) GetHasmore() *bool {
	return s.Hasmore
}

func (s *ListLiveMessageGroupMessagesResponseBody) GetMessageList() []*ListLiveMessageGroupMessagesResponseBodyMessageList {
	return s.MessageList
}

func (s *ListLiveMessageGroupMessagesResponseBody) GetNextPageToken() *int64 {
	return s.NextPageToken
}

func (s *ListLiveMessageGroupMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveMessageGroupMessagesResponseBody) SetGroupId(v string) *ListLiveMessageGroupMessagesResponseBody {
	s.GroupId = &v
	return s
}

func (s *ListLiveMessageGroupMessagesResponseBody) SetHasmore(v bool) *ListLiveMessageGroupMessagesResponseBody {
	s.Hasmore = &v
	return s
}

func (s *ListLiveMessageGroupMessagesResponseBody) SetMessageList(v []*ListLiveMessageGroupMessagesResponseBodyMessageList) *ListLiveMessageGroupMessagesResponseBody {
	s.MessageList = v
	return s
}

func (s *ListLiveMessageGroupMessagesResponseBody) SetNextPageToken(v int64) *ListLiveMessageGroupMessagesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListLiveMessageGroupMessagesResponseBody) SetRequestId(v string) *ListLiveMessageGroupMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveMessageGroupMessagesResponseBody) Validate() error {
	if s.MessageList != nil {
		for _, item := range s.MessageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLiveMessageGroupMessagesResponseBodyMessageList struct {
	// The message body.
	//
	// example:
	//
	// step2 helo, cc group
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The message identifier.
	//
	// example:
	//
	// c-1-1-0
	MsgTid *string `json:"MsgTid,omitempty" xml:"MsgTid,omitempty"`
	// The message type.
	//
	// example:
	//
	// 2
	MsgType *int64 `json:"MsgType,omitempty" xml:"MsgType,omitempty"`
	// The message sender details.
	Sender *ListLiveMessageGroupMessagesResponseBodyMessageListSender `json:"Sender,omitempty" xml:"Sender,omitempty" type:"Struct"`
	// The message sequence number.
	//
	// example:
	//
	// 1
	SeqNumber *int64 `json:"SeqNumber,omitempty" xml:"SeqNumber,omitempty"`
	// The time when the message was created, expressed as a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1697081134
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The total number of messages.
	//
	// example:
	//
	// 1
	TotalMessages *int64 `json:"TotalMessages,omitempty" xml:"TotalMessages,omitempty"`
}

func (s ListLiveMessageGroupMessagesResponseBodyMessageList) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupMessagesResponseBodyMessageList) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) GetBody() *string {
	return s.Body
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) GetMsgTid() *string {
	return s.MsgTid
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) GetMsgType() *int64 {
	return s.MsgType
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) GetSender() *ListLiveMessageGroupMessagesResponseBodyMessageListSender {
	return s.Sender
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) GetSeqNumber() *int64 {
	return s.SeqNumber
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) GetTotalMessages() *int64 {
	return s.TotalMessages
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) SetBody(v string) *ListLiveMessageGroupMessagesResponseBodyMessageList {
	s.Body = &v
	return s
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) SetMsgTid(v string) *ListLiveMessageGroupMessagesResponseBodyMessageList {
	s.MsgTid = &v
	return s
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) SetMsgType(v int64) *ListLiveMessageGroupMessagesResponseBodyMessageList {
	s.MsgType = &v
	return s
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) SetSender(v *ListLiveMessageGroupMessagesResponseBodyMessageListSender) *ListLiveMessageGroupMessagesResponseBodyMessageList {
	s.Sender = v
	return s
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) SetSeqNumber(v int64) *ListLiveMessageGroupMessagesResponseBodyMessageList {
	s.SeqNumber = &v
	return s
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) SetTimestamp(v int64) *ListLiveMessageGroupMessagesResponseBodyMessageList {
	s.Timestamp = &v
	return s
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) SetTotalMessages(v int64) *ListLiveMessageGroupMessagesResponseBodyMessageList {
	s.TotalMessages = &v
	return s
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageList) Validate() error {
	if s.Sender != nil {
		if err := s.Sender.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLiveMessageGroupMessagesResponseBodyMessageListSender struct {
	// The ID of the user who sent the message.
	//
	// example:
	//
	// uid2
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The extended information of the user who sent the message.
	//
	// example:
	//
	// testusermeta2
	UserInfo *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListLiveMessageGroupMessagesResponseBodyMessageListSender) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupMessagesResponseBodyMessageListSender) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageListSender) GetUserId() *string {
	return s.UserId
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageListSender) GetUserInfo() *string {
	return s.UserInfo
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageListSender) SetUserId(v string) *ListLiveMessageGroupMessagesResponseBodyMessageListSender {
	s.UserId = &v
	return s
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageListSender) SetUserInfo(v string) *ListLiveMessageGroupMessagesResponseBodyMessageListSender {
	s.UserInfo = &v
	return s
}

func (s *ListLiveMessageGroupMessagesResponseBodyMessageListSender) Validate() error {
	return dara.Validate(s)
}
