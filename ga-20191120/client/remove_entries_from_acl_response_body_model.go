// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEntriesFromAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *RemoveEntriesFromAclResponseBody
	GetAclId() *string
	SetRequestId(v string) *RemoveEntriesFromAclResponseBody
	GetRequestId() *string
}

type RemoveEntriesFromAclResponseBody struct {
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
	// 64ADAB1E-0B7F-4FD8-A404-3BECC0E9CCFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveEntriesFromAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveEntriesFromAclResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *RemoveEntriesFromAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveEntriesFromAclResponseBody) SetAclId(v string) *RemoveEntriesFromAclResponseBody {
	s.AclId = &v
	return s
}

func (s *RemoveEntriesFromAclResponseBody) SetRequestId(v string) *RemoveEntriesFromAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveEntriesFromAclResponseBody) Validate() error {
	return dara.Validate(s)
}
