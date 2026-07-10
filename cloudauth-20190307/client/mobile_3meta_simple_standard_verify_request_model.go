// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaSimpleStandardVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifyNum(v string) *Mobile3MetaSimpleStandardVerifyRequest
	GetIdentifyNum() *string
	SetMobile(v string) *Mobile3MetaSimpleStandardVerifyRequest
	GetMobile() *string
	SetParamType(v string) *Mobile3MetaSimpleStandardVerifyRequest
	GetParamType() *string
	SetUserName(v string) *Mobile3MetaSimpleStandardVerifyRequest
	GetUserName() *string
}

type Mobile3MetaSimpleStandardVerifyRequest struct {
	// The ID card number. Valid values:
	//
	// - If ParamType is set to normal, pass in the ID card number in plaintext.
	//
	// - If ParamType is set to md5, pass in the MD5-encrypted ID card number.
	//
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// The phone number. Valid values:
	//
	// - If ParamType is set to normal, pass in the phone number in plaintext.
	//
	// - If ParamType is set to md5, pass in the MD5-encrypted phone number.
	//
	// example:
	//
	// 130********
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
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
	// The name. Valid values:
	//
	// - If ParamType is set to normal, pass in the name in plaintext.
	//
	// - If ParamType is set to md5, pass in the MD5-encrypted name.
	//
	// example:
	//
	// 张*
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Mobile3MetaSimpleStandardVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaSimpleStandardVerifyRequest) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleStandardVerifyRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *Mobile3MetaSimpleStandardVerifyRequest) GetMobile() *string {
	return s.Mobile
}

func (s *Mobile3MetaSimpleStandardVerifyRequest) GetParamType() *string {
	return s.ParamType
}

func (s *Mobile3MetaSimpleStandardVerifyRequest) GetUserName() *string {
	return s.UserName
}

func (s *Mobile3MetaSimpleStandardVerifyRequest) SetIdentifyNum(v string) *Mobile3MetaSimpleStandardVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Mobile3MetaSimpleStandardVerifyRequest) SetMobile(v string) *Mobile3MetaSimpleStandardVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *Mobile3MetaSimpleStandardVerifyRequest) SetParamType(v string) *Mobile3MetaSimpleStandardVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Mobile3MetaSimpleStandardVerifyRequest) SetUserName(v string) *Mobile3MetaSimpleStandardVerifyRequest {
	s.UserName = &v
	return s
}

func (s *Mobile3MetaSimpleStandardVerifyRequest) Validate() error {
	return dara.Validate(s)
}
