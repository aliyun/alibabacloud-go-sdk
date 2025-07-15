// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIPv6TranslatorAclListEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclEntryId(v string) *AddIPv6TranslatorAclListEntryResponseBody
	GetAclEntryId() *string
	SetRequestId(v string) *AddIPv6TranslatorAclListEntryResponseBody
	GetRequestId() *string
}

type AddIPv6TranslatorAclListEntryResponseBody struct {
	// The ID of the ACL entry.
	//
	// example:
	//
	// ipv6transaclentry-bp105jrs****
	AclEntryId *string `json:"AclEntryId,omitempty" xml:"AclEntryId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8B2F5262-6B57-43F2-defr345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIPv6TranslatorAclListEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddIPv6TranslatorAclListEntryResponseBody) GoString() string {
	return s.String()
}

func (s *AddIPv6TranslatorAclListEntryResponseBody) GetAclEntryId() *string {
	return s.AclEntryId
}

func (s *AddIPv6TranslatorAclListEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddIPv6TranslatorAclListEntryResponseBody) SetAclEntryId(v string) *AddIPv6TranslatorAclListEntryResponseBody {
	s.AclEntryId = &v
	return s
}

func (s *AddIPv6TranslatorAclListEntryResponseBody) SetRequestId(v string) *AddIPv6TranslatorAclListEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddIPv6TranslatorAclListEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
