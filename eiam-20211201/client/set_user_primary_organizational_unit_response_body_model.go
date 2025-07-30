// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserPrimaryOrganizationalUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetUserPrimaryOrganizationalUnitResponseBody
	GetRequestId() *string
}

type SetUserPrimaryOrganizationalUnitResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetUserPrimaryOrganizationalUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetUserPrimaryOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *SetUserPrimaryOrganizationalUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetUserPrimaryOrganizationalUnitResponseBody) SetRequestId(v string) *SetUserPrimaryOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
