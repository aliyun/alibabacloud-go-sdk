// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatappMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *ListChatappMessageRequest
	GetBusinessNumber() *string
	SetChannelType(v string) *ListChatappMessageRequest
	GetChannelType() *string
	SetClientAcceptStatus(v string) *ListChatappMessageRequest
	GetClientAcceptStatus() *string
	SetCustSpaceId(v string) *ListChatappMessageRequest
	GetCustSpaceId() *string
	SetEndTime(v int64) *ListChatappMessageRequest
	GetEndTime() *int64
	SetEventAction(v string) *ListChatappMessageRequest
	GetEventAction() *string
	SetGroupMessageId(v string) *ListChatappMessageRequest
	GetGroupMessageId() *string
	SetMessageStatus(v string) *ListChatappMessageRequest
	GetMessageStatus() *string
	SetOwnerId(v int64) *ListChatappMessageRequest
	GetOwnerId() *int64
	SetPage(v *ListChatappMessageRequestPage) *ListChatappMessageRequest
	GetPage() *ListChatappMessageRequestPage
	SetResourceOwnerAccount(v string) *ListChatappMessageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListChatappMessageRequest
	GetResourceOwnerId() *int64
	SetStartTime(v int64) *ListChatappMessageRequest
	GetStartTime() *int64
	SetTemplateCode(v string) *ListChatappMessageRequest
	GetTemplateCode() *string
	SetUserNumber(v string) *ListChatappMessageRequest
	GetUserNumber() *string
}

type ListChatappMessageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8613800****
	BusinessNumber *string `json:"BusinessNumber,omitempty" xml:"BusinessNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WHATSAPP
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// success
	ClientAcceptStatus *string `json:"ClientAcceptStatus,omitempty" xml:"ClientAcceptStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// 1727057232686
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// UP
	EventAction *string `json:"EventAction,omitempty" xml:"EventAction,omitempty"`
	// example:
	//
	// 9292****
	GroupMessageId *string `json:"GroupMessageId,omitempty" xml:"GroupMessageId,omitempty"`
	// example:
	//
	// success
	MessageStatus *string `json:"MessageStatus,omitempty" xml:"MessageStatus,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	Page                 *ListChatappMessageRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	ResourceOwnerAccount *string                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                         `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 1727057232686
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 9938***
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// 86138***
	UserNumber *string `json:"UserNumber,omitempty" xml:"UserNumber,omitempty"`
}

func (s ListChatappMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatappMessageRequest) GoString() string {
	return s.String()
}

func (s *ListChatappMessageRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *ListChatappMessageRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *ListChatappMessageRequest) GetClientAcceptStatus() *string {
	return s.ClientAcceptStatus
}

func (s *ListChatappMessageRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListChatappMessageRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListChatappMessageRequest) GetEventAction() *string {
	return s.EventAction
}

func (s *ListChatappMessageRequest) GetGroupMessageId() *string {
	return s.GroupMessageId
}

func (s *ListChatappMessageRequest) GetMessageStatus() *string {
	return s.MessageStatus
}

func (s *ListChatappMessageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListChatappMessageRequest) GetPage() *ListChatappMessageRequestPage {
	return s.Page
}

func (s *ListChatappMessageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListChatappMessageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListChatappMessageRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListChatappMessageRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ListChatappMessageRequest) GetUserNumber() *string {
	return s.UserNumber
}

func (s *ListChatappMessageRequest) SetBusinessNumber(v string) *ListChatappMessageRequest {
	s.BusinessNumber = &v
	return s
}

func (s *ListChatappMessageRequest) SetChannelType(v string) *ListChatappMessageRequest {
	s.ChannelType = &v
	return s
}

func (s *ListChatappMessageRequest) SetClientAcceptStatus(v string) *ListChatappMessageRequest {
	s.ClientAcceptStatus = &v
	return s
}

func (s *ListChatappMessageRequest) SetCustSpaceId(v string) *ListChatappMessageRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListChatappMessageRequest) SetEndTime(v int64) *ListChatappMessageRequest {
	s.EndTime = &v
	return s
}

func (s *ListChatappMessageRequest) SetEventAction(v string) *ListChatappMessageRequest {
	s.EventAction = &v
	return s
}

func (s *ListChatappMessageRequest) SetGroupMessageId(v string) *ListChatappMessageRequest {
	s.GroupMessageId = &v
	return s
}

func (s *ListChatappMessageRequest) SetMessageStatus(v string) *ListChatappMessageRequest {
	s.MessageStatus = &v
	return s
}

func (s *ListChatappMessageRequest) SetOwnerId(v int64) *ListChatappMessageRequest {
	s.OwnerId = &v
	return s
}

func (s *ListChatappMessageRequest) SetPage(v *ListChatappMessageRequestPage) *ListChatappMessageRequest {
	s.Page = v
	return s
}

func (s *ListChatappMessageRequest) SetResourceOwnerAccount(v string) *ListChatappMessageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListChatappMessageRequest) SetResourceOwnerId(v int64) *ListChatappMessageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListChatappMessageRequest) SetStartTime(v int64) *ListChatappMessageRequest {
	s.StartTime = &v
	return s
}

func (s *ListChatappMessageRequest) SetTemplateCode(v string) *ListChatappMessageRequest {
	s.TemplateCode = &v
	return s
}

func (s *ListChatappMessageRequest) SetUserNumber(v string) *ListChatappMessageRequest {
	s.UserNumber = &v
	return s
}

func (s *ListChatappMessageRequest) Validate() error {
	return dara.Validate(s)
}

type ListChatappMessageRequestPage struct {
	// This parameter is required.
	//
	// example:
	//
	// 49
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 78
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListChatappMessageRequestPage) String() string {
	return dara.Prettify(s)
}

func (s ListChatappMessageRequestPage) GoString() string {
	return s.String()
}

func (s *ListChatappMessageRequestPage) GetIndex() *int64 {
	return s.Index
}

func (s *ListChatappMessageRequestPage) GetSize() *int64 {
	return s.Size
}

func (s *ListChatappMessageRequestPage) SetIndex(v int64) *ListChatappMessageRequestPage {
	s.Index = &v
	return s
}

func (s *ListChatappMessageRequestPage) SetSize(v int64) *ListChatappMessageRequestPage {
	s.Size = &v
	return s
}

func (s *ListChatappMessageRequestPage) Validate() error {
	return dara.Validate(s)
}
