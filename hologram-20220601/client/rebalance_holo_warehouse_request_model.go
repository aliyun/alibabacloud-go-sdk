// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebalanceHoloWarehouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *RebalanceHoloWarehouseRequest
	GetName() *string
}

type RebalanceHoloWarehouseRequest struct {
	// The name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_oss
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RebalanceHoloWarehouseRequest) String() string {
	return dara.Prettify(s)
}

func (s RebalanceHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *RebalanceHoloWarehouseRequest) GetName() *string {
	return s.Name
}

func (s *RebalanceHoloWarehouseRequest) SetName(v string) *RebalanceHoloWarehouseRequest {
	s.Name = &v
	return s
}

func (s *RebalanceHoloWarehouseRequest) Validate() error {
	return dara.Validate(s)
}
