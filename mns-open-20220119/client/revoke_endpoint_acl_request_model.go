// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeEndpointAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclStrategy(v string) *RevokeEndpointAclRequest
	GetAclStrategy() *string
	SetCidrList(v []*string) *RevokeEndpointAclRequest
	GetCidrList() []*string
	SetEndpointType(v string) *RevokeEndpointAclRequest
	GetEndpointType() *string
}

type RevokeEndpointAclRequest struct {
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
	CidrList []*string `json:"CidrList,omitempty" xml:"CidrList,omitempty" type:"Repeated"`
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

func (s RevokeEndpointAclRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeEndpointAclRequest) GoString() string {
	return s.String()
}

func (s *RevokeEndpointAclRequest) GetAclStrategy() *string {
	return s.AclStrategy
}

func (s *RevokeEndpointAclRequest) GetCidrList() []*string {
	return s.CidrList
}

func (s *RevokeEndpointAclRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *RevokeEndpointAclRequest) SetAclStrategy(v string) *RevokeEndpointAclRequest {
	s.AclStrategy = &v
	return s
}

func (s *RevokeEndpointAclRequest) SetCidrList(v []*string) *RevokeEndpointAclRequest {
	s.CidrList = v
	return s
}

func (s *RevokeEndpointAclRequest) SetEndpointType(v string) *RevokeEndpointAclRequest {
	s.EndpointType = &v
	return s
}

func (s *RevokeEndpointAclRequest) Validate() error {
	return dara.Validate(s)
}
