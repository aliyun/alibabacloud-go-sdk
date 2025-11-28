// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeDeviceInfoRequest
	GetRegionId() *string
}

type DescribeDeviceInfoRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDeviceInfoRequest) SetRegionId(v string) *DescribeDeviceInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}
