// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationalUnitDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateOrganizationalUnitDescriptionResponseBody
	GetRequestId() *string
}

type UpdateOrganizationalUnitDescriptionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOrganizationalUnitDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationalUnitDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOrganizationalUnitDescriptionResponseBody) SetRequestId(v string) *UpdateOrganizationalUnitDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOrganizationalUnitDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
