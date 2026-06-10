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
	// The governance action. Valid values:
	//
	// - `deny`: Denies deployments that violate the rule.
	//
	// - `warn`: Generates an alert for deployments that violate the rule.
	//
	// example:
	//
	// deny
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The namespaces where the policy applies. If you omit this parameter, the policy applies to all namespaces.
	Namespaces []*string `json:"namespaces,omitempty" xml:"namespaces,omitempty" type:"Repeated"`
	// For details on the parameters supported by each policy governance rule, see [Container security policy rules](https://www.alibabacloud.com/help/doc-detail/359819.html).
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
