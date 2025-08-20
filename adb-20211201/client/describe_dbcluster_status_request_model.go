// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeDBClusterStatusRequest
	GetRegionId() *string
}

type DescribeDBClusterStatusRequest struct {
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBClusterStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterStatusRequest) SetRegionId(v string) *DescribeDBClusterStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterStatusRequest) Validate() error {
	return dara.Validate(s)
}
