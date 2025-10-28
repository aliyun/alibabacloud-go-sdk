// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateK8sServiceRequest
	GetAppId() *string
	SetExternalTrafficPolicy(v string) *UpdateK8sServiceRequest
	GetExternalTrafficPolicy() *string
	SetName(v string) *UpdateK8sServiceRequest
	GetName() *string
	SetServicePorts(v string) *UpdateK8sServiceRequest
	GetServicePorts() *string
	SetType(v string) *UpdateK8sServiceRequest
	GetType() *string
}

type UpdateK8sServiceRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5a166fbd-****-****-a286-781659d9f54c
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The policy used for external traffic management. Valid values:
	//
	// 	- Local: local mode
	//
	// 	- Cluster: cluster mode
	//
	// Default value: Local.
	//
	// example:
	//
	// Local
	ExternalTrafficPolicy *string `json:"ExternalTrafficPolicy,omitempty" xml:"ExternalTrafficPolicy,omitempty"`
	// The name of the service in a Kubernetes cluster.
	//
	// 	- The name can contain lowercase letters, digits, and hyphens (-).
	//
	// 	- It must start with a letter and end with a letter or digit.
	//
	// 	- The name can be 2 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-http
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The mappings between service ports. Set this parameter to a JSON array. The following parameters are included in the configurations:
	//
	// 	- **protocol**: the protocol used by the service. Valid values: TCP and UDP. This parameter is required.
	//
	// 	- **port**: the frontend service port. Valid values: 1 to 65535. This parameter is required.
	//
	// 	- **targetPort**: the backend container port. Valid values: 1 to 65535. This parameter is required.
	//
	// Example: `[{"protocol": "TCP", "port": 80, "targetPort": 8080},{"protocol": "TCP", "port": 81, "targetPort": 8081}]`
	//
	// example:
	//
	// [{"protocol":"TCP","port":80,"targetPort":8080}]
	ServicePorts *string `json:"ServicePorts,omitempty" xml:"ServicePorts,omitempty"`
	// The type of the service in a Kubernetes cluster. Set the value to ClusterIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// ClusterIP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateK8sServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sServiceRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateK8sServiceRequest) GetExternalTrafficPolicy() *string {
	return s.ExternalTrafficPolicy
}

func (s *UpdateK8sServiceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateK8sServiceRequest) GetServicePorts() *string {
	return s.ServicePorts
}

func (s *UpdateK8sServiceRequest) GetType() *string {
	return s.Type
}

func (s *UpdateK8sServiceRequest) SetAppId(v string) *UpdateK8sServiceRequest {
	s.AppId = &v
	return s
}

func (s *UpdateK8sServiceRequest) SetExternalTrafficPolicy(v string) *UpdateK8sServiceRequest {
	s.ExternalTrafficPolicy = &v
	return s
}

func (s *UpdateK8sServiceRequest) SetName(v string) *UpdateK8sServiceRequest {
	s.Name = &v
	return s
}

func (s *UpdateK8sServiceRequest) SetServicePorts(v string) *UpdateK8sServiceRequest {
	s.ServicePorts = &v
	return s
}

func (s *UpdateK8sServiceRequest) SetType(v string) *UpdateK8sServiceRequest {
	s.Type = &v
	return s
}

func (s *UpdateK8sServiceRequest) Validate() error {
	return dara.Validate(s)
}
