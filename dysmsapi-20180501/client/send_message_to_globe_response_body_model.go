// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageToGlobeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *SendMessageToGlobeResponseBody
	GetFrom() *string
	SetMessageId(v string) *SendMessageToGlobeResponseBody
	GetMessageId() *string
	SetNumberDetail(v *SendMessageToGlobeResponseBodyNumberDetail) *SendMessageToGlobeResponseBody
	GetNumberDetail() *SendMessageToGlobeResponseBodyNumberDetail
	SetRequestId(v string) *SendMessageToGlobeResponseBody
	GetRequestId() *string
	SetResponseCode(v string) *SendMessageToGlobeResponseBody
	GetResponseCode() *string
	SetResponseDescription(v string) *SendMessageToGlobeResponseBody
	GetResponseDescription() *string
	SetSegments(v string) *SendMessageToGlobeResponseBody
	GetSegments() *string
	SetTo(v string) *SendMessageToGlobeResponseBody
	GetTo() *string
}

type SendMessageToGlobeResponseBody struct {
	// The sender ID returned.
	//
	// example:
	//
	// Alicloud321
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The ID of the message.
	//
	// example:
	//
	// 1008030300****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The details about the mobile phone number of the recipient.
	NumberDetail *SendMessageToGlobeResponseBodyNumberDetail `json:"NumberDetail,omitempty" xml:"NumberDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
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
	// The number of messages that incurred fees.
	//
	// example:
	//
	// 1
	Segments *string `json:"Segments,omitempty" xml:"Segments,omitempty"`
	// The mobile phone number to which the message was sent.
	//
	// example:
	//
	// 1380000****
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SendMessageToGlobeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendMessageToGlobeResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeResponseBody) GetFrom() *string {
	return s.From
}

func (s *SendMessageToGlobeResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *SendMessageToGlobeResponseBody) GetNumberDetail() *SendMessageToGlobeResponseBodyNumberDetail {
	return s.NumberDetail
}

func (s *SendMessageToGlobeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendMessageToGlobeResponseBody) GetResponseCode() *string {
	return s.ResponseCode
}

func (s *SendMessageToGlobeResponseBody) GetResponseDescription() *string {
	return s.ResponseDescription
}

func (s *SendMessageToGlobeResponseBody) GetSegments() *string {
	return s.Segments
}

func (s *SendMessageToGlobeResponseBody) GetTo() *string {
	return s.To
}

func (s *SendMessageToGlobeResponseBody) SetFrom(v string) *SendMessageToGlobeResponseBody {
	s.From = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetMessageId(v string) *SendMessageToGlobeResponseBody {
	s.MessageId = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetNumberDetail(v *SendMessageToGlobeResponseBodyNumberDetail) *SendMessageToGlobeResponseBody {
	s.NumberDetail = v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetRequestId(v string) *SendMessageToGlobeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetResponseCode(v string) *SendMessageToGlobeResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetResponseDescription(v string) *SendMessageToGlobeResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetSegments(v string) *SendMessageToGlobeResponseBody {
	s.Segments = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) SetTo(v string) *SendMessageToGlobeResponseBody {
	s.To = &v
	return s
}

func (s *SendMessageToGlobeResponseBody) Validate() error {
	if s.NumberDetail != nil {
		if err := s.NumberDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendMessageToGlobeResponseBodyNumberDetail struct {
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

func (s SendMessageToGlobeResponseBodyNumberDetail) String() string {
	return dara.Prettify(s)
}

func (s SendMessageToGlobeResponseBodyNumberDetail) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) GetCarrier() *string {
	return s.Carrier
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) GetCountry() *string {
	return s.Country
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) GetRegion() *string {
	return s.Region
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) SetCarrier(v string) *SendMessageToGlobeResponseBodyNumberDetail {
	s.Carrier = &v
	return s
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) SetCountry(v string) *SendMessageToGlobeResponseBodyNumberDetail {
	s.Country = &v
	return s
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) SetRegion(v string) *SendMessageToGlobeResponseBodyNumberDetail {
	s.Region = &v
	return s
}

func (s *SendMessageToGlobeResponseBodyNumberDetail) Validate() error {
	return dara.Validate(s)
}
