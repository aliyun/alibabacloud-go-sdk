// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiGroupVpcWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribeApiGroupVpcWhitelistRequest
	GetGroupId() *string
	SetSecurityToken(v string) *DescribeApiGroupVpcWhitelistRequest
	GetSecurityToken() *string
}

type DescribeApiGroupVpcWhitelistRequest struct {
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// f6c64214dc27400196eef954575d60d3
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApiGroupVpcWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupVpcWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupVpcWhitelistRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiGroupVpcWhitelistRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApiGroupVpcWhitelistRequest) SetGroupId(v string) *DescribeApiGroupVpcWhitelistRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiGroupVpcWhitelistRequest) SetSecurityToken(v string) *DescribeApiGroupVpcWhitelistRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiGroupVpcWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
