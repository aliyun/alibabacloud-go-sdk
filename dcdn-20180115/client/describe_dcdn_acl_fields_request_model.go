// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnAclFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDcdnAclFieldsRequest
	GetLang() *string
}

type DescribeDcdnAclFieldsRequest struct {
	// The access language. Default value: en. Valid values:
	//
	// 	- **en**: English
	//
	// 	- **zh**: Chinese
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeDcdnAclFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnAclFieldsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnAclFieldsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDcdnAclFieldsRequest) SetLang(v string) *DescribeDcdnAclFieldsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDcdnAclFieldsRequest) Validate() error {
	return dara.Validate(s)
}
