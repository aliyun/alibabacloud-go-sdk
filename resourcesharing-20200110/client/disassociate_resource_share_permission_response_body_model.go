// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateResourceSharePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisassociateResourceSharePermissionResponseBody
	GetRequestId() *string
}

type DisassociateResourceSharePermissionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 111FB84A-60A9-403E-9067-E55D7EE95BD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateResourceSharePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisassociateResourceSharePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateResourceSharePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisassociateResourceSharePermissionResponseBody) SetRequestId(v string) *DisassociateResourceSharePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociateResourceSharePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
