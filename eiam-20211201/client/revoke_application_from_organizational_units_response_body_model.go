// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeApplicationFromOrganizationalUnitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeApplicationFromOrganizationalUnitsResponseBody
	GetRequestId() *string
}

type RevokeApplicationFromOrganizationalUnitsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeApplicationFromOrganizationalUnitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeApplicationFromOrganizationalUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromOrganizationalUnitsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeApplicationFromOrganizationalUnitsResponseBody) SetRequestId(v string) *RevokeApplicationFromOrganizationalUnitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeApplicationFromOrganizationalUnitsResponseBody) Validate() error {
	return dara.Validate(s)
}
