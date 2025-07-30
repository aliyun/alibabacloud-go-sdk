// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOrganizationalUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOrganizationalUnitResponseBody
	GetRequestId() *string
}

type DeleteOrganizationalUnitResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOrganizationalUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOrganizationalUnitResponseBody) SetRequestId(v string) *DeleteOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOrganizationalUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
