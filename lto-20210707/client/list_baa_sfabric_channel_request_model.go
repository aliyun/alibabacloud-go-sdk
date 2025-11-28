// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSFabricChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaaSFabricConsortiumId(v string) *ListBaaSFabricChannelRequest
	GetBaaSFabricConsortiumId() *string
	SetRegionId(v string) *ListBaaSFabricChannelRequest
	GetRegionId() *string
}

type ListBaaSFabricChannelRequest struct {
	// This parameter is required.
	BaaSFabricConsortiumId *string `json:"BaaSFabricConsortiumId,omitempty" xml:"BaaSFabricConsortiumId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBaaSFabricChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSFabricChannelRequest) GoString() string {
	return s.String()
}

func (s *ListBaaSFabricChannelRequest) GetBaaSFabricConsortiumId() *string {
	return s.BaaSFabricConsortiumId
}

func (s *ListBaaSFabricChannelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBaaSFabricChannelRequest) SetBaaSFabricConsortiumId(v string) *ListBaaSFabricChannelRequest {
	s.BaaSFabricConsortiumId = &v
	return s
}

func (s *ListBaaSFabricChannelRequest) SetRegionId(v string) *ListBaaSFabricChannelRequest {
	s.RegionId = &v
	return s
}

func (s *ListBaaSFabricChannelRequest) Validate() error {
	return dara.Validate(s)
}
