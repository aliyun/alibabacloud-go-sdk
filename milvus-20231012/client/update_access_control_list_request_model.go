// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccessControlListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *UpdateAccessControlListRequest
	GetAclId() *string
	SetCidr(v string) *UpdateAccessControlListRequest
	GetCidr() *string
	SetInstanceId(v string) *UpdateAccessControlListRequest
	GetInstanceId() *string
}

type UpdateAccessControlListRequest struct {
	// The ID of public network access control
	//
	// example:
	//
	// acl-123xxx
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The CIDR blocks.
	//
	// example:
	//
	// 192.168.1.0/24,172.16.0.0/16
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateAccessControlListRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessControlListRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccessControlListRequest) GetAclId() *string {
	return s.AclId
}

func (s *UpdateAccessControlListRequest) GetCidr() *string {
	return s.Cidr
}

func (s *UpdateAccessControlListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAccessControlListRequest) SetAclId(v string) *UpdateAccessControlListRequest {
	s.AclId = &v
	return s
}

func (s *UpdateAccessControlListRequest) SetCidr(v string) *UpdateAccessControlListRequest {
	s.Cidr = &v
	return s
}

func (s *UpdateAccessControlListRequest) SetInstanceId(v string) *UpdateAccessControlListRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAccessControlListRequest) Validate() error {
	return dara.Validate(s)
}
