// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolicyInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolicyInstancesResponse
	GetStatusCode() *int32
	SetBody(v []*DescribePolicyInstancesResponseBody) *DescribePolicyInstancesResponse
	GetBody() []*DescribePolicyInstancesResponseBody
}

type DescribePolicyInstancesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*DescribePolicyInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s DescribePolicyInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolicyInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolicyInstancesResponse) GetBody() []*DescribePolicyInstancesResponseBody {
	return s.Body
}

func (s *DescribePolicyInstancesResponse) SetHeaders(v map[string]*string) *DescribePolicyInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyInstancesResponse) SetStatusCode(v int32) *DescribePolicyInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyInstancesResponse) SetBody(v []*DescribePolicyInstancesResponseBody) *DescribePolicyInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribePolicyInstancesResponse) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyInstancesResponseBody struct {
	// The UID of the Alibaba Cloud account that is used to deploy the policy instance.
	//
	// example:
	//
	// 16298168****
	AliUid *string `json:"ali_uid,omitempty" xml:"ali_uid,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// c8155823d057948c69a****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The name of the policy instance.
	//
	// example:
	//
	// no-env-var-secrets-****
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// ACKPSPCapabilities
	PolicyName *string `json:"policy_name,omitempty" xml:"policy_name,omitempty"`
	// The type of the policy.
	//
	// example:
	//
	// k8s-general
	PolicyCategory *string `json:"policy_category,omitempty" xml:"policy_category,omitempty"`
	// The description of the policy template.
	//
	// example:
	//
	// Restricts secrets used in pod envs
	PolicyDescription *string `json:"policy_description,omitempty" xml:"policy_description,omitempty"`
	// The parameters of the policy instance.
	//
	// example:
	//
	// "restrictedNamespaces": [ "test" ]
	PolicyParameters *string `json:"policy_parameters,omitempty" xml:"policy_parameters,omitempty"`
	// The severity level of the policy instance.
	//
	// example:
	//
	// low
	PolicySeverity *string `json:"policy_severity,omitempty" xml:"policy_severity,omitempty"`
	// The applicable scope of the policy instance.
	//
	// A value of \\	- indicates all namespaces in the cluster. This is the default value.
	//
	// Multiple namespaces are separated by commas (,).
	//
	// example:
	//
	// *
	PolicyScope *string `json:"policy_scope,omitempty" xml:"policy_scope,omitempty"`
	// The action of the policy. Valid values:
	//
	// 	- `deny`: Deployments that match the policy are denied.
	//
	// 	- `warn`: Alerts are generated for deployments that match the policy.
	//
	// example:
	//
	// deny
	PolicyAction *string `json:"policy_action,omitempty" xml:"policy_action,omitempty"`
	// Deprecated
	//
	// The creation time of the instance. This parameter is deprecated.
	//
	// example:
	//
	// 2024-10-29T18:09:12+08:00
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	// Deprecated
	//
	// The update time of the instance. This parameter is deprecated.
	//
	// example:
	//
	// 2024-10-29T18:09:12+08:00
	Updated *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
	// Deprecated
	//
	// The ID of the resource. This parameter is deprecated.
	//
	// example:
	//
	// 123456***
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// Deprecated
	//
	// The number of violations processed in the cluster. This parameter is deprecated.
	//
	// example:
	//
	// 0
	TotalViolations *int64 `json:"total_violations,omitempty" xml:"total_violations,omitempty"`
	// Deprecated
	//
	// The status of the deletion. This parameter is deprecated.
	//
	// example:
	//
	// 0
	IsDeleted *int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
}

func (s DescribePolicyInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesResponseBody) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribePolicyInstancesResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribePolicyInstancesResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribePolicyInstancesResponseBody) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribePolicyInstancesResponseBody) GetPolicyCategory() *string {
	return s.PolicyCategory
}

func (s *DescribePolicyInstancesResponseBody) GetPolicyDescription() *string {
	return s.PolicyDescription
}

func (s *DescribePolicyInstancesResponseBody) GetPolicyParameters() *string {
	return s.PolicyParameters
}

func (s *DescribePolicyInstancesResponseBody) GetPolicySeverity() *string {
	return s.PolicySeverity
}

func (s *DescribePolicyInstancesResponseBody) GetPolicyScope() *string {
	return s.PolicyScope
}

func (s *DescribePolicyInstancesResponseBody) GetPolicyAction() *string {
	return s.PolicyAction
}

func (s *DescribePolicyInstancesResponseBody) GetCreated() *string {
	return s.Created
}

func (s *DescribePolicyInstancesResponseBody) GetUpdated() *string {
	return s.Updated
}

func (s *DescribePolicyInstancesResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribePolicyInstancesResponseBody) GetTotalViolations() *int64 {
	return s.TotalViolations
}

func (s *DescribePolicyInstancesResponseBody) GetIsDeleted() *int64 {
	return s.IsDeleted
}

func (s *DescribePolicyInstancesResponseBody) SetAliUid(v string) *DescribePolicyInstancesResponseBody {
	s.AliUid = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetClusterId(v string) *DescribePolicyInstancesResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetInstanceName(v string) *DescribePolicyInstancesResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetPolicyName(v string) *DescribePolicyInstancesResponseBody {
	s.PolicyName = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetPolicyCategory(v string) *DescribePolicyInstancesResponseBody {
	s.PolicyCategory = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetPolicyDescription(v string) *DescribePolicyInstancesResponseBody {
	s.PolicyDescription = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetPolicyParameters(v string) *DescribePolicyInstancesResponseBody {
	s.PolicyParameters = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetPolicySeverity(v string) *DescribePolicyInstancesResponseBody {
	s.PolicySeverity = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetPolicyScope(v string) *DescribePolicyInstancesResponseBody {
	s.PolicyScope = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetPolicyAction(v string) *DescribePolicyInstancesResponseBody {
	s.PolicyAction = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetCreated(v string) *DescribePolicyInstancesResponseBody {
	s.Created = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetUpdated(v string) *DescribePolicyInstancesResponseBody {
	s.Updated = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetResourceId(v string) *DescribePolicyInstancesResponseBody {
	s.ResourceId = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetTotalViolations(v int64) *DescribePolicyInstancesResponseBody {
	s.TotalViolations = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) SetIsDeleted(v int64) *DescribePolicyInstancesResponseBody {
	s.IsDeleted = &v
	return s
}

func (s *DescribePolicyInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
