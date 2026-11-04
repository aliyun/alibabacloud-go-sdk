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
	// Note: Only the numbers of second-generation resident identity cards and residence permits for Hong Kong, Macao, and Taiwan residents are supported.
	//
	// - If you set paramType to normal, enter the ID card number in plaintext.
	//
	// - If you set paramType to md5, enter the first 6 digits of the ID card number in plaintext, the date of birth in ciphertext, and the last 4 digits of the ID card number in plaintext.
	//
	// example:
	//
	// Plaintext: 429001********8211
	//
	// Ciphertext:
	//
	// 42900132fa7bcd874161bea8ec8fd98f390ec98211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// The encryption method. Valid values:
	//
	// - normal: plaintext without encryption
	//
	// - md5: MD5 encryption
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The name.
	//
	// - If you set paramType to normal, enter the name in plaintext.
	//
	// - If you set paramType to md5, enter the first character of the name in ciphertext and the rest of the name in plaintext.
	//
	// example:
	//
	// Plaintext: Zhang San
	//
	// Ciphertext:
	//
	// 6499fc7409049355527ef6a2ba5706b8San
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
