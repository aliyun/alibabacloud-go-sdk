// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDatabaseResponseBody
	GetRequestId() *string
}

type DeleteDatabaseResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 81500666-d7f5-4143-8329-0223cc738105
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatabaseResponseBody) SetRequestId(v string) *DeleteDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
