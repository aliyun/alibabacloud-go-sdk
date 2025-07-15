// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeApiRequest
	GetApiId() *string
	SetGroupId(v string) *DescribeApiRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DescribeApiRequest
	GetSecurityToken() *string
}

type DescribeApiRequest struct {
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8afff6c8c4c6447abb035812e4d66b65
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 123
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApiRequest) SetApiId(v string) *DescribeApiRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiRequest) SetGroupId(v string) *DescribeApiRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiRequest) SetSecurityToken(v string) *DescribeApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiRequest) Validate() error {
	return dara.Validate(s)
}
