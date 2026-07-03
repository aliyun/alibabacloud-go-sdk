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
	// The region of the Data Management hub. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: For assets in the Chinese mainland and China (Hong Kong).
	//
	// - ap-southeast-1: For assets in regions outside China.
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
