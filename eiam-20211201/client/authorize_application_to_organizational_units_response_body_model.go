// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeApplicationToOrganizationalUnitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuthorizeApplicationToOrganizationalUnitsResponseBody
	GetRequestId() *string
}

type AuthorizeApplicationToOrganizationalUnitsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeApplicationToOrganizationalUnitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeApplicationToOrganizationalUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToOrganizationalUnitsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeApplicationToOrganizationalUnitsResponseBody) SetRequestId(v string) *AuthorizeApplicationToOrganizationalUnitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeApplicationToOrganizationalUnitsResponseBody) Validate() error {
	return dara.Validate(s)
}
