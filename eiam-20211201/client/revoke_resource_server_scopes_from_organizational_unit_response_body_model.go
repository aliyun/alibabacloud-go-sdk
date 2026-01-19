// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourceServerScopesFromOrganizationalUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeResourceServerScopesFromOrganizationalUnitResponseBody
	GetRequestId() *string
}

type RevokeResourceServerScopesFromOrganizationalUnitResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeResourceServerScopesFromOrganizationalUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourceServerScopesFromOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitResponseBody) SetRequestId(v string) *RevokeResourceServerScopesFromOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
