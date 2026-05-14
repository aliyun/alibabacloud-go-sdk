// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWarehouseAutoScaleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DisableWarehouseAutoScaleRequest
	GetName() *string
}

type DisableWarehouseAutoScaleRequest struct {
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DisableWarehouseAutoScaleRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableWarehouseAutoScaleRequest) GoString() string {
	return s.String()
}

func (s *DisableWarehouseAutoScaleRequest) GetName() *string {
	return s.Name
}

func (s *DisableWarehouseAutoScaleRequest) SetName(v string) *DisableWarehouseAutoScaleRequest {
	s.Name = &v
	return s
}

func (s *DisableWarehouseAutoScaleRequest) Validate() error {
	return dara.Validate(s)
}
