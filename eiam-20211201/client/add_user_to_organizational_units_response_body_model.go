// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToOrganizationalUnitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUserToOrganizationalUnitsResponseBody
	GetRequestId() *string
}

type AddUserToOrganizationalUnitsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserToOrganizationalUnitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserToOrganizationalUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToOrganizationalUnitsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserToOrganizationalUnitsResponseBody) SetRequestId(v string) *AddUserToOrganizationalUnitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserToOrganizationalUnitsResponseBody) Validate() error {
	return dara.Validate(s)
}
