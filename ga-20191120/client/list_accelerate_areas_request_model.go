// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccelerateAreasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListAccelerateAreasRequest
	GetRegionId() *string
}

type ListAccelerateAreasRequest struct {
	// The region ID of the Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAccelerateAreasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccelerateAreasRequest) GoString() string {
	return s.String()
}

func (s *ListAccelerateAreasRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAccelerateAreasRequest) SetRegionId(v string) *ListAccelerateAreasRequest {
	s.RegionId = &v
	return s
}

func (s *ListAccelerateAreasRequest) Validate() error {
	return dara.Validate(s)
}
