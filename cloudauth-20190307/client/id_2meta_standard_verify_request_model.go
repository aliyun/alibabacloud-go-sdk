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
	// ID number:
	//
	// - When `paramType` is normal: enter the plain text of the ID number.
	//
	// - When `paramType` is md5:
	//
	// The first 6 digits (plain text) + date of birth (encrypted) + last 4 digits (plain text).
	//
	// example:
	//
	// 4****************1
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// Parameter type:
	//
	// - normal: unencrypted.
	//
	// - md5: md5 encrypted.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Name:
	//
	// - When `paramType` is normal: enter the plain text of the name.
	//
	// - When `paramType` is md5: the first character of the name (encrypted) + the rest of the name (plain text).
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
