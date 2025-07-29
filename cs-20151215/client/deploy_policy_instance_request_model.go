// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployPolicyInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *DeployPolicyInstanceRequest
	GetAction() *string
	SetNamespaces(v []*string) *DeployPolicyInstanceRequest
	GetNamespaces() []*string
	SetParameters(v map[string]interface{}) *DeployPolicyInstanceRequest
	GetParameters() map[string]interface{}
}

type DeployPolicyInstanceRequest struct {
	// The action of the policy. Valid values:
	//
	// 	- `deny`: Deployments that match the policy are denied.
	//
	// 	- `warn`: Alerts are generated for Deployments that match the policy.
	//
	// example:
	//
	// deny
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The namespaces to which the policy applies. If you leave this parameter empty, the policy is applicable to all namespaces of the cluster.
	Namespaces []*string `json:"namespaces,omitempty" xml:"namespaces,omitempty" type:"Repeated"`
	// The parameter settings of the policy. For more information about the parameters supported by each policy, see [Predefined security policies of ACK](https://www.alibabacloud.com/help/doc-detail/359819.html).
	//
	// example:
	//
	// {"restrictedNamespaces": [ "test" ]}
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s DeployPolicyInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployPolicyInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeployPolicyInstanceRequest) GetAction() *string {
	return s.Action
}

func (s *DeployPolicyInstanceRequest) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *DeployPolicyInstanceRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *DeployPolicyInstanceRequest) SetAction(v string) *DeployPolicyInstanceRequest {
	s.Action = &v
	return s
}

func (s *DeployPolicyInstanceRequest) SetNamespaces(v []*string) *DeployPolicyInstanceRequest {
	s.Namespaces = v
	return s
}

func (s *DeployPolicyInstanceRequest) SetParameters(v map[string]interface{}) *DeployPolicyInstanceRequest {
	s.Parameters = v
	return s
}

func (s *DeployPolicyInstanceRequest) Validate() error {
	return dara.Validate(s)
}
