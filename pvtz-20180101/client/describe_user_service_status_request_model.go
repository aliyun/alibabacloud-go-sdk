// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeUserServiceStatusRequest
	GetLang() *string
}

type DescribeUserServiceStatusRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeUserServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserServiceStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeUserServiceStatusRequest) SetLang(v string) *DescribeUserServiceStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
