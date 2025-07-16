// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSecFuncInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeCdnSecFuncInfoRequest
	GetLang() *string
	SetSecFuncType(v string) *DescribeCdnSecFuncInfoRequest
	GetSecFuncType() *string
}

type DescribeCdnSecFuncInfoRequest struct {
	// The language.
	//
	// 	- en: English
	//
	// 	- zh: Chinese
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the security feature. Valid values:
	//
	// 	- CipherSuiteGroupCustomize: custom cipher suite.
	//
	// 	- CipherSuiteGroupStrict: dustom cipher suite.
	//
	// This parameter is required.
	//
	// example:
	//
	// CipherSuiteGroupCustomize
	SecFuncType *string `json:"SecFuncType,omitempty" xml:"SecFuncType,omitempty"`
}

func (s DescribeCdnSecFuncInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSecFuncInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnSecFuncInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCdnSecFuncInfoRequest) GetSecFuncType() *string {
	return s.SecFuncType
}

func (s *DescribeCdnSecFuncInfoRequest) SetLang(v string) *DescribeCdnSecFuncInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCdnSecFuncInfoRequest) SetSecFuncType(v string) *DescribeCdnSecFuncInfoRequest {
	s.SecFuncType = &v
	return s
}

func (s *DescribeCdnSecFuncInfoRequest) Validate() error {
	return dara.Validate(s)
}
