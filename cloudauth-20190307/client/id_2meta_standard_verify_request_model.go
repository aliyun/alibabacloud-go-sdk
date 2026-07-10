// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaStandardVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifyNum(v string) *Id2MetaStandardVerifyRequest
	GetIdentifyNum() *string
	SetParamType(v string) *Id2MetaStandardVerifyRequest
	GetParamType() *string
	SetUserName(v string) *Id2MetaStandardVerifyRequest
	GetUserName() *string
}

type Id2MetaStandardVerifyRequest struct {
	// The ID card number.
	//
	// - If ParamType is set to normal, enter the ID card number in plaintext.
	//
	// - If ParamType is set to md5, the format is: first 6 digits of the ID card number (plaintext) + date of birth (ciphertext) + last 4 digits of the ID card number (plaintext).
	//
	// example:
	//
	// 4****************1
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// The parameter type. Valid values:
	//
	// - normal: not encrypted.
	//
	// - md5: MD5-encrypted.
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
	// 张*
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Id2MetaStandardVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaStandardVerifyRequest) GoString() string {
	return s.String()
}

func (s *Id2MetaStandardVerifyRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *Id2MetaStandardVerifyRequest) GetParamType() *string {
	return s.ParamType
}

func (s *Id2MetaStandardVerifyRequest) GetUserName() *string {
	return s.UserName
}

func (s *Id2MetaStandardVerifyRequest) SetIdentifyNum(v string) *Id2MetaStandardVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Id2MetaStandardVerifyRequest) SetParamType(v string) *Id2MetaStandardVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Id2MetaStandardVerifyRequest) SetUserName(v string) *Id2MetaStandardVerifyRequest {
	s.UserName = &v
	return s
}

func (s *Id2MetaStandardVerifyRequest) Validate() error {
	return dara.Validate(s)
}
