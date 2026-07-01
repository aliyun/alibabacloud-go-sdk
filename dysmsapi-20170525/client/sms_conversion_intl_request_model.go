// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmsConversionIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversionTime(v int64) *SmsConversionIntlRequest
	GetConversionTime() *int64
	SetDelivered(v bool) *SmsConversionIntlRequest
	GetDelivered() *bool
	SetMessageId(v string) *SmsConversionIntlRequest
	GetMessageId() *string
	SetOwnerId(v int64) *SmsConversionIntlRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SmsConversionIntlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SmsConversionIntlRequest
	GetResourceOwnerId() *int64
}

type SmsConversionIntlRequest struct {
	// The timestamp at which the message reached the recipient. Must be a Unix timestamp, expressed as a long integer in milliseconds.
	//
	// - If this field is not specified: defaults to the current timestamp.
	//
	// - If this field is specified: the timestamp must be greater than the send time and less than the current timestamp.
	//
	// example:
	//
	// 1349055900000
	ConversionTime *int64 `json:"ConversionTime,omitempty" xml:"ConversionTime,omitempty"`
	// Set to true if your user replied to the message you sent. Otherwise, set to false.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Delivered *bool `json:"Delivered,omitempty" xml:"Delivered,omitempty"`
	// Message ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1008030300****
	MessageId            *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SmsConversionIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s SmsConversionIntlRequest) GoString() string {
	return s.String()
}

func (s *SmsConversionIntlRequest) GetConversionTime() *int64 {
	return s.ConversionTime
}

func (s *SmsConversionIntlRequest) GetDelivered() *bool {
	return s.Delivered
}

func (s *SmsConversionIntlRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *SmsConversionIntlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SmsConversionIntlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SmsConversionIntlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SmsConversionIntlRequest) SetConversionTime(v int64) *SmsConversionIntlRequest {
	s.ConversionTime = &v
	return s
}

func (s *SmsConversionIntlRequest) SetDelivered(v bool) *SmsConversionIntlRequest {
	s.Delivered = &v
	return s
}

func (s *SmsConversionIntlRequest) SetMessageId(v string) *SmsConversionIntlRequest {
	s.MessageId = &v
	return s
}

func (s *SmsConversionIntlRequest) SetOwnerId(v int64) *SmsConversionIntlRequest {
	s.OwnerId = &v
	return s
}

func (s *SmsConversionIntlRequest) SetResourceOwnerAccount(v string) *SmsConversionIntlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SmsConversionIntlRequest) SetResourceOwnerId(v int64) *SmsConversionIntlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SmsConversionIntlRequest) Validate() error {
	return dara.Validate(s)
}
