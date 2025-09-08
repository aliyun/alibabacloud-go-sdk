// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeServiceStatusRequest
	GetRegionId() *string
}

type DescribeServiceStatusRequest struct {
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServiceStatusRequest) SetRegionId(v string) *DescribeServiceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
