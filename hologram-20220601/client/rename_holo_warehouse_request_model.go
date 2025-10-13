// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameHoloWarehouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *RenameHoloWarehouseRequest
	GetName() *string
	SetNewWarehouseName(v string) *RenameHoloWarehouseRequest
	GetNewWarehouseName() *string
}

type RenameHoloWarehouseRequest struct {
	// The original name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The new name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// new_name
	NewWarehouseName *string `json:"newWarehouseName,omitempty" xml:"newWarehouseName,omitempty"`
}

func (s RenameHoloWarehouseRequest) String() string {
	return dara.Prettify(s)
}

func (s RenameHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *RenameHoloWarehouseRequest) GetName() *string {
	return s.Name
}

func (s *RenameHoloWarehouseRequest) GetNewWarehouseName() *string {
	return s.NewWarehouseName
}

func (s *RenameHoloWarehouseRequest) SetName(v string) *RenameHoloWarehouseRequest {
	s.Name = &v
	return s
}

func (s *RenameHoloWarehouseRequest) SetNewWarehouseName(v string) *RenameHoloWarehouseRequest {
	s.NewWarehouseName = &v
	return s
}

func (s *RenameHoloWarehouseRequest) Validate() error {
	return dara.Validate(s)
}
