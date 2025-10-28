// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateK8sServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateK8sServiceRequest
	GetAppId() *string
	SetExternalTrafficPolicy(v string) *CreateK8sServiceRequest
	GetExternalTrafficPolicy() *string
	SetName(v string) *CreateK8sServiceRequest
	GetName() *string
	SetServicePorts(v string) *CreateK8sServiceRequest
	GetServicePorts() *string
	SetType(v string) *CreateK8sServiceRequest
	GetType() *string
}

type CreateK8sServiceRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a166fbd-****-****-a286-781659d9f54c
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The policy used for external traffic management. Valid values:
	//
	// 	- Local: The network traffic can be routed to pods on the node where the Service is deployed.
	//
	// 	- Cluster: The network traffic can be routed to pods on other nodes in the cluster.
	//
	// Default value: Local.
	//
	// example:
	//
	// Local
	ExternalTrafficPolicy *string `json:"ExternalTrafficPolicy,omitempty" xml:"ExternalTrafficPolicy,omitempty"`
	// The name of the Kubernetes Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-http
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port mapping of the Kubernetes Service. Set this parameter to a JSON array. The following parameters are included in the configurations:
	//
	// 	- **protocol**: the protocol used by the Service. Valid values: TCP and UDP. This parameter is mandatory.
	//
	// 	- **port**: the frontend service port. Valid values: 1 to 65535. This parameter is mandatory.
	//
	// 	- **targetPort**: the backend container port. Valid values: 1 to 65535. This parameter is mandatory.
	//
	// Example: `[{"protocol": "TCP", "port": 80, "targetPort": 8080},{"protocol": "TCP", "port": 81, "targetPort": 8081}]`
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"protocol":"TCP","port":80,"targetPort":8080}]
	ServicePorts *string `json:"ServicePorts,omitempty" xml:"ServicePorts,omitempty"`
	// The type of the Kubernetes Service. Set the value to ClusterIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// ClusterIP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateK8sServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateK8sServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateK8sServiceRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateK8sServiceRequest) GetExternalTrafficPolicy() *string {
	return s.ExternalTrafficPolicy
}

func (s *CreateK8sServiceRequest) GetName() *string {
	return s.Name
}

func (s *CreateK8sServiceRequest) GetServicePorts() *string {
	return s.ServicePorts
}

func (s *CreateK8sServiceRequest) GetType() *string {
	return s.Type
}

func (s *CreateK8sServiceRequest) SetAppId(v string) *CreateK8sServiceRequest {
	s.AppId = &v
	return s
}

func (s *CreateK8sServiceRequest) SetExternalTrafficPolicy(v string) *CreateK8sServiceRequest {
	s.ExternalTrafficPolicy = &v
	return s
}

func (s *CreateK8sServiceRequest) SetName(v string) *CreateK8sServiceRequest {
	s.Name = &v
	return s
}

func (s *CreateK8sServiceRequest) SetServicePorts(v string) *CreateK8sServiceRequest {
	s.ServicePorts = &v
	return s
}

func (s *CreateK8sServiceRequest) SetType(v string) *CreateK8sServiceRequest {
	s.Type = &v
	return s
}

func (s *CreateK8sServiceRequest) Validate() error {
	return dara.Validate(s)
}
