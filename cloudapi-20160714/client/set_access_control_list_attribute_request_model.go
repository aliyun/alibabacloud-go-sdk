// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccessControlListAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *SetAccessControlListAttributeRequest
	GetAclId() *string
	SetAclName(v string) *SetAccessControlListAttributeRequest
	GetAclName() *string
	SetSecurityToken(v string) *SetAccessControlListAttributeRequest
	GetSecurityToken() *string
}

type SetAccessControlListAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// acl-bp1ohqkonqybecf4llbrc
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testAcl
	AclName       *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetAccessControlListAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAccessControlListAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetAccessControlListAttributeRequest) GetAclId() *string {
	return s.AclId
}

func (s *SetAccessControlListAttributeRequest) GetAclName() *string {
	return s.AclName
}

func (s *SetAccessControlListAttributeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetAccessControlListAttributeRequest) SetAclId(v string) *SetAccessControlListAttributeRequest {
	s.AclId = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) SetAclName(v string) *SetAccessControlListAttributeRequest {
	s.AclName = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) SetSecurityToken(v string) *SetAccessControlListAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetAccessControlListAttributeRequest) Validate() error {
	return dara.Validate(s)
}
