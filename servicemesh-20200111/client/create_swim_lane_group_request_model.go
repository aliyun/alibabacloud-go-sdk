// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSwimLaneGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *CreateSwimLaneGroupRequest
	GetGroupName() *string
	SetIngressGatewayName(v string) *CreateSwimLaneGroupRequest
	GetIngressGatewayName() *string
	SetIngressGatewayNamespace(v string) *CreateSwimLaneGroupRequest
	GetIngressGatewayNamespace() *string
	SetIngressType(v string) *CreateSwimLaneGroupRequest
	GetIngressType() *string
	SetIsPermissive(v bool) *CreateSwimLaneGroupRequest
	GetIsPermissive() *bool
	SetRouteHeader(v string) *CreateSwimLaneGroupRequest
	GetRouteHeader() *string
	SetServiceMeshId(v string) *CreateSwimLaneGroupRequest
	GetServiceMeshId() *string
	SetServicesList(v string) *CreateSwimLaneGroupRequest
	GetServicesList() *string
	SetTraceHeader(v string) *CreateSwimLaneGroupRequest
	GetTraceHeader() *string
}

type CreateSwimLaneGroupRequest struct {
	// The name of the lane group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the ingress gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ingressgateway
	IngressGatewayName *string `json:"IngressGatewayName,omitempty" xml:"IngressGatewayName,omitempty"`
	// example:
	//
	// istio-system
	IngressGatewayNamespace *string `json:"IngressGatewayNamespace,omitempty" xml:"IngressGatewayNamespace,omitempty"`
	// The type of the gateway for ingress traffic. Only ASM ingress gateways are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// ASM
	IngressType *string `json:"IngressType,omitempty" xml:"IngressType,omitempty"`
	// Specifies whether the permissive mode is enabled for the lane group to be created.
	//
	// example:
	//
	// false
	IsPermissive *bool `json:"IsPermissive,omitempty" xml:"IsPermissive,omitempty"`
	// The request routing header of the lane group if you plan to create a lane group in permissive mode. This parameter must be specified when IsPermissive is set to true.
	//
	// example:
	//
	// x-asm-prefer-tag
	RouteHeader *string `json:"RouteHeader,omitempty" xml:"RouteHeader,omitempty"`
	// The ID of the Alibaba Cloud Service Mesh (ASM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// A list of services associated with the lane group. The value is a JSON array. The format of a service is `$Cluster name/$Cluster ID/$Namespace/$Service name`.
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mocka\\",\\"sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mockb\\",\\"sh01/c089443ea9e50403fa4f0a6237d11e0a9/default/mockc\\"]
	ServicesList *string `json:"ServicesList,omitempty" xml:"ServicesList,omitempty"`
	// The end-to-end (E2E) pass-through request header of the lane group if you plan to create a lane group in permissive mode. This parameter must be specified when IsPermissive is set to true.
	//
	// example:
	//
	// my-request-id
	TraceHeader *string `json:"TraceHeader,omitempty" xml:"TraceHeader,omitempty"`
}

func (s CreateSwimLaneGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSwimLaneGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSwimLaneGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateSwimLaneGroupRequest) GetIngressGatewayName() *string {
	return s.IngressGatewayName
}

func (s *CreateSwimLaneGroupRequest) GetIngressGatewayNamespace() *string {
	return s.IngressGatewayNamespace
}

func (s *CreateSwimLaneGroupRequest) GetIngressType() *string {
	return s.IngressType
}

func (s *CreateSwimLaneGroupRequest) GetIsPermissive() *bool {
	return s.IsPermissive
}

func (s *CreateSwimLaneGroupRequest) GetRouteHeader() *string {
	return s.RouteHeader
}

func (s *CreateSwimLaneGroupRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *CreateSwimLaneGroupRequest) GetServicesList() *string {
	return s.ServicesList
}

func (s *CreateSwimLaneGroupRequest) GetTraceHeader() *string {
	return s.TraceHeader
}

func (s *CreateSwimLaneGroupRequest) SetGroupName(v string) *CreateSwimLaneGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetIngressGatewayName(v string) *CreateSwimLaneGroupRequest {
	s.IngressGatewayName = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetIngressGatewayNamespace(v string) *CreateSwimLaneGroupRequest {
	s.IngressGatewayNamespace = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetIngressType(v string) *CreateSwimLaneGroupRequest {
	s.IngressType = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetIsPermissive(v bool) *CreateSwimLaneGroupRequest {
	s.IsPermissive = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetRouteHeader(v string) *CreateSwimLaneGroupRequest {
	s.RouteHeader = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetServiceMeshId(v string) *CreateSwimLaneGroupRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetServicesList(v string) *CreateSwimLaneGroupRequest {
	s.ServicesList = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) SetTraceHeader(v string) *CreateSwimLaneGroupRequest {
	s.TraceHeader = &v
	return s
}

func (s *CreateSwimLaneGroupRequest) Validate() error {
	return dara.Validate(s)
}
