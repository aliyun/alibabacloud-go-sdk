// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberOperatorAttributeAnnualResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBody
	GetCode() *string
	SetData(v *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) *DescribePhoneNumberOperatorAttributeAnnualResponseBody
	GetData() *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData
	SetMessage(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribePhoneNumberOperatorAttributeAnnualResponseBody
	GetSuccess() *bool
}

type DescribePhoneNumberOperatorAttributeAnnualResponseBody struct {
	Code      *string                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) GetData() *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData {
	return s.Data
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) SetCode(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) SetData(v *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) *DescribePhoneNumberOperatorAttributeAnnualResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) SetMessage(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) SetRequestId(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) SetSuccess(v bool) *DescribePhoneNumberOperatorAttributeAnnualResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhoneNumberOperatorAttributeAnnualResponseBodyData struct {
	BasicCarrier      *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	Carrier           *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	City              *string `json:"City,omitempty" xml:"City,omitempty"`
	Mts               *string `json:"Mts,omitempty" xml:"Mts,omitempty"`
	NumberPortability *bool   `json:"NumberPortability,omitempty" xml:"NumberPortability,omitempty"`
	Province          *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) GetBasicCarrier() *string {
	return s.BasicCarrier
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) GetCarrier() *string {
	return s.Carrier
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) GetCity() *string {
	return s.City
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) GetMts() *string {
	return s.Mts
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) GetNumberPortability() *bool {
	return s.NumberPortability
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) GetProvince() *string {
	return s.Province
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) SetBasicCarrier(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) SetCarrier(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) SetCity(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData {
	s.City = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) SetMts(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData {
	s.Mts = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) SetNumberPortability(v bool) *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData {
	s.NumberPortability = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) SetProvince(v string) *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData {
	s.Province = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponseBodyData) Validate() error {
	return dara.Validate(s)
}
