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
	// The governance action of the rule. Valid values:
	//
	// - `deny`: blocks non-compliant deployments.
	//
	// - `warn`: generates alerts.
	//
	// example:
	//
	// deny
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The namespaces to which the policy is restricted. An empty value indicates all namespaces.
	Namespaces []*string `json:"namespaces,omitempty" xml:"namespaces,omitempty" type:"Repeated"`
	// The parameter settings of the current rule instance. For the parameters supported by each policy governance rule and the corresponding metric description, see [Security policy rule library](https://www.alibabacloud.com/help/doc-detail/359819.html).
	//
	// example:
	//
	// {   "repos": [     "registry-vpc.cn-hangzhou.aliyuncs.com/acs/",     "registry.cn-hangzhou.aliyuncs.com/acs/"   ] }
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
