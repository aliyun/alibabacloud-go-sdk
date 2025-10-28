// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sIngressRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteK8sIngressRuleRequest
	GetClusterId() *string
	SetName(v string) *DeleteK8sIngressRuleRequest
	GetName() *string
	SetNamespace(v string) *DeleteK8sIngressRuleRequest
	GetNamespace() *string
}

type DeleteK8sIngressRuleRequest struct {
	// The ID of the Kubernetes cluster.
	//
	// example:
	//
	// 5b2b4ab4-efbc-4a81-9c45-xxxxxxxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

func (s DeleteK8sIngressRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sIngressRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteK8sIngressRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteK8sIngressRuleRequest) GetName() *string {
	return s.Name
}

func (s *DeleteK8sIngressRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteK8sIngressRuleRequest) SetClusterId(v string) *DeleteK8sIngressRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteK8sIngressRuleRequest) SetName(v string) *DeleteK8sIngressRuleRequest {
	s.Name = &v
	return s
}

func (s *DeleteK8sIngressRuleRequest) SetNamespace(v string) *DeleteK8sIngressRuleRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteK8sIngressRuleRequest) Validate() error {
	return dara.Validate(s)
}
