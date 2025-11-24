// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIstioGatewayDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHosts(v string) *DeleteIstioGatewayDomainsRequest
	GetHosts() *string
	SetIstioGatewayName(v string) *DeleteIstioGatewayDomainsRequest
	GetIstioGatewayName() *string
	SetLimit(v string) *DeleteIstioGatewayDomainsRequest
	GetLimit() *string
	SetNamespace(v string) *DeleteIstioGatewayDomainsRequest
	GetNamespace() *string
	SetPortName(v string) *DeleteIstioGatewayDomainsRequest
	GetPortName() *string
	SetServiceMeshId(v string) *DeleteIstioGatewayDomainsRequest
	GetServiceMeshId() *string
}

type DeleteIstioGatewayDomainsRequest struct {
	// The domain names of the one or more hosts that are exposed by the ASM gateway. Separate multiple domain names with commas (,).
	//
	// example:
	//
	// example.com,demo.com
	Hosts *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	// The name of the ASM gateway.
	//
	// example:
	//
	// ingressgateway
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The maximum number of Istio gateways to query.
	//
	// example:
	//
	// 10
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The port name.
	//
	// example:
	//
	// https
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DeleteIstioGatewayDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIstioGatewayDomainsRequest) GoString() string {
	return s.String()
}

func (s *DeleteIstioGatewayDomainsRequest) GetHosts() *string {
	return s.Hosts
}

func (s *DeleteIstioGatewayDomainsRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *DeleteIstioGatewayDomainsRequest) GetLimit() *string {
	return s.Limit
}

func (s *DeleteIstioGatewayDomainsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteIstioGatewayDomainsRequest) GetPortName() *string {
	return s.PortName
}

func (s *DeleteIstioGatewayDomainsRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DeleteIstioGatewayDomainsRequest) SetHosts(v string) *DeleteIstioGatewayDomainsRequest {
	s.Hosts = &v
	return s
}

func (s *DeleteIstioGatewayDomainsRequest) SetIstioGatewayName(v string) *DeleteIstioGatewayDomainsRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *DeleteIstioGatewayDomainsRequest) SetLimit(v string) *DeleteIstioGatewayDomainsRequest {
	s.Limit = &v
	return s
}

func (s *DeleteIstioGatewayDomainsRequest) SetNamespace(v string) *DeleteIstioGatewayDomainsRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteIstioGatewayDomainsRequest) SetPortName(v string) *DeleteIstioGatewayDomainsRequest {
	s.PortName = &v
	return s
}

func (s *DeleteIstioGatewayDomainsRequest) SetServiceMeshId(v string) *DeleteIstioGatewayDomainsRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DeleteIstioGatewayDomainsRequest) Validate() error {
	return dara.Validate(s)
}
