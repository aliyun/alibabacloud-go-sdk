// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIstioGatewayDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIstioGatewayName(v string) *DescribeIstioGatewayDomainsRequest
	GetIstioGatewayName() *string
	SetLimit(v string) *DescribeIstioGatewayDomainsRequest
	GetLimit() *string
	SetNamespace(v string) *DescribeIstioGatewayDomainsRequest
	GetNamespace() *string
	SetServiceMeshId(v string) *DescribeIstioGatewayDomainsRequest
	GetServiceMeshId() *string
}

type DescribeIstioGatewayDomainsRequest struct {
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
	// The namespace in which the ASM gateway resides.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeIstioGatewayDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayDomainsRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *DescribeIstioGatewayDomainsRequest) GetLimit() *string {
	return s.Limit
}

func (s *DescribeIstioGatewayDomainsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeIstioGatewayDomainsRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeIstioGatewayDomainsRequest) SetIstioGatewayName(v string) *DescribeIstioGatewayDomainsRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *DescribeIstioGatewayDomainsRequest) SetLimit(v string) *DescribeIstioGatewayDomainsRequest {
	s.Limit = &v
	return s
}

func (s *DescribeIstioGatewayDomainsRequest) SetNamespace(v string) *DescribeIstioGatewayDomainsRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeIstioGatewayDomainsRequest) SetServiceMeshId(v string) *DescribeIstioGatewayDomainsRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeIstioGatewayDomainsRequest) Validate() error {
	return dara.Validate(s)
}
