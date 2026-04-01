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
	// The ID of the ACK Edge cluster in which the RDS Custom instance resides.
	//
	// example:
	//
	// c463aaa89e2b84cacacfbf23c4867****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The node pool ID.
	//
	// example:
	//
	// np31da1b38983f4511b490fc62108a****
	NodePoolId *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
