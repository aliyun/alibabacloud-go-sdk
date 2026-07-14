// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatappMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListChatappMessageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListChatappMessageResponseBody
	GetCode() *string
	SetData(v []*ListChatappMessageResponseBodyData) *ListChatappMessageResponseBody
	GetData() []*ListChatappMessageResponseBodyData
	SetMessage(v string) *ListChatappMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListChatappMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListChatappMessageResponseBody
	GetSuccess() *bool
}

type ListChatappMessageResponseBody struct {
	// The access denied details.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// - OK: The request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of returned data objects.
	Data []*ListChatappMessageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 608F9CCA-B5EB-3D72-8047-B25D6D75BDEC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListChatappMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChatappMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatappMessageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListChatappMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChatappMessageResponseBody) GetData() []*ListChatappMessageResponseBodyData {
	return s.Data
}

func (s *ListChatappMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChatappMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatappMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListChatappMessageResponseBody) SetAccessDeniedDetail(v string) *ListChatappMessageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListChatappMessageResponseBody) SetCode(v string) *ListChatappMessageResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatappMessageResponseBody) SetData(v []*ListChatappMessageResponseBodyData) *ListChatappMessageResponseBody {
	s.Data = v
	return s
}

func (s *ListChatappMessageResponseBody) SetMessage(v string) *ListChatappMessageResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatappMessageResponseBody) SetRequestId(v string) *ListChatappMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatappMessageResponseBody) SetSuccess(v bool) *ListChatappMessageResponseBody {
	s.Success = &v
	return s
}

func (s *ListChatappMessageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListChatappMessageResponseBodyData struct {
	// The business phone number.
	//
	// example:
	//
	// 86183********
	BusinessNumber *string `json:"BusinessNumber,omitempty" xml:"BusinessNumber,omitempty"`
	// The channel type.
	//
	// example:
	//
	// WHATSAPP
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// The name of the message receiving status.
	//
	// example:
	//
	// Success
	ClientAcceptStatusName *string `json:"ClientAcceptStatusName,omitempty" xml:"ClientAcceptStatusName,omitempty"`
	// The message read status.
	//
	// example:
	//
	// success
	ClientReadStatus *string `json:"ClientReadStatus,omitempty" xml:"ClientReadStatus,omitempty"`
	// The message read status name.
	//
	// example:
	//
	// Success
	ClientReadStatusName *string `json:"ClientReadStatusName,omitempty" xml:"ClientReadStatusName,omitempty"`
	// The conversation ID.
	//
	// example:
	//
	// 805a66**************************
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The inbound or outbound message type.
	//
	// example:
	//
	// DOWN
	EventAction *string `json:"EventAction,omitempty" xml:"EventAction,omitempty"`
	// The name of the inbound or outbound message type. Valid values:
	//
	// - DOWN: outbound message.
	//
	// - UP: inbound message.
	//
	// example:
	//
	// DOWN
	EventActionName *string `json:"EventActionName,omitempty" xml:"EventActionName,omitempty"`
	// The fallback content.
	//
	// example:
	//
	// None
	FailBackContent *string `json:"FailBackContent,omitempty" xml:"FailBackContent,omitempty"`
	// Indicates whether the message falls back to SMS. Valid values:
	//
	// - Y: Yes.
	//
	// - N: No.
	//
	// example:
	//
	// Y
	FailBackFlag *string `json:"FailBackFlag,omitempty" xml:"FailBackFlag,omitempty"`
	// The reason for the sending failure.
	//
	// example:
	//
	// timeout
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// The template language. For more languages, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// example:
	//
	// en
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	// The message content.
	//
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The message ID.
	//
	// example:
	//
	// 202509*******************
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The message source.
	//
	// example:
	//
	// api
	MessageSource *string `json:"MessageSource,omitempty" xml:"MessageSource,omitempty"`
	// The message status.
	//
	// example:
	//
	// success
	MessageStatus *string `json:"MessageStatus,omitempty" xml:"MessageStatus,omitempty"`
	// The message status name.
	//
	// example:
	//
	// Success
	MessageStatusName *string `json:"MessageStatusName,omitempty" xml:"MessageStatusName,omitempty"`
	// The message type.
	//
	// example:
	//
	// INTERACTIVE
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The message type name.
	//
	// example:
	//
	// interactive
	MessageTypeName *string `json:"MessageTypeName,omitempty" xml:"MessageTypeName,omitempty"`
	// The month of the message.
	//
	// example:
	//
	// 202507
	Month *string `json:"Month,omitempty" xml:"Month,omitempty"`
	// The sending time.
	//
	// example:
	//
	// 2025-07-11T01:16:49.761+00:00
	SendTime *string `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	// The template code.
	//
	// example:
	//
	// 1103***************
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template name.
	//
	// example:
	//
	// picture_template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type.
	//
	// example:
	//
	// message
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The unique message ID.
	//
	// example:
	//
	// 20250911******************************
	UniqueMessageId *string `json:"UniqueMessageId,omitempty" xml:"UniqueMessageId,omitempty"`
	// The user phone number.
	//
	// example:
	//
	// 86177********
	UserNumber *string `json:"UserNumber,omitempty" xml:"UserNumber,omitempty"`
}

func (s ListChatappMessageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListChatappMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChatappMessageResponseBodyData) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *ListChatappMessageResponseBodyData) GetChannelType() *string {
	return s.ChannelType
}

func (s *ListChatappMessageResponseBodyData) GetClientAcceptStatusName() *string {
	return s.ClientAcceptStatusName
}

func (s *ListChatappMessageResponseBodyData) GetClientReadStatus() *string {
	return s.ClientReadStatus
}

func (s *ListChatappMessageResponseBodyData) GetClientReadStatusName() *string {
	return s.ClientReadStatusName
}

func (s *ListChatappMessageResponseBodyData) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListChatappMessageResponseBodyData) GetEventAction() *string {
	return s.EventAction
}

func (s *ListChatappMessageResponseBodyData) GetEventActionName() *string {
	return s.EventActionName
}

func (s *ListChatappMessageResponseBodyData) GetFailBackContent() *string {
	return s.FailBackContent
}

func (s *ListChatappMessageResponseBodyData) GetFailBackFlag() *string {
	return s.FailBackFlag
}

func (s *ListChatappMessageResponseBodyData) GetFailReason() *string {
	return s.FailReason
}

func (s *ListChatappMessageResponseBodyData) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *ListChatappMessageResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ListChatappMessageResponseBodyData) GetMessageId() *string {
	return s.MessageId
}

func (s *ListChatappMessageResponseBodyData) GetMessageSource() *string {
	return s.MessageSource
}

func (s *ListChatappMessageResponseBodyData) GetMessageStatus() *string {
	return s.MessageStatus
}

func (s *ListChatappMessageResponseBodyData) GetMessageStatusName() *string {
	return s.MessageStatusName
}

func (s *ListChatappMessageResponseBodyData) GetMessageType() *string {
	return s.MessageType
}

func (s *ListChatappMessageResponseBodyData) GetMessageTypeName() *string {
	return s.MessageTypeName
}

func (s *ListChatappMessageResponseBodyData) GetMonth() *string {
	return s.Month
}

func (s *ListChatappMessageResponseBodyData) GetSendTime() *string {
	return s.SendTime
}

func (s *ListChatappMessageResponseBodyData) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ListChatappMessageResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListChatappMessageResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListChatappMessageResponseBodyData) GetUniqueMessageId() *string {
	return s.UniqueMessageId
}

func (s *ListChatappMessageResponseBodyData) GetUserNumber() *string {
	return s.UserNumber
}

func (s *ListChatappMessageResponseBodyData) SetBusinessNumber(v string) *ListChatappMessageResponseBodyData {
	s.BusinessNumber = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetChannelType(v string) *ListChatappMessageResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetClientAcceptStatusName(v string) *ListChatappMessageResponseBodyData {
	s.ClientAcceptStatusName = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetClientReadStatus(v string) *ListChatappMessageResponseBodyData {
	s.ClientReadStatus = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetClientReadStatusName(v string) *ListChatappMessageResponseBodyData {
	s.ClientReadStatusName = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetConversationId(v string) *ListChatappMessageResponseBodyData {
	s.ConversationId = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetEventAction(v string) *ListChatappMessageResponseBodyData {
	s.EventAction = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetEventActionName(v string) *ListChatappMessageResponseBodyData {
	s.EventActionName = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetFailBackContent(v string) *ListChatappMessageResponseBodyData {
	s.FailBackContent = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetFailBackFlag(v string) *ListChatappMessageResponseBodyData {
	s.FailBackFlag = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetFailReason(v string) *ListChatappMessageResponseBodyData {
	s.FailReason = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetLanguageCode(v string) *ListChatappMessageResponseBodyData {
	s.LanguageCode = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetMessage(v string) *ListChatappMessageResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetMessageId(v string) *ListChatappMessageResponseBodyData {
	s.MessageId = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetMessageSource(v string) *ListChatappMessageResponseBodyData {
	s.MessageSource = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetMessageStatus(v string) *ListChatappMessageResponseBodyData {
	s.MessageStatus = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetMessageStatusName(v string) *ListChatappMessageResponseBodyData {
	s.MessageStatusName = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetMessageType(v string) *ListChatappMessageResponseBodyData {
	s.MessageType = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetMessageTypeName(v string) *ListChatappMessageResponseBodyData {
	s.MessageTypeName = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetMonth(v string) *ListChatappMessageResponseBodyData {
	s.Month = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetSendTime(v string) *ListChatappMessageResponseBodyData {
	s.SendTime = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetTemplateCode(v string) *ListChatappMessageResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetTemplateName(v string) *ListChatappMessageResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetType(v string) *ListChatappMessageResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetUniqueMessageId(v string) *ListChatappMessageResponseBodyData {
	s.UniqueMessageId = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) SetUserNumber(v string) *ListChatappMessageResponseBodyData {
	s.UserNumber = &v
	return s
}

func (s *ListChatappMessageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
