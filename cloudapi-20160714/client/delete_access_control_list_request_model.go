// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessControlListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DeleteAccessControlListRequest
	GetAclId() *string
	SetSecurityToken(v string) *DeleteAccessControlListRequest
	GetSecurityToken() *string
}

type DeleteAccessControlListRequest struct {
	// The ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-uf6fpfdg3b5muska7uqem
	AclId         *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteAccessControlListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessControlListRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessControlListRequest) GetAclId() *string {
	return s.AclId
}

func (s *DeleteAccessControlListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteAccessControlListRequest) SetAclId(v string) *DeleteAccessControlListRequest {
	s.AclId = &v
	return s
}

func (s *DeleteAccessControlListRequest) SetSecurityToken(v string) *DeleteAccessControlListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteAccessControlListRequest) Validate() error {
	return dara.Validate(s)
}
