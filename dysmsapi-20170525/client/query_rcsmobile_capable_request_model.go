// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRCSMobileCapableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPhoneNumbers(v string) *QueryRCSMobileCapableRequest
	GetPhoneNumbers() *string
	SetSignName(v string) *QueryRCSMobileCapableRequest
	GetSignName() *string
	SetTemplateCode(v string) *QueryRCSMobileCapableRequest
	GetTemplateCode() *string
}

type QueryRCSMobileCapableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	PhoneNumbers *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s QueryRCSMobileCapableRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRCSMobileCapableRequest) GoString() string {
	return s.String()
}

func (s *QueryRCSMobileCapableRequest) GetPhoneNumbers() *string {
	return s.PhoneNumbers
}

func (s *QueryRCSMobileCapableRequest) GetSignName() *string {
	return s.SignName
}

func (s *QueryRCSMobileCapableRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *QueryRCSMobileCapableRequest) SetPhoneNumbers(v string) *QueryRCSMobileCapableRequest {
	s.PhoneNumbers = &v
	return s
}

func (s *QueryRCSMobileCapableRequest) SetSignName(v string) *QueryRCSMobileCapableRequest {
	s.SignName = &v
	return s
}

func (s *QueryRCSMobileCapableRequest) SetTemplateCode(v string) *QueryRCSMobileCapableRequest {
	s.TemplateCode = &v
	return s
}

func (s *QueryRCSMobileCapableRequest) Validate() error {
	return dara.Validate(s)
}
