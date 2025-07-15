// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeApiDocRequest
	GetApiId() *string
	SetGroupId(v string) *DescribeApiDocRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DescribeApiDocRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeApiDocRequest
	GetStageName() *string
}

type DescribeApiDocRequest struct {
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3b81fd160f5645e097cc8855d75a1cf6
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 123
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The environment name. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **TEST**
	//
	// If this parameter is not specified, the default value RELEASE is used.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiDocRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiDocRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiDocRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiDocRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiDocRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApiDocRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApiDocRequest) SetApiId(v string) *DescribeApiDocRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiDocRequest) SetGroupId(v string) *DescribeApiDocRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiDocRequest) SetSecurityToken(v string) *DescribeApiDocRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiDocRequest) SetStageName(v string) *DescribeApiDocRequest {
	s.StageName = &v
	return s
}

func (s *DescribeApiDocRequest) Validate() error {
	return dara.Validate(s)
}
