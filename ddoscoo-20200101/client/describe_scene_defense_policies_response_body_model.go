// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneDefensePoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicies(v []*DescribeSceneDefensePoliciesResponseBodyPolicies) *DescribeSceneDefensePoliciesResponseBody
	GetPolicies() []*DescribeSceneDefensePoliciesResponseBodyPolicies
	SetRequestId(v string) *DescribeSceneDefensePoliciesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSceneDefensePoliciesResponseBody
	GetSuccess() *bool
}

type DescribeSceneDefensePoliciesResponseBody struct {
	// An array that consists of the configurations of the scenario-specific custom policy.
	Policies []*DescribeSceneDefensePoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F65DF043-E0EB-4796-9467-23DDCDF92C1D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSceneDefensePoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneDefensePoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefensePoliciesResponseBody) GetPolicies() []*DescribeSceneDefensePoliciesResponseBodyPolicies {
	return s.Policies
}

func (s *DescribeSceneDefensePoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSceneDefensePoliciesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSceneDefensePoliciesResponseBody) SetPolicies(v []*DescribeSceneDefensePoliciesResponseBodyPolicies) *DescribeSceneDefensePoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBody) SetRequestId(v string) *DescribeSceneDefensePoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBody) SetSuccess(v bool) *DescribeSceneDefensePoliciesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBody) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSceneDefensePoliciesResponseBodyPolicies struct {
	// The execution status of the policy. Valid values:
	//
	// 	- **1**: not executed or execution completed
	//
	// 	- **0**: being executed
	//
	// 	- **-1**: execution failed
	//
	// example:
	//
	// 1
	Done *int32 `json:"Done,omitempty" xml:"Done,omitempty"`
	// The time at which the policy expires. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1586016000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// testpolicy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of objects that are protected by the policy.
	//
	// example:
	//
	// 1
	ObjectCount *int32 `json:"ObjectCount,omitempty" xml:"ObjectCount,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// 321a-fd31-df51-****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The running rules of the policy.
	RuntimePolicies []*DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies `json:"RuntimePolicies,omitempty" xml:"RuntimePolicies,omitempty" type:"Repeated"`
	// The time at which the policy takes effect. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1585670400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the policy. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: pending enabling
	//
	// 	- **2**: enabled
	//
	// 	- **3**: expired
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the template that is used to create the policy. Valid values:
	//
	// 	- **promotion**: the Important Activity template
	//
	// 	- **bypass**: the Forward All template
	//
	// example:
	//
	// promotion
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s DescribeSceneDefensePoliciesResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneDefensePoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) GetDone() *int32 {
	return s.Done
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) GetName() *string {
	return s.Name
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) GetObjectCount() *int32 {
	return s.ObjectCount
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) GetRuntimePolicies() []*DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies {
	return s.RuntimePolicies
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) GetTemplate() *string {
	return s.Template
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetDone(v int32) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.Done = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetEndTime(v int64) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.EndTime = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetName(v string) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.Name = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetObjectCount(v int32) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.ObjectCount = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetPolicyId(v string) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetRuntimePolicies(v []*DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.RuntimePolicies = v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetStartTime(v int64) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.StartTime = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetStatus(v int32) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.Status = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetTemplate(v string) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.Template = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) Validate() error {
	if s.RuntimePolicies != nil {
		for _, item := range s.RuntimePolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies struct {
	// The protection rule that is applied when the policy takes effect.
	//
	// If you set **PolicyType*	- to **1**, the value is **{"cc_rule_enable": false }**. The value indicates that the Frequency Control policy is disabled.
	//
	// If you set **PolicyType*	- to **2**, the value is **{"ai_rule_enable": 0}**. The value indicates that the Intelligent Protection policy is disabled.
	//
	// example:
	//
	// {"cc_rule_enable": false }
	NewValue *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	// The protection policy whose status is changed when the policy takes effect. Valid values:
	//
	// 	- **1**: indicates that the Frequency Control policy is changed.
	//
	// 	- **2**: indicates that the Intelligent Protection policy is changed.
	//
	// example:
	//
	// 1
	PolicyType *int32 `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The running status of the policy. Valid values:
	//
	// 	- **0**: The policy has not been issued or is restored.
	//
	// 	- **1**: The policy is pending.
	//
	// 	- **2**: The policy is being restored.
	//
	// 	- **3**: The policy takes effect.
	//
	// 	- **4**: The policy fails to take effect.
	//
	// 	- **5**:The policy fails to be restored.
	//
	// 	- **6**: The configurations of the protected objects for the policy does not exist because the configurations may be deleted.
	//
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The protection rule that is applied before the policy takes effect.
	//
	// If you set **PolicyType*	- to **1**, the value is **{"cc_rule_enable": true}**. The value indicates that the Frequency Control policy is enabled.
	//
	// If you set **PolicyType*	- to **2**, the value is **{"ai_rule_enable": 1}**. The value indicates that the Intelligent Protection policy is enabled.
	//
	// example:
	//
	// {"cc_rule_enable": true}
	OldValue *string `json:"oldValue,omitempty" xml:"oldValue,omitempty"`
}

func (s DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) GetNewValue() *string {
	return s.NewValue
}

func (s *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) GetPolicyType() *int32 {
	return s.PolicyType
}

func (s *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) GetOldValue() *string {
	return s.OldValue
}

func (s *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) SetNewValue(v string) *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies {
	s.NewValue = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) SetPolicyType(v int32) *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies {
	s.PolicyType = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) SetStatus(v int32) *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies {
	s.Status = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) SetOldValue(v string) *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies {
	s.OldValue = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) Validate() error {
	return dara.Validate(s)
}
