// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrganizationalUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationalUnitId(v string) *CreateOrganizationalUnitResponseBody
	GetOrganizationalUnitId() *string
	SetRequestId(v string) *CreateOrganizationalUnitResponseBody
	GetRequestId() *string
}

type CreateOrganizationalUnitResponseBody struct {
	// The organization ID.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrganizationalUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitResponseBody) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *CreateOrganizationalUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrganizationalUnitResponseBody) SetOrganizationalUnitId(v string) *CreateOrganizationalUnitResponseBody {
	s.OrganizationalUnitId = &v
	return s
}

func (s *CreateOrganizationalUnitResponseBody) SetRequestId(v string) *CreateOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrganizationalUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
