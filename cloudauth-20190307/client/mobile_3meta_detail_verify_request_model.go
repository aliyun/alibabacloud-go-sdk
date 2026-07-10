// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaDetailVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifyNum(v string) *Mobile3MetaDetailVerifyRequest
	GetIdentifyNum() *string
	SetMobile(v string) *Mobile3MetaDetailVerifyRequest
	GetMobile() *string
	SetParamType(v string) *Mobile3MetaDetailVerifyRequest
	GetParamType() *string
	SetUserName(v string) *Mobile3MetaDetailVerifyRequest
	GetUserName() *string
}

type Mobile3MetaDetailVerifyRequest struct {
	// The ID card number.
	//
	// > Note: Only second-generation resident ID card numbers and Hong Kong, Macao, or Taiwan residence permit numbers are supported.
	//
	// - If ParamType is set to normal, pass in the ID card number in plaintext.
	//
	// - If ParamType is set to md5, pass in the ID card number in ciphertext.
	//
	// example:
	//
	// 明文：429001********8211
	//
	// 密文：
	//
	// 32fa7bcd874161bea8ec8fd98f390ec9
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// The phone number.
	//
	// - If ParamType is set to normal, pass in the phone number in plaintext.
	//
	// - If ParamType is set to md5, pass in the phone number in ciphertext.
	//
	// example:
	//
	// 明文：186****2055
	//
	// 密文：
	//
	// 849169cd3b20621c1c78bd61a11a4fc2
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The encryption method. Valid values:
	//
	// - normal: plaintext without encryption
	//
	// - md5: MD5 encryption.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The name.
	//
	// - If ParamType is set to normal, pass in the name in plaintext.
	//
	// - If ParamType is set to md5, pass in the name in ciphertext.
	//
	// example:
	//
	// 明文：张三
	//
	// 密文：
	//
	// 32fa7bcd874161bea8ec8fd98f390ec9
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Mobile3MetaDetailVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaDetailVerifyRequest) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailVerifyRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *Mobile3MetaDetailVerifyRequest) GetMobile() *string {
	return s.Mobile
}

func (s *Mobile3MetaDetailVerifyRequest) GetParamType() *string {
	return s.ParamType
}

func (s *Mobile3MetaDetailVerifyRequest) GetUserName() *string {
	return s.UserName
}

func (s *Mobile3MetaDetailVerifyRequest) SetIdentifyNum(v string) *Mobile3MetaDetailVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Mobile3MetaDetailVerifyRequest) SetMobile(v string) *Mobile3MetaDetailVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *Mobile3MetaDetailVerifyRequest) SetParamType(v string) *Mobile3MetaDetailVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Mobile3MetaDetailVerifyRequest) SetUserName(v string) *Mobile3MetaDetailVerifyRequest {
	s.UserName = &v
	return s
}

func (s *Mobile3MetaDetailVerifyRequest) Validate() error {
	return dara.Validate(s)
}
