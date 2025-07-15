// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaVerifyIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifyNum(v string) *Mobile3MetaVerifyIntlRequest
	GetIdentifyNum() *string
	SetMobile(v string) *Mobile3MetaVerifyIntlRequest
	GetMobile() *string
	SetParamType(v string) *Mobile3MetaVerifyIntlRequest
	GetParamType() *string
	SetProductCode(v string) *Mobile3MetaVerifyIntlRequest
	GetProductCode() *string
	SetUserName(v string) *Mobile3MetaVerifyIntlRequest
	GetUserName() *string
}

type Mobile3MetaVerifyIntlRequest struct {
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// example:
	//
	// 186****1234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// example:
	//
	// MOBILE_3META
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Mobile3MetaVerifyIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *Mobile3MetaVerifyIntlRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *Mobile3MetaVerifyIntlRequest) GetMobile() *string {
	return s.Mobile
}

func (s *Mobile3MetaVerifyIntlRequest) GetParamType() *string {
	return s.ParamType
}

func (s *Mobile3MetaVerifyIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *Mobile3MetaVerifyIntlRequest) GetUserName() *string {
	return s.UserName
}

func (s *Mobile3MetaVerifyIntlRequest) SetIdentifyNum(v string) *Mobile3MetaVerifyIntlRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Mobile3MetaVerifyIntlRequest) SetMobile(v string) *Mobile3MetaVerifyIntlRequest {
	s.Mobile = &v
	return s
}

func (s *Mobile3MetaVerifyIntlRequest) SetParamType(v string) *Mobile3MetaVerifyIntlRequest {
	s.ParamType = &v
	return s
}

func (s *Mobile3MetaVerifyIntlRequest) SetProductCode(v string) *Mobile3MetaVerifyIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *Mobile3MetaVerifyIntlRequest) SetUserName(v string) *Mobile3MetaVerifyIntlRequest {
	s.UserName = &v
	return s
}

func (s *Mobile3MetaVerifyIntlRequest) Validate() error {
	return dara.Validate(s)
}
