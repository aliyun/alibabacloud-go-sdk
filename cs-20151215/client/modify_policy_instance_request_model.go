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
	// The action of the policy. Valid values:
	//
	// 	- `deny`: Deployments that match the policy are denied.
	//
	// 	- `warn`: Alerts are generated for deployments that match the policy.
	//
	// example:
	//
	// deny
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The ID of the policy instance.
	//
	// example:
	//
	// allowed-repos-cbhhb
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name,omitempty"`
	// The namespaces to which the policy is applied. The policy is applied to all namespaces if this parameter is left empty.
	Namespaces []*string `json:"namespaces,omitempty" xml:"namespaces,omitempty" type:"Repeated"`
	// The parameters of the policy instance. For more information, see [Predefined security policies of ACK](https://help.aliyun.com/document_detail/359819.html).
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
