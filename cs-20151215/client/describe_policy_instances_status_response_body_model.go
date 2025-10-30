// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyInstancesStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstancesSeverityCount(v map[string]interface{}) *DescribePolicyInstancesStatusResponseBody
	GetInstancesSeverityCount() map[string]interface{}
	SetPolicyInstances(v []*DescribePolicyInstancesStatusResponseBodyPolicyInstances) *DescribePolicyInstancesStatusResponseBody
	GetPolicyInstances() []*DescribePolicyInstancesStatusResponseBodyPolicyInstances
}

type DescribePolicyInstancesStatusResponseBody struct {
	// The number of policy instances that are deployed in the cluster at different severity levels.
	//
	// example:
	//
	// { "high": 11,     "medium": 1  }
	InstancesSeverityCount map[string]interface{} `json:"instances_severity_count,omitempty" xml:"instances_severity_count,omitempty"`
	// The number of policy instances of each policy type.
	PolicyInstances []*DescribePolicyInstancesStatusResponseBodyPolicyInstances `json:"policy_instances,omitempty" xml:"policy_instances,omitempty" type:"Repeated"`
}

func (s DescribePolicyInstancesStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyInstancesStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesStatusResponseBody) GetInstancesSeverityCount() map[string]interface{} {
	return s.InstancesSeverityCount
}

func (s *DescribePolicyInstancesStatusResponseBody) GetPolicyInstances() []*DescribePolicyInstancesStatusResponseBodyPolicyInstances {
	return s.PolicyInstances
}

func (s *DescribePolicyInstancesStatusResponseBody) SetInstancesSeverityCount(v map[string]interface{}) *DescribePolicyInstancesStatusResponseBody {
	s.InstancesSeverityCount = v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBody) SetPolicyInstances(v []*DescribePolicyInstancesStatusResponseBodyPolicyInstances) *DescribePolicyInstancesStatusResponseBody {
	s.PolicyInstances = v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBody) Validate() error {
	if s.PolicyInstances != nil {
		for _, item := range s.PolicyInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolicyInstancesStatusResponseBodyPolicyInstances struct {
	// The type of the policy. For more information about different types of policies and their descriptions, see [Predefined security policies of ACK](https://help.aliyun.com/document_detail/359819.html).
	//
	// example:
	//
	// compliance
	PolicyCategory *string `json:"policy_category,omitempty" xml:"policy_category,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// Restricts use of the cluster-admin role.
	PolicyDescription *string `json:"policy_description,omitempty" xml:"policy_description,omitempty"`
	// The number of policy instances that are deployed. If this parameter is empty, no policy instance is deployed.
	//
	// example:
	//
	// 1
	PolicyInstancesCount *int64 `json:"policy_instances_count,omitempty" xml:"policy_instances_count,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// ACKRestrictRoleBindings
	PolicyName *string `json:"policy_name,omitempty" xml:"policy_name,omitempty"`
	// The severity level of the policy.
	//
	// example:
	//
	// medium
	PolicySeverity *string `json:"policy_severity,omitempty" xml:"policy_severity,omitempty"`
}

func (s DescribePolicyInstancesStatusResponseBodyPolicyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyInstancesStatusResponseBodyPolicyInstances) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesStatusResponseBodyPolicyInstances) GetPolicyCategory() *string {
	return s.PolicyCategory
}

func (s *DescribePolicyInstancesStatusResponseBodyPolicyInstances) GetPolicyDescription() *string {
	return s.PolicyDescription
}

func (s *DescribePolicyInstancesStatusResponseBodyPolicyInstances) GetPolicyInstancesCount() *int64 {
	return s.PolicyInstancesCount
}

func (s *DescribePolicyInstancesStatusResponseBodyPolicyInstances) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribePolicyInstancesStatusResponseBodyPolicyInstances) GetPolicySeverity() *string {
	return s.PolicySeverity
}

func (s *DescribePolicyInstancesStatusResponseBodyPolicyInstances) SetPolicyCategory(v string) *DescribePolicyInstancesStatusResponseBodyPolicyInstances {
	s.PolicyCategory = &v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBodyPolicyInstances) SetPolicyDescription(v string) *DescribePolicyInstancesStatusResponseBodyPolicyInstances {
	s.PolicyDescription = &v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBodyPolicyInstances) SetPolicyInstancesCount(v int64) *DescribePolicyInstancesStatusResponseBodyPolicyInstances {
	s.PolicyInstancesCount = &v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBodyPolicyInstances) SetPolicyName(v string) *DescribePolicyInstancesStatusResponseBodyPolicyInstances {
	s.PolicyName = &v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBodyPolicyInstances) SetPolicySeverity(v string) *DescribePolicyInstancesStatusResponseBodyPolicyInstances {
	s.PolicySeverity = &v
	return s
}

func (s *DescribePolicyInstancesStatusResponseBodyPolicyInstances) Validate() error {
	return dara.Validate(s)
}
