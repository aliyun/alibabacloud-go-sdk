// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateResourceSharePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateResourceSharePermissionResponseBody
	GetRequestId() *string
}

type AssociateResourceSharePermissionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 111FB84A-60A9-403E-9067-E55D7EE95BD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateResourceSharePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateResourceSharePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateResourceSharePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateResourceSharePermissionResponseBody) SetRequestId(v string) *AssociateResourceSharePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateResourceSharePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
