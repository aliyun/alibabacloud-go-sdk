// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationalUnitParentIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateOrganizationalUnitParentIdResponseBody
	GetRequestId() *string
}

type UpdateOrganizationalUnitParentIdResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOrganizationalUnitParentIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationalUnitParentIdResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitParentIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOrganizationalUnitParentIdResponseBody) SetRequestId(v string) *UpdateOrganizationalUnitParentIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOrganizationalUnitParentIdResponseBody) Validate() error {
	return dara.Validate(s)
}
