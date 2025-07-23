// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeEndpointAclShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclStrategy(v string) *RevokeEndpointAclShrinkRequest
	GetAclStrategy() *string
	SetCidrListShrink(v string) *RevokeEndpointAclShrinkRequest
	GetCidrListShrink() *string
	SetEndpointType(v string) *RevokeEndpointAclShrinkRequest
	GetEndpointType() *string
}

type RevokeEndpointAclShrinkRequest struct {
	// The ACL policy. Value:
	//
	// 	- **allow**: indicates that this operation is included in the Cidr whitelist. (Only the allow is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// allow
	AclStrategy *string `json:"AclStrategy,omitempty" xml:"AclStrategy,omitempty"`
	// The list of CIDR block.
	//
	// This parameter is required.
	CidrListShrink *string `json:"CidrList,omitempty" xml:"CidrList,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **public**: indicates public endpoint. (Only the public is supported.)
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s RevokeEndpointAclShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeEndpointAclShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevokeEndpointAclShrinkRequest) GetAclStrategy() *string {
	return s.AclStrategy
}

func (s *RevokeEndpointAclShrinkRequest) GetCidrListShrink() *string {
	return s.CidrListShrink
}

func (s *RevokeEndpointAclShrinkRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *RevokeEndpointAclShrinkRequest) SetAclStrategy(v string) *RevokeEndpointAclShrinkRequest {
	s.AclStrategy = &v
	return s
}

func (s *RevokeEndpointAclShrinkRequest) SetCidrListShrink(v string) *RevokeEndpointAclShrinkRequest {
	s.CidrListShrink = &v
	return s
}

func (s *RevokeEndpointAclShrinkRequest) SetEndpointType(v string) *RevokeEndpointAclShrinkRequest {
	s.EndpointType = &v
	return s
}

func (s *RevokeEndpointAclShrinkRequest) Validate() error {
	return dara.Validate(s)
}
