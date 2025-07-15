// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessControlListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclName(v string) *CreateAccessControlListRequest
	GetAclName() *string
	SetAddressIPVersion(v string) *CreateAccessControlListRequest
	GetAddressIPVersion() *string
	SetSecurityToken(v string) *CreateAccessControlListRequest
	GetSecurityToken() *string
}

type CreateAccessControlListRequest struct {
	// The name of the ACL. The name must be 1 to 30 characters in length, and can contain letters, digits, periods (.), hyphens (-), forward slashes (/), and underscores (_). The name must be unique within the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// testAcl
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// The IP protocol version of the ACL. Valid values:
	//
	// 	- **IPv4**
	//
	// 	- **IPv6**
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateAccessControlListRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessControlListRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessControlListRequest) GetAclName() *string {
	return s.AclName
}

func (s *CreateAccessControlListRequest) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *CreateAccessControlListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateAccessControlListRequest) SetAclName(v string) *CreateAccessControlListRequest {
	s.AclName = &v
	return s
}

func (s *CreateAccessControlListRequest) SetAddressIPVersion(v string) *CreateAccessControlListRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *CreateAccessControlListRequest) SetSecurityToken(v string) *CreateAccessControlListRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateAccessControlListRequest) Validate() error {
	return dara.Validate(s)
}
