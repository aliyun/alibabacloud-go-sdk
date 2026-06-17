// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *DescribeRegionsRequest
	GetRegion() *string
}

type DescribeRegionsRequest struct {
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeRegionsRequest) SetRegion(v string) *DescribeRegionsRequest {
	s.Region = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
