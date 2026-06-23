// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *ModifyPolicyInstanceRequest
	GetAction() *string
	SetInstanceName(v string) *ModifyPolicyInstanceRequest
	GetInstanceName() *string
	SetNamespaces(v []*string) *ModifyPolicyInstanceRequest
	GetNamespaces() []*string
	SetParameters(v map[string]interface{}) *ModifyPolicyInstanceRequest
	GetParameters() map[string]interface{}
}

type ModifyPolicyInstanceRequest struct {
	// The governance action of the rule. Valid values:
	//
	// - `deny`: Blocks non-compliant deployments.
	//
	// - `warn`: Generates an alert.
	//
	// example:
	//
	// deny
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The instance ID of the policy rule.
	//
	// example:
	//
	// allowed-repos-cbhhb
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name,omitempty"`
	// The namespaces to which the policy applies. If this parameter is empty, the policy applies to all namespaces.
	Namespaces []*string `json:"namespaces,omitempty" xml:"namespaces,omitempty" type:"Repeated"`
	// The configuration parameters of the current rule instance. For more information about parameter settings rules, see [Container security policy rules](https://help.aliyun.com/document_detail/359819.html).
	//
	// example:
	//
	// "restrictedNamespaces": [ "test" ]
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s ModifyPolicyInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyInstanceRequest) GetAction() *string {
	return s.Action
}

func (s *ModifyPolicyInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyPolicyInstanceRequest) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *ModifyPolicyInstanceRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *ModifyPolicyInstanceRequest) SetAction(v string) *ModifyPolicyInstanceRequest {
	s.Action = &v
	return s
}

func (s *ModifyPolicyInstanceRequest) SetInstanceName(v string) *ModifyPolicyInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyPolicyInstanceRequest) SetNamespaces(v []*string) *ModifyPolicyInstanceRequest {
	s.Namespaces = v
	return s
}

func (s *ModifyPolicyInstanceRequest) SetParameters(v map[string]interface{}) *ModifyPolicyInstanceRequest {
	s.Parameters = v
	return s
}

func (s *ModifyPolicyInstanceRequest) Validate() error {
	return dara.Validate(s)
}
