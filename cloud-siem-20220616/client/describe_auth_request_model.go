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
	// The region of the data management center for Threat Analysis. Select the management center region based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Select this value if your assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Select this value if your assets are located outside China.
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
