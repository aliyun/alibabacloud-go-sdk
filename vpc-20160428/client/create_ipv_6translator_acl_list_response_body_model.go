// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIPv6TranslatorAclListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *CreateIPv6TranslatorAclListResponseBody
	GetAclId() *string
	SetRequestId(v string) *CreateIPv6TranslatorAclListResponseBody
	GetRequestId() *string
}

type CreateIPv6TranslatorAclListResponseBody struct {
	// The ACL ID.
	//
	// example:
	//
	// ipv6transacl-bp1de2xxxx
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8B2F5262-6B57-43F2-xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIPv6TranslatorAclListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIPv6TranslatorAclListResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIPv6TranslatorAclListResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *CreateIPv6TranslatorAclListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIPv6TranslatorAclListResponseBody) SetAclId(v string) *CreateIPv6TranslatorAclListResponseBody {
	s.AclId = &v
	return s
}

func (s *CreateIPv6TranslatorAclListResponseBody) SetRequestId(v string) *CreateIPv6TranslatorAclListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIPv6TranslatorAclListResponseBody) Validate() error {
	return dara.Validate(s)
}
