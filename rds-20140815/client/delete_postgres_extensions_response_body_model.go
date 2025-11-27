// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePostgresExtensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePostgresExtensionsResponseBody
	GetRequestId() *string
}

type DeletePostgresExtensionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7E4448A6-9FE6-4474-A0C1-AA7CFC772CAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePostgresExtensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePostgresExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePostgresExtensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePostgresExtensionsResponseBody) SetRequestId(v string) *DeletePostgresExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePostgresExtensionsResponseBody) Validate() error {
	return dara.Validate(s)
}
