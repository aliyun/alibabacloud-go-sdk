// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIstioGatewayDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredential(v string) *CreateIstioGatewayDomainsRequest
	GetCredential() *string
	SetForceHttps(v bool) *CreateIstioGatewayDomainsRequest
	GetForceHttps() *bool
	SetHosts(v string) *CreateIstioGatewayDomainsRequest
	GetHosts() *string
	SetIstioGatewayName(v string) *CreateIstioGatewayDomainsRequest
	GetIstioGatewayName() *string
	SetLimit(v string) *CreateIstioGatewayDomainsRequest
	GetLimit() *string
	SetNamespace(v string) *CreateIstioGatewayDomainsRequest
	GetNamespace() *string
	SetNumber(v int32) *CreateIstioGatewayDomainsRequest
	GetNumber() *int32
	SetPortName(v string) *CreateIstioGatewayDomainsRequest
	GetPortName() *string
	SetProtocol(v string) *CreateIstioGatewayDomainsRequest
	GetProtocol() *string
	SetServiceMeshId(v string) *CreateIstioGatewayDomainsRequest
	GetServiceMeshId() *string
}

type CreateIstioGatewayDomainsRequest struct {
	// The name of the secret that contains the Transport Layer Security (TLS) certificate and certificate authority (CA) certificate.
	//
	// example:
	//
	// bookinfo-secret
	Credential *string `json:"Credential,omitempty" xml:"Credential,omitempty"`
	// Specifies whether to forcibly use TLS to protect connection security.
	//
	// 	- `true`: forcibly uses TLS to protect connection security.
	//
	// 	- `false`: does not forcibly use TLS to protect connection security.
	//
	// example:
	//
	// true
	ForceHttps *bool `json:"ForceHttps,omitempty" xml:"ForceHttps,omitempty"`
	// The one or more domain names that are exposed by the ASM gateway. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,demo.com
	Hosts *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	// The name of the ASM gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ingressgateway
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The maximum number of ASM gateways to query.
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
	// The port that is provided by the ASM gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// 443
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// The name of the port.
	//
	// This parameter is required.
	//
	// example:
	//
	// http-demo
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The type of the protocol. Valid values: `HTTP`, `HTTPS`, `GRPC`, `HTTP2`, `MONGO`, `TCP`, and `TLS`.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTPS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s CreateIstioGatewayDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayDomainsRequest) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayDomainsRequest) GetCredential() *string {
	return s.Credential
}

func (s *CreateIstioGatewayDomainsRequest) GetForceHttps() *bool {
	return s.ForceHttps
}

func (s *CreateIstioGatewayDomainsRequest) GetHosts() *string {
	return s.Hosts
}

func (s *CreateIstioGatewayDomainsRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *CreateIstioGatewayDomainsRequest) GetLimit() *string {
	return s.Limit
}

func (s *CreateIstioGatewayDomainsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateIstioGatewayDomainsRequest) GetNumber() *int32 {
	return s.Number
}

func (s *CreateIstioGatewayDomainsRequest) GetPortName() *string {
	return s.PortName
}

func (s *CreateIstioGatewayDomainsRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateIstioGatewayDomainsRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *CreateIstioGatewayDomainsRequest) SetCredential(v string) *CreateIstioGatewayDomainsRequest {
	s.Credential = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetForceHttps(v bool) *CreateIstioGatewayDomainsRequest {
	s.ForceHttps = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetHosts(v string) *CreateIstioGatewayDomainsRequest {
	s.Hosts = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetIstioGatewayName(v string) *CreateIstioGatewayDomainsRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetLimit(v string) *CreateIstioGatewayDomainsRequest {
	s.Limit = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetNamespace(v string) *CreateIstioGatewayDomainsRequest {
	s.Namespace = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetNumber(v int32) *CreateIstioGatewayDomainsRequest {
	s.Number = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetPortName(v string) *CreateIstioGatewayDomainsRequest {
	s.PortName = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetProtocol(v string) *CreateIstioGatewayDomainsRequest {
	s.Protocol = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) SetServiceMeshId(v string) *CreateIstioGatewayDomainsRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *CreateIstioGatewayDomainsRequest) Validate() error {
	return dara.Validate(s)
}
