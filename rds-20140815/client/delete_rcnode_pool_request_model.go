// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCNodePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteRCNodePoolRequest
	GetClusterId() *string
	SetNodePoolId(v string) *DeleteRCNodePoolRequest
	GetNodePoolId() *string
	SetRegionId(v string) *DeleteRCNodePoolRequest
	GetRegionId() *string
}

type DeleteRCNodePoolRequest struct {
	// The ID of the ACK cluster to which the RDS Custom instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// c463aaa89e2b84cacacfbf23c4867****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The node pool ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// np31da1b38983f4511b490fc62108a****
	NodePoolId *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteRCNodePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCNodePoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteRCNodePoolRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteRCNodePoolRequest) GetNodePoolId() *string {
	return s.NodePoolId
}

func (s *DeleteRCNodePoolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRCNodePoolRequest) SetClusterId(v string) *DeleteRCNodePoolRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteRCNodePoolRequest) SetNodePoolId(v string) *DeleteRCNodePoolRequest {
	s.NodePoolId = &v
	return s
}

func (s *DeleteRCNodePoolRequest) SetRegionId(v string) *DeleteRCNodePoolRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRCNodePoolRequest) Validate() error {
	return dara.Validate(s)
}
