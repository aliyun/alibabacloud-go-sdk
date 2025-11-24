// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIstioGatewayDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewaySecretDetails(v []*DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) *DescribeIstioGatewayDomainsResponseBody
	GetGatewaySecretDetails() []*DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails
	SetRequestId(v string) *DescribeIstioGatewayDomainsResponseBody
	GetRequestId() *string
}

type DescribeIstioGatewayDomainsResponseBody struct {
	// The domain names that are exposed by the ASM gateway.
	GatewaySecretDetails []*DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails `json:"GatewaySecretDetails,omitempty" xml:"GatewaySecretDetails,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIstioGatewayDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayDomainsResponseBody) GetGatewaySecretDetails() []*DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	return s.GatewaySecretDetails
}

func (s *DescribeIstioGatewayDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIstioGatewayDomainsResponseBody) SetGatewaySecretDetails(v []*DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) *DescribeIstioGatewayDomainsResponseBody {
	s.GatewaySecretDetails = v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBody) SetRequestId(v string) *DescribeIstioGatewayDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBody) Validate() error {
	if s.GatewaySecretDetails != nil {
		for _, item := range s.GatewaySecretDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails struct {
	// The name of the secret that contains the Transport Layer Security (TLS) certificate and certificate authority (CA) certificate.
	//
	// example:
	//
	// bookinfo-secret
	CredentialName *string `json:"CredentialName,omitempty" xml:"CredentialName,omitempty"`
	// The details of the domain name in the JSON format.
	//
	// example:
	//
	// {   "servers": [     {       "port": {         "number": 27018,         "name": "mongo",         "protocol": "MONGO"       },       "hosts": [         "*"       ]     }   ] }
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The list of domain names.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The name of the Istio gateway.
	//
	// example:
	//
	// ingressgateway
	GatewayCRName *string `json:"GatewayCRName,omitempty" xml:"GatewayCRName,omitempty"`
	// The namespace in which the ASM gateway resides.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The port name.
	//
	// example:
	//
	// https-demo
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The type of the protocol. Valid values: `HTTP`, `HTTPS`, `GRPC`, `HTTP2`, `MONGO`, `TCP`, and `TLS`.
	//
	// example:
	//
	// HTTPS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) GetCredentialName() *string {
	return s.CredentialName
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) GetDetail() *string {
	return s.Detail
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) GetGatewayCRName() *string {
	return s.GatewayCRName
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) GetPortName() *string {
	return s.PortName
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) SetCredentialName(v string) *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	s.CredentialName = &v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) SetDetail(v string) *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	s.Detail = &v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) SetDomains(v []*string) *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	s.Domains = v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) SetGatewayCRName(v string) *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	s.GatewayCRName = &v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) SetNamespace(v string) *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	s.Namespace = &v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) SetPortName(v string) *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	s.PortName = &v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) SetProtocol(v string) *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails {
	s.Protocol = &v
	return s
}

func (s *DescribeIstioGatewayDomainsResponseBodyGatewaySecretDetails) Validate() error {
	return dara.Validate(s)
}
