// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateACLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *CreateACLResponseBody
	GetAclId() *string
	SetAclType(v string) *CreateACLResponseBody
	GetAclType() *string
	SetRequestId(v string) *CreateACLResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateACLResponseBody
	GetResourceGroupId() *string
}

type CreateACLResponseBody struct {
	// The ID of the ACL.
	//
	// example:
	//
	// acl-o6yol7zowii5n2****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The type of SAG instance to be associated with the ACL.
	//
	// example:
	//
	// acl-hardware
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EE837E9F-BD50-4C2B-9E47-260F9D848480
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the ACL belongs.
	//
	// example:
	//
	// rg-acfm2iu4fnc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateACLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateACLResponseBody) GoString() string {
	return s.String()
}

func (s *CreateACLResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *CreateACLResponseBody) GetAclType() *string {
	return s.AclType
}

func (s *CreateACLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateACLResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateACLResponseBody) SetAclId(v string) *CreateACLResponseBody {
	s.AclId = &v
	return s
}

func (s *CreateACLResponseBody) SetAclType(v string) *CreateACLResponseBody {
	s.AclType = &v
	return s
}

func (s *CreateACLResponseBody) SetRequestId(v string) *CreateACLResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateACLResponseBody) SetResourceGroupId(v string) *CreateACLResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateACLResponseBody) Validate() error {
	return dara.Validate(s)
}
