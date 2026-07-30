// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualBridgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBridgeId(v string) *DeleteVirtualBridgeRequest
	GetBridgeId() *string
	SetRegionId(v string) *DeleteVirtualBridgeRequest
	GetRegionId() *string
}

type DeleteVirtualBridgeRequest struct {
	// The virtual bridge ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vb-fjsidhfishfiu****
	BridgeId *string `json:"BridgeId,omitempty" xml:"BridgeId,omitempty"`
	// The region ID. Call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteVirtualBridgeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualBridgeRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualBridgeRequest) GetBridgeId() *string {
	return s.BridgeId
}

func (s *DeleteVirtualBridgeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVirtualBridgeRequest) SetBridgeId(v string) *DeleteVirtualBridgeRequest {
	s.BridgeId = &v
	return s
}

func (s *DeleteVirtualBridgeRequest) SetRegionId(v string) *DeleteVirtualBridgeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVirtualBridgeRequest) Validate() error {
	return dara.Validate(s)
}
