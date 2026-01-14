// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableAccelerateAreasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListAvailableAccelerateAreasRequest
	GetAcceleratorId() *string
	SetAccessMode(v string) *ListAvailableAccelerateAreasRequest
	GetAccessMode() *string
	SetRegionId(v string) *ListAvailableAccelerateAreasRequest
	GetRegionId() *string
}

type ListAvailableAccelerateAreasRequest struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1yeeq8yfoyszmqy****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	AccessMode    *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAvailableAccelerateAreasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableAccelerateAreasRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableAccelerateAreasRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListAvailableAccelerateAreasRequest) GetAccessMode() *string {
	return s.AccessMode
}

func (s *ListAvailableAccelerateAreasRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAvailableAccelerateAreasRequest) SetAcceleratorId(v string) *ListAvailableAccelerateAreasRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListAvailableAccelerateAreasRequest) SetAccessMode(v string) *ListAvailableAccelerateAreasRequest {
	s.AccessMode = &v
	return s
}

func (s *ListAvailableAccelerateAreasRequest) SetRegionId(v string) *ListAvailableAccelerateAreasRequest {
	s.RegionId = &v
	return s
}

func (s *ListAvailableAccelerateAreasRequest) Validate() error {
	return dara.Validate(s)
}
