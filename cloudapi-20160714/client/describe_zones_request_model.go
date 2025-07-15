// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *DescribeZonesRequest
	GetLanguage() *string
	SetSecurityToken(v string) *DescribeZonesRequest
	GetSecurityToken() *string
}

type DescribeZonesRequest struct {
	// The language in which you want to use to return the description of the system policy. Valid values:
	//
	// 	- en: English
	//
	// 	- zh-CN: Chinese
	//
	// 	- ja: Japanese
	//
	// example:
	//
	// zh
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeZonesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeZonesRequest) SetLanguage(v string) *DescribeZonesRequest {
	s.Language = &v
	return s
}

func (s *DescribeZonesRequest) SetSecurityToken(v string) *DescribeZonesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeZonesRequest) Validate() error {
	return dara.Validate(s)
}
