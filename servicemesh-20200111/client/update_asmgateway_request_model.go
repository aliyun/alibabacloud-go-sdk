// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateASMGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdateASMGatewayRequest
	GetBody() *string
	SetIstioGatewayName(v string) *UpdateASMGatewayRequest
	GetIstioGatewayName() *string
	SetServiceMeshId(v string) *UpdateASMGatewayRequest
	GetServiceMeshId() *string
}

type UpdateASMGatewayRequest struct {
	// The new YAML content of the ASM gateway.
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

func (s UpdateASMGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateASMGatewayRequest) GoString() string {
	return s.String()
}

func (s *UpdateASMGatewayRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateASMGatewayRequest) GetIstioGatewayName() *string {
	return s.IstioGatewayName
}

func (s *UpdateASMGatewayRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateASMGatewayRequest) SetBody(v string) *UpdateASMGatewayRequest {
	s.Body = &v
	return s
}

func (s *UpdateASMGatewayRequest) SetIstioGatewayName(v string) *UpdateASMGatewayRequest {
	s.IstioGatewayName = &v
	return s
}

func (s *UpdateASMGatewayRequest) SetServiceMeshId(v string) *UpdateASMGatewayRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateASMGatewayRequest) Validate() error {
	return dara.Validate(s)
}
