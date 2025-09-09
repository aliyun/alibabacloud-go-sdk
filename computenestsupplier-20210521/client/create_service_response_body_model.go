// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDryRunResult(v *CreateServiceResponseBodyDryRunResult) *CreateServiceResponseBody
	GetDryRunResult() *CreateServiceResponseBodyDryRunResult
	SetRequestId(v string) *CreateServiceResponseBody
	GetRequestId() *string
	SetServiceId(v string) *CreateServiceResponseBody
	GetServiceId() *string
	SetStatus(v string) *CreateServiceResponseBody
	GetStatus() *string
	SetVersion(v string) *CreateServiceResponseBody
	GetVersion() *string
}

type CreateServiceResponseBody struct {
	// The dry run result.
	DryRunResult *CreateServiceResponseBodyDryRunResult `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The status of the service.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The service version.
	//
	// example:
	//
	// draft
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) GetDryRunResult() *CreateServiceResponseBodyDryRunResult {
	return s.DryRunResult
}

func (s *CreateServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateServiceResponseBody) GetVersion() *string {
	return s.Version
}

func (s *CreateServiceResponseBody) SetDryRunResult(v *CreateServiceResponseBodyDryRunResult) *CreateServiceResponseBody {
	s.DryRunResult = v
	return s
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceResponseBody) SetServiceId(v string) *CreateServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceResponseBody) SetStatus(v string) *CreateServiceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateServiceResponseBody) SetVersion(v string) *CreateServiceResponseBody {
	s.Version = &v
	return s
}

func (s *CreateServiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateServiceResponseBodyDryRunResult struct {
	// The required ram policy for deploying role.
	RolePolicy *CreateServiceResponseBodyDryRunResultRolePolicy `json:"RolePolicy,omitempty" xml:"RolePolicy,omitempty" type:"Struct"`
}

func (s CreateServiceResponseBodyDryRunResult) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceResponseBodyDryRunResult) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBodyDryRunResult) GetRolePolicy() *CreateServiceResponseBodyDryRunResultRolePolicy {
	return s.RolePolicy
}

func (s *CreateServiceResponseBodyDryRunResult) SetRolePolicy(v *CreateServiceResponseBodyDryRunResultRolePolicy) *CreateServiceResponseBodyDryRunResult {
	s.RolePolicy = v
	return s
}

func (s *CreateServiceResponseBodyDryRunResult) Validate() error {
	return dara.Validate(s)
}

type CreateServiceResponseBodyDryRunResultRolePolicy struct {
	// The missing ram policy for deploying role.
	MissingPolicy []*CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy `json:"MissingPolicy,omitempty" xml:"MissingPolicy,omitempty" type:"Repeated"`
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

func (s CreateServiceResponseBodyDryRunResultRolePolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceResponseBodyDryRunResultRolePolicy) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicy) GetMissingPolicy() []*CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy {
	return s.MissingPolicy
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicy) GetPolicy() *string {
	return s.Policy
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicy) SetMissingPolicy(v []*CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) *CreateServiceResponseBodyDryRunResultRolePolicy {
	s.MissingPolicy = v
	return s
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicy) SetPolicy(v string) *CreateServiceResponseBodyDryRunResultRolePolicy {
	s.Policy = &v
	return s
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicy) Validate() error {
	return dara.Validate(s)
}

type CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy struct {
	// The Actions.
	Action []*string `json:"Action,omitempty" xml:"Action,omitempty" type:"Repeated"`
	// Resource in ram policy.
	//
	// example:
	//
	// *
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The service name in ram policy.
	//
	// example:
	//
	// ecs
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) GetAction() []*string {
	return s.Action
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) GetResource() *string {
	return s.Resource
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) SetAction(v []*string) *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy {
	s.Action = v
	return s
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) SetResource(v string) *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy {
	s.Resource = &v
	return s
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) SetServiceName(v string) *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy {
	s.ServiceName = &v
	return s
}

func (s *CreateServiceResponseBodyDryRunResultRolePolicyMissingPolicy) Validate() error {
	return dara.Validate(s)
}
