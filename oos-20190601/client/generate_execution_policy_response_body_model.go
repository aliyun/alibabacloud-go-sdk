// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateExecutionPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMissingPolicy(v string) *GenerateExecutionPolicyResponseBody
	GetMissingPolicy() *string
	SetPolicy(v string) *GenerateExecutionPolicyResponseBody
	GetPolicy() *string
	SetRequestId(v string) *GenerateExecutionPolicyResponseBody
	GetRequestId() *string
}

type GenerateExecutionPolicyResponseBody struct {
	// The policies that are missing.
	//
	// example:
	//
	// [{\\"Action\\": [\\"ecs:DescribeInvocationResults\\", \\"ecs:DescribeInstances\\", \\"ecs:RunCommand\\", \\"ecs:DescribeInvocations\\"], \\"ServiceName\\": \\"ecs\\", \\"Resources\\": \\"*\\"}]
	MissingPolicy *string `json:"MissingPolicy,omitempty" xml:"MissingPolicy,omitempty"`
	// The RAM policy.
	//
	// example:
	//
	// {}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateExecutionPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateExecutionPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateExecutionPolicyResponseBody) GetMissingPolicy() *string {
	return s.MissingPolicy
}

func (s *GenerateExecutionPolicyResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *GenerateExecutionPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateExecutionPolicyResponseBody) SetMissingPolicy(v string) *GenerateExecutionPolicyResponseBody {
	s.MissingPolicy = &v
	return s
}

func (s *GenerateExecutionPolicyResponseBody) SetPolicy(v string) *GenerateExecutionPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GenerateExecutionPolicyResponseBody) SetRequestId(v string) *GenerateExecutionPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateExecutionPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
