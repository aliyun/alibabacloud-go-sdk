// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostgresExtensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePostgresExtensionsResponseBody
	GetRequestId() *string
}

type UpdatePostgresExtensionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7E4448A6-9FE6-4474-A0C1-AA7CFC772CAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePostgresExtensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostgresExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePostgresExtensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePostgresExtensionsResponseBody) SetRequestId(v string) *UpdatePostgresExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePostgresExtensionsResponseBody) Validate() error {
	return dara.Validate(s)
}
