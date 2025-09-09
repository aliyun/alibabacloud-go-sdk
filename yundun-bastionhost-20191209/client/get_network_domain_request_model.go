// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetNetworkDomainRequest
	GetInstanceId() *string
	SetNetworkDomainId(v string) *GetNetworkDomainRequest
	GetNetworkDomainId() *string
	SetRegionId(v string) *GetNetworkDomainRequest
	GetRegionId() *string
}

type GetNetworkDomainRequest struct {
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances ](https://help.aliyun.com/document_detail/153281.html)operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-i7m2btk6g48
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network domain to query.
	//
	// > You can call the [ListNetworkDomains ](https://help.aliyun.com/document_detail/2758827.html)operation to query the network domain ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
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

func (s GetNetworkDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkDomainRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNetworkDomainRequest) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *GetNetworkDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNetworkDomainRequest) SetInstanceId(v string) *GetNetworkDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNetworkDomainRequest) SetNetworkDomainId(v string) *GetNetworkDomainRequest {
	s.NetworkDomainId = &v
	return s
}

func (s *GetNetworkDomainRequest) SetRegionId(v string) *GetNetworkDomainRequest {
	s.RegionId = &v
	return s
}

func (s *GetNetworkDomainRequest) Validate() error {
	return dara.Validate(s)
}
