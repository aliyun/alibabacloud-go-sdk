// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDocTypesRequest
	GetLang() *string
}

type DescribeDocTypesRequest struct {
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeDocTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDocTypesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDocTypesRequest) SetLang(v string) *DescribeDocTypesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDocTypesRequest) Validate() error {
	return dara.Validate(s)
}
