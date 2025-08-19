// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile2MetaVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobile(v string) *Mobile2MetaVerifyRequest
	GetMobile() *string
	SetParamType(v string) *Mobile2MetaVerifyRequest
	GetParamType() *string
	SetUserName(v string) *Mobile2MetaVerifyRequest
	GetUserName() *string
}

type Mobile2MetaVerifyRequest struct {
	// Phone number:
	//
	// - When paramType is normal: input the plaintext phone number.
	//
	// - When paramType is md5: input the encrypted phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// ● 明文：186****2055
	//
	// ● 密文：
	//
	// 849169cd3b20621c1c78bd61a11a4fc2
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Encryption method:
	//
	// - normal: plaintext without encryption
	//
	// - md5: MD5 encryption
	//
	// This parameter is required.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Name:
	//
	// - When paramType is normal: input the plaintext name.
	//
	// - When paramType is md5: input the encrypted name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ● 明文：张三
	//
	// ● 密文：
	//
	// 32fa7bcd874161bea8ec8fd98f390ec9
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Mobile2MetaVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s Mobile2MetaVerifyRequest) GoString() string {
	return s.String()
}

func (s *Mobile2MetaVerifyRequest) GetMobile() *string {
	return s.Mobile
}

func (s *Mobile2MetaVerifyRequest) GetParamType() *string {
	return s.ParamType
}

func (s *Mobile2MetaVerifyRequest) GetUserName() *string {
	return s.UserName
}

func (s *Mobile2MetaVerifyRequest) SetMobile(v string) *Mobile2MetaVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *Mobile2MetaVerifyRequest) SetParamType(v string) *Mobile2MetaVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Mobile2MetaVerifyRequest) SetUserName(v string) *Mobile2MetaVerifyRequest {
	s.UserName = &v
	return s
}

func (s *Mobile2MetaVerifyRequest) Validate() error {
	return dara.Validate(s)
}
