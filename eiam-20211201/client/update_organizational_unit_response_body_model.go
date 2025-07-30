// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationalUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateOrganizationalUnitResponseBody
	GetRequestId() *string
}

type UpdateOrganizationalUnitResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOrganizationalUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOrganizationalUnitResponseBody) SetRequestId(v string) *UpdateOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOrganizationalUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
