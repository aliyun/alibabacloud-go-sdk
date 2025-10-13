// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendHoloWarehouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *SuspendHoloWarehouseRequest
	GetName() *string
}

type SuspendHoloWarehouseRequest struct {
	// The name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s SuspendHoloWarehouseRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *SuspendHoloWarehouseRequest) GetName() *string {
	return s.Name
}

func (s *SuspendHoloWarehouseRequest) SetName(v string) *SuspendHoloWarehouseRequest {
	s.Name = &v
	return s
}

func (s *SuspendHoloWarehouseRequest) Validate() error {
	return dara.Validate(s)
}
