// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateASMGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *CreateASMGatewayRequest
	GetBody() *string
	SetIstioGatewayName(v string) *CreateASMGatewayRequest
	GetIstioGatewayName() *string
	SetServiceMeshId(v string) *CreateASMGatewayRequest
	GetServiceMeshId() *string
}

type CreateASMGatewayRequest struct {
	// The YAML content that is used to create the ASM gateway.
	//
	// example:
	//
	// {"apiVersion":"istio.alibabacloud.com/v1beta1","kind":"IstioGateway","metadata":{"name":"ingressgateway","namespace":"istio-system"},"spec":{"gatewayType":"ingress","clusterIds":["xxxxx"],"ports":[{"name":"http-0","port":80,"targetPort":80,"protocol":"TCP"},{"name":"https-1","port":443,"targetPort":443,"protocol":"TCP"}],"serviceAnnotations":{"service.beta.kubernetes.io/alicloud-loadbalancer-address-type":"internet","service.beta.kubernetes.io/alibaba-cloud-loadbalancer-spec":"slb.s1.small"},"replicaCount":2,"resources":{"limits":{"cpu":"2","memory":"4G"},"requests":{"cpu":"200m","memory":"256Mi"}},"serviceType":"LoadBalancer","maxReplicas":2,"minReplicas":2}}
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The name of the ASM gateway.
	//
	// example:
	//
	// ingressgateway
	IstioGatewayName *string `json:"IstioGatewayName,omitempty" xml:"IstioGatewayName,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s CreateASMGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateASMGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateASMGatewayRequest) GetBody() *string {
	return s.Body
}

func (s *CreateASMGatewayRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *CreateASMGatewayRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *CreateASMGatewayRequest) SetBody(v string) *CreateASMGatewayRequest {
	s.Body = &v
	return s
}

func (s *CreateASMGatewayRequest) SetIstioGatewayName(v string) *CreateASMGatewayRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *CreateASMGatewayRequest) SetServiceMeshId(v string) *CreateASMGatewayRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *CreateASMGatewayRequest) Validate() error {
	return dara.Validate(s)
}
