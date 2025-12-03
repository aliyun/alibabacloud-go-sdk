// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessControlListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *CreateAccessControlListResponseBody
	GetAclId() *string
	SetRequestId(v string) *CreateAccessControlListResponseBody
	GetRequestId() *string
}

type CreateAccessControlListResponseBody struct {
	// The IP version. Valid values: **ipv4*	- and **ipv6**.
	//
	// example:
	//
	// acl-rj9xpxzcwxrukois****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 988CB45E-1643-48C0-87B4-928DDF77EA49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessControlListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessControlListResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessControlListResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *CreateAccessControlListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccessControlListResponseBody) SetAclId(v string) *CreateAccessControlListResponseBody {
	s.AclId = &v
	return s
}

func (s *CreateAccessControlListResponseBody) SetRequestId(v string) *CreateAccessControlListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccessControlListResponseBody) Validate() error {
	return dara.Validate(s)
}
