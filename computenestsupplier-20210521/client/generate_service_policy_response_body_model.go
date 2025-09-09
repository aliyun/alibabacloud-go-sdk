// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateServicePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMissingPolicy(v []*GenerateServicePolicyResponseBodyMissingPolicy) *GenerateServicePolicyResponseBody
	GetMissingPolicy() []*GenerateServicePolicyResponseBodyMissingPolicy
	SetPolicy(v string) *GenerateServicePolicyResponseBody
	GetPolicy() *string
	SetRequestId(v string) *GenerateServicePolicyResponseBody
	GetRequestId() *string
}

type GenerateServicePolicyResponseBody struct {
	// The policies that are missing.
	MissingPolicy []*GenerateServicePolicyResponseBodyMissingPolicy `json:"MissingPolicy,omitempty" xml:"MissingPolicy,omitempty" type:"Repeated"`
	// The RAM policy.
	//
	// example:
	//
	// {Statement": [{ "Action": ["oos:*"], "Effect": "Allow", "Resource": "*"},{ "Action": ["ecs:DescribeInstances"], "Effect": "Allow", "Resource": "*"},{ "Action": ["ecs:RunInstance"], "Effect": "Allow", "Resource": "*"}], "Version": "1"}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5040BE9E-8DA2-5C9D-9B70-0EE6027A14BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateServicePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateServicePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateServicePolicyResponseBody) GetMissingPolicy() []*GenerateServicePolicyResponseBodyMissingPolicy {
	return s.MissingPolicy
}

func (s *GenerateServicePolicyResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *GenerateServicePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateServicePolicyResponseBody) SetMissingPolicy(v []*GenerateServicePolicyResponseBodyMissingPolicy) *GenerateServicePolicyResponseBody {
	s.MissingPolicy = v
	return s
}

func (s *GenerateServicePolicyResponseBody) SetPolicy(v string) *GenerateServicePolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GenerateServicePolicyResponseBody) SetRequestId(v string) *GenerateServicePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateServicePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type GenerateServicePolicyResponseBodyMissingPolicy struct {
	// Operations on specific resources.
	Action []*string `json:"Action,omitempty" xml:"Action,omitempty" type:"Repeated"`
	// The specific objects authorized. An asterisk (*) denotes all resources.
	//
	// example:
	//
	// *
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// ecs
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GenerateServicePolicyResponseBodyMissingPolicy) String() string {
	return dara.Prettify(s)
}

func (s GenerateServicePolicyResponseBodyMissingPolicy) GoString() string {
	return s.String()
}

func (s *GenerateServicePolicyResponseBodyMissingPolicy) GetAction() []*string {
	return s.Action
}

func (s *GenerateServicePolicyResponseBodyMissingPolicy) GetResource() *string {
	return s.Resource
}

func (s *GenerateServicePolicyResponseBodyMissingPolicy) GetServiceName() *string {
	return s.ServiceName
}

func (s *GenerateServicePolicyResponseBodyMissingPolicy) SetAction(v []*string) *GenerateServicePolicyResponseBodyMissingPolicy {
	s.Action = v
	return s
}

func (s *GenerateServicePolicyResponseBodyMissingPolicy) SetResource(v string) *GenerateServicePolicyResponseBodyMissingPolicy {
	s.Resource = &v
	return s
}

func (s *GenerateServicePolicyResponseBodyMissingPolicy) SetServiceName(v string) *GenerateServicePolicyResponseBodyMissingPolicy {
	s.ServiceName = &v
	return s
}

func (s *GenerateServicePolicyResponseBodyMissingPolicy) Validate() error {
	return dara.Validate(s)
}
