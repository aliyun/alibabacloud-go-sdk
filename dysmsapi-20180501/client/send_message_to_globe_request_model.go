// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageToGlobeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *SendMessageToGlobeRequest
	GetChannelId() *string
	SetFrom(v string) *SendMessageToGlobeRequest
	GetFrom() *string
	SetMessage(v string) *SendMessageToGlobeRequest
	GetMessage() *string
	SetTaskId(v string) *SendMessageToGlobeRequest
	GetTaskId() *string
	SetTo(v string) *SendMessageToGlobeRequest
	GetTo() *string
	SetType(v string) *SendMessageToGlobeRequest
	GetType() *string
	SetValidityPeriod(v int64) *SendMessageToGlobeRequest
	GetValidityPeriod() *int64
}

type SendMessageToGlobeRequest struct {
	// The ID of the channel.
	//
	// example:
	//
	// 3790
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The mobile phone number of the sender. You can also specify a sender ID. The sender ID can contain both letters and digits. If it does, the ID must be between 1 to 11 characters in length. If the sender ID contains only digits, it must be 1 to 15 characters in length.
	//
	// example:
	//
	// Alicloud321
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The content of the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hello
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the messaging campaign. It must be 1 to 255 characters in length. The ID is the value of the TaskId field in the delivery receipt of the message.
	//
	// example:
	//
	// 123****789
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The mobile phone number to which the message is sent. You must add the dialing code to the beginning of the mobile phone number. Example: 8521245567\\*\\*\\*\\*.
	//
	// For more information, see [Dialing codes](https://www.alibabacloud.com/help/en/sms/product-overview/dialing-codes?spm=a2c63.p38356.0.0.48b940a1PFYRMz).
	//
	// >  You cannot call the SendMessageToGlobe operation to send messages to the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8521245567****
	To   *string `json:"To,omitempty" xml:"To,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The validity period of the message. Unit: seconds.
	//
	// example:
	//
	// 600
	ValidityPeriod *int64 `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty"`
}

func (s SendMessageToGlobeRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMessageToGlobeRequest) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *SendMessageToGlobeRequest) GetFrom() *string {
	return s.From
}

func (s *SendMessageToGlobeRequest) GetMessage() *string {
	return s.Message
}

func (s *SendMessageToGlobeRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SendMessageToGlobeRequest) GetTo() *string {
	return s.To
}

func (s *SendMessageToGlobeRequest) GetType() *string {
	return s.Type
}

func (s *SendMessageToGlobeRequest) GetValidityPeriod() *int64 {
	return s.ValidityPeriod
}

func (s *SendMessageToGlobeRequest) SetChannelId(v string) *SendMessageToGlobeRequest {
	s.ChannelId = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetFrom(v string) *SendMessageToGlobeRequest {
	s.From = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetMessage(v string) *SendMessageToGlobeRequest {
	s.Message = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetTaskId(v string) *SendMessageToGlobeRequest {
	s.TaskId = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetTo(v string) *SendMessageToGlobeRequest {
	s.To = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetType(v string) *SendMessageToGlobeRequest {
	s.Type = &v
	return s
}

func (s *SendMessageToGlobeRequest) SetValidityPeriod(v int64) *SendMessageToGlobeRequest {
	s.ValidityPeriod = &v
	return s
}

func (s *SendMessageToGlobeRequest) Validate() error {
	return dara.Validate(s)
}
