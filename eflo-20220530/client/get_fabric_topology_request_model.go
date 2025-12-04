// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFabricTopologyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetFabricTopologyRequest
	GetClusterId() *string
	SetLniIds(v []*string) *GetFabricTopologyRequest
	GetLniIds() []*string
	SetNodeIds(v []*string) *GetFabricTopologyRequest
	GetNodeIds() []*string
	SetRegionId(v string) *GetFabricTopologyRequest
	GetRegionId() *string
	SetVpcId(v string) *GetFabricTopologyRequest
	GetVpcId() *string
	SetVpdId(v string) *GetFabricTopologyRequest
	GetVpdId() *string
}

type GetFabricTopologyRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i-169263721924****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Lingjun network interface controller ID List
	LniIds []*string `json:"LniIds,omitempty" xml:"LniIds,omitempty" type:"Repeated"`
	// Node ID list
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-k8i0g9fk68t7u0u2w****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Lingjun CIDR block ID
	//
	// example:
	//
	// vpd-aof7****
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s GetFabricTopologyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFabricTopologyRequest) GoString() string {
	return s.String()
}

func (s *GetFabricTopologyRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetFabricTopologyRequest) GetLniIds() []*string {
	return s.LniIds
}

func (s *GetFabricTopologyRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *GetFabricTopologyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetFabricTopologyRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *GetFabricTopologyRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *GetFabricTopologyRequest) SetClusterId(v string) *GetFabricTopologyRequest {
	s.ClusterId = &v
	return s
}

func (s *GetFabricTopologyRequest) SetLniIds(v []*string) *GetFabricTopologyRequest {
	s.LniIds = v
	return s
}

func (s *GetFabricTopologyRequest) SetNodeIds(v []*string) *GetFabricTopologyRequest {
	s.NodeIds = v
	return s
}

func (s *GetFabricTopologyRequest) SetRegionId(v string) *GetFabricTopologyRequest {
	s.RegionId = &v
	return s
}

func (s *GetFabricTopologyRequest) SetVpcId(v string) *GetFabricTopologyRequest {
	s.VpcId = &v
	return s
}

func (s *GetFabricTopologyRequest) SetVpdId(v string) *GetFabricTopologyRequest {
	s.VpdId = &v
	return s
}

func (s *GetFabricTopologyRequest) Validate() error {
	return dara.Validate(s)
}
