// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatabaseAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDatabaseAccountResponseBody
	GetRequestId() *string
}

type DeleteDatabaseAccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 81500666-d7f5-4143-8329-0223cc738105
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatabaseAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatabaseAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatabaseAccountResponseBody) SetRequestId(v string) *DeleteDatabaseAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatabaseAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
