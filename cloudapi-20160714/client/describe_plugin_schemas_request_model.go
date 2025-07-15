// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *DescribePluginSchemasRequest
	GetLanguage() *string
	SetSecurityToken(v string) *DescribePluginSchemasRequest
	GetSecurityToken() *string
}

type DescribePluginSchemasRequest struct {
	// example:
	//
	// en
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribePluginSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginSchemasRequest) GoString() string {
	return s.String()
}

func (s *DescribePluginSchemasRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribePluginSchemasRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribePluginSchemasRequest) SetLanguage(v string) *DescribePluginSchemasRequest {
	s.Language = &v
	return s
}

func (s *DescribePluginSchemasRequest) SetSecurityToken(v string) *DescribePluginSchemasRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePluginSchemasRequest) Validate() error {
	return dara.Validate(s)
}
