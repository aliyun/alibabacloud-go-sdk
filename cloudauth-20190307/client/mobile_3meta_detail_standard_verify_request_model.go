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
	// ID number:
	//
	// - When `paramType` is `normal`: Input the plain text of the ID number.
	//
	// - When `paramType` is `md5`: Input the encrypted text of the ID number.
	//
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// Phone number:
	//
	// - When `paramType` is `normal`: Input the plain text of the phone number.
	//
	// - When `paramType` is `md5`: Input the encrypted text of the phone number.
	//
	// example:
	//
	// 138********
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Parameter type:
	//
	// - normal: Unencrypted.
	//
	// - md5: Encrypted with MD5.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Name:
	//
	// - When `paramType` is `normal`: Input the plain text of the name.
	//
	// - When `paramType` is `md5`: Input the encrypted text of the name.
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
