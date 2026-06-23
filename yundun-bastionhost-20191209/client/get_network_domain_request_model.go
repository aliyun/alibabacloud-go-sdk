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
	SetRegionId(v string) *GetNetworkDomainRequest
	GetRegionId() *string
}

type GetNetworkDomainRequest struct {
	// Indicates whether to immediately recheck the status of the proxy server. Valid values:
	//
	// - **true**: Immediately rechecks the status of the proxy server and returns the latest ProxyState and ProxyStateErrorCode.
	//
	// - **false**: (Default) Returns the currently recorded status without rechecking the proxy server.
	//
	// example:
	//
	// false
	CheckProxyState *string `json:"CheckProxyState,omitempty" xml:"CheckProxyState,omitempty"`
	// The ID of the Bastionhost instance.
	//
	// > Call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to get this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-i7m2btk6g48
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network domain to query.
	//
	// > Call the [ListNetworkDomains](https://help.aliyun.com/document_detail/2758827.html) operation to get this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about region IDs, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *GetNetworkDomainRequest) SetRegionId(v string) *GetNetworkDomainRequest {
	s.RegionId = &v
	return s
}

func (s *GetNetworkDomainRequest) Validate() error {
	return dara.Validate(s)
}
