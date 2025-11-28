// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStatMemberDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeStatMemberDeviceInfoRequest
	GetRegionId() *string
}

type DescribeStatMemberDeviceInfoRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeStatMemberDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatMemberDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeStatMemberDeviceInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStatMemberDeviceInfoRequest) SetRegionId(v string) *DescribeStatMemberDeviceInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStatMemberDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}
