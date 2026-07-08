// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarehouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWarehouseInstanceId(v string) *DeleteWarehouseRequest
	GetWarehouseInstanceId() *string
}

type DeleteWarehouseRequest struct {
	// The certificate warehouse instance ID.
	//
	// example:
	//
	// cas-wh-Q7ID6V
	WarehouseInstanceId *string `json:"WarehouseInstanceId,omitempty" xml:"WarehouseInstanceId,omitempty"`
}

func (s DeleteWarehouseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarehouseRequest) GoString() string {
	return s.String()
}

func (s *DeleteWarehouseRequest) GetWarehouseInstanceId() *string {
	return s.WarehouseInstanceId
}

func (s *DeleteWarehouseRequest) SetWarehouseInstanceId(v string) *DeleteWarehouseRequest {
	s.WarehouseInstanceId = &v
	return s
}

func (s *DeleteWarehouseRequest) Validate() error {
	return dara.Validate(s)
}
