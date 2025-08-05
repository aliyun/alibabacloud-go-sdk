// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuid(v string) *AddControlPolicyResponseBody
	GetAclUuid() *string
	SetRequestId(v string) *AddControlPolicyResponseBody
	GetRequestId() *string
}

type AddControlPolicyResponseBody struct {
	// The ID of the access control policy that is created on the Internet firewall.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df221ad84c
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AddControlPolicyResponseBody) GetAclUuid() *string {
	return s.AclUuid
}

func (s *AddControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddControlPolicyResponseBody) SetAclUuid(v string) *AddControlPolicyResponseBody {
	s.AclUuid = &v
	return s
}

func (s *AddControlPolicyResponseBody) SetRequestId(v string) *AddControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
