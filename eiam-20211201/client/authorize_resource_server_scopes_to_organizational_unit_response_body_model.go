// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceServerScopesToOrganizationalUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuthorizeResourceServerScopesToOrganizationalUnitResponseBody
	GetRequestId() *string
}

type AuthorizeResourceServerScopesToOrganizationalUnitResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeResourceServerScopesToOrganizationalUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceServerScopesToOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitResponseBody) SetRequestId(v string) *AuthorizeResourceServerScopesToOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
