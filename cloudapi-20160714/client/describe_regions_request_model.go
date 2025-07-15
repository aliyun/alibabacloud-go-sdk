// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *DescribeRegionsRequest
	GetLanguage() *string
	SetSecurityToken(v string) *DescribeRegionsRequest
	GetSecurityToken() *string
}

type DescribeRegionsRequest struct {
	// The language in which you want to return the descriptions of the access control policies. Valid values:
	//
	// 	- zh-CN (default value): Chinese
	//
	// 	- en: English
	//
	// 	- ja: Japanese
	//
	// example:
	//
	// zh
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeRegionsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeRegionsRequest) SetLanguage(v string) *DescribeRegionsRequest {
	s.Language = &v
	return s
}

func (s *DescribeRegionsRequest) SetSecurityToken(v string) *DescribeRegionsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
