// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaVerifyIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifyNum(v string) *Id2MetaVerifyIntlRequest
	GetIdentifyNum() *string
	SetParamType(v string) *Id2MetaVerifyIntlRequest
	GetParamType() *string
	SetProductCode(v string) *Id2MetaVerifyIntlRequest
	GetProductCode() *string
	SetUserName(v string) *Id2MetaVerifyIntlRequest
	GetUserName() *string
}

type Id2MetaVerifyIntlRequest struct {
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// example:
	//
	// ID_2META
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Id2MetaVerifyIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyIntlRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *Id2MetaVerifyIntlRequest) GetParamType() *string {
	return s.ParamType
}

func (s *Id2MetaVerifyIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *Id2MetaVerifyIntlRequest) GetUserName() *string {
	return s.UserName
}

func (s *Id2MetaVerifyIntlRequest) SetIdentifyNum(v string) *Id2MetaVerifyIntlRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Id2MetaVerifyIntlRequest) SetParamType(v string) *Id2MetaVerifyIntlRequest {
	s.ParamType = &v
	return s
}

func (s *Id2MetaVerifyIntlRequest) SetProductCode(v string) *Id2MetaVerifyIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *Id2MetaVerifyIntlRequest) SetUserName(v string) *Id2MetaVerifyIntlRequest {
	s.UserName = &v
	return s
}

func (s *Id2MetaVerifyIntlRequest) Validate() error {
	return dara.Validate(s)
}
