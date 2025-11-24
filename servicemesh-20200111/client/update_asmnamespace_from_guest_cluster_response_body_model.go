// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateASMNamespaceFromGuestClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateASMNamespaceFromGuestClusterResponseBody
	GetRequestId() *string
}

type UpdateASMNamespaceFromGuestClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9DEC6122-ACEC-183D-8451-8E0A1A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateASMNamespaceFromGuestClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateASMNamespaceFromGuestClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateASMNamespaceFromGuestClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateASMNamespaceFromGuestClusterResponseBody) SetRequestId(v string) *UpdateASMNamespaceFromGuestClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateASMNamespaceFromGuestClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
