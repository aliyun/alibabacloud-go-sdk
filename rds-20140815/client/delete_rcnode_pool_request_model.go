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
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	NodePoolId *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	// This parameter is required.
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
