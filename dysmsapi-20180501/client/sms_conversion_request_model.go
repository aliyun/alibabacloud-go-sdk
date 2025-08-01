// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmsConversionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversionTime(v int64) *SmsConversionRequest
	GetConversionTime() *int64
	SetDelivered(v bool) *SmsConversionRequest
	GetDelivered() *bool
	SetMessageId(v string) *SmsConversionRequest
	GetMessageId() *string
	SetTo(v string) *SmsConversionRequest
	GetTo() *string
}

type SmsConversionRequest struct {
	// The time when the OTP message was delivered. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// 	- If you leave the parameter empty, the current timestamp is specified by default.
	//
	// 	- If you specify the parameter, the timestamp must be greater than the message sending time and less than the current timestamp.
	//
	// example:
	//
	// 1349055900000
	ConversionTime *int64 `json:"ConversionTime,omitempty" xml:"ConversionTime,omitempty"`
	// Specifies whether customers replied to the OTP message. Valid values: true and false.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Delivered *bool `json:"Delivered,omitempty" xml:"Delivered,omitempty"`
	// The ID of the OTP message.
	//
	// example:
	//
	// 1008030300****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The mobile phone number of the recipient. You must add the dialing code to the beginning of the mobile phone number.
	//
	// For more information, see [Dialing codes](https://help.aliyun.com/document_detail/158400.html).
	//
	// example:
	//
	// 8521245567****
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SmsConversionRequest) String() string {
	return dara.Prettify(s)
}

func (s SmsConversionRequest) GoString() string {
	return s.String()
}

func (s *SmsConversionRequest) GetConversionTime() *int64 {
	return s.ConversionTime
}

func (s *SmsConversionRequest) GetDelivered() *bool {
	return s.Delivered
}

func (s *SmsConversionRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *SmsConversionRequest) GetTo() *string {
	return s.To
}

func (s *SmsConversionRequest) SetConversionTime(v int64) *SmsConversionRequest {
	s.ConversionTime = &v
	return s
}

func (s *SmsConversionRequest) SetDelivered(v bool) *SmsConversionRequest {
	s.Delivered = &v
	return s
}

func (s *SmsConversionRequest) SetMessageId(v string) *SmsConversionRequest {
	s.MessageId = &v
	return s
}

func (s *SmsConversionRequest) SetTo(v string) *SmsConversionRequest {
	s.To = &v
	return s
}

func (s *SmsConversionRequest) Validate() error {
	return dara.Validate(s)
}
