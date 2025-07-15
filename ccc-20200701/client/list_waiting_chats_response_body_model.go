// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaitingChatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWaitingChatsResponseBody
	GetCode() *string
	SetData(v []*ListWaitingChatsResponseBodyData) *ListWaitingChatsResponseBody
	GetData() []*ListWaitingChatsResponseBodyData
	SetHttpStatusCode(v int32) *ListWaitingChatsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListWaitingChatsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListWaitingChatsResponseBody
	GetRequestId() *string
}

type ListWaitingChatsResponseBody struct {
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListWaitingChatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 03C67DAD-EB26-41D8-949D-9B0C470FB716
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListWaitingChatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingChatsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWaitingChatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWaitingChatsResponseBody) GetData() []*ListWaitingChatsResponseBodyData {
	return s.Data
}

func (s *ListWaitingChatsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListWaitingChatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWaitingChatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWaitingChatsResponseBody) SetCode(v string) *ListWaitingChatsResponseBody {
	s.Code = &v
	return s
}

func (s *ListWaitingChatsResponseBody) SetData(v []*ListWaitingChatsResponseBodyData) *ListWaitingChatsResponseBody {
	s.Data = v
	return s
}

func (s *ListWaitingChatsResponseBody) SetHttpStatusCode(v int32) *ListWaitingChatsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListWaitingChatsResponseBody) SetMessage(v string) *ListWaitingChatsResponseBody {
	s.Message = &v
	return s
}

func (s *ListWaitingChatsResponseBody) SetRequestId(v string) *ListWaitingChatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWaitingChatsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListWaitingChatsResponseBodyData struct {
	// example:
	//
	// 843073c2-*****-49fb-a616-738ddddfebdc
	AccessChannelId *string `json:"AccessChannelId,omitempty" xml:"AccessChannelId,omitempty"`
	// example:
	//
	// Web
	AccessChannelType *string `json:"AccessChannelType,omitempty" xml:"AccessChannelType,omitempty"`
	// example:
	//
	// false
	BeingAssigned *bool `json:"BeingAssigned,omitempty" xml:"BeingAssigned,omitempty"`
	// example:
	//
	// $23086709$EAUNIT
	ChatConversationId *string `json:"ChatConversationId,omitempty" xml:"ChatConversationId,omitempty"`
	// example:
	//
	// 1718868572094
	EnqueueTime *int64 `json:"EnqueueTime,omitempty" xml:"EnqueueTime,omitempty"`
	// example:
	//
	// chat-434537064047960064
	JobId        *string                                     `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Messages     []*ListWaitingChatsResponseBodyDataMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	SkillGroupId *string                                     `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	UserList     []*ListWaitingChatsResponseBodyDataUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s ListWaitingChatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingChatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWaitingChatsResponseBodyData) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *ListWaitingChatsResponseBodyData) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *ListWaitingChatsResponseBodyData) GetBeingAssigned() *bool {
	return s.BeingAssigned
}

func (s *ListWaitingChatsResponseBodyData) GetChatConversationId() *string {
	return s.ChatConversationId
}

func (s *ListWaitingChatsResponseBodyData) GetEnqueueTime() *int64 {
	return s.EnqueueTime
}

func (s *ListWaitingChatsResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *ListWaitingChatsResponseBodyData) GetMessages() []*ListWaitingChatsResponseBodyDataMessages {
	return s.Messages
}

func (s *ListWaitingChatsResponseBodyData) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListWaitingChatsResponseBodyData) GetUserList() []*ListWaitingChatsResponseBodyDataUserList {
	return s.UserList
}

func (s *ListWaitingChatsResponseBodyData) SetAccessChannelId(v string) *ListWaitingChatsResponseBodyData {
	s.AccessChannelId = &v
	return s
}

func (s *ListWaitingChatsResponseBodyData) SetAccessChannelType(v string) *ListWaitingChatsResponseBodyData {
	s.AccessChannelType = &v
	return s
}

func (s *ListWaitingChatsResponseBodyData) SetBeingAssigned(v bool) *ListWaitingChatsResponseBodyData {
	s.BeingAssigned = &v
	return s
}

func (s *ListWaitingChatsResponseBodyData) SetChatConversationId(v string) *ListWaitingChatsResponseBodyData {
	s.ChatConversationId = &v
	return s
}

func (s *ListWaitingChatsResponseBodyData) SetEnqueueTime(v int64) *ListWaitingChatsResponseBodyData {
	s.EnqueueTime = &v
	return s
}

func (s *ListWaitingChatsResponseBodyData) SetJobId(v string) *ListWaitingChatsResponseBodyData {
	s.JobId = &v
	return s
}

func (s *ListWaitingChatsResponseBodyData) SetMessages(v []*ListWaitingChatsResponseBodyDataMessages) *ListWaitingChatsResponseBodyData {
	s.Messages = v
	return s
}

func (s *ListWaitingChatsResponseBodyData) SetSkillGroupId(v string) *ListWaitingChatsResponseBodyData {
	s.SkillGroupId = &v
	return s
}

func (s *ListWaitingChatsResponseBodyData) SetUserList(v []*ListWaitingChatsResponseBodyDataUserList) *ListWaitingChatsResponseBodyData {
	s.UserList = v
	return s
}

func (s *ListWaitingChatsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListWaitingChatsResponseBodyDataMessages struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// c361765f-******-4e07-b81c-4b5d9183fac6
	SenderId   *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	SenderType *string `json:"SenderType,omitempty" xml:"SenderType,omitempty"`
}

func (s ListWaitingChatsResponseBodyDataMessages) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingChatsResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *ListWaitingChatsResponseBodyDataMessages) GetContent() *string {
	return s.Content
}

func (s *ListWaitingChatsResponseBodyDataMessages) GetSenderId() *string {
	return s.SenderId
}

func (s *ListWaitingChatsResponseBodyDataMessages) GetSenderType() *string {
	return s.SenderType
}

func (s *ListWaitingChatsResponseBodyDataMessages) SetContent(v string) *ListWaitingChatsResponseBodyDataMessages {
	s.Content = &v
	return s
}

func (s *ListWaitingChatsResponseBodyDataMessages) SetSenderId(v string) *ListWaitingChatsResponseBodyDataMessages {
	s.SenderId = &v
	return s
}

func (s *ListWaitingChatsResponseBodyDataMessages) SetSenderType(v string) *ListWaitingChatsResponseBodyDataMessages {
	s.SenderType = &v
	return s
}

func (s *ListWaitingChatsResponseBodyDataMessages) Validate() error {
	return dara.Validate(s)
}

type ListWaitingChatsResponseBodyDataUserList struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// example:
	//
	// c361765f-******-4e07-b81c-4b5d9183fac6
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// example:
	//
	// CUSTOMER
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s ListWaitingChatsResponseBodyDataUserList) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingChatsResponseBodyDataUserList) GoString() string {
	return s.String()
}

func (s *ListWaitingChatsResponseBodyDataUserList) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListWaitingChatsResponseBodyDataUserList) GetUserId() *string {
	return s.UserId
}

func (s *ListWaitingChatsResponseBodyDataUserList) GetUserName() *string {
	return s.UserName
}

func (s *ListWaitingChatsResponseBodyDataUserList) GetUserType() *string {
	return s.UserType
}

func (s *ListWaitingChatsResponseBodyDataUserList) SetAvatarUrl(v string) *ListWaitingChatsResponseBodyDataUserList {
	s.AvatarUrl = &v
	return s
}

func (s *ListWaitingChatsResponseBodyDataUserList) SetUserId(v string) *ListWaitingChatsResponseBodyDataUserList {
	s.UserId = &v
	return s
}

func (s *ListWaitingChatsResponseBodyDataUserList) SetUserName(v string) *ListWaitingChatsResponseBodyDataUserList {
	s.UserName = &v
	return s
}

func (s *ListWaitingChatsResponseBodyDataUserList) SetUserType(v string) *ListWaitingChatsResponseBodyDataUserList {
	s.UserType = &v
	return s
}

func (s *ListWaitingChatsResponseBodyDataUserList) Validate() error {
	return dara.Validate(s)
}
