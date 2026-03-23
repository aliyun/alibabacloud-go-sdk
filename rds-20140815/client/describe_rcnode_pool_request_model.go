// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeRCNodePoolRequest
	GetClusterId() *string
	SetNodePoolId(v string) *DescribeRCNodePoolRequest
	GetNodePoolId() *string
	SetRegionId(v string) *DescribeRCNodePoolRequest
	GetRegionId() *string
}

type DescribeRCNodePoolRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	NodePoolId *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRCNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCNodePoolRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCNodePoolRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeRCNodePoolRequest) GetNodePoolId() *string {
	return s.NodePoolId
}

func (s *DescribeRCNodePoolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCNodePoolRequest) SetClusterId(v string) *DescribeRCNodePoolRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRCNodePoolRequest) SetNodePoolId(v string) *DescribeRCNodePoolRequest {
	s.NodePoolId = &v
	return s
}

func (s *DescribeRCNodePoolRequest) SetRegionId(v string) *DescribeRCNodePoolRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCNodePoolRequest) Validate() error {
	return dara.Validate(s)
}
