// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeCoordinatePrivilegeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeCoordinatePrivilegeResponseBody
	GetRequestId() *string
}

type RevokeCoordinatePrivilegeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeCoordinatePrivilegeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeCoordinatePrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeCoordinatePrivilegeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeCoordinatePrivilegeResponseBody) SetRequestId(v string) *RevokeCoordinatePrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeCoordinatePrivilegeResponseBody) Validate() error {
	return dara.Validate(s)
}
