// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveHostsToNetworkDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostIds(v []*string) *MoveHostsToNetworkDomainRequest
	GetHostIds() []*string
	SetInstanceId(v string) *MoveHostsToNetworkDomainRequest
	GetInstanceId() *string
	SetNetworkDomainId(v string) *MoveHostsToNetworkDomainRequest
	GetNetworkDomainId() *string
	SetRegionId(v string) *MoveHostsToNetworkDomainRequest
	GetRegionId() *string
}

type MoveHostsToNetworkDomainRequest struct {
	// The IDs of the hosts that you want to add to the network domain.
	//
	// This parameter is required.
	HostIds []*string `json:"HostIds,omitempty" xml:"HostIds,omitempty" type:"Repeated"`
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-uax2zmx8005
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network domain to which you want to add hosts.
	//
	// >  You can call the [ListNetworkDomains](https://help.aliyun.com/document_detail/2758827.html) operation to query the network domain ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s MoveHostsToNetworkDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveHostsToNetworkDomainRequest) GoString() string {
	return s.String()
}

func (s *MoveHostsToNetworkDomainRequest) GetHostIds() []*string {
	return s.HostIds
}

func (s *MoveHostsToNetworkDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MoveHostsToNetworkDomainRequest) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *MoveHostsToNetworkDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MoveHostsToNetworkDomainRequest) SetHostIds(v []*string) *MoveHostsToNetworkDomainRequest {
	s.HostIds = v
	return s
}

func (s *MoveHostsToNetworkDomainRequest) SetInstanceId(v string) *MoveHostsToNetworkDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *MoveHostsToNetworkDomainRequest) SetNetworkDomainId(v string) *MoveHostsToNetworkDomainRequest {
	s.NetworkDomainId = &v
	return s
}

func (s *MoveHostsToNetworkDomainRequest) SetRegionId(v string) *MoveHostsToNetworkDomainRequest {
	s.RegionId = &v
	return s
}

func (s *MoveHostsToNetworkDomainRequest) Validate() error {
	return dara.Validate(s)
}
