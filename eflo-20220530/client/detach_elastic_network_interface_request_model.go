// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachElasticNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticNetworkInterfaceId(v string) *DetachElasticNetworkInterfaceRequest
	GetElasticNetworkInterfaceId() *string
	SetNodeId(v string) *DetachElasticNetworkInterfaceRequest
	GetNodeId() *string
	SetRegionId(v string) *DetachElasticNetworkInterfaceRequest
	GetRegionId() *string
}

type DetachElasticNetworkInterfaceRequest struct {
	// The ID of the ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// e01-cn-zxu2zp3****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachElasticNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachElasticNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *DetachElasticNetworkInterfaceRequest) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *DetachElasticNetworkInterfaceRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DetachElasticNetworkInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachElasticNetworkInterfaceRequest) SetElasticNetworkInterfaceId(v string) *DetachElasticNetworkInterfaceRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *DetachElasticNetworkInterfaceRequest) SetNodeId(v string) *DetachElasticNetworkInterfaceRequest {
	s.NodeId = &v
	return s
}

func (s *DetachElasticNetworkInterfaceRequest) SetRegionId(v string) *DetachElasticNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *DetachElasticNetworkInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
