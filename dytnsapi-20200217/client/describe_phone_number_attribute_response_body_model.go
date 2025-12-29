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
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **InvalidParameter**: The specified phone number is invalid or the parameter format is invalid.
	//
	// 	- **PhoneNumberNotfound**: No attribute information can be found for the specified phone number.
	//
	// 	- **isp.UNKNOWN**: An unknown exception occurred.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The attribute information about the phone number.
	PhoneNumberAttribute *DescribePhoneNumberAttributeResponseBodyPhoneNumberAttribute `json:"PhoneNumberAttribute,omitempty" xml:"PhoneNumberAttribute,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The basic carrier. Valid values:
	//
	// 	- **China Mobile**
	//
	// 	- **China Unicom**
	//
	// 	- **China Telecom**
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	// The actual carrier, including the virtual network operator (VNO). If the phone number involves mobile number portability, the value of this parameter is the carrier after mobile number portability.
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The city where the phone number is registered.
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// Indicates whether the phone number involves mobile number portability. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsNumberPortability *bool `json:"IsNumberPortability,omitempty" xml:"IsNumberPortability,omitempty"`
	// The number segment to which the phone number belongs.
	//
	// example:
	//
	// 139
	NumberSegment *int64 `json:"NumberSegment,omitempty" xml:"NumberSegment,omitempty"`
	// The province where the phone number is registered.
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
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
