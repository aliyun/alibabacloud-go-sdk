// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaSimpleVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifyNum(v string) *Mobile3MetaSimpleVerifyRequest
	GetIdentifyNum() *string
	SetMobile(v string) *Mobile3MetaSimpleVerifyRequest
	GetMobile() *string
	SetParamType(v string) *Mobile3MetaSimpleVerifyRequest
	GetParamType() *string
	SetUserName(v string) *Mobile3MetaSimpleVerifyRequest
	GetUserName() *string
}

type Mobile3MetaSimpleVerifyRequest struct {
	// ID number:
	//
	// Note
	//
	// Only supports the ID numbers of second-generation resident IDs and Hong Kong, Macao, and Taiwan residence permits.
	//
	// - When paramType is normal: enter the plaintext ID number.
	//
	// - When paramType is md5: enter the encrypted ID number.
	//
	// example:
	//
	// 明文：429001********8211
	//
	// 密文：
	//
	// 32fa7bcd874161bea8ec8fd98f390ec9
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// Mobile phone number:
	//
	// - When paramType is normal: enter the plaintext mobile phone number.
	//
	// - When paramType is md5: enter the encrypted mobile phone number.
	//
	// example:
	//
	// 明文：186****2055
	//
	// 密文：
	//
	// 849169cd3b20621c1c78bd61a11a4fc2
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Encryption method:
	//
	// - normal: plaintext, not encrypted
	//
	// - md5: MD5 encryption
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Name:
	//
	// - When paramType is normal: enter the plaintext name.
	//
	// - When paramType is md5: enter the encrypted name.
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

func (s Mobile3MetaSimpleVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaSimpleVerifyRequest) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleVerifyRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *Mobile3MetaSimpleVerifyRequest) GetMobile() *string {
	return s.Mobile
}

func (s *Mobile3MetaSimpleVerifyRequest) GetParamType() *string {
	return s.ParamType
}

func (s *Mobile3MetaSimpleVerifyRequest) GetUserName() *string {
	return s.UserName
}

func (s *Mobile3MetaSimpleVerifyRequest) SetIdentifyNum(v string) *Mobile3MetaSimpleVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyRequest) SetMobile(v string) *Mobile3MetaSimpleVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyRequest) SetParamType(v string) *Mobile3MetaSimpleVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyRequest) SetUserName(v string) *Mobile3MetaSimpleVerifyRequest {
	s.UserName = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyRequest) Validate() error {
	return dara.Validate(s)
}
