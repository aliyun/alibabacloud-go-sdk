// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSendMessageToGlobeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *BatchSendMessageToGlobeRequest
	GetChannelId() *string
	SetFrom(v string) *BatchSendMessageToGlobeRequest
	GetFrom() *string
	SetMessage(v string) *BatchSendMessageToGlobeRequest
	GetMessage() *string
	SetTaskId(v string) *BatchSendMessageToGlobeRequest
	GetTaskId() *string
	SetTo(v string) *BatchSendMessageToGlobeRequest
	GetTo() *string
	SetType(v string) *BatchSendMessageToGlobeRequest
	GetType() *string
	SetValidityPeriod(v int64) *BatchSendMessageToGlobeRequest
	GetValidityPeriod() *int64
}

type BatchSendMessageToGlobeRequest struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// sms-djnfjn344
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The mobile phone number of the sender. You can specify the sender ID when you call the API operation. The sender ID can contain only digits and letters. If the sender ID contains letters, it can be a maximum of 11 characters in length. If the sender ID contains only digits, it can be a maximum of 15 characters in length.
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
	// [\\"message to 931520581****\\",\\"message to 931530581****\\",\\"message to 931540581****\\",\\"message to 931550581****\\"]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the messaging campaign. It must be 1 to 255 characters in length. The ID is the value of the TaskId field in the delivery receipt of the message.
	//
	// example:
	//
	// 123789****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The mobile phone number of the recipient. You must add the dialing code to the beginning of each mobile phone number.
	//
	// For more information, see [Dialing codes](https://help.aliyun.com/document_detail/158400.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"931520581****\\",\\"931530581****\\",\\"931540581****\\",\\"931550581****\\"]
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// The type of the message. Valid values:
	//
	// 	- **NOTIFY**: notification
	//
	// 	- **MKT**: promotional message
	//
	// example:
	//
	// NOTIFY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The validity period of the message. Unit: seconds.
	//
	// example:
	//
	// 600
	ValidityPeriod *int64 `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty"`
}

func (s BatchSendMessageToGlobeRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMessageToGlobeRequest) GoString() string {
	return s.String()
}

func (s *BatchSendMessageToGlobeRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *BatchSendMessageToGlobeRequest) GetFrom() *string {
	return s.From
}

func (s *BatchSendMessageToGlobeRequest) GetMessage() *string {
	return s.Message
}

func (s *BatchSendMessageToGlobeRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchSendMessageToGlobeRequest) GetTo() *string {
	return s.To
}

func (s *BatchSendMessageToGlobeRequest) GetType() *string {
	return s.Type
}

func (s *BatchSendMessageToGlobeRequest) GetValidityPeriod() *int64 {
	return s.ValidityPeriod
}

func (s *BatchSendMessageToGlobeRequest) SetChannelId(v string) *BatchSendMessageToGlobeRequest {
	s.ChannelId = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetFrom(v string) *BatchSendMessageToGlobeRequest {
	s.From = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetMessage(v string) *BatchSendMessageToGlobeRequest {
	s.Message = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetTaskId(v string) *BatchSendMessageToGlobeRequest {
	s.TaskId = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetTo(v string) *BatchSendMessageToGlobeRequest {
	s.To = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetType(v string) *BatchSendMessageToGlobeRequest {
	s.Type = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) SetValidityPeriod(v int64) *BatchSendMessageToGlobeRequest {
	s.ValidityPeriod = &v
	return s
}

func (s *BatchSendMessageToGlobeRequest) Validate() error {
	return dara.Validate(s)
}
