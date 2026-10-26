// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteNetworkDomainRequest
	GetInstanceId() *string
	SetNetworkDomainId(v string) *DeleteNetworkDomainRequest
	GetNetworkDomainId() *string
	SetProjectId(v int64) *DeleteNetworkDomainRequest
	GetProjectId() *int64
	SetRegionId(v string) *DeleteNetworkDomainRequest
	GetRegionId() *string
}

type DeleteNetworkDomainRequest struct {
	// The instance ID of the bastion host where the network domain that you want to delete resides.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost_std_intl-sg-uq833e2dz02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network domain that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host instance where the network domain that you want to delete resides.
	//
	// > For information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNetworkDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteNetworkDomainRequest) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *DeleteNetworkDomainRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteNetworkDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNetworkDomainRequest) SetInstanceId(v string) *DeleteNetworkDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNetworkDomainRequest) SetNetworkDomainId(v string) *DeleteNetworkDomainRequest {
	s.NetworkDomainId = &v
	return s
}

func (s *DeleteNetworkDomainRequest) SetProjectId(v int64) *DeleteNetworkDomainRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteNetworkDomainRequest) SetRegionId(v string) *DeleteNetworkDomainRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNetworkDomainRequest) Validate() error {
	return dara.Validate(s)
}
