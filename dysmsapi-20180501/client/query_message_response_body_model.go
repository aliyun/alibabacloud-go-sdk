// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *QueryMessageResponseBody
	GetErrorCode() *string
	SetErrorDescription(v string) *QueryMessageResponseBody
	GetErrorDescription() *string
	SetMessage(v string) *QueryMessageResponseBody
	GetMessage() *string
	SetMessageId(v string) *QueryMessageResponseBody
	GetMessageId() *string
	SetNumberDetail(v *QueryMessageResponseBodyNumberDetail) *QueryMessageResponseBody
	GetNumberDetail() *QueryMessageResponseBodyNumberDetail
	SetReceiveDate(v string) *QueryMessageResponseBody
	GetReceiveDate() *string
	SetRequestId(v string) *QueryMessageResponseBody
	GetRequestId() *string
	SetResponseCode(v string) *QueryMessageResponseBody
	GetResponseCode() *string
	SetResponseDescription(v string) *QueryMessageResponseBody
	GetResponseDescription() *string
	SetSendDate(v string) *QueryMessageResponseBody
	GetSendDate() *string
	SetStatus(v string) *QueryMessageResponseBody
	GetStatus() *string
	SetTo(v string) *QueryMessageResponseBody
	GetTo() *string
}

type QueryMessageResponseBody struct {
	// The status code of the message.
	//
	// example:
	//
	// DELIVERED
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// success
	ErrorDescription *string `json:"ErrorDescription,omitempty" xml:"ErrorDescription,omitempty"`
	// The content of the message.
	//
	// example:
	//
	// Hello!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the message.
	//
	// example:
	//
	// 1008030xxx3003
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The details about the mobile phone number.
	NumberDetail *QueryMessageResponseBodyNumberDetail `json:"NumberDetail,omitempty" xml:"NumberDetail,omitempty" type:"Struct"`
	// The time when the delivery receipt was received from the carrier.
	//
	// example:
	//
	// Mon, 24 Dec 2018 16:58:22 +0800
	ReceiveDate *string `json:"ReceiveDate,omitempty" xml:"ReceiveDate,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8D28D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code of the delivery request.
	//
	// example:
	//
	// OK
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// The description of the delivery request status.
	//
	// example:
	//
	// The SMS Send Request was accepted
	ResponseDescription *string `json:"ResponseDescription,omitempty" xml:"ResponseDescription,omitempty"`
	// The time when the message was sent to the carrier.
	//
	// example:
	//
	// Mon, 24 Dec 2018 16:58:22 +0800
	SendDate *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	// The delivery status of the message.
	//
	// 	- 1: The message was sent.
	//
	// 	- 2: The message failed to be sent.
	//
	// 	- 3: The message is being sent.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The mobile phone number to which the message was sent.
	//
	// example:
	//
	// 6581xxx810
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s QueryMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMessageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryMessageResponseBody) GetErrorDescription() *string {
	return s.ErrorDescription
}

func (s *QueryMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMessageResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *QueryMessageResponseBody) GetNumberDetail() *QueryMessageResponseBodyNumberDetail {
	return s.NumberDetail
}

func (s *QueryMessageResponseBody) GetReceiveDate() *string {
	return s.ReceiveDate
}

func (s *QueryMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMessageResponseBody) GetResponseCode() *string {
	return s.ResponseCode
}

func (s *QueryMessageResponseBody) GetResponseDescription() *string {
	return s.ResponseDescription
}

func (s *QueryMessageResponseBody) GetSendDate() *string {
	return s.SendDate
}

func (s *QueryMessageResponseBody) GetStatus() *string {
	return s.Status
}

func (s *QueryMessageResponseBody) GetTo() *string {
	return s.To
}

func (s *QueryMessageResponseBody) SetErrorCode(v string) *QueryMessageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMessageResponseBody) SetErrorDescription(v string) *QueryMessageResponseBody {
	s.ErrorDescription = &v
	return s
}

func (s *QueryMessageResponseBody) SetMessage(v string) *QueryMessageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMessageResponseBody) SetMessageId(v string) *QueryMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *QueryMessageResponseBody) SetNumberDetail(v *QueryMessageResponseBodyNumberDetail) *QueryMessageResponseBody {
	s.NumberDetail = v
	return s
}

func (s *QueryMessageResponseBody) SetReceiveDate(v string) *QueryMessageResponseBody {
	s.ReceiveDate = &v
	return s
}

func (s *QueryMessageResponseBody) SetRequestId(v string) *QueryMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMessageResponseBody) SetResponseCode(v string) *QueryMessageResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *QueryMessageResponseBody) SetResponseDescription(v string) *QueryMessageResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *QueryMessageResponseBody) SetSendDate(v string) *QueryMessageResponseBody {
	s.SendDate = &v
	return s
}

func (s *QueryMessageResponseBody) SetStatus(v string) *QueryMessageResponseBody {
	s.Status = &v
	return s
}

func (s *QueryMessageResponseBody) SetTo(v string) *QueryMessageResponseBody {
	s.To = &v
	return s
}

func (s *QueryMessageResponseBody) Validate() error {
	if s.NumberDetail != nil {
		if err := s.NumberDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMessageResponseBodyNumberDetail struct {
	// The carrier that owns the mobile phone number.
	//
	// example:
	//
	// CMI
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The country to which the mobile phone number belongs.
	//
	// example:
	//
	// China
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The region to which the mobile phone number belongs.
	//
	// example:
	//
	// HongKong
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s QueryMessageResponseBodyNumberDetail) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageResponseBodyNumberDetail) GoString() string {
	return s.String()
}

func (s *QueryMessageResponseBodyNumberDetail) GetCarrier() *string {
	return s.Carrier
}

func (s *QueryMessageResponseBodyNumberDetail) GetCountry() *string {
	return s.Country
}

func (s *QueryMessageResponseBodyNumberDetail) GetRegion() *string {
	return s.Region
}

func (s *QueryMessageResponseBodyNumberDetail) SetCarrier(v string) *QueryMessageResponseBodyNumberDetail {
	s.Carrier = &v
	return s
}

func (s *QueryMessageResponseBodyNumberDetail) SetCountry(v string) *QueryMessageResponseBodyNumberDetail {
	s.Country = &v
	return s
}

func (s *QueryMessageResponseBodyNumberDetail) SetRegion(v string) *QueryMessageResponseBodyNumberDetail {
	s.Region = &v
	return s
}

func (s *QueryMessageResponseBodyNumberDetail) Validate() error {
	return dara.Validate(s)
}
