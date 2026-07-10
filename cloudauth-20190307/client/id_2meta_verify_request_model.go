// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifyNum(v string) *Id2MetaVerifyRequest
	GetIdentifyNum() *string
	SetParamType(v string) *Id2MetaVerifyRequest
	GetParamType() *string
	SetUserName(v string) *Id2MetaVerifyRequest
	GetUserName() *string
}

type Id2MetaVerifyRequest struct {
	// The ID card number.
	//
	// Note
	//
	// Only second-generation resident ID card numbers and Hong Kong, Macao, and Taiwan residence permit numbers are supported.
	//
	// - If ParamType is set to normal, enter the ID card number in plaintext.
	//
	// - If ParamType is set to md5, the format is: first 6 digits of the ID card number (plaintext) + date of birth (ciphertext) + last 4 digits of the ID card number (plaintext).
	//
	// example:
	//
	// 明文：429001********8211
	//
	// 密文：
	//
	// 42900132fa7bcd874161bea8ec8fd98f390ec98211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
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
	// - If ParamType is set to normal, enter the name in plaintext.
	//
	// - If ParamType is set to md5, the format is: first character of the name (ciphertext) + remaining characters of the name (plaintext).
	//
	// example:
	//
	// 明文：张三
	//
	// 密文：
	//
	// 6499fc7409049355527ef6a2ba5706b8三​
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Id2MetaVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaVerifyRequest) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *Id2MetaVerifyRequest) GetParamType() *string {
	return s.ParamType
}

func (s *Id2MetaVerifyRequest) GetUserName() *string {
	return s.UserName
}

func (s *Id2MetaVerifyRequest) SetIdentifyNum(v string) *Id2MetaVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Id2MetaVerifyRequest) SetParamType(v string) *Id2MetaVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Id2MetaVerifyRequest) SetUserName(v string) *Id2MetaVerifyRequest {
	s.UserName = &v
	return s
}

func (s *Id2MetaVerifyRequest) Validate() error {
	return dara.Validate(s)
}
