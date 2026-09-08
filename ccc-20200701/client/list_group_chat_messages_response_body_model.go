// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupChatMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGroupChatMessagesResponseBody
	GetCode() *string
	SetData(v *ListGroupChatMessagesResponseBodyData) *ListGroupChatMessagesResponseBody
	GetData() *ListGroupChatMessagesResponseBodyData
	SetHttpStatusCode(v int32) *ListGroupChatMessagesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGroupChatMessagesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGroupChatMessagesResponseBody
	GetRequestId() *string
}

type ListGroupChatMessagesResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *ListGroupChatMessagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code returned in the response. A value of 200 indicates a successful request.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 2263B273-AC1B-44EB-BA98-87F2322C6780
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGroupChatMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupChatMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupChatMessagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGroupChatMessagesResponseBody) GetData() *ListGroupChatMessagesResponseBodyData {
	return s.Data
}

func (s *ListGroupChatMessagesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGroupChatMessagesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGroupChatMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupChatMessagesResponseBody) SetCode(v string) *ListGroupChatMessagesResponseBody {
	s.Code = &v
	return s
}

func (s *ListGroupChatMessagesResponseBody) SetData(v *ListGroupChatMessagesResponseBodyData) *ListGroupChatMessagesResponseBody {
	s.Data = v
	return s
}

func (s *ListGroupChatMessagesResponseBody) SetHttpStatusCode(v int32) *ListGroupChatMessagesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGroupChatMessagesResponseBody) SetMessage(v string) *ListGroupChatMessagesResponseBody {
	s.Message = &v
	return s
}

func (s *ListGroupChatMessagesResponseBody) SetRequestId(v string) *ListGroupChatMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupChatMessagesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGroupChatMessagesResponseBodyData struct {
	// Message list.
	Messages []*ListGroupChatMessagesResponseBodyDataMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// Token for the next page.
	//
	// example:
	//
	// 54d1a616d95a4a01ba58967a9115b649
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
}

func (s ListGroupChatMessagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGroupChatMessagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGroupChatMessagesResponseBodyData) GetMessages() []*ListGroupChatMessagesResponseBodyDataMessages {
	return s.Messages
}

func (s *ListGroupChatMessagesResponseBodyData) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListGroupChatMessagesResponseBodyData) SetMessages(v []*ListGroupChatMessagesResponseBodyDataMessages) *ListGroupChatMessagesResponseBodyData {
	s.Messages = v
	return s
}

func (s *ListGroupChatMessagesResponseBodyData) SetNextPageToken(v string) *ListGroupChatMessagesResponseBodyData {
	s.NextPageToken = &v
	return s
}

func (s *ListGroupChatMessagesResponseBodyData) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupChatMessagesResponseBodyDataMessages struct {
	// Message content.
	//
	// example:
	//
	// {"variables":{},"text":"<p>好的，不客气</p>","contentType":"Text","subContentType":"richtext"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Call ID.
	//
	// example:
	//
	// chat-65382141036853491
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Indicates whether the message was revoked.
	//
	// example:
	//
	// false
	Recalled *bool `json:"Recalled,omitempty" xml:"Recalled,omitempty"`
	// Sender profile picture URL.
	//
	// example:
	//
	// http://xxxxx.com
	SenderAvatarUrl *string `json:"SenderAvatarUrl,omitempty" xml:"SenderAvatarUrl,omitempty"`
	// User ID of the message sender.
	//
	// example:
	//
	// 64bb4ececc34fc5ec1ca1153
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// Sender name.
	//
	// example:
	//
	// test-agent@test-instanceId
	SenderName *string `json:"SenderName,omitempty" xml:"SenderName,omitempty"`
	// Sender type.
	//
	// Valid values:
	//
	// - **ADMIN**: system
	//
	// - **CUSTOMER**: visitor
	//
	// - **AGENT**: agent
	//
	// example:
	//
	// CUSTOMER
	SenderType *string `json:"SenderType,omitempty" xml:"SenderType,omitempty"`
	// Message timestamp, in Unix timestamp format, measured in milliseconds.
	//
	// example:
	//
	// 1696126980371
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListGroupChatMessagesResponseBodyDataMessages) String() string {
	return dara.Prettify(s)
}

func (s ListGroupChatMessagesResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) GetContent() *string {
	return s.Content
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) GetJobId() *string {
	return s.JobId
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) GetRecalled() *bool {
	return s.Recalled
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) GetSenderAvatarUrl() *string {
	return s.SenderAvatarUrl
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) GetSenderId() *string {
	return s.SenderId
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) GetSenderName() *string {
	return s.SenderName
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) GetSenderType() *string {
	return s.SenderType
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) SetContent(v string) *ListGroupChatMessagesResponseBodyDataMessages {
	s.Content = &v
	return s
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) SetJobId(v string) *ListGroupChatMessagesResponseBodyDataMessages {
	s.JobId = &v
	return s
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) SetRecalled(v bool) *ListGroupChatMessagesResponseBodyDataMessages {
	s.Recalled = &v
	return s
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) SetSenderAvatarUrl(v string) *ListGroupChatMessagesResponseBodyDataMessages {
	s.SenderAvatarUrl = &v
	return s
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) SetSenderId(v string) *ListGroupChatMessagesResponseBodyDataMessages {
	s.SenderId = &v
	return s
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) SetSenderName(v string) *ListGroupChatMessagesResponseBodyDataMessages {
	s.SenderName = &v
	return s
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) SetSenderType(v string) *ListGroupChatMessagesResponseBodyDataMessages {
	s.SenderType = &v
	return s
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) SetTimestamp(v int64) *ListGroupChatMessagesResponseBodyDataMessages {
	s.Timestamp = &v
	return s
}

func (s *ListGroupChatMessagesResponseBodyDataMessages) Validate() error {
	return dara.Validate(s)
}
