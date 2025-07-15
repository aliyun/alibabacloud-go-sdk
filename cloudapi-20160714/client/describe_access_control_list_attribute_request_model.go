// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessControlListAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DescribeAccessControlListAttributeRequest
	GetAclId() *string
	SetSecurityToken(v string) *DescribeAccessControlListAttributeRequest
	GetSecurityToken() *string
}

type DescribeAccessControlListAttributeRequest struct {
	// The ID of the access control policy.
	//
	// example:
	//
	// acl-3nsohdozz0ru8fi5onwz1
	AclId         *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAccessControlListAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListAttributeRequest) GetAclId() *string {
	return s.AclId
}

func (s *DescribeAccessControlListAttributeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeAccessControlListAttributeRequest) SetAclId(v string) *DescribeAccessControlListAttributeRequest {
	s.AclId = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) SetSecurityToken(v string) *DescribeAccessControlListAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAccessControlListAttributeRequest) Validate() error {
	return dara.Validate(s)
}
