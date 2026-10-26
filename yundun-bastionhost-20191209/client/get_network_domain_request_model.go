// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckProxyState(v string) *GetNetworkDomainRequest
	GetCheckProxyState() *string
	SetInstanceId(v string) *GetNetworkDomainRequest
	GetInstanceId() *string
	SetNetworkDomainId(v string) *GetNetworkDomainRequest
	GetNetworkDomainId() *string
	SetProjectId(v int64) *GetNetworkDomainRequest
	GetProjectId() *int64
	SetRegionId(v string) *GetNetworkDomainRequest
	GetRegionId() *string
}

type GetNetworkDomainRequest struct {
	// Specifies whether to immediately re-check the proxy server status. Valid values:
	//
	// - **true**: Immediately re-checks the proxy server status and returns the latest ProxyState/ProxyStateErrorCode.
	//
	// - **false**: Default value. Does not re-check the proxy server status and directly returns the currently recorded status.
	//
	// example:
	//
	// false
	CheckProxyState *string `json:"CheckProxyState,omitempty" xml:"CheckProxyState,omitempty"`
	// The instance ID of the bastion host.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-i7m2btk6g48
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network domain to query.
	//
	// > You can call the [ListNetworkDomains](https://help.aliyun.com/document_detail/2758827.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *GetNetworkDomainRequest) GetCheckProxyState() *string {
	return s.CheckProxyState
}

func (s *GetNetworkDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNetworkDomainRequest) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *GetNetworkDomainRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetNetworkDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNetworkDomainRequest) SetCheckProxyState(v string) *GetNetworkDomainRequest {
	s.CheckProxyState = &v
	return s
}

func (s *GetNetworkDomainRequest) SetInstanceId(v string) *GetNetworkDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNetworkDomainRequest) SetNetworkDomainId(v string) *GetNetworkDomainRequest {
	s.NetworkDomainId = &v
	return s
}

func (s *GetNetworkDomainRequest) SetProjectId(v int64) *GetNetworkDomainRequest {
	s.ProjectId = &v
	return s
}

func (s *GetNetworkDomainRequest) SetRegionId(v string) *GetNetworkDomainRequest {
	s.RegionId = &v
	return s
}

func (s *GetNetworkDomainRequest) Validate() error {
	return dara.Validate(s)
}
