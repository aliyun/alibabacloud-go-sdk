// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEntriesToAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *AddEntriesToAclResponseBody
	GetAclId() *string
	SetRequestId(v string) *AddEntriesToAclResponseBody
	GetRequestId() *string
}

type AddEntriesToAclResponseBody struct {
	// The ACL ID.
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

func (s AddEntriesToAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddEntriesToAclResponseBody) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *AddEntriesToAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddEntriesToAclResponseBody) SetAclId(v string) *AddEntriesToAclResponseBody {
	s.AclId = &v
	return s
}

func (s *AddEntriesToAclResponseBody) SetRequestId(v string) *AddEntriesToAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEntriesToAclResponseBody) Validate() error {
	return dara.Validate(s)
}
