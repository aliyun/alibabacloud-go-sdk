// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberOperatorAttributeAnnualUseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody
	GetCode() *string
	SetData(v *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody
	GetData() *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData
	SetMessage(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody
	GetRequestId() *string
}

type DescribePhoneNumberOperatorAttributeAnnualUseResponseBody struct {
	AccessDeniedDetail *string                                                        `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) GetData() *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData {
	return s.Data
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) SetAccessDeniedDetail(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) SetCode(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) SetData(v *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) SetMessage(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) SetRequestId(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData struct {
	BasicCarrier        *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	Carrier             *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	City                *string `json:"City,omitempty" xml:"City,omitempty"`
	IsNumberPortability *bool   `json:"IsNumberPortability,omitempty" xml:"IsNumberPortability,omitempty"`
	NumberSegment       *int64  `json:"NumberSegment,omitempty" xml:"NumberSegment,omitempty"`
	Province            *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) GetBasicCarrier() *string {
	return s.BasicCarrier
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) GetCarrier() *string {
	return s.Carrier
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) GetCity() *string {
	return s.City
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) GetIsNumberPortability() *bool {
	return s.IsNumberPortability
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) GetNumberSegment() *int64 {
	return s.NumberSegment
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) GetProvince() *string {
	return s.Province
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) SetBasicCarrier(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) SetCarrier(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) SetCity(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData {
	s.City = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) SetIsNumberPortability(v bool) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData {
	s.IsNumberPortability = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) SetNumberSegment(v int64) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData {
	s.NumberSegment = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) SetProvince(v string) *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData {
	s.Province = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
