// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeKvstorePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InitializeKvstorePermissionResponseBody
	GetRequestId() *string
}

type InitializeKvstorePermissionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitializeKvstorePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitializeKvstorePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeKvstorePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitializeKvstorePermissionResponseBody) SetRequestId(v string) *InitializeKvstorePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeKvstorePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
