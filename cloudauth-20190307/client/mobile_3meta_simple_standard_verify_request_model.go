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
	// ID number:
	//
	// - When `paramType` is `normal`: Input the plain text of the ID number.
	//
	// - When `paramType` is `md5`: Input the encrypted ID number.
	//
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// Phone number:
	//
	// - When `paramType` is `normal`: Input the plain text of the phone number.
	//
	// - When `paramType` is `md5`: Input the encrypted phone number.
	//
	// example:
	//
	// 130********
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Parameter type:
	//
	// - normal: Unencrypted.
	//
	// - md5: MD5 encrypted.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Name:
	//
	// - When `paramType` is `normal`: Input the plain text of the name.
	//
	// - When `paramType` is `md5`: Input the encrypted name.
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
