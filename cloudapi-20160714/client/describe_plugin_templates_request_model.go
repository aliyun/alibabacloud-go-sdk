// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *DescribePluginTemplatesRequest
	GetLanguage() *string
	SetPluginName(v string) *DescribePluginTemplatesRequest
	GetPluginName() *string
	SetSecurityToken(v string) *DescribePluginTemplatesRequest
	GetSecurityToken() *string
}

type DescribePluginTemplatesRequest struct {
	// The language that is used to return the description of the system policy. Valid values:
	//
	// 	- en: English
	//
	// 	- zh-CN: Chinese.
	//
	// 	- ja: Japanese
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the plug-in.
	//
	// example:
	//
	// jwtAuth
	PluginName    *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribePluginTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribePluginTemplatesRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribePluginTemplatesRequest) GetPluginName() *string {
	return s.PluginName
}

func (s *DescribePluginTemplatesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribePluginTemplatesRequest) SetLanguage(v string) *DescribePluginTemplatesRequest {
	s.Language = &v
	return s
}

func (s *DescribePluginTemplatesRequest) SetPluginName(v string) *DescribePluginTemplatesRequest {
	s.PluginName = &v
	return s
}

func (s *DescribePluginTemplatesRequest) SetSecurityToken(v string) *DescribePluginTemplatesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePluginTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
