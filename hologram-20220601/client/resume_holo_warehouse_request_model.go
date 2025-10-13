// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeHoloWarehouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ResumeHoloWarehouseRequest
	GetName() *string
}

type ResumeHoloWarehouseRequest struct {
	// The name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ResumeHoloWarehouseRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *ResumeHoloWarehouseRequest) GetName() *string {
	return s.Name
}

func (s *ResumeHoloWarehouseRequest) SetName(v string) *ResumeHoloWarehouseRequest {
	s.Name = &v
	return s
}

func (s *ResumeHoloWarehouseRequest) Validate() error {
	return dara.Validate(s)
}
