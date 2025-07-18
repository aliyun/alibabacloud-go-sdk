// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile2MetaVerifyIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobile(v string) *Mobile2MetaVerifyIntlRequest
	GetMobile() *string
	SetParamType(v string) *Mobile2MetaVerifyIntlRequest
	GetParamType() *string
	SetProductCode(v string) *Mobile2MetaVerifyIntlRequest
	GetProductCode() *string
	SetUserName(v string) *Mobile2MetaVerifyIntlRequest
	GetUserName() *string
}

type Mobile2MetaVerifyIntlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 186****1234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MOBILE_2META
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// This parameter is required.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Mobile2MetaVerifyIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s Mobile2MetaVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *Mobile2MetaVerifyIntlRequest) GetMobile() *string {
	return s.Mobile
}

func (s *Mobile2MetaVerifyIntlRequest) GetParamType() *string {
	return s.ParamType
}

func (s *Mobile2MetaVerifyIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *Mobile2MetaVerifyIntlRequest) GetUserName() *string {
	return s.UserName
}

func (s *Mobile2MetaVerifyIntlRequest) SetMobile(v string) *Mobile2MetaVerifyIntlRequest {
	s.Mobile = &v
	return s
}

func (s *Mobile2MetaVerifyIntlRequest) SetParamType(v string) *Mobile2MetaVerifyIntlRequest {
	s.ParamType = &v
	return s
}

func (s *Mobile2MetaVerifyIntlRequest) SetProductCode(v string) *Mobile2MetaVerifyIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *Mobile2MetaVerifyIntlRequest) SetUserName(v string) *Mobile2MetaVerifyIntlRequest {
	s.UserName = &v
	return s
}

func (s *Mobile2MetaVerifyIntlRequest) Validate() error {
	return dara.Validate(s)
}
