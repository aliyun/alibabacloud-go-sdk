// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberOperatorAttributeAnnualUseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *DescribePhoneNumberOperatorAttributeAnnualUseRequest
	GetAuthCode() *string
	SetInputNumber(v string) *DescribePhoneNumberOperatorAttributeAnnualUseRequest
	GetInputNumber() *string
	SetMask(v string) *DescribePhoneNumberOperatorAttributeAnnualUseRequest
	GetMask() *string
}

type DescribePhoneNumberOperatorAttributeAnnualUseRequest struct {
	// This parameter is required.
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// This parameter is required.
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseRequest) GetMask() *string {
	return s.Mask
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseRequest) SetAuthCode(v string) *DescribePhoneNumberOperatorAttributeAnnualUseRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseRequest) SetInputNumber(v string) *DescribePhoneNumberOperatorAttributeAnnualUseRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseRequest) SetMask(v string) *DescribePhoneNumberOperatorAttributeAnnualUseRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseRequest) Validate() error {
	return dara.Validate(s)
}
