// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiMarketAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeApiMarketAttributesRequest
	GetApiId() *string
	SetGroupId(v string) *DescribeApiMarketAttributesRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DescribeApiMarketAttributesRequest
	GetSecurityToken() *string
}

type DescribeApiMarketAttributesRequest struct {
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1f9b5e9ba80943099cac52e040b7e160
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// b693252f3f19445e9a01dac177c4454c
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApiMarketAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiMarketAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiMarketAttributesRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiMarketAttributesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiMarketAttributesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApiMarketAttributesRequest) SetApiId(v string) *DescribeApiMarketAttributesRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiMarketAttributesRequest) SetGroupId(v string) *DescribeApiMarketAttributesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiMarketAttributesRequest) SetSecurityToken(v string) *DescribeApiMarketAttributesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiMarketAttributesRequest) Validate() error {
	return dara.Validate(s)
}
