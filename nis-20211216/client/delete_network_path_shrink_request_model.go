// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkPathShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkPathIdsShrink(v string) *DeleteNetworkPathShrinkRequest
	GetNetworkPathIdsShrink() *string
	SetRegionId(v string) *DeleteNetworkPathShrinkRequest
	GetRegionId() *string
}

type DeleteNetworkPathShrinkRequest struct {
	// The IDs of network paths.
	//
	// This parameter is required.
	NetworkPathIdsShrink *string `json:"NetworkPathIds,omitempty" xml:"NetworkPathIds,omitempty"`
	// The region ID of the network path that you want to delete.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNetworkPathShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkPathShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPathShrinkRequest) GetNetworkPathIdsShrink() *string {
	return s.NetworkPathIdsShrink
}

func (s *DeleteNetworkPathShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNetworkPathShrinkRequest) SetNetworkPathIdsShrink(v string) *DeleteNetworkPathShrinkRequest {
	s.NetworkPathIdsShrink = &v
	return s
}

func (s *DeleteNetworkPathShrinkRequest) SetRegionId(v string) *DeleteNetworkPathShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNetworkPathShrinkRequest) Validate() error {
	return dara.Validate(s)
}
