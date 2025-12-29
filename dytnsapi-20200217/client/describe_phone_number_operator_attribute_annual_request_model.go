// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberOperatorAttributeAnnualRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *DescribePhoneNumberOperatorAttributeAnnualRequest
	GetAuthCode() *string
	SetMask(v string) *DescribePhoneNumberOperatorAttributeAnnualRequest
	GetMask() *string
	SetNumber(v string) *DescribePhoneNumberOperatorAttributeAnnualRequest
	GetNumber() *string
}

type DescribePhoneNumberOperatorAttributeAnnualRequest struct {
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	Mask     *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribePhoneNumberOperatorAttributeAnnualRequest) GetMask() *string {
	return s.Mask
}

func (s *DescribePhoneNumberOperatorAttributeAnnualRequest) GetNumber() *string {
	return s.Number
}

func (s *DescribePhoneNumberOperatorAttributeAnnualRequest) SetAuthCode(v string) *DescribePhoneNumberOperatorAttributeAnnualRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualRequest) SetMask(v string) *DescribePhoneNumberOperatorAttributeAnnualRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualRequest) SetNumber(v string) *DescribePhoneNumberOperatorAttributeAnnualRequest {
	s.Number = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualRequest) Validate() error {
	return dara.Validate(s)
}
