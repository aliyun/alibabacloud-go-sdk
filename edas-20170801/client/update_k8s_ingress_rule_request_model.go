// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sIngressRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotations(v string) *UpdateK8sIngressRuleRequest
	GetAnnotations() *string
	SetClusterId(v string) *UpdateK8sIngressRuleRequest
	GetClusterId() *string
	SetIngressConf(v map[string]interface{}) *UpdateK8sIngressRuleRequest
	GetIngressConf() map[string]interface{}
	SetLabels(v string) *UpdateK8sIngressRuleRequest
	GetLabels() *string
	SetName(v string) *UpdateK8sIngressRuleRequest
	GetName() *string
	SetNamespace(v string) *UpdateK8sIngressRuleRequest
	GetNamespace() *string
}

type UpdateK8sIngressRuleRequest struct {
	// The annotations.
	//
	// example:
	//
	// {\\"nginx.ingress.kubernetes.io/ssl-redirect\\":\\"true\\",\\"nginx.ingress.kubernetes.io/configuration-snippet\\":\\"set $test value\\"}
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The ID of the Kubernetes cluster.
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
	// {\\"test-label\\":\\"test-label-value\\"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The name of the Ingress. The name can contain lowercase letters, digits, and hyphens (-). It must start with a lowercase letter but cannot end with a hyphen (-). The name can be up to 63 characters in length.
	//
	// example:
	//
	// my-ingress-rule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace of the Kubernetes cluster.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s UpdateK8sIngressRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sIngressRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sIngressRuleRequest) GetAnnotations() *string {
	return s.Annotations
}

func (s *UpdateK8sIngressRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateK8sIngressRuleRequest) GetIngressConf() map[string]interface{} {
	return s.IngressConf
}

func (s *UpdateK8sIngressRuleRequest) GetLabels() *string {
	return s.Labels
}

func (s *UpdateK8sIngressRuleRequest) GetName() *string {
	return s.Name
}

func (s *UpdateK8sIngressRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateK8sIngressRuleRequest) SetAnnotations(v string) *UpdateK8sIngressRuleRequest {
	s.Annotations = &v
	return s
}

func (s *UpdateK8sIngressRuleRequest) SetClusterId(v string) *UpdateK8sIngressRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateK8sIngressRuleRequest) SetIngressConf(v map[string]interface{}) *UpdateK8sIngressRuleRequest {
	s.IngressConf = v
	return s
}

func (s *UpdateK8sIngressRuleRequest) SetLabels(v string) *UpdateK8sIngressRuleRequest {
	s.Labels = &v
	return s
}

func (s *UpdateK8sIngressRuleRequest) SetName(v string) *UpdateK8sIngressRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateK8sIngressRuleRequest) SetNamespace(v string) *UpdateK8sIngressRuleRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateK8sIngressRuleRequest) Validate() error {
	return dara.Validate(s)
}
