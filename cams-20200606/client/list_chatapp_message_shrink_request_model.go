// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatappMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessNumber(v string) *ListChatappMessageShrinkRequest
	GetBusinessNumber() *string
	SetChannelType(v string) *ListChatappMessageShrinkRequest
	GetChannelType() *string
	SetClientAcceptStatus(v string) *ListChatappMessageShrinkRequest
	GetClientAcceptStatus() *string
	SetCustSpaceId(v string) *ListChatappMessageShrinkRequest
	GetCustSpaceId() *string
	SetEndTime(v int64) *ListChatappMessageShrinkRequest
	GetEndTime() *int64
	SetEndTimeStr(v string) *ListChatappMessageShrinkRequest
	GetEndTimeStr() *string
	SetEventAction(v string) *ListChatappMessageShrinkRequest
	GetEventAction() *string
	SetGroupMessageId(v string) *ListChatappMessageShrinkRequest
	GetGroupMessageId() *string
	SetMessageStatus(v string) *ListChatappMessageShrinkRequest
	GetMessageStatus() *string
	SetOwnerId(v int64) *ListChatappMessageShrinkRequest
	GetOwnerId() *int64
	SetPageShrink(v string) *ListChatappMessageShrinkRequest
	GetPageShrink() *string
	SetResourceOwnerAccount(v string) *ListChatappMessageShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListChatappMessageShrinkRequest
	GetResourceOwnerId() *int64
	SetStartTime(v int64) *ListChatappMessageShrinkRequest
	GetStartTime() *int64
	SetStartTimeStr(v string) *ListChatappMessageShrinkRequest
	GetStartTimeStr() *string
	SetTemplateCode(v string) *ListChatappMessageShrinkRequest
	GetTemplateCode() *string
	SetUserNumber(v string) *ListChatappMessageShrinkRequest
	GetUserNumber() *string
}

type ListChatappMessageShrinkRequest struct {
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
	// 2024-01-30 00:00:00
	EndTimeStr *string `json:"EndTimeStr,omitempty" xml:"EndTimeStr,omitempty"`
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
	PageShrink           *string `json:"Page,omitempty" xml:"Page,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 1727057232686
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	StartTimeStr *string `json:"StartTimeStr,omitempty" xml:"StartTimeStr,omitempty"`
	// example:
	//
	// 9938***
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// 86138***
	UserNumber *string `json:"UserNumber,omitempty" xml:"UserNumber,omitempty"`
}

func (s ListChatappMessageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatappMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListChatappMessageShrinkRequest) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *ListChatappMessageShrinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *ListChatappMessageShrinkRequest) GetClientAcceptStatus() *string {
	return s.ClientAcceptStatus
}

func (s *ListChatappMessageShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListChatappMessageShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListChatappMessageShrinkRequest) GetEndTimeStr() *string {
	return s.EndTimeStr
}

func (s *ListChatappMessageShrinkRequest) GetEventAction() *string {
	return s.EventAction
}

func (s *ListChatappMessageShrinkRequest) GetGroupMessageId() *string {
	return s.GroupMessageId
}

func (s *ListChatappMessageShrinkRequest) GetMessageStatus() *string {
	return s.MessageStatus
}

func (s *ListChatappMessageShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListChatappMessageShrinkRequest) GetPageShrink() *string {
	return s.PageShrink
}

func (s *ListChatappMessageShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListChatappMessageShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListChatappMessageShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListChatappMessageShrinkRequest) GetStartTimeStr() *string {
	return s.StartTimeStr
}

func (s *ListChatappMessageShrinkRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ListChatappMessageShrinkRequest) GetUserNumber() *string {
	return s.UserNumber
}

func (s *ListChatappMessageShrinkRequest) SetBusinessNumber(v string) *ListChatappMessageShrinkRequest {
	s.BusinessNumber = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetChannelType(v string) *ListChatappMessageShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetClientAcceptStatus(v string) *ListChatappMessageShrinkRequest {
	s.ClientAcceptStatus = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetCustSpaceId(v string) *ListChatappMessageShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetEndTime(v int64) *ListChatappMessageShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetEndTimeStr(v string) *ListChatappMessageShrinkRequest {
	s.EndTimeStr = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetEventAction(v string) *ListChatappMessageShrinkRequest {
	s.EventAction = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetGroupMessageId(v string) *ListChatappMessageShrinkRequest {
	s.GroupMessageId = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetMessageStatus(v string) *ListChatappMessageShrinkRequest {
	s.MessageStatus = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetOwnerId(v int64) *ListChatappMessageShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetPageShrink(v string) *ListChatappMessageShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetResourceOwnerAccount(v string) *ListChatappMessageShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetResourceOwnerId(v int64) *ListChatappMessageShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetStartTime(v int64) *ListChatappMessageShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetStartTimeStr(v string) *ListChatappMessageShrinkRequest {
	s.StartTimeStr = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetTemplateCode(v string) *ListChatappMessageShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) SetUserNumber(v string) *ListChatappMessageShrinkRequest {
	s.UserNumber = &v
	return s
}

func (s *ListChatappMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
