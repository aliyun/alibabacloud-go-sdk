// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoloWarehouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteHoloWarehouseRequest
	GetName() *string
}

type DeleteHoloWarehouseRequest struct {
	// The name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DeleteHoloWarehouseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *DeleteHoloWarehouseRequest) GetName() *string {
	return s.Name
}

func (s *DeleteHoloWarehouseRequest) SetName(v string) *DeleteHoloWarehouseRequest {
	s.Name = &v
	return s
}

func (s *DeleteHoloWarehouseRequest) Validate() error {
	return dara.Validate(s)
}
