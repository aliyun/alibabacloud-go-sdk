// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeDrdsInstanceRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *DescribeDrdsInstanceRequest
	GetRegionId() *string
}

type DescribeDrdsInstanceRequest struct {
	// The ID of the instance that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbga1138****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region in which the instance is created.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDrdsInstanceRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceRequest) SetRegionId(v string) *DescribeDrdsInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstanceRequest) Validate() error {
	return dara.Validate(s)
}
