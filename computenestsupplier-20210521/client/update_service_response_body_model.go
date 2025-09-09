// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDryRunResult(v *UpdateServiceResponseBodyDryRunResult) *UpdateServiceResponseBody
	GetDryRunResult() *UpdateServiceResponseBodyDryRunResult
	SetRequestId(v string) *UpdateServiceResponseBody
	GetRequestId() *string
}

type UpdateServiceResponseBody struct {
	// The dry run result.
	DryRunResult *UpdateServiceResponseBodyDryRunResult `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty" type:"Struct"`
	// The hosted O\\&M configurations.
	//
	// example:
	//
	// DF0F666F-FBBC-55C3-A368-C955DE7B4839
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBody) GetDryRunResult() *UpdateServiceResponseBodyDryRunResult {
	return s.DryRunResult
}

func (s *UpdateServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceResponseBody) SetDryRunResult(v *UpdateServiceResponseBodyDryRunResult) *UpdateServiceResponseBody {
	s.DryRunResult = v
	return s
}

func (s *UpdateServiceResponseBody) SetRequestId(v string) *UpdateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceResponseBodyDryRunResult struct {
	// The required ram policy for deploying role.
	RolePolicy *UpdateServiceResponseBodyDryRunResultRolePolicy `json:"RolePolicy,omitempty" xml:"RolePolicy,omitempty" type:"Struct"`
}

func (s UpdateServiceResponseBodyDryRunResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceResponseBodyDryRunResult) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBodyDryRunResult) GetRolePolicy() *UpdateServiceResponseBodyDryRunResultRolePolicy {
	return s.RolePolicy
}

func (s *UpdateServiceResponseBodyDryRunResult) SetRolePolicy(v *UpdateServiceResponseBodyDryRunResultRolePolicy) *UpdateServiceResponseBodyDryRunResult {
	s.RolePolicy = v
	return s
}

func (s *UpdateServiceResponseBodyDryRunResult) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceResponseBodyDryRunResultRolePolicy struct {
	// The missing  ram policy for deploying role.
	MissingPolicy []*UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy `json:"MissingPolicy,omitempty" xml:"MissingPolicy,omitempty" type:"Repeated"`
	// The required ram policy for deploying role.
	//
	// example:
	//
	// {
	//
	// 	"Statement": [{
	//
	// 		"Action": ["oos:CancelExecutions", "oos:DeleteExecutions", "oos:GetTemplate", "oos:ListExecutions", "oos:ListTemplates", "oos:NotifyExecution", "oos:StartExecution"],
	//
	// 		"Effect": "Allow",
	//
	// 		"Resource": "*"
	//
	// 	}, {
	//
	// 		"Action": ["ram:PassRole"],
	//
	// 		"Effect": "Allow",
	//
	// 		"Resource": "*"
	//
	// 	}, {
	//
	// 		"Action": ["ros:CreateStack", "ros:GetStack", "ros:UpdateStack", "ros:ListStackEvents", "ros:ListStackResources", "ros:ListStackResources", "ros:DeleteStack", "ram:GetRole"],
	//
	// 		"Effect": "Allow",
	//
	// 		"Resource": "*"
	//
	// 	}],
	//
	// 	"Version": "1"
	//
	// }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s UpdateServiceResponseBodyDryRunResultRolePolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceResponseBodyDryRunResultRolePolicy) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicy) GetMissingPolicy() []*UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy {
	return s.MissingPolicy
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicy) GetPolicy() *string {
	return s.Policy
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicy) SetMissingPolicy(v []*UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) *UpdateServiceResponseBodyDryRunResultRolePolicy {
	s.MissingPolicy = v
	return s
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicy) SetPolicy(v string) *UpdateServiceResponseBodyDryRunResultRolePolicy {
	s.Policy = &v
	return s
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicy) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy struct {
	// The Actions.
	Action []*string `json:"Action,omitempty" xml:"Action,omitempty" type:"Repeated"`
	// The responses.
	//
	// example:
	//
	// *
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The service name.
	//
	// example:
	//
	// ecs
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) GetAction() []*string {
	return s.Action
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) GetResource() *string {
	return s.Resource
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) GetServiceName() *string {
	return s.ServiceName
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) SetAction(v []*string) *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy {
	s.Action = v
	return s
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) SetResource(v string) *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy {
	s.Resource = &v
	return s
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) SetServiceName(v string) *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy {
	s.ServiceName = &v
	return s
}

func (s *UpdateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) Validate() error {
	return dara.Validate(s)
}
