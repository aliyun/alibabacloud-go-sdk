// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStatDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeStatDeviceInfoRequest
	GetRegionId() *string
}

type DescribeStatDeviceInfoRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeStatDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeStatDeviceInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStatDeviceInfoRequest) SetRegionId(v string) *DescribeStatDeviceInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStatDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}
