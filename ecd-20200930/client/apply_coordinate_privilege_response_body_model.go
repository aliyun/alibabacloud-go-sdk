// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCoordinatePrivilegeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyCoordinatePrivilegeResponseBody
	GetRequestId() *string
}

type ApplyCoordinatePrivilegeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C43EEAC3-84F8-5C1E-A067-4751C3D1422E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyCoordinatePrivilegeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinatePrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyCoordinatePrivilegeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyCoordinatePrivilegeResponseBody) SetRequestId(v string) *ApplyCoordinatePrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyCoordinatePrivilegeResponseBody) Validate() error {
	return dara.Validate(s)
}
