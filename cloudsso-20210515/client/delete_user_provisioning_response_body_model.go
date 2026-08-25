// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserProvisioningResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserProvisioningResponseBody
	GetRequestId() *string
}

type DeleteUserProvisioningResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F6F90F3D-4502-5877-B80B-97476F6AE2CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserProvisioningResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserProvisioningResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserProvisioningResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserProvisioningResponseBody) SetRequestId(v string) *DeleteUserProvisioningResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserProvisioningResponseBody) Validate() error {
	return dara.Validate(s)
}
