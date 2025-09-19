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
	// The mobile number.
	//
	// >
	//
	// > - If **paramType*	- is set to **normal**, enter the plaintext value.
	//
	// > - If **paramType*	- is set to **md5**, enter the 32-bit lowercase MD5 string.
	//
	// This parameter is required.
	//
	// example:
	//
	// 186****1234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The parameter type:
	//
	// - **normal**: plaintext
	//
	// - **md5**: MD5-encrypted
	//
	// This parameter is required.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The product to use. Set this parameter to the static value **MOBILE_2META**.
	//
	// This parameter is required.
	//
	// example:
	//
	// MOBILE_2META
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The name.
	//
	// >
	//
	// > - If **paramType*	- is set to **normal**, enter the plaintext value.
	//
	// > - If **paramType*	- is set to **md5**, enter the 32-bit lowercase MD5 string.
	//
	// This parameter is required.
	//
	// example:
	//
	// Zhang*
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
