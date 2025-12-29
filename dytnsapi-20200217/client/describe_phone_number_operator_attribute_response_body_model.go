// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberOperatorAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribePhoneNumberOperatorAttributeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribePhoneNumberOperatorAttributeResponseBody
	GetCode() *string
	SetData(v *DescribePhoneNumberOperatorAttributeResponseBodyData) *DescribePhoneNumberOperatorAttributeResponseBody
	GetData() *DescribePhoneNumberOperatorAttributeResponseBodyData
	SetMessage(v string) *DescribePhoneNumberOperatorAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePhoneNumberOperatorAttributeResponseBody
	GetRequestId() *string
}

type DescribePhoneNumberOperatorAttributeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
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
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *DescribePhoneNumberOperatorAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD135850177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) GetData() *DescribePhoneNumberOperatorAttributeResponseBodyData {
	return s.Data
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) SetAccessDeniedDetail(v string) *DescribePhoneNumberOperatorAttributeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) SetCode(v string) *DescribePhoneNumberOperatorAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) SetData(v *DescribePhoneNumberOperatorAttributeResponseBodyData) *DescribePhoneNumberOperatorAttributeResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) SetMessage(v string) *DescribePhoneNumberOperatorAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) SetRequestId(v string) *DescribePhoneNumberOperatorAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhoneNumberOperatorAttributeResponseBodyData struct {
	// The basic carrier. Valid values:
	//
	// 	- **China Mobile**
	//
	// 	- **China Unicom**
	//
	// 	- **China Telecom**
	//
	// 	- **China Broadnet**
	//
	// example:
	//
	// China Mobile
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	// The actual carrier, including the virtual network operator (VNO). If the phone number involves mobile number portability, the value of this parameter is the carrier after mobile number portability.
	//
	// example:
	//
	// China Mobile
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The city where the phone number is registered.
	//
	// example:
	//
	// hangzhou
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
	//
	// example:
	//
	// zhejiang
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) GetBasicCarrier() *string {
	return s.BasicCarrier
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) GetCarrier() *string {
	return s.Carrier
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) GetCity() *string {
	return s.City
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) GetIsNumberPortability() *bool {
	return s.IsNumberPortability
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) GetNumberSegment() *int64 {
	return s.NumberSegment
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) GetProvince() *string {
	return s.Province
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetBasicCarrier(v string) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetCarrier(v string) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetCity(v string) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.City = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetIsNumberPortability(v bool) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.IsNumberPortability = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetNumberSegment(v int64) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.NumberSegment = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) SetProvince(v string) *DescribePhoneNumberOperatorAttributeResponseBodyData {
	s.Province = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
