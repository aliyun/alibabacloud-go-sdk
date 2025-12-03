// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageWithTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *SendMessageWithTemplateResponseBody
	GetMessageId() *string
	SetNumberDetail(v *SendMessageWithTemplateResponseBodyNumberDetail) *SendMessageWithTemplateResponseBody
	GetNumberDetail() *SendMessageWithTemplateResponseBodyNumberDetail
	SetRequestId(v string) *SendMessageWithTemplateResponseBody
	GetRequestId() *string
	SetResponseCode(v string) *SendMessageWithTemplateResponseBody
	GetResponseCode() *string
	SetResponseDescription(v string) *SendMessageWithTemplateResponseBody
	GetResponseDescription() *string
	SetSegments(v string) *SendMessageWithTemplateResponseBody
	GetSegments() *string
	SetTo(v string) *SendMessageWithTemplateResponseBody
	GetTo() *string
}

type SendMessageWithTemplateResponseBody struct {
	// The ID of the message.
	//
	// example:
	//
	// 1**************3
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The details about the mobile phone number of the recipient.
	NumberDetail *SendMessageWithTemplateResponseBodyNumberDetail `json:"NumberDetail,omitempty" xml:"NumberDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8D23D6
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
	// The mobile phone number to which the message was sent. The dialing code was added to the beginning of the mobile phone number. Example: 861503871\\*\\*\\*\\*.
	//
	// example:
	//
	// 861503871****
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SendMessageWithTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendMessageWithTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageWithTemplateResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *SendMessageWithTemplateResponseBody) GetNumberDetail() *SendMessageWithTemplateResponseBodyNumberDetail {
	return s.NumberDetail
}

func (s *SendMessageWithTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendMessageWithTemplateResponseBody) GetResponseCode() *string {
	return s.ResponseCode
}

func (s *SendMessageWithTemplateResponseBody) GetResponseDescription() *string {
	return s.ResponseDescription
}

func (s *SendMessageWithTemplateResponseBody) GetSegments() *string {
	return s.Segments
}

func (s *SendMessageWithTemplateResponseBody) GetTo() *string {
	return s.To
}

func (s *SendMessageWithTemplateResponseBody) SetMessageId(v string) *SendMessageWithTemplateResponseBody {
	s.MessageId = &v
	return s
}

func (s *SendMessageWithTemplateResponseBody) SetNumberDetail(v *SendMessageWithTemplateResponseBodyNumberDetail) *SendMessageWithTemplateResponseBody {
	s.NumberDetail = v
	return s
}

func (s *SendMessageWithTemplateResponseBody) SetRequestId(v string) *SendMessageWithTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageWithTemplateResponseBody) SetResponseCode(v string) *SendMessageWithTemplateResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *SendMessageWithTemplateResponseBody) SetResponseDescription(v string) *SendMessageWithTemplateResponseBody {
	s.ResponseDescription = &v
	return s
}

func (s *SendMessageWithTemplateResponseBody) SetSegments(v string) *SendMessageWithTemplateResponseBody {
	s.Segments = &v
	return s
}

func (s *SendMessageWithTemplateResponseBody) SetTo(v string) *SendMessageWithTemplateResponseBody {
	s.To = &v
	return s
}

func (s *SendMessageWithTemplateResponseBody) Validate() error {
	if s.NumberDetail != nil {
		if err := s.NumberDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendMessageWithTemplateResponseBodyNumberDetail struct {
	// The carrier that owns the mobile phone number.
	//
	// example:
	//
	// China Mobile
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
	// Nanjing, Jiangsu
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SendMessageWithTemplateResponseBodyNumberDetail) String() string {
	return dara.Prettify(s)
}

func (s SendMessageWithTemplateResponseBodyNumberDetail) GoString() string {
	return s.String()
}

func (s *SendMessageWithTemplateResponseBodyNumberDetail) GetCarrier() *string {
	return s.Carrier
}

func (s *SendMessageWithTemplateResponseBodyNumberDetail) GetCountry() *string {
	return s.Country
}

func (s *SendMessageWithTemplateResponseBodyNumberDetail) GetRegion() *string {
	return s.Region
}

func (s *SendMessageWithTemplateResponseBodyNumberDetail) SetCarrier(v string) *SendMessageWithTemplateResponseBodyNumberDetail {
	s.Carrier = &v
	return s
}

func (s *SendMessageWithTemplateResponseBodyNumberDetail) SetCountry(v string) *SendMessageWithTemplateResponseBodyNumberDetail {
	s.Country = &v
	return s
}

func (s *SendMessageWithTemplateResponseBodyNumberDetail) SetRegion(v string) *SendMessageWithTemplateResponseBodyNumberDetail {
	s.Region = &v
	return s
}

func (s *SendMessageWithTemplateResponseBodyNumberDetail) Validate() error {
	return dara.Validate(s)
}
