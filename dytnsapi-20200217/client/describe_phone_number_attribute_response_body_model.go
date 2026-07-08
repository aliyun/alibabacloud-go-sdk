// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePhoneNumberAttributeResponseBody
	GetCode() *string
	SetMessage(v string) *DescribePhoneNumberAttributeResponseBody
	GetMessage() *string
	SetPhoneNumberAttribute(v *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) *DescribePhoneNumberAttributeResponseBody
	GetPhoneNumberAttribute() *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute
	SetRequestId(v string) *DescribePhoneNumberAttributeResponseBody
	GetRequestId() *string
}

type DescribePhoneNumberAttributeResponseBody struct {
	Code                 *string                                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message              *string                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PhoneNumberAttribute *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute `json:"PhoneNumberAttribute,omitempty" xml:"PhoneNumberAttribute,omitempty" type:"Struct"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePhoneNumberAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePhoneNumberAttributeResponseBody) GetPhoneNumberAttribute() *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	return s.PhoneNumberAttribute
}

func (s *DescribePhoneNumberAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePhoneNumberAttributeResponseBody) SetCode(v string) *DescribePhoneNumberAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBody) SetMessage(v string) *DescribePhoneNumberAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBody) SetPhoneNumberAttribute(v *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) *DescribePhoneNumberAttributeResponseBody {
	s.PhoneNumberAttribute = v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBody) SetRequestId(v string) *DescribePhoneNumberAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBody) Validate() error {
	if s.PhoneNumberAttribute != nil {
		if err := s.PhoneNumberAttribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute struct {
	BasicCarrier        *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	Carrier             *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	City                *string `json:"City,omitempty" xml:"City,omitempty"`
	IsNumberPortability *bool   `json:"IsNumberPortability,omitempty" xml:"IsNumberPortability,omitempty"`
	NumberSegment       *int64  `json:"NumberSegment,omitempty" xml:"NumberSegment,omitempty"`
	Province            *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) GetBasicCarrier() *string {
	return s.BasicCarrier
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) GetCarrier() *string {
	return s.Carrier
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) GetCity() *string {
	return s.City
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) GetIsNumberPortability() *bool {
	return s.IsNumberPortability
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) GetNumberSegment() *int64 {
	return s.NumberSegment
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) GetProvince() *string {
	return s.Province
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetBasicCarrier(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.BasicCarrier = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetCarrier(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetCity(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.City = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetIsNumberPortability(v bool) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.IsNumberPortability = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetNumberSegment(v int64) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.NumberSegment = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) SetProvince(v string) *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute {
	s.Province = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute) Validate() error {
	return dara.Validate(s)
}
