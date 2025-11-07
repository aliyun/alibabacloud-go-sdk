// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkPathIds(v []*string) *DeleteNetworkPathRequest
	GetNetworkPathIds() []*string
	SetRegionId(v string) *DeleteNetworkPathRequest
	GetRegionId() *string
}

type DeleteNetworkPathRequest struct {
	// The IDs of network paths.
	//
	// This parameter is required.
	NetworkPathIds []*string `json:"NetworkPathIds,omitempty" xml:"NetworkPathIds,omitempty" type:"Repeated"`
	// The region ID of the network path that you want to delete.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNetworkPathRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkPathRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPathRequest) GetNetworkPathIds() []*string {
	return s.NetworkPathIds
}

func (s *DeleteNetworkPathRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNetworkPathRequest) SetNetworkPathIds(v []*string) *DeleteNetworkPathRequest {
	s.NetworkPathIds = v
	return s
}

func (s *DeleteNetworkPathRequest) SetRegionId(v string) *DeleteNetworkPathRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNetworkPathRequest) Validate() error {
	return dara.Validate(s)
}
