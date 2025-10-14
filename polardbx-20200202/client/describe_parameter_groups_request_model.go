// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeParameterGroupsRequest
	GetRegionId() *string
}

type DescribeParameterGroupsRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeParameterGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeParameterGroupsRequest) SetRegionId(v string) *DescribeParameterGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParameterGroupsRequest) Validate() error {
	return dara.Validate(s)
}
