// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarehouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWarehouseResponseBody
	GetRequestId() *string
}

type DeleteWarehouseResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5C53F7EC-7C47-5DA5-8B6D-FE6B1F934E82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWarehouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWarehouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWarehouseResponseBody) SetRequestId(v string) *DeleteWarehouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWarehouseResponseBody) Validate() error {
	return dara.Validate(s)
}
