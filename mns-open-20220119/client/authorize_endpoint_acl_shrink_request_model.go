// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeEndpointAclShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclStrategy(v string) *AuthorizeEndpointAclShrinkRequest
	GetAclStrategy() *string
	SetCidrListShrink(v string) *AuthorizeEndpointAclShrinkRequest
	GetCidrListShrink() *string
	SetEndpointType(v string) *AuthorizeEndpointAclShrinkRequest
	GetEndpointType() *string
}

type AuthorizeEndpointAclShrinkRequest struct {
	// The access control list (ACL) policy. Valid value:
	//
	// - **allow**: A CIDR whitelist. Only allow is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// allow
	AclStrategy *string `json:"AclStrategy,omitempty" xml:"AclStrategy,omitempty"`
	// A list of CIDR blocks.
	//
	// This parameter is required.
	CidrListShrink *string `json:"CidrList,omitempty" xml:"CidrList,omitempty"`
	// The type of the endpoint. Valid value:
	//
	// - **public**: An internet endpoint. Only public is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s AuthorizeEndpointAclShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeEndpointAclShrinkRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeEndpointAclShrinkRequest) GetAclStrategy() *string {
	return s.AclStrategy
}

func (s *AuthorizeEndpointAclShrinkRequest) GetCidrListShrink() *string {
	return s.CidrListShrink
}

func (s *AuthorizeEndpointAclShrinkRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *AuthorizeEndpointAclShrinkRequest) SetAclStrategy(v string) *AuthorizeEndpointAclShrinkRequest {
	s.AclStrategy = &v
	return s
}

func (s *AuthorizeEndpointAclShrinkRequest) SetCidrListShrink(v string) *AuthorizeEndpointAclShrinkRequest {
	s.CidrListShrink = &v
	return s
}

func (s *AuthorizeEndpointAclShrinkRequest) SetEndpointType(v string) *AuthorizeEndpointAclShrinkRequest {
	s.EndpointType = &v
	return s
}

func (s *AuthorizeEndpointAclShrinkRequest) Validate() error {
	return dara.Validate(s)
}
