// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOrganizationalUnitChildrenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOrganizationalUnitChildrenResponseBody
	GetRequestId() *string
}

type DeleteOrganizationalUnitChildrenResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOrganizationalUnitChildrenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOrganizationalUnitChildrenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitChildrenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOrganizationalUnitChildrenResponseBody) SetRequestId(v string) *DeleteOrganizationalUnitChildrenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOrganizationalUnitChildrenResponseBody) Validate() error {
	return dara.Validate(s)
}
