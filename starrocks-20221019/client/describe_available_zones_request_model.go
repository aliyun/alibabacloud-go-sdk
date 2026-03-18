// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeAvailableZonesRequest
	GetRegionId() *string
}

type DescribeAvailableZonesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAvailableZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableZonesRequest) SetRegionId(v string) *DescribeAvailableZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableZonesRequest) Validate() error {
	return dara.Validate(s)
}
