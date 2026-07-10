// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaDetailStandardVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifyNum(v string) *Mobile3MetaDetailStandardVerifyRequest
	GetIdentifyNum() *string
	SetMobile(v string) *Mobile3MetaDetailStandardVerifyRequest
	GetMobile() *string
	SetParamType(v string) *Mobile3MetaDetailStandardVerifyRequest
	GetParamType() *string
	SetUserName(v string) *Mobile3MetaDetailStandardVerifyRequest
	GetUserName() *string
}

type Mobile3MetaDetailStandardVerifyRequest struct {
	// The ID card number.
	//
	// - If ParamType is set to normal, pass in the ID card number in plaintext.
	//
	// - If ParamType is set to md5, pass in the MD5-encrypted ID card number.
	//
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// The phone number.
	//
	// - If ParamType is set to normal, pass in the phone number in plaintext.
	//
	// - If ParamType is set to md5, pass in the MD5-encrypted phone number.
	//
	// example:
	//
	// 138********
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
	// The name.
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

func (s Mobile3MetaDetailStandardVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaDetailStandardVerifyRequest) GoString() string {
	return s.String()
}

func (s *Mobile3MetaDetailStandardVerifyRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *Mobile3MetaDetailStandardVerifyRequest) GetMobile() *string {
	return s.Mobile
}

func (s *Mobile3MetaDetailStandardVerifyRequest) GetParamType() *string {
	return s.ParamType
}

func (s *Mobile3MetaDetailStandardVerifyRequest) GetUserName() *string {
	return s.UserName
}

func (s *Mobile3MetaDetailStandardVerifyRequest) SetIdentifyNum(v string) *Mobile3MetaDetailStandardVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Mobile3MetaDetailStandardVerifyRequest) SetMobile(v string) *Mobile3MetaDetailStandardVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *Mobile3MetaDetailStandardVerifyRequest) SetParamType(v string) *Mobile3MetaDetailStandardVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Mobile3MetaDetailStandardVerifyRequest) SetUserName(v string) *Mobile3MetaDetailStandardVerifyRequest {
	s.UserName = &v
	return s
}

func (s *Mobile3MetaDetailStandardVerifyRequest) Validate() error {
	return dara.Validate(s)
}
