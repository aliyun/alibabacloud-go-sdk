// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisitorChatMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVisitorChatMessagesResponseBody
	GetCode() *string
	SetData(v *ListVisitorChatMessagesResponseBodyData) *ListVisitorChatMessagesResponseBody
	GetData() *ListVisitorChatMessagesResponseBodyData
	SetHttpStatusCode(v int32) *ListVisitorChatMessagesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListVisitorChatMessagesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVisitorChatMessagesResponseBody
	GetRequestId() *string
}

type ListVisitorChatMessagesResponseBody struct {
	// Response code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data
	Data *ListVisitorChatMessagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 8707EB29-BAED-4302-B999-40BA61877437
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVisitorChatMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVisitorChatMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVisitorChatMessagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVisitorChatMessagesResponseBody) GetData() *ListVisitorChatMessagesResponseBodyData {
	return s.Data
}

func (s *ListVisitorChatMessagesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListVisitorChatMessagesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVisitorChatMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVisitorChatMessagesResponseBody) SetCode(v string) *ListVisitorChatMessagesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVisitorChatMessagesResponseBody) SetData(v *ListVisitorChatMessagesResponseBodyData) *ListVisitorChatMessagesResponseBody {
	s.Data = v
	return s
}

func (s *ListVisitorChatMessagesResponseBody) SetHttpStatusCode(v int32) *ListVisitorChatMessagesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListVisitorChatMessagesResponseBody) SetMessage(v string) *ListVisitorChatMessagesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVisitorChatMessagesResponseBody) SetRequestId(v string) *ListVisitorChatMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVisitorChatMessagesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVisitorChatMessagesResponseBodyData struct {
	// Message list
	Messages []*ListVisitorChatMessagesResponseBodyDataMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// Token for the next page
	//
	// example:
	//
	// 1737193352340::7463707254.EAUNIT
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
}

func (s ListVisitorChatMessagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVisitorChatMessagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVisitorChatMessagesResponseBodyData) GetMessages() []*ListVisitorChatMessagesResponseBodyDataMessages {
	return s.Messages
}

func (s *ListVisitorChatMessagesResponseBodyData) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListVisitorChatMessagesResponseBodyData) SetMessages(v []*ListVisitorChatMessagesResponseBodyDataMessages) *ListVisitorChatMessagesResponseBodyData {
	s.Messages = v
	return s
}

func (s *ListVisitorChatMessagesResponseBodyData) SetNextPageToken(v string) *ListVisitorChatMessagesResponseBodyData {
	s.NextPageToken = &v
	return s
}

func (s *ListVisitorChatMessagesResponseBodyData) Validate() error {
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

type ListVisitorChatMessagesResponseBodyDataMessages struct {
	// Message content
	//
	// example:
	//
	// {"variables":{},"text":"Hi，我是小云~ 很高兴遇见你！","contentType":"Text"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Call ID
	//
	// example:
	//
	// chat-65382141036853491
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Sender profile picture URL
	//
	// example:
	//
	// http://xxxxx.com/avatar.png
	SenderAvatarUrl *string `json:"SenderAvatarUrl,omitempty" xml:"SenderAvatarUrl,omitempty"`
	// Sender user ID
	//
	// example:
	//
	// fcd020fe-****-1a272a174a7d
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// Sender name
	//
	// example:
	//
	// 访客1
	SenderName *string `json:"SenderName,omitempty" xml:"SenderName,omitempty"`
	// Sender type
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
	// UNIX timestamp
	//
	// example:
	//
	// 1696126980371
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListVisitorChatMessagesResponseBodyDataMessages) String() string {
	return dara.Prettify(s)
}

func (s ListVisitorChatMessagesResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) GetContent() *string {
	return s.Content
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) GetJobId() *string {
	return s.JobId
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) GetSenderAvatarUrl() *string {
	return s.SenderAvatarUrl
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) GetSenderId() *string {
	return s.SenderId
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) GetSenderName() *string {
	return s.SenderName
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) GetSenderType() *string {
	return s.SenderType
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) SetContent(v string) *ListVisitorChatMessagesResponseBodyDataMessages {
	s.Content = &v
	return s
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) SetJobId(v string) *ListVisitorChatMessagesResponseBodyDataMessages {
	s.JobId = &v
	return s
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) SetSenderAvatarUrl(v string) *ListVisitorChatMessagesResponseBodyDataMessages {
	s.SenderAvatarUrl = &v
	return s
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) SetSenderId(v string) *ListVisitorChatMessagesResponseBodyDataMessages {
	s.SenderId = &v
	return s
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) SetSenderName(v string) *ListVisitorChatMessagesResponseBodyDataMessages {
	s.SenderName = &v
	return s
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) SetSenderType(v string) *ListVisitorChatMessagesResponseBodyDataMessages {
	s.SenderType = &v
	return s
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) SetTimestamp(v int64) *ListVisitorChatMessagesResponseBodyDataMessages {
	s.Timestamp = &v
	return s
}

func (s *ListVisitorChatMessagesResponseBodyDataMessages) Validate() error {
	return dara.Validate(s)
}
