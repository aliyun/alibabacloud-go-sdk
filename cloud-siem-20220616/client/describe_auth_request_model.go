// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeAuthRequest
	GetRegionId() *string
}

type DescribeAuthRequest struct {
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

func (s DescribeAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAuthRequest) SetRegionId(v string) *DescribeAuthRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAuthRequest) Validate() error {
	return dara.Validate(s)
}
