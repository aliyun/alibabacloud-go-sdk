// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateK8sIngressRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotations(v string) *CreateK8sIngressRuleRequest
	GetAnnotations() *string
	SetClusterId(v string) *CreateK8sIngressRuleRequest
	GetClusterId() *string
	SetIngressConf(v map[string]interface{}) *CreateK8sIngressRuleRequest
	GetIngressConf() map[string]interface{}
	SetLabels(v string) *CreateK8sIngressRuleRequest
	GetLabels() *string
	SetName(v string) *CreateK8sIngressRuleRequest
	GetName() *string
	SetNamespace(v string) *CreateK8sIngressRuleRequest
	GetNamespace() *string
}

type CreateK8sIngressRuleRequest struct {
	// The annotations.
	//
	// example:
	//
	// {\\"alb.ingress.kubernetes.io/rewrite-target\\":\\"/consumer-echo/test\\"}
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The ID of the Kubernetes cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5b2b4ab4-efbc-4a81-9c45-xxxxxxxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The routing rules of the Ingress. Set this parameter to a JSON string in the following format:
	//
	//     {
	//
	//       "rules": [
	//
	//         {
	//
	//           "host": "abc.com",
	//
	//           "secretName": "tls-secret",
	//
	//           "paths": [
	//
	//             {
	//
	//               "path": "/path",
	//
	//               "backend": {
	//
	//                 "servicePort": 80,
	//
	//                 "serviceName": "xxx"
	//
	//               }
	//
	//             }
	//
	//           ]
	//
	//         }
	//
	//       ]
	//
	//     }
	//
	// Parameter description:
	//
	// 	- rules: the list of routing rules.
	//
	// 	- host: the domain name to be accessed.
	//
	// 	- secretName: the name of the Secret that stores the information about the Transport Layer Security (TLS) certificate. The certificate is required if you need to use the HTTPS protocol.
	//
	// 	- paths: the list of paths to be accessed.
	//
	// 	- path: the path to be accessed.
	//
	// 	- backend: the configuration of the backend service. You can specify a service that is created in the Enterprise Distributed Application Service (EDAS) console.
	//
	// 	- serviceName: the name of the backend service.
	//
	// 	- servicePort: the port of the backend service.
	//
	// example:
	//
	// {"rules":[{"host":"abc.com","secretName":"tls-secret","paths":[{"path":"/path","backend":{"servicePort":80,"serviceName":"xxx"}}]}]}
	IngressConf map[string]interface{} `json:"IngressConf,omitempty" xml:"IngressConf,omitempty"`
	// The labels.
	//
	// example:
	//
	// {\\"test-labels\\":\\"test-value\\"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The name of the Ingress. The name can contain lowercase letters, digits, and hyphens (-). It must start with a lowercase letter but cannot end with a hyphen (-). The name can be up to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-ingress-rule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace of the Kubernetes cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s CreateK8sIngressRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateK8sIngressRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateK8sIngressRuleRequest) GetAnnotations() *string {
	return s.Annotations
}

func (s *CreateK8sIngressRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateK8sIngressRuleRequest) GetIngressConf() map[string]interface{} {
	return s.IngressConf
}

func (s *CreateK8sIngressRuleRequest) GetLabels() *string {
	return s.Labels
}

func (s *CreateK8sIngressRuleRequest) GetName() *string {
	return s.Name
}

func (s *CreateK8sIngressRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateK8sIngressRuleRequest) SetAnnotations(v string) *CreateK8sIngressRuleRequest {
	s.Annotations = &v
	return s
}

func (s *CreateK8sIngressRuleRequest) SetClusterId(v string) *CreateK8sIngressRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateK8sIngressRuleRequest) SetIngressConf(v map[string]interface{}) *CreateK8sIngressRuleRequest {
	s.IngressConf = v
	return s
}

func (s *CreateK8sIngressRuleRequest) SetLabels(v string) *CreateK8sIngressRuleRequest {
	s.Labels = &v
	return s
}

func (s *CreateK8sIngressRuleRequest) SetName(v string) *CreateK8sIngressRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateK8sIngressRuleRequest) SetNamespace(v string) *CreateK8sIngressRuleRequest {
	s.Namespace = &v
	return s
}

func (s *CreateK8sIngressRuleRequest) Validate() error {
	return dara.Validate(s)
}
