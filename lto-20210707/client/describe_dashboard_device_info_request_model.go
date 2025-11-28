// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeDashboardDeviceInfoRequest
	GetRegionId() *string
}

type DescribeDashboardDeviceInfoRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDashboardDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDashboardDeviceInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDashboardDeviceInfoRequest) SetRegionId(v string) *DescribeDashboardDeviceInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDashboardDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}
