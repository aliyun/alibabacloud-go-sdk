// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachElasticNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticNetworkInterfaceId(v string) *AttachElasticNetworkInterfaceRequest
	GetElasticNetworkInterfaceId() *string
	SetNodeId(v string) *AttachElasticNetworkInterfaceRequest
	GetNodeId() *string
	SetRegionId(v string) *AttachElasticNetworkInterfaceRequest
	GetRegionId() *string
}

type AttachElasticNetworkInterfaceRequest struct {
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
	// e01-cn-lbj3aej****
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

func (s AttachElasticNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachElasticNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *AttachElasticNetworkInterfaceRequest) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *AttachElasticNetworkInterfaceRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *AttachElasticNetworkInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachElasticNetworkInterfaceRequest) SetElasticNetworkInterfaceId(v string) *AttachElasticNetworkInterfaceRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *AttachElasticNetworkInterfaceRequest) SetNodeId(v string) *AttachElasticNetworkInterfaceRequest {
	s.NodeId = &v
	return s
}

func (s *AttachElasticNetworkInterfaceRequest) SetRegionId(v string) *AttachElasticNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *AttachElasticNetworkInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
