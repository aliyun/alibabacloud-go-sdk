// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCVClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeRCVClusterRequest
	GetClusterId() *string
	SetRegionId(v string) *DescribeRCVClusterRequest
	GetRegionId() *string
}

type DescribeRCVClusterRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRCVClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCVClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCVClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeRCVClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCVClusterRequest) SetClusterId(v string) *DescribeRCVClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRCVClusterRequest) SetRegionId(v string) *DescribeRCVClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCVClusterRequest) Validate() error {
	return dara.Validate(s)
}
