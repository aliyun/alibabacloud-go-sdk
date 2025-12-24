// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnumItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnumType(v string) *DescribeEnumItemsRequest
	GetEnumType() *string
	SetLang(v string) *DescribeEnumItemsRequest
	GetLang() *string
}

type DescribeEnumItemsRequest struct {
	// The type of the enumeration item. Valid values:
	//
	// 	- **process**: scenarios
	//
	// This parameter is required.
	//
	// example:
	//
	// process
	EnumType *string `json:"EnumType,omitempty" xml:"EnumType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese (default)
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeEnumItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnumItemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnumItemsRequest) GetEnumType() *string {
	return s.EnumType
}

func (s *DescribeEnumItemsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEnumItemsRequest) SetEnumType(v string) *DescribeEnumItemsRequest {
	s.EnumType = &v
	return s
}

func (s *DescribeEnumItemsRequest) SetLang(v string) *DescribeEnumItemsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEnumItemsRequest) Validate() error {
	return dara.Validate(s)
}
