// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendRCSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutId(v string) *SendRCSRequest
	GetOutId() *string
	SetPhoneNumbers(v string) *SendRCSRequest
	GetPhoneNumbers() *string
	SetSignName(v string) *SendRCSRequest
	GetSignName() *string
	SetTemplateCode(v string) *SendRCSRequest
	GetTemplateCode() *string
	SetTemplateParam(v string) *SendRCSRequest
	GetTemplateParam() *string
}

type SendRCSRequest struct {
	// example:
	//
	// 示例值
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
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
	// 示例值示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// 示例值示例值
	TemplateParam *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
}

func (s SendRCSRequest) String() string {
	return dara.Prettify(s)
}

func (s SendRCSRequest) GoString() string {
	return s.String()
}

func (s *SendRCSRequest) GetOutId() *string {
	return s.OutId
}

func (s *SendRCSRequest) GetPhoneNumbers() *string {
	return s.PhoneNumbers
}

func (s *SendRCSRequest) GetSignName() *string {
	return s.SignName
}

func (s *SendRCSRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *SendRCSRequest) GetTemplateParam() *string {
	return s.TemplateParam
}

func (s *SendRCSRequest) SetOutId(v string) *SendRCSRequest {
	s.OutId = &v
	return s
}

func (s *SendRCSRequest) SetPhoneNumbers(v string) *SendRCSRequest {
	s.PhoneNumbers = &v
	return s
}

func (s *SendRCSRequest) SetSignName(v string) *SendRCSRequest {
	s.SignName = &v
	return s
}

func (s *SendRCSRequest) SetTemplateCode(v string) *SendRCSRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendRCSRequest) SetTemplateParam(v string) *SendRCSRequest {
	s.TemplateParam = &v
	return s
}

func (s *SendRCSRequest) Validate() error {
	return dara.Validate(s)
}
