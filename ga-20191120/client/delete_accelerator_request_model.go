// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DeleteAcceleratorRequest
	GetAcceleratorId() *string
	SetRegionId(v string) *DeleteAcceleratorRequest
	GetRegionId() *string
}

type DeleteAcceleratorRequest struct {
	// The ID of the GA instance that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// Deprecated
	//
	// The ID of the region where your GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *DeleteAcceleratorRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DeleteAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAcceleratorRequest) SetAcceleratorId(v string) *DeleteAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteAcceleratorRequest) SetRegionId(v string) *DeleteAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAcceleratorRequest) Validate() error {
	return dara.Validate(s)
}
