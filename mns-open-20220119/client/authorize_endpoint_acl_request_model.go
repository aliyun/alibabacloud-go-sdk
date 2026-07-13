// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeEndpointAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclStrategy(v string) *AuthorizeEndpointAclRequest
	GetAclStrategy() *string
	SetCidrList(v []*string) *AuthorizeEndpointAclRequest
	GetCidrList() []*string
	SetEndpointType(v string) *AuthorizeEndpointAclRequest
	GetEndpointType() *string
}

type AuthorizeEndpointAclRequest struct {
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
	CidrList []*string `json:"CidrList,omitempty" xml:"CidrList,omitempty" type:"Repeated"`
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

func (s AuthorizeEndpointAclRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeEndpointAclRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeEndpointAclRequest) GetAclStrategy() *string {
	return s.AclStrategy
}

func (s *AuthorizeEndpointAclRequest) GetCidrList() []*string {
	return s.CidrList
}

func (s *AuthorizeEndpointAclRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *AuthorizeEndpointAclRequest) SetAclStrategy(v string) *AuthorizeEndpointAclRequest {
	s.AclStrategy = &v
	return s
}

func (s *AuthorizeEndpointAclRequest) SetCidrList(v []*string) *AuthorizeEndpointAclRequest {
	s.CidrList = v
	return s
}

func (s *AuthorizeEndpointAclRequest) SetEndpointType(v string) *AuthorizeEndpointAclRequest {
	s.EndpointType = &v
	return s
}

func (s *AuthorizeEndpointAclRequest) Validate() error {
	return dara.Validate(s)
}
