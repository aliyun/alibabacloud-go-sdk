// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSecFuncInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDcdnSecFuncInfoRequest
	GetLang() *string
	SetSecFuncType(v string) *DescribeDcdnSecFuncInfoRequest
	GetSecFuncType() *string
}

type DescribeDcdnSecFuncInfoRequest struct {
	// The language. Valid values: en and zh. Default value: en.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the drop-down list. Valid values: RobotRuleName and RobotObject.
	//
	// This parameter is required.
	//
	// example:
	//
	// RobotRuleName
	SecFuncType *string `json:"SecFuncType,omitempty" xml:"SecFuncType,omitempty"`
}

func (s DescribeDcdnSecFuncInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSecFuncInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecFuncInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDcdnSecFuncInfoRequest) GetSecFuncType() *string {
	return s.SecFuncType
}

func (s *DescribeDcdnSecFuncInfoRequest) SetLang(v string) *DescribeDcdnSecFuncInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoRequest) SetSecFuncType(v string) *DescribeDcdnSecFuncInfoRequest {
	s.SecFuncType = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoRequest) Validate() error {
	return dara.Validate(s)
}
