// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeDrdsInstanceVersionRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *DescribeDrdsInstanceVersionRequest
	GetRegionId() *string
}

type DescribeDrdsInstanceVersionRequest struct {
	// The ID of the PolarDB-X 1.0 instance whose version you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbga1138****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsInstanceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceVersionRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsInstanceVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDrdsInstanceVersionRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceVersionRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceVersionRequest) SetRegionId(v string) *DescribeDrdsInstanceVersionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstanceVersionRequest) Validate() error {
	return dara.Validate(s)
}
