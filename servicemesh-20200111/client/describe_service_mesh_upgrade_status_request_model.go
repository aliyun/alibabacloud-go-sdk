// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshUpgradeStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllIstioGatewayFullNames(v string) *DescribeServiceMeshUpgradeStatusRequest
	GetAllIstioGatewayFullNames() *string
	SetGuestClusterIds(v string) *DescribeServiceMeshUpgradeStatusRequest
	GetGuestClusterIds() *string
	SetServiceMeshId(v string) *DescribeServiceMeshUpgradeStatusRequest
	GetServiceMeshId() *string
}

type DescribeServiceMeshUpgradeStatusRequest struct {
	// The fully qualified names of ingress gateways in the ASM instance. Separate multiple names with commas (,).
	//
	// example:
	//
	// istio-system:ingressgateway1,istio-system:ingressgateway2
	AllIstioGatewayFullNames *string `json:"AllIstioGatewayFullNames,omitempty" xml:"AllIstioGatewayFullNames,omitempty"`
	// The IDs of the clusters on the data plane of the ASM instance. Separate multiple clusters with commas (,).
	//
	// example:
	//
	// caeac85a793c94afbbb0a4bb20320****
	GuestClusterIds *string `json:"GuestClusterIds,omitempty" xml:"GuestClusterIds,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11fd0027-c27e-41bb-a565-75583054****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshUpgradeStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshUpgradeStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshUpgradeStatusRequest) GetAllIstioGatewayFullNames() *string {
	return s.AllIstioGatewayFullNames
}

func (s *DescribeServiceMeshUpgradeStatusRequest) GetGuestClusterIds() *string {
	return s.GuestClusterIds
}

func (s *DescribeServiceMeshUpgradeStatusRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeServiceMeshUpgradeStatusRequest) SetAllIstioGatewayFullNames(v string) *DescribeServiceMeshUpgradeStatusRequest {
	s.AllIstioGatewayFullNames = &v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusRequest) SetGuestClusterIds(v string) *DescribeServiceMeshUpgradeStatusRequest {
	s.GuestClusterIds = &v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusRequest) SetServiceMeshId(v string) *DescribeServiceMeshUpgradeStatusRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusRequest) Validate() error {
	return dara.Validate(s)
}
