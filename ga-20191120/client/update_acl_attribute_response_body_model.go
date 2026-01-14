// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAclAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *UpdateAclAttributeResponseBody
	GetAclId() *string
	SetRequestId(v string) *UpdateAclAttributeResponseBody
	GetRequestId() *string
}

type UpdateAclAttributeResponseBody struct {
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

func (s UpdateAclAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAclAttributeResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *UpdateAclAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAclAttributeResponseBody) SetAclId(v string) *UpdateAclAttributeResponseBody {
	s.AclId = &v
	return s
}

func (s *UpdateAclAttributeResponseBody) SetRequestId(v string) *UpdateAclAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAclAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
