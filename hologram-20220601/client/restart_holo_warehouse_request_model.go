// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartHoloWarehouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *RestartHoloWarehouseRequest
	GetName() *string
}

type RestartHoloWarehouseRequest struct {
	// The name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RestartHoloWarehouseRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *RestartHoloWarehouseRequest) GetName() *string {
	return s.Name
}

func (s *RestartHoloWarehouseRequest) SetName(v string) *RestartHoloWarehouseRequest {
	s.Name = &v
	return s
}

func (s *RestartHoloWarehouseRequest) Validate() error {
	return dara.Validate(s)
}
