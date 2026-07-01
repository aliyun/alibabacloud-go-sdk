// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCSMobileCapableTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPhoneNumbersFile(v string) *CreateRCSMobileCapableTaskRequest
	GetPhoneNumbersFile() *string
	SetSignName(v string) *CreateRCSMobileCapableTaskRequest
	GetSignName() *string
	SetTemplateCode(v string) *CreateRCSMobileCapableTaskRequest
	GetTemplateCode() *string
}

type CreateRCSMobileCapableTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	PhoneNumbersFile *string `json:"PhoneNumbersFile,omitempty" xml:"PhoneNumbersFile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s CreateRCSMobileCapableTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRCSMobileCapableTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRCSMobileCapableTaskRequest) GetPhoneNumbersFile() *string {
	return s.PhoneNumbersFile
}

func (s *CreateRCSMobileCapableTaskRequest) GetSignName() *string {
	return s.SignName
}

func (s *CreateRCSMobileCapableTaskRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateRCSMobileCapableTaskRequest) SetPhoneNumbersFile(v string) *CreateRCSMobileCapableTaskRequest {
	s.PhoneNumbersFile = &v
	return s
}

func (s *CreateRCSMobileCapableTaskRequest) SetSignName(v string) *CreateRCSMobileCapableTaskRequest {
	s.SignName = &v
	return s
}

func (s *CreateRCSMobileCapableTaskRequest) SetTemplateCode(v string) *CreateRCSMobileCapableTaskRequest {
	s.TemplateCode = &v
	return s
}

func (s *CreateRCSMobileCapableTaskRequest) Validate() error {
	return dara.Validate(s)
}
