// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DeleteAclResponseBody
	GetAclId() *string
	SetRequestId(v string) *DeleteAclResponseBody
	GetRequestId() *string
}

type DeleteAclResponseBody struct {
	// The ID of the ACL.
	//
	// example:
	//
	// nacl-hp34s2h0xx1ht4nwo****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAclResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAclResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *DeleteAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAclResponseBody) SetAclId(v string) *DeleteAclResponseBody {
	s.AclId = &v
	return s
}

func (s *DeleteAclResponseBody) SetRequestId(v string) *DeleteAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAclResponseBody) Validate() error {
	return dara.Validate(s)
}
