// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarehouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWarehouseResponseBody
	GetRequestId() *string
	SetWarehouseInstanceId(v string) *CreateWarehouseResponseBody
	GetWarehouseInstanceId() *string
}

type CreateWarehouseResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F0206B77-14B9-584C-8A3A-09D5827FBC50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The instance ID of the certificate application repository.
	//
	// example:
	//
	// cas-wh-typ-serial
	WarehouseInstanceId *string `json:"WarehouseInstanceId,omitempty" xml:"WarehouseInstanceId,omitempty"`
}

func (s CreateWarehouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWarehouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWarehouseResponseBody) GetWarehouseInstanceId() *string {
	return s.WarehouseInstanceId
}

func (s *CreateWarehouseResponseBody) SetRequestId(v string) *CreateWarehouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWarehouseResponseBody) SetWarehouseInstanceId(v string) *CreateWarehouseResponseBody {
	s.WarehouseInstanceId = &v
	return s
}

func (s *CreateWarehouseResponseBody) Validate() error {
	return dara.Validate(s)
}
