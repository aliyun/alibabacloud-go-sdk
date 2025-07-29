// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DescribePolicyInstancesRequest
	GetInstanceName() *string
	SetPolicyName(v string) *DescribePolicyInstancesRequest
	GetPolicyName() *string
}

type DescribePolicyInstancesRequest struct {
	// The name of the policy instance that you want to query.
	//
	// example:
	//
	// allowed-repos-cz4s2
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name,omitempty"`
	// The name of the policy that you want to query.
	//
	// example:
	//
	// ACKPSPCapabilities
	PolicyName *string `json:"policy_name,omitempty" xml:"policy_name,omitempty"`
}

func (s DescribePolicyInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribePolicyInstancesRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribePolicyInstancesRequest) SetInstanceName(v string) *DescribePolicyInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribePolicyInstancesRequest) SetPolicyName(v string) *DescribePolicyInstancesRequest {
	s.PolicyName = &v
	return s
}

func (s *DescribePolicyInstancesRequest) Validate() error {
	return dara.Validate(s)
}
